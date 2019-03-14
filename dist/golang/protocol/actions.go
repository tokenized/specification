package protocol

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	// CodeAssetDefinition identifies data as a AssetDefinition message.
	CodeAssetDefinition = "A1"

	// CodeAssetCreation identifies data as a AssetCreation message.
	CodeAssetCreation = "A2"

	// CodeAssetModification identifies data as a AssetModification message.
	CodeAssetModification = "A3"

	// CodeContractOffer identifies data as a ContractOffer message.
	CodeContractOffer = "C1"

	// CodeContractFormation identifies data as a ContractFormation message.
	CodeContractFormation = "C2"

	// CodeContractAmendment identifies data as a ContractAmendment message.
	CodeContractAmendment = "C3"

	// CodeStaticContractFormation identifies data as a
	// StaticContractFormation message.
	CodeStaticContractFormation = "C4"

	// CodeOrder identifies data as a Order message.
	CodeOrder = "E1"

	// CodeFreeze identifies data as a Freeze message.
	CodeFreeze = "E2"

	// CodeThaw identifies data as a Thaw message.
	CodeThaw = "E3"

	// CodeConfiscation identifies data as a Confiscation message.
	CodeConfiscation = "E4"

	// CodeReconciliation identifies data as a Reconciliation message.
	CodeReconciliation = "E5"

	// CodeInitiative identifies data as a Initiative message.
	CodeInitiative = "G1"

	// CodeReferendum identifies data as a Referendum message.
	CodeReferendum = "G2"

	// CodeVote identifies data as a Vote message.
	CodeVote = "G3"

	// CodeBallotCast identifies data as a BallotCast message.
	CodeBallotCast = "G4"

	// CodeBallotCounted identifies data as a BallotCounted message.
	CodeBallotCounted = "G5"

	// CodeResult identifies data as a Result message.
	CodeResult = "G6"

	// CodeMessage identifies data as a Message message.
	CodeMessage = "M1"

	// CodeRejection identifies data as a Rejection message.
	CodeRejection = "M2"

	// CodeEstablishment identifies data as a Establishment message.
	CodeEstablishment = "R1"

	// CodeAddition identifies data as a Addition message.
	CodeAddition = "R2"

	// CodeAlteration identifies data as a Alteration message.
	CodeAlteration = "R3"

	// CodeRemoval identifies data as a Removal message.
	CodeRemoval = "R4"

	// CodeTransfer identifies data as a Transfer message.
	CodeTransfer = "T1"

	// CodeSettlement identifies data as a Settlement message.
	CodeSettlement = "T4"
)

// TypeMapping holds a mapping of message codes to message types.
var (
	TypeMapping = map[string]OpReturnMessage{
		CodeAssetDefinition:         &AssetDefinition{},
		CodeAssetCreation:           &AssetCreation{},
		CodeAssetModification:       &AssetModification{},
		CodeContractOffer:           &ContractOffer{},
		CodeContractFormation:       &ContractFormation{},
		CodeContractAmendment:       &ContractAmendment{},
		CodeStaticContractFormation: &StaticContractFormation{},
		CodeOrder:                   &Order{},
		CodeFreeze:                  &Freeze{},
		CodeThaw:                    &Thaw{},
		CodeConfiscation:            &Confiscation{},
		CodeReconciliation:          &Reconciliation{},
		CodeInitiative:              &Initiative{},
		CodeReferendum:              &Referendum{},
		CodeVote:                    &Vote{},
		CodeBallotCast:              &BallotCast{},
		CodeBallotCounted:           &BallotCounted{},
		CodeResult:                  &Result{},
		CodeMessage:                 &Message{},
		CodeRejection:               &Rejection{},
		CodeEstablishment:           &Establishment{},
		CodeAddition:                &Addition{},
		CodeAlteration:              &Alteration{},
		CodeRemoval:                 &Removal{},
		CodeTransfer:                &Transfer{},
		CodeSettlement:              &Settlement{},
	}

	// ProtocolID is the current protocol ID
	ProtocolID = []byte("tokenized.com")
)

// AssetDefinition Asset Definition Action - This action is used by the
// issuer to define the properties/characteristics of the Asset (token)
// that it wants to create.
type AssetDefinition struct {
	Header                      Header
	AssetType                   []byte
	AssetID                     []byte
	AssetAuthFlags              []byte
	TransfersPermitted          bool
	TradeRestrictions           []byte
	EnforcementOrdersPermitted  bool
	VoteMultiplier              uint8
	ReferendumProposal          bool
	InitiativeProposal          bool
	AssetModificationGovernance bool
	TokenQty                    uint64
	ContractFeeCurrency         []byte
	ContractFeeVar              float32
	ContractFeeFixed            float32
	AssetPayloadLen             uint16
	AssetPayload                []byte
}

// NewAssetDefinition returns a new AssetDefinition with defaults set.
func NewAssetDefinition() AssetDefinition {
	return AssetDefinition{}
}

// Type returns the type identifer for this message.
func (m AssetDefinition) Type() string {
	return CodeAssetDefinition
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m AssetDefinition) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m AssetDefinition) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, pad(m.AssetType, 3)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AssetID, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AssetAuthFlags, 8)); err != nil {
		return nil, err
	}

	if err := write(buf, m.TransfersPermitted); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.TradeRestrictions, 3)); err != nil {
		return nil, err
	}

	if err := write(buf, m.EnforcementOrdersPermitted); err != nil {
		return nil, err
	}

	if err := write(buf, m.VoteMultiplier); err != nil {
		return nil, err
	}

	if err := write(buf, m.ReferendumProposal); err != nil {
		return nil, err
	}

	if err := write(buf, m.InitiativeProposal); err != nil {
		return nil, err
	}

	if err := write(buf, m.AssetModificationGovernance); err != nil {
		return nil, err
	}

	if err := write(buf, m.TokenQty); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.ContractFeeCurrency, 3)); err != nil {
		return nil, err
	}

	if err := write(buf, m.ContractFeeVar); err != nil {
		return nil, err
	}

	if err := write(buf, m.ContractFeeFixed); err != nil {
		return nil, err
	}

	if err := write(buf, m.AssetPayloadLen); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.AssetPayloadLen); i++ {
		if err := write(buf, m.AssetPayload[i]); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeAssetDefinition, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *AssetDefinition) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	m.AssetType = make([]byte, 3)
	if err := readLen(buf, m.AssetType); err != nil {
		return 0, err
	}

	m.AssetID = make([]byte, 32)
	if err := readLen(buf, m.AssetID); err != nil {
		return 0, err
	}

	m.AssetAuthFlags = make([]byte, 8)
	if err := readLen(buf, m.AssetAuthFlags); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TransfersPermitted); err != nil {
		return 0, err
	}

	m.TradeRestrictions = make([]byte, 3)
	if err := readLen(buf, m.TradeRestrictions); err != nil {
		return 0, err
	}

	if err := read(buf, &m.EnforcementOrdersPermitted); err != nil {
		return 0, err
	}

	if err := read(buf, &m.VoteMultiplier); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ReferendumProposal); err != nil {
		return 0, err
	}

	if err := read(buf, &m.InitiativeProposal); err != nil {
		return 0, err
	}

	if err := read(buf, &m.AssetModificationGovernance); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TokenQty); err != nil {
		return 0, err
	}

	m.ContractFeeCurrency = make([]byte, 3)
	if err := readLen(buf, m.ContractFeeCurrency); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ContractFeeVar); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ContractFeeFixed); err != nil {
		return 0, err
	}

	if err := read(buf, &m.AssetPayloadLen); err != nil {
		return 0, err
	}

	m.AssetPayload = make([]byte, m.AssetPayloadLen, m.AssetPayloadLen)
	if err := read(buf, &m.AssetPayload); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m AssetDefinition) PayloadMessage() (PayloadMessage, error) {
	p, err := NewPayloadMessageFromCode(m.AssetType)
	if p == nil || err != nil {
		return nil, err
	}

	if _, err := p.Write(m.AssetPayload); err != nil {
		return nil, err
	}

	return p, nil
}

