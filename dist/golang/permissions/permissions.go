package permissions

import (
	"bytes"
	"fmt"

	"github.com/dgryski/go-bitstream"
	"github.com/pkg/errors"
)

var (
	ErrMissingElementIndex = errors.New("Missing element index")
)

// Permission represents the permissions assigned to a specific set of fields.
// If no fields are associated with a permission object, Fields is empty, then the permission object
//   is used as the default, and associated with all fields not specifically associated with another
//   permission object.
// If no permission object is associated with a field via the Fields, and no default permission
//   object is specified (with no fields), then the protocol default is applied, which is everything
//   is disabled except "Permitted". Meaning the administrator can change the field at will without
//   a vote.
type Permission struct {
	Permitted              bool // No Vote / Administration Amends
	AdministrationProposal bool // Administration Initiates / Members Vote / Administration Amend
	HolderProposal         bool // Members Initiate / Members Vote / Administration Amends
	AdministrativeMatter   bool // Administration Initiates / Administrators Vote / Administration Amends
	VotingSystemsAllowed   []bool
	Fields                 []FieldIndexPath // The fields that this permission applies to
}

// Permissions is a list of permission objects for a set of fields.
type Permissions []Permission

func (p Permission) Copy() Permission {
	result := p

	result.VotingSystemsAllowed = make([]bool, len(p.VotingSystemsAllowed))
	copy(result.VotingSystemsAllowed, p.VotingSystemsAllowed)

	result.Fields = make([]FieldIndexPath, len(p.Fields))
	for i, fip := range p.Fields {
		result.Fields[i] = make(FieldIndexPath, len(fip))
		copy(result.Fields[i], fip)
	}

	return result
}

func (ps Permissions) Copy() Permissions {
	result := make(Permissions, len(ps))
	for i, p := range ps {
		result[i] = p.Copy()
	}
	return result
}

func (p Permission) Equal(other Permission) bool {
	if p.Permitted != other.Permitted {
		return false
	}
	if p.AdministrationProposal != other.AdministrationProposal {
		return false
	}
	if p.HolderProposal != other.HolderProposal {
		return false
	}
	if p.AdministrativeMatter != other.AdministrativeMatter {
		return false
	}

	if len(p.VotingSystemsAllowed) != len(other.VotingSystemsAllowed) {
		return false
	}
	for i, vsa := range p.VotingSystemsAllowed {
		if vsa != other.VotingSystemsAllowed[i] {
			return false
		}
	}

	if len(p.Fields) != len(other.Fields) {
		return false
	}
	for i, fip := range p.Fields {
		if !fip.Equal(other.Fields[i]) {
			return false
		}
	}

	return true
}

func (ps Permissions) Equal(other Permissions) bool {
	if len(ps) != len(other) {
		return false
	}
	for i, p := range ps {
		if !p.Equal(other[i]) {
			return false
		}
	}

	return true
}

// PermissionOf returns the permission object associated with a specified field index path.
func (ps Permissions) PermissionOf(fip FieldIndexPath) Permission {
	result := Permission{
		Permitted: true,
	}
	depth := 0

	for _, p := range ps {
		// Find deepest path match specified in permissions.
		for _, pfip := range p.Fields {
			match := pfip.MatchDepth(fip)
			if match > depth {
				depth = match
				result = p
			}
		}
		if depth == 0 && len(p.Fields) == 0 {
			result = p
		}
	}

	return result
}

// SubPermissions returns all permissions related to the specified path and operation. isList
// specified if the current field is a list of items. It is used to narrow down permissions when
// working through an amendment.
// The permissions are returned in order with the first item the most specific to the field. The
// most specific to the field is the one with the closest path.
func (ps Permissions) SubPermissions(fip FieldIndexPath, operation uint32,
	isList bool) (Permissions, error) {

	var result Permissions
	for _, permission := range ps {
		pathMatch := false
		var match *Permission
		for _, pfip := range permission.Fields {
			if len(pfip) == 0 {
				continue
			}

			if pfip[0] == fip[0] {
				if isList {
					if operation == 0 || operation == 2 {
						// modify or delete (refers to specific element)
						if len(fip) < 2 {
							// No element index
							return nil, ErrMissingElementIndex
						}

						if len(pfip) != 1 && // permissions specify this deep
							pfip[1] != 0 && // permissions don't specify all elements
							pfip[1]-1 != fip[1] { // permissions don't specify this element
							continue // not a match
						}
					} else {
						// add (refers to new element)
						if len(pfip) != 1 && // permissions specify this deep
							pfip[1] != 0 { // permissions don't specify all elements
							continue // not a match
						}
					}
				}

				if match == nil {
					newPermission := permission.Copy()
					newPermission.Fields = nil
					match = &newPermission
					pathMatch = true
				}

				// Remove matching field index so it is lined up for subfields
				toRemove := 1
				if isList {
					toRemove++
				}
				if len(pfip) > toRemove {
					match.Fields = append(match.Fields, pfip[toRemove:])
				}
			}
		}

		if match == nil && len(permission.Fields) == 0 {
			// General match
			newPermission := permission.Copy()
			match = &newPermission
		}

		if match != nil {
			if pathMatch {
				result = append(Permissions{*match}, result...) // add to beginning
			} else {
				result = append(result, *match) // add to end
			}
		}
	}

	if len(result) == 0 {
		// No matches so fall back to default
		result = append(result, Permission{
			Permitted: true,
		})

	}

	return result, nil
}

