package actions

import (
	"encoding/json"
	"testing"

	"github.com/pkg/errors"
	"github.com/tokenized/pkg/bitcoin"
)

func TestContractCreateAmendments(t *testing.T) {
	key1, err := bitcoin.GenerateKey(bitcoin.MainNet)
	if err != nil {
		t.Fatalf("Failed to generate key : %s", err)
	}
	publicKey1 := key1.PublicKey()
	address1, err := key1.RawAddress()
	if err != nil {
		t.Fatalf("Failed to generate address : %s", err)
	}

	key2, err := bitcoin.GenerateKey(bitcoin.MainNet)
	if err != nil {
		t.Fatalf("Failed to generate key : %s", err)
	}
	publicKey2 := key2.PublicKey()

	tests := []struct {
		name       string
		current    *ContractFormation
		newValue   *ContractOffer
		err        error
		amendments []*AmendmentField
	}{
		{
			name: "Change name",
			current: &ContractFormation{
				ContractName:        "Original Name",
				MasterAddress:       address1.Bytes(),
				RestrictedQtyAssets: 1,
			},
			newValue: &ContractOffer{
				ContractName:        "New Name",
				MasterAddress:       address1.Bytes(),
				RestrictedQtyAssets: 1,
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{1, byte(ContractFieldContractName)},
					Operation:      0,
					Data:           []byte("New Name"),
				},
			},
		},
		{
			name: "Change restricted quantity",
			current: &ContractFormation{
				ContractName:        "Original Name",
				MasterAddress:       address1.Bytes(),
				RestrictedQtyAssets: 1,
			},
			newValue: &ContractOffer{
				ContractName:        "Original Name",
				MasterAddress:       address1.Bytes(),
				RestrictedQtyAssets: 2,
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{1, byte(ContractFieldRestrictedQtyAssets)},
					Operation:      0,
					Data:           []byte{2},
				},
			},
		},
		{
			name: "Change 2 fields",
			current: &ContractFormation{
				ContractName:        "Original Name",
				MasterAddress:       address1.Bytes(),
				RestrictedQtyAssets: 1,
			},
			newValue: &ContractOffer{
				ContractName:        "New Name",
				MasterAddress:       address1.Bytes(),
				RestrictedQtyAssets: 2,
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{1, byte(ContractFieldContractName)},
					Operation:      0,
					Data:           []byte("New Name"),
				},
				&AmendmentField{
					FieldIndexPath: []byte{1, byte(ContractFieldRestrictedQtyAssets)},
					Operation:      0,
					Data:           []byte{2},
				},
			},
		},
		{
			name: "Change service key",
			current: &ContractFormation{
				ContractName:        "Original Name",
				MasterAddress:       address1.Bytes(),
				RestrictedQtyAssets: 1,
				Services: []*ServiceField{
					&ServiceField{
						Type:      ServiceTypeIdentityOracle,
						URL:       "test.identity.tokenized.com",
						PublicKey: publicKey1.Bytes(),
					},
				},
			},
			newValue: &ContractOffer{
				ContractName:        "Original Name",
				MasterAddress:       address1.Bytes(),
				RestrictedQtyAssets: 1,
				Services: []*ServiceField{
					&ServiceField{
						Type:      ServiceTypeIdentityOracle,
						URL:       "test.identity.tokenized.com",
						PublicKey: publicKey2.Bytes(),
					},
				},
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{3, byte(ContractFieldServices), 0,
						byte(ServiceFieldPublicKey)},
					Operation: 0,
					Data:      publicKey2.Bytes(),
				},
			},
		},
		{
			name: "Remove last service",
			current: &ContractFormation{
				ContractName:        "Original Name",
				MasterAddress:       address1.Bytes(),
				RestrictedQtyAssets: 1,
				Services: []*ServiceField{
					&ServiceField{
						Type:      ServiceTypeIdentityOracle,
						URL:       "test.identity.tokenized.com",
						PublicKey: publicKey1.Bytes(),
					},
					&ServiceField{
						Type:      ServiceTypeAuthorityOracle,
						URL:       "test.authority.tokenized.com",
						PublicKey: publicKey2.Bytes(),
					},
				},
			},
			newValue: &ContractOffer{
				ContractName:        "Original Name",
				MasterAddress:       address1.Bytes(),
				RestrictedQtyAssets: 1,
				Services: []*ServiceField{
					&ServiceField{
						Type:      ServiceTypeIdentityOracle,
						URL:       "test.identity.tokenized.com",
						PublicKey: publicKey1.Bytes(),
					},
				},
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{2, byte(ContractFieldServices), 1},
					Operation:      2,
				},
			},
		},
		{
			// Note: Without making the difference checking too complicated, because of limited
			// template functionality, the difference is detected as modifying the first item, then
			// removing the second item.
			name: "Remove first service",
			current: &ContractFormation{
				ContractName:        "Original Name",
				MasterAddress:       address1.Bytes(),
				RestrictedQtyAssets: 1,
				Services: []*ServiceField{
					&ServiceField{
						Type:      ServiceTypeIdentityOracle,
						URL:       "test.identity.tokenized.com",
						PublicKey: publicKey1.Bytes(),
					},
					&ServiceField{
						Type:      ServiceTypeAuthorityOracle,
						URL:       "test.authority.tokenized.com",
						PublicKey: publicKey2.Bytes(),
					},
				},
			},
			newValue: &ContractOffer{
				ContractName:        "Original Name",
				MasterAddress:       address1.Bytes(),
				RestrictedQtyAssets: 1,
				Services: []*ServiceField{
					&ServiceField{
						Type:      ServiceTypeAuthorityOracle,
						URL:       "test.authority.tokenized.com",
						PublicKey: publicKey2.Bytes(),
					},
				},
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{3, byte(ContractFieldServices), 0,
						byte(ServiceFieldType)},
					Operation: 0,
					Data:      []byte{byte(ServiceTypeAuthorityOracle)},
				},
				&AmendmentField{
					FieldIndexPath: []byte{3, byte(ContractFieldServices), 0,
						byte(ServiceFieldURL)},
					Operation: 0,
					Data:      []byte("test.authority.tokenized.com"),
				},
				&AmendmentField{
					FieldIndexPath: []byte{3, byte(ContractFieldServices), 0,
						byte(ServiceFieldPublicKey)},
					Operation: 0,
					Data:      publicKey2.Bytes(),
				},
				&AmendmentField{ // remove second service
					FieldIndexPath: []byte{2, byte(ContractFieldServices), 1},
					Operation:      2,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amendments, err := tt.current.CreateAmendments(tt.newValue)
			if err != nil {
				if tt.err == nil {
					t.Errorf("Failed to create amendments : %s", err)
					return
				}

				if tt.err.Error() != err.Error() {
					t.Errorf("Wrong error : got %s, want %s", err, tt.err)
					return
				}

				return
			}

			if tt.err != nil {
				t.Errorf("Error not returned : want %s", tt.err)
				return
			}

			if len(amendments) != len(tt.amendments) {
				t.Errorf("Wrong amendment count : got %d, want %d\n%+v", len(amendments),
					len(tt.amendments), amendments)
				return
			}

			for i := range amendments {
				if !amendments[i].Equal(tt.amendments[i]) {
					t.Errorf("Wrong amendment %d : \n  got  %+v\n  want %+v", i, *amendments[i],
						*tt.amendments[i])
					return
				}
			}

			amended := &ContractFormation{}
			if err := convert(amended, tt.current); err != nil {
				t.Errorf("Failed to convert current value : %s", err)
				return
			}
			for i, amendment := range amendments {
				fip, err := FieldIndexPathFromBytes(amendment.FieldIndexPath)
				if err != nil {
					t.Errorf("Failed to parse FIP : %s", err)
					return
				}

				_, err = amended.ApplyAmendment(fip, amendment.Operation, amendment.Data)
				if err != nil {
					t.Errorf("Failed to apply amendment %d : %s", i, err)
					return
				}
			}

			newValue := &ContractFormation{}
			if err := convert(newValue, tt.newValue); err != nil {
				t.Errorf("Failed to convert new value : %s", err)
				return
			}
			if !amended.Equal(newValue) {
				t.Errorf("Amended value doesn't match : \n  got  %+v\n  want %+v", *amended,
					*newValue)
				return
			}
		})
	}
}