func (m AssetDefinition) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("AssetType:%#x", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#x", m.AssetID))
	vals = append(vals, fmt.Sprintf("AssetAuthFlags:%#x", m.AssetAuthFlags))
	vals = append(vals, fmt.Sprintf("TransfersPermitted:%#+v", m.TransfersPermitted))
	vals = append(vals, fmt.Sprintf("TradeRestrictions:%#x", m.TradeRestrictions))
	vals = append(vals, fmt.Sprintf("EnforcementOrdersPermitted:%#+v", m.EnforcementOrdersPermitted))
	vals = append(vals, fmt.Sprintf("VoteMultiplier:%v", m.VoteMultiplier))
	vals = append(vals, fmt.Sprintf("ReferendumProposal:%#+v", m.ReferendumProposal))
	vals = append(vals, fmt.Sprintf("InitiativeProposal:%#+v", m.InitiativeProposal))
	vals = append(vals, fmt.Sprintf("AssetModificationGovernance:%#+v", m.AssetModificationGovernance))
	vals = append(vals, fmt.Sprintf("TokenQty:%v", m.TokenQty))
	vals = append(vals, fmt.Sprintf("ContractFeeCurrency:%#x", m.ContractFeeCurrency))
	vals = append(vals, fmt.Sprintf("ContractFeeVar:%v", m.ContractFeeVar))
	vals = append(vals, fmt.Sprintf("ContractFeeFixed:%v", m.ContractFeeFixed))
	vals = append(vals, fmt.Sprintf("AssetPayloadLen:%v", m.AssetPayloadLen))
	vals = append(vals, fmt.Sprintf("AssetPayload:%#x", m.AssetPayload))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// AssetCreation Asset Creation Action - This action creates an Asset in
// response to the Issuer's instructions in the Definition Action.
type AssetCreation struct {
	Header                      Header
	AssetType                   []byte
	AssetID                     []byte
	AssetAuthFlags              []byte
	TransfersPermitted          bool
	TradeRestrictions           []byte
	EnforcementOrdersPermitted  bool
	VoteMultiplier              uint8
	ReferendumProposal          bool
	InitiativeProposal          bool
	AssetModificationGovernance bool
	TokenQty                    uint64
	ContractFeeCurrency         []byte
	ContractFeeVar              float32
	ContractFeeFixed            float32
	AssetPayloadLen             uint16
	AssetPayload                []byte
	AssetRevision               uint64
	Timestamp                   uint64
}

// NewAssetCreation returns a new AssetCreation with defaults set.
func NewAssetCreation() AssetCreation {
	return AssetCreation{}
}

// Type returns the type identifer for this message.
func (m AssetCreation) Type() string {
	return CodeAssetCreation
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m AssetCreation) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m AssetCreation) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, pad(m.AssetType, 3)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AssetID, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AssetAuthFlags, 8)); err != nil {
		return nil, err
	}

	if err := write(buf, m.TransfersPermitted); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.TradeRestrictions, 3)); err != nil {
		return nil, err
	}

	if err := write(buf, m.EnforcementOrdersPermitted); err != nil {
		return nil, err
	}

	if err := write(buf, m.VoteMultiplier); err != nil {
		return nil, err
	}

	if err := write(buf, m.ReferendumProposal); err != nil {
		return nil, err
	}

	if err := write(buf, m.InitiativeProposal); err != nil {
		return nil, err
	}

	if err := write(buf, m.AssetModificationGovernance); err != nil {
		return nil, err
	}

	if err := write(buf, m.TokenQty); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.ContractFeeCurrency, 3)); err != nil {
		return nil, err
	}

	if err := write(buf, m.ContractFeeVar); err != nil {
		return nil, err
	}

	if err := write(buf, m.ContractFeeFixed); err != nil {
		return nil, err
	}

	if err := write(buf, m.AssetPayloadLen); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.AssetPayloadLen); i++ {
		if err := write(buf, m.AssetPayload[i]); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.AssetRevision); err != nil {
		return nil, err
	}

	if err := write(buf, m.Timestamp); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeAssetCreation, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *AssetCreation) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	m.AssetType = make([]byte, 3)
	if err := readLen(buf, m.AssetType); err != nil {
		return 0, err
	}

	m.AssetID = make([]byte, 32)
	if err := readLen(buf, m.AssetID); err != nil {
		return 0, err
	}

	m.AssetAuthFlags = make([]byte, 8)
	if err := readLen(buf, m.AssetAuthFlags); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TransfersPermitted); err != nil {
		return 0, err
	}

	m.TradeRestrictions = make([]byte, 3)
	if err := readLen(buf, m.TradeRestrictions); err != nil {
		return 0, err
	}

	if err := read(buf, &m.EnforcementOrdersPermitted); err != nil {
		return 0, err
	}

	if err := read(buf, &m.VoteMultiplier); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ReferendumProposal); err != nil {
		return 0, err
	}

	if err := read(buf, &m.InitiativeProposal); err != nil {
		return 0, err
	}

	if err := read(buf, &m.AssetModificationGovernance); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TokenQty); err != nil {
		return 0, err
	}

	m.ContractFeeCurrency = make([]byte, 3)
	if err := readLen(buf, m.ContractFeeCurrency); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ContractFeeVar); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ContractFeeFixed); err != nil {
		return 0, err
	}

	if err := read(buf, &m.AssetPayloadLen); err != nil {
		return 0, err
	}

	m.AssetPayload = make([]byte, m.AssetPayloadLen, m.AssetPayloadLen)
	if err := read(buf, &m.AssetPayload); err != nil {
		return 0, err
	}

	if err := read(buf, &m.AssetRevision); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Timestamp); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m AssetCreation) PayloadMessage() (PayloadMessage, error) {
	p, err := NewPayloadMessageFromCode(m.AssetType)
	if p == nil || err != nil {
		return nil, err
	}

	if _, err := p.Write(m.AssetPayload); err != nil {
		return nil, err
	}

	return p, nil
}

