package protocol

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
)

var (
	defaultEndian = binary.BigEndian
)

// pad returns a []byte with the given length.
func pad(b []byte, l int) []byte {
	if len(b) == l {
		return b
	}

	padding := []byte{}
	c := l - len(b)

	for i := 0; i < c; i++ {
		padding = append(padding, 0)
	}

	return append(b, padding...)
}

// lpad returns a []byte with the given length, padded on the left.
func lpad(b []byte, l int) []byte {
	if len(b) == l {
		return b
	}

	padding := []byte{}
	c := l - len(b)

	for i := 0; i < c; i++ {
		padding = append(padding, 0)
	}

	return append(padding, b...)
}

// write writes the  value to the buffer.
func write(buf io.Writer, v interface{}) error {
	return binary.Write(buf, defaultEndian, v)
}

// read fills the value with the appropriate number of bytes from the buffer.
//
// This is useful for fixed size types such as int, float etc.
func read(buf io.Reader, v interface{}) error {
	return binary.Read(buf, defaultEndian, v)
}

// readLen reads the number of bytes from the buffer to fill the slice of
// []byte.
func readLen(buf io.Reader, b []byte) error {
	_, err := io.ReadFull(buf, b)
	return err
}

func uintToBytes(i uint64) ([]byte, error) {
	l := 0

	if i <= math.MaxUint8 {
		l = 1
	} else if i <= math.MaxUint16 {
		l = 2
	} else if i < math.MaxUint32 {
		l = 4
	} else {
		l = 8
	}

	buf := new(bytes.Buffer)

	if err := binary.Write(buf, defaultEndian, i); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	if len(b) == l {
		return b, nil
	}

	// trim off the leading 0, because we have an 8 byte representation here,
	// and we want a representation of size l
	b = bytes.TrimLeftFunc(b, func(v rune) bool {
		return v == rune(0)
	})

	// pad the bytes out to the correct length
	return lpad(b, l), nil

}

func lenForOpPushdata(code byte) int {
	switch code {
	case 76:
		return 1
	case 77:
		return 2
	case 78:
		return 4
	}

	return 0
}