// PermissionsFromBytes reads raw auth flag data into an array of permission structs.
func PermissionsFromBytes(b []byte, votingSystemCount int) (Permissions, error) {
	if len(b) == 0 {
		return Permissions{}, nil
	}

	buf := bytes.NewBuffer(b)
	r := bitstream.NewReader(buf)

	count, err := readBase128VarInt(r)
	if err != nil {
		return nil, err
	}

	result := make(Permissions, count)
	var fips []FieldIndexPath
	for i, _ := range result {
		if err := result[i].read(r, votingSystemCount); err != nil {
			return nil, err
		}

		// Check for duplicates
		for _, fieldIndexPath := range result[i].Fields {
			for _, fip := range fips {
				matches := true
				for i, index := range fieldIndexPath {
					if i >= len(fip) {
						if i > 0 {
							return nil, fmt.Errorf("Duplicate partial field index path %v",
								fieldIndexPath)
						}
						break
					}
					if index != fip[i] {
						matches = false
						break
					}
				}
				if matches {
					return nil, fmt.Errorf("Duplicate field index path %v", fieldIndexPath)
				}
			}
			fips = append(fips, fieldIndexPath)
		}
	}

	if buf.Len() > 0 {
		return nil, errors.New("Bytes remaining")
	}

	return result, nil
}

// Read reads a permission object from the buffer.
func (p *Permission) read(r *bitstream.BitReader, votingSystemCount int) error {
	// Permitted
	bit, err := r.ReadBit()
	if err != nil {
		return err
	}
	p.Permitted = bool(bit)

	// AdministrationProposal
	bit, err = r.ReadBit()
	if err != nil {
		return err
	}
	p.AdministrationProposal = bool(bit)

	// HolderProposal
	bit, err = r.ReadBit()
	if err != nil {
		return err
	}
	p.HolderProposal = bool(bit)

	// AdministrativeMatter
	bit, err = r.ReadBit()
	if err != nil {
		return err
	}
	p.AdministrativeMatter = bool(bit)

	// Voting Systems
	if p.AdministrationProposal || p.HolderProposal || p.AdministrativeMatter {
		p.VotingSystemsAllowed = make([]bool, votingSystemCount)
		for i, _ := range p.VotingSystemsAllowed {
			bit, err = r.ReadBit()
			if err != nil {
				return err
			}
			p.VotingSystemsAllowed[i] = bool(bit)
		}
	}

	// Field index paths (base 128 var ints)
	fieldCount, err := readBase128VarInt(r)
	if err != nil {
		return err
	}
	p.Fields = make([]FieldIndexPath, fieldCount)
	for i, _ := range p.Fields {
		indexCount, err := readBase128VarInt(r)
		if err != nil {
			return err
		}
		p.Fields[i] = make(FieldIndexPath, indexCount)
		for j, _ := range p.Fields[i] {
			p.Fields[i][j], err = readBase128VarInt(r)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Bytes writes an array of permission structs into raw auth flag data.
func (ps Permissions) Bytes() ([]byte, error) {
	var buf bytes.Buffer
	w := bitstream.NewWriter(&buf)

	// Verify all permission objects have the same number of voting systems.
	if len(ps) > 0 {
		voteCount := len(ps[0].VotingSystemsAllowed)
		for _, p := range ps {
			if len(p.VotingSystemsAllowed) != voteCount {
				return nil, errors.New("Voting system counts inconsistent")
			}
		}
	}

	if err := writeBase128VarInt(w, uint32(len(ps))); err != nil {
		return nil, err
	}

	for _, permission := range ps {
		if err := permission.write(w); err != nil {
			return nil, err
		}
	}

	if err := w.Flush(bitstream.Zero); err != nil {
		return nil, errors.Wrap(err, "Failed to write permissions")
	}
	return buf.Bytes(), nil
}

// Write writes a permission object to the buffer.
func (p Permission) write(w *bitstream.BitWriter) error {
	// Permitted
	err := w.WriteBit(bitstream.Bit(p.Permitted))
	if err != nil {
		return err
	}

	// AdministrationProposal
	err = w.WriteBit(bitstream.Bit(p.AdministrationProposal))
	if err != nil {
		return err
	}

	// HolderProposal
	err = w.WriteBit(bitstream.Bit(p.HolderProposal))
	if err != nil {
		return err
	}

	// AdministrativeMatter
	err = w.WriteBit(bitstream.Bit(p.AdministrativeMatter))
	if err != nil {
		return err
	}

	// Voting Systems
	if p.AdministrationProposal || p.HolderProposal || p.AdministrativeMatter {
		for _, votingSystem := range p.VotingSystemsAllowed {
			err = w.WriteBit(bitstream.Bit(votingSystem))
			if err != nil {
				return err
			}
		}
	}

	// Field index paths (base 128 var ints)
	if err := writeBase128VarInt(w, uint32(len(p.Fields))); err != nil {
		return err
	}
	for _, field := range p.Fields {
		if err := field.write(w); err != nil {
			return err
		}
	}

	return nil
}