func (m AssetCreation) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("AssetType:%#x", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#x", m.AssetID))
	vals = append(vals, fmt.Sprintf("AssetAuthFlags:%#x", m.AssetAuthFlags))
	vals = append(vals, fmt.Sprintf("TransfersPermitted:%#+v", m.TransfersPermitted))
	vals = append(vals, fmt.Sprintf("TradeRestrictions:%#x", m.TradeRestrictions))
	vals = append(vals, fmt.Sprintf("EnforcementOrdersPermitted:%#+v", m.EnforcementOrdersPermitted))
	vals = append(vals, fmt.Sprintf("VoteMultiplier:%v", m.VoteMultiplier))
	vals = append(vals, fmt.Sprintf("ReferendumProposal:%#+v", m.ReferendumProposal))
	vals = append(vals, fmt.Sprintf("InitiativeProposal:%#+v", m.InitiativeProposal))
	vals = append(vals, fmt.Sprintf("AssetModificationGovernance:%#+v", m.AssetModificationGovernance))
	vals = append(vals, fmt.Sprintf("TokenQty:%v", m.TokenQty))
	vals = append(vals, fmt.Sprintf("ContractFeeCurrency:%#x", m.ContractFeeCurrency))
	vals = append(vals, fmt.Sprintf("ContractFeeVar:%v", m.ContractFeeVar))
	vals = append(vals, fmt.Sprintf("ContractFeeFixed:%v", m.ContractFeeFixed))
	vals = append(vals, fmt.Sprintf("AssetPayloadLen:%v", m.AssetPayloadLen))
	vals = append(vals, fmt.Sprintf("AssetPayload:%#x", m.AssetPayload))
	vals = append(vals, fmt.Sprintf("AssetRevision:%v", m.AssetRevision))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// AssetModification Asset Modification Action - Token Dilutions, Call
// Backs/Revocations, burning etc.
type AssetModification struct {
	Header            Header
	AssetType         []byte
	AssetID           []byte
	AssetRevision     uint64
	ModificationCount uint8
	Modifications     []Amendment
	RefTxID           []byte
}

// NewAssetModification returns a new AssetModification with defaults set.
func NewAssetModification() AssetModification {
	return AssetModification{}
}

// Type returns the type identifer for this message.
func (m AssetModification) Type() string {
	return CodeAssetModification
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m AssetModification) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m AssetModification) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, pad(m.AssetType, 3)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AssetID, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, m.AssetRevision); err != nil {
		return nil, err
	}

	if err := write(buf, m.ModificationCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.ModificationCount); i++ {
		b, err := m.Modifications[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, pad(m.RefTxID, 32)); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeAssetModification, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *AssetModification) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	m.AssetType = make([]byte, 3)
	if err := readLen(buf, m.AssetType); err != nil {
		return 0, err
	}

	m.AssetID = make([]byte, 32)
	if err := readLen(buf, m.AssetID); err != nil {
		return 0, err
	}

	if err := read(buf, &m.AssetRevision); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ModificationCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.ModificationCount); i++ {
		x := &Amendment{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.Modifications = append(m.Modifications, *x)
	}

	m.RefTxID = make([]byte, 32)
	if err := readLen(buf, m.RefTxID); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m AssetModification) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m AssetModification) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("AssetType:%#x", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#x", m.AssetID))
	vals = append(vals, fmt.Sprintf("AssetRevision:%v", m.AssetRevision))
	vals = append(vals, fmt.Sprintf("ModificationCount:%v", m.ModificationCount))
	vals = append(vals, fmt.Sprintf("Modifications:%#+v", m.Modifications))
	vals = append(vals, fmt.Sprintf("RefTxID:%#x", m.RefTxID))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// ContractOffer The Contract Offer action allows the Issuer to tell the
// smart contract what they want the details (labels, data, T&C's, etc.) of
// the Contract to be on-chain in a public and immutable way. The Contract
// Offer action 'initializes' a generic smart contract that has been spun
// up by either the Smart Contract Operator or the Issuer. This on-chain
// action allows for the positive response from the smart contract with
// either a Contract Formation Action or a Rejection Action.
type ContractOffer struct {
	Header                     Header
	ContractName               Nvarchar8
	ContractFileType           uint8
	LenContractFile            uint32
	ContractFile               []byte
	GoverningLaw               []byte
	Jurisdiction               []byte
	ContractExpiration         uint64
	ContractURI                Nvarchar8
	IssuerName                 Nvarchar8
	IssuerType                 byte
	IssuerLogoURL              Nvarchar8
	ContractOperatorID         Nvarchar8
	ContractAuthFlags          []byte
	VotingSystemCount          uint8
	VotingSystems              []VotingSystem
	RestrictedQtyAssets        uint64
	ReferendumProposal         bool
	InitiativeProposal         bool
	RegistryCount              uint8
	Registries                 []Registry
	IssuerAddress              bool
	UnitNumber                 Nvarchar8
	BuildingNumber             Nvarchar8
	Street                     Nvarchar16
	SuburbCity                 Nvarchar8
	TerritoryStateProvinceCode []byte
	CountryCode                []byte
	PostalZIPCode              Nvarchar8
	EmailAddress               Nvarchar8
	PhoneNumber                Nvarchar8
	KeyRolesCount              uint8
	KeyRoles                   []KeyRole
	NotableRolesCount          uint8
	NotableRoles               []NotableRole
}

// NewContractOffer returns a new ContractOffer with defaults set.
func NewContractOffer() ContractOffer {
	return ContractOffer{}
}

// Type returns the type identifer for this message.
func (m ContractOffer) Type() string {
	return CodeContractOffer
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m ContractOffer) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m ContractOffer) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	{
		b, err := m.ContractName.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.ContractFileType); err != nil {
		return nil, err
	}

	if err := write(buf, m.LenContractFile); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.ContractFile, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.GoverningLaw, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.Jurisdiction, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, m.ContractExpiration); err != nil {
		return nil, err
	}

	{
		b, err := m.ContractURI.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.IssuerName.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.IssuerType); err != nil {
		return nil, err
	}

	{
		b, err := m.IssuerLogoURL.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.ContractOperatorID.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, pad(m.ContractAuthFlags, 16)); err != nil {
		return nil, err
	}

	if err := write(buf, m.VotingSystemCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.VotingSystemCount); i++ {
		b, err := m.VotingSystems[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.RestrictedQtyAssets); err != nil {
		return nil, err
	}

	if err := write(buf, m.ReferendumProposal); err != nil {
		return nil, err
	}

	if err := write(buf, m.InitiativeProposal); err != nil {
		return nil, err
	}

	if err := write(buf, m.RegistryCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.RegistryCount); i++ {
		b, err := m.Registries[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.IssuerAddress); err != nil {
		return nil, err
	}

	{
		b, err := m.UnitNumber.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.BuildingNumber.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.Street.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.SuburbCity.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, pad(m.TerritoryStateProvinceCode, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.CountryCode, 3)); err != nil {
		return nil, err
	}

	{
		b, err := m.PostalZIPCode.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.EmailAddress.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.PhoneNumber.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.KeyRolesCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.KeyRolesCount); i++ {
		b, err := m.KeyRoles[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.NotableRolesCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.NotableRolesCount); i++ {
		b, err := m.NotableRoles[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeContractOffer, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *ContractOffer) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := m.ContractName.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ContractFileType); err != nil {
		return 0, err
	}

	if err := read(buf, &m.LenContractFile); err != nil {
		return 0, err
	}

	m.ContractFile = make([]byte, 32)
	if err := readLen(buf, m.ContractFile); err != nil {
		return 0, err
	}

	m.GoverningLaw = make([]byte, 5)
	if err := readLen(buf, m.GoverningLaw); err != nil {
		return 0, err
	}

	m.Jurisdiction = make([]byte, 5)
	if err := readLen(buf, m.Jurisdiction); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ContractExpiration); err != nil {
		return 0, err
	}

	if err := m.ContractURI.Write(buf); err != nil {
		return 0, err
	}

	if err := m.IssuerName.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.IssuerType); err != nil {
		return 0, err
	}

	if err := m.IssuerLogoURL.Write(buf); err != nil {
		return 0, err
	}

	if err := m.ContractOperatorID.Write(buf); err != nil {
		return 0, err
	}

	m.ContractAuthFlags = make([]byte, 16)
	if err := readLen(buf, m.ContractAuthFlags); err != nil {
		return 0, err
	}

	if err := read(buf, &m.VotingSystemCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.VotingSystemCount); i++ {
		x := &VotingSystem{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.VotingSystems = append(m.VotingSystems, *x)
	}

	if err := read(buf, &m.RestrictedQtyAssets); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ReferendumProposal); err != nil {
		return 0, err
	}

	if err := read(buf, &m.InitiativeProposal); err != nil {
		return 0, err
	}

	if err := read(buf, &m.RegistryCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.RegistryCount); i++ {
		x := &Registry{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.Registries = append(m.Registries, *x)
	}

	if err := read(buf, &m.IssuerAddress); err != nil {
		return 0, err
	}

	if err := m.UnitNumber.Write(buf); err != nil {
		return 0, err
	}

	if err := m.BuildingNumber.Write(buf); err != nil {
		return 0, err
	}

	if err := m.Street.Write(buf); err != nil {
		return 0, err
	}

	if err := m.SuburbCity.Write(buf); err != nil {
		return 0, err
	}

	m.TerritoryStateProvinceCode = make([]byte, 5)
	if err := readLen(buf, m.TerritoryStateProvinceCode); err != nil {
		return 0, err
	}

	m.CountryCode = make([]byte, 3)
	if err := readLen(buf, m.CountryCode); err != nil {
		return 0, err
	}

	if err := m.PostalZIPCode.Write(buf); err != nil {
		return 0, err
	}

	if err := m.EmailAddress.Write(buf); err != nil {
		return 0, err
	}

	if err := m.PhoneNumber.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.KeyRolesCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.KeyRolesCount); i++ {
		x := &KeyRole{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.KeyRoles = append(m.KeyRoles, *x)
	}

	if err := read(buf, &m.NotableRolesCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.NotableRolesCount); i++ {
		x := &NotableRole{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.NotableRoles = append(m.NotableRoles, *x)
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m ContractOffer) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m ContractOffer) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("ContractName:%#+v", m.ContractName))
	vals = append(vals, fmt.Sprintf("ContractFileType:%v", m.ContractFileType))
	vals = append(vals, fmt.Sprintf("LenContractFile:%v", m.LenContractFile))
	vals = append(vals, fmt.Sprintf("ContractFile:%#x", m.ContractFile))
	vals = append(vals, fmt.Sprintf("GoverningLaw:%#x", m.GoverningLaw))
	vals = append(vals, fmt.Sprintf("Jurisdiction:%#x", m.Jurisdiction))
	vals = append(vals, fmt.Sprintf("ContractExpiration:%v", m.ContractExpiration))
	vals = append(vals, fmt.Sprintf("ContractURI:%#+v", m.ContractURI))
	vals = append(vals, fmt.Sprintf("IssuerName:%#+v", m.IssuerName))
	vals = append(vals, fmt.Sprintf("IssuerType:%#+v", m.IssuerType))
	vals = append(vals, fmt.Sprintf("IssuerLogoURL:%#+v", m.IssuerLogoURL))
	vals = append(vals, fmt.Sprintf("ContractOperatorID:%#+v", m.ContractOperatorID))
	vals = append(vals, fmt.Sprintf("ContractAuthFlags:%#x", m.ContractAuthFlags))
	vals = append(vals, fmt.Sprintf("VotingSystemCount:%v", m.VotingSystemCount))
	vals = append(vals, fmt.Sprintf("VotingSystems:%#+v", m.VotingSystems))
	vals = append(vals, fmt.Sprintf("RestrictedQtyAssets:%v", m.RestrictedQtyAssets))
	vals = append(vals, fmt.Sprintf("ReferendumProposal:%#+v", m.ReferendumProposal))
	vals = append(vals, fmt.Sprintf("InitiativeProposal:%#+v", m.InitiativeProposal))
	vals = append(vals, fmt.Sprintf("RegistryCount:%v", m.RegistryCount))
	vals = append(vals, fmt.Sprintf("Registries:%#+v", m.Registries))
	vals = append(vals, fmt.Sprintf("IssuerAddress:%#+v", m.IssuerAddress))
	vals = append(vals, fmt.Sprintf("UnitNumber:%#+v", m.UnitNumber))
	vals = append(vals, fmt.Sprintf("BuildingNumber:%#+v", m.BuildingNumber))
	vals = append(vals, fmt.Sprintf("Street:%#+v", m.Street))
	vals = append(vals, fmt.Sprintf("SuburbCity:%#+v", m.SuburbCity))
	vals = append(vals, fmt.Sprintf("TerritoryStateProvinceCode:%#x", m.TerritoryStateProvinceCode))
	vals = append(vals, fmt.Sprintf("CountryCode:%#x", m.CountryCode))
	vals = append(vals, fmt.Sprintf("PostalZIPCode:%#+v", m.PostalZIPCode))
	vals = append(vals, fmt.Sprintf("EmailAddress:%#+v", m.EmailAddress))
	vals = append(vals, fmt.Sprintf("PhoneNumber:%#+v", m.PhoneNumber))
	vals = append(vals, fmt.Sprintf("KeyRolesCount:%v", m.KeyRolesCount))
	vals = append(vals, fmt.Sprintf("KeyRoles:%#+v", m.KeyRoles))
	vals = append(vals, fmt.Sprintf("NotableRolesCount:%v", m.NotableRolesCount))
	vals = append(vals, fmt.Sprintf("NotableRoles:%#+v", m.NotableRoles))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// ContractFormation This txn is created by the Contract (smart
// contract/off-chain agent/token contract) upon receipt of a valid
// Contract Offer Action from the issuer. The Smart Contract will execute
// on a server controlled by the Issuer. or a Smart Contract Operator on
// their behalf.
type ContractFormation struct {
	Header                     Header
	ContractName               Nvarchar8
	ContractFileType           uint8
	LenContractFile            uint32
	ContractFile               []byte
	GoverningLaw               []byte
	Jurisdiction               []byte
	ContractExpiration         uint64
	ContractURI                Nvarchar8
	IssuerName                 Nvarchar8
	IssuerType                 byte
	IssuerLogoURL              Nvarchar8
	ContractOperatorID         Nvarchar8
	ContractAuthFlags          []byte
	VotingSystemCount          uint8
	VotingSystems              []VotingSystem
	RestrictedQtyAssets        uint64
	ReferendumProposal         bool
	InitiativeProposal         bool
	RegistryCount              uint8
	Registries                 []Registry
	IssuerAddress              bool
	UnitNumber                 Nvarchar8
	BuildingNumber             Nvarchar8
	Street                     Nvarchar16
	SuburbCity                 Nvarchar8
	TerritoryStateProvinceCode []byte
	CountryCode                []byte
	PostalZIPCode              Nvarchar8
	EmailAddress               Nvarchar8
	PhoneNumber                Nvarchar8
	KeyRolesCount              uint8
	KeyRoles                   []KeyRole
	NotableRolesCount          uint8
	NotableRoles               []NotableRole
	ContractRevision           uint64
	Timestamp                  uint64
}

// NewContractFormation returns a new ContractFormation with defaults set.
func NewContractFormation() ContractFormation {
	return ContractFormation{}
}

// Type returns the type identifer for this message.
func (m ContractFormation) Type() string {
	return CodeContractFormation
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m ContractFormation) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m ContractFormation) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	{
		b, err := m.ContractName.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.ContractFileType); err != nil {
		return nil, err
	}

	if err := write(buf, m.LenContractFile); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.ContractFile, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.GoverningLaw, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.Jurisdiction, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, m.ContractExpiration); err != nil {
		return nil, err
	}

	{
		b, err := m.ContractURI.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.IssuerName.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.IssuerType); err != nil {
		return nil, err
	}

	{
		b, err := m.IssuerLogoURL.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.ContractOperatorID.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, pad(m.ContractAuthFlags, 16)); err != nil {
		return nil, err
	}

	if err := write(buf, m.VotingSystemCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.VotingSystemCount); i++ {
		b, err := m.VotingSystems[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.RestrictedQtyAssets); err != nil {
		return nil, err
	}

	if err := write(buf, m.ReferendumProposal); err != nil {
		return nil, err
	}

	if err := write(buf, m.InitiativeProposal); err != nil {
		return nil, err
	}

	if err := write(buf, m.RegistryCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.RegistryCount); i++ {
		b, err := m.Registries[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.IssuerAddress); err != nil {
		return nil, err
	}

	{
		b, err := m.UnitNumber.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.BuildingNumber.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.Street.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.SuburbCity.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, pad(m.TerritoryStateProvinceCode, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.CountryCode, 3)); err != nil {
		return nil, err
	}

	{
		b, err := m.PostalZIPCode.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.EmailAddress.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.PhoneNumber.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.KeyRolesCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.KeyRolesCount); i++ {
		b, err := m.KeyRoles[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.NotableRolesCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.NotableRolesCount); i++ {
		b, err := m.NotableRoles[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.ContractRevision); err != nil {
		return nil, err
	}

	if err := write(buf, m.Timestamp); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeContractFormation, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *ContractFormation) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := m.ContractName.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ContractFileType); err != nil {
		return 0, err
	}

	if err := read(buf, &m.LenContractFile); err != nil {
		return 0, err
	}

	m.ContractFile = make([]byte, 32)
	if err := readLen(buf, m.ContractFile); err != nil {
		return 0, err
	}

	m.GoverningLaw = make([]byte, 5)
	if err := readLen(buf, m.GoverningLaw); err != nil {
		return 0, err
	}

	m.Jurisdiction = make([]byte, 5)
	if err := readLen(buf, m.Jurisdiction); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ContractExpiration); err != nil {
		return 0, err
	}

	if err := m.ContractURI.Write(buf); err != nil {
		return 0, err
	}

	if err := m.IssuerName.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.IssuerType); err != nil {
		return 0, err
	}

	if err := m.IssuerLogoURL.Write(buf); err != nil {
		return 0, err
	}

	if err := m.ContractOperatorID.Write(buf); err != nil {
		return 0, err
	}

	m.ContractAuthFlags = make([]byte, 16)
	if err := readLen(buf, m.ContractAuthFlags); err != nil {
		return 0, err
	}

	if err := read(buf, &m.VotingSystemCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.VotingSystemCount); i++ {
		x := &VotingSystem{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.VotingSystems = append(m.VotingSystems, *x)
	}

	if err := read(buf, &m.RestrictedQtyAssets); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ReferendumProposal); err != nil {
		return 0, err
	}

	if err := read(buf, &m.InitiativeProposal); err != nil {
		return 0, err
	}

	if err := read(buf, &m.RegistryCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.RegistryCount); i++ {
		x := &Registry{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.Registries = append(m.Registries, *x)
	}

	if err := read(buf, &m.IssuerAddress); err != nil {
		return 0, err
	}

	if err := m.UnitNumber.Write(buf); err != nil {
		return 0, err
	}

	if err := m.BuildingNumber.Write(buf); err != nil {
		return 0, err
	}

	if err := m.Street.Write(buf); err != nil {
		return 0, err
	}

	if err := m.SuburbCity.Write(buf); err != nil {
		return 0, err
	}

	m.TerritoryStateProvinceCode = make([]byte, 5)
	if err := readLen(buf, m.TerritoryStateProvinceCode); err != nil {
		return 0, err
	}

	m.CountryCode = make([]byte, 3)
	if err := readLen(buf, m.CountryCode); err != nil {
		return 0, err
	}

	if err := m.PostalZIPCode.Write(buf); err != nil {
		return 0, err
	}

	if err := m.EmailAddress.Write(buf); err != nil {
		return 0, err
	}

	if err := m.PhoneNumber.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.KeyRolesCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.KeyRolesCount); i++ {
		x := &KeyRole{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.KeyRoles = append(m.KeyRoles, *x)
	}

	if err := read(buf, &m.NotableRolesCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.NotableRolesCount); i++ {
		x := &NotableRole{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.NotableRoles = append(m.NotableRoles, *x)
	}

	if err := read(buf, &m.ContractRevision); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Timestamp); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m ContractFormation) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m ContractFormation) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("ContractName:%#+v", m.ContractName))
	vals = append(vals, fmt.Sprintf("ContractFileType:%v", m.ContractFileType))
	vals = append(vals, fmt.Sprintf("LenContractFile:%v", m.LenContractFile))
	vals = append(vals, fmt.Sprintf("ContractFile:%#x", m.ContractFile))
	vals = append(vals, fmt.Sprintf("GoverningLaw:%#x", m.GoverningLaw))
	vals = append(vals, fmt.Sprintf("Jurisdiction:%#x", m.Jurisdiction))
	vals = append(vals, fmt.Sprintf("ContractExpiration:%v", m.ContractExpiration))
	vals = append(vals, fmt.Sprintf("ContractURI:%#+v", m.ContractURI))
	vals = append(vals, fmt.Sprintf("IssuerName:%#+v", m.IssuerName))
	vals = append(vals, fmt.Sprintf("IssuerType:%#+v", m.IssuerType))
	vals = append(vals, fmt.Sprintf("IssuerLogoURL:%#+v", m.IssuerLogoURL))
	vals = append(vals, fmt.Sprintf("ContractOperatorID:%#+v", m.ContractOperatorID))
	vals = append(vals, fmt.Sprintf("ContractAuthFlags:%#x", m.ContractAuthFlags))
	vals = append(vals, fmt.Sprintf("VotingSystemCount:%v", m.VotingSystemCount))
	vals = append(vals, fmt.Sprintf("VotingSystems:%#+v", m.VotingSystems))
	vals = append(vals, fmt.Sprintf("RestrictedQtyAssets:%v", m.RestrictedQtyAssets))
	vals = append(vals, fmt.Sprintf("ReferendumProposal:%#+v", m.ReferendumProposal))
	vals = append(vals, fmt.Sprintf("InitiativeProposal:%#+v", m.InitiativeProposal))
	vals = append(vals, fmt.Sprintf("RegistryCount:%v", m.RegistryCount))
	vals = append(vals, fmt.Sprintf("Registries:%#+v", m.Registries))
	vals = append(vals, fmt.Sprintf("IssuerAddress:%#+v", m.IssuerAddress))
	vals = append(vals, fmt.Sprintf("UnitNumber:%#+v", m.UnitNumber))
	vals = append(vals, fmt.Sprintf("BuildingNumber:%#+v", m.BuildingNumber))
	vals = append(vals, fmt.Sprintf("Street:%#+v", m.Street))
	vals = append(vals, fmt.Sprintf("SuburbCity:%#+v", m.SuburbCity))
	vals = append(vals, fmt.Sprintf("TerritoryStateProvinceCode:%#x", m.TerritoryStateProvinceCode))
	vals = append(vals, fmt.Sprintf("CountryCode:%#x", m.CountryCode))
	vals = append(vals, fmt.Sprintf("PostalZIPCode:%#+v", m.PostalZIPCode))
	vals = append(vals, fmt.Sprintf("EmailAddress:%#+v", m.EmailAddress))
	vals = append(vals, fmt.Sprintf("PhoneNumber:%#+v", m.PhoneNumber))
	vals = append(vals, fmt.Sprintf("KeyRolesCount:%v", m.KeyRolesCount))
	vals = append(vals, fmt.Sprintf("KeyRoles:%#+v", m.KeyRoles))
	vals = append(vals, fmt.Sprintf("NotableRolesCount:%v", m.NotableRolesCount))
	vals = append(vals, fmt.Sprintf("NotableRoles:%#+v", m.NotableRoles))
	vals = append(vals, fmt.Sprintf("ContractRevision:%v", m.ContractRevision))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// ContractAmendment Contract Amendment Action - the issuer can initiate an
