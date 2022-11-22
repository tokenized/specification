package actions

import (
	"encoding/binary"
	"io"

	"github.com/pkg/errors"
	"github.com/tokenized/pkg/bitcoin"
)

// WriteDeterministic writes data from a ContractOffer in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *ContractOffer) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write([]byte(a.ContractName)); err != nil {
		return errors.Wrap(err, "ContractName")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.BodyOfAgreementType)); err != nil {
		return errors.Wrap(err, "BodyOfAgreementType")
	}

	if _, err := w.Write(a.BodyOfAgreement); err != nil {
		return errors.Wrap(err, "BodyOfAgreement")
	}

	for i, item := range a.SupportingDocs {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "SupportingDocs %d", i)
		}
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.ContractExpiration)); err != nil {
		return errors.Wrap(err, "ContractExpiration")
	}

	if _, err := w.Write([]byte(a.ContractURI)); err != nil {
		return errors.Wrap(err, "ContractURI")
	}

	if err := a.Issuer.WriteDeterministic(w); err != nil {
		return errors.Wrap(err, "Issuer")
	}

	if err := binary.Write(w, binary.LittleEndian, a.ContractOperatorIncluded); err != nil {
		return errors.Wrap(err, "ContractOperatorIncluded")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.ContractFee)); err != nil {
		return errors.Wrap(err, "ContractFee")
	}

	for i, item := range a.VotingSystems {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "VotingSystems %d", i)
		}
	}

	if _, err := w.Write(a.ContractPermissions); err != nil {
		return errors.Wrap(err, "ContractPermissions")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.RestrictedQtyInstruments)); err != nil {
		return errors.Wrap(err, "RestrictedQtyInstruments")
	}

	if err := binary.Write(w, binary.LittleEndian, a.AdministrationProposal); err != nil {
		return errors.Wrap(err, "AdministrationProposal")
	}

	if err := binary.Write(w, binary.LittleEndian, a.HolderProposal); err != nil {
		return errors.Wrap(err, "HolderProposal")
	}

	for i, item := range a.Oracles {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Oracles %d", i)
		}
	}

	if _, err := w.Write(a.MasterAddress); err != nil {
		return errors.Wrap(err, "MasterAddress")
	}

	if _, err := w.Write(a.EntityContract); err != nil {
		return errors.Wrap(err, "EntityContract")
	}

	if _, err := w.Write(a.OperatorEntityContract); err != nil {
		return errors.Wrap(err, "OperatorEntityContract")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.ContractType)); err != nil {
		return errors.Wrap(err, "ContractType")
	}

	for i, item := range a.Services {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Services %d", i)
		}
	}

	for i, item := range a.AdminIdentityCertificates {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "AdminIdentityCertificates %d", i)
		}
	}

	if _, err := w.Write([]byte(a.GoverningLaw)); err != nil {
		return errors.Wrap(err, "GoverningLaw")
	}

	if _, err := w.Write([]byte(a.Jurisdiction)); err != nil {
		return errors.Wrap(err, "Jurisdiction")
	}

	return nil
}

