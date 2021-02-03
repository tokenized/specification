package permissions

import (
	"bytes"
	"strconv"

	"github.com/dgryski/go-bitstream"
	"github.com/pkg/errors"
)

// FieldIndexPath is an "index path" to the field being specified. The index for each field is
// specified. Each item in the list is a level. The first item is the top level field index and each
// item following is an index to sub-field within that field or an element with a list if it is a
// list.
//
// Permission field index paths:
// List indexes a zero means it applies to all elements of the list, the first element is referred
// to by 1 and so on.
//
// Amendment field index paths:
// List indexes in amendments are normal zero based indexes to items.
//
// Examples:
//   First field [ 0 ]
//   Third field [ 2 ]
//   Second field of the fifth top level field [ 4, 1 ]
//   All elements of the list that is the 4th field [ 5, 0 ]
//   The second item in the list that is the 5th field [ 6, 2 ]
type FieldIndexPath []uint32

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

func (fip FieldIndexPath) Equal(other FieldIndexPath) bool {
	if len(fip) != len(other) {
		return false
	}
	for i, index := range fip {
		if index != other[i] {
			return false
		}
	}

	return true
}

func FieldIndexPathFromBytes(b []byte) (FieldIndexPath, error) {
	buf := bytes.NewBuffer(b)
	r := bitstream.NewReader(buf)
	return readFieldIndexPath(r)
}

func readFieldIndexPath(r *bitstream.BitReader) (FieldIndexPath, error) {
	count, err := readBase128VarInt(r)
	if err != nil {
		return nil, err
	}
	fip := make(FieldIndexPath, count)
	for i, _ := range fip {
		fip[i], err = readBase128VarInt(r)
		if err != nil {
			return nil, err
		}
	}
	return fip, nil
}

func (fip FieldIndexPath) Copy() FieldIndexPath {
	result := make(FieldIndexPath, len(fip))
	copy(result, fip)
	return result
}

// Bytes writes a field index path into bytes.
func (fip FieldIndexPath) Bytes() ([]byte, error) {
	var buf bytes.Buffer
	w := bitstream.NewWriter(&buf)

	err := fip.write(w)
	if err != nil {
		return nil, err
	}

	if err := w.Flush(bitstream.Zero); err != nil {
		return nil, errors.Wrap(err, "Failed to write field index path")
	}
	return buf.Bytes(), nil
}

func (fip FieldIndexPath) write(w *bitstream.BitWriter) error {
	if err := writeBase128VarInt(w, uint32(len(fip))); err != nil {
		return err
	}
	for _, index := range fip {
		if err := writeBase128VarInt(w, index); err != nil {
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

func readBase128VarInt(r *bitstream.BitReader) (uint32, error) {
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

func writeBase128VarInt(w *bitstream.BitWriter, value uint32) error {
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
