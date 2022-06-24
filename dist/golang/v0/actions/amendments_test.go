package actions

import (
	"bytes"
	"testing"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/pkg/json"
	"github.com/tokenized/specification/dist/golang/v0/instruments"
	"github.com/tokenized/specification/dist/golang/v0/permissions"

	"github.com/golang/protobuf/proto"
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

	entityKey, err := bitcoin.GenerateKey(bitcoin.MainNet)
	if err != nil {
		t.Fatalf("Failed to generate key : %s", err)
	}
	entityAddress, err := entityKey.RawAddress()
	if err != nil {
		t.Fatalf("Failed to generate address : %s", err)
	}

	var buf bytes.Buffer
	if err := bitcoin.WriteBase128VarInt(&buf, uint64(1)); err != nil {
		t.Fatalf("Failed to write var int 1 : %s", err)
	}
	varInt1 := buf.Bytes()

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
				ContractName:             "Original Name",
				MasterAddress:            address1.Bytes(),
				RestrictedQtyInstruments: 1,
			},
			newValue: &ContractOffer{
				ContractName:             "New Name",
				MasterAddress:            address1.Bytes(),
				RestrictedQtyInstruments: 1,
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
				ContractName:             "Original Name",
				MasterAddress:            address1.Bytes(),
				RestrictedQtyInstruments: 1,
				Issuer: &EntityField{
					Name: "Original",
				},
			},
			newValue: &ContractOffer{
				ContractName:             "Original Name",
				MasterAddress:            address1.Bytes(),
				RestrictedQtyInstruments: 1,
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
				ContractName:             "Original Name",
				MasterAddress:            address1.Bytes(),
				RestrictedQtyInstruments: 1,
			},
			newValue: &ContractOffer{
				ContractName:             "Original Name",
				MasterAddress:            address1.Bytes(),
				RestrictedQtyInstruments: 2,
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{1, byte(ContractFieldRestrictedQtyInstruments)},
					Operation:      0,
					Data:           []byte{2},
				},
			},
		},
		{
			name: "Change 2 fields",
			current: &ContractFormation{
				ContractName:             "Original Name",
				MasterAddress:            address1.Bytes(),
				RestrictedQtyInstruments: 1,
			},
			newValue: &ContractOffer{
				ContractName:             "New Name",
				MasterAddress:            address1.Bytes(),
				RestrictedQtyInstruments: 2,
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{1, byte(ContractFieldContractName)},
					Operation:      0,
					Data:           []byte("New Name"),
				},
				&AmendmentField{
					FieldIndexPath: []byte{1, byte(ContractFieldRestrictedQtyInstruments)},
					Operation:      0,
					Data:           []byte{2},
				},
			},
		},
		{
			name: "Change service key",
			current: &ContractFormation{
				ContractName:             "Original Name",
				MasterAddress:            address1.Bytes(),
				RestrictedQtyInstruments: 1,
				Services: []*ServiceField{
					&ServiceField{
						Type:      ServiceTypeIdentityOracle,
						URL:       "test.identity.tokenized.com",
						PublicKey: publicKey1.Bytes(),
					},
				},
			},
			newValue: &ContractOffer{
				ContractName:             "Original Name",
				MasterAddress:            address1.Bytes(),
				RestrictedQtyInstruments: 1,
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
				ContractName:             "Original Name",
				MasterAddress:            address1.Bytes(),
				RestrictedQtyInstruments: 1,
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
				ContractName:             "Original Name",
				MasterAddress:            address1.Bytes(),
				RestrictedQtyInstruments: 1,
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
				ContractName:             "Original Name",
				MasterAddress:            address1.Bytes(),
				RestrictedQtyInstruments: 1,
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
				ContractName:             "Original Name",
				MasterAddress:            address1.Bytes(),
				RestrictedQtyInstruments: 1,
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
		{
			name: "Change oracle type",
			current: &ContractFormation{
				ContractName: "Original Name",
				Oracles: []*OracleField{
					&OracleField{
						OracleTypes:    []uint32{0},
						EntityContract: entityAddress.Bytes(),
					},
				},
				MasterAddress:            address1.Bytes(),
				RestrictedQtyInstruments: 1,
			},
			newValue: &ContractOffer{
				ContractName: "Original Name",
				Oracles: []*OracleField{
					&OracleField{
						OracleTypes:    []uint32{1},
						EntityContract: entityAddress.Bytes(),
					},
				},
				MasterAddress:            address1.Bytes(),
				RestrictedQtyInstruments: 1,
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{4, byte(ContractFieldOracles), 0,
						byte(OracleFieldOracleTypes), 0},
					Operation: 0,
					Data:      varInt1,
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

func TestContractApplyAmendments(t *testing.T) {
	tests := []struct {
		name        string
		current     *ContractFormation
		newValue    *ContractOffer
		permissions permissions.Permissions
		err         error
	}{
		{
			name: "Change Name",
			current: &ContractFormation{
				ContractName: "Name",
			},
			newValue: &ContractOffer{
				ContractName: "NewName",
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
			err: nil,
		},
		{
			name: "Remove Issuer",
			current: &ContractFormation{
				Issuer: &EntityField{
					Name: "Issuer Name",
					Type: EntitiesIndividual,
				},
			},
			newValue: &ContractOffer{},
			permissions: permissions.Permissions{
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
			name: "Remove Administration",
			current: &ContractFormation{
				Issuer: &EntityField{
					Name: "Issuer Name",
					Type: EntitiesIndividual,
					Administration: []*AdministratorField{
						&AdministratorField{
							Name: "Admin Name",
						},
					},
				},
			},
			newValue: &ContractOffer{
				Issuer: &EntityField{
					Name: "Issuer Name",
					Type: EntitiesIndividual,
				},
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

				_, err = amended.ApplyAmendment(fip, amendment.Operation, amendment.Data,
					tt.permissions)
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

func TestBodyOfAgreementCreateAmendments(t *testing.T) {

	newTerm := &DefinedTermField{
		Term:       "New Term",
		Definition: "Definition of new term",
	}

	newTermBytes, err := proto.Marshal(newTerm)
	if err != nil {
		t.Fatalf("Failed to marshal new term : %s", err)
	}

	tests := []struct {
		name       string
		current    *BodyOfAgreementFormation
		newValue   *BodyOfAgreementOffer
		err        error
		amendments []*AmendmentField
	}{
		{
			name: "Change Title",
			current: &BodyOfAgreementFormation{
				Chapters: []*ChapterField{
					&ChapterField{
						Title: "Title 1",
						// Preamble             string
						// Articles             []*ClauseField
					},
				},
			},
			newValue: &BodyOfAgreementOffer{
				Chapters: []*ChapterField{
					&ChapterField{
						Title: "Title 2",
						// Preamble             string
						// Articles             []*ClauseField
					},
				},
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{
						byte(3), // number of items in field index path
						byte(BodyOfAgreementFieldChapters),
						byte(0), // first chapter
						byte(ChapterFieldTitle),
					},
					Operation: 0,
					Data:      []byte("Title 2"),
				},
			},
		},
		{
			name: "Add Definition",
			current: &BodyOfAgreementFormation{
				Chapters: []*ChapterField{
					&ChapterField{
						Title: "Title 1",
					},
				},
			},
			newValue: &BodyOfAgreementOffer{
				Chapters: []*ChapterField{
					&ChapterField{
						Title:    "Title 1",
						Preamble: "Use of [New Term]().",
					},
				},
				Definitions: []*DefinedTermField{
					&DefinedTermField{
						Term:       "New Term",
						Definition: "Definition of new term",
					},
				},
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{
						byte(3), // number of items in field index path
						byte(BodyOfAgreementFieldChapters),
						byte(0), // first chapter
						byte(ChapterFieldPreamble),
					},
					Operation: 0,
					Data:      []byte("Use of [New Term]()."),
				},
				&AmendmentField{
					FieldIndexPath: []byte{
						byte(2), // number of items in field index path
						byte(BodyOfAgreementFieldDefinitions),
						byte(0), // first term
					},
					Operation: AmendmentOperationAddElement,
					Data:      newTermBytes,
				},
			},
		},
		{
			name: "Remove Definition",
			current: &BodyOfAgreementFormation{
				Chapters: []*ChapterField{
					&ChapterField{
						Title:    "Title 1",
						Preamble: "Use of [Term]().",
					},
				},
				Definitions: []*DefinedTermField{
					&DefinedTermField{
						Term:       "Term",
						Definition: "Definition of term",
					},
				},
			},
			newValue: &BodyOfAgreementOffer{
				Chapters: []*ChapterField{
					&ChapterField{
						Title: "Title 1",
					},
				},
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{
						byte(3), // number of items in field index path
						byte(BodyOfAgreementFieldChapters),
						byte(0), // first chapter
						byte(ChapterFieldPreamble),
					},
					Operation: 0,
					Data:      nil,
				},
				&AmendmentField{
					FieldIndexPath: []byte{
						byte(2), // number of items in field index path
						byte(BodyOfAgreementFieldDefinitions),
						byte(0), // first term
					},
					Operation: AmendmentOperationRemoveElement,
				},
			},
		},
		{
			name: "Add First Article",
			current: &BodyOfAgreementFormation{
				Chapters: []*ChapterField{
					&ChapterField{
						Title:    "Title 1",
						Preamble: "Use of [Term]().",
					},
				},
				Definitions: []*DefinedTermField{
					&DefinedTermField{
						Term:       "Term",
						Definition: "Definition of term",
					},
				},
			},
			newValue: &BodyOfAgreementOffer{
				Chapters: []*ChapterField{
					&ChapterField{
						Title:    "Title 1",
						Preamble: "Use of [Term]().",
						Articles: []*ClauseField{
							{
								Title: "Article 1",
								Body:  "Body 1",
							},
						},
					},
				},
				Definitions: []*DefinedTermField{
					&DefinedTermField{
						Term:       "Term",
						Definition: "Definition of term",
					},
				},
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{
						byte(4), // number of items in field index path
						byte(BodyOfAgreementFieldChapters),
						byte(0), // first chapter
						byte(ChapterFieldArticles),
						byte(0), // first article
					},
					Operation: 1,
					Data: []byte{10, 9, 65, 114, 116, 105, 99, 108, 101, 32, 49, 18, 6, 66,
						111, 100, 121, 32, 49},
				},
			},
		},
		{
			name: "Add Second Article",
			current: &BodyOfAgreementFormation{
				Chapters: []*ChapterField{
					&ChapterField{
						Title:    "Title 1",
						Preamble: "Use of [Term]().",
						Articles: []*ClauseField{
							{
								Title: "Article 1",
								Body:  "Body 1",
							},
						},
					},
				},
				Definitions: []*DefinedTermField{
					&DefinedTermField{
						Term:       "Term",
						Definition: "Definition of term",
					},
				},
			},
			newValue: &BodyOfAgreementOffer{
				Chapters: []*ChapterField{
					&ChapterField{
						Title:    "Title 1",
						Preamble: "Use of [Term]().",
						Articles: []*ClauseField{
							{
								Title: "Article 1",
								Body:  "Body 1",
							},
							{
								Title: "Article 2",
								Body:  "Body 2",
							},
						},
					},
				},
				Definitions: []*DefinedTermField{
					&DefinedTermField{
						Term:       "Term",
						Definition: "Definition of term",
					},
				},
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{
						byte(4), // number of items in field index path
						byte(BodyOfAgreementFieldChapters),
						byte(0), // first chapter
						byte(ChapterFieldArticles),
						byte(1), // second article
					},
					Operation: 1,
					Data: []byte{10, 9, 65, 114, 116, 105, 99, 108, 101, 32, 50, 18, 6, 66,
						111, 100, 121, 32, 50},
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

				_, err = amended.ApplyAmendment(fip, amendment.Operation, amendment.Data, nil)
				if err != nil {
					t.Errorf("Failed to apply amendment %d : %s", i, err)
					return
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

func TestInstrumentCreateAmendments(t *testing.T) {
	currency := &instruments.Currency{
		CurrencyCode: "USD",
		Precision:    100,
	}
	currencyPayload, _ := currency.Bytes()

	tests := []struct {
		name       string
		current    *InstrumentCreation
		newValue   *InstrumentDefinition
		err        error
		amendments []*AmendmentField
	}{
		{
			name: "Change transfers permitted",
			current: &InstrumentCreation{
				EnforcementOrdersPermitted: false,
				TradeRestrictions:          []string{"AUS"},
				AuthorizedTokenQty:         10000,
				InstrumentType:             "CCY",
				InstrumentPayload:          currencyPayload,
			},
			newValue: &InstrumentDefinition{
				EnforcementOrdersPermitted: true,
				TradeRestrictions:          []string{"AUS"},
				AuthorizedTokenQty:         10000,
				InstrumentType:             "CCY",
				InstrumentPayload:          currencyPayload,
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{1, byte(InstrumentFieldEnforcementOrdersPermitted)},
					Operation:      0,
					Data:           []byte{1},
				},
			},
		},
		{
			name: "Change token quantity",
			current: &InstrumentCreation{
				TradeRestrictions:  []string{"AUS"},
				AuthorizedTokenQty: 10000,
				InstrumentType:     "CCY",
				InstrumentPayload:  currencyPayload,
			},
			newValue: &InstrumentDefinition{
				TradeRestrictions:  []string{"AUS"},
				AuthorizedTokenQty: 100000,
				InstrumentType:     "CCY",
				InstrumentPayload:  currencyPayload,
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{1, byte(InstrumentFieldAuthorizedTokenQty)},
					Operation:      0,
					Data:           []byte{160, 141, 6}, // base 128 var int encoding of 100000
				},
			},
		},
		{
			name: "Change two fields",
			current: &InstrumentCreation{
				EnforcementOrdersPermitted: false,
				TradeRestrictions:          []string{"AUS"},
				AuthorizedTokenQty:         10000,
				InstrumentType:             "CCY",
				InstrumentPayload:          currencyPayload,
			},
			newValue: &InstrumentDefinition{
				EnforcementOrdersPermitted: true,
				TradeRestrictions:          []string{"AUS"},
				AuthorizedTokenQty:         100000,
				InstrumentType:             "CCY",
				InstrumentPayload:          currencyPayload,
			},
			err: nil,
			amendments: []*AmendmentField{
				&AmendmentField{
					FieldIndexPath: []byte{1, byte(InstrumentFieldEnforcementOrdersPermitted)},
					Operation:      0,
					Data:           []byte{1},
				},
				&AmendmentField{
					FieldIndexPath: []byte{1, byte(InstrumentFieldAuthorizedTokenQty)},
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

			amended := &InstrumentCreation{}
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

			newValue := &InstrumentCreation{}
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

func TestInstrumentCreateAmendmentsCouponName(t *testing.T) {
	currentCoupon := &instruments.Coupon{
		FaceValue: &instruments.CurrencyValueField{
			Value:        100,
			CurrencyCode: "USD",
			Precision:    2,
		},
		CouponName: "Test Coupon",
	}

	cb, _ := currentCoupon.Bytes()

	current := &InstrumentCreation{
		TradeRestrictions:  []string{"AUS"},
		AuthorizedTokenQty: 10000,
		InstrumentType:     instruments.CodeCoupon,
		InstrumentPayload:  cb,
	}

	newCoupon := &instruments.Coupon{
		FaceValue: &instruments.CurrencyValueField{
			Value:        100,
			CurrencyCode: "USD",
			Precision:    2,
		},
		CouponName: "New Test Coupon",
	}

	nb, _ := newCoupon.Bytes()

	newValue := &InstrumentDefinition{
		TradeRestrictions:  []string{"AUS"},
		AuthorizedTokenQty: 10000,
		InstrumentType:     instruments.CodeCoupon,
		InstrumentPayload:  nb,
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
		FieldIndexPath: []byte{2, byte(InstrumentFieldInstrumentPayload), byte(instruments.CouponFieldCouponName)},
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
