package protocol

import (
	"bytes"

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

// FieldIndexPath is an "index path" to the field being specified. The index for each field is
//   specified in protobuf. Each item in the list is a level. The first item is the top level field
//   index and each item following is an index to sub-field within that field.
// Examples:
//   First field [ 0 ]
//   Third field [ 2 ]
//   Second field of the fifth top level field [ 4, 1 ]
type FieldIndexPath []uint32

// ReadAuthFlags reads raw auth flag data into an array of permission structs.
func ReadAuthFlags(authFlags []byte, votingSystemCount int) ([]Permission, error) {
	buf := bytes.NewBuffer(authFlags)
	r := bitstream.NewReader(buf)

	count, err := ReadBase128VarInt(r)
	if err != nil {
		return nil, err
	}

	result := make([]Permission, count)
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

// Write writes a permission object to the buffer.
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

// WriteAuthFlags writes an array of permission structs into raw auth flag data.
func WriteAuthFlags(permissions []Permission) ([]byte, error) {
	var buf bytes.Buffer
	w := bitstream.NewWriter(&buf)

	// Verify all permission objects have the same number of voting systems.
	if len(permissions) > 0 {
		voteCount := len(permissions[0].VotingSystemsAllowed)
		for _, p := range permissions {
			if len(p.VotingSystemsAllowed) != voteCount {
				return nil, errors.New("Voting system counts inconsistent")
			}
		}
	}

	if err := WriteBase128VarInt(w, uint32(len(permissions))); err != nil {
		return nil, err
	}

	for _, permission := range permissions {
		if err := permission.Write(w); err != nil {
			return nil, err
		}
	}

	if err := w.Flush(bitstream.Zero); err != nil {
		return nil, errors.Wrap(err, "Failed to write auth flags")
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
		if err := WriteBase128VarInt(w, uint32(len(field))); err != nil {
			return err
		}
		for _, index := range field {
			if err := WriteBase128VarInt(w, index); err != nil {
				return err
			}
		}
	}

	return nil
}

func WriteBase128VarInt(w *bitstream.BitWriter, value uint32) error {
	for value != 0 {
		if value < 128 {
			return w.WriteByte(byte(value))
		}
		subValue := (byte(value&0x7f) | 0x80) // Get last 7 bits and set high bit
		if err := w.WriteByte(subValue); err != nil {
			return err
		}
		value = value >> 7
	}

	return nil
}
