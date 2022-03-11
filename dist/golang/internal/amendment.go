package internal

import "bytes"

// Amendment is an internal data structure used to communicate amendment data between actions and
// instruments, since instruments can't see the actions data structure.
type Amendment struct {
	FIP       []uint32
	Operation uint32
	Data      []byte
}

// Equal returns true if the objects are equal.
func (a Amendment) Equal(b Amendment) bool {
	if len(a.FIP) != len(b.FIP) {
		return false
	}

	for i := range a.FIP {
		if a.FIP[i] != b.FIP[i] {
			return false
		}
	}

	if a.Operation != b.Operation {
		return false
	}

	if !bytes.Equal(a.Data, b.Data) {
		return false
	}

	return true
}
