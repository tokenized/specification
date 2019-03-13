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
	Header                      Header  // Common header data for all actions
	AssetType                   string  // eg. Share - Common
	AssetID                     string  // Randomly generated base58 string.  Each Asset ID should be unique.  However, an Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type + Asset ID = Asset Code.  An Asset Code is a human readable idenitfier that can be used in a similar way to a Bitcoin (BSV) address, a vanity identifying label.
	AssetAuthFlags              []byte  // Authorization Flags,  bitwise operation
	TransfersPermitted          bool    // 1 = Transfers are permitted.  0 = Transfers are not permitted.
	TradeRestrictions           string  // Asset can only be traded within the trade restrictions.  Eg. AUS - Australian residents only.  EU - European Union residents only.
	EnforcementOrdersPermitted  bool    // 1 = Enforcement Orders are permitted. 0 = Enforcement Orders are not permitted.
	VoteMultiplier              uint8   // Multiplies the vote by the integer. 1 token = 1 vote with a 1 for vote multipler (normal).  1 token = 3 votes with a multiplier of 3, for example.
	ReferendumProposal          bool    // A Referendum is permitted for Asset-Wide Proposals (outside of smart contract scope) if also permitted by the contract. If the contract has proposals by referendum restricted, then this flag is meaningless.
	InitiativeProposal          bool    // An initiative is permitted for Asset-Wide Proposals (outside of smart contract scope) if also permitted by the contract. If the contract has proposals by initiative restricted, then this flag is meaningless.
	AssetModificationGovernance bool    // 1 - Contract-wide Asset Governance.  0 - Asset-wide Asset Governance.  If a referendum or initiative is used to propose a modification to a subfield controlled by the asset auth flags, then the vote will either be a contract-wide vote (all assets vote on the referendum/initiative) or an asset-wide vote (all assets vote on the referendum/initiative) depending on the value in this subfield.  The voting system specifies the voting rules.
	TokenQty                    uint64  // Quantity of token - 0 is valid. Fungible 'shares' of the Asset. 1 is used for non-fungible tokens.  Asset IDs become the non-fungible Asset ID and many Asset IDs can be associated with a particular Contract.
	ContractFeeCurrency         string  // BSV, USD, AUD, EUR, etc.
	ContractFeeVar              float32 // Percent of the value of the transaction
	ContractFeeFixed            float32 // Fixed fee (payment made in BSV)
	AssetPayloadLen             uint16  // Size of the asset payload in bytes.
	AssetPayload                []byte  // Payload length is dependent on the asset type. Each asset is made up of a defined set of information pertaining to the specific asset type, and may contain fields of variable length type (nvarchar8, 16, 32)
}

// NewAssetDefinition returns a new empty AssetDefinition.
func NewEmptyAssetDefinition() *AssetDefinition {
	result := AssetDefinition{}
	return &result
}