// WriteDeterministic writes data from a ContractFormation in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *ContractFormation) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write([]byte(a.ContractName)); err != nil {
		return errors.Wrap(err, "ContractName")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.BodyOfAgreementType)); err != nil {
		return errors.Wrap(err, "BodyOfAgreementType")
	}

	if _, err := w.Write(a.BodyOfAgreement); err != nil {
		return errors.Wrap(err, "BodyOfAgreement")
	}

	for i, item := range a.SupportingDocs {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "SupportingDocs %d", i)
		}
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.ContractExpiration)); err != nil {
		return errors.Wrap(err, "ContractExpiration")
	}

	if _, err := w.Write([]byte(a.ContractURI)); err != nil {
		return errors.Wrap(err, "ContractURI")
	}

	if err := a.Issuer.WriteDeterministic(w); err != nil {
		return errors.Wrap(err, "Issuer")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.ContractFee)); err != nil {
		return errors.Wrap(err, "ContractFee")
	}

	for i, item := range a.VotingSystems {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "VotingSystems %d", i)
		}
	}

	if _, err := w.Write(a.ContractPermissions); err != nil {
		return errors.Wrap(err, "ContractPermissions")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.RestrictedQtyInstruments)); err != nil {
		return errors.Wrap(err, "RestrictedQtyInstruments")
	}

	if err := binary.Write(w, binary.LittleEndian, a.AdministrationProposal); err != nil {
		return errors.Wrap(err, "AdministrationProposal")
	}

	if err := binary.Write(w, binary.LittleEndian, a.HolderProposal); err != nil {
		return errors.Wrap(err, "HolderProposal")
	}

	for i, item := range a.Oracles {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Oracles %d", i)
		}
	}

	if _, err := w.Write(a.MasterAddress); err != nil {
		return errors.Wrap(err, "MasterAddress")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.ContractRevision)); err != nil {
		return errors.Wrap(err, "ContractRevision")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Timestamp)); err != nil {
		return errors.Wrap(err, "Timestamp")
	}

	if _, err := w.Write(a.EntityContract); err != nil {
		return errors.Wrap(err, "EntityContract")
	}

	if _, err := w.Write(a.OperatorEntityContract); err != nil {
		return errors.Wrap(err, "OperatorEntityContract")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.ContractType)); err != nil {
		return errors.Wrap(err, "ContractType")
	}

	for i, item := range a.Services {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Services %d", i)
		}
	}

	for i, item := range a.AdminIdentityCertificates {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "AdminIdentityCertificates %d", i)
		}
	}

	if _, err := w.Write(a.AdminAddress); err != nil {
		return errors.Wrap(err, "AdminAddress")
	}

	if _, err := w.Write(a.OperatorAddress); err != nil {
		return errors.Wrap(err, "OperatorAddress")
	}

	if _, err := w.Write([]byte(a.GoverningLaw)); err != nil {
		return errors.Wrap(err, "GoverningLaw")
	}

	if _, err := w.Write([]byte(a.Jurisdiction)); err != nil {
		return errors.Wrap(err, "Jurisdiction")
	}

	return nil
}

// WriteDeterministic writes data from a ContractAmendment in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *ContractAmendment) WriteDeterministic(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, a.ChangeAdministrationAddress); err != nil {
		return errors.Wrap(err, "ChangeAdministrationAddress")
	}

	if err := binary.Write(w, binary.LittleEndian, a.ChangeOperatorAddress); err != nil {
		return errors.Wrap(err, "ChangeOperatorAddress")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.ContractRevision)); err != nil {
		return errors.Wrap(err, "ContractRevision")
	}

	for i, item := range a.Amendments {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Amendments %d", i)
		}
	}

	if _, err := w.Write(a.RefTxID[:]); err != nil {
		return errors.Wrap(err, "RefTxID")
	}

	return nil
}

// WriteDeterministic writes data from a StaticContractFormation in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *StaticContractFormation) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write([]byte(a.ContractName)); err != nil {
		return errors.Wrap(err, "ContractName")
	}

	if _, err := w.Write(a.ContractCode[:]); err != nil {
		return errors.Wrap(err, "ContractCode")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.BodyOfAgreementType)); err != nil {
		return errors.Wrap(err, "BodyOfAgreementType")
	}

	if _, err := w.Write(a.BodyOfAgreement); err != nil {
		return errors.Wrap(err, "BodyOfAgreement")
	}

	if _, err := w.Write([]byte(a.ContractType)); err != nil {
		return errors.Wrap(err, "ContractType")
	}

	for i, item := range a.SupportingDocs {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "SupportingDocs %d", i)
		}
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.ContractRevision)); err != nil {
		return errors.Wrap(err, "ContractRevision")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.EffectiveDate)); err != nil {
		return errors.Wrap(err, "EffectiveDate")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.ContractExpiration)); err != nil {
		return errors.Wrap(err, "ContractExpiration")
	}

	if _, err := w.Write([]byte(a.ContractURI)); err != nil {
		return errors.Wrap(err, "ContractURI")
	}

	if _, err := w.Write(a.PrevRevTxID[:]); err != nil {
		return errors.Wrap(err, "PrevRevTxID")
	}

	for i, item := range a.Entities {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Entities %d", i)
		}
	}

	if err := a.EntityOracle.WriteDeterministic(w); err != nil {
		return errors.Wrap(err, "EntityOracle")
	}

	if _, err := w.Write(a.EntityOracleSignature); err != nil {
		return errors.Wrap(err, "EntityOracleSignature")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.EntityOracleSigBlockHeight)); err != nil {
		return errors.Wrap(err, "EntityOracleSigBlockHeight")
	}

	if _, err := w.Write([]byte(a.GoverningLaw)); err != nil {
		return errors.Wrap(err, "GoverningLaw")
	}

	if _, err := w.Write([]byte(a.Jurisdiction)); err != nil {
		return errors.Wrap(err, "Jurisdiction")
	}

	return nil
}

// WriteDeterministic writes data from a ContractAddressChange in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *ContractAddressChange) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write(a.NewContractAddress); err != nil {
		return errors.Wrap(err, "NewContractAddress")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Timestamp)); err != nil {
		return errors.Wrap(err, "Timestamp")
	}

	return nil
}