// amendment to the contract establishment metadata. The ability to make an
// amendment to the contract is restricted by the Authorization Flag set on
// the current revision of Contract Formation action.
type ContractAmendment struct {
	Header                Header
	ChangeIssuerAddress   bool
	ChangeOperatorAddress bool
	ContractRevision      uint16
	AmendmentsCount       uint8
	Amendments            []Amendment
	RefTxID               []byte
}

// NewContractAmendment returns a new ContractAmendment with defaults set.
func NewContractAmendment() ContractAmendment {
	return ContractAmendment{}
}

// Type returns the type identifer for this message.
func (m ContractAmendment) Type() string {
	return CodeContractAmendment
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m ContractAmendment) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m ContractAmendment) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.ChangeIssuerAddress); err != nil {
		return nil, err
	}

	if err := write(buf, m.ChangeOperatorAddress); err != nil {
		return nil, err
	}

	if err := write(buf, m.ContractRevision); err != nil {
		return nil, err
	}

	if err := write(buf, m.AmendmentsCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.AmendmentsCount); i++ {
		b, err := m.Amendments[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, pad(m.RefTxID, 32)); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeContractAmendment, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *ContractAmendment) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ChangeIssuerAddress); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ChangeOperatorAddress); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ContractRevision); err != nil {
		return 0, err
	}

	if err := read(buf, &m.AmendmentsCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.AmendmentsCount); i++ {
		x := &Amendment{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.Amendments = append(m.Amendments, *x)
	}

	m.RefTxID = make([]byte, 32)
	if err := readLen(buf, m.RefTxID); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m ContractAmendment) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m ContractAmendment) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("ChangeIssuerAddress:%#+v", m.ChangeIssuerAddress))
	vals = append(vals, fmt.Sprintf("ChangeOperatorAddress:%#+v", m.ChangeOperatorAddress))
	vals = append(vals, fmt.Sprintf("ContractRevision:%v", m.ContractRevision))
	vals = append(vals, fmt.Sprintf("AmendmentsCount:%v", m.AmendmentsCount))
	vals = append(vals, fmt.Sprintf("Amendments:%#+v", m.Amendments))
	vals = append(vals, fmt.Sprintf("RefTxID:%#x", m.RefTxID))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// StaticContractFormation Static Contract Formation Action
