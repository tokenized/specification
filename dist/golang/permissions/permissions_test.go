package permissions

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"

	bitstream "github.com/dgryski/go-bitstream"
)

func Test_Base128(t *testing.T) {
	tests := []struct {
		value uint32
		bits  []byte
	}{
		{
			value: 0,
			bits:  []byte{0},
		},
		{
			value: 1,
			bits:  []byte{1},
		},
		{
			value: 300,
			bits:  []byte{172, 2},
		},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(int(tt.value)), func(t *testing.T) {
			var buf bytes.Buffer
			w := bitstream.NewWriter(&buf)

			if err := writeBase128VarInt(w, tt.value); err != nil {
				t.Fatalf("Failed to write base 128 var int : %s", err)
			}

			b := buf.Bytes()

			if !bytes.Equal(b, tt.bits) {
				t.Fatalf("Wrong bits : got 0x%x, want 0x%x", b, tt.bits)
			}

			rbuf := bytes.NewBuffer(b)
			r := bitstream.NewReader(rbuf)

			result, err := readBase128VarInt(r)
			if err != nil {
				t.Fatalf("Failed to read base 128 var int : %s", err)
			}

			if result != tt.value {
				t.Fatalf("Wrong value returned : got %d, want %d", result, tt.value)
			}
		})
	}
}

func Test_SerializePermissions(t *testing.T) {

	tests := []Permissions{
		Permissions{
			Permission{
				Permitted:              true,
				AdministrationProposal: false,
				HolderProposal:         true,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{true, false},
				Fields: []FieldIndexPath{
					FieldIndexPath{1},
				},
			},
		},
		Permissions{
			Permission{
				Permitted:              true,
				AdministrationProposal: true,
				HolderProposal:         true,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{},
				Fields: []FieldIndexPath{
					FieldIndexPath{1},
				},
			},
			Permission{
				Permitted:              false,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   true,
				VotingSystemsAllowed:   []bool{},
				Fields: []FieldIndexPath{
					FieldIndexPath{2},
					FieldIndexPath{3, 4, 5},
					FieldIndexPath{800, 1},
				},
			},
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b, err := tt.Bytes()
			if err != nil {
				t.Fatalf("Failed to write permission : %s", err)
			}

			result, err := PermissionsFromBytes(b, len(tt[0].VotingSystemsAllowed))
			if err != nil {
				t.Fatalf("Failed to read permission : %s", err)
			}

			if !reflect.DeepEqual(result, tt) {
				t.Fatalf("Wrong result permission : \n got %+v\n want %+v", result, tt)
			}
		})
	}
}

func Test_FailSerializePermissions(t *testing.T) {

	tests := []Permissions{
		Permissions{
			Permission{
				Permitted:              true,
				AdministrationProposal: false,
				HolderProposal:         true,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{true, false},
				Fields: []FieldIndexPath{
					FieldIndexPath{1},
					FieldIndexPath{1, 2},
				},
			},
		},
		Permissions{
			Permission{
				Permitted:              true,
				AdministrationProposal: true,
				HolderProposal:         true,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{},
				Fields: []FieldIndexPath{
					FieldIndexPath{1},
				},
			},
			Permission{
				Permitted:              false,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   true,
				VotingSystemsAllowed:   []bool{},
				Fields: []FieldIndexPath{
					FieldIndexPath{2},
					FieldIndexPath{3, 4, 5},
					FieldIndexPath{800, 1},
					FieldIndexPath{1},
				},
			},
		},
		Permissions{
			Permission{
				Permitted:              true,
				AdministrationProposal: true,
				HolderProposal:         true,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{},
				Fields: []FieldIndexPath{
					FieldIndexPath{1},
				},
			},
			Permission{
				Permitted:              false,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   true,
				VotingSystemsAllowed:   []bool{},
				Fields: []FieldIndexPath{
					FieldIndexPath{2},
					FieldIndexPath{3, 4, 5},
					FieldIndexPath{800, 1},
					FieldIndexPath{1, 2},
				},
			},
		},
		Permissions{
			Permission{
				Permitted:              true,
				AdministrationProposal: true,
				HolderProposal:         true,
				AdministrativeMatter:   false,
				VotingSystemsAllowed:   []bool{},
				Fields: []FieldIndexPath{
					FieldIndexPath{1, 2},
				},
			},
			Permission{
				Permitted:              false,
				AdministrationProposal: false,
				HolderProposal:         false,
				AdministrativeMatter:   true,
				VotingSystemsAllowed:   []bool{},
				Fields: []FieldIndexPath{
					FieldIndexPath{2},
					FieldIndexPath{3, 4, 5},
					FieldIndexPath{800, 1},
					FieldIndexPath{1},
				},
			},
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b, err := tt.Bytes()
			if err != nil {
				t.Fatalf("Failed to write permission : %s", err)
			}

			_, err = PermissionsFromBytes(b, len(tt[0].VotingSystemsAllowed))
			if err == nil {
				t.Fatalf("Failed to reject invalid permission")
			}
		})
	}
}