// WriteDeterministic writes data from a BodyOfAgreementOffer in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *BodyOfAgreementOffer) WriteDeterministic(w io.Writer) error {
	for i, item := range a.Chapters {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Chapters %d", i)
		}
	}

	for i, item := range a.Definitions {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Definitions %d", i)
		}
	}

	return nil
}

// WriteDeterministic writes data from a BodyOfAgreementFormation in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *BodyOfAgreementFormation) WriteDeterministic(w io.Writer) error {
	for i, item := range a.Chapters {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Chapters %d", i)
		}
	}

	for i, item := range a.Definitions {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Definitions %d", i)
		}
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Revision)); err != nil {
		return errors.Wrap(err, "Revision")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Timestamp)); err != nil {
		return errors.Wrap(err, "Timestamp")
	}

	return nil
}

// WriteDeterministic writes data from a BodyOfAgreementAmendment in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *BodyOfAgreementAmendment) WriteDeterministic(w io.Writer) error {
	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Revision)); err != nil {
		return errors.Wrap(err, "Revision")
	}

	for i, item := range a.Amendments {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Amendments %d", i)
		}
	}

	if _, err := w.Write(a.RefTxID[:]); err != nil {
		return errors.Wrap(err, "RefTxID")
	}

	return nil
}

// WriteDeterministic writes data from a InstrumentDefinition in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *InstrumentDefinition) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write(a.InstrumentPermissions); err != nil {
		return errors.Wrap(err, "InstrumentPermissions")
	}

	if err := binary.Write(w, binary.LittleEndian, a.EnforcementOrdersPermitted); err != nil {
		return errors.Wrap(err, "EnforcementOrdersPermitted")
	}

	if err := binary.Write(w, binary.LittleEndian, a.VotingRights); err != nil {
		return errors.Wrap(err, "VotingRights")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.VoteMultiplier)); err != nil {
		return errors.Wrap(err, "VoteMultiplier")
	}

	if err := binary.Write(w, binary.LittleEndian, a.AdministrationProposal); err != nil {
		return errors.Wrap(err, "AdministrationProposal")
	}

	if err := binary.Write(w, binary.LittleEndian, a.HolderProposal); err != nil {
		return errors.Wrap(err, "HolderProposal")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.InstrumentModificationGovernance)); err != nil {
		return errors.Wrap(err, "InstrumentModificationGovernance")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.AuthorizedTokenQty)); err != nil {
		return errors.Wrap(err, "AuthorizedTokenQty")
	}

	if _, err := w.Write([]byte(a.InstrumentType)); err != nil {
		return errors.Wrap(err, "InstrumentType")
	}

	if _, err := w.Write(a.InstrumentPayload); err != nil {
		return errors.Wrap(err, "InstrumentPayload")
	}

	for i, item := range a.TradeRestrictions {
		if _, err := w.Write([]byte(item)); err != nil {
			return errors.Wrapf(err, "TradeRestrictions %d", i)
		}
	}

	return nil
}

// WriteDeterministic writes data from a InstrumentCreation in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *InstrumentCreation) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write(a.InstrumentCode[:]); err != nil {
		return errors.Wrap(err, "InstrumentCode")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.InstrumentIndex)); err != nil {
		return errors.Wrap(err, "InstrumentIndex")
	}

	if _, err := w.Write(a.InstrumentPermissions); err != nil {
		return errors.Wrap(err, "InstrumentPermissions")
	}

	if err := binary.Write(w, binary.LittleEndian, a.EnforcementOrdersPermitted); err != nil {
		return errors.Wrap(err, "EnforcementOrdersPermitted")
	}

	if err := binary.Write(w, binary.LittleEndian, a.VotingRights); err != nil {
		return errors.Wrap(err, "VotingRights")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.VoteMultiplier)); err != nil {
		return errors.Wrap(err, "VoteMultiplier")
	}

	if err := binary.Write(w, binary.LittleEndian, a.AdministrationProposal); err != nil {
		return errors.Wrap(err, "AdministrationProposal")
	}

	if err := binary.Write(w, binary.LittleEndian, a.HolderProposal); err != nil {
		return errors.Wrap(err, "HolderProposal")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.InstrumentModificationGovernance)); err != nil {
		return errors.Wrap(err, "InstrumentModificationGovernance")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.AuthorizedTokenQty)); err != nil {
		return errors.Wrap(err, "AuthorizedTokenQty")
	}

	if _, err := w.Write([]byte(a.InstrumentType)); err != nil {
		return errors.Wrap(err, "InstrumentType")
	}

	if _, err := w.Write(a.InstrumentPayload); err != nil {
		return errors.Wrap(err, "InstrumentPayload")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.InstrumentRevision)); err != nil {
		return errors.Wrap(err, "InstrumentRevision")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Timestamp)); err != nil {
		return errors.Wrap(err, "Timestamp")
	}

	for i, item := range a.TradeRestrictions {
		if _, err := w.Write([]byte(item)); err != nil {
			return errors.Wrapf(err, "TradeRestrictions %d", i)
		}
	}

	return nil
}