type StaticContractFormation struct {
	Header             Header
	ContractName       Nvarchar8
	ContractType       Nvarchar8
	ContractFileType   uint8
	LenContractFile    uint32
	ContractFile       []byte
	ContractRevision   uint16
	GoverningLaw       []byte
	Jurisdiction       []byte
	EffectiveDate      uint64
	ContractExpiration uint64
	ContractURI        Nvarchar8
	PrevRevTxID        []byte
	EntityCount        uint8
	Entities           []Entity
}

// NewStaticContractFormation returns a new StaticContractFormation with defaults set.
func NewStaticContractFormation() StaticContractFormation {
	return StaticContractFormation{}
}

// Type returns the type identifer for this message.
func (m StaticContractFormation) Type() string {
	return CodeStaticContractFormation
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m StaticContractFormation) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m StaticContractFormation) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	{
		b, err := m.ContractName.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.ContractType.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.ContractFileType); err != nil {
		return nil, err
	}

	if err := write(buf, m.LenContractFile); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.ContractFile, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, m.ContractRevision); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.GoverningLaw, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.Jurisdiction, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, m.EffectiveDate); err != nil {
		return nil, err
	}

	if err := write(buf, m.ContractExpiration); err != nil {
		return nil, err
	}

	{
		b, err := m.ContractURI.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, pad(m.PrevRevTxID, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, m.EntityCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.EntityCount); i++ {
		b, err := m.Entities[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeStaticContractFormation, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *StaticContractFormation) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := m.ContractName.Write(buf); err != nil {
		return 0, err
	}

	if err := m.ContractType.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ContractFileType); err != nil {
		return 0, err
	}

	if err := read(buf, &m.LenContractFile); err != nil {
		return 0, err
	}

	m.ContractFile = make([]byte, 32)
	if err := readLen(buf, m.ContractFile); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ContractRevision); err != nil {
		return 0, err
	}

	m.GoverningLaw = make([]byte, 5)
	if err := readLen(buf, m.GoverningLaw); err != nil {
		return 0, err
	}

	m.Jurisdiction = make([]byte, 5)
	if err := readLen(buf, m.Jurisdiction); err != nil {
		return 0, err
	}

	if err := read(buf, &m.EffectiveDate); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ContractExpiration); err != nil {
		return 0, err
	}

	if err := m.ContractURI.Write(buf); err != nil {
		return 0, err
	}

	m.PrevRevTxID = make([]byte, 32)
	if err := readLen(buf, m.PrevRevTxID); err != nil {
		return 0, err
	}

	if err := read(buf, &m.EntityCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.EntityCount); i++ {
		x := &Entity{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.Entities = append(m.Entities, *x)
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m StaticContractFormation) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m StaticContractFormation) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("ContractName:%#+v", m.ContractName))
	vals = append(vals, fmt.Sprintf("ContractType:%#+v", m.ContractType))
	vals = append(vals, fmt.Sprintf("ContractFileType:%v", m.ContractFileType))
	vals = append(vals, fmt.Sprintf("LenContractFile:%v", m.LenContractFile))
	vals = append(vals, fmt.Sprintf("ContractFile:%#x", m.ContractFile))
	vals = append(vals, fmt.Sprintf("ContractRevision:%v", m.ContractRevision))
	vals = append(vals, fmt.Sprintf("GoverningLaw:%#x", m.GoverningLaw))
	vals = append(vals, fmt.Sprintf("Jurisdiction:%#x", m.Jurisdiction))
	vals = append(vals, fmt.Sprintf("EffectiveDate:%v", m.EffectiveDate))
	vals = append(vals, fmt.Sprintf("ContractExpiration:%v", m.ContractExpiration))
	vals = append(vals, fmt.Sprintf("ContractURI:%#+v", m.ContractURI))
	vals = append(vals, fmt.Sprintf("PrevRevTxID:%#x", m.PrevRevTxID))
	vals = append(vals, fmt.Sprintf("EntityCount:%v", m.EntityCount))
	vals = append(vals, fmt.Sprintf("Entities:%#+v", m.Entities))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Order Order Action - Issuer to signal to the smart contract that the
// tokens that a particular public address(es) owns are to be confiscated,
// frozen, thawed or reconciled.
type Order struct {
	Header                 Header
	AssetType              []byte
	AssetID                []byte
	ComplianceAction       byte
	TargetAddressCount     uint16
	TargetAddresses        []TargetAddress
	DepositAddress         Nvarchar8
	AuthorityName          Nvarchar8
	SigAlgoAddressList     uint8
	AuthorityPublicKey     Nvarchar8
	OrderSignature         Nvarchar8
	SupportingEvidenceHash []byte
	RefTxnID               []byte
	FreezePeriod           uint64
	Message                Nvarchar32
}

// NewOrder returns a new Order with defaults set.
func NewOrder() Order {
	return Order{}
}

// Type returns the type identifer for this message.
func (m Order) Type() string {
	return CodeOrder
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Order) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Order) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, pad(m.AssetType, 3)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AssetID, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, m.ComplianceAction); err != nil {
		return nil, err
	}

	if err := write(buf, m.TargetAddressCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.TargetAddressCount); i++ {
		b, err := m.TargetAddresses[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.DepositAddress.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.AuthorityName.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.SigAlgoAddressList); err != nil {
		return nil, err
	}

	{
		b, err := m.AuthorityPublicKey.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.OrderSignature.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, pad(m.SupportingEvidenceHash, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.RefTxnID, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, m.FreezePeriod); err != nil {
		return nil, err
	}

	{
		b, err := m.Message.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeOrder, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *Order) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	m.AssetType = make([]byte, 3)
	if err := readLen(buf, m.AssetType); err != nil {
		return 0, err
	}

	m.AssetID = make([]byte, 32)
	if err := readLen(buf, m.AssetID); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ComplianceAction); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TargetAddressCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.TargetAddressCount); i++ {
		x := &TargetAddress{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.TargetAddresses = append(m.TargetAddresses, *x)
	}

	if err := m.DepositAddress.Write(buf); err != nil {
		return 0, err
	}

	if err := m.AuthorityName.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.SigAlgoAddressList); err != nil {
		return 0, err
	}

	if err := m.AuthorityPublicKey.Write(buf); err != nil {
		return 0, err
	}

	if err := m.OrderSignature.Write(buf); err != nil {
		return 0, err
	}

	m.SupportingEvidenceHash = make([]byte, 32)
	if err := readLen(buf, m.SupportingEvidenceHash); err != nil {
		return 0, err
	}

	m.RefTxnID = make([]byte, 32)
	if err := readLen(buf, m.RefTxnID); err != nil {
		return 0, err
	}

	if err := read(buf, &m.FreezePeriod); err != nil {
		return 0, err
	}

	if err := m.Message.Write(buf); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Order) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Order) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("AssetType:%#x", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#x", m.AssetID))
	vals = append(vals, fmt.Sprintf("ComplianceAction:%#+v", m.ComplianceAction))
	vals = append(vals, fmt.Sprintf("TargetAddressCount:%v", m.TargetAddressCount))
	vals = append(vals, fmt.Sprintf("TargetAddresses:%#+v", m.TargetAddresses))
	vals = append(vals, fmt.Sprintf("DepositAddress:%#+v", m.DepositAddress))
	vals = append(vals, fmt.Sprintf("AuthorityName:%#+v", m.AuthorityName))
	vals = append(vals, fmt.Sprintf("SigAlgoAddressList:%v", m.SigAlgoAddressList))
	vals = append(vals, fmt.Sprintf("AuthorityPublicKey:%#+v", m.AuthorityPublicKey))
	vals = append(vals, fmt.Sprintf("OrderSignature:%#+v", m.OrderSignature))
	vals = append(vals, fmt.Sprintf("SupportingEvidenceHash:%#x", m.SupportingEvidenceHash))
	vals = append(vals, fmt.Sprintf("RefTxnID:%#x", m.RefTxnID))
	vals = append(vals, fmt.Sprintf("FreezePeriod:%v", m.FreezePeriod))
	vals = append(vals, fmt.Sprintf("Message:%#+v", m.Message))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Freeze Freeze Action - To be used to comply with
// contractual/legal/issuer requirements. The target public address(es)
// will be marked as frozen. However the Freeze action publishes this fact
// to the public blockchain for transparency. The Contract will not respond
// to any actions requested by the frozen address.
type Freeze struct {
	Header       Header
	AddressCount uint16
	Addresses    []Address
	Timestamp    uint64
}

// NewFreeze returns a new Freeze with defaults set.
func NewFreeze() Freeze {
	return Freeze{}
}

// Type returns the type identifer for this message.
func (m Freeze) Type() string {
	return CodeFreeze
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Freeze) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Freeze) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.AddressCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.AddressCount); i++ {
		b, err := m.Addresses[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.Timestamp); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeFreeze, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *Freeze) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.AddressCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.AddressCount); i++ {
		x := &Address{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.Addresses = append(m.Addresses, *x)
	}

	if err := read(buf, &m.Timestamp); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Freeze) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Freeze) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("AddressCount:%v", m.AddressCount))
	vals = append(vals, fmt.Sprintf("Addresses:%#+v", m.Addresses))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Thaw Thaw Action - to be used to comply with contractual obligations or
