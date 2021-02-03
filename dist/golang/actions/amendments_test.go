package actions

import (
	"testing"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/pkg/json"
	"github.com/tokenized/specification/dist/golang/assets"
	"github.com/tokenized/specification/dist/golang/permissions"

	"github.com/pkg/errors"
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
			name: "Change issuer name",
			current: &ContractFormation{
				ContractName:        "Original Name",
				MasterAddress:       address1.Bytes(),
				RestrictedQtyAssets: 1,
				Issuer: &EntityField{
					Name: "Original",
				},
			},
			newValue: &ContractOffer{
				ContractName:        "Original Name",
				MasterAddress:       address1.Bytes(),
				RestrictedQtyAssets: 1,
				Issuer: &EntityField{
					Name: "New",
				},
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{2, byte(ContractFieldIssuer), byte(EntityFieldName)},
					Operation:      0,
					Data:           []byte("New"),
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
				fip, err := permissions.FieldIndexPathFromBytes(amendment.FieldIndexPath)
				if err != nil {
					t.Errorf("Failed to parse FIP : %s", err)
					return
				}

				_, err = amended.ApplyAmendment(fip, amendment.Operation, amendment.Data, nil)
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

func TestBodyOfAgreementApplyAmendments(t *testing.T) {
	tests := []struct {
		name              string
		current           *BodyOfAgreementFormation
		newValue          *BodyOfAgreementOffer
		permissions       permissions.Permissions
		resultPermissions permissions.Permissions
		err               error
	}{
		{
			name: "Change Title",
			current: &BodyOfAgreementFormation{
				Chapters: []*ChapterField{
					&ChapterField{
						Title:    "Agreement Title",
						Preamble: "",
						Articles: []*ClauseField{},
					},
				},
				Definitions: []*DefinedTermField{},
			},
			newValue: &BodyOfAgreementOffer{
				Chapters: []*ChapterField{
					&ChapterField{
						Title:    "Agreement Title 2",
						Preamble: "",
						Articles: []*ClauseField{},
					},
				},
				Definitions: []*DefinedTermField{},
			},
			permissions: permissions.Permissions{
				permissions.Permission{
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         true,
					AdministrativeMatter:   false,
					Fields:                 []permissions.FieldIndexPath{},
				},
			},
			resultPermissions: permissions.Permissions{
				permissions.Permission{
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         true,
					AdministrativeMatter:   false,
					Fields:                 []permissions.FieldIndexPath{},
				},
			},
			err: nil,
		},
		{
			name: "Change Preamble Permissions",
			current: &BodyOfAgreementFormation{
				Chapters: []*ChapterField{
					&ChapterField{
						Title:    "Agreement Title",
						Preamble: "Preamble",
						Articles: []*ClauseField{},
					},
				},
				Definitions: []*DefinedTermField{},
			},
			newValue: &BodyOfAgreementOffer{
				Chapters: []*ChapterField{
					&ChapterField{
						Title:    "Agreement Title",
						Preamble: "Preamble 1",
						Articles: []*ClauseField{},
					},
				},
				Definitions: []*DefinedTermField{},
			},
			permissions: permissions.Permissions{
				permissions.Permission{ // match all
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         true,
					AdministrativeMatter:   false,
					Fields:                 []permissions.FieldIndexPath{},
				},
				permissions.Permission{
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         false,
					AdministrativeMatter:   true,
					Fields: []permissions.FieldIndexPath{
						permissions.FieldIndexPath{
							BodyOfAgreementFieldChapters,
							0,
							ChapterFieldPreamble,
						},
					},
				},
			},
			resultPermissions: permissions.Permissions{
				permissions.Permission{ // specific match
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         false,
					AdministrativeMatter:   true,
					Fields:                 []permissions.FieldIndexPath{},
				},
				permissions.Permission{ // general match
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         true,
					AdministrativeMatter:   false,
					Fields:                 []permissions.FieldIndexPath{},
				},
			},
			err: nil,
		},
		{
			name: "Change Article Title",
			current: &BodyOfAgreementFormation{
				Chapters: []*ChapterField{
					&ChapterField{
						Title:    "Agreement Title",
						Preamble: "Preamble",
						Articles: []*ClauseField{
							&ClauseField{
								Title: "Article Title",
							},
						},
					},
				},
				Definitions: []*DefinedTermField{},
			},
			newValue: &BodyOfAgreementOffer{
				Chapters: []*ChapterField{
					&ChapterField{
						Title:    "Agreement Title",
						Preamble: "Preamble",
						Articles: []*ClauseField{
							&ClauseField{
								Title: "Article Title 2",
							},
						},
					},
				},
				Definitions: []*DefinedTermField{},
			},
			permissions: permissions.Permissions{
				permissions.Permission{ // non-match
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         true,
					AdministrativeMatter:   false,
					Fields: []permissions.FieldIndexPath{
						permissions.FieldIndexPath{8},
					},
				},
				permissions.Permission{
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         false,
					AdministrativeMatter:   true,
					Fields: []permissions.FieldIndexPath{
						permissions.FieldIndexPath{
							BodyOfAgreementFieldChapters,
							0,
							ChapterFieldArticles,
							ClauseFieldTitle,
						},
					},
				},
			},
			resultPermissions: permissions.Permissions{
				permissions.Permission{ // specific match
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         false,
					AdministrativeMatter:   true,
					Fields:                 []permissions.FieldIndexPath{},
				},
			},
			err: nil,
		},
		{
			name: "Change Clause Body",
			current: &BodyOfAgreementFormation{
				Chapters: []*ChapterField{
					&ChapterField{
						Title:    "Agreement Title",
						Preamble: "Preamble",
						Articles: []*ClauseField{
							&ClauseField{
								Title: "Article Title",
								Children: []*ClauseField{
									&ClauseField{
										Body: "Clause Title",
									},
								},
							},
						},
					},
				},
				Definitions: []*DefinedTermField{},
			},
			newValue: &BodyOfAgreementOffer{
				Chapters: []*ChapterField{
					&ChapterField{
						Title:    "Agreement Title",
						Preamble: "Preamble",
						Articles: []*ClauseField{
							&ClauseField{
								Title: "Article Title",
								Children: []*ClauseField{
									&ClauseField{
										Body: "Clause Title 2",
									},
								},
							},
						},
					},
				},
				Definitions: []*DefinedTermField{},
			},
			permissions: permissions.Permissions{
				permissions.Permission{ // non-match
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         true,
					AdministrativeMatter:   false,
					Fields: []permissions.FieldIndexPath{
						permissions.FieldIndexPath{8},
					},
				},
				permissions.Permission{
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         false,
					AdministrativeMatter:   true,
					Fields: []permissions.FieldIndexPath{
						permissions.FieldIndexPath{
							BodyOfAgreementFieldChapters,
							0,
							ChapterFieldArticles,
							0,
							ClauseFieldChildren,
							0,
							ClauseFieldBody,
						},
					},
				},
			},
			resultPermissions: permissions.Permissions{
				permissions.Permission{ // specific match
					Permitted:              false,
					AdministrationProposal: false,
					HolderProposal:         false,
					AdministrativeMatter:   true,
					Fields:                 []permissions.FieldIndexPath{},
				},
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amendments, err := tt.current.CreateAmendments(tt.newValue)
			if err != nil {
				t.Errorf("Failed to create amendments : %s", err)
				return
			}

			amended := &BodyOfAgreementFormation{}
			if err := convert(amended, tt.current); err != nil {
				t.Errorf("Failed to convert current value : %s", err)
				return
			}

			for i, amendment := range amendments {
				fip, err := permissions.FieldIndexPathFromBytes(amendment.FieldIndexPath)
				if err != nil {
					t.Errorf("Failed to parse FIP : %s", err)
					return
				}

				resultPermissions, err := amended.ApplyAmendment(fip, amendment.Operation,
					amendment.Data, tt.permissions)
				if err != nil {
					t.Errorf("Failed to apply amendment %d : %s", i, err)
					return
				}

				if !resultPermissions.Equal(tt.resultPermissions) {
					t.Errorf("Wrong result permissions : \ngot  %+v\nwant %+v", resultPermissions,
						tt.resultPermissions)
				}
			}

			newValue := &BodyOfAgreementFormation{}
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
	currency := &assets.Currency{
		CurrencyCode: "USD",
		Precision:    100,
	}
	currencyPayload, _ := currency.Bytes()

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
				EnforcementOrdersPermitted: false,
				TradeRestrictions:          []string{"AUS"},
				AuthorizedTokenQty:         10000,
				AssetType:                  "CCY",
				AssetPayload:               currencyPayload,
			},
			newValue: &AssetDefinition{
				EnforcementOrdersPermitted: true,
				TradeRestrictions:          []string{"AUS"},
				AuthorizedTokenQty:         10000,
				AssetType:                  "CCY",
				AssetPayload:               currencyPayload,
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{1, byte(AssetFieldEnforcementOrdersPermitted)},
					Operation:      0,
					Data:           []byte{1},
				},
			},
		},
		{
			name: "Change token quantity",
			current: &AssetCreation{
				TradeRestrictions:  []string{"AUS"},
				AuthorizedTokenQty: 10000,
				AssetType:          "CCY",
				AssetPayload:       currencyPayload,
			},
			newValue: &AssetDefinition{
				TradeRestrictions:  []string{"AUS"},
				AuthorizedTokenQty: 100000,
				AssetType:          "CCY",
				AssetPayload:       currencyPayload,
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{1, byte(AssetFieldAuthorizedTokenQty)},
					Operation:      0,
					Data:           []byte{160, 141, 6}, // base 128 var int encoding of 100000
				},
			},
		},
		{
			name: "Change two fields",
			current: &AssetCreation{
				EnforcementOrdersPermitted: false,
				TradeRestrictions:          []string{"AUS"},
				AuthorizedTokenQty:         10000,
				AssetType:                  "CCY",
				AssetPayload:               currencyPayload,
			},
			newValue: &AssetDefinition{
				EnforcementOrdersPermitted: true,
				TradeRestrictions:          []string{"AUS"},
				AuthorizedTokenQty:         100000,
				AssetType:                  "CCY",
				AssetPayload:               currencyPayload,
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{1, byte(AssetFieldEnforcementOrdersPermitted)},
					Operation:      0,
					Data:           []byte{1},
				},
				&AmendmentField{
					FieldIndexPath: []byte{1, byte(AssetFieldAuthorizedTokenQty)},
					Operation:      0,
					Data:           []byte{160, 141, 6}, // base 128 var int encoding of 100000
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
				fip, err := permissions.FieldIndexPathFromBytes(amendment.FieldIndexPath)
				if err != nil {
					t.Errorf("Failed to parse FIP : %s", err)
					return
				}

				_, err = amended.ApplyAmendment(fip, amendment.Operation, amendment.Data, nil)
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

func TestAssetCreateAmendmentsCouponName(t *testing.T) {
	currentCoupon := &assets.Coupon{
		Currency:    "USD",
		Value:       1,
		Precision:   100,
		Description: "Test Coupon",
	}

	cb, _ := currentCoupon.Bytes()

	current := &AssetCreation{
		TradeRestrictions:  []string{"AUS"},
		AuthorizedTokenQty: 10000,
		AssetType:          assets.CodeCoupon,
		AssetPayload:       cb,
	}

	newCoupon := &assets.Coupon{
		Currency:    "USD",
		Value:       1,
		Precision:   100,
		Description: "New Test Coupon",
	}

	nb, _ := newCoupon.Bytes()

	newValue := &AssetDefinition{
		TradeRestrictions:  []string{"AUS"},
		AuthorizedTokenQty: 10000,
		AssetType:          assets.CodeCoupon,
		AssetPayload:       nb,
	}

	amendments, err := current.CreateAmendments(newValue)
	if err != nil {
		t.Fatalf("Failed to create amendments : %s", err)
	}

	if len(amendments) != 1 {
		t.Fatalf("Wrong amendment count : got %d, want %d", len(amendments), 1)
	}

	// Check amendment
	expectedAmendment := &AmendmentField{
		FieldIndexPath: []byte{2, byte(AssetFieldAssetPayload), byte(assets.CouponFieldDescription)},
		Operation:      0,
		Data:           []byte("New Test Coupon"),
	}

	if !amendments[0].Equal(expectedAmendment) {
		t.Fatalf("Wrong amendment : \n   got %+v\n  want %+v", amendments[0], expectedAmendment)
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