// WriteDeterministic writes data from a InstrumentModification in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *InstrumentModification) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write([]byte(a.InstrumentType)); err != nil {
		return errors.Wrap(err, "InstrumentType")
	}

	if _, err := w.Write(a.InstrumentCode[:]); err != nil {
		return errors.Wrap(err, "InstrumentCode")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.InstrumentRevision)); err != nil {
		return errors.Wrap(err, "InstrumentRevision")
	}

	for i, item := range a.Amendments {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Amendments %d", i)
		}
	}

	if _, err := w.Write(a.RefTxID[:]); err != nil {
		return errors.Wrap(err, "RefTxID")
	}

	return nil
}

// WriteDeterministic writes data from a Transfer in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *Transfer) WriteDeterministic(w io.Writer) error {
	for i, item := range a.Instruments {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Instruments %d", i)
		}
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.OfferExpiry)); err != nil {
		return errors.Wrap(err, "OfferExpiry")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.ExchangeFee)); err != nil {
		return errors.Wrap(err, "ExchangeFee")
	}

	if _, err := w.Write(a.ExchangeFeeAddress); err != nil {
		return errors.Wrap(err, "ExchangeFeeAddress")
	}

	return nil
}

// WriteDeterministic writes data from a Settlement in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *Settlement) WriteDeterministic(w io.Writer) error {
	for i, item := range a.Instruments {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Instruments %d", i)
		}
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Timestamp)); err != nil {
		return errors.Wrap(err, "Timestamp")
	}

	return nil
}

// WriteDeterministic writes data from a RectificationSettlement in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *RectificationSettlement) WriteDeterministic(w io.Writer) error {
	if err := a.Transfer.WriteDeterministic(w); err != nil {
		return errors.Wrap(err, "Transfer")
	}

	for i, item := range a.Instruments {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Instruments %d", i)
		}
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Timestamp)); err != nil {
		return errors.Wrap(err, "Timestamp")
	}

	return nil
}

// WriteDeterministic writes data from a Proposal in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *Proposal) WriteDeterministic(w io.Writer) error {
	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Type)); err != nil {
		return errors.Wrap(err, "Type")
	}

	if _, err := w.Write([]byte(a.InstrumentType)); err != nil {
		return errors.Wrap(err, "InstrumentType")
	}

	if _, err := w.Write(a.InstrumentCode[:]); err != nil {
		return errors.Wrap(err, "InstrumentCode")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.VoteSystem)); err != nil {
		return errors.Wrap(err, "VoteSystem")
	}

	for i, item := range a.ProposedAmendments {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "ProposedAmendments %d", i)
		}
	}

	if _, err := w.Write([]byte(a.VoteOptions)); err != nil {
		return errors.Wrap(err, "VoteOptions")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.VoteMax)); err != nil {
		return errors.Wrap(err, "VoteMax")
	}

	if _, err := w.Write([]byte(a.ProposalDescription)); err != nil {
		return errors.Wrap(err, "ProposalDescription")
	}

	if _, err := w.Write(a.ProposalDocumentHash[:]); err != nil {
		return errors.Wrap(err, "ProposalDocumentHash")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.VoteCutOffTimestamp)); err != nil {
		return errors.Wrap(err, "VoteCutOffTimestamp")
	}

	return nil
}

// WriteDeterministic writes data from a Vote in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *Vote) WriteDeterministic(w io.Writer) error {
	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Timestamp)); err != nil {
		return errors.Wrap(err, "Timestamp")
	}

	return nil
}

// WriteDeterministic writes data from a BallotCast in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *BallotCast) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write(a.VoteTxId[:]); err != nil {
		return errors.Wrap(err, "VoteTxId")
	}

	if _, err := w.Write([]byte(a.Vote)); err != nil {
		return errors.Wrap(err, "Vote")
	}

	return nil
}