// legal requirements. The Alleged Offender's tokens will be unfrozen to
// allow them to resume normal exchange and governance activities.
type Thaw struct {
	Header       Header
	AddressCount uint16
	Addresses    []Address
	RefTxnID     []byte
	Timestamp    uint64
}

// NewThaw returns a new Thaw with defaults set.
func NewThaw() Thaw {
	return Thaw{}
}

// Type returns the type identifer for this message.
func (m Thaw) Type() string {
	return CodeThaw
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Thaw) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Thaw) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.AddressCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.AddressCount); i++ {
		b, err := m.Addresses[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, pad(m.RefTxnID, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, m.Timestamp); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeThaw, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *Thaw) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.AddressCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.AddressCount); i++ {
		x := &Address{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.Addresses = append(m.Addresses, *x)
	}

	m.RefTxnID = make([]byte, 32)
	if err := readLen(buf, m.RefTxnID); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Timestamp); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Thaw) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Thaw) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("AddressCount:%v", m.AddressCount))
	vals = append(vals, fmt.Sprintf("Addresses:%#+v", m.Addresses))
	vals = append(vals, fmt.Sprintf("RefTxnID:%#x", m.RefTxnID))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Confiscation Confiscation Action - to be used to comply with contractual
// obligations, legal and/or issuer requirements.
type Confiscation struct {
	Header       Header
	AddressCount uint16
	Addresses    []Address
	DepositQty   uint64
	Timestamp    uint64
}

// NewConfiscation returns a new Confiscation with defaults set.
func NewConfiscation() Confiscation {
	return Confiscation{}
}

// Type returns the type identifer for this message.
func (m Confiscation) Type() string {
	return CodeConfiscation
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Confiscation) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Confiscation) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.AddressCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.AddressCount); i++ {
		b, err := m.Addresses[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.DepositQty); err != nil {
		return nil, err
	}

	if err := write(buf, m.Timestamp); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeConfiscation, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *Confiscation) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.AddressCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.AddressCount); i++ {
		x := &Address{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.Addresses = append(m.Addresses, *x)
	}

	if err := read(buf, &m.DepositQty); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Timestamp); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Confiscation) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Confiscation) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("AddressCount:%v", m.AddressCount))
	vals = append(vals, fmt.Sprintf("Addresses:%#+v", m.Addresses))
	vals = append(vals, fmt.Sprintf("DepositQty:%v", m.DepositQty))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Reconciliation Reconciliation Action - to be used at the direction of
// the issuer to fix record keeping errors with bitcoin and token balances.
type Reconciliation struct {
	Header       Header
	AddressCount uint16
	Addresses    []Address
	Timestamp    uint64
}

// NewReconciliation returns a new Reconciliation with defaults set.
func NewReconciliation() Reconciliation {
	return Reconciliation{}
}

// Type returns the type identifer for this message.
func (m Reconciliation) Type() string {
	return CodeReconciliation
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Reconciliation) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Reconciliation) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.AddressCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.AddressCount); i++ {
		b, err := m.Addresses[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.Timestamp); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeReconciliation, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *Reconciliation) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.AddressCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.AddressCount); i++ {
		x := &Address{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.Addresses = append(m.Addresses, *x)
	}

	if err := read(buf, &m.Timestamp); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Reconciliation) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Reconciliation) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("AddressCount:%v", m.AddressCount))
	vals = append(vals, fmt.Sprintf("Addresses:%#+v", m.Addresses))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Initiative Initiative Action - Allows Token Owners to propose an
// Initiative (aka Initiative/Shareholder vote). A significant cost -
// specified in the Contract Formation - can be attached to this action to
// reduce spam, as the resulting vote will be put to all token owners.
type Initiative struct {
	Header               Header
	AssetType            []byte
	AssetID              []byte
	VoteSystem           uint8
	Proposal             bool
	ProposedChangesCount uint8
	ProposedChanges      []Amendment
	VoteOptions          Nvarchar8
	VoteMax              uint8
	ProposalDescription  Nvarchar32
	ProposalDocumentHash []byte
	VoteCutOffTimestamp  uint64
}

// NewInitiative returns a new Initiative with defaults set.
func NewInitiative() Initiative {
	return Initiative{}
}

