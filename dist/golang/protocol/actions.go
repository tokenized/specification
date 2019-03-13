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
	AssetPayload                []byte  // Payload length is dependent on the asset type. Each asset is made up of a defined set of information pertaining to the specific asset type, and may contain fields of variable length type (nvarchar8, 16, 32)
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

	// AssetType
	if err := WriteFixedChar(buf, m.AssetType, 3); err != nil {
		return nil, err
	}

	// AssetID
	if err := WriteFixedChar(buf, m.AssetID, 32); err != nil {
		return nil, err
	}

	// AssetAuthFlags
	if err := write(buf, pad(m.AssetAuthFlags, 8)); err != nil {
		return nil, err
	}

	// TransfersPermitted
	if err := write(buf, m.TransfersPermitted); err != nil {
		return nil, err
	}

	// TradeRestrictions
	if err := WriteFixedChar(buf, m.TradeRestrictions, 3); err != nil {
		return nil, err
	}

	// EnforcementOrdersPermitted
	if err := write(buf, m.EnforcementOrdersPermitted); err != nil {
		return nil, err
	}

	// VoteMultiplier
	if err := write(buf, m.VoteMultiplier); err != nil {
		return nil, err
	}

	// ReferendumProposal
	if err := write(buf, m.ReferendumProposal); err != nil {
		return nil, err
	}

	// InitiativeProposal
	if err := write(buf, m.InitiativeProposal); err != nil {
		return nil, err
	}

	// AssetModificationGovernance
	if err := write(buf, m.AssetModificationGovernance); err != nil {
		return nil, err
	}

	// TokenQty
	if err := write(buf, m.TokenQty); err != nil {
		return nil, err
	}

	// ContractFeeCurrency
	if err := WriteFixedChar(buf, m.ContractFeeCurrency, 3); err != nil {
		return nil, err
	}

	// ContractFeeVar
	if err := write(buf, m.ContractFeeVar); err != nil {
		return nil, err
	}

	// ContractFeeFixed
	if err := write(buf, m.ContractFeeFixed); err != nil {
		return nil, err
	}

	// AssetPayload
	if err := WriteVarBin(buf, m.AssetPayload, 16); err != nil {
		return nil, err
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// AssetType
	{
		var err error
		m.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// AssetID
	{
		var err error
		m.AssetID, err = ReadFixedChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// AssetAuthFlags
	m.AssetAuthFlags = make([]byte, 8)
	if err := readLen(buf, m.AssetAuthFlags); err != nil {
		return 0, err
	}

	// TransfersPermitted
	if err := read(buf, &m.TransfersPermitted); err != nil {
		return 0, err
	}

	// TradeRestrictions
	{
		var err error
		m.TradeRestrictions, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// EnforcementOrdersPermitted
	if err := read(buf, &m.EnforcementOrdersPermitted); err != nil {
		return 0, err
	}

	// VoteMultiplier
	if err := read(buf, &m.VoteMultiplier); err != nil {
		return 0, err
	}

	// ReferendumProposal
	if err := read(buf, &m.ReferendumProposal); err != nil {
		return 0, err
	}

	// InitiativeProposal
	if err := read(buf, &m.InitiativeProposal); err != nil {
		return 0, err
	}

	// AssetModificationGovernance
	if err := read(buf, &m.AssetModificationGovernance); err != nil {
		return 0, err
	}

	// TokenQty
	if err := read(buf, &m.TokenQty); err != nil {
		return 0, err
	}

	// ContractFeeCurrency
	{
		var err error
		m.ContractFeeCurrency, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// ContractFeeVar
	if err := read(buf, &m.ContractFeeVar); err != nil {
		return 0, err
	}

	// ContractFeeFixed
	if err := read(buf, &m.ContractFeeFixed); err != nil {
		return 0, err
	}

	// AssetPayload
	{
		var err error
		m.AssetPayload, err = ReadVarBin(buf, 16)
		if err != nil {
			return 0, err
		}
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
	AssetPayload                []byte  // Payload length is dependent on the asset type. Each asset is made up of a defined set of information pertaining to the specific asset type, and may contain fields of variable length type (nvarchar8, 16, 32)
	AssetRevision               uint64  // Counter 0 - 65,535
	Timestamp                   uint64  // Timestamp in nanoseconds of when the smart contract created the action.
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

	// AssetType
	if err := WriteFixedChar(buf, m.AssetType, 3); err != nil {
		return nil, err
	}

	// AssetID
	if err := WriteFixedChar(buf, m.AssetID, 32); err != nil {
		return nil, err
	}

	// AssetAuthFlags
	if err := write(buf, pad(m.AssetAuthFlags, 8)); err != nil {
		return nil, err
	}

	// TransfersPermitted
	if err := write(buf, m.TransfersPermitted); err != nil {
		return nil, err
	}

	// TradeRestrictions
	if err := WriteFixedChar(buf, m.TradeRestrictions, 3); err != nil {
		return nil, err
	}

	// EnforcementOrdersPermitted
	if err := write(buf, m.EnforcementOrdersPermitted); err != nil {
		return nil, err
	}

	// VoteMultiplier
	if err := write(buf, m.VoteMultiplier); err != nil {
		return nil, err
	}

	// ReferendumProposal
	if err := write(buf, m.ReferendumProposal); err != nil {
		return nil, err
	}

	// InitiativeProposal
	if err := write(buf, m.InitiativeProposal); err != nil {
		return nil, err
	}

	// AssetModificationGovernance
	if err := write(buf, m.AssetModificationGovernance); err != nil {
		return nil, err
	}

	// TokenQty
	if err := write(buf, m.TokenQty); err != nil {
		return nil, err
	}

	// ContractFeeCurrency
	if err := WriteFixedChar(buf, m.ContractFeeCurrency, 3); err != nil {
		return nil, err
	}

	// ContractFeeVar
	if err := write(buf, m.ContractFeeVar); err != nil {
		return nil, err
	}

	// ContractFeeFixed
	if err := write(buf, m.ContractFeeFixed); err != nil {
		return nil, err
	}

	// AssetPayload
	if err := WriteVarBin(buf, m.AssetPayload, 16); err != nil {
		return nil, err
	}

	// AssetRevision
	if err := write(buf, m.AssetRevision); err != nil {
		return nil, err
	}

	// Timestamp
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// AssetType
	{
		var err error
		m.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// AssetID
	{
		var err error
		m.AssetID, err = ReadFixedChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// AssetAuthFlags
	m.AssetAuthFlags = make([]byte, 8)
	if err := readLen(buf, m.AssetAuthFlags); err != nil {
		return 0, err
	}

	// TransfersPermitted
	if err := read(buf, &m.TransfersPermitted); err != nil {
		return 0, err
	}

	// TradeRestrictions
	{
		var err error
		m.TradeRestrictions, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// EnforcementOrdersPermitted
	if err := read(buf, &m.EnforcementOrdersPermitted); err != nil {
		return 0, err
	}

	// VoteMultiplier
	if err := read(buf, &m.VoteMultiplier); err != nil {
		return 0, err
	}

	// ReferendumProposal
	if err := read(buf, &m.ReferendumProposal); err != nil {
		return 0, err
	}

	// InitiativeProposal
	if err := read(buf, &m.InitiativeProposal); err != nil {
		return 0, err
	}

	// AssetModificationGovernance
	if err := read(buf, &m.AssetModificationGovernance); err != nil {
		return 0, err
	}

	// TokenQty
	if err := read(buf, &m.TokenQty); err != nil {
		return 0, err
	}

	// ContractFeeCurrency
	{
		var err error
		m.ContractFeeCurrency, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// ContractFeeVar
	if err := read(buf, &m.ContractFeeVar); err != nil {
		return 0, err
	}

	// ContractFeeFixed
	if err := read(buf, &m.ContractFeeFixed); err != nil {
		return 0, err
	}

	// AssetPayload
	{
		var err error
		m.AssetPayload, err = ReadVarBin(buf, 16)
		if err != nil {
			return 0, err
		}
	}

	// AssetRevision
	if err := read(buf, &m.AssetRevision); err != nil {
		return 0, err
	}

	// Timestamp
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
	vals = append(vals, fmt.Sprintf("AssetPayload:%#x", m.AssetPayload))
	vals = append(vals, fmt.Sprintf("AssetRevision:%v", m.AssetRevision))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// AssetModification Asset Modification Action - Token Dilutions, Call
// Backs/Revocations, burning etc.
type AssetModification struct {
	Header        Header      // Common header data for all actions
	AssetType     string      // eg. Share - Common
	AssetID       string      // Randomly generated base58 string.  Each Asset ID should be unique.  However, a Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type + Asset ID = Asset Code.  An Asset Code is a human readable idenitfier that can be used in a similar way to a Bitcoin (BSV) address, a vanity identifying label.
	AssetRevision uint64      // Counter. (Subfield cannot be manually changed by Asset Modification Action.  Only SC can increment by 1 with each AC action. SC will reject AM actions where the wrong asset revision has been selected.
	Modifications []Amendment //
	RefTxID       [32]byte    // Tx-ID of the associated Result action (governance) that permitted the modifications.
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

	// AssetType
	if err := WriteFixedChar(buf, m.AssetType, 3); err != nil {
		return nil, err
	}

	// AssetID
	if err := WriteFixedChar(buf, m.AssetID, 32); err != nil {
		return nil, err
	}

	// AssetRevision
	if err := write(buf, m.AssetRevision); err != nil {
		return nil, err
	}

	// Modifications
	if err := WriteVariableSize(buf, uint64(len(m.Modifications)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range m.Modifications {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	// RefTxID
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// AssetType
	{
		var err error
		m.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// AssetID
	{
		var err error
		m.AssetID, err = ReadFixedChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// AssetRevision
	if err := read(buf, &m.AssetRevision); err != nil {
		return 0, err
	}

	// Modifications
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		m.Modifications = make([]Amendment, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue Amendment
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			m.Modifications = append(m.Modifications, newValue)
		}
	}

	// RefTxID
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
	ContractFile               []byte         // SHA-256 hash of the contract file or markdown data for contract file specific to the smart contract and relevant Assets.  Legal and technical information. (eg. pdf)
	GoverningLaw               string         // 5 Letter Code to Identify which governing law the contract will adhere to.  Disputes are to be settled by this law in the jurisdiction specified below. Private dispute resolution organizations can be used as well.  A custom code just needs to be defined.
	Jurisdiction               string         // Legal proceedings/arbitration will take place using the specified Governing Law in this location.
	ContractExpiration         uint64         // All actions related to the contract will cease to work after this timestamp. The smart contract will stop running.  This will allow many token use cases to be able to calculate total smart contract running costs for the entire life of the contract. Eg. an issuer is creating tickets for an event on the 5th of June 2018.  The smart contract will facilitate exchange and send transactions up until the 6th of June.  Wallets can use this to forget tokens that are no longer valid - or at least store them in an 'Expired' folder.
	ContractURI                string         // Points to an information page that also has a copy of the Contract.  Anyone can go to the website to have a look at the price/token, information about the Issuer (company), information about the Asset, legal information, etc.  There will also be a way for Token Owners to vote on this page and contact details with the Issuer/tokenized companies. Could be a IPv6/IPv4, an IPFS address (hash) or txn-id for on-chain information or even a public address (DNS).
	IssuerName                 string         // Length 0-255 bytes. 0 is not valid.Issuing entity (company, organization, individual).  Can be any unique identifying string, including human readable names for branding/vanity purposes.
	IssuerType                 byte           // P - Public Company Limited by Shares, C - Private Company Limited by Shares, I - Individual, L - Limited Partnership, U -Unlimited Partnership, T - Sole Proprietorship, S - Statutory Company, O - Non-Profit Organization, N - Nation State, G - Government Agency, U - Unit Trust, D - Discretionary Trust.  Found in 'Entities' (Specification/Resources).
	IssuerLogoURL              string         // The URL of the Issuers logo.
	ContractOperatorID         string         // Length 0-255 bytes. 0 is valid. Smart Contract Operator identifier. Can be any unique identifying string, including human readable names for branding/vanity purposes. Can also be null or the Issuer.
	ContractAuthFlags          []byte         // Authorization Flags aka Terms and Conditions that the smart contract can enforce.  Other terms and conditions that are out of the smart contract's control are listed in the actual Contract File.
	VotingSystems              []VotingSystem // A list of voting systems.
	RestrictedQtyAssets        uint64         // Number of Assets (non-fungible) permitted on this contract. 0 if unlimited which will display an infinity symbol in UI
	ReferendumProposal         bool           // A Referendum is permitted for Proposals (outside of smart contract scope).
	InitiativeProposal         bool           // An initiative is permitted for Proposals (outside of smart contract scope).
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
	KeyRoles                   []KeyRole      // A list of Key Roles.
	NotableRoles               []NotableRole  // A list of Notable Roles.
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

	// ContractName
	if err := WriteVarChar(buf, m.ContractName, 8); err != nil {
		return nil, err
	}

	// ContractFileType
	if err := write(buf, m.ContractFileType); err != nil {
		return nil, err
	}

	// ContractFile
	if err := WriteVarBin(buf, m.ContractFile, 32); err != nil {
		return nil, err
	}

	// GoverningLaw
	if err := WriteFixedChar(buf, m.GoverningLaw, 5); err != nil {
		return nil, err
	}

	// Jurisdiction
	if err := WriteFixedChar(buf, m.Jurisdiction, 5); err != nil {
		return nil, err
	}

	// ContractExpiration
	if err := write(buf, m.ContractExpiration); err != nil {
		return nil, err
	}

	// ContractURI
	if err := WriteVarChar(buf, m.ContractURI, 8); err != nil {
		return nil, err
	}

	// IssuerName
	if err := WriteVarChar(buf, m.IssuerName, 8); err != nil {
		return nil, err
	}

	// IssuerType
	if err := write(buf, m.IssuerType); err != nil {
		return nil, err
	}

	// IssuerLogoURL
	if err := WriteVarChar(buf, m.IssuerLogoURL, 8); err != nil {
		return nil, err
	}

	// ContractOperatorID
	if err := WriteVarChar(buf, m.ContractOperatorID, 8); err != nil {
		return nil, err
	}

	// ContractAuthFlags
	if err := write(buf, pad(m.ContractAuthFlags, 16)); err != nil {
		return nil, err
	}

	// VotingSystems
	if err := WriteVariableSize(buf, uint64(len(m.VotingSystems)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range m.VotingSystems {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	// RestrictedQtyAssets
	if err := write(buf, m.RestrictedQtyAssets); err != nil {
		return nil, err
	}

	// ReferendumProposal
	if err := write(buf, m.ReferendumProposal); err != nil {
		return nil, err
	}

	// InitiativeProposal
	if err := write(buf, m.InitiativeProposal); err != nil {
		return nil, err
	}

	// Registries
	if err := WriteVariableSize(buf, uint64(len(m.Registries)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range m.Registries {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	// IssuerAddress
	if err := write(buf, m.IssuerAddress); err != nil {
		return nil, err
	}

	// UnitNumber
	if err := WriteVarChar(buf, m.UnitNumber, 8); err != nil {
		return nil, err
	}

	// BuildingNumber
	if err := WriteVarChar(buf, m.BuildingNumber, 8); err != nil {
		return nil, err
	}

	// Street
	if err := WriteVarChar(buf, m.Street, 16); err != nil {
		return nil, err
	}

	// SuburbCity
	if err := WriteVarChar(buf, m.SuburbCity, 8); err != nil {
		return nil, err
	}

	// TerritoryStateProvinceCode
	if err := WriteFixedChar(buf, m.TerritoryStateProvinceCode, 5); err != nil {
		return nil, err
	}

	// CountryCode
	if err := WriteFixedChar(buf, m.CountryCode, 3); err != nil {
		return nil, err
	}

	// PostalZIPCode
	if err := WriteVarChar(buf, m.PostalZIPCode, 8); err != nil {
		return nil, err
	}

	// EmailAddress
	if err := WriteVarChar(buf, m.EmailAddress, 8); err != nil {
		return nil, err
	}

	// PhoneNumber
	if err := WriteVarChar(buf, m.PhoneNumber, 8); err != nil {
		return nil, err
	}

	// KeyRoles
	if err := WriteVariableSize(buf, uint64(len(m.KeyRoles)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range m.KeyRoles {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	// NotableRoles
	if err := WriteVariableSize(buf, uint64(len(m.NotableRoles)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range m.NotableRoles {
		b, err := value.Serialize()
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// ContractName
	{
		var err error
		m.ContractName, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// ContractFileType
	if err := read(buf, &m.ContractFileType); err != nil {
		return 0, err
	}

	// ContractFile
	{
		var err error
		m.ContractFile, err = ReadVarBin(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// GoverningLaw
	{
		var err error
		m.GoverningLaw, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	// Jurisdiction
	{
		var err error
		m.Jurisdiction, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	// ContractExpiration
	if err := read(buf, &m.ContractExpiration); err != nil {
		return 0, err
	}

	// ContractURI
	{
		var err error
		m.ContractURI, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// IssuerName
	{
		var err error
		m.IssuerName, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// IssuerType
	if err := read(buf, &m.IssuerType); err != nil {
		return 0, err
	}

	// IssuerLogoURL
	{
		var err error
		m.IssuerLogoURL, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// ContractOperatorID
	{
		var err error
		m.ContractOperatorID, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// ContractAuthFlags
	m.ContractAuthFlags = make([]byte, 16)
	if err := readLen(buf, m.ContractAuthFlags); err != nil {
		return 0, err
	}

	// VotingSystems
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		m.VotingSystems = make([]VotingSystem, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue VotingSystem
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			m.VotingSystems = append(m.VotingSystems, newValue)
		}
	}

	// RestrictedQtyAssets
	if err := read(buf, &m.RestrictedQtyAssets); err != nil {
		return 0, err
	}

	// ReferendumProposal
	if err := read(buf, &m.ReferendumProposal); err != nil {
		return 0, err
	}

	// InitiativeProposal
	if err := read(buf, &m.InitiativeProposal); err != nil {
		return 0, err
	}

	// Registries
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		m.Registries = make([]Registry, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue Registry
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			m.Registries = append(m.Registries, newValue)
		}
	}

	// IssuerAddress
	if err := read(buf, &m.IssuerAddress); err != nil {
		return 0, err
	}

	// UnitNumber
	{
		var err error
		m.UnitNumber, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// BuildingNumber
	{
		var err error
		m.BuildingNumber, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// Street
	{
		var err error
		m.Street, err = ReadVarChar(buf, 16)
		if err != nil {
			return 0, err
		}
	}

	// SuburbCity
	{
		var err error
		m.SuburbCity, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// TerritoryStateProvinceCode
	{
		var err error
		m.TerritoryStateProvinceCode, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	// CountryCode
	{
		var err error
		m.CountryCode, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// PostalZIPCode
	{
		var err error
		m.PostalZIPCode, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// EmailAddress
	{
		var err error
		m.EmailAddress, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// PhoneNumber
	{
		var err error
		m.PhoneNumber, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// KeyRoles
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		m.KeyRoles = make([]KeyRole, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue KeyRole
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			m.KeyRoles = append(m.KeyRoles, newValue)
		}
	}

	// NotableRoles
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		m.NotableRoles = make([]NotableRole, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue NotableRole
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			m.NotableRoles = append(m.NotableRoles, newValue)
		}
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
	vals = append(vals, fmt.Sprintf("ContractFile:%#x", m.ContractFile))
	vals = append(vals, fmt.Sprintf("GoverningLaw:%#+v", m.GoverningLaw))
	vals = append(vals, fmt.Sprintf("Jurisdiction:%#+v", m.Jurisdiction))
	vals = append(vals, fmt.Sprintf("ContractExpiration:%v", m.ContractExpiration))
	vals = append(vals, fmt.Sprintf("ContractURI:%#+v", m.ContractURI))
	vals = append(vals, fmt.Sprintf("IssuerName:%#+v", m.IssuerName))
	vals = append(vals, fmt.Sprintf("IssuerType:%#+v", m.IssuerType))
	vals = append(vals, fmt.Sprintf("IssuerLogoURL:%#+v", m.IssuerLogoURL))
	vals = append(vals, fmt.Sprintf("ContractOperatorID:%#+v", m.ContractOperatorID))
	vals = append(vals, fmt.Sprintf("ContractAuthFlags:%#x", m.ContractAuthFlags))
	vals = append(vals, fmt.Sprintf("VotingSystems:%#+v", m.VotingSystems))
	vals = append(vals, fmt.Sprintf("RestrictedQtyAssets:%v", m.RestrictedQtyAssets))
	vals = append(vals, fmt.Sprintf("ReferendumProposal:%#+v", m.ReferendumProposal))
	vals = append(vals, fmt.Sprintf("InitiativeProposal:%#+v", m.InitiativeProposal))
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
	vals = append(vals, fmt.Sprintf("KeyRoles:%#+v", m.KeyRoles))
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
	ContractFile               []byte         // SHA-256 hash of the contract file or markdown data for contract file specific to the smart contract and relevant Assets.  Legal and technical information. (eg. pdf)
	GoverningLaw               string         // 5 Letter Code to Identify which governing law the contract will adhere to.  Disputes are to be settled by this law in the jurisdiction specified below. Private dispute resolution organizations can be used as well.  A custom code just needs to be defined.
	Jurisdiction               string         // Legal proceedings/arbitration will take place using the specified Governing Law in this location.
	ContractExpiration         uint64         // All actions related to the contract will cease to work after this timestamp. The smart contract will stop running.  This will allow many token use cases to be able to calculate smart contract running costs. Eg. an issuer is creating tickets for an event on the 5th of June 2018.  The smart contract will facilitate exchange and send transactions up until the 6th of June.  Wallets can use this to forget tokens that are no longer valid - or at least store them in an 'Expired' folder.
	ContractURI                string         // Length 0-255 bytes.  0 is valid. Points to an information page that also has a copy of the Contract.  Anyone can go to the website to have a look at the price/token, information about the Issuer (company), information about the Asset, legal information, etc.  There will also be a way for Token Owners to vote on this page and contact details with the Issuer/tokenized companies. Could be a IPv6/IPv4, an IPFS address (hash) or txn-id for on chain information or even a public address (DNS).
	IssuerName                 string         // Length 0-255 bytes. 0 is not valid. Issuing entity (company, organization, individual).  Can be any unique identifying string, including human readable names for branding/vanity purposes.
	IssuerType                 byte           // P - Public Company Limited by Shares, C - Private Company Limited by Shares, I - Individual, L - Limited Partnership, U -Unlimited Partnership, T - Sole Proprietorship, S - Statutory Company, O - Non-Profit Organization, N - Nation State, G - Government Agency, U - Unit Trust, D - Discretionary Trust.  Found in 'Entities' (Specification/Resources).
	IssuerLogoURL              string         // The URL of the Issuers logo.
	ContractOperatorID         string         // Length 0-255 bytes. 0 is valid. Smart Contract Operator identifier. Can be any unique identifying string, including human readable names for branding/vanity purposes. Can also be null or the Issuer.
	ContractAuthFlags          []byte         // Authorization Flags aka Terms and Conditions that the smart contract can enforce.  Other terms and conditions that are out of the smart contract's control are listed in the actual Contract File.
	VotingSystems              []VotingSystem // A list voting systems.
	RestrictedQtyAssets        uint64         // Number of Assets (non-fungible) permitted on this contract. 0 if unlimited which will display an infinity symbol in UI
	ReferendumProposal         bool           // A Referendum is permitted for Contract-Wide Proposals (outside of smart contract scope).
	InitiativeProposal         bool           // An initiative is permitted for Contract-Wide Proposals (outside of smart contract scope).
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
	KeyRoles                   []KeyRole      // A list of Key Roles.
	NotableRoles               []NotableRole  // A list of Notable Roles.
	ContractRevision           uint64         // Counter. Cannot be manually changed by issuer.  Can only be incremented by 1 by SC when CF action is published.
	Timestamp                  uint64         // Timestamp in nanoseconds of when the smart contract created the action.
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

	// ContractName
	if err := WriteVarChar(buf, m.ContractName, 8); err != nil {
		return nil, err
	}

	// ContractFileType
	if err := write(buf, m.ContractFileType); err != nil {
		return nil, err
	}

	// ContractFile
	if err := WriteVarBin(buf, m.ContractFile, 32); err != nil {
		return nil, err
	}

	// GoverningLaw
	if err := WriteFixedChar(buf, m.GoverningLaw, 5); err != nil {
		return nil, err
	}

	// Jurisdiction
	if err := WriteFixedChar(buf, m.Jurisdiction, 5); err != nil {
		return nil, err
	}

	// ContractExpiration
	if err := write(buf, m.ContractExpiration); err != nil {
		return nil, err
	}

	// ContractURI
	if err := WriteVarChar(buf, m.ContractURI, 8); err != nil {
		return nil, err
	}

	// IssuerName
	if err := WriteVarChar(buf, m.IssuerName, 8); err != nil {
		return nil, err
	}

	// IssuerType
	if err := write(buf, m.IssuerType); err != nil {
		return nil, err
	}

	// IssuerLogoURL
	if err := WriteVarChar(buf, m.IssuerLogoURL, 8); err != nil {
		return nil, err
	}

	// ContractOperatorID
	if err := WriteVarChar(buf, m.ContractOperatorID, 8); err != nil {
		return nil, err
	}

	// ContractAuthFlags
	if err := write(buf, pad(m.ContractAuthFlags, 16)); err != nil {
		return nil, err
	}

	// VotingSystems
	if err := WriteVariableSize(buf, uint64(len(m.VotingSystems)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range m.VotingSystems {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	// RestrictedQtyAssets
	if err := write(buf, m.RestrictedQtyAssets); err != nil {
		return nil, err
	}

	// ReferendumProposal
	if err := write(buf, m.ReferendumProposal); err != nil {
		return nil, err
	}

	// InitiativeProposal
	if err := write(buf, m.InitiativeProposal); err != nil {
		return nil, err
	}

	// Registries
	if err := WriteVariableSize(buf, uint64(len(m.Registries)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range m.Registries {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	// IssuerAddress
	if err := write(buf, m.IssuerAddress); err != nil {
		return nil, err
	}

	// UnitNumber
	if err := WriteVarChar(buf, m.UnitNumber, 8); err != nil {
		return nil, err
	}

	// BuildingNumber
	if err := WriteVarChar(buf, m.BuildingNumber, 8); err != nil {
		return nil, err
	}

	// Street
	if err := WriteVarChar(buf, m.Street, 16); err != nil {
		return nil, err
	}

	// SuburbCity
	if err := WriteVarChar(buf, m.SuburbCity, 8); err != nil {
		return nil, err
	}

	// TerritoryStateProvinceCode
	if err := WriteFixedChar(buf, m.TerritoryStateProvinceCode, 5); err != nil {
		return nil, err
	}

	// CountryCode
	if err := WriteFixedChar(buf, m.CountryCode, 3); err != nil {
		return nil, err
	}

	// PostalZIPCode
	if err := WriteVarChar(buf, m.PostalZIPCode, 8); err != nil {
		return nil, err
	}

	// EmailAddress
	if err := WriteVarChar(buf, m.EmailAddress, 8); err != nil {
		return nil, err
	}

	// PhoneNumber
	if err := WriteVarChar(buf, m.PhoneNumber, 8); err != nil {
		return nil, err
	}

	// KeyRoles
	if err := WriteVariableSize(buf, uint64(len(m.KeyRoles)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range m.KeyRoles {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	// NotableRoles
	if err := WriteVariableSize(buf, uint64(len(m.NotableRoles)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range m.NotableRoles {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	// ContractRevision
	if err := write(buf, m.ContractRevision); err != nil {
		return nil, err
	}

	// Timestamp
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// ContractName
	{
		var err error
		m.ContractName, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// ContractFileType
	if err := read(buf, &m.ContractFileType); err != nil {
		return 0, err
	}

	// ContractFile
	{
		var err error
		m.ContractFile, err = ReadVarBin(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// GoverningLaw
	{
		var err error
		m.GoverningLaw, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	// Jurisdiction
	{
		var err error
		m.Jurisdiction, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	// ContractExpiration
	if err := read(buf, &m.ContractExpiration); err != nil {
		return 0, err
	}

	// ContractURI
	{
		var err error
		m.ContractURI, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// IssuerName
	{
		var err error
		m.IssuerName, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// IssuerType
	if err := read(buf, &m.IssuerType); err != nil {
		return 0, err
	}

	// IssuerLogoURL
	{
		var err error
		m.IssuerLogoURL, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// ContractOperatorID
	{
		var err error
		m.ContractOperatorID, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// ContractAuthFlags
	m.ContractAuthFlags = make([]byte, 16)
	if err := readLen(buf, m.ContractAuthFlags); err != nil {
		return 0, err
	}

	// VotingSystems
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		m.VotingSystems = make([]VotingSystem, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue VotingSystem
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			m.VotingSystems = append(m.VotingSystems, newValue)
		}
	}

	// RestrictedQtyAssets
	if err := read(buf, &m.RestrictedQtyAssets); err != nil {
		return 0, err
	}

	// ReferendumProposal
	if err := read(buf, &m.ReferendumProposal); err != nil {
		return 0, err
	}

	// InitiativeProposal
	if err := read(buf, &m.InitiativeProposal); err != nil {
		return 0, err
	}

	// Registries
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		m.Registries = make([]Registry, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue Registry
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			m.Registries = append(m.Registries, newValue)
		}
	}

	// IssuerAddress
	if err := read(buf, &m.IssuerAddress); err != nil {
		return 0, err
	}

	// UnitNumber
	{
		var err error
		m.UnitNumber, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// BuildingNumber
	{
		var err error
		m.BuildingNumber, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// Street
	{
		var err error
		m.Street, err = ReadVarChar(buf, 16)
		if err != nil {
			return 0, err
		}
	}

	// SuburbCity
	{
		var err error
		m.SuburbCity, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// TerritoryStateProvinceCode
	{
		var err error
		m.TerritoryStateProvinceCode, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	// CountryCode
	{
		var err error
		m.CountryCode, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// PostalZIPCode
	{
		var err error
		m.PostalZIPCode, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// EmailAddress
	{
		var err error
		m.EmailAddress, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// PhoneNumber
	{
		var err error
		m.PhoneNumber, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// KeyRoles
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		m.KeyRoles = make([]KeyRole, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue KeyRole
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			m.KeyRoles = append(m.KeyRoles, newValue)
		}
	}

	// NotableRoles
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		m.NotableRoles = make([]NotableRole, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue NotableRole
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			m.NotableRoles = append(m.NotableRoles, newValue)
		}
	}

	// ContractRevision
	if err := read(buf, &m.ContractRevision); err != nil {
		return 0, err
	}

	// Timestamp
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
	vals = append(vals, fmt.Sprintf("ContractFile:%#x", m.ContractFile))
	vals = append(vals, fmt.Sprintf("GoverningLaw:%#+v", m.GoverningLaw))
	vals = append(vals, fmt.Sprintf("Jurisdiction:%#+v", m.Jurisdiction))
	vals = append(vals, fmt.Sprintf("ContractExpiration:%v", m.ContractExpiration))
	vals = append(vals, fmt.Sprintf("ContractURI:%#+v", m.ContractURI))
	vals = append(vals, fmt.Sprintf("IssuerName:%#+v", m.IssuerName))
	vals = append(vals, fmt.Sprintf("IssuerType:%#+v", m.IssuerType))
	vals = append(vals, fmt.Sprintf("IssuerLogoURL:%#+v", m.IssuerLogoURL))
	vals = append(vals, fmt.Sprintf("ContractOperatorID:%#+v", m.ContractOperatorID))
	vals = append(vals, fmt.Sprintf("ContractAuthFlags:%#x", m.ContractAuthFlags))
	vals = append(vals, fmt.Sprintf("VotingSystems:%#+v", m.VotingSystems))
	vals = append(vals, fmt.Sprintf("RestrictedQtyAssets:%v", m.RestrictedQtyAssets))
	vals = append(vals, fmt.Sprintf("ReferendumProposal:%#+v", m.ReferendumProposal))
	vals = append(vals, fmt.Sprintf("InitiativeProposal:%#+v", m.InitiativeProposal))
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
	vals = append(vals, fmt.Sprintf("KeyRoles:%#+v", m.KeyRoles))
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
	Amendments            []Amendment //
	RefTxID               [32]byte    // Tx-ID of the associated Result action (governance) that permitted the modifications.
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

	// ChangeIssuerAddress
	if err := write(buf, m.ChangeIssuerAddress); err != nil {
		return nil, err
	}

	// ChangeOperatorAddress
	if err := write(buf, m.ChangeOperatorAddress); err != nil {
		return nil, err
	}

	// ContractRevision
	if err := write(buf, m.ContractRevision); err != nil {
		return nil, err
	}

	// Amendments
	if err := WriteVariableSize(buf, uint64(len(m.Amendments)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range m.Amendments {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	// RefTxID
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// ChangeIssuerAddress
	if err := read(buf, &m.ChangeIssuerAddress); err != nil {
		return 0, err
	}

	// ChangeOperatorAddress
	if err := read(buf, &m.ChangeOperatorAddress); err != nil {
		return 0, err
	}

	// ContractRevision
	if err := read(buf, &m.ContractRevision); err != nil {
		return 0, err
	}

	// Amendments
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		m.Amendments = make([]Amendment, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue Amendment
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			m.Amendments = append(m.Amendments, newValue)
		}
	}

	// RefTxID
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
	ContractFile       []byte   // SHA-256 hash of the contract file or markdown data for contract file specific to the smart contract and relevant Assets.  Legal and technical information. (eg. pdf)
	ContractRevision   uint16   // Counter 0 - 65,535
	GoverningLaw       string   // 5 Letter Code to Identify which governing law the contract will adhere to.  Disputes are to be settled by this law in the jurisdiction specified below. Private dispute resolution organizations can be used as well.  A custom code just needs to be defined.
	Jurisdiction       string   // Legal proceedings/arbitration will take place using the specified Governing Law in this location.
	EffectiveDate      uint64   // Start date of the contract.
	ContractExpiration uint64   // All actions related to the contract will cease to work after this timestamp. The smart contract will stop running.  This will allow many token use cases to be able to calculate smart contract running costs. Eg. an issuer is creating tickets for an event on the 5th of June 2018.  The smart contract will facilitate exchange and send transactions up until the 6th of June.  Wallets can use this to forget tokens that are no longer valid - or at least store them in an 'Expired' folder.
	ContractURI        string   // Length 0-255 bytes. Points to an information page that also has a copy of the Contract.  Anyone can go to the website to have a look at the price/token, information about the Issuer (company), information about the Asset, legal information, etc.  There will also be a way for Token Owners to vote on this page and contact details with the Issuer/tokenized companies. Could be a IPv6/IPv4, an IPFS address (hash) or txn-id for on chain information or even a public address (DNS).
	PrevRevTxID        [32]byte // The Tx-ID of the previous contract revision.
	Entities           []Entity //
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

	// ContractName
	if err := WriteVarChar(buf, m.ContractName, 8); err != nil {
		return nil, err
	}

	// ContractType
	if err := WriteVarChar(buf, m.ContractType, 8); err != nil {
		return nil, err
	}

	// ContractFileType
	if err := write(buf, m.ContractFileType); err != nil {
		return nil, err
	}

	// ContractFile
	if err := WriteVarBin(buf, m.ContractFile, 32); err != nil {
		return nil, err
	}

	// ContractRevision
	if err := write(buf, m.ContractRevision); err != nil {
		return nil, err
	}

	// GoverningLaw
	if err := WriteFixedChar(buf, m.GoverningLaw, 5); err != nil {
		return nil, err
	}

	// Jurisdiction
	if err := WriteFixedChar(buf, m.Jurisdiction, 5); err != nil {
		return nil, err
	}

	// EffectiveDate
	if err := write(buf, m.EffectiveDate); err != nil {
		return nil, err
	}

	// ContractExpiration
	if err := write(buf, m.ContractExpiration); err != nil {
		return nil, err
	}

	// ContractURI
	if err := WriteVarChar(buf, m.ContractURI, 8); err != nil {
		return nil, err
	}

	// PrevRevTxID
	if err := write(buf, m.PrevRevTxID); err != nil {
		return nil, err
	}

	// Entities
	if err := WriteVariableSize(buf, uint64(len(m.Entities)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range m.Entities {
		b, err := value.Serialize()
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// ContractName
	{
		var err error
		m.ContractName, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// ContractType
	{
		var err error
		m.ContractType, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// ContractFileType
	if err := read(buf, &m.ContractFileType); err != nil {
		return 0, err
	}

	// ContractFile
	{
		var err error
		m.ContractFile, err = ReadVarBin(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// ContractRevision
	if err := read(buf, &m.ContractRevision); err != nil {
		return 0, err
	}

	// GoverningLaw
	{
		var err error
		m.GoverningLaw, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	// Jurisdiction
	{
		var err error
		m.Jurisdiction, err = ReadFixedChar(buf, 5)
		if err != nil {
			return 0, err
		}
	}

	// EffectiveDate
	if err := read(buf, &m.EffectiveDate); err != nil {
		return 0, err
	}

	// ContractExpiration
	if err := read(buf, &m.ContractExpiration); err != nil {
		return 0, err
	}

	// ContractURI
	{
		var err error
		m.ContractURI, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// PrevRevTxID
	if err := read(buf, &m.PrevRevTxID); err != nil {
		return 0, err
	}

	// Entities
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		m.Entities = make([]Entity, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue Entity
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			m.Entities = append(m.Entities, newValue)
		}
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
	vals = append(vals, fmt.Sprintf("ContractFile:%#x", m.ContractFile))
	vals = append(vals, fmt.Sprintf("ContractRevision:%v", m.ContractRevision))
	vals = append(vals, fmt.Sprintf("GoverningLaw:%#+v", m.GoverningLaw))
	vals = append(vals, fmt.Sprintf("Jurisdiction:%#+v", m.Jurisdiction))
	vals = append(vals, fmt.Sprintf("EffectiveDate:%v", m.EffectiveDate))
	vals = append(vals, fmt.Sprintf("ContractExpiration:%v", m.ContractExpiration))
	vals = append(vals, fmt.Sprintf("ContractURI:%#+v", m.ContractURI))
	vals = append(vals, fmt.Sprintf("PrevRevTxID:%#+v", m.PrevRevTxID))
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
	TargetAddresses        []TargetAddress //
	DepositAddress         Address         // The public address for confiscated tokens to be deposited in.  Null for Freeze, Thaw, actions. For Reconciliation actions the deposit address is who receives bitcoin.
	AuthorityName          string          // Length 0-255 bytes. Enforcement Authority Name (eg. Issuer, Queensland Police Service, Tokenized, etc.)
	SigAlgoAddressList     uint8           // 0 = No Registry-signed Message, 1 = ECDSA+secp256k1
	AuthorityPublicKey     string          // Length 0-255 bytes. Public Key associated with the Enforcement Authority
	OrderSignature         string          // Length 0-255 bytes. Signature for a message that lists out the target addresses and deposit address. Signature of (Contract Address, Asset Code, Compliance Action, Supporting Evidence Hash, Time Out Expiration, TargetAddress1, TargetAddress1Qty, TargetAddressX, TargetAddressXQty,...,DepositAddress)
	SupportingEvidenceHash [32]byte        // SHA-256: warrant, court order, etc.
	RefTxnID               [32]byte        // The settlement action that was dropped from the network.  Not applicable for Freeze, Thaw, and Confiscation orders.  Only applicable for reconcilliation actions.  No subfield when F, T, R is selected as the Compliance Action subfield.
	FreezePeriod           uint64          // Used for a 'time out'.  Tokens are automatically unfrozen after the expiration timestamp without requiring a Thaw Action. Null value for Thaw, Confiscation and Reconciallitaion orders.
	Message                string          // Length only limited by the Bitcoin protocol.
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

	// AssetType
	if err := WriteFixedChar(buf, m.AssetType, 3); err != nil {
		return nil, err
	}

	// AssetID
	if err := WriteFixedChar(buf, m.AssetID, 32); err != nil {
		return nil, err
	}

	// ComplianceAction
	if err := write(buf, m.ComplianceAction); err != nil {
		return nil, err
	}

	// TargetAddresses
	if err := WriteVariableSize(buf, uint64(len(m.TargetAddresses)), 16, 8); err != nil {
		return nil, err
	}
	for _, value := range m.TargetAddresses {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	// DepositAddress
	{
		b, err := m.DepositAddress.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	// AuthorityName
	if err := WriteVarChar(buf, m.AuthorityName, 8); err != nil {
		return nil, err
	}

	// SigAlgoAddressList
	if err := write(buf, m.SigAlgoAddressList); err != nil {
		return nil, err
	}

	// AuthorityPublicKey
	if err := WriteVarChar(buf, m.AuthorityPublicKey, 8); err != nil {
		return nil, err
	}

	// OrderSignature
	if err := WriteVarChar(buf, m.OrderSignature, 8); err != nil {
		return nil, err
	}

	// SupportingEvidenceHash
	if err := write(buf, m.SupportingEvidenceHash); err != nil {
		return nil, err
	}

	// RefTxnID
	if err := write(buf, m.RefTxnID); err != nil {
		return nil, err
	}

	// FreezePeriod
	if err := write(buf, m.FreezePeriod); err != nil {
		return nil, err
	}

	// Message
	if err := WriteVarChar(buf, m.Message, 32); err != nil {
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// AssetType
	{
		var err error
		m.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// AssetID
	{
		var err error
		m.AssetID, err = ReadFixedChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// ComplianceAction
	if err := read(buf, &m.ComplianceAction); err != nil {
		return 0, err
	}

	// TargetAddresses
	{
		size, err := ReadVariableSize(buf, 16, 8)
		if err != nil {
			return 0, err
		}
		m.TargetAddresses = make([]TargetAddress, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue TargetAddress
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			m.TargetAddresses = append(m.TargetAddresses, newValue)
		}
	}

	// DepositAddress
	if err := m.DepositAddress.Write(buf); err != nil {
		return 0, err
	}

	// AuthorityName
	{
		var err error
		m.AuthorityName, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// SigAlgoAddressList
	if err := read(buf, &m.SigAlgoAddressList); err != nil {
		return 0, err
	}

	// AuthorityPublicKey
	{
		var err error
		m.AuthorityPublicKey, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// OrderSignature
	{
		var err error
		m.OrderSignature, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// SupportingEvidenceHash
	if err := read(buf, &m.SupportingEvidenceHash); err != nil {
		return 0, err
	}

	// RefTxnID
	if err := read(buf, &m.RefTxnID); err != nil {
		return 0, err
	}

	// FreezePeriod
	if err := read(buf, &m.FreezePeriod); err != nil {
		return 0, err
	}

	// Message
	{
		var err error
		m.Message, err = ReadVarChar(buf, 32)
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
	Header    Header    // Common header data for all actions
	Addresses []Address // Addresses holding tokens to be frozen.
	Timestamp uint64    // Timestamp in nanoseconds of when the smart contract created the action.
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

	// Addresses
	if err := WriteVariableSize(buf, uint64(len(m.Addresses)), 16, 8); err != nil {
		return nil, err
	}
	for _, value := range m.Addresses {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	// Timestamp
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// Addresses
	{
		size, err := ReadVariableSize(buf, 16, 8)
		if err != nil {
			return 0, err
		}
		m.Addresses = make([]Address, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue Address
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			m.Addresses = append(m.Addresses, newValue)
		}
	}

	// Timestamp
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
	vals = append(vals, fmt.Sprintf("Addresses:%#+v", m.Addresses))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Thaw Thaw Action - to be used to comply with contractual obligations or
// legal requirements. The Alleged Offender's tokens will be unfrozen to
// allow them to resume normal exchange and governance activities.
type Thaw struct {
	Header    Header    // Common header data for all actions
	Addresses []Address // Addresses holding tokens to be thawed.
	RefTxnID  [32]byte  // The related freeze action.
	Timestamp uint64    // Timestamp in nanoseconds of when the smart contract created the action.
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

	// Addresses
	if err := WriteVariableSize(buf, uint64(len(m.Addresses)), 16, 8); err != nil {
		return nil, err
	}
	for _, value := range m.Addresses {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	// RefTxnID
	if err := write(buf, m.RefTxnID); err != nil {
		return nil, err
	}

	// Timestamp
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// Addresses
	{
		size, err := ReadVariableSize(buf, 16, 8)
		if err != nil {
			return 0, err
		}
		m.Addresses = make([]Address, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue Address
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			m.Addresses = append(m.Addresses, newValue)
		}
	}

	// RefTxnID
	if err := read(buf, &m.RefTxnID); err != nil {
		return 0, err
	}

	// Timestamp
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
	vals = append(vals, fmt.Sprintf("Addresses:%#+v", m.Addresses))
	vals = append(vals, fmt.Sprintf("RefTxnID:%#+v", m.RefTxnID))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Confiscation Confiscation Action - to be used to comply with contractual
// obligations, legal and/or issuer requirements.
type Confiscation struct {
	Header     Header    // Common header data for all actions
	Addresses  []Address // Addresses holding tokens to be confiscated.
	DepositQty uint64    // Custodian's token balance after confiscation.
	Timestamp  uint64    // Timestamp in nanoseconds of when the smart contract created the action.
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

	// Addresses
	if err := WriteVariableSize(buf, uint64(len(m.Addresses)), 16, 8); err != nil {
		return nil, err
	}
	for _, value := range m.Addresses {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	// DepositQty
	if err := write(buf, m.DepositQty); err != nil {
		return nil, err
	}

	// Timestamp
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// Addresses
	{
		size, err := ReadVariableSize(buf, 16, 8)
		if err != nil {
			return 0, err
		}
		m.Addresses = make([]Address, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue Address
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			m.Addresses = append(m.Addresses, newValue)
		}
	}

	// DepositQty
	if err := read(buf, &m.DepositQty); err != nil {
		return 0, err
	}

	// Timestamp
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
	vals = append(vals, fmt.Sprintf("Addresses:%#+v", m.Addresses))
	vals = append(vals, fmt.Sprintf("DepositQty:%v", m.DepositQty))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Reconciliation Reconciliation Action - to be used at the direction of
// the issuer to fix record keeping errors with bitcoin and token balances.
type Reconciliation struct {
	Header    Header    // Common header data for all actions
	Addresses []Address // Addresses holding tokens to be reconciled.
	Timestamp uint64    // Timestamp in nanoseconds of when the smart contract created the action.
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

	// Addresses
	if err := WriteVariableSize(buf, uint64(len(m.Addresses)), 16, 8); err != nil {
		return nil, err
	}
	for _, value := range m.Addresses {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	// Timestamp
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// Addresses
	{
		size, err := ReadVariableSize(buf, 16, 8)
		if err != nil {
			return 0, err
		}
		m.Addresses = make([]Address, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue Address
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			m.Addresses = append(m.Addresses, newValue)
		}
	}

	// Timestamp
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
	ProposedChanges      []Amendment // Each element contains details of which fields to modify, or delete. Because the number of fields in a Contract and Asset is dynamic due to some fields being able to be repeated, the index value of the field needs to be calculated against the Contract or Asset the changes are to apply to. In the event of a Vote being created from this Initiative, the changes will be applied to the version of the Contract or Asset at that time.
	VoteOptions          string      // Length 1-255 bytes. 0 is not valid. Each byte allows for a different vote option.  Typical votes will likely be multiple choice or Y/N. Vote instances are identified by the Tx-ID. AB000000000 would be chosen for Y/N (binary) type votes.
	VoteMax              uint8       // Range: 1-X. How many selections can a voter make in a Ballot Cast.  1 is selected for Y/N (binary)
	ProposalDescription  string      // Length restricted by the Bitcoin protocol. 0 is valid. Description or details of the vote
	ProposalDocumentHash [32]byte    // Hash of the proposal document to be distributed to voters.
	VoteCutOffTimestamp  uint64      // Ballot casts after this timestamp will not be included. The vote has finished.
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

	// AssetType
	if err := WriteFixedChar(buf, m.AssetType, 3); err != nil {
		return nil, err
	}

	// AssetID
	if err := WriteFixedChar(buf, m.AssetID, 32); err != nil {
		return nil, err
	}

	// VoteSystem
	if err := write(buf, m.VoteSystem); err != nil {
		return nil, err
	}

	// Proposal
	if err := write(buf, m.Proposal); err != nil {
		return nil, err
	}

	// ProposedChanges
	if err := WriteVariableSize(buf, uint64(len(m.ProposedChanges)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range m.ProposedChanges {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	// VoteOptions
	if err := WriteVarChar(buf, m.VoteOptions, 8); err != nil {
		return nil, err
	}

	// VoteMax
	if err := write(buf, m.VoteMax); err != nil {
		return nil, err
	}

	// ProposalDescription
	if err := WriteVarChar(buf, m.ProposalDescription, 32); err != nil {
		return nil, err
	}

	// ProposalDocumentHash
	if err := write(buf, m.ProposalDocumentHash); err != nil {
		return nil, err
	}

	// VoteCutOffTimestamp
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// AssetType
	{
		var err error
		m.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// AssetID
	{
		var err error
		m.AssetID, err = ReadFixedChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// VoteSystem
	if err := read(buf, &m.VoteSystem); err != nil {
		return 0, err
	}

	// Proposal
	if err := read(buf, &m.Proposal); err != nil {
		return 0, err
	}

	// ProposedChanges
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		m.ProposedChanges = make([]Amendment, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue Amendment
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			m.ProposedChanges = append(m.ProposedChanges, newValue)
		}
	}

	// VoteOptions
	{
		var err error
		m.VoteOptions, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// VoteMax
	if err := read(buf, &m.VoteMax); err != nil {
		return 0, err
	}

	// ProposalDescription
	{
		var err error
		m.ProposalDescription, err = ReadVarChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// ProposalDocumentHash
	if err := read(buf, &m.ProposalDocumentHash); err != nil {
		return 0, err
	}

	// VoteCutOffTimestamp
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
	AssetSpecificVote    bool        // 1 - Yes, 0 - No.  No Asset Type/AssetID subfields for N - No.
	AssetType            string      // eg. Share, Bond, Ticket
	AssetID              string      // Randomly generated base58 string.  Each Asset ID should be unique.  However, an Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type can be the leading bytes - a convention - to make it easy to identify that it is a token by humans.
	VoteSystem           uint8       // X for Vote System X. (1-255, 0 is not valid.)
	Proposal             bool        // 1 for a Proposal, 0 for an initiative that is requesting changes to specific subfields for modification. If this field is true, the subfields should be empty.  The smart contract cannot interpret the results of a vote when Proposal = 1.  All meaning is interpreted by the token owners and smart contract simply facilates the record keeping.  When Proposal = 0, the smart contract always assumes the first choice is a 'yes', or 'pass', if the threshold is met, and will process the proposed changes accordingly.
	ProposedChanges      []Amendment // Each element contains details of which fields to modify, or delete. Because the number of fields in a Contract and Asset is dynamic due to some fields being able to be repeated, the index value of the field needs to be calculated against the Contract or Asset the changes are to apply to. In the event of a Vote being created from this Initiative, the changes will be applied to the version of the Contract or Asset at that time.
	VoteOptions          string      // Length 1-255 bytes. 0 is not valid. Each byte allows for a different vote option.  Typical votes will likely be multiple choice or Y/N. Vote instances are identified by the Tx-ID. AB000000000 would be chosen for Y/N (binary) type votes. Only applicable if Proposal Type is set to P for Proposal.  All other Proposal Types will be binary.  Pass/Fail.
	VoteMax              uint8       // Range: 1-15. How many selections can a voter make in a Ballot Cast.  1 is selected for Y/N (binary)
	ProposalDescription  string      // Length restricted by the Bitcoin protocol. 0 is valid. Description of the vote.
	ProposalDocumentHash [32]byte    // Hash of the proposal document to be distributed to voters
	VoteCutOffTimestamp  uint64      // Ballot casts after this timestamp will not be included. The vote has finished.
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

	// AssetSpecificVote
	if err := write(buf, m.AssetSpecificVote); err != nil {
		return nil, err
	}

	// AssetType
	if err := WriteFixedChar(buf, m.AssetType, 3); err != nil {
		return nil, err
	}

	// AssetID
	if err := WriteFixedChar(buf, m.AssetID, 32); err != nil {
		return nil, err
	}

	// VoteSystem
	if err := write(buf, m.VoteSystem); err != nil {
		return nil, err
	}

	// Proposal
	if err := write(buf, m.Proposal); err != nil {
		return nil, err
	}

	// ProposedChanges
	if err := WriteVariableSize(buf, uint64(len(m.ProposedChanges)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range m.ProposedChanges {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	// VoteOptions
	if err := WriteVarChar(buf, m.VoteOptions, 8); err != nil {
		return nil, err
	}

	// VoteMax
	if err := write(buf, m.VoteMax); err != nil {
		return nil, err
	}

	// ProposalDescription
	if err := WriteVarChar(buf, m.ProposalDescription, 32); err != nil {
		return nil, err
	}

	// ProposalDocumentHash
	if err := write(buf, m.ProposalDocumentHash); err != nil {
		return nil, err
	}

	// VoteCutOffTimestamp
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// AssetSpecificVote
	if err := read(buf, &m.AssetSpecificVote); err != nil {
		return 0, err
	}

	// AssetType
	{
		var err error
		m.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// AssetID
	{
		var err error
		m.AssetID, err = ReadFixedChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// VoteSystem
	if err := read(buf, &m.VoteSystem); err != nil {
		return 0, err
	}

	// Proposal
	if err := read(buf, &m.Proposal); err != nil {
		return 0, err
	}

	// ProposedChanges
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		m.ProposedChanges = make([]Amendment, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue Amendment
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			m.ProposedChanges = append(m.ProposedChanges, newValue)
		}
	}

	// VoteOptions
	{
		var err error
		m.VoteOptions, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// VoteMax
	if err := read(buf, &m.VoteMax); err != nil {
		return 0, err
	}

	// ProposalDescription
	{
		var err error
		m.ProposalDescription, err = ReadVarChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// ProposalDocumentHash
	if err := read(buf, &m.ProposalDocumentHash); err != nil {
		return 0, err
	}

	// VoteCutOffTimestamp
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
	vals = append(vals, fmt.Sprintf("AssetType:%#+v", m.AssetType))
	vals = append(vals, fmt.Sprintf("AssetID:%#+v", m.AssetID))
	vals = append(vals, fmt.Sprintf("VoteSystem:%v", m.VoteSystem))
	vals = append(vals, fmt.Sprintf("Proposal:%#+v", m.Proposal))
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

	// Timestamp
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// Timestamp
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

	// AssetType
	if err := WriteFixedChar(buf, m.AssetType, 3); err != nil {
		return nil, err
	}

	// AssetID
	if err := WriteFixedChar(buf, m.AssetID, 32); err != nil {
		return nil, err
	}

	// VoteTxnID
	if err := write(buf, m.VoteTxnID); err != nil {
		return nil, err
	}

	// Vote
	if err := WriteVarChar(buf, m.Vote, 8); err != nil {
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// AssetType
	{
		var err error
		m.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// AssetID
	{
		var err error
		m.AssetID, err = ReadFixedChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// VoteTxnID
	if err := read(buf, &m.VoteTxnID); err != nil {
		return 0, err
	}

	// Vote
	{
		var err error
		m.Vote, err = ReadVarChar(buf, 8)
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

	// Timestamp
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// Timestamp
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
	Header           Header      // Common header data for all actions
	AssetType        string      // eg. Share, Bond, Ticket
	AssetID          string      // Randomly generated base58 string.  Each Asset ID should be unique.  However, a Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type can be the leading bytes - a convention - to make it easy to identify that it is a token by humans. If its a Contract vote then can be null.
	Proposal         bool        // 1 for a Proposal, 0 for an initiative that is requesting changes to specific subfields for modification. If this field is true, the subfields should be empty.  The smart contract cannot interpret the results of a vote when Proposal = 1.  All meaning is interpreted by the token owners and smart contract simply facilates the record keeping.  When Proposal = 0, the smart contract always assumes the first choice is a 'yes', or 'pass', if the threshold is met, and will process the proposed changes accordingly.
	ProposedChanges  []Amendment // Each element contains details of which fields to modify, or delete. Because the number of fields in a Contract and Asset is dynamic due to some fields being able to be repeated, the index value of the field needs to be calculated against the Contract or Asset the changes are to apply to. In the event of a Vote being created from this Initiative, the changes will be applied to the version of the Contract or Asset at that time.
	VoteTxnID        [32]byte    // Link to the Vote Action txn.
	VoteOptionsCount uint8       // Number of Vote Options to follow.
	OptionXTally     uint64      // Number of valid votes counted for Option X
	Result           string      // Length 1-255 bytes. 0 is not valid. The Option with the most votes. In the event of a draw for 1st place, all winning options are listed.
	Timestamp        uint64      // Timestamp in nanoseconds of when the smart contract created the action.
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

	// AssetType
	if err := WriteFixedChar(buf, m.AssetType, 3); err != nil {
		return nil, err
	}

	// AssetID
	if err := WriteFixedChar(buf, m.AssetID, 32); err != nil {
		return nil, err
	}

	// Proposal
	if err := write(buf, m.Proposal); err != nil {
		return nil, err
	}

	// ProposedChanges
	if err := WriteVariableSize(buf, uint64(len(m.ProposedChanges)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range m.ProposedChanges {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	// VoteTxnID
	if err := write(buf, m.VoteTxnID); err != nil {
		return nil, err
	}

	// VoteOptionsCount
	if err := write(buf, m.VoteOptionsCount); err != nil {
		return nil, err
	}

	// OptionXTally
	if err := write(buf, m.OptionXTally); err != nil {
		return nil, err
	}

	// Result
	if err := WriteVarChar(buf, m.Result, 8); err != nil {
		return nil, err
	}

	// Timestamp
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// AssetType
	{
		var err error
		m.AssetType, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// AssetID
	{
		var err error
		m.AssetID, err = ReadFixedChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// Proposal
	if err := read(buf, &m.Proposal); err != nil {
		return 0, err
	}

	// ProposedChanges
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		m.ProposedChanges = make([]Amendment, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue Amendment
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			m.ProposedChanges = append(m.ProposedChanges, newValue)
		}
	}

	// VoteTxnID
	if err := read(buf, &m.VoteTxnID); err != nil {
		return 0, err
	}

	// VoteOptionsCount
	if err := read(buf, &m.VoteOptionsCount); err != nil {
		return 0, err
	}

	// OptionXTally
	if err := read(buf, &m.OptionXTally); err != nil {
		return 0, err
	}

	// Result
	{
		var err error
		m.Result, err = ReadVarChar(buf, 8)
		if err != nil {
			return 0, err
		}
	}

	// Timestamp
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
	Header         Header   // Common header data for all actions
	AddressIndexes []uint16 // Associates the message to a particular output by the index.
	MessageType    string   // Potential for up to 65,535 different message types
	MessagePayload string   // Length only limited by the Bitcoin protocol. Public or private (RSA public key, Diffie-Hellman). Issuers/Contracts can send the signifying amount of satoshis to themselves for public announcements or private 'notes' if encrypted. See Message Types for a full list of potential use cases.

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

	// AddressIndexes
	if err := WriteVariableSize(buf, uint64(len(m.AddressIndexes)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range m.AddressIndexes {
		if err := write(buf, value); err != nil {
			return nil, err
		}
	}

	// MessageType
	if err := WriteFixedChar(buf, m.MessageType, 2); err != nil {
		return nil, err
	}

	// MessagePayload
	if err := WriteVarChar(buf, m.MessagePayload, 32); err != nil {
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// AddressIndexes
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		m.AddressIndexes = make([]uint16, size, size)
		if err := read(buf, &m.AddressIndexes); err != nil {
			return 0, err
		}
	}

	// MessageType
	{
		var err error
		m.MessageType, err = ReadFixedChar(buf, 2)
		if err != nil {
			return 0, err
		}
	}

	// MessagePayload
	{
		var err error
		m.MessagePayload, err = ReadVarChar(buf, 32)
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
	Header         Header   // Common header data for all actions
	AddressIndexes []uint16 // Associates the message to a particular output by the index.
	RejectionType  uint8    // Classifies the rejection by a type.
	MessagePayload string   // Length 0-65,535 bytes. Message that explains the reasoning for a rejection, if needed.  Most rejection types will be captured by the Rejection Type Subfield.
	Timestamp      uint64   // Timestamp in nanoseconds of when the smart contract created the action.
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

	// AddressIndexes
	if err := WriteVariableSize(buf, uint64(len(m.AddressIndexes)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range m.AddressIndexes {
		if err := write(buf, value); err != nil {
			return nil, err
		}
	}

	// RejectionType
	if err := write(buf, m.RejectionType); err != nil {
		return nil, err
	}

	// MessagePayload
	if err := WriteVarChar(buf, m.MessagePayload, 32); err != nil {
		return nil, err
	}

	// Timestamp
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// AddressIndexes
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		m.AddressIndexes = make([]uint16, size, size)
		if err := read(buf, &m.AddressIndexes); err != nil {
			return 0, err
		}
	}

	// RejectionType
	if err := read(buf, &m.RejectionType); err != nil {
		return 0, err
	}

	// MessagePayload
	{
		var err error
		m.MessagePayload, err = ReadVarChar(buf, 32)
		if err != nil {
			return 0, err
		}
	}

	// Timestamp
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

	// Message
	if err := WriteVarChar(buf, m.Message, 32); err != nil {
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// Message
	{
		var err error
		m.Message, err = ReadVarChar(buf, 32)
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

	// Message
	if err := WriteVarChar(buf, m.Message, 32); err != nil {
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// Message
	{
		var err error
		m.Message, err = ReadVarChar(buf, 32)
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

	// Message
	if err := WriteVarChar(buf, m.Message, 32); err != nil {
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// Message
	{
		var err error
		m.Message, err = ReadVarChar(buf, 32)
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

	// Message
	if err := WriteVarChar(buf, m.Message, 32); err != nil {
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// Message
	{
		var err error
		m.Message, err = ReadVarChar(buf, 32)
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
	Assets              []AssetTransfer // The Assets involved in the Transfer Action.
	OfferExpiry         uint64          // This prevents any party from holding on to the partially signed message as a form of an option.  Eg. the exchange at this price is valid for 30 mins.
	ExchangeFeeCurrency string          // BSV, USD, AUD, EUR, etc.
	ExchangeFeeVar      float32         // Percent of the value of the transaction
	ExchangeFeeFixed    float32         // Fixed fee
	ExchangeFeeAddress  Address         // Identifies the public address that the exchange fee should be paid to.
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

	// Assets
	if err := WriteVariableSize(buf, uint64(len(m.Assets)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range m.Assets {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	// OfferExpiry
	if err := write(buf, m.OfferExpiry); err != nil {
		return nil, err
	}

	// ExchangeFeeCurrency
	if err := WriteFixedChar(buf, m.ExchangeFeeCurrency, 3); err != nil {
		return nil, err
	}

	// ExchangeFeeVar
	if err := write(buf, m.ExchangeFeeVar); err != nil {
		return nil, err
	}

	// ExchangeFeeFixed
	if err := write(buf, m.ExchangeFeeFixed); err != nil {
		return nil, err
	}

	// ExchangeFeeAddress
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// Assets
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		m.Assets = make([]AssetTransfer, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue AssetTransfer
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			m.Assets = append(m.Assets, newValue)
		}
	}

	// OfferExpiry
	if err := read(buf, &m.OfferExpiry); err != nil {
		return 0, err
	}

	// ExchangeFeeCurrency
	{
		var err error
		m.ExchangeFeeCurrency, err = ReadFixedChar(buf, 3)
		if err != nil {
			return 0, err
		}
	}

	// ExchangeFeeVar
	if err := read(buf, &m.ExchangeFeeVar); err != nil {
		return 0, err
	}

	// ExchangeFeeFixed
	if err := read(buf, &m.ExchangeFeeFixed); err != nil {
		return 0, err
	}

	// ExchangeFeeAddress
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
	vals = append(vals, fmt.Sprintf("Assets:%#+v", m.Assets))
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
	Header    Header            // Common header data for all actions
	Assets    []AssetSettlement // The Assets settled by the transfer action.
	Timestamp uint64            // Timestamp in nanoseconds of when the smart contract created the action.
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

	// Assets
	if err := WriteVariableSize(buf, uint64(len(m.Assets)), 0, 8); err != nil {
		return nil, err
	}
	for _, value := range m.Assets {
		b, err := value.Serialize()
		if err != nil {
			return nil, err
		}

		if err := write(buf, b); err != nil {
			return nil, err
		}
	}

	// Timestamp
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

	// Header
	if err := m.Header.Write(buf); err != nil {
		return 0, err
	}

	// Assets
	{
		size, err := ReadVariableSize(buf, 0, 8)
		if err != nil {
			return 0, err
		}
		m.Assets = make([]AssetSettlement, 0, size)
		for i := uint64(0); i < size; i++ {
			var newValue AssetSettlement
			if err := newValue.Write(buf); err != nil {
				return 0, err
			}

			m.Assets = append(m.Assets, newValue)
		}
	}

	// Timestamp
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
	vals = append(vals, fmt.Sprintf("Assets:%#+v", m.Assets))
	vals = append(vals, fmt.Sprintf("Timestamp:%#+v", m.Timestamp))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