// WriteDeterministic writes data from a BallotCounted in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *BallotCounted) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write(a.VoteTxId[:]); err != nil {
		return errors.Wrap(err, "VoteTxId")
	}

	if _, err := w.Write([]byte(a.Vote)); err != nil {
		return errors.Wrap(err, "Vote")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Quantity)); err != nil {
		return errors.Wrap(err, "Quantity")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Timestamp)); err != nil {
		return errors.Wrap(err, "Timestamp")
	}

	return nil
}

// WriteDeterministic writes data from a Result in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *Result) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write([]byte(a.InstrumentType)); err != nil {
		return errors.Wrap(err, "InstrumentType")
	}

	if _, err := w.Write(a.InstrumentCode[:]); err != nil {
		return errors.Wrap(err, "InstrumentCode")
	}

	for i, item := range a.ProposedAmendments {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "ProposedAmendments %d", i)
		}
	}

	if _, err := w.Write(a.VoteTxId[:]); err != nil {
		return errors.Wrap(err, "VoteTxId")
	}

	for i, item := range a.OptionTally {
		if err := bitcoin.WriteBase128VarInt(w, uint64(item)); err != nil {
			return errors.Wrapf(err, "OptionTally %d", i)
		}
	}

	if _, err := w.Write([]byte(a.Result)); err != nil {
		return errors.Wrap(err, "Result")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Timestamp)); err != nil {
		return errors.Wrap(err, "Timestamp")
	}

	return nil
}

// WriteDeterministic writes data from a Order in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *Order) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write([]byte(a.ComplianceAction)); err != nil {
		return errors.Wrap(err, "ComplianceAction")
	}

	if _, err := w.Write([]byte(a.InstrumentType)); err != nil {
		return errors.Wrap(err, "InstrumentType")
	}

	if _, err := w.Write(a.InstrumentCode[:]); err != nil {
		return errors.Wrap(err, "InstrumentCode")
	}

	for i, item := range a.TargetAddresses {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "TargetAddresses %d", i)
		}
	}

	if _, err := w.Write(a.FreezeTxId[:]); err != nil {
		return errors.Wrap(err, "FreezeTxId")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.FreezePeriod)); err != nil {
		return errors.Wrap(err, "FreezePeriod")
	}

	if _, err := w.Write(a.DepositAddress); err != nil {
		return errors.Wrap(err, "DepositAddress")
	}

	if _, err := w.Write([]byte(a.AuthorityName)); err != nil {
		return errors.Wrap(err, "AuthorityName")
	}

	if _, err := w.Write(a.AuthorityPublicKey); err != nil {
		return errors.Wrap(err, "AuthorityPublicKey")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.SignatureAlgorithm)); err != nil {
		return errors.Wrap(err, "SignatureAlgorithm")
	}

	if _, err := w.Write(a.OrderSignature); err != nil {
		return errors.Wrap(err, "OrderSignature")
	}

	for i, item := range a.BitcoinDispersions {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "BitcoinDispersions %d", i)
		}
	}

	if _, err := w.Write([]byte(a.Message)); err != nil {
		return errors.Wrap(err, "Message")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.SupportingEvidenceFormat)); err != nil {
		return errors.Wrap(err, "SupportingEvidenceFormat")
	}

	if _, err := w.Write(a.SupportingEvidence); err != nil {
		return errors.Wrap(err, "SupportingEvidence")
	}

	for i, item := range a.ReferenceTransactions {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "ReferenceTransactions %d", i)
		}
	}

	return nil
}

// WriteDeterministic writes data from a Freeze in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *Freeze) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write([]byte(a.InstrumentType)); err != nil {
		return errors.Wrap(err, "InstrumentType")
	}

	if _, err := w.Write(a.InstrumentCode[:]); err != nil {
		return errors.Wrap(err, "InstrumentCode")
	}

	for i, item := range a.Quantities {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Quantities %d", i)
		}
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.FreezePeriod)); err != nil {
		return errors.Wrap(err, "FreezePeriod")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Timestamp)); err != nil {
		return errors.Wrap(err, "Timestamp")
	}

	return nil
}

// WriteDeterministic writes data from a Thaw in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *Thaw) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write(a.FreezeTxId[:]); err != nil {
		return errors.Wrap(err, "FreezeTxId")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Timestamp)); err != nil {
		return errors.Wrap(err, "Timestamp")
	}

	return nil
}

// WriteDeterministic writes data from a Confiscation in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *Confiscation) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write([]byte(a.InstrumentType)); err != nil {
		return errors.Wrap(err, "InstrumentType")
	}

	if _, err := w.Write(a.InstrumentCode[:]); err != nil {
		return errors.Wrap(err, "InstrumentCode")
	}

	for i, item := range a.Quantities {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Quantities %d", i)
		}
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.DepositQty)); err != nil {
		return errors.Wrap(err, "DepositQty")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Timestamp)); err != nil {
		return errors.Wrap(err, "Timestamp")
	}

	return nil
}