// Type returns the type identifer for this message.
func (m Initiative) Type() string {
	return CodeInitiative
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Initiative) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Initiative) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, pad(m.AssetType, 3)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AssetID, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, m.VoteSystem); err != nil {
		return nil, err
	}

	if err := write(buf, m.Proposal); err != nil {
		return nil, err
	}

	if err := write(buf, m.ProposedChangesCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.ProposedChangesCount); i++ {
		b, err := m.ProposedChanges[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.VoteOptions.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.VoteMax); err != nil {
		return nil, err
	}

	{
		b, err := m.ProposalDescription.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, pad(m.ProposalDocumentHash, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, m.VoteCutOffTimestamp); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeInitiative, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *Initiative) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	m.AssetType = make([]byte, 3)
	if err := readLen(buf, m.AssetType); err != nil {
		return 0, err
	}

	m.AssetID = make([]byte, 32)
	if err := readLen(buf, m.AssetID); err != nil {
		return 0, err
	}

	if err := read(buf, &m.VoteSystem); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Proposal); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ProposedChangesCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.ProposedChangesCount); i++ {
		x := &Amendment{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.ProposedChanges = append(m.ProposedChanges, *x)
	}

	if err := m.VoteOptions.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.VoteMax); err != nil {
		return 0, err
	}

	if err := m.ProposalDescription.Write(buf); err != nil {
		return 0, err
	}

	m.ProposalDocumentHash = make([]byte, 32)
	if err := readLen(buf, m.ProposalDocumentHash); err != nil {
		return 0, err
	}

	if err := read(buf, &m.VoteCutOffTimestamp); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Initiative) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Initiative) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("AssetType:%#x", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#x", m.AssetID))
	vals = append(vals, fmt.Sprintf("VoteSystem:%v", m.VoteSystem))
	vals = append(vals, fmt.Sprintf("Proposal:%#+v", m.Proposal))
	vals = append(vals, fmt.Sprintf("ProposedChangesCount:%v", m.ProposedChangesCount))
	vals = append(vals, fmt.Sprintf("ProposedChanges:%#+v", m.ProposedChanges))
	vals = append(vals, fmt.Sprintf("VoteOptions:%#+v", m.VoteOptions))
	vals = append(vals, fmt.Sprintf("VoteMax:%v", m.VoteMax))
	vals = append(vals, fmt.Sprintf("ProposalDescription:%#+v", m.ProposalDescription))
	vals = append(vals, fmt.Sprintf("ProposalDocumentHash:%#x", m.ProposalDocumentHash))
	vals = append(vals, fmt.Sprintf("VoteCutOffTimestamp:%v", m.VoteCutOffTimestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Referendum Referendum Action - Issuer instructs the Contract to Initiate
// a Token Owner Vote. Usually used for contract amendments, organizational
// governance, etc.
type Referendum struct {
	Header               Header
	AssetSpecificVote    bool
	AssetType            []byte
	AssetID              []byte
	VoteSystem           uint8
	Proposal             bool
	ProposedChangesCount uint8
	ProposedChanges      []Amendment
	VoteOptions          Nvarchar8
	VoteMax              uint8
	ProposalDescription  Nvarchar32
	ProposalDocumentHash []byte
	VoteCutOffTimestamp  uint64
}

// NewReferendum returns a new Referendum with defaults set.
func NewReferendum() Referendum {
	return Referendum{}
}

// Type returns the type identifer for this message.
func (m Referendum) Type() string {
	return CodeReferendum
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Referendum) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Referendum) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.AssetSpecificVote); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AssetType, 3)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AssetID, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, m.VoteSystem); err != nil {
		return nil, err
	}

	if err := write(buf, m.Proposal); err != nil {
		return nil, err
	}

	if err := write(buf, m.ProposedChangesCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.ProposedChangesCount); i++ {
		b, err := m.ProposedChanges[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	{
		b, err := m.VoteOptions.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.VoteMax); err != nil {
		return nil, err
	}

	{
		b, err := m.ProposalDescription.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, pad(m.ProposalDocumentHash, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, m.VoteCutOffTimestamp); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeReferendum, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *Referendum) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.AssetSpecificVote); err != nil {
		return 0, err
	}

	m.AssetType = make([]byte, 3)
	if err := readLen(buf, m.AssetType); err != nil {
		return 0, err
	}

	m.AssetID = make([]byte, 32)
	if err := readLen(buf, m.AssetID); err != nil {
		return 0, err
	}

	if err := read(buf, &m.VoteSystem); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Proposal); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ProposedChangesCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.ProposedChangesCount); i++ {
		x := &Amendment{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.ProposedChanges = append(m.ProposedChanges, *x)
	}

	if err := m.VoteOptions.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.VoteMax); err != nil {
		return 0, err
	}

	if err := m.ProposalDescription.Write(buf); err != nil {
		return 0, err
	}

	m.ProposalDocumentHash = make([]byte, 32)
	if err := readLen(buf, m.ProposalDocumentHash); err != nil {
		return 0, err
	}

	if err := read(buf, &m.VoteCutOffTimestamp); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Referendum) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Referendum) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("AssetSpecificVote:%#+v", m.AssetSpecificVote))
	vals = append(vals, fmt.Sprintf("AssetType:%#x", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#x", m.AssetID))
	vals = append(vals, fmt.Sprintf("VoteSystem:%v", m.VoteSystem))
	vals = append(vals, fmt.Sprintf("Proposal:%#+v", m.Proposal))
	vals = append(vals, fmt.Sprintf("ProposedChangesCount:%v", m.ProposedChangesCount))
	vals = append(vals, fmt.Sprintf("ProposedChanges:%#+v", m.ProposedChanges))
	vals = append(vals, fmt.Sprintf("VoteOptions:%#+v", m.VoteOptions))
	vals = append(vals, fmt.Sprintf("VoteMax:%v", m.VoteMax))
	vals = append(vals, fmt.Sprintf("ProposalDescription:%#+v", m.ProposalDescription))
	vals = append(vals, fmt.Sprintf("ProposalDocumentHash:%#x", m.ProposalDocumentHash))
	vals = append(vals, fmt.Sprintf("VoteCutOffTimestamp:%v", m.VoteCutOffTimestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Vote Vote Action - A vote is created by the Contract in response to a
// valid Referendum (Issuer) or Initiative (User) Action.
type Vote struct {
	Header    Header
	Timestamp uint64
}

// NewVote returns a new Vote with defaults set.
func NewVote() Vote {
	return Vote{}
}

// Type returns the type identifer for this message.
func (m Vote) Type() string {
	return CodeVote
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Vote) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Vote) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Timestamp); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeVote, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *Vote) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Timestamp); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Vote) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Vote) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// BallotCast Ballot Cast Action - Used by Token Owners to cast their
// ballot (vote) on proposals raised by the Issuer (Referendum) or other
// token holders (Initiative). 1 Vote per token unless a vote multiplier is
// specified in the relevant Asset Definition action.
type BallotCast struct {
	Header    Header
	AssetType []byte
	AssetID   []byte
	VoteTxnID []byte
	Vote      Nvarchar8
}

// NewBallotCast returns a new BallotCast with defaults set.
func NewBallotCast() BallotCast {
	return BallotCast{}
}

// Type returns the type identifer for this message.
func (m BallotCast) Type() string {
	return CodeBallotCast
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m BallotCast) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m BallotCast) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, pad(m.AssetType, 3)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AssetID, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.VoteTxnID, 32)); err != nil {
		return nil, err
	}

	{
		b, err := m.Vote.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeBallotCast, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *BallotCast) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	m.AssetType = make([]byte, 3)
	if err := readLen(buf, m.AssetType); err != nil {
		return 0, err
	}

	m.AssetID = make([]byte, 32)
	if err := readLen(buf, m.AssetID); err != nil {
		return 0, err
	}

	m.VoteTxnID = make([]byte, 32)
	if err := readLen(buf, m.VoteTxnID); err != nil {
		return 0, err
	}

	if err := m.Vote.Write(buf); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m BallotCast) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m BallotCast) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("AssetType:%#x", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#x", m.AssetID))
	vals = append(vals, fmt.Sprintf("VoteTxnID:%#x", m.VoteTxnID))
	vals = append(vals, fmt.Sprintf("Vote:%#+v", m.Vote))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// BallotCounted Ballot Counted Action - The smart contract will respond to
// a Ballot Cast action with a Ballot Counted action if the Ballot Cast is
// valid. If the Ballot Cast is not valid, then the smart contract will
// respond with a Rejection Action.
type BallotCounted struct {
	Header    Header
	Timestamp uint64
}

// NewBallotCounted returns a new BallotCounted with defaults set.
func NewBallotCounted() BallotCounted {
	return BallotCounted{}
}

// Type returns the type identifer for this message.
func (m BallotCounted) Type() string {
	return CodeBallotCounted
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m BallotCounted) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m BallotCounted) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Timestamp); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeBallotCounted, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *BallotCounted) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Timestamp); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m BallotCounted) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m BallotCounted) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Result Result Action - Once a vote has been completed the results are
// published.
type Result struct {
	Header               Header
	AssetType            []byte
	AssetID              []byte
	Proposal             bool
	ProposedChangesCount uint8
	ProposedChanges      []Amendment
	VoteTxnID            []byte
	VoteOptionsCount     uint8
	OptionXTally         uint64
	Result               Nvarchar8
	Timestamp            uint64
}

// NewResult returns a new Result with defaults set.
func NewResult() Result {
	return Result{}
}

// Type returns the type identifer for this message.
func (m Result) Type() string {
	return CodeResult
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Result) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Result) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, pad(m.AssetType, 3)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AssetID, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, m.Proposal); err != nil {
		return nil, err
	}

	if err := write(buf, m.ProposedChangesCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.ProposedChangesCount); i++ {
		b, err := m.ProposedChanges[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, pad(m.VoteTxnID, 32)); err != nil {
		return nil, err
	}

	if err := write(buf, m.VoteOptionsCount); err != nil {
		return nil, err
	}

	if err := write(buf, m.OptionXTally); err != nil {
		return nil, err
	}

	{
		b, err := m.Result.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.Timestamp); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeResult, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *Result) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	m.AssetType = make([]byte, 3)
	if err := readLen(buf, m.AssetType); err != nil {
		return 0, err
	}

	m.AssetID = make([]byte, 32)
	if err := readLen(buf, m.AssetID); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Proposal); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ProposedChangesCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.ProposedChangesCount); i++ {
		x := &Amendment{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.ProposedChanges = append(m.ProposedChanges, *x)
	}

	m.VoteTxnID = make([]byte, 32)
	if err := readLen(buf, m.VoteTxnID); err != nil {
		return 0, err
	}

	if err := read(buf, &m.VoteOptionsCount); err != nil {
		return 0, err
	}

	if err := read(buf, &m.OptionXTally); err != nil {
		return 0, err
	}

	if err := m.Result.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Timestamp); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Result) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Result) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("AssetType:%#x", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#x", m.AssetID))
	vals = append(vals, fmt.Sprintf("Proposal:%#+v", m.Proposal))
	vals = append(vals, fmt.Sprintf("ProposedChangesCount:%v", m.ProposedChangesCount))
	vals = append(vals, fmt.Sprintf("ProposedChanges:%#+v", m.ProposedChanges))
	vals = append(vals, fmt.Sprintf("VoteTxnID:%#x", m.VoteTxnID))
	vals = append(vals, fmt.Sprintf("VoteOptionsCount:%v", m.VoteOptionsCount))
	vals = append(vals, fmt.Sprintf("OptionXTally:%v", m.OptionXTally))
	vals = append(vals, fmt.Sprintf("Result:%#+v", m.Result))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Message Message Action - the message action is a general purpose
// communication action. 'Twitter/SMS' for Issuers/Investors/Users. The
// message txn can also be used for passing partially signed txns on-chain,
// establishing private communication channels and EDI (receipting,
// invoices, PO, and private offers/bids). The messages are broken down by
// type for easy filtering in the a user's wallet. The Message Types are
// listed in the Message Types table.
type Message struct {
	Header                Header
	QtyReceivingAddresses uint8
	AddressIndexes        []uint16
	MessageType           []byte
	MessagePayload        Nvarchar32
}

// NewMessage returns a new Message with defaults set.
func NewMessage() Message {
	return Message{}
}

// Type returns the type identifer for this message.
func (m Message) Type() string {
	return CodeMessage
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Message) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Message) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.QtyReceivingAddresses); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.QtyReceivingAddresses); i++ {
		if err := write(buf, m.AddressIndexes[i]); err != nil {
			return nil, err
		}
	}

	if err := write(buf, pad(m.MessageType, 2)); err != nil {
		return nil, err
	}

	{
		b, err := m.MessagePayload.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeMessage, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *Message) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.QtyReceivingAddresses); err != nil {
		return 0, err
	}

	m.AddressIndexes = make([]uint16, m.QtyReceivingAddresses, m.QtyReceivingAddresses)
	if err := read(buf, &m.AddressIndexes); err != nil {
		return 0, err
	}

	m.MessageType = make([]byte, 2)
	if err := readLen(buf, m.MessageType); err != nil {
		return 0, err
	}

	if err := m.MessagePayload.Write(buf); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Message) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Message) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("QtyReceivingAddresses:%v", m.QtyReceivingAddresses))
	vals = append(vals, fmt.Sprintf("AddressIndexes:%v", m.AddressIndexes))
	vals = append(vals, fmt.Sprintf("MessageType:%#x", m.MessageType))
	vals = append(vals, fmt.Sprintf("MessagePayload:%#+v", m.MessagePayload))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Rejection Rejection Action - used to reject request actions that do not