func Test_SubPermissions(t *testing.T) {

	tests := []struct {
		name              string
		permissions       Permissions
		fip               FieldIndexPath
		operation         uint8
		isList            bool
		resultPermissions Permissions
	}{
		{
			name: "empty",
			permissions: Permissions{
				Permission{
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         true,
					AdministrativeMatter:   false,
					Fields:                 []FieldIndexPath{},
				},
			},
			fip:       FieldIndexPath{1},
			operation: 0,
			isList:    false,
			resultPermissions: Permissions{
				Permission{
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         true,
					AdministrativeMatter:   false,
					Fields:                 []FieldIndexPath{},
				},
			},
		},
		{
			name: "matching entry",
			permissions: Permissions{
				Permission{
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         true,
					AdministrativeMatter:   false,
					Fields: []FieldIndexPath{
						FieldIndexPath{3},
					},
				},
			},
			fip:       FieldIndexPath{3},
			operation: 0,
			isList:    false,
			resultPermissions: Permissions{
				Permission{
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         true,
					AdministrativeMatter:   false,
					Fields:                 []FieldIndexPath{},
				},
			},
		},
		{
			name: "deeper match",
			permissions: Permissions{
				Permission{
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         true,
					AdministrativeMatter:   false,
					Fields: []FieldIndexPath{
						FieldIndexPath{3, 1},
					},
				},
			},
			fip:       FieldIndexPath{3},
			operation: 0,
			isList:    false,
			resultPermissions: Permissions{
				Permission{
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         true,
					AdministrativeMatter:   false,
					Fields: []FieldIndexPath{
						FieldIndexPath{1},
					},
				},
			},
		},
		{
			name: "non-matching paths in match",
			permissions: Permissions{
				Permission{
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         true,
					AdministrativeMatter:   false,
					Fields: []FieldIndexPath{
						FieldIndexPath{3, 1},
						FieldIndexPath{4},
						FieldIndexPath{5, 3},
					},
				},
			},
			fip:       FieldIndexPath{3},
			operation: 0,
			isList:    false,
			resultPermissions: Permissions{
				Permission{
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         true,
					AdministrativeMatter:   false,
					Fields: []FieldIndexPath{
						FieldIndexPath{1},
					},
				},
			},
		},
		{
			name: "list match all",
			permissions: Permissions{
				Permission{
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         true,
					AdministrativeMatter:   false,
					Fields: []FieldIndexPath{
						FieldIndexPath{3, 0, 5}, // 0 means match all list indexes
					},
				},
			},
			fip:       FieldIndexPath{3, 2},
			operation: 0,
			isList:    true,
			resultPermissions: Permissions{
				Permission{
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         true,
					AdministrativeMatter:   false,
					Fields: []FieldIndexPath{
						FieldIndexPath{5},
					},
				},
			},
		},
		{
			name: "list match specific",
			permissions: Permissions{
				Permission{
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         true,
					AdministrativeMatter:   false,
					Fields: []FieldIndexPath{
						FieldIndexPath{3, 2, 5}, // 2 is a specific list index match
					},
				},
			},
			fip:       FieldIndexPath{3, 1}, // permission index minus one for "match all" zero offset
			operation: 0,
			isList:    true,
			resultPermissions: Permissions{
				Permission{
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         true,
					AdministrativeMatter:   false,
					Fields: []FieldIndexPath{
						FieldIndexPath{5},
					},
				},
			},
		},
		{
			name: "list non-match",
			permissions: Permissions{
				Permission{
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         true,
					AdministrativeMatter:   false,
					Fields: []FieldIndexPath{
						FieldIndexPath{3, 2, 5}, // 2 is a specific list index match
					},
				},
			},
			fip:       FieldIndexPath{3, 3},
			operation: 0,
			isList:    true,
			resultPermissions: Permissions{
				Permission{ // default match
					Permitted: true,
				},
			},
		},
		{
			name: "list default match",
			permissions: Permissions{
				Permission{
					Permitted:              false,
					AdministrationProposal: true,
					HolderProposal:         false,
					AdministrativeMatter:   false,
					Fields:                 []FieldIndexPath{},
				},
				Permission{
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         true,
					AdministrativeMatter:   false,
					Fields: []FieldIndexPath{
						FieldIndexPath{3, 2, 5}, // 2 is a specific list index match
					},
				},
			},
			fip:       FieldIndexPath{3, 3},
			operation: 0,
			isList:    true,
			resultPermissions: Permissions{
				Permission{
					Permitted:              false,
					AdministrationProposal: true,
					HolderProposal:         false,
					AdministrativeMatter:   false,
					Fields:                 []FieldIndexPath{},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultPermissions, err := tt.permissions.SubPermissions(tt.fip, tt.operation, tt.isList)
			if err != nil {
				t.Errorf("Failed to get sub permissions : %s", err)
				return
			}

			if !resultPermissions.Equal(tt.resultPermissions) {
				t.Errorf("Wrong result permissions : \ngot  %+v\nwant %+v", resultPermissions,
					tt.resultPermissions)
			}
		})
	}
}
