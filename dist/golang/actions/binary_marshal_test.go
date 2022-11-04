package actions

import (
	"bytes"
	"testing"
	"time"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/pkg/json"
)

func Test_BinaryMarshal_ContractOffer(t *testing.T) {
	key, _ := bitcoin.GenerateKey(bitcoin.TestNet)
	ra, _ := key.RawAddress()

	offer := &ContractOffer{
		ContractName:        "Contract Name",
		BodyOfAgreementType: ContractBodyOfAgreementTypeNone,
		// BodyOfAgreement           []byte
		// SupportingDocs            []*DocumentField
		ContractExpiration: uint64(time.Now().UnixNano()),
		ContractURI:        "mock://test.contract",
		Issuer: &EntityField{
			Name: "Issuer Name",
		},
		// ContractOperatorIncluded  bool
		ContractFee: 2000,
		VotingSystems: []*VotingSystemField{
			{
				Name:                    "Vote System 1",
				VoteType:                "A",
				TallyLogic:              0,
				ThresholdPercentage:     50,
				VoteMultiplierPermitted: true,
				HolderProposalFee:       20000,
			},
		},
		// ContractPermissions       []byte
		// RestrictedQtyInstruments  uint64
		// AdministrationProposal    bool
		HolderProposal: true,
		// Oracles                   []*OracleField
		MasterAddress: ra.Bytes(),
		// EntityContract            []byte
		// OperatorEntityContract    []byte
		ContractType: ContractTypeEntity,
		// Services                  []*ServiceField
		// AdminIdentityCertificates []*AdminIdentityCertificateField
		// GoverningLaw              string
		// Jurisdiction              string
	}

	js, _ := json.MarshalIndent(offer, "", "  ")
	t.Logf("Original Offer : %s", js)

	b, err := offer.MarshalBinary()
	if err != nil {
		t.Fatalf("Failed to binary marshal contract offer : %s", err)
	}

	if len(b) == 0 {
		t.Fatalf("Empty marshalled binary")
	}

	if b[0] != binaryVersion {
		t.Fatalf("Wrong binary version : got %d, want %d", b[0], binaryVersion)
	}

	newOffer := &ContractOffer{}
	if err := newOffer.UnmarshalBinary(b); err != nil {
		t.Fatalf("Failed to binary unmarshal : %s", err)
	}

	js, _ = json.MarshalIndent(newOffer, "", "  ")
	t.Logf("New Offer : %s", js)

	if !newOffer.Equal(offer) {
		t.Errorf("Offer not equal")
	}

	if newOffer.ContractName != offer.ContractName {
		t.Errorf("Contract name is wrong : got %s, want %s", newOffer.ContractName,
			offer.ContractName)
	}

	if newOffer.ContractFee != offer.ContractFee {
		t.Errorf("Contract fee is wrong : got %d, want %d", newOffer.ContractFee,
			offer.ContractFee)
	}

	if newOffer.Issuer == nil {
		t.Errorf("Missing issuer")
	} else if newOffer.Issuer.Name != offer.Issuer.Name {
		t.Errorf("Wrong issuer name : got %s, want %s", newOffer.Issuer.Name, offer.Issuer.Name)
	}

	if len(newOffer.VotingSystems) != 1 {
		t.Errorf("Wrong voting system count : got %d, want %d", len(newOffer.VotingSystems), 1)
	} else {
		if newOffer.VotingSystems[0].Name != offer.VotingSystems[0].Name {
			t.Errorf("Wrong voting system name : got %s, want %s", newOffer.VotingSystems[0].Name,
				offer.VotingSystems[0].Name)
		}

		if newOffer.VotingSystems[0].HolderProposalFee != offer.VotingSystems[0].HolderProposalFee {
			t.Errorf("Wrong voting system holder proposal fee : got %d, want %d",
				newOffer.VotingSystems[0].HolderProposalFee,
				offer.VotingSystems[0].HolderProposalFee)
		}
	}

	if newOffer.HolderProposal != offer.HolderProposal {
		t.Errorf("Contract holder proposal is wrong : got %t, want %t", newOffer.HolderProposal,
			offer.HolderProposal)
	}

	if len(newOffer.Oracles) != 0 {
		t.Errorf("Wrong oracle count : got %d, want %d", len(newOffer.Oracles), 0)
	}

	if len(newOffer.Services) != 0 {
		t.Errorf("Wrong service count : got %d, want %d", len(newOffer.Services), 0)
	}

	if !bytes.Equal(offer.MasterAddress, newOffer.MasterAddress) {
		t.Errorf("Wrong master address : got %x, want %x", newOffer.MasterAddress,
			offer.MasterAddress)
	}

	if newOffer.ContractType != offer.ContractType {
		t.Errorf("Contract type is wrong : got %d, want %d", newOffer.ContractType,
			offer.ContractType)
	}
}