// comply with the Contract. If money is to be returned to a User then it
// is used in lieu of the Settlement Action to properly account for token
// balances. All Issuer/User request Actions must be responded to by the
// Contract with an Action. The only exception to this rule is when there
// is not enough fees in the first Action for the Contract response action
// to remain revenue neutral. If not enough fees are attached to pay for
// the Contract response then the Contract will not respond.
type Rejection struct {
	Header                Header
	QtyReceivingAddresses uint8
	AddressIndexes        []uint16
	RejectionType         uint8
	MessagePayload        Nvarchar32
	Timestamp             uint64
}

// NewRejection returns a new Rejection with defaults set.
func NewRejection() Rejection {
	return Rejection{}
}

// Type returns the type identifer for this message.
func (m Rejection) Type() string {
	return CodeRejection
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Rejection) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Rejection) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.QtyReceivingAddresses); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.QtyReceivingAddresses); i++ {
		if err := write(buf, m.AddressIndexes[i]); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.RejectionType); err != nil {
		return nil, err
	}

	{
		b, err := m.MessagePayload.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.Timestamp); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeRejection, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *Rejection) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.QtyReceivingAddresses); err != nil {
		return 0, err
	}

	m.AddressIndexes = make([]uint16, m.QtyReceivingAddresses, m.QtyReceivingAddresses)
	if err := read(buf, &m.AddressIndexes); err != nil {
		return 0, err
	}

	if err := read(buf, &m.RejectionType); err != nil {
		return 0, err
	}

	if err := m.MessagePayload.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Timestamp); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Rejection) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Rejection) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("QtyReceivingAddresses:%v", m.QtyReceivingAddresses))
	vals = append(vals, fmt.Sprintf("AddressIndexes:%v", m.AddressIndexes))
	vals = append(vals, fmt.Sprintf("RejectionType:%v", m.RejectionType))
	vals = append(vals, fmt.Sprintf("MessagePayload:%#+v", m.MessagePayload))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Establishment Establishment Action - Establishes an on-chain register.
type Establishment struct {
	Header  Header
	Message Nvarchar32
}

// NewEstablishment returns a new Establishment with defaults set.
func NewEstablishment() Establishment {
	return Establishment{}
}

// Type returns the type identifer for this message.
func (m Establishment) Type() string {
	return CodeEstablishment
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Establishment) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Establishment) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	{
		b, err := m.Message.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeEstablishment, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *Establishment) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := m.Message.Write(buf); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Establishment) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Establishment) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("Message:%#+v", m.Message))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Addition Addition Action - Adds an entry to the Register.
type Addition struct {
	Header  Header
	Message Nvarchar32
}

// NewAddition returns a new Addition with defaults set.
func NewAddition() Addition {
	return Addition{}
}

// Type returns the type identifer for this message.
func (m Addition) Type() string {
	return CodeAddition
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Addition) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Addition) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	{
		b, err := m.Message.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeAddition, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *Addition) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := m.Message.Write(buf); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Addition) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Addition) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("Message:%#+v", m.Message))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Alteration Alteration Action - A register entry/record can be altered.
type Alteration struct {
	Header  Header
	Message Nvarchar32
}

// NewAlteration returns a new Alteration with defaults set.
func NewAlteration() Alteration {
	return Alteration{}
}

// Type returns the type identifer for this message.
func (m Alteration) Type() string {
	return CodeAlteration
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Alteration) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Alteration) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	{
		b, err := m.Message.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeAlteration, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *Alteration) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := m.Message.Write(buf); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Alteration) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Alteration) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("Message:%#+v", m.Message))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Removal Removal Action - Removes an entry/record from the Register.
type Removal struct {
	Header  Header
	Message Nvarchar32
}

// NewRemoval returns a new Removal with defaults set.
func NewRemoval() Removal {
	return Removal{}
}

// Type returns the type identifer for this message.
func (m Removal) Type() string {
	return CodeRemoval
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Removal) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Removal) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	{
		b, err := m.Message.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeRemoval, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *Removal) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := m.Message.Write(buf); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Removal) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Removal) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("Message:%#+v", m.Message))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Transfer A Token Owner(s) Sends, Excahnges or Swaps a token(s) or
// Bitcoin for a token(s) or Bitcoin. Can be as simple as sending a single
// token to a receiver. Or can be as complex as many senders sending many
// different assets - controlled by many different smart contracts - to a
// number of receivers. This action also supports atomic swaps (tokens for
// tokens). Since many parties and contracts can be involved in a transfer
// and the corresponding settlement action, the partially signed T1 and T2
// actions will need to be passed around on-chain with an M1 action, or
// off-chain.
type Transfer struct {
	Header              Header
	AssetTransferCount  uint8
	AssetTransfers      []AssetTransfer
	OfferExpiry         uint64
	ExchangeFeeCurrency []byte
	ExchangeFeeVar      float32
	ExchangeFeeFixed    float32
	ExchangeFeeAddress  []byte
}

// NewTransfer returns a new Transfer with defaults set.
func NewTransfer() Transfer {
	return Transfer{}
}

// Type returns the type identifer for this message.
func (m Transfer) Type() string {
	return CodeTransfer
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Transfer) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Transfer) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.AssetTransferCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.AssetTransferCount); i++ {
		b, err := m.AssetTransfers[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.OfferExpiry); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.ExchangeFeeCurrency, 3)); err != nil {
		return nil, err
	}

	if err := write(buf, m.ExchangeFeeVar); err != nil {
		return nil, err
	}

	if err := write(buf, m.ExchangeFeeFixed); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.ExchangeFeeAddress, 34)); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeTransfer, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *Transfer) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.AssetTransferCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.AssetTransferCount); i++ {
		x := &AssetTransfer{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.AssetTransfers = append(m.AssetTransfers, *x)
	}

	if err := read(buf, &m.OfferExpiry); err != nil {
		return 0, err
	}

	m.ExchangeFeeCurrency = make([]byte, 3)
	if err := readLen(buf, m.ExchangeFeeCurrency); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ExchangeFeeVar); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ExchangeFeeFixed); err != nil {
		return 0, err
	}

	m.ExchangeFeeAddress = make([]byte, 34)
	if err := readLen(buf, m.ExchangeFeeAddress); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Transfer) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Transfer) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("AssetTransferCount:%v", m.AssetTransferCount))
	vals = append(vals, fmt.Sprintf("AssetTransfers:%#+v", m.AssetTransfers))
	vals = append(vals, fmt.Sprintf("OfferExpiry:%v", m.OfferExpiry))
	vals = append(vals, fmt.Sprintf("ExchangeFeeCurrency:%#x", m.ExchangeFeeCurrency))
	vals = append(vals, fmt.Sprintf("ExchangeFeeVar:%v", m.ExchangeFeeVar))
	vals = append(vals, fmt.Sprintf("ExchangeFeeFixed:%v", m.ExchangeFeeFixed))
	vals = append(vals, fmt.Sprintf("ExchangeFeeAddress:%#x", m.ExchangeFeeAddress))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Settlement Settlement Action - Settles the transfer request of bitcoins
// and tokens from transfer (T1) actions.
type Settlement struct {
	Header               Header
	AssetSettlementCount uint8
	AssetSettlements     []AssetSettlement
	Timestamp            uint64
}

// NewSettlement returns a new Settlement with defaults set.
func NewSettlement() Settlement {
	return Settlement{}
}

// Type returns the type identifer for this message.
func (m Settlement) Type() string {
	return CodeSettlement
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Settlement) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Settlement) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.AssetSettlementCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.AssetSettlementCount); i++ {
		b, err := m.AssetSettlements[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.Timestamp); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode(CodeSettlement, len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m *Settlement) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	if err := read(buf, &m.AssetSettlementCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.AssetSettlementCount); i++ {
		x := &AssetSettlement{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.AssetSettlements = append(m.AssetSettlements, *x)
	}

	if err := read(buf, &m.Timestamp); err != nil {
		return 0, err
	}

	return len(b), nil
}

// PayloadMessage returns the PayloadMessage, if any.
func (m Settlement) PayloadMessage() (PayloadMessage, error) {
	return nil, nil
}

func (m Settlement) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Header:%#+v", m.Header))
	vals = append(vals, fmt.Sprintf("AssetSettlementCount:%v", m.AssetSettlementCount))
	vals = append(vals, fmt.Sprintf("AssetSettlements:%#+v", m.AssetSettlements))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
