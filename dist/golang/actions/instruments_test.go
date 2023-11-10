package actions

import (
	"testing"

	"github.com/tokenized/pkg/bitcoin"
)

func Test_InstrumentDefinition_TransferFee_Validate(t *testing.T) {
	key, _ := bitcoin.GenerateKey(bitcoin.MainNet)
	ls, _ := key.LockingScript()
	ra, _ := bitcoin.RawAddressFromLockingScript(ls)

	tests := []struct {
		name       string
		definition *InstrumentDefinition
		valid      bool
	}{
		{
			name: "no fee",
			definition: &InstrumentDefinition{
				InstrumentType:    "CRN",
				InstrumentPayload: []byte{1, 2, 3},
			},
			valid: true,
		},
		{
			name: "missing address",
			definition: &InstrumentDefinition{
				InstrumentType:    "CRN",
				InstrumentPayload: []byte{1, 2, 3},
				TransferFee: &FeeField{
					Quantity: 1,
				},
			},
			valid: false,
		},
		{
			name: "includes address",
			definition: &InstrumentDefinition{
				InstrumentType:    "CRN",
				InstrumentPayload: []byte{1, 2, 3},
				TransferFee: &FeeField{
					Address:  ra.Bytes(),
					Quantity: 1,
				},
			},
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.definition.Validate()
			if tt.valid {
				if err != nil {
					t.Errorf("Failed validate : %s", err)
				} else {
					t.Logf("Valid")
				}
			} else {
				if err == nil {
					t.Errorf("No error on invalid")
				} else {
					t.Logf("Invalid : %s", err)
				}
			}
		})
	}
}

func Test_InstrumentCreation_TransferFee_Validate(t *testing.T) {
	key, _ := bitcoin.GenerateKey(bitcoin.MainNet)
	ls, _ := key.LockingScript()
	ra, _ := bitcoin.RawAddressFromLockingScript(ls)

	tests := []struct {
		name     string
		creation *InstrumentCreation
		valid    bool
	}{
		{
			name: "no fee",
			creation: &InstrumentCreation{
				InstrumentType:    "CRN",
				InstrumentPayload: []byte{1, 2, 3},
			},
			valid: true,
		},
		{
			name: "missing address",
			creation: &InstrumentCreation{
				InstrumentType:    "CRN",
				InstrumentPayload: []byte{1, 2, 3},
				TransferFee: &FeeField{
					Quantity: 1,
				},
			},
			valid: false,
		},
		{
			name: "includes address",
			creation: &InstrumentCreation{
				InstrumentType:    "CRN",
				InstrumentPayload: []byte{1, 2, 3},
				TransferFee: &FeeField{
					Address:  ra.Bytes(),
					Quantity: 1,
				},
			},
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.creation.Validate()
			if tt.valid {
				if err != nil {
					t.Errorf("Failed validate : %s", err)
				} else {
					t.Logf("Valid")
				}
			} else {
				if err == nil {
					t.Errorf("No error on invalid")
				} else {
					t.Logf("Invalid : %s", err)
				}
			}
		})
	}
}