func TestAssetCreateAmendments(t *testing.T) {
	tests := []struct {
		name       string
		current    *AssetCreation
		newValue   *AssetDefinition
		err        error
		amendments []*AmendmentField
	}{
		{
			name: "Change transfers permitted",
			current: &AssetCreation{
				TransfersPermitted: false,
				TradeRestrictions:  []string{"AUS"},
				TokenQty:           10000,
				AssetType:          "CUR",
				AssetPayload:       []byte{1},
			},
			newValue: &AssetDefinition{
				TransfersPermitted: true,
				TradeRestrictions:  []string{"AUS"},
				TokenQty:           10000,
				AssetType:          "CUR",
				AssetPayload:       []byte{1},
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{1, byte(AssetFieldTransfersPermitted)},
					Operation:      0,
					Data:           []byte{1},
				},
			},
		},
		{
			name: "Change asset type",
			current: &AssetCreation{
				TransfersPermitted: false,
				TradeRestrictions:  []string{"AUS"},
				TokenQty:           10000,
				AssetType:          "CUR",
				AssetPayload:       []byte{1},
			},
			newValue: &AssetDefinition{
				TransfersPermitted: false,
				TradeRestrictions:  []string{"AUS"},
				TokenQty:           10000,
				AssetType:          "COU",
				AssetPayload:       []byte{1},
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{1, byte(AssetFieldAssetType)},
					Operation:      0,
					Data:           []byte("COU"),
				},
			},
		},
		{
			name: "Change token quantity",
			current: &AssetCreation{
				TransfersPermitted: false,
				TradeRestrictions:  []string{"AUS"},
				TokenQty:           10000,
				AssetType:          "CUR",
				AssetPayload:       []byte{1},
			},
			newValue: &AssetDefinition{
				TransfersPermitted: false,
				TradeRestrictions:  []string{"AUS"},
				TokenQty:           100000,
				AssetType:          "CUR",
				AssetPayload:       []byte{1},
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{1, byte(AssetFieldTokenQty)},
					Operation:      0,
					Data:           []byte{160, 141, 6}, // base 128 var int encoding of 100000
				},
			},
		},
		{
			name: "Change two fields",
			current: &AssetCreation{
				TransfersPermitted: false,
				TradeRestrictions:  []string{"AUS"},
				TokenQty:           10000,
				AssetType:          "CUR",
				AssetPayload:       []byte{1},
			},
			newValue: &AssetDefinition{
				TransfersPermitted: false,
				TradeRestrictions:  []string{"AUS"},
				TokenQty:           100000,
				AssetType:          "COU",
				AssetPayload:       []byte{1},
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{1, byte(AssetFieldTokenQty)},
					Operation:      0,
					Data:           []byte{160, 141, 6}, // base 128 var int encoding of 100000
				},
				&AmendmentField{
					FieldIndexPath: []byte{1, byte(AssetFieldAssetType)},
					Operation:      0,
					Data:           []byte("COU"),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amendments, err := tt.current.CreateAmendments(tt.newValue)
			if err != nil {
				if tt.err == nil {
					t.Errorf("Failed to create amendments : %s", err)
					return
				}

				if tt.err.Error() != err.Error() {
					t.Errorf("Wrong error : got %s, want %s", err, tt.err)
					return
				}

				return
			}

			if tt.err != nil {
				t.Errorf("Error not returned : want %s", tt.err)
				return
			}

			if len(amendments) != len(tt.amendments) {
				t.Errorf("Wrong amendment count : got %d, want %d\n%+v", len(amendments),
					len(tt.amendments), amendments)
				return
			}

			for i := range amendments {
				if !amendments[i].Equal(tt.amendments[i]) {
					t.Errorf("Wrong amendment %d : \n  got  %+v\n  want %+v", i, *amendments[i],
						*tt.amendments[i])
					return
				}
			}

			amended := &AssetCreation{}
			if err := convert(amended, tt.current); err != nil {
				t.Errorf("Failed to convert current value : %s", err)
				return
			}
			for i, amendment := range amendments {
				fip, err := FieldIndexPathFromBytes(amendment.FieldIndexPath)
				if err != nil {
					t.Errorf("Failed to parse FIP : %s", err)
					return
				}

				_, err = amended.ApplyAmendment(fip, amendment.Operation, amendment.Data)
				if err != nil {
					t.Errorf("Failed to apply amendment %d : %s", i, err)
					return
				}
			}

			newValue := &AssetCreation{}
			if err := convert(newValue, tt.newValue); err != nil {
				t.Errorf("Failed to convert new value : %s", err)
				return
			}
			if !amended.Equal(newValue) {
				t.Errorf("Amended value doesn't match : \n  got  %+v\n  want %+v", *amended,
					*newValue)
				return
			}
		})
	}
}

func convert(dst, src interface{}) error {
	// Marshal source object to json.
	var data []byte
	var err error
	data, err = json.Marshal(src)
	if err != nil {
		return errors.Wrap(err, "marshal")
	}

	// Unmarshal json back into destination object.
	err = json.Unmarshal(data, dst)
	if err != nil {
		return errors.Wrap(err, "unmarshal")
	}

	return nil
}