// NewAssetDefinition returns a new AssetDefinition with specified values set.
func NewAssetDefinition(assetType string) *AssetDefinition {
	result := AssetDefinition{}
	result.AssetType = assetType
	return &result
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

	if err := WriteFixedChar(buf, m.AssetType, 3); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.AssetID, 32); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AssetAuthFlags, 8)); err != nil {
		return nil, err
	}

	if err := write(buf, m.TransfersPermitted); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.TradeRestrictions, 3); err != nil {
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

	if err := WriteFixedChar(buf, m.ContractFeeCurrency, 3); err != nil {
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

	header, err := NewHeaderForCode([]byte(CodeAssetDefinition), len(b))
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

	{
		var err error
		m.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.AssetID, err = ReadFixedChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	m.AssetAuthFlags = make([]byte, 8)
	if err := readLen(buf, m.AssetAuthFlags); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TransfersPermitted); err != nil {
		return 0, err
	}

	{
		var err error
		m.TradeRestrictions, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
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

	{
		var err error
		m.ContractFeeCurrency, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
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
	p, err := NewPayloadMessageFromCode([]byte(m.AssetType))
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
	vals = append(vals, fmt.Sprintf("AssetType:%#+v", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#+v", m.AssetID))
	vals = append(vals, fmt.Sprintf("AssetAuthFlags:%#x", m.AssetAuthFlags))
	vals = append(vals, fmt.Sprintf("TransfersPermitted:%#+v", m.TransfersPermitted))
	vals = append(vals, fmt.Sprintf("TradeRestrictions:%#+v", m.TradeRestrictions))
	vals = append(vals, fmt.Sprintf("EnforcementOrdersPermitted:%#+v", m.EnforcementOrdersPermitted))
	vals = append(vals, fmt.Sprintf("VoteMultiplier:%v", m.VoteMultiplier))
	vals = append(vals, fmt.Sprintf("ReferendumProposal:%#+v", m.ReferendumProposal))
	vals = append(vals, fmt.Sprintf("InitiativeProposal:%#+v", m.InitiativeProposal))
	vals = append(vals, fmt.Sprintf("AssetModificationGovernance:%#+v", m.AssetModificationGovernance))
	vals = append(vals, fmt.Sprintf("TokenQty:%v", m.TokenQty))
	vals = append(vals, fmt.Sprintf("ContractFeeCurrency:%#+v", m.ContractFeeCurrency))
	vals = append(vals, fmt.Sprintf("ContractFeeVar:%v", m.ContractFeeVar))
	vals = append(vals, fmt.Sprintf("ContractFeeFixed:%v", m.ContractFeeFixed))
	vals = append(vals, fmt.Sprintf("AssetPayloadLen:%v", m.AssetPayloadLen))
	vals = append(vals, fmt.Sprintf("AssetPayload:%#x", m.AssetPayload))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// AssetCreation Asset Creation Action - This action creates an Asset in
// response to the Issuer's instructions in the Definition Action.
type AssetCreation struct {
	Header                      Header  // Common header data for all actions
	AssetType                   string  // eg. Share - Common
	AssetID                     string  // Randomly generated base58 string.  Each Asset ID should be unique.  However, a Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type + Asset ID = Asset Code.  An Asset Code is a human readable idenitfier that can be used in a similar way to a Bitcoin (BSV) address, a vanity identifying label.
	AssetAuthFlags              []byte  // Authorization Flags,  bitwise operation
	TransfersPermitted          bool    // 1 = Transfers are permitted.  0 = Transfers are not permitted.
	TradeRestrictions           string  // Asset can only be traded within the trade restrictions.  Eg. AUS - Australian residents only.  EU - European Union residents only.
	EnforcementOrdersPermitted  bool    // 1 = Enforcement Orders are permitted. 0 = Enforcement Orders are not permitted.
	VoteMultiplier              uint8   // Multiplies the vote by the integer. 1 token = 1 vote with a 1 for vote multipler (normal).  1 token = 3 votes with a multiplier of 3, for example.
	ReferendumProposal          bool    // A Referendum is permitted for Asset-Wide Proposals (outside of smart contract scope) if also permitted by the contract. If the contract has proposals by referendum restricted, then this flag is meaningless.
	InitiativeProposal          bool    // An initiative is permitted for Asset-Wide Proposals (outside of smart contract scope) if also permitted by the contract. If the contract has proposals by initiative restricted, then this flag is meaningless.
	AssetModificationGovernance bool    // 1 - Contract-wide Asset Governance.  0 - Asset-wide Asset Governance.  If a referendum or initiative is used to propose a modification to a subfield controlled by the asset auth flags, then the vote will either be a contract-wide vote (all assets vote on the referendum/initiative) or an asset-wide vote (all assets vote on the referendum/initiative).  The voting system specifies the voting rules.
	TokenQty                    uint64  // Quantity of token - 0 is valid. Fungible 'shares' of the Asset. 1 is used for non-fungible tokens.  Asset IDs become the non-fungible Asset ID and many Asset IDs can be associated with a particular Contract.
	ContractFeeCurrency         string  // BSV, USD, AUD, EUR, etc.
	ContractFeeVar              float32 // Percent of the value of the transaction
	ContractFeeFixed            float32 // Fixed fee (payment made in BSV)
	AssetPayloadLen             uint16  // Size of the asset payload in bytes.
	AssetPayload                []byte  // Payload length is dependent on the asset type. Each asset is made up of a defined set of information pertaining to the specific asset type, and may contain fields of variable length type (nvarchar8, 16, 32)
	AssetRevision               uint64  // Counter 0 - 65,535
	Timestamp                   uint64  // Timestamp in nanoseconds of when the smart contract created the action.
}

// NewAssetCreation returns a new empty AssetCreation.
func NewEmptyAssetCreation() *AssetCreation {
	result := AssetCreation{}
	return &result
}

// NewAssetCreation returns a new AssetCreation with specified values set.
func NewAssetCreation() *AssetCreation {
	result := AssetCreation{}
	return &result
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

	if err := WriteFixedChar(buf, m.AssetType, 3); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.AssetID, 32); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AssetAuthFlags, 8)); err != nil {
		return nil, err
	}

	if err := write(buf, m.TransfersPermitted); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.TradeRestrictions, 3); err != nil {
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

	if err := WriteFixedChar(buf, m.ContractFeeCurrency, 3); err != nil {
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

	header, err := NewHeaderForCode([]byte(CodeAssetCreation), len(b))
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

	{
		var err error
		m.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.AssetID, err = ReadFixedChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	m.AssetAuthFlags = make([]byte, 8)
	if err := readLen(buf, m.AssetAuthFlags); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TransfersPermitted); err != nil {
		return 0, err
	}

	{
		var err error
		m.TradeRestrictions, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
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

	{
		var err error
		m.ContractFeeCurrency, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
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
	p, err := NewPayloadMessageFromCode([]byte(m.AssetType))
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
	vals = append(vals, fmt.Sprintf("AssetType:%#+v", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#+v", m.AssetID))
	vals = append(vals, fmt.Sprintf("AssetAuthFlags:%#x", m.AssetAuthFlags))
	vals = append(vals, fmt.Sprintf("TransfersPermitted:%#+v", m.TransfersPermitted))
	vals = append(vals, fmt.Sprintf("TradeRestrictions:%#+v", m.TradeRestrictions))
	vals = append(vals, fmt.Sprintf("EnforcementOrdersPermitted:%#+v", m.EnforcementOrdersPermitted))
	vals = append(vals, fmt.Sprintf("VoteMultiplier:%v", m.VoteMultiplier))
	vals = append(vals, fmt.Sprintf("ReferendumProposal:%#+v", m.ReferendumProposal))
	vals = append(vals, fmt.Sprintf("InitiativeProposal:%#+v", m.InitiativeProposal))
	vals = append(vals, fmt.Sprintf("AssetModificationGovernance:%#+v", m.AssetModificationGovernance))
	vals = append(vals, fmt.Sprintf("TokenQty:%v", m.TokenQty))
	vals = append(vals, fmt.Sprintf("ContractFeeCurrency:%#+v", m.ContractFeeCurrency))
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
	Header            Header      // Common header data for all actions
	AssetType         string      // eg. Share - Common
	AssetID           string      // Randomly generated base58 string.  Each Asset ID should be unique.  However, a Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type + Asset ID = Asset Code.  An Asset Code is a human readable idenitfier that can be used in a similar way to a Bitcoin (BSV) address, a vanity identifying label.
	AssetRevision     uint64      // Counter. (Subfield cannot be manually changed by Asset Modification Action.  Only SC can increment by 1 with each AC action. SC will reject AM actions where the wrong asset revision has been selected.
	ModificationCount uint8       // Number of Modifications. Must be less than the max Subfield Index of CF.
	Modifications     []Amendment //
	RefTxID           [32]byte    // Tx-ID of the associated Result action (governance) that permitted the modifications.
}

// NewAssetModification returns a new empty AssetModification.
func NewEmptyAssetModification() *AssetModification {
	result := AssetModification{}
	return &result
}

// NewAssetModification returns a new AssetModification with specified values set.
func NewAssetModification() *AssetModification {
	result := AssetModification{}
	return &result
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

	if err := WriteFixedChar(buf, m.AssetType, 3); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.AssetID, 32); err != nil {
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

	if err := write(buf, m.RefTxID); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode([]byte(CodeAssetModification), len(b))
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

	{
		var err error
		m.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.AssetID, err = ReadFixedChar(buf, 32)
		if err != nil {
			return 0, err
		}
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

	if err := read(buf, &m.RefTxID); err != nil {
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
	vals = append(vals, fmt.Sprintf("AssetType:%#+v", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#+v", m.AssetID))
	vals = append(vals, fmt.Sprintf("AssetRevision:%v", m.AssetRevision))
	vals = append(vals, fmt.Sprintf("ModificationCount:%v", m.ModificationCount))
	vals = append(vals, fmt.Sprintf("Modifications:%#+v", m.Modifications))
	vals = append(vals, fmt.Sprintf("RefTxID:%#+v", m.RefTxID))

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
	Header                     Header         // Common header data for all actions
	ContractName               string         // Can be any unique identifying string, including human readable names for branding/vanity purposes.   [Contract identifier (instance) is the bitcoin public key hash address. If the Public Address is lost, then the issuer will have to reissue the entire contract, Asset definition and tokens with the new public address.]. Smart contracts can be branded and specialized to suit any terms and conditions.
	ContractFileType           uint8          // 1 - SHA-256 Hash, 2 - Markdown
	LenContractFile            uint32         // Max size is the max transaction size - other data in the txn.
	ContractFile               string         // SHA-256 hash of the contract file or markdown data for contract file specific to the smart contract and relevant Assets.  Legal and technical information. (eg. pdf)
	GoverningLaw               string         // 5 Letter Code to Identify which governing law the contract will adhere to.  Disputes are to be settled by this law in the jurisdiction specified below. Private dispute resolution organizations can be used as well.  A custom code just needs to be defined.
	Jurisdiction               string         // Legal proceedings/arbitration will take place using the specified Governing Law in this location.
	ContractExpiration         uint64         // All actions related to the contract will cease to work after this timestamp. The smart contract will stop running.  This will allow many token use cases to be able to calculate total smart contract running costs for the entire life of the contract. Eg. an issuer is creating tickets for an event on the 5th of June 2018.  The smart contract will facilitate exchange and send transactions up until the 6th of June.  Wallets can use this to forget tokens that are no longer valid - or at least store them in an 'Expired' folder.
	ContractURI                string         // Points to an information page that also has a copy of the Contract.  Anyone can go to the website to have a look at the price/token, information about the Issuer (company), information about the Asset, legal information, etc.  There will also be a way for Token Owners to vote on this page and contact details with the Issuer/tokenized companies. Could be a IPv6/IPv4, an IPFS address (hash) or txn-id for on-chain information or even a public address (DNS).
	IssuerName                 string         // Length 0-255 bytes. 0 is not valid.Issuing entity (company, organization, individual).  Can be any unique identifying string, including human readable names for branding/vanity purposes.
	IssuerType                 string         // P - Public Company Limited by Shares, C - Private Company Limited by Shares, I - Individual, L - Limited Partnership, U -Unlimited Partnership, T - Sole Proprietorship, S - Statutory Company, O - Non-Profit Organization, N - Nation State, G - Government Agency, U - Unit Trust, D - Discretionary Trust.  Found in 'Entities' (Specification/Resources).
	IssuerLogoURL              string         // The URL of the Issuers logo.
	ContractOperatorID         string         // Length 0-255 bytes. 0 is valid. Smart Contract Operator identifier. Can be any unique identifying string, including human readable names for branding/vanity purposes. Can also be null or the Issuer.
	ContractAuthFlags          []byte         // Authorization Flags aka Terms and Conditions that the smart contract can enforce.  Other terms and conditions that are out of the smart contract's control are listed in the actual Contract File.
	VotingSystemCount          uint8          // 0-255 voting systems. If 0, Voting System and associated subfields (InitiativeThreshold, InitiativeThresholdCurrency) will be null.
	VotingSystems              []VotingSystem // A list of voting systems.
	RestrictedQtyAssets        uint64         // Number of Assets (non-fungible) permitted on this contract. 0 if unlimited which will display an infinity symbol in UI
	ReferendumProposal         bool           // A Referendum is permitted for Proposals (outside of smart contract scope).
	InitiativeProposal         bool           // An initiative is permitted for Proposals (outside of smart contract scope).
	RegistryCount              uint8          // Number of registries (eg. KYC registry/database/whitelist/identity database/etc - managed by a Registrar (oracle)) the smart contract is permitted to interact with. 0-255. 0 is valid (no registry subfields).
	Registries                 []Registry     // A list Registries
	IssuerAddress              bool           // Physical/mailing address. Y/N, N means there is no issuer address.
	UnitNumber                 string         // Issuer Address Details (eg. HQ)
	BuildingNumber             string         //
	Street                     string         //
	SuburbCity                 string         //
	TerritoryStateProvinceCode string         //
	CountryCode                string         //
	PostalZIPCode              string         //
	EmailAddress               string         // Length 0-255 bytes. 0 is valid (no ContactAddress). Address for text-based communication: eg. email address, Bitcoin address
	PhoneNumber                string         // Length 0-50 bytes. 0 is valid (no Phone subfield).Phone Number for Entity.
	KeyRolesCount              uint8          // Number of key roles associated with the issuing entity.  (eg. Directors, etc.) 0-255. 0 is valid.
	KeyRoles                   []KeyRole      // A list of Key Roles.
	NotableRolesCount          uint8          // Number of notable roles associated with the issuing entity.  (eg. Corporate Officers, Managers, etc.) 0-255. 0 is valid.
	NotableRoles               []NotableRole  // A list of Notable Roles.
}

// NewContractOffer returns a new empty ContractOffer.
func NewEmptyContractOffer() *ContractOffer {
	result := ContractOffer{}
	return &result
}

// NewContractOffer returns a new ContractOffer with specified values set.
func NewContractOffer(name string, issuerName string) *ContractOffer {
	result := ContractOffer{}
	result.ContractName = name
	result.IssuerName = issuerName
	return &result
}

// SetIssuerAddress Sets the issuer's mailing address on the contract..
func (action *ContractOffer) SetIssuerAddress(unit string, building string, street string, city string, state string, countryCode string, postalCode string) {
	action.UnitNumber = unit
	action.BuildingNumber = building
	action.Street = street
	action.SuburbCity = city
	action.TerritoryStateProvinceCode = state
	action.CountryCode = countryCode
	action.PostalZIPCode = postalCode
	action.IssuerAddress = true
}

// AddKeyRole Adds a key role to the contract.
func (action *ContractOffer) AddKeyRole(roleType uint8, name string) {
	action.KeyRoles = append(action.KeyRoles, *NewKeyRole(roleType, name))
}

// AddNotableRole Adds a notable role to the contract.
func (action *ContractOffer) AddNotableRole(roleType uint8, name string) {
	action.NotableRoles = append(action.NotableRoles, *NewNotableRole(roleType, name))
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

	if err := WriteVarChar(buf, m.ContractName, 255); err != nil {
		return nil, err
	}

	if err := write(buf, m.ContractFileType); err != nil {
		return nil, err
	}

	if err := write(buf, m.LenContractFile); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.ContractFile, 4294967295); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.GoverningLaw, 5); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.Jurisdiction, 5); err != nil {
		return nil, err
	}

	if err := write(buf, m.ContractExpiration); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.ContractURI, 255); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.IssuerName, 255); err != nil {
		return nil, err
	}

	if err := write(buf, m.IssuerType); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.IssuerLogoURL, 255); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.ContractOperatorID, 255); err != nil {
		return nil, err
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

	if err := WriteVarChar(buf, m.UnitNumber, 255); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.BuildingNumber, 255); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.Street, 65535); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.SuburbCity, 255); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.TerritoryStateProvinceCode, 5); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.CountryCode, 3); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.PostalZIPCode, 255); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.EmailAddress, 255); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.PhoneNumber, 255); err != nil {
		return nil, err
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

	header, err := NewHeaderForCode([]byte(CodeContractOffer), len(b))
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

	{
		var err error
		m.ContractName, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	if err := read(buf, &m.ContractFileType); err != nil {
		return 0, err
	}

	if err := read(buf, &m.LenContractFile); err != nil {
		return 0, err
	}

	{
		var err error
		m.ContractFile, err = ReadVarChar(buf, 4294967295)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.GoverningLaw, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.Jurisdiction, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	if err := read(buf, &m.ContractExpiration); err != nil {
		return 0, err
	}

	{
		var err error
		m.ContractURI, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.IssuerName, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	if err := read(buf, &m.IssuerType); err != nil {
		return 0, err
	}

	{
		var err error
		m.IssuerLogoURL, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.ContractOperatorID, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
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

	{
		var err error
		m.UnitNumber, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.BuildingNumber, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.Street, err = ReadVarChar(buf, 65535)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.SuburbCity, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.TerritoryStateProvinceCode, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.CountryCode, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.PostalZIPCode, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.EmailAddress, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.PhoneNumber, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
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
	vals = append(vals, fmt.Sprintf("ContractFile:%#+v", m.ContractFile))
	vals = append(vals, fmt.Sprintf("GoverningLaw:%#+v", m.GoverningLaw))
	vals = append(vals, fmt.Sprintf("Jurisdiction:%#+v", m.Jurisdiction))
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
	vals = append(vals, fmt.Sprintf("TerritoryStateProvinceCode:%#+v", m.TerritoryStateProvinceCode))
	vals = append(vals, fmt.Sprintf("CountryCode:%#+v", m.CountryCode))
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
	Header                     Header         // Common header data for all actions
	ContractName               string         // Can be any unique identifying string, including human readable names for branding/vanity purposes.   [Contract identifier (instance) is the bitcoin public key hash address. If the Public Address is lost, then the issuer will have to reissue the entire contract, Asset definition and tokens with the new public address.]. Smart contracts can be branded and specialized to suit any terms and conditions.
	ContractFileType           uint8          // 1 - SHA-256 Hash, 2 - Markdown file
	LenContractFile            uint32         // Max size is the max transaction size - other data in the txn.
	ContractFile               [32]byte       // SHA-256 hash of the Contract file specific to the smart contract and relevant Assets.  Legal and technical information. (eg. pdf)
	GoverningLaw               string         // 5 Letter Code to Identify which governing law the contract will adhere to.  Disputes are to be settled by this law in the jurisdiction specified below. Private dispute resolution organizations can be used as well.  A custom code just needs to be defined.
	Jurisdiction               string         // Legal proceedings/arbitration will take place using the specified Governing Law in this location.
	ContractExpiration         uint64         // All actions related to the contract will cease to work after this timestamp. The smart contract will stop running.  This will allow many token use cases to be able to calculate smart contract running costs. Eg. an issuer is creating tickets for an event on the 5th of June 2018.  The smart contract will facilitate exchange and send transactions up until the 6th of June.  Wallets can use this to forget tokens that are no longer valid - or at least store them in an 'Expired' folder.
	ContractURI                string         // Length 0-255 bytes.  0 is valid. Points to an information page that also has a copy of the Contract.  Anyone can go to the website to have a look at the price/token, information about the Issuer (company), information about the Asset, legal information, etc.  There will also be a way for Token Owners to vote on this page and contact details with the Issuer/tokenized companies. Could be a IPv6/IPv4, an IPFS address (hash) or txn-id for on chain information or even a public address (DNS).
	IssuerName                 string         // Length 0-255 bytes. 0 is not valid. Issuing entity (company, organization, individual).  Can be any unique identifying string, including human readable names for branding/vanity purposes.
	IssuerType                 byte           // P - Public Company Limited by Shares, C - Private Company Limited by Shares, I - Individual, L - Limited Partnership, U -Unlimited Partnership, T - Sole Proprietorship, S - Statutory Company, O - Non-Profit Organization, N - Nation State, G - Government Agency, U - Unit Trust, D - Discretionary Trust.  Found in 'Entities' (Specification/Resources).
	IssuerLogoURL              string         // The URL of the Issuers logo.
	ContractOperatorID         string         // Length 0-255 bytes. 0 is valid. Smart Contract Operator identifier. Can be any unique identifying string, including human readable names for branding/vanity purposes. Can also be null or the Issuer.
	ContractAuthFlags          []byte         // Authorization Flags aka Terms and Conditions that the smart contract can enforce.  Other terms and conditions that are out of the smart contract's control are listed in the actual Contract File.
	VotingSystemCount          uint8          // 0-255 voting systems. If 0, Voting System and associated subfields (InitiativeThreshold, InitiativeThresholdCurrency) will be null.
	VotingSystems              []VotingSystem // A list voting systems.
	RestrictedQtyAssets        uint64         // Number of Assets (non-fungible) permitted on this contract. 0 if unlimited which will display an infinity symbol in UI
	ReferendumProposal         bool           // A Referendum is permitted for Contract-Wide Proposals (outside of smart contract scope).
	InitiativeProposal         bool           // An initiative is permitted for Contract-Wide Proposals (outside of smart contract scope).
	RegistryCount              uint8          // Number of registries (eg. KYC registry/database/whitelist/identity database/etc - managed by a Registrar (oracle)) the smart contract is permitted to interact with. 0-255. 0 is valid (no registry subfields).
	Registries                 []Registry     // A list Registries
	IssuerAddress              bool           // Physical/mailing address. Y/N, N means there is no issuer address.
	UnitNumber                 string         // Issuer Address Details (eg. HQ)
	BuildingNumber             string         //
	Street                     string         //
	SuburbCity                 string         //
	TerritoryStateProvinceCode string         //
	CountryCode                string         //
	PostalZIPCode              string         //
	EmailAddress               string         // Address for text-based communication: eg. email address, Bitcoin address
	PhoneNumber                string         // Phone Number for Entity. Max acceptable size: 50.
	KeyRolesCount              uint8          // Number of key roles associated with the issuing entity.  (eg. Directors, etc.) 0-255. 0 is valid.
	KeyRoles                   []KeyRole      // A list of Key Roles.
	NotableRolesCount          uint8          // Number of notable roles associated with the issuing entity.  (eg. Corporate Officers, Managers, etc.) 0-255. 0 is valid.
	NotableRoles               []NotableRole  // A list of Notable Roles.
	ContractRevision           uint64         // Counter. Cannot be manually changed by issuer.  Can only be incremented by 1 by SC when CF action is published.
	Timestamp                  uint64         // Timestamp in nanoseconds of when the smart contract created the action.
}

// NewContractFormation returns a new empty ContractFormation.
func NewEmptyContractFormation() *ContractFormation {
	result := ContractFormation{}
	return &result
}

// NewContractFormation returns a new ContractFormation with specified values set.
func NewContractFormation() *ContractFormation {
	result := ContractFormation{}
	return &result
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

	if err := WriteVarChar(buf, m.ContractName, 255); err != nil {
		return nil, err
	}

	if err := write(buf, m.ContractFileType); err != nil {
		return nil, err
	}

	if err := write(buf, m.LenContractFile); err != nil {
		return nil, err
	}

	if err := write(buf, m.ContractFile); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.GoverningLaw, 5); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.Jurisdiction, 5); err != nil {
		return nil, err
	}

	if err := write(buf, m.ContractExpiration); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.ContractURI, 255); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.IssuerName, 255); err != nil {
		return nil, err
	}

	if err := write(buf, m.IssuerType); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.IssuerLogoURL, 255); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.ContractOperatorID, 255); err != nil {
		return nil, err
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

	if err := WriteVarChar(buf, m.UnitNumber, 255); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.BuildingNumber, 255); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.Street, 65535); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.SuburbCity, 255); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.TerritoryStateProvinceCode, 5); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.CountryCode, 3); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.PostalZIPCode, 255); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.EmailAddress, 255); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.PhoneNumber, 255); err != nil {
		return nil, err
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

	header, err := NewHeaderForCode([]byte(CodeContractFormation), len(b))
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

	{
		var err error
		m.ContractName, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	if err := read(buf, &m.ContractFileType); err != nil {
		return 0, err
	}

	if err := read(buf, &m.LenContractFile); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ContractFile); err != nil {
		return 0, err
	}

	{
		var err error
		m.GoverningLaw, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.Jurisdiction, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	if err := read(buf, &m.ContractExpiration); err != nil {
		return 0, err
	}

	{
		var err error
		m.ContractURI, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.IssuerName, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	if err := read(buf, &m.IssuerType); err != nil {
		return 0, err
	}

	{
		var err error
		m.IssuerLogoURL, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.ContractOperatorID, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
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

	{
		var err error
		m.UnitNumber, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.BuildingNumber, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.Street, err = ReadVarChar(buf, 65535)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.SuburbCity, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.TerritoryStateProvinceCode, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.CountryCode, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.PostalZIPCode, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.EmailAddress, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.PhoneNumber, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
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
	vals = append(vals, fmt.Sprintf("ContractFile:%#+v", m.ContractFile))
	vals = append(vals, fmt.Sprintf("GoverningLaw:%#+v", m.GoverningLaw))
	vals = append(vals, fmt.Sprintf("Jurisdiction:%#+v", m.Jurisdiction))
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
	vals = append(vals, fmt.Sprintf("TerritoryStateProvinceCode:%#+v", m.TerritoryStateProvinceCode))
	vals = append(vals, fmt.Sprintf("CountryCode:%#+v", m.CountryCode))
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
	Header                Header      // Common header data for all actions
	ChangeIssuerAddress   bool        // 1 - Yes, 0 - No.  Used to change the issuer address.  The new issuer address must be in the input[1] position.
	ChangeOperatorAddress bool        // 1 - Yes, 0 - No.  Used to change the smart contract operator address.  The new operator address must be in the input[1] position.
	ContractRevision      uint16      // Counter 0 - 65,535
	AmendmentsCount       uint8       // Number of Amendments. Must be less than the max Subfield Index of CF.
	Amendments            []Amendment //
	RefTxID               [32]byte    // Tx-ID of the associated Result action (governance) that permitted the modifications.
}

// NewContractAmendment returns a new empty ContractAmendment.
func NewEmptyContractAmendment() *ContractAmendment {
	result := ContractAmendment{}
	return &result
}

// NewContractAmendment returns a new ContractAmendment with specified values set.
func NewContractAmendment() *ContractAmendment {
	result := ContractAmendment{}
	return &result
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

	if err := write(buf, m.RefTxID); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode([]byte(CodeContractAmendment), len(b))
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

	if err := read(buf, &m.RefTxID); err != nil {
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
	vals = append(vals, fmt.Sprintf("RefTxID:%#+v", m.RefTxID))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// StaticContractFormation Static Contract Formation Action
type StaticContractFormation struct {
	Header             Header   // Common header data for all actions
	ContractName       string   // Length 0-255 bytes. Can be any unique identifying string, including human readable names for branding/vanity purposes.   [Contract identifier (instance) is the bitcoin public address. If the Public Address is lost, then the issuer will have to reissue the entire contract, Asset definition and tokens with the new public address.]. Smart contracts can be branded and specialized to suit any terms and conditions.
	ContractType       string   //
	ContractFileType   uint8    // 1 - SHA-256 Hash, 2 - Markdown file
	LenContractFile    uint32   // Max size is the max transaction size - other data in the txn.
	ContractFile       string   // SHA-256 hash of the Contract file specific to the smart contract and relevant Assets.  Legal and technical information. (eg. pdf)
	ContractRevision   uint16   // Counter 0 - 65,535
	GoverningLaw       string   // 5 Letter Code to Identify which governing law the contract will adhere to.  Disputes are to be settled by this law in the jurisdiction specified below. Private dispute resolution organizations can be used as well.  A custom code just needs to be defined.
	Jurisdiction       string   // Legal proceedings/arbitration will take place using the specified Governing Law in this location.
	EffectiveDate      uint64   // Start date of the contract.
	ContractExpiration uint64   // All actions related to the contract will cease to work after this timestamp. The smart contract will stop running.  This will allow many token use cases to be able to calculate smart contract running costs. Eg. an issuer is creating tickets for an event on the 5th of June 2018.  The smart contract will facilitate exchange and send transactions up until the 6th of June.  Wallets can use this to forget tokens that are no longer valid - or at least store them in an 'Expired' folder.
	ContractURI        string   // Length 0-255 bytes. Points to an information page that also has a copy of the Contract.  Anyone can go to the website to have a look at the price/token, information about the Issuer (company), information about the Asset, legal information, etc.  There will also be a way for Token Owners to vote on this page and contact details with the Issuer/tokenized companies. Could be a IPv6/IPv4, an IPFS address (hash) or txn-id for on chain information or even a public address (DNS).
	PrevRevTxID        [32]byte // The Tx-ID of the previous contract revision.
	EntityCount        uint8    // Number of entities involved in the contract as contracting parties.
	Entities           []Entity //
}

// NewStaticContractFormation returns a new empty StaticContractFormation.
func NewEmptyStaticContractFormation() *StaticContractFormation {
	result := StaticContractFormation{}
	return &result
}

// NewStaticContractFormation returns a new StaticContractFormation with specified values set.
func NewStaticContractFormation() *StaticContractFormation {
	result := StaticContractFormation{}
	return &result
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

	if err := WriteVarChar(buf, m.ContractName, 30); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.ContractType, 30); err != nil {
		return nil, err
	}

	if err := write(buf, m.ContractFileType); err != nil {
		return nil, err
	}

	if err := write(buf, m.LenContractFile); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.ContractFile, 32); err != nil {
		return nil, err
	}

	if err := write(buf, m.ContractRevision); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.GoverningLaw, 5); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.Jurisdiction, 5); err != nil {
		return nil, err
	}

	if err := write(buf, m.EffectiveDate); err != nil {
		return nil, err
	}

	if err := write(buf, m.ContractExpiration); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.ContractURI, 53); err != nil {
		return nil, err
	}

	if err := write(buf, m.PrevRevTxID); err != nil {
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

	header, err := NewHeaderForCode([]byte(CodeStaticContractFormation), len(b))
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

	{
		var err error
		m.ContractName, err = ReadVarChar(buf, 30)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.ContractType, err = ReadVarChar(buf, 30)
		if err != nil {
			return 0, err
		}
	}

	if err := read(buf, &m.ContractFileType); err != nil {
		return 0, err
	}

	if err := read(buf, &m.LenContractFile); err != nil {
		return 0, err
	}

	{
		var err error
		m.ContractFile, err = ReadFixedChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	if err := read(buf, &m.ContractRevision); err != nil {
		return 0, err
	}

	{
		var err error
		m.GoverningLaw, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.Jurisdiction, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	if err := read(buf, &m.EffectiveDate); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ContractExpiration); err != nil {
		return 0, err
	}

	{
		var err error
		m.ContractURI, err = ReadVarChar(buf, 53)
		if err != nil {
			return 0, err
		}
	}

	if err := read(buf, &m.PrevRevTxID); err != nil {
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
	vals = append(vals, fmt.Sprintf("ContractFile:%#+v", m.ContractFile))
	vals = append(vals, fmt.Sprintf("ContractRevision:%v", m.ContractRevision))
	vals = append(vals, fmt.Sprintf("GoverningLaw:%#+v", m.GoverningLaw))
	vals = append(vals, fmt.Sprintf("Jurisdiction:%#+v", m.Jurisdiction))
	vals = append(vals, fmt.Sprintf("EffectiveDate:%v", m.EffectiveDate))
	vals = append(vals, fmt.Sprintf("ContractExpiration:%v", m.ContractExpiration))
	vals = append(vals, fmt.Sprintf("ContractURI:%#+v", m.ContractURI))
	vals = append(vals, fmt.Sprintf("PrevRevTxID:%#+v", m.PrevRevTxID))
	vals = append(vals, fmt.Sprintf("EntityCount:%v", m.EntityCount))
	vals = append(vals, fmt.Sprintf("Entities:%#+v", m.Entities))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Order Order Action - Issuer to signal to the smart contract that the
// tokens that a particular public address(es) owns are to be confiscated,
// frozen, thawed or reconciled.
type Order struct {
	Header                 Header          // Common header data for all actions
	AssetType              string          // eg. Share, Bond, Ticket
	AssetID                string          // Randomly generated base58 string.  Each Asset ID should be unique.  However, a Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type can be the leading bytes - a convention - to make it easy to identify that it is a token by humans.
	ComplianceAction       byte            // Freeze (F), Thaw (T), Confiscate (C), Reconciliation (R)
	TargetAddressCount     uint16          // 0 - 65,535
	TargetAddresses        []TargetAddress //
	DepositAddress         string          // Length 1-255 bytes. The public address for confiscated tokens to be deposited in.  Null for Freeze, Thaw, actions. For Reconciliation actions the deposit address is who receives bitcoin.
	AuthorityName          string          // Length 0-255 bytes. Enforcement Authority Name (eg. Issuer, Queensland Police Service, Tokenized, etc.)
	SigAlgoAddressList     uint8           // 0 = No Registry-signed Message, 1 = ECDSA+secp256k1
	AuthorityPublicKey     string          // Length 0-255 bytes. Public Key associated with the Enforcement Authority
	OrderSignature         string          // Length 0-255 bytes. Signature for a message that lists out the target addresses and deposit address. Signature of (Contract Address, Asset Code, Compliance Action, Supporting Evidence Hash, Time Out Expiration, TargetAddress1, TargetAddress1Qty, TargetAddressX, TargetAddressXQty,...,DepositAddress)
	SupportingEvidenceHash [32]byte        // SHA-256: warrant, court order, etc.
	RefTxnID               [32]byte        // The settlement action that was dropped from the network.  Not applicable for Freeze, Thaw, and Confiscation orders.  Only applicable for reconcilliation actions.  No subfield when F, T, R is selected as the Compliance Action subfield.
	FreezePeriod           uint64          // Used for a 'time out'.  Tokens are automatically unfrozen after the expiration timestamp without requiring a Thaw Action. Null value for Thaw, Confiscation and Reconciallitaion orders.
	Message                string          // Length only limited by the Bitcoin protocol.
}

// NewOrder returns a new empty Order.
func NewEmptyOrder() *Order {
	result := Order{}
	return &result
}

// NewOrder returns a new Order with specified values set.
func NewOrder() *Order {
	result := Order{}
	return &result
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

	if err := WriteFixedChar(buf, m.AssetType, 3); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.AssetID, 32); err != nil {
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

	if err := WriteVarChar(buf, m.DepositAddress, 255); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.AuthorityName, 255); err != nil {
		return nil, err
	}

	if err := write(buf, m.SigAlgoAddressList); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.AuthorityPublicKey, 255); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.OrderSignature, 255); err != nil {
		return nil, err
	}

	if err := write(buf, m.SupportingEvidenceHash); err != nil {
		return nil, err
	}

	if err := write(buf, m.RefTxnID); err != nil {
		return nil, err
	}

	if err := write(buf, m.FreezePeriod); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.Message, 4294967295); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode([]byte(CodeOrder), len(b))
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

	{
		var err error
		m.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.AssetID, err = ReadFixedChar(buf, 32)
		if err != nil {
			return 0, err
		}
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

	{
		var err error
		m.DepositAddress, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.AuthorityName, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	if err := read(buf, &m.SigAlgoAddressList); err != nil {
		return 0, err
	}

	{
		var err error
		m.AuthorityPublicKey, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.OrderSignature, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	if err := read(buf, &m.SupportingEvidenceHash); err != nil {
		return 0, err
	}

	if err := read(buf, &m.RefTxnID); err != nil {
		return 0, err
	}

	if err := read(buf, &m.FreezePeriod); err != nil {
		return 0, err
	}

	{
		var err error
		m.Message, err = ReadVarChar(buf, 4294967295)
		if err != nil {
			return 0, err
		}
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
	vals = append(vals, fmt.Sprintf("AssetType:%#+v", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#+v", m.AssetID))
	vals = append(vals, fmt.Sprintf("ComplianceAction:%#+v", m.ComplianceAction))
	vals = append(vals, fmt.Sprintf("TargetAddressCount:%v", m.TargetAddressCount))
	vals = append(vals, fmt.Sprintf("TargetAddresses:%#+v", m.TargetAddresses))
	vals = append(vals, fmt.Sprintf("DepositAddress:%#+v", m.DepositAddress))
	vals = append(vals, fmt.Sprintf("AuthorityName:%#+v", m.AuthorityName))
	vals = append(vals, fmt.Sprintf("SigAlgoAddressList:%v", m.SigAlgoAddressList))
	vals = append(vals, fmt.Sprintf("AuthorityPublicKey:%#+v", m.AuthorityPublicKey))
	vals = append(vals, fmt.Sprintf("OrderSignature:%#+v", m.OrderSignature))
	vals = append(vals, fmt.Sprintf("SupportingEvidenceHash:%#+v", m.SupportingEvidenceHash))
	vals = append(vals, fmt.Sprintf("RefTxnID:%#+v", m.RefTxnID))
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
	Header       Header    // Common header data for all actions
	AddressCount uint16    // 0 - 65,535
	Addresses    []Address // Addresses holding tokens to be frozen.
	Timestamp    uint64    // Timestamp in nanoseconds of when the smart contract created the action.
}

// NewFreeze returns a new empty Freeze.
func NewEmptyFreeze() *Freeze {
	result := Freeze{}
	return &result
}

// NewFreeze returns a new Freeze with specified values set.
func NewFreeze() *Freeze {
	result := Freeze{}
	return &result
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

	header, err := NewHeaderForCode([]byte(CodeFreeze), len(b))
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
	Header       Header    // Common header data for all actions
	AddressCount uint16    // 0 - 65,535
	Addresses    []Address // Addresses holding tokens to be thawed.
	RefTxnID     [32]byte  // The related freeze action.
	Timestamp    uint64    // Timestamp in nanoseconds of when the smart contract created the action.
}

// NewThaw returns a new empty Thaw.
func NewEmptyThaw() *Thaw {
	result := Thaw{}
	return &result
}

// NewThaw returns a new Thaw with specified values set.
func NewThaw() *Thaw {
	result := Thaw{}
	return &result
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

	if err := write(buf, m.RefTxnID); err != nil {
		return nil, err
	}

	if err := write(buf, m.Timestamp); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode([]byte(CodeThaw), len(b))
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

	if err := read(buf, &m.RefTxnID); err != nil {
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
	vals = append(vals, fmt.Sprintf("RefTxnID:%#+v", m.RefTxnID))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Confiscation Confiscation Action - to be used to comply with contractual
// obligations, legal and/or issuer requirements.
type Confiscation struct {
	Header       Header    // Common header data for all actions
	AddressCount uint16    // 0 - 65,535
	Addresses    []Address // Addresses holding tokens to be confiscated.
	DepositQty   uint64    // Custodian's token balance after confiscation.
	Timestamp    uint64    // Timestamp in nanoseconds of when the smart contract created the action.
}

// NewConfiscation returns a new empty Confiscation.
func NewEmptyConfiscation() *Confiscation {
	result := Confiscation{}
	return &result
}

// NewConfiscation returns a new Confiscation with specified values set.
func NewConfiscation() *Confiscation {
	result := Confiscation{}
	return &result
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

	header, err := NewHeaderForCode([]byte(CodeConfiscation), len(b))
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
	Header       Header    // Common header data for all actions
	AddressCount uint16    // 0 - 65,535
	Addresses    []Address // Addresses holding tokens to be reconciled.
	Timestamp    uint64    // Timestamp in nanoseconds of when the smart contract created the action.
}

// NewReconciliation returns a new empty Reconciliation.
func NewEmptyReconciliation() *Reconciliation {
	result := Reconciliation{}
	return &result
}

// NewReconciliation returns a new Reconciliation with specified values set.
func NewReconciliation() *Reconciliation {
	result := Reconciliation{}
	return &result
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

	header, err := NewHeaderForCode([]byte(CodeReconciliation), len(b))
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
	Header               Header      // Common header data for all actions
	AssetType            string      // eg. Share, Bond, Ticket
	AssetID              string      // Randomly generated base58 string.  Each Asset ID should be unique.  However, an Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type can be the leading bytes - a convention - to make it easy to identify that it is a token by humans.
	VoteSystem           uint8       // X for Vote System X. (1-255, 0 is not valid.)
	Proposal             bool        // 1 for a Proposal, 0 for an initiative that is requesting changes to specific subfields for modification. If this field is true, the subfields should be empty.  The smart contract cannot interpret the results of a vote when Proposal = 1.  All meaning is interpreted by the token owners and smart contract simply facilates the record keeping.  When Proposal = 0, the smart contract always assumes the first choice is a 'yes', or 'pass', if the threshold is met, and will process the proposed changes accordingly.
	ProposedChangesCount uint8       //
	ProposedChanges      []Amendment // Each element contains details of which fields to modify, or delete. Because the number of fields in a Contract and Asset is dynamic due to some fields being able to be repeated, the index value of the field needs to be calculated against the Contract or Asset the changes are to apply to. In the event of a Vote being created from this Initiative, the changes will be applied to the version of the Contract or Asset at that time.
	VoteOptions          string      // Length 1-255 bytes. 0 is not valid. Each byte allows for a different vote option.  Typical votes will likely be multiple choice or Y/N. Vote instances are identified by the Tx-ID. AB000000000 would be chosen for Y/N (binary) type votes.
	VoteMax              uint8       // Range: 1-X. How many selections can a voter make in a Ballot Cast.  1 is selected for Y/N (binary)
	ProposalDescription  string      // Length restricted by the Bitcoin protocol. 0 is valid. Description or details of the vote
	ProposalDocumentHash [32]byte    // Hash of the proposal document to be distributed to voters.
	VoteCutOffTimestamp  uint64      // Ballot casts after this timestamp will not be included. The vote has finished.
}

// NewInitiative returns a new empty Initiative.
func NewEmptyInitiative() *Initiative {
	result := Initiative{}
	return &result
}

// NewInitiative returns a new Initiative with specified values set.
func NewInitiative() *Initiative {
	result := Initiative{}
	return &result
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

	if err := WriteFixedChar(buf, m.AssetType, 3); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.AssetID, 32); err != nil {
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

	if err := WriteVarChar(buf, m.VoteOptions, 255); err != nil {
		return nil, err
	}

	if err := write(buf, m.VoteMax); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.ProposalDescription, 4294967295); err != nil {
		return nil, err
	}

	if err := write(buf, m.ProposalDocumentHash); err != nil {
		return nil, err
	}

	if err := write(buf, m.VoteCutOffTimestamp); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode([]byte(CodeInitiative), len(b))
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

	{
		var err error
		m.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.AssetID, err = ReadFixedChar(buf, 32)
		if err != nil {
			return 0, err
		}
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

	{
		var err error
		m.VoteOptions, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	if err := read(buf, &m.VoteMax); err != nil {
		return 0, err
	}

	{
		var err error
		m.ProposalDescription, err = ReadVarChar(buf, 4294967295)
		if err != nil {
			return 0, err
		}
	}

	if err := read(buf, &m.ProposalDocumentHash); err != nil {
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
	vals = append(vals, fmt.Sprintf("AssetType:%#+v", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#+v", m.AssetID))
	vals = append(vals, fmt.Sprintf("VoteSystem:%v", m.VoteSystem))
	vals = append(vals, fmt.Sprintf("Proposal:%#+v", m.Proposal))
	vals = append(vals, fmt.Sprintf("ProposedChangesCount:%v", m.ProposedChangesCount))
	vals = append(vals, fmt.Sprintf("ProposedChanges:%#+v", m.ProposedChanges))
	vals = append(vals, fmt.Sprintf("VoteOptions:%#+v", m.VoteOptions))
	vals = append(vals, fmt.Sprintf("VoteMax:%v", m.VoteMax))
	vals = append(vals, fmt.Sprintf("ProposalDescription:%#+v", m.ProposalDescription))
	vals = append(vals, fmt.Sprintf("ProposalDocumentHash:%#+v", m.ProposalDocumentHash))
	vals = append(vals, fmt.Sprintf("VoteCutOffTimestamp:%v", m.VoteCutOffTimestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Referendum Referendum Action - Issuer instructs the Contract to Initiate
// a Token Owner Vote. Usually used for contract amendments, organizational
// governance, etc.
type Referendum struct {
	Header               Header      // Common header data for all actions
	TextEncoding         uint8       // 0 = ASCII, 1 = UTF-8, 2 = UTF-16, 3 = Unicode.  Encoding applies to all 'text' data types. All 'string' types will always be encoded with ASCII.  Where string is selected, all fields will be ASCII.
	AssetSpecificVote    bool        // 1 - Yes, 0 - No.  No Asset Type/AssetID subfields for N - No.
	AssetType            string      // eg. Share, Bond, Ticket
	AssetID              string      // Randomly generated base58 string.  Each Asset ID should be unique.  However, an Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type can be the leading bytes - a convention - to make it easy to identify that it is a token by humans.
	VoteSystem           uint8       // X for Vote System X. (1-255, 0 is not valid.)
	Proposal             bool        // 1 for a Proposal, 0 for an initiative that is requesting changes to specific subfields for modification. If this field is true, the subfields should be empty.  The smart contract cannot interpret the results of a vote when Proposal = 1.  All meaning is interpreted by the token owners and smart contract simply facilates the record keeping.  When Proposal = 0, the smart contract always assumes the first choice is a 'yes', or 'pass', if the threshold is met, and will process the proposed changes accordingly.
	ProposedChangesCount uint8       //
	ProposedChanges      []Amendment // Each element contains details of which fields to modify, or delete. Because the number of fields in a Contract and Asset is dynamic due to some fields being able to be repeated, the index value of the field needs to be calculated against the Contract or Asset the changes are to apply to. In the event of a Vote being created from this Initiative, the changes will be applied to the version of the Contract or Asset at that time.
	VoteOptions          string      // Length 1-255 bytes. 0 is not valid. Each byte allows for a different vote option.  Typical votes will likely be multiple choice or Y/N. Vote instances are identified by the Tx-ID. AB000000000 would be chosen for Y/N (binary) type votes. Only applicable if Proposal Type is set to P for Proposal.  All other Proposal Types will be binary.  Pass/Fail.
	VoteMax              uint8       // Range: 1-15. How many selections can a voter make in a Ballot Cast.  1 is selected for Y/N (binary)
	ProposalDescription  string      // Length restricted by the Bitcoin protocol. 0 is valid. Description of the vote.
	ProposalDocumentHash [32]byte    // Hash of the proposal document to be distributed to voters
	VoteCutOffTimestamp  uint64      // Ballot casts after this timestamp will not be included. The vote has finished.
}

// NewReferendum returns a new empty Referendum.
func NewEmptyReferendum() *Referendum {
	result := Referendum{}
	return &result
}

// NewReferendum returns a new Referendum with specified values set.
func NewReferendum() *Referendum {
	result := Referendum{}
	return &result
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

	if err := write(buf, m.TextEncoding); err != nil {
		return nil, err
	}

	if err := write(buf, m.AssetSpecificVote); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.AssetType, 3); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.AssetID, 32); err != nil {
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

	if err := WriteVarChar(buf, m.VoteOptions, 255); err != nil {
		return nil, err
	}

	if err := write(buf, m.VoteMax); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.ProposalDescription, 4294967295); err != nil {
		return nil, err
	}

	if err := write(buf, m.ProposalDocumentHash); err != nil {
		return nil, err
	}

	if err := write(buf, m.VoteCutOffTimestamp); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode([]byte(CodeReferendum), len(b))
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

	if err := read(buf, &m.TextEncoding); err != nil {
		return 0, err
	}

	if err := read(buf, &m.AssetSpecificVote); err != nil {
		return 0, err
	}

	{
		var err error
		m.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.AssetID, err = ReadFixedChar(buf, 32)
		if err != nil {
			return 0, err
		}
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

	{
		var err error
		m.VoteOptions, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
	}

	if err := read(buf, &m.VoteMax); err != nil {
		return 0, err
	}

	{
		var err error
		m.ProposalDescription, err = ReadVarChar(buf, 4294967295)
		if err != nil {
			return 0, err
		}
	}

	if err := read(buf, &m.ProposalDocumentHash); err != nil {
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
	vals = append(vals, fmt.Sprintf("TextEncoding:%v", m.TextEncoding))
	vals = append(vals, fmt.Sprintf("AssetSpecificVote:%#+v", m.AssetSpecificVote))
	vals = append(vals, fmt.Sprintf("AssetType:%#+v", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#+v", m.AssetID))
	vals = append(vals, fmt.Sprintf("VoteSystem:%v", m.VoteSystem))
	vals = append(vals, fmt.Sprintf("Proposal:%#+v", m.Proposal))
	vals = append(vals, fmt.Sprintf("ProposedChangesCount:%v", m.ProposedChangesCount))
	vals = append(vals, fmt.Sprintf("ProposedChanges:%#+v", m.ProposedChanges))
	vals = append(vals, fmt.Sprintf("VoteOptions:%#+v", m.VoteOptions))
	vals = append(vals, fmt.Sprintf("VoteMax:%v", m.VoteMax))
	vals = append(vals, fmt.Sprintf("ProposalDescription:%#+v", m.ProposalDescription))
	vals = append(vals, fmt.Sprintf("ProposalDocumentHash:%#+v", m.ProposalDocumentHash))
	vals = append(vals, fmt.Sprintf("VoteCutOffTimestamp:%v", m.VoteCutOffTimestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Vote Vote Action - A vote is created by the Contract in response to a
// valid Referendum (Issuer) or Initiative (User) Action.
type Vote struct {
	Header    Header // Common header data for all actions
	Timestamp uint64 // Timestamp in nanoseconds of when the smart contract created the action.
}

// NewVote returns a new empty Vote.
func NewEmptyVote() *Vote {
	result := Vote{}
	return &result
}

// NewVote returns a new Vote with specified values set.
func NewVote() *Vote {
	result := Vote{}
	return &result
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

	header, err := NewHeaderForCode([]byte(CodeVote), len(b))
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
	Header    Header   // Common header data for all actions
	AssetType string   // eg. Share, Bond, Ticket
	AssetID   string   // Randomly generated base58 string.  Each Asset ID should be unique.  However, a Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type can be the leading bytes - a convention - to make it easy to identify that it is a token by humans.
	VoteTxnID [32]byte // Tx-ID of the Vote the Ballot Cast is being made for.
	Vote      string   // Length 1-255 bytes. 0 is not valid. Accept, Reject, Abstain, Spoiled, Multiple Choice, or Preference List. 15 options total. Order of preference.  1st position = 1st choice. 2nd position = 2nd choice, etc.  A is always Accept and B is always reject in a Y/N votes.
}

// NewBallotCast returns a new empty BallotCast.
func NewEmptyBallotCast() *BallotCast {
	result := BallotCast{}
	return &result
}

// NewBallotCast returns a new BallotCast with specified values set.
func NewBallotCast() *BallotCast {
	result := BallotCast{}
	return &result
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

	if err := WriteFixedChar(buf, m.AssetType, 3); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.AssetID, 32); err != nil {
		return nil, err
	}

	if err := write(buf, m.VoteTxnID); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.Vote, 255); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode([]byte(CodeBallotCast), len(b))
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

	{
		var err error
		m.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.AssetID, err = ReadFixedChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	if err := read(buf, &m.VoteTxnID); err != nil {
		return 0, err
	}

	{
		var err error
		m.Vote, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
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
	vals = append(vals, fmt.Sprintf("AssetType:%#+v", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#+v", m.AssetID))
	vals = append(vals, fmt.Sprintf("VoteTxnID:%#+v", m.VoteTxnID))
	vals = append(vals, fmt.Sprintf("Vote:%#+v", m.Vote))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// BallotCounted Ballot Counted Action - The smart contract will respond to
// a Ballot Cast action with a Ballot Counted action if the Ballot Cast is
// valid. If the Ballot Cast is not valid, then the smart contract will
// respond with a Rejection Action.
type BallotCounted struct {
	Header    Header // Common header data for all actions
	Timestamp uint64 // Timestamp in nanoseconds of when the smart contract created the action.
}

// NewBallotCounted returns a new empty BallotCounted.
func NewEmptyBallotCounted() *BallotCounted {
	result := BallotCounted{}
	return &result
}

// NewBallotCounted returns a new BallotCounted with specified values set.
func NewBallotCounted() *BallotCounted {
	result := BallotCounted{}
	return &result
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

	header, err := NewHeaderForCode([]byte(CodeBallotCounted), len(b))
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
	Header               Header      // Common header data for all actions
	AssetType            string      // eg. Share, Bond, Ticket
	AssetID              string      // Randomly generated base58 string.  Each Asset ID should be unique.  However, a Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type can be the leading bytes - a convention - to make it easy to identify that it is a token by humans. If its a Contract vote then can be null.
	Proposal             bool        // 1 for a Proposal, 0 for an initiative that is requesting changes to specific subfields for modification. If this field is true, the subfields should be empty.  The smart contract cannot interpret the results of a vote when Proposal = 1.  All meaning is interpreted by the token owners and smart contract simply facilates the record keeping.  When Proposal = 0, the smart contract always assumes the first choice is a 'yes', or 'pass', if the threshold is met, and will process the proposed changes accordingly.
	ProposedChangesCount uint8       //
	ProposedChanges      []Amendment // Each element contains details of which fields to modify, or delete. Because the number of fields in a Contract and Asset is dynamic due to some fields being able to be repeated, the index value of the field needs to be calculated against the Contract or Asset the changes are to apply to. In the event of a Vote being created from this Initiative, the changes will be applied to the version of the Contract or Asset at that time.
	VoteTxnID            [32]byte    // Link to the Vote Action txn.
	VoteOptionsCount     uint8       // Number of Vote Options to follow.
	OptionXTally         uint64      // Number of valid votes counted for Option X
	Result               string      // Length 1-255 bytes. 0 is not valid. The Option with the most votes. In the event of a draw for 1st place, all winning options are listed.
	Timestamp            uint64      // Timestamp in nanoseconds of when the smart contract created the action.
}

// NewResult returns a new empty Result.
func NewEmptyResult() *Result {
	result := Result{}
	return &result
}

// NewResult returns a new Result with specified values set.
func NewResult() *Result {
	result := Result{}
	return &result
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

	if err := WriteFixedChar(buf, m.AssetType, 3); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.AssetID, 32); err != nil {
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

	if err := write(buf, m.VoteTxnID); err != nil {
		return nil, err
	}

	if err := write(buf, m.VoteOptionsCount); err != nil {
		return nil, err
	}

	if err := write(buf, m.OptionXTally); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.Result, 255); err != nil {
		return nil, err
	}

	if err := write(buf, m.Timestamp); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode([]byte(CodeResult), len(b))
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

	{
		var err error
		m.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.AssetID, err = ReadFixedChar(buf, 32)
		if err != nil {
			return 0, err
		}
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

	if err := read(buf, &m.VoteTxnID); err != nil {
		return 0, err
	}

	if err := read(buf, &m.VoteOptionsCount); err != nil {
		return 0, err
	}

	if err := read(buf, &m.OptionXTally); err != nil {
		return 0, err
	}

	{
		var err error
		m.Result, err = ReadVarChar(buf, 255)
		if err != nil {
			return 0, err
		}
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
	vals = append(vals, fmt.Sprintf("AssetType:%#+v", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#+v", m.AssetID))
	vals = append(vals, fmt.Sprintf("Proposal:%#+v", m.Proposal))
	vals = append(vals, fmt.Sprintf("ProposedChangesCount:%v", m.ProposedChangesCount))
	vals = append(vals, fmt.Sprintf("ProposedChanges:%#+v", m.ProposedChanges))
	vals = append(vals, fmt.Sprintf("VoteTxnID:%#+v", m.VoteTxnID))
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
	Header                Header   // Common header data for all actions
	QtyReceivingAddresses uint8    // 0-255 Message Receiving Addresses
	AddressIndexes        []uint16 // Associates the message to a particular output by the index.
	MessageType           string   // Potential for up to 65,535 different message types
	MessagePayload        string   // Length only limited by the Bitcoin protocol. Public or private (RSA public key, Diffie-Hellman). Issuers/Contracts can send the signifying amount of satoshis to themselves for public announcements or private 'notes' if encrypted. See Message Types for a full list of potential use cases.

}

// NewMessage returns a new empty Message.
func NewEmptyMessage() *Message {
	result := Message{}
	return &result
}

// NewMessage returns a new Message with specified values set.
func NewMessage() *Message {
	result := Message{}
	return &result
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

	if err := WriteFixedChar(buf, m.MessageType, 2); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.MessagePayload, 4294967295); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode([]byte(CodeMessage), len(b))
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

	{
		var err error
		m.MessageType, err = ReadFixedChar(buf, 2)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.MessagePayload, err = ReadVarChar(buf, 4294967295)
		if err != nil {
			return 0, err
		}
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
	vals = append(vals, fmt.Sprintf("MessageType:%#+v", m.MessageType))
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
	Header                Header   // Common header data for all actions
	QtyReceivingAddresses uint8    // 0-255 Message Receiving Addresses
	AddressIndexes        []uint16 // Associates the message to a particular output by the index.
	RejectionType         uint8    // Classifies the rejection by a type.
	MessagePayload        string   // Length 0-65,535 bytes. Message that explains the reasoning for a rejection, if needed.  Most rejection types will be captured by the Rejection Type Subfield.
	Timestamp             uint64   // Timestamp in nanoseconds of when the smart contract created the action.
}

// NewRejection returns a new empty Rejection.
func NewEmptyRejection() *Rejection {
	result := Rejection{}
	return &result
}

// NewRejection returns a new Rejection with specified values set.
func NewRejection() *Rejection {
	result := Rejection{}
	return &result
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

	if err := WriteVarChar(buf, m.MessagePayload, 4294967295); err != nil {
		return nil, err
	}

	if err := write(buf, m.Timestamp); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode([]byte(CodeRejection), len(b))
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

	{
		var err error
		m.MessagePayload, err = ReadVarChar(buf, 4294967295)
		if err != nil {
			return 0, err
		}
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
	Header  Header // Common header data for all actions
	Message string // Length only limited by Bitcoin protocol.
}

// NewEstablishment returns a new empty Establishment.
func NewEmptyEstablishment() *Establishment {
	result := Establishment{}
	return &result
}

// NewEstablishment returns a new Establishment with specified values set.
func NewEstablishment() *Establishment {
	result := Establishment{}
	return &result
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

	if err := WriteVarChar(buf, m.Message, 25); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode([]byte(CodeEstablishment), len(b))
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

	{
		var err error
		m.Message, err = ReadVarChar(buf, 25)
		if err != nil {
			return 0, err
		}
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
	Header  Header // Common header data for all actions
	Message string // Length only limited by Bitcoin protocol.
}

// NewAddition returns a new empty Addition.
func NewEmptyAddition() *Addition {
	result := Addition{}
	return &result
}

// NewAddition returns a new Addition with specified values set.
func NewAddition() *Addition {
	result := Addition{}
	return &result
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

	if err := WriteVarChar(buf, m.Message, 4294967295); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode([]byte(CodeAddition), len(b))
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

	{
		var err error
		m.Message, err = ReadVarChar(buf, 4294967295)
		if err != nil {
			return 0, err
		}
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
	Header  Header // Common header data for all actions
	Message string // Length only limited by the Bitcoin protocol.
}

// NewAlteration returns a new empty Alteration.
func NewEmptyAlteration() *Alteration {
	result := Alteration{}
	return &result
}

// NewAlteration returns a new Alteration with specified values set.
func NewAlteration() *Alteration {
	result := Alteration{}
	return &result
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

	if err := WriteVarChar(buf, m.Message, 4294967295); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode([]byte(CodeAlteration), len(b))
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

	{
		var err error
		m.Message, err = ReadVarChar(buf, 4294967295)
		if err != nil {
			return 0, err
		}
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
	Header  Header // Common header data for all actions
	Message string // Length only limited by the Bitcoin protocol.
}

// NewRemoval returns a new empty Removal.
func NewEmptyRemoval() *Removal {
	result := Removal{}
	return &result
}

// NewRemoval returns a new Removal with specified values set.
func NewRemoval() *Removal {
	result := Removal{}
	return &result
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

	if err := WriteVarChar(buf, m.Message, 4294967295); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode([]byte(CodeRemoval), len(b))
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

	{
		var err error
		m.Message, err = ReadVarChar(buf, 4294967295)
		if err != nil {
			return 0, err
		}
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
	Header              Header          // Common header data for all actions
	AssetCount          uint8           // The number of Assets involved in the Transfer Action.
	AssetTypeX          string          // eg. Share, Bond, Ticket. All characters must be capitalised.
	AssetIDX            string          // Randomly generated base58 string.  Each Asset ID should be unique.  However, a Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type can be the leading bytes - a convention - to make it easy to identify that it is a token by humans.
	AssetXSenderCount   uint8           // Number inputs sending tokens. 1-255, 0 is not valid.
	AssetXSenders       []QuantityIndex // Each element has the value of tokens to be spent from the input address, which is referred to by the index.
	AssetXReceiverCount uint8           // Number of outputs receiving tokens. 1-255. 0 is not valid.
	AssetXReceivers     []TokenReceiver // Each element has the value of tokens to be received by the output address, which is referred to by the index.
	OfferExpiry         uint64          // This prevents any party from holding on to the partially signed message as a form of an option.  Eg. the exchange at this price is valid for 30 mins.
	ExchangeFeeCurrency string          // BSV, USD, AUD, EUR, etc.
	ExchangeFeeVar      float32         // Percent of the value of the transaction
	ExchangeFeeFixed    float32         // Fixed fee
	ExchangeFeeAddress  Address         // Identifies the public address that the exchange fee should be paid to.
}

// NewTransfer returns a new empty Transfer.
func NewEmptyTransfer() *Transfer {
	result := Transfer{}
	return &result
}

// NewTransfer returns a new Transfer with specified values set.
func NewTransfer() *Transfer {
	result := Transfer{}
	return &result
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

	if err := write(buf, m.AssetCount); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.AssetTypeX, 3); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.AssetIDX, 32); err != nil {
		return nil, err
	}

	if err := write(buf, m.AssetXSenderCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.AssetXSenderCount); i++ {
		b, err := m.AssetXSenders[i].Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	if err := write(buf, m.AssetXReceiverCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.AssetXReceiverCount); i++ {
		b, err := m.AssetXReceivers[i].Serialize()
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

	if err := WriteFixedChar(buf, m.ExchangeFeeCurrency, 3); err != nil {
		return nil, err
	}

	if err := write(buf, m.ExchangeFeeVar); err != nil {
		return nil, err
	}

	if err := write(buf, m.ExchangeFeeFixed); err != nil {
		return nil, err
	}

	{
		b, err := m.ExchangeFeeAddress.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode([]byte(CodeTransfer), len(b))
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

	if err := read(buf, &m.AssetCount); err != nil {
		return 0, err
	}

	{
		var err error
		m.AssetTypeX, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.AssetIDX, err = ReadFixedChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	if err := read(buf, &m.AssetXSenderCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.AssetXSenderCount); i++ {
		x := &QuantityIndex{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.AssetXSenders = append(m.AssetXSenders, *x)
	}

	if err := read(buf, &m.AssetXReceiverCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.AssetXReceiverCount); i++ {
		x := &TokenReceiver{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.AssetXReceivers = append(m.AssetXReceivers, *x)
	}

	if err := read(buf, &m.OfferExpiry); err != nil {
		return 0, err
	}

	{
		var err error
		m.ExchangeFeeCurrency, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	if err := read(buf, &m.ExchangeFeeVar); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ExchangeFeeFixed); err != nil {
		return 0, err
	}

	if err := m.ExchangeFeeAddress.Write(buf); err != nil {
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
	vals = append(vals, fmt.Sprintf("AssetCount:%v", m.AssetCount))
	vals = append(vals, fmt.Sprintf("AssetTypeX:%#+v", m.AssetTypeX))
	vals = append(vals, fmt.Sprintf("AssetIDX:%#+v", m.AssetIDX))
	vals = append(vals, fmt.Sprintf("AssetXSenderCount:%v", m.AssetXSenderCount))
	vals = append(vals, fmt.Sprintf("AssetXSenders:%#+v", m.AssetXSenders))
	vals = append(vals, fmt.Sprintf("AssetXReceiverCount:%v", m.AssetXReceiverCount))
	vals = append(vals, fmt.Sprintf("AssetXReceivers:%#+v", m.AssetXReceivers))
	vals = append(vals, fmt.Sprintf("OfferExpiry:%v", m.OfferExpiry))
	vals = append(vals, fmt.Sprintf("ExchangeFeeCurrency:%#+v", m.ExchangeFeeCurrency))
	vals = append(vals, fmt.Sprintf("ExchangeFeeVar:%v", m.ExchangeFeeVar))
	vals = append(vals, fmt.Sprintf("ExchangeFeeFixed:%v", m.ExchangeFeeFixed))
	vals = append(vals, fmt.Sprintf("ExchangeFeeAddress:%#+v", m.ExchangeFeeAddress))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Settlement Settlement Action - Settles the transfer request of bitcoins
// and tokens from transfer (T1) actions.
type Settlement struct {
	Header                 Header          // Common header data for all actions
	AssetCount             uint8           // The number of Assets specified by the Transfer action to be settled.
	AssetTypeX             string          // eg. Share, Bond, Ticket
	AssetIDX               string          // Randomly generated base58 string.  Each Asset ID should be unique.  However, a Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type can be the leading bytes - a convention - to make it easy to identify that it is a token by humans.
	AssetXSettlementsCount uint8           // Number of settlements for Asset X.
	AssetXAddressesXQty    []QuantityIndex // Each element contains the resulting token balance of Asset X for the output Address, which is referred to by the index.
	Timestamp              uint64          // Timestamp in nanoseconds of when the smart contract created the action.
}

// NewSettlement returns a new empty Settlement.
func NewEmptySettlement() *Settlement {
	result := Settlement{}
	return &result
}

// NewSettlement returns a new Settlement with specified values set.
func NewSettlement() *Settlement {
	result := Settlement{}
	return &result
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

	if err := write(buf, m.AssetCount); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.AssetTypeX, 3); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.AssetIDX, 32); err != nil {
		return nil, err
	}

	if err := write(buf, m.AssetXSettlementsCount); err != nil {
		return nil, err
	}

	for i := 0; i < int(m.AssetXSettlementsCount); i++ {
		b, err := m.AssetXAddressesXQty[i].Serialize()
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

	header, err := NewHeaderForCode([]byte(CodeSettlement), len(b))
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

	if err := read(buf, &m.AssetCount); err != nil {
		return 0, err
	}

	{
		var err error
		m.AssetTypeX, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	{
		var err error
		m.AssetIDX, err = ReadFixedChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	if err := read(buf, &m.AssetXSettlementsCount); err != nil {
		return 0, err
	}

	for i := 0; i < int(m.AssetXSettlementsCount); i++ {
		x := &QuantityIndex{}
		if err := x.Write(buf); err != nil {
			return 0, err
		}

		m.AssetXAddressesXQty = append(m.AssetXAddressesXQty, *x)
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
	vals = append(vals, fmt.Sprintf("AssetCount:%v", m.AssetCount))
	vals = append(vals, fmt.Sprintf("AssetTypeX:%#+v", m.AssetTypeX))
	vals = append(vals, fmt.Sprintf("AssetIDX:%#+v", m.AssetIDX))
	vals = append(vals, fmt.Sprintf("AssetXSettlementsCount:%v", m.AssetXSettlementsCount))
	vals = append(vals, fmt.Sprintf("AssetXAddressesXQty:%#+v", m.AssetXAddressesXQty))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