// WriteDeterministic writes data from a DeprecatedReconciliation in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *DeprecatedReconciliation) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write([]byte(a.InstrumentType)); err != nil {
		return errors.Wrap(err, "InstrumentType")
	}

	if _, err := w.Write(a.InstrumentCode[:]); err != nil {
		return errors.Wrap(err, "InstrumentCode")
	}

	for i, item := range a.Quantities {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Quantities %d", i)
		}
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Timestamp)); err != nil {
		return errors.Wrap(err, "Timestamp")
	}

	return nil
}

// WriteDeterministic writes data from a Establishment in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *Establishment) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write([]byte(a.Message)); err != nil {
		return errors.Wrap(err, "Message")
	}

	return nil
}

// WriteDeterministic writes data from a Addition in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *Addition) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write([]byte(a.Message)); err != nil {
		return errors.Wrap(err, "Message")
	}

	return nil
}

// WriteDeterministic writes data from a Alteration in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *Alteration) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write(a.EntryTxID[:]); err != nil {
		return errors.Wrap(err, "EntryTxID")
	}

	if _, err := w.Write([]byte(a.Message)); err != nil {
		return errors.Wrap(err, "Message")
	}

	return nil
}

// WriteDeterministic writes data from a Removal in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *Removal) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write(a.EntryTxID[:]); err != nil {
		return errors.Wrap(err, "EntryTxID")
	}

	if _, err := w.Write([]byte(a.Message)); err != nil {
		return errors.Wrap(err, "Message")
	}

	return nil
}

// WriteDeterministic writes data from a Message in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *Message) WriteDeterministic(w io.Writer) error {
	for i, item := range a.SenderIndexes {
		if err := bitcoin.WriteBase128VarInt(w, uint64(item)); err != nil {
			return errors.Wrapf(err, "SenderIndexes %d", i)
		}
	}

	for i, item := range a.ReceiverIndexes {
		if err := bitcoin.WriteBase128VarInt(w, uint64(item)); err != nil {
			return errors.Wrapf(err, "ReceiverIndexes %d", i)
		}
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.MessageCode)); err != nil {
		return errors.Wrap(err, "MessageCode")
	}

	if _, err := w.Write(a.MessagePayload); err != nil {
		return errors.Wrap(err, "MessagePayload")
	}

	return nil
}

// WriteDeterministic writes data from a Rejection in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *Rejection) WriteDeterministic(w io.Writer) error {
	for i, item := range a.AddressIndexes {
		if err := bitcoin.WriteBase128VarInt(w, uint64(item)); err != nil {
			return errors.Wrapf(err, "AddressIndexes %d", i)
		}
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.RejectAddressIndex)); err != nil {
		return errors.Wrap(err, "RejectAddressIndex")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.RejectionCode)); err != nil {
		return errors.Wrap(err, "RejectionCode")
	}

	if _, err := w.Write([]byte(a.Message)); err != nil {
		return errors.Wrap(err, "Message")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Timestamp)); err != nil {
		return errors.Wrap(err, "Timestamp")
	}

	return nil
}

// WriteDeterministic writes data from a Administrator in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *AdministratorField) WriteDeterministic(w io.Writer) error {
	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Type)); err != nil {
		return errors.Wrap(err, "Type")
	}

	if _, err := w.Write([]byte(a.Name)); err != nil {
		return errors.Wrap(err, "Name")
	}

	return nil
}

// WriteDeterministic writes data from a AdminIdentityCertificate in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *AdminIdentityCertificateField) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write(a.EntityContract); err != nil {
		return errors.Wrap(err, "EntityContract")
	}

	if _, err := w.Write(a.Signature); err != nil {
		return errors.Wrap(err, "Signature")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.BlockHeight)); err != nil {
		return errors.Wrap(err, "BlockHeight")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Expiration)); err != nil {
		return errors.Wrap(err, "Expiration")
	}

	return nil
}

// WriteDeterministic writes data from a Amendment in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *AmendmentField) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write(a.FieldIndexPath); err != nil {
		return errors.Wrap(err, "FieldIndexPath")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Operation)); err != nil {
		return errors.Wrap(err, "Operation")
	}

	if _, err := w.Write(a.Data); err != nil {
		return errors.Wrap(err, "Data")
	}

	return nil
}

