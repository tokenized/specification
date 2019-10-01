package protocol

import (
	"bytes"
	"strconv"

	"github.com/dgryski/go-bitstream"
	"github.com/pkg/errors"
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

// FieldIndexPath is an "index path" to the field being specified. The index for each field is
//   specified in protobuf. Each item in the list is a level. The first item is the top level field
//   index and each item following is an index to sub-field within that field.
// Examples:
//   First field [ 0 ]
//   Third field [ 2 ]
//   Second field of the fifth top level field [ 4, 1 ]
type FieldIndexPath []uint32

// PermissionOf returns the permission object associated with a specified field index path.
func (ps Permissions) PermissionOf(fip FieldIndexPath) Permission {
	result := Permission{
		Permitted: true,
	}
	depth := 0

	for _, p := range ps {
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

// PermissionsFromBytes reads raw auth flag data into an array of permission structs.
func PermissionsFromBytes(b []byte, votingSystemCount int) (Permissions, error) {
	buf := bytes.NewBuffer(b)
	r := bitstream.NewReader(buf)

	count, err := ReadBase128VarInt(r)
	if err != nil {
		return nil, err
	}

	result := make(Permissions, count)
	for i, _ := range result {
		if err := result[i].Read(r, votingSystemCount); err != nil {
			return nil, err
		}
	}

	if buf.Len() > 0 {
		return nil, errors.New("Bytes remaining")
	}

	return result, nil
}

// Read reads a permission object from the buffer.
func (p *Permission) Read(r *bitstream.BitReader, votingSystemCount int) error {
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
	fieldCount, err := ReadBase128VarInt(r)
	if err != nil {
		return err
	}
	p.Fields = make([]FieldIndexPath, fieldCount)
	for i, _ := range p.Fields {
		indexCount, err := ReadBase128VarInt(r)
		if err != nil {
			return err
		}
		p.Fields[i] = make(FieldIndexPath, indexCount)
		for j, _ := range p.Fields[i] {
			p.Fields[i][j], err = ReadBase128VarInt(r)
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

	if err := WriteBase128VarInt(w, uint32(len(ps))); err != nil {
		return nil, err
	}

	for _, permission := range ps {
		if err := permission.Write(w); err != nil {
			return nil, err
		}
	}

	if err := w.Flush(bitstream.Zero); err != nil {
		return nil, errors.Wrap(err, "Failed to write permissions")
	}
	return buf.Bytes(), nil
}

// Write writes a permission object to the buffer.
func (p Permission) Write(w *bitstream.BitWriter) error {
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
	if err := WriteBase128VarInt(w, uint32(len(p.Fields))); err != nil {
		return err
	}
	for _, field := range p.Fields {
		if err := field.Write(w); err != nil {
			return err
		}
	}

	return nil
}

func (fip FieldIndexPath) MatchDepth(rfip FieldIndexPath) int {
	depth := 0
	for i, val := range fip {
		if i >= len(rfip) {
			break
		}
		if val != rfip[i] {
			break
		}
		depth++
	}
	return depth
}

func FieldIndexPathFromBytes(b []byte) (FieldIndexPath, error) {
	buf := bytes.NewBuffer(b)
	r := bitstream.NewReader(buf)
	return ReadFieldIndexPath(r)
}

func ReadFieldIndexPath(r *bitstream.BitReader) (FieldIndexPath, error) {
	count, err := ReadBase128VarInt(r)
	if err != nil {
		return nil, err
	}
	fip := make(FieldIndexPath, count)
	for i, _ := range fip {
		fip[i], err = ReadBase128VarInt(r)
		if err != nil {
			return nil, err
		}
	}
	return fip, nil
}

// Bytes writes a field index path into bytes.
func (fip FieldIndexPath) Bytes() ([]byte, error) {
	var buf bytes.Buffer
	w := bitstream.NewWriter(&buf)

	err := fip.Write(w)
	if err != nil {
		return nil, err
	}

	if err := w.Flush(bitstream.Zero); err != nil {
		return nil, errors.Wrap(err, "Failed to write field index path")
	}
	return buf.Bytes(), nil
}

func (fip FieldIndexPath) Write(w *bitstream.BitWriter) error {
	if err := WriteBase128VarInt(w, uint32(len(fip))); err != nil {
		return err
	}
	for _, index := range fip {
		if err := WriteBase128VarInt(w, index); err != nil {
			return err
		}
	}
	return nil
}

func (fip FieldIndexPath) String() string {
	result := ""
	first := true
	for _, index := range fip {
		if first {
			first = false
		} else {
			result += " / "
		}
		result += strconv.Itoa(int(index))
	}
	return result
}

func ReadBase128VarInt(r *bitstream.BitReader) (uint32, error) {
	value := uint32(0)
	done := false
	bitOffset := uint32(0)
	for !done {
		subValue, err := r.ReadByte()
		if err != nil {
			return value, err
		}

		done = (subValue & 0x80) == 0 // High bit not set
		subValue = subValue & 0x7f    // Remove high bit

		value += uint32(subValue) << bitOffset
		bitOffset += 7
	}

	return value, nil
}

func WriteBase128VarInt(w *bitstream.BitWriter, value uint32) error {
	for {
		if value < 128 {
			return w.WriteByte(byte(value))
		}
		subValue := (byte(value&0x7f) | 0x80) // Get last 7 bits and set high bit
		if err := w.WriteByte(subValue); err != nil {
			return err
		}
		value = value >> 7
	}
}