// WriteDeterministic writes data from a InstrumentReceiver in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *InstrumentReceiverField) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write(a.Address); err != nil {
		return errors.Wrap(err, "Address")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Quantity)); err != nil {
		return errors.Wrap(err, "Quantity")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.OracleSigAlgorithm)); err != nil {
		return errors.Wrap(err, "OracleSigAlgorithm")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.OracleIndex)); err != nil {
		return errors.Wrap(err, "OracleIndex")
	}

	if _, err := w.Write(a.OracleConfirmationSig); err != nil {
		return errors.Wrap(err, "OracleConfirmationSig")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.OracleSigBlockHeight)); err != nil {
		return errors.Wrap(err, "OracleSigBlockHeight")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.OracleSigExpiry)); err != nil {
		return errors.Wrap(err, "OracleSigExpiry")
	}

	return nil
}

// WriteDeterministic writes data from a InstrumentSettlement in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *InstrumentSettlementField) WriteDeterministic(w io.Writer) error {
	if err := bitcoin.WriteBase128VarInt(w, uint64(a.ContractIndex)); err != nil {
		return errors.Wrap(err, "ContractIndex")
	}

	if _, err := w.Write([]byte(a.InstrumentType)); err != nil {
		return errors.Wrap(err, "InstrumentType")
	}

	if _, err := w.Write(a.InstrumentCode[:]); err != nil {
		return errors.Wrap(err, "InstrumentCode")
	}

	for i, item := range a.Settlements {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Settlements %d", i)
		}
	}

	return nil
}

// WriteDeterministic writes data from a InstrumentTransfer in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *InstrumentTransferField) WriteDeterministic(w io.Writer) error {
	if err := bitcoin.WriteBase128VarInt(w, uint64(a.ContractIndex)); err != nil {
		return errors.Wrap(err, "ContractIndex")
	}

	if _, err := w.Write([]byte(a.InstrumentType)); err != nil {
		return errors.Wrap(err, "InstrumentType")
	}

	if _, err := w.Write(a.InstrumentCode[:]); err != nil {
		return errors.Wrap(err, "InstrumentCode")
	}

	for i, item := range a.InstrumentSenders {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "InstrumentSenders %d", i)
		}
	}

	for i, item := range a.InstrumentReceivers {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "InstrumentReceivers %d", i)
		}
	}

	return nil
}

// WriteDeterministic writes data from a Chapter in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *ChapterField) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write([]byte(a.Title)); err != nil {
		return errors.Wrap(err, "Title")
	}

	if _, err := w.Write([]byte(a.Preamble)); err != nil {
		return errors.Wrap(err, "Preamble")
	}

	for i, item := range a.Articles {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Articles %d", i)
		}
	}

	return nil
}

// WriteDeterministic writes data from a Clause in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *ClauseField) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write([]byte(a.Title)); err != nil {
		return errors.Wrap(err, "Title")
	}

	if _, err := w.Write([]byte(a.Body)); err != nil {
		return errors.Wrap(err, "Body")
	}

	for i, item := range a.Children {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Children %d", i)
		}
	}

	return nil
}

// WriteDeterministic writes data from a DefinedTerm in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *DefinedTermField) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write([]byte(a.Term)); err != nil {
		return errors.Wrap(err, "Term")
	}

	if _, err := w.Write([]byte(a.Definition)); err != nil {
		return errors.Wrap(err, "Definition")
	}

	return nil
}

// WriteDeterministic writes data from a Document in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *DocumentField) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write([]byte(a.Name)); err != nil {
		return errors.Wrap(err, "Name")
	}

	if _, err := w.Write([]byte(a.Type)); err != nil {
		return errors.Wrap(err, "Type")
	}

	if _, err := w.Write(a.Contents); err != nil {
		return errors.Wrap(err, "Contents")
	}

	return nil
}

// WriteDeterministic writes data from a Entity in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *EntityField) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write([]byte(a.Name)); err != nil {
		return errors.Wrap(err, "Name")
	}

	if _, err := w.Write([]byte(a.Type)); err != nil {
		return errors.Wrap(err, "Type")
	}

	if _, err := w.Write([]byte(a.LEI)); err != nil {
		return errors.Wrap(err, "LEI")
	}

	if _, err := w.Write([]byte(a.UnitNumber)); err != nil {
		return errors.Wrap(err, "UnitNumber")
	}

	if _, err := w.Write([]byte(a.BuildingNumber)); err != nil {
		return errors.Wrap(err, "BuildingNumber")
	}

	if _, err := w.Write([]byte(a.Street)); err != nil {
		return errors.Wrap(err, "Street")
	}

	if _, err := w.Write([]byte(a.SuburbCity)); err != nil {
		return errors.Wrap(err, "SuburbCity")
	}

	if _, err := w.Write([]byte(a.TerritoryStateProvinceCode)); err != nil {
		return errors.Wrap(err, "TerritoryStateProvinceCode")
	}

	if _, err := w.Write([]byte(a.CountryCode)); err != nil {
		return errors.Wrap(err, "CountryCode")
	}

	if _, err := w.Write([]byte(a.PostalZIPCode)); err != nil {
		return errors.Wrap(err, "PostalZIPCode")
	}

	if _, err := w.Write([]byte(a.EmailAddress)); err != nil {
		return errors.Wrap(err, "EmailAddress")
	}

	if _, err := w.Write([]byte(a.PhoneNumber)); err != nil {
		return errors.Wrap(err, "PhoneNumber")
	}

	for i, item := range a.Administration {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Administration %d", i)
		}
	}

	for i, item := range a.Management {
		if err := item.WriteDeterministic(w); err != nil {
			return errors.Wrapf(err, "Management %d", i)
		}
	}

	if _, err := w.Write([]byte(a.DomainName)); err != nil {
		return errors.Wrap(err, "DomainName")
	}

	if _, err := w.Write([]byte(a.PaymailHandle)); err != nil {
		return errors.Wrap(err, "PaymailHandle")
	}

	return nil
}

// WriteDeterministic writes data from a Manager in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *ManagerField) WriteDeterministic(w io.Writer) error {
	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Type)); err != nil {
		return errors.Wrap(err, "Type")
	}

	if _, err := w.Write([]byte(a.Name)); err != nil {
		return errors.Wrap(err, "Name")
	}

	return nil
}

// WriteDeterministic writes data from a Oracle in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *OracleField) WriteDeterministic(w io.Writer) error {

	for i, item := range a.OracleTypes {
		if err := bitcoin.WriteBase128VarInt(w, uint64(item)); err != nil {
			return errors.Wrapf(err, "OracleTypes %d", i)
		}
	}

	if _, err := w.Write(a.EntityContract); err != nil {
		return errors.Wrap(err, "EntityContract")
	}

	return nil
}

// WriteDeterministic writes data from a QuantityIndex in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *QuantityIndexField) WriteDeterministic(w io.Writer) error {
	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Index)); err != nil {
		return errors.Wrap(err, "Index")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Quantity)); err != nil {
		return errors.Wrap(err, "Quantity")
	}

	return nil
}

// WriteDeterministic writes data from a ReferenceTransaction in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *ReferenceTransactionField) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write(a.Transaction); err != nil {
		return errors.Wrap(err, "Transaction")
	}

	for i, item := range a.Outputs {
		if _, err := w.Write(item); err != nil {
			return errors.Wrapf(err, "Outputs %d", i)
		}
	}

	return nil
}

// WriteDeterministic writes data from a Service in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *ServiceField) WriteDeterministic(w io.Writer) error {
	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Type)); err != nil {
		return errors.Wrap(err, "Type")
	}

	if _, err := w.Write([]byte(a.URL)); err != nil {
		return errors.Wrap(err, "URL")
	}

	if _, err := w.Write(a.PublicKey[:]); err != nil {
		return errors.Wrap(err, "PublicKey")
	}

	return nil
}

// WriteDeterministic writes data from a TargetAddress in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *TargetAddressField) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write(a.Address); err != nil {
		return errors.Wrap(err, "Address")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.Quantity)); err != nil {
		return errors.Wrap(err, "Quantity")
	}

	return nil
}

// WriteDeterministic writes data from a VotingSystem in a deterministic way so the data can
// be used to sign an object. The data output can not be parsed back into an object.
func (a *VotingSystemField) WriteDeterministic(w io.Writer) error {
	if _, err := w.Write([]byte(a.Name)); err != nil {
		return errors.Wrap(err, "Name")
	}

	if _, err := w.Write([]byte(a.VoteType)); err != nil {
		return errors.Wrap(err, "VoteType")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.TallyLogic)); err != nil {
		return errors.Wrap(err, "TallyLogic")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.ThresholdPercentage)); err != nil {
		return errors.Wrap(err, "ThresholdPercentage")
	}

	if err := binary.Write(w, binary.LittleEndian, a.VoteMultiplierPermitted); err != nil {
		return errors.Wrap(err, "VoteMultiplierPermitted")
	}

	if err := bitcoin.WriteBase128VarInt(w, uint64(a.HolderProposalFee)); err != nil {
		return errors.Wrap(err, "HolderProposalFee")
	}

	return nil
}
