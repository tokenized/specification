// Package protocol provides base level structs and validation for
// the protocol.
//
// The code in this file is auto-generated. Do not edit it by hand as it will
// be overwritten when code is regenerated.

import {sprintf} from 'sprintf-js';
import _ from '@keyring/util';
import {write, read, ReadVarChar, ReadVariableSize, ReadVarBin, ReadFixedChar,
	WriteVarChar, WriteVariableSize, WriteFixedChar, WriteVarBin, char} from './bytes';
import { TxId, InstrumentCode, Timestamp, ContractCode, PublicKeyHash } from './protocol_types';
import { Document, Amendment, VotingSystem, Oracle, Entity, TargetAddress,
	QuantityIndex, InstrumentTransfer, InstrumentSettlement } from './field_types';
import { Resources } from './resources';
import { OpReturnMessage } from './protocol';

export enum ActionCode {
	// CodeInstrumentDefinition identifies data as a InstrumentDefinition message.
	CodeInstrumentDefinition = 'A1',

	// CodeInstrumentCreation identifies data as a InstrumentCreation message.
	CodeInstrumentCreation = 'A2',

	// CodeInstrumentModification identifies data as a InstrumentModification message.
	CodeInstrumentModification = 'A3',

	// CodeContractOffer identifies data as a ContractOffer message.
	CodeContractOffer = 'C1',

	// CodeContractFormation identifies data as a ContractFormation message.
	CodeContractFormation = 'C2',

	// CodeContractAmendment identifies data as a ContractAmendment message.
	CodeContractAmendment = 'C3',

	// CodeStaticContractFormation identifies data as a
	// StaticContractFormation message.
	CodeStaticContractFormation = 'C4',

	// CodeContractAddressChange identifies data as a ContractAddressChange
	// message.
	CodeContractAddressChange = 'C5',

	// CodeOrder identifies data as a Order message.
	CodeOrder = 'E1',

	// CodeFreeze identifies data as a Freeze message.
	CodeFreeze = 'E2',

	// CodeThaw identifies data as a Thaw message.
	CodeThaw = 'E3',

	// CodeConfiscation identifies data as a Confiscation message.
	CodeConfiscation = 'E4',

	// CodeReconciliation identifies data as a Reconciliation message.
	CodeReconciliation = 'E5',

	// CodeProposal identifies data as a Proposal message.
	CodeProposal = 'G1',

	// CodeVote identifies data as a Vote message.
	CodeVote = 'G2',

	// CodeBallotCast identifies data as a BallotCast message.
	CodeBallotCast = 'G3',

	// CodeBallotCounted identifies data as a BallotCounted message.
	CodeBallotCounted = 'G4',

	// CodeResult identifies data as a Result message.
	CodeResult = 'G5',

	// CodeMessage identifies data as a Message message.
	CodeMessage = 'M1',

	// CodeRejection identifies data as a Rejection message.
	CodeRejection = 'M2',

	// CodeEstablishment identifies data as a Establishment message.
	CodeEstablishment = 'R1',

	// CodeAddition identifies data as a Addition message.
	CodeAddition = 'R2',

	// CodeAlteration identifies data as a Alteration message.
	CodeAlteration = 'R3',

	// CodeRemoval identifies data as a Removal message.
	CodeRemoval = 'R4',

	// CodeTransfer identifies data as a Transfer message.
	CodeTransfer = 'T1',

	// CodeSettlement identifies data as a Settlement message.
	CodeSettlement = 'T2',


	// ComplianceActionFreeze identifies a freeze type
	ComplianceActionFreeze = 'F',

	// ComplianceActionThaw identifies a thaw type
	ComplianceActionThaw = 'T',

	// ComplianceActionConfiscation identifies a confiscation type
	ComplianceActionConfiscation = 'C',

	// ComplianceActionReconciliation identifies a reconcilation type
	ComplianceActionReconciliation = 'R',
}

// TypeMapping holds a mapping of action codes to action types.
export function TypeMapping(code: string): OpReturnMessage {
	switch (code) {
	case ActionCode.CodeInstrumentDefinition:
		return new InstrumentDefinition();
	case ActionCode.CodeInstrumentCreation:
		return new InstrumentCreation();
	case ActionCode.CodeInstrumentModification:
		return new InstrumentModification();
	case ActionCode.CodeContractOffer:
		return new ContractOffer();
	case ActionCode.CodeContractFormation:
		return new ContractFormation();
	case ActionCode.CodeContractAmendment:
		return new ContractAmendment();
	case ActionCode.CodeStaticContractFormation:
		return new StaticContractFormation();
	case ActionCode.CodeContractAddressChange:
		return new ContractAddressChange();
	case ActionCode.CodeOrder:
		return new Order();
	case ActionCode.CodeFreeze:
		return new Freeze();
	case ActionCode.CodeThaw:
		return new Thaw();
	case ActionCode.CodeConfiscation:
		return new Confiscation();
	case ActionCode.CodeReconciliation:
		return new Reconciliation();
	case ActionCode.CodeProposal:
		return new Proposal();
	case ActionCode.CodeVote:
		return new Vote();
	case ActionCode.CodeBallotCast:
		return new BallotCast();
	case ActionCode.CodeBallotCounted:
		return new BallotCounted();
	case ActionCode.CodeResult:
		return new Result();
	case ActionCode.CodeMessage:
		return new Message();
	case ActionCode.CodeRejection:
		return new Rejection();
	case ActionCode.CodeEstablishment:
		return new Establishment();
	case ActionCode.CodeAddition:
		return new Addition();
	case ActionCode.CodeAlteration:
		return new Alteration();
	case ActionCode.CodeRemoval:
		return new Removal();
	case ActionCode.CodeTransfer:
		return new Transfer();
	case ActionCode.CodeSettlement:
		return new Settlement();
	default:
		return null;
	}
}



	// InstrumentDefinition This action is used by the administration to define the
// properties/characteristics of the instrument (token) that it wants to create.
// An instrument has a unique identifier called InstrumentID = InstrumentType +
// base58(InstrumentCode + checksum). An instrument is always linked to a contract
// that is identified by the public address of the Contract wallet.
	export class InstrumentDefinition extends OpReturnMessage {
		type = ActionCode.CodeInstrumentDefinition;
		typeStr = 'InstrumentDefinition';

	
		// 	Three letter character that specifies the instrument type.
		instrument_type;
	
		// 	A set of switches that define the authorization rules for this instrument.
	// See the Authorization Flags documentation for more detail.
		instrument_auth_flags;
	
		// 	Set to true if transfers are permitted between two parties, otherwise
	// set to false to prevent peer-to-peer transfers.
		transfers_permitted;
	
		// 	If specified, the instrument can only be traded within the specified trade
	// restriction zone. For example, AUS would restrict to Australian
	// residents only.
		trade_restrictions;
	
		// 	Set to true if the administration is permitted to make enforcement
	// orders on user tokens and balances, otherwise set to false to disable
	// this feature.
		enforcement_orders_permitted;
	
		// 	When false holders of this instrument will not be able to vote for tokens
	// of this instrument even on voting systems in which vote multiplers are not
	// permitted.
		voting_rights;
	
		// 	Multiplies a vote by the specified integer. Where 1 token is equal to
	// 1 vote with a 1 for vote multipler (normal), 1 token = 3 votes with a
	// multiplier of 3, for example. If zero, then holders of this instrument don't
	// get any votes for their tokens.
		vote_multiplier;
	
		// 	Set to true if the administration is permitted to make proposals
	// outside of the smart contract scope.
		administration_proposal;
	
		// 	Set to true if a holder is permitted to make proposals outside of the
	// smart contract scope.
		holder_proposal;
	
		// 	Supported values: 1 - Contract-wide Instrument Governance. 0 - Instrument-wide
	// Instrument Governance. If a referendum or initiative is used to propose a
	// modification to a subfield controlled by the instrument auth flags, then the
	// vote will either be a contract-wide vote (all instruments vote on the
	// referendum/initiative) or an instrument-wide vote (only this instrument votes on
	// the referendum/initiative) depending on the value in this subfield. The
	// voting system specifies the voting rules.
		instrument_modification_governance;
	
		// 	The number of tokens to issue with this instrument. Set to greater than
	// zero for fungible tokens. If the value is 1 then it becomes a
	// non-fungible token, where the contract would have many instruments. Set to 0
	// to represent a placeholder instrument, where tokens are to be issued later,
	// or to represent a decomissioned instrument where all tokens have been
	// revoked.
		token_qty;
	
		// 	A custom payload that contains meta data about this instrument. Payload
	// structure and length is dependent on the instrument type chosen. See instrument
	// documentation for more details.
		instrument_payload;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeInstrumentDefinition;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// instrument_type (string)
		{
			WriteFixedChar(buf, this.instrument_type, 3);
		}

		// instrument_auth_flags ([]byte)
		{
			WriteVarBin(buf, this.instrument_auth_flags, 8);
		}

		// transfers_permitted (bool)
		{
			write(buf, this.transfers_permitted, 'bool');
		}

		// trade_restrictions ([][3]byte)
		{
			// IsNativeTypeArray [][3]byte
			const type = '[][3]byte'.slice(2);
			WriteVariableSize(buf, this.trade_restrictions.length, 16, 8);
			this.trade_restrictions.forEach(value => {
				write(buf, value, type); // TODO
			});
		}

		// enforcement_orders_permitted (bool)
		{
			write(buf, this.enforcement_orders_permitted, 'bool');
		}

		// voting_rights (bool)
		{
			write(buf, this.voting_rights, 'bool');
		}

		// vote_multiplier (uint8)
		{
			write(buf, this.vote_multiplier, 'uint8');
		}

		// administration_proposal (bool)
		{
			write(buf, this.administration_proposal, 'bool');
		}

		// holder_proposal (bool)
		{
			write(buf, this.holder_proposal, 'bool');
		}

		// instrument_modification_governance (uint8)
		{
			write(buf, this.instrument_modification_governance, 'uint8');
		}

		// token_qty (uint64)
		{
			write(buf, this.token_qty, 'uint64');
		}

		// instrument_payload ([]byte)
		{
			WriteVarBin(buf, this.instrument_payload, 16);
		}

		return buf.buf;
	}

	// write populates the fields in InstrumentDefinition from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// instrument_type (string)
		{
			this.instrument_type = ReadFixedChar(buf, 3);
		}

		// instrument_auth_flags ([]byte)
		{
			this.instrument_auth_flags = ReadVarBin(buf, 8);
		}

		// transfers_permitted (bool)
		{
			this.transfers_permitted = read(buf, 'bool');
		}

		// trade_restrictions ([][3]byte)
		{
			const size = ReadVariableSize(buf, 16, 8);
			const type = '[][3]byte'.slice(2);
			this.trade_restrictions = [...Array(size)].map(() => read(buf, type));
		}

		// enforcement_orders_permitted (bool)
		{
			this.enforcement_orders_permitted = read(buf, 'bool');
		}

		// voting_rights (bool)
		{
			this.voting_rights = read(buf, 'bool');
		}

		// vote_multiplier (uint8)
		{
			this.vote_multiplier = read(buf, 'uint8');
		}

		// administration_proposal (bool)
		{
			this.administration_proposal = read(buf, 'bool');
		}

		// holder_proposal (bool)
		{
			this.holder_proposal = read(buf, 'bool');
		}

		// instrument_modification_governance (uint8)
		{
			this.instrument_modification_governance = read(buf, 'uint8');
		}

		// token_qty (uint64)
		{
			this.token_qty = read(buf, 'uint64');
		}

		// instrument_payload ([]byte)
		{
			this.instrument_payload = ReadVarBin(buf, 16);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// InstrumentType (string)
		{
			if (this.instrument_type.length > 3) {
				return sprintf('fixedchar field instrument_type too long %d/%d', this.instrument_type.length, 3);
			}
		}

		// InstrumentAuthFlags ([]byte)
		{
			if (this.instrument_auth_flags.length >= (2 ** 8)) {
				return sprintf('varbin field instrument_auth_flags too long %d/%d', this.instrument_auth_flags.length, (2 ** 8) - 1);
			}
		}

		// TransfersPermitted (bool) 
		{
		}

		// TradeRestrictions ([][3]byte) 
		{
			if (this.trade_restrictions.length >= (2 ** 16)) {
				return sprintf('list field trade_restrictions has too many items %d/%d', this.trade_restrictions.length, (2 ** 16) - 1);
			}

			let err = null;
			this.trade_restrictions.find((value) => {
				const str = Buffer.from(value).toString('ascii');
				if (!Resources.GetPolityType(str)) {
					err = sprintf('trade_restrictions: Invalid polity value : %s', str);
					return true;
				}
			});
			if (err) return err;
		}

		// EnforcementOrdersPermitted (bool) 
		{
		}

		// VotingRights (bool) 
		{
		}

		// VoteMultiplier (uint8) 
		{
		}

		// AdministrationProposal (bool) 
		{
		}

		// HolderProposal (bool) 
		{
		}

		// InstrumentModificationGovernance (uint8)
		{
			// $field.IntValues [0 1]
			if ( this.instrument_modification_governance !== 0 && this.instrument_modification_governance !== 1) {
				return sprintf('field instrument_modification_governance value is invalid : %d', this.instrument_modification_governance);
			}

		}

		// TokenQty (uint64) 
		{
		}

		// InstrumentPayload ([]byte)
		{
			if (this.instrument_payload.length >= (2 ** 16)) {
				return sprintf('varbin field instrument_payload too long %d/%d', this.instrument_payload.length, (2 ** 16) - 1);
			}
		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('instrument_type:%s', this.instrument_type.toString()));
			vals.push(sprintf('instrument_auth_flags:%s', this.instrument_auth_flags.toString()));
			vals.push(sprintf('transfers_permitted:%s', this.transfers_permitted.toString()));
			vals.push(sprintf('trade_restrictions:%s', this.trade_restrictions.toString()));
			vals.push(sprintf('enforcement_orders_permitted:%s', this.enforcement_orders_permitted.toString()));
			vals.push(sprintf('voting_rights:%s', this.voting_rights.toString()));
			vals.push(sprintf('vote_multiplier:%d', this.vote_multiplier));
			vals.push(sprintf('administration_proposal:%s', this.administration_proposal.toString()));
			vals.push(sprintf('holder_proposal:%s', this.holder_proposal.toString()));
			vals.push(sprintf('instrument_modification_governance:%d', this.instrument_modification_governance));
			vals.push(sprintf('token_qty:%d', this.token_qty));
			vals.push(sprintf('instrument_payload:%s', this.instrument_payload.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for InstrumentDefinition.', e);
			throw e;
		}
		// this.instrument_type = parsed.instrumentType (fixedchar)
		
		// this.instrument_auth_flags = parsed.instrumentAuthFlags (varbin)
		
		// this.transfers_permitted = parsed.transfersPermitted (bool)
		
		// this.trade_restrictions = parsed.tradeRestrictions (Polity[])
		
		// this.enforcement_orders_permitted = parsed.enforcementOrdersPermitted (bool)
		
		// this.voting_rights = parsed.votingRights (bool)
		
		this.vote_multiplier = parsed.voteMultiplier;
		
		// this.administration_proposal = parsed.administrationProposal (bool)
		
		// this.holder_proposal = parsed.holderProposal (bool)
		
		this.instrument_modification_governance = parsed.instrumentModificationGovernance;
		
		this.token_qty = parsed.tokenQty;
		
		// this.instrument_payload = parsed.instrumentPayload (varbin)
		
		return this.Validate();
	}
}

	// InstrumentCreation This action creates an instrument in response to the
// administration's instructions in the Definition Action.
	export class InstrumentCreation extends OpReturnMessage {
		type = ActionCode.CodeInstrumentCreation;
		typeStr = 'InstrumentCreation';

	
		// 	Three letter character that specifies the instrument type.
		instrument_type;
	
		// 	A unique code that is used to identify the instrument. It is generated by
	// hashing the contract public key hash and the instrument index.
	// SHA256(contract PKH + instrument index)
		instrument_code;
	
		// 	The index of the instrument within the contract. First instrument is zero,
	// second is one. Used to derive the instrument code.
		instrument_index;
	
		// 	A set of switches that define the authorization rules for this instrument.
	// See the Authorization Flags documentation for more detail.
		instrument_auth_flags;
	
		// 	Set to true if transfers are permitted between two parties, otherwise
	// set to false to prevent peer-to-peer transfers.
		transfers_permitted;
	
		// 	If specified, the instrument can only be traded within the specified trade
	// restriction zone. For example, AUS would restrict to Australian
	// residents only.
		trade_restrictions;
	
		// 	Set to true if the administration is permitted to make enforcement
	// orders on user tokens and balances, otherwise set to false to disable
	// this feature.
		enforcement_orders_permitted;
	
		// 	When false holders of this instrument will not be able to vote for tokens
	// of this instrument even on voting systems in which vote multiplers are not
	// permitted.
		voting_rights;
	
		// 	Multiplies a vote by the specified integer. Where 1 token is equal to
	// 1 vote with a 1 for vote multipler (normal), 1 token = 3 votes with a
	// multiplier of 3, for example. If zero, then holders of this instrument don't
	// get any votes for their tokens.
		vote_multiplier;
	
		// 	Set to true if the administration is permitted to make proposals
	// outside of the smart contract scope.
		administration_proposal;
	
		// 	Set to true if a holder is permitted to make proposals outside of the
	// smart contract scope.
		holder_proposal;
	
		// 	Supported values: 1 - Contract-wide Instrument Governance. 0 - Instrument-wide
	// Instrument Governance. If a referendum or initiative is used to propose a
	// modification to a subfield controlled by the instrument auth flags, then the
	// vote will either be a contract-wide vote (all instruments vote on the
	// referendum/initiative) or an instrument-wide vote (only this instrument votes on
	// the referendum/initiative) depending on the value in this subfield. The
	// voting system specifies the voting rules.
		instrument_modification_governance;
	
		// 	The number of tokens to issue with this instrument. Set to greater than
	// zero for fungible tokens. If the value is 1 then it becomes a
	// non-fungible token, where the contract would have many instruments. Set to 0
	// to represent a placeholder instrument, where tokens are to be issued later,
	// or to represent a decomissioned instrument where all tokens have been
	// revoked.
		token_qty;
	
		// 	A custom payload that contains meta data about this instrument. Payload
	// structure and length is dependent on the instrument type chosen. See instrument
	// documentation for more details.
		instrument_payload;
	
		// 	A counter for the number of times this instrument has been revised using a
	// modification action.
		instrument_revision;
	
		// 	Timestamp in nanoseconds of when the smart contract created the
	// action.
		timestamp;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeInstrumentCreation;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// instrument_type (string)
		{
			WriteFixedChar(buf, this.instrument_type, 3);
		}

		// instrument_code (InstrumentCode)
		{
			buf.write(this.instrument_code.Serialize());
		}

		// instrument_index (uint64)
		{
			write(buf, this.instrument_index, 'uint64');
		}

		// instrument_auth_flags ([]byte)
		{
			WriteVarBin(buf, this.instrument_auth_flags, 8);
		}

		// transfers_permitted (bool)
		{
			write(buf, this.transfers_permitted, 'bool');
		}

		// trade_restrictions ([][3]byte)
		{
			// IsNativeTypeArray [][3]byte
			const type = '[][3]byte'.slice(2);
			WriteVariableSize(buf, this.trade_restrictions.length, 16, 8);
			this.trade_restrictions.forEach(value => {
				write(buf, value, type); // TODO
			});
		}

		// enforcement_orders_permitted (bool)
		{
			write(buf, this.enforcement_orders_permitted, 'bool');
		}

		// voting_rights (bool)
		{
			write(buf, this.voting_rights, 'bool');
		}

		// vote_multiplier (uint8)
		{
			write(buf, this.vote_multiplier, 'uint8');
		}

		// administration_proposal (bool)
		{
			write(buf, this.administration_proposal, 'bool');
		}

		// holder_proposal (bool)
		{
			write(buf, this.holder_proposal, 'bool');
		}

		// instrument_modification_governance (uint8)
		{
			write(buf, this.instrument_modification_governance, 'uint8');
		}

		// token_qty (uint64)
		{
			write(buf, this.token_qty, 'uint64');
		}

		// instrument_payload ([]byte)
		{
			WriteVarBin(buf, this.instrument_payload, 16);
		}

		// instrument_revision (uint32)
		{
			write(buf, this.instrument_revision, 'uint32');
		}

		// timestamp (Timestamp)
		{
			buf.write(this.timestamp.Serialize());
		}

		return buf.buf;
	}

	// write populates the fields in InstrumentCreation from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// instrument_type (string)
		{
			this.instrument_type = ReadFixedChar(buf, 3);
		}

		// instrument_code (InstrumentCode)
		{
			this.instrument_code = new InstrumentCode();
			this.instrument_code.Write(buf);
		}

		// instrument_index (uint64)
		{
			this.instrument_index = read(buf, 'uint64');
		}

		// instrument_auth_flags ([]byte)
		{
			this.instrument_auth_flags = ReadVarBin(buf, 8);
		}

		// transfers_permitted (bool)
		{
			this.transfers_permitted = read(buf, 'bool');
		}

		// trade_restrictions ([][3]byte)
		{
			const size = ReadVariableSize(buf, 16, 8);
			const type = '[][3]byte'.slice(2);
			this.trade_restrictions = [...Array(size)].map(() => read(buf, type));
		}

		// enforcement_orders_permitted (bool)
		{
			this.enforcement_orders_permitted = read(buf, 'bool');
		}

		// voting_rights (bool)
		{
			this.voting_rights = read(buf, 'bool');
		}

		// vote_multiplier (uint8)
		{
			this.vote_multiplier = read(buf, 'uint8');
		}

		// administration_proposal (bool)
		{
			this.administration_proposal = read(buf, 'bool');
		}

		// holder_proposal (bool)
		{
			this.holder_proposal = read(buf, 'bool');
		}

		// instrument_modification_governance (uint8)
		{
			this.instrument_modification_governance = read(buf, 'uint8');
		}

		// token_qty (uint64)
		{
			this.token_qty = read(buf, 'uint64');
		}

		// instrument_payload ([]byte)
		{
			this.instrument_payload = ReadVarBin(buf, 16);
		}

		// instrument_revision (uint32)
		{
			this.instrument_revision = read(buf, 'uint32');
		}

		// timestamp (Timestamp)
		{
			this.timestamp = new Timestamp();
			this.timestamp.Write(buf);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// InstrumentType (string)
		{
			if (this.instrument_type.length > 3) {
				return sprintf('fixedchar field instrument_type too long %d/%d', this.instrument_type.length, 3);
			}
		}

		// InstrumentCode (InstrumentCode)
		{
			// IsInternalType
			const err = this.instrument_code.Validate();
			if  (err) return sprintf('field instrument_code is invalid : %s', err);

		}

		// InstrumentIndex (uint64)
		{
		}

		// InstrumentAuthFlags ([]byte)
		{
			if (this.instrument_auth_flags.length >= (2 ** 8)) {
				return sprintf('varbin field instrument_auth_flags too long %d/%d', this.instrument_auth_flags.length, (2 ** 8) - 1);
			}
		}

		// TransfersPermitted (bool) 
		{
		}

		// TradeRestrictions ([][3]byte) 
		{
			if (this.trade_restrictions.length >= (2 ** 16)) {
				return sprintf('list field trade_restrictions has too many items %d/%d', this.trade_restrictions.length, (2 ** 16) - 1);
			}

			let err = null;
			this.trade_restrictions.find((value) => {
				const str = Buffer.from(value).toString('ascii');
				if (!Resources.GetPolityType(str)) {
					err = sprintf('trade_restrictions: Invalid polity value : %s', str);
					return true;
				}
			});
			if (err) return err;
		}

		// EnforcementOrdersPermitted (bool) 
		{
		}

		// VotingRights (bool) 
		{
		}

		// VoteMultiplier (uint8) 
		{
		}

		// AdministrationProposal (bool) 
		{
		}

		// HolderProposal (bool) 
		{
		}

		// InstrumentModificationGovernance (uint8)
		{
			// $field.IntValues [0 1]
			if ( this.instrument_modification_governance !== 0 && this.instrument_modification_governance !== 1) {
				return sprintf('field instrument_modification_governance value is invalid : %d', this.instrument_modification_governance);
			}

		}

		// TokenQty (uint64) 
		{
		}

		// InstrumentPayload ([]byte)
		{
			if (this.instrument_payload.length >= (2 ** 16)) {
				return sprintf('varbin field instrument_payload too long %d/%d', this.instrument_payload.length, (2 ** 16) - 1);
			}
		}

		// InstrumentRevision (uint32)
		{
		}

		// Timestamp (Timestamp) 
		{
			// IsInternalType
			const err = this.timestamp.Validate();
			if  (err) return sprintf('field timestamp is invalid : %s', err);

		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('instrument_type:%s', this.instrument_type.toString()));
			vals.push(sprintf('instrument_code:%s', this.instrument_code.toString()));
			vals.push(sprintf('instrument_index:%d', this.instrument_index));
			vals.push(sprintf('instrument_auth_flags:%s', this.instrument_auth_flags.toString()));
			vals.push(sprintf('transfers_permitted:%s', this.transfers_permitted.toString()));
			vals.push(sprintf('trade_restrictions:%s', this.trade_restrictions.toString()));
			vals.push(sprintf('enforcement_orders_permitted:%s', this.enforcement_orders_permitted.toString()));
			vals.push(sprintf('voting_rights:%s', this.voting_rights.toString()));
			vals.push(sprintf('vote_multiplier:%d', this.vote_multiplier));
			vals.push(sprintf('administration_proposal:%s', this.administration_proposal.toString()));
			vals.push(sprintf('holder_proposal:%s', this.holder_proposal.toString()));
			vals.push(sprintf('instrument_modification_governance:%d', this.instrument_modification_governance));
			vals.push(sprintf('token_qty:%d', this.token_qty));
			vals.push(sprintf('instrument_payload:%s', this.instrument_payload.toString()));
			vals.push(sprintf('instrument_revision:%d', this.instrument_revision));
			vals.push(sprintf('timestamp:%s', this.timestamp.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for InstrumentCreation.', e);
			throw e;
		}
		// this.instrument_type = parsed.instrumentType (fixedchar)
		
		this.instrument_code = new InstrumentCode();
		this.instrument_code.fromJSON(JSON.stringify(parsed.instrumentCode));
		
		this.instrument_index = parsed.instrumentIndex;
		
		// this.instrument_auth_flags = parsed.instrumentAuthFlags (varbin)
		
		// this.transfers_permitted = parsed.transfersPermitted (bool)
		
		// this.trade_restrictions = parsed.tradeRestrictions (Polity[])
		
		// this.enforcement_orders_permitted = parsed.enforcementOrdersPermitted (bool)
		
		// this.voting_rights = parsed.votingRights (bool)
		
		this.vote_multiplier = parsed.voteMultiplier;
		
		// this.administration_proposal = parsed.administrationProposal (bool)
		
		// this.holder_proposal = parsed.holderProposal (bool)
		
		this.instrument_modification_governance = parsed.instrumentModificationGovernance;
		
		this.token_qty = parsed.tokenQty;
		
		// this.instrument_payload = parsed.instrumentPayload (varbin)
		
		this.instrument_revision = parsed.instrumentRevision;
		
		this.timestamp = new Timestamp();
		this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
		
		return this.Validate();
	}
}

	// InstrumentModification Token Dilutions, Call Backs/Revocations, burning etc.
	export class InstrumentModification extends OpReturnMessage {
		type = ActionCode.CodeInstrumentModification;
		typeStr = 'InstrumentModification';

	
		// 	Three letter character that specifies the instrument type.
		instrument_type;
	
		// 	A unique code that is used to identify the instrument. It is generated by
	// hashing the contract public key hash and the instrument index.
	// SHA256(contract PKH + instrument index)
		instrument_code;
	
		// 	The current revision figure to ensure the modification provided is
	// based on the latest version.
		instrument_revision;
	
		// 	A collection of modifications to perform on this instrument.
		amendments;
	
		// 	The Bitcoin transaction ID of the associated result action that
	// permitted the modifications. See Governance for more details.
		ref_tx_id;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeInstrumentModification;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// instrument_type (string)
		{
			WriteFixedChar(buf, this.instrument_type, 3);
		}

		// instrument_code (InstrumentCode)
		{
			buf.write(this.instrument_code.Serialize());
		}

		// instrument_revision (uint32)
		{
			write(buf, this.instrument_revision, 'uint32');
		}

		// amendments ([]Amendment)
		{
			WriteVariableSize(buf, this.amendments.length, 8, 8);
			this.amendments.forEach((value) => {
				buf.write(value.Serialize());
			});
		}

		// ref_tx_id (TxId)
		{
			buf.write(this.ref_tx_id.Serialize());
		}

		return buf.buf;
	}

	// write populates the fields in InstrumentModification from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// instrument_type (string)
		{
			this.instrument_type = ReadFixedChar(buf, 3);
		}

		// instrument_code (InstrumentCode)
		{
			this.instrument_code = new InstrumentCode();
			this.instrument_code.Write(buf);
		}

		// instrument_revision (uint32)
		{
			this.instrument_revision = read(buf, 'uint32');
		}

		// amendments ([]Amendment)
		{
			// IsInternalTypeArray
			const size = ReadVariableSize(buf, 8, 8);
			this.amendments = [];
			for (let i = 0; i < size; i++) {
				const newValue = new Amendment();
				newValue.Write(buf);

				this.amendments.push(newValue);
			}
		}

		// ref_tx_id (TxId)
		{
			this.ref_tx_id = new TxId();
			this.ref_tx_id.Write(buf);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// InstrumentType (string)
		{
			if (this.instrument_type.length > 3) {
				return sprintf('fixedchar field instrument_type too long %d/%d', this.instrument_type.length, 3);
			}
		}

		// InstrumentCode (InstrumentCode)
		{
			// IsInternalType
			const err = this.instrument_code.Validate();
			if  (err) return sprintf('field instrument_code is invalid : %s', err);

		}

		// InstrumentRevision (uint32)
		{
		}

		// Amendments ([]Amendment) 
		{
			if (this.amendments.length > (2 ** 8) - 1) {
				return sprintf('list field amendments has too many items %d/%d', this.amendments.length, (2 ** 8) - 1);
			}

			const err = this.amendments.find((value, i) => {
				const err2 = value.Validate();
				if (err2) return sprintf('list field amendments[%d] is invalid : %s', i, err2);
			});
			if (err) return err;
		}

		// RefTxID (TxId) 
		{
			// IsInternalType
			const err = this.ref_tx_id.Validate();
			if  (err) return sprintf('field ref_tx_id is invalid : %s', err);

		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('instrument_type:%s', this.instrument_type.toString()));
			vals.push(sprintf('instrument_code:%s', this.instrument_code.toString()));
			vals.push(sprintf('instrument_revision:%d', this.instrument_revision));
			vals.push(sprintf('amendments:%s', this.amendments.toString()));
			vals.push(sprintf('ref_tx_id:%s', this.ref_tx_id.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for InstrumentModification.', e);
			throw e;
		}
		// this.instrument_type = parsed.instrumentType (fixedchar)
		
		this.instrument_code = new InstrumentCode();
		this.instrument_code.fromJSON(JSON.stringify(parsed.instrumentCode));
		
		this.instrument_revision = parsed.instrumentRevision;
		
		// this.amendments (Amendment[])
		
		this.ref_tx_id = new TxId();
		this.ref_tx_id.fromJSON(JSON.stringify(parsed.refTxID));
		
		return this.Validate();
	}
}

	// ContractOffer Allows the administration to tell the smart contract what
// they want the details (labels, data, T&C's, etc.) of the Contract to be
// on-chain in a public and immutable way. The Contract Offer action
// 'initializes' a generic smart contract that has been spun up by either
// the smart contract operator or the administration. This on-chain action
// allows for the positive response from the smart contract with either a
// Contract Formation Action or a Rejection Action.
	export class ContractOffer extends OpReturnMessage {
		type = ActionCode.CodeContractOffer;
		typeStr = 'ContractOffer';

	
		// 	Can be any unique identifying string, including human readable names
	// for branding/vanity purposes. Contract identifier (instance) is the
	// bitcoin public key hash address. If the public address is lost, then
	// the administration will have to reissue the entire contract, Instrument
	// Definition and tokens with the new public address. Smart contracts can
	// be branded and specialized to suit any terms and conditions.
		contract_name;
	
		// 	1 - SHA-256 Hash, 2 - Tokenized Body of Agreement Format
		body_of_agreement_type;
	
		// 	SHA-256 hash of the body of the agreement (full contract in pdf format
	// or the like) or the full terms and conditions of an agreement in the
	// Tokenized Body of Agreement format. This is specific to the smart
	// contract and relevant Instruments. Legal and technical information.
		body_of_agreement;
	
		// 	Describes the purpose of the contract.
		contract_type;
	
		// 	Supporting documents that are important to the contract.
		supporting_docs;
	
		// 	5 Letter Code to identify which governing law the contract will adhere
	// to. Disputes are to be settled by this law in the jurisdiction
	// specified below. Private dispute resolution organizations can be used
	// as well. A custom code just needs to be defined.
		governing_law;
	
		// 	Legal proceedings/arbitration will take place using the specified
	// Governing Law in this location.
		jurisdiction;
	
		// 	All actions related to the contract will cease to work after this
	// timestamp. The smart contract will stop running. This will allow many
	// token use cases to be able to calculate total smart contract running
	// costs for the entire life of the contract. Eg. an issuer is creating
	// tickets for an event on the 5th of June 2018. The smart contract will
	// facilitate exchange and send transactions up until the 6th of June.
	// Wallets can use this to forget tokens that are no longer valid - or at
	// least store them in an 'Expired' folder.
		contract_expiration;
	
		// 	Points to an information page that also has a copy of the Contract.
	// Anyone can go to the website to have a look at the price/token,
	// information about the issuer (company), information about the instrument,
	// legal information, etc. There will also be a way for token owners to
	// vote on this page and contact details with the issuer/tokenized
	// companies. Could be a IPv6/IPv4, or txn-id for on-chain information or
	// even a public address (DNS).
		contract_uri;
	
		// 	The issuer of this contract.
		issuer;
	
		// 	The URL of the issuer's logo.
		issuer_logo_url;
	
		// 	If true, then the second input is a contract operator. If false, then
	// all additional inputs are just funding and "includes" fields are
	// skipped in serialization.
		contract_operator_included;
	
		// 	An additional entity with operator access to the contract.
		contract_operator;
	
		// 	Satoshis required to be paid to the contract for each instrument action.
		contract_fee;
	
		// 	A list of voting systems.
		voting_systems;
	
		// 	A set of switches that define the authorization rules for this
	// contract. See the Authorization Flags documentation for more detail.
	// Other terms and conditions that are out of the smart contract's control
	// should be listed in the Body of Agreement.
		contract_auth_flags;
	
		// 	Number of Instruments (non-fungible) permitted on this contract. 0 if
	// unlimited which will display an infinity symbol in UI
		restricted_qty_instruments;
	
		// 	Set to true if the administration is permitted to make proposals
	// outside of the smart contract scope.
		administration_proposal;
	
		// 	Set to true if a holder is permitted to make proposals outside of the
	// smart contract scope.
		holder_proposal;
	
		// 	A list of oracles that provide approval for all token transfers for
	// all instruments under the contract.
		oracles;
	
		// 	The public key hash of the contract's master key. This key has the
	// ability to change the active contract address in case of a security
	// failure with the active contract key.
		master_pkh;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeContractOffer;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// contract_name (string)
		{
			WriteVarChar(buf, this.contract_name, 8);
		}

		// body_of_agreement_type (uint8)
		{
			write(buf, this.body_of_agreement_type, 'uint8');
		}

		// body_of_agreement ([]byte)
		{
			WriteVarBin(buf, this.body_of_agreement, 32);
		}

		// contract_type (string)
		{
			WriteVarChar(buf, this.contract_type, 8);
		}

		// supporting_docs ([]Document)
		{
			WriteVariableSize(buf, this.supporting_docs.length, 8, 8);
			this.supporting_docs.forEach((value) => {
				buf.write(value.Serialize());
			});
		}

		// governing_law (string)
		{
			WriteFixedChar(buf, this.governing_law, 5);
		}

		// jurisdiction (string)
		{
			WriteFixedChar(buf, this.jurisdiction, 5);
		}

		// contract_expiration (Timestamp)
		{
			buf.write(this.contract_expiration.Serialize());
		}

		// contract_uri (string)
		{
			WriteVarChar(buf, this.contract_uri, 8);
		}

		// issuer (Entity)
		{
			buf.write(this.issuer.Serialize());
		}

		// issuer_logo_url (string)
		{
			WriteVarChar(buf, this.issuer_logo_url, 8);
		}

		// contract_operator_included (bool)
		{
			write(buf, this.contract_operator_included, 'bool');
		}

		// contract_operator (Entity)
		if (this.contract_operator_included) {
			buf.write(this.contract_operator.Serialize());
		}

		// contract_fee (uint64)
		{
			write(buf, this.contract_fee, 'uint64');
		}

		// voting_systems ([]VotingSystem)
		{
			WriteVariableSize(buf, this.voting_systems.length, 8, 8);
			this.voting_systems.forEach((value) => {
				buf.write(value.Serialize());
			});
		}

		// contract_auth_flags ([]byte)
		{
			WriteVarBin(buf, this.contract_auth_flags, 16);
		}

		// restricted_qty_instruments (uint64)
		{
			write(buf, this.restricted_qty_instruments, 'uint64');
		}

		// administration_proposal (bool)
		{
			write(buf, this.administration_proposal, 'bool');
		}

		// holder_proposal (bool)
		{
			write(buf, this.holder_proposal, 'bool');
		}

		// oracles ([]Oracle)
		{
			WriteVariableSize(buf, this.oracles.length, 8, 8);
			this.oracles.forEach((value) => {
				buf.write(value.Serialize());
			});
		}

		// master_pkh (PublicKeyHash)
		{
			buf.write(this.master_pkh.Serialize());
		}

		return buf.buf;
	}

	// write populates the fields in ContractOffer from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// contract_name (string)
		{
			this.contract_name = ReadVarChar(buf, 8);
		}

		// body_of_agreement_type (uint8)
		{
			this.body_of_agreement_type = read(buf, 'uint8');
		}

		// body_of_agreement ([]byte)
		{
			this.body_of_agreement = ReadVarBin(buf, 32);
		}

		// contract_type (string)
		{
			this.contract_type = ReadVarChar(buf, 8);
		}

		// supporting_docs ([]Document)
		{
			// IsInternalTypeArray
			const size = ReadVariableSize(buf, 8, 8);
			this.supporting_docs = [];
			for (let i = 0; i < size; i++) {
				const newValue = new Document();
				newValue.Write(buf);

				this.supporting_docs.push(newValue);
			}
		}

		// governing_law (string)
		{
			this.governing_law = ReadFixedChar(buf, 5);
		}

		// jurisdiction (string)
		{
			this.jurisdiction = ReadFixedChar(buf, 5);
		}

		// contract_expiration (Timestamp)
		{
			this.contract_expiration = new Timestamp();
			this.contract_expiration.Write(buf);
		}

		// contract_uri (string)
		{
			this.contract_uri = ReadVarChar(buf, 8);
		}

		// issuer (Entity)
		{
			this.issuer = new Entity();
			this.issuer.Write(buf);
		}

		// issuer_logo_url (string)
		{
			this.issuer_logo_url = ReadVarChar(buf, 8);
		}

		// contract_operator_included (bool)
		{
			this.contract_operator_included = read(buf, 'bool');
		}

		// contract_operator (Entity)
		if (this.contract_operator_included) {
			this.contract_operator = new Entity();
			this.contract_operator.Write(buf);
		}

		// contract_fee (uint64)
		{
			this.contract_fee = read(buf, 'uint64');
		}

		// voting_systems ([]VotingSystem)
		{
			// IsInternalTypeArray
			const size = ReadVariableSize(buf, 8, 8);
			this.voting_systems = [];
			for (let i = 0; i < size; i++) {
				const newValue = new VotingSystem();
				newValue.Write(buf);

				this.voting_systems.push(newValue);
			}
		}

		// contract_auth_flags ([]byte)
		{
			this.contract_auth_flags = ReadVarBin(buf, 16);
		}

		// restricted_qty_instruments (uint64)
		{
			this.restricted_qty_instruments = read(buf, 'uint64');
		}

		// administration_proposal (bool)
		{
			this.administration_proposal = read(buf, 'bool');
		}

		// holder_proposal (bool)
		{
			this.holder_proposal = read(buf, 'bool');
		}

		// oracles ([]Oracle)
		{
			// IsInternalTypeArray
			const size = ReadVariableSize(buf, 8, 8);
			this.oracles = [];
			for (let i = 0; i < size; i++) {
				const newValue = new Oracle();
				newValue.Write(buf);

				this.oracles.push(newValue);
			}
		}

		// master_pkh (PublicKeyHash)
		{
			this.master_pkh = new PublicKeyHash();
			this.master_pkh.Write(buf);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// ContractName (string) 
		{
			if (this.contract_name.length > (2 ** 8)) {
				return sprintf('varchar field contract_name too long %d/%d', this.contract_name.length, (2 ** 8) - 1);
			}
		}

		// BodyOfAgreementType (uint8) 
		{
			// $field.IntValues [1 2]
			if ( this.body_of_agreement_type !== 1 && this.body_of_agreement_type !== 2) {
				return sprintf('field body_of_agreement_type value is invalid : %d', this.body_of_agreement_type);
			}

		}

		// BodyOfAgreement ([]byte) 
		{
			if (this.body_of_agreement.length >= (2 ** 32)) {
				return sprintf('varbin field body_of_agreement too long %d/%d', this.body_of_agreement.length, (2 ** 32) - 1);
			}
		}

		// ContractType (string) 
		{
			if (this.contract_type.length > (2 ** 8)) {
				return sprintf('varchar field contract_type too long %d/%d', this.contract_type.length, (2 ** 8) - 1);
			}
		}

		// SupportingDocs ([]Document) 
		{
			if (this.supporting_docs.length > (2 ** 8) - 1) {
				return sprintf('list field supporting_docs has too many items %d/%d', this.supporting_docs.length, (2 ** 8) - 1);
			}

			const err = this.supporting_docs.find((value, i) => {
				const err2 = value.Validate();
				if (err2) return sprintf('list field supporting_docs[%d] is invalid : %s', i, err2);
			});
			if (err) return err;
		}

		// GoverningLaw (string) 
		{
			if (this.governing_law.length > 5) {
				return sprintf('fixedchar field governing_law too long %d/%d', this.governing_law.length, 5);
			}
		}

		// Jurisdiction (string) 
		{
			if (this.jurisdiction.length > 5) {
				return sprintf('fixedchar field jurisdiction too long %d/%d', this.jurisdiction.length, 5);
			}
		}

		// ContractExpiration (Timestamp) 
		{
			// IsInternalType
			const err = this.contract_expiration.Validate();
			if  (err) return sprintf('field contract_expiration is invalid : %s', err);

		}

		// ContractURI (string) 
		{
			if (this.contract_uri.length > (2 ** 8)) {
				return sprintf('varchar field contract_uri too long %d/%d', this.contract_uri.length, (2 ** 8) - 1);
			}
		}

		// Issuer (Entity) 
		{
			// IsInternalType
			const err = this.issuer.Validate();
			if  (err) return sprintf('field issuer is invalid : %s', err);

		}

		// IssuerLogoURL (string) 
		{
			if (this.issuer_logo_url.length > (2 ** 8)) {
				return sprintf('varchar field issuer_logo_url too long %d/%d', this.issuer_logo_url.length, (2 ** 8) - 1);
			}
		}

		// ContractOperatorIncluded (bool) 
		{
		}

		// ContractOperator (Entity)
		// IncludeIfTrue
		if (this.contract_operator_included) {
			// IsInternalType
			const err = this.contract_operator.Validate();
			if  (err) return sprintf('field contract_operator is invalid : %s', err);

		}

		// ContractFee (uint64) 
		{
		}

		// VotingSystems ([]VotingSystem) 
		{
			if (this.voting_systems.length > (2 ** 8) - 1) {
				return sprintf('list field voting_systems has too many items %d/%d', this.voting_systems.length, (2 ** 8) - 1);
			}

			const err = this.voting_systems.find((value, i) => {
				const err2 = value.Validate();
				if (err2) return sprintf('list field voting_systems[%d] is invalid : %s', i, err2);
			});
			if (err) return err;
		}

		// ContractAuthFlags ([]byte) 
		{
			if (this.contract_auth_flags.length >= (2 ** 16)) {
				return sprintf('varbin field contract_auth_flags too long %d/%d', this.contract_auth_flags.length, (2 ** 16) - 1);
			}
		}

		// RestrictedQtyInstruments (uint64)
		{
		}

		// AdministrationProposal (bool) 
		{
		}

		// HolderProposal (bool) 
		{
		}

		// Oracles ([]Oracle) 
		{
			if (this.oracles.length > (2 ** 8) - 1) {
				return sprintf('list field oracles has too many items %d/%d', this.oracles.length, (2 ** 8) - 1);
			}

			const err = this.oracles.find((value, i) => {
				const err2 = value.Validate();
				if (err2) return sprintf('list field oracles[%d] is invalid : %s', i, err2);
			});
			if (err) return err;
		}

		// MasterPKH (PublicKeyHash) 
		{
			// IsInternalType
			const err = this.master_pkh.Validate();
			if  (err) return sprintf('field master_pkh is invalid : %s', err);

		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('contract_name:%s', this.contract_name.toString()));
			vals.push(sprintf('body_of_agreement_type:%d', this.body_of_agreement_type));
			vals.push(sprintf('body_of_agreement:%s', this.body_of_agreement.toString()));
			vals.push(sprintf('contract_type:%s', this.contract_type.toString()));
			vals.push(sprintf('supporting_docs:%s', this.supporting_docs.toString()));
			vals.push(sprintf('governing_law:%s', this.governing_law.toString()));
			vals.push(sprintf('jurisdiction:%s', this.jurisdiction.toString()));
			vals.push(sprintf('contract_expiration:%s', this.contract_expiration.toString()));
			vals.push(sprintf('contract_uri:%s', this.contract_uri.toString()));
			vals.push(sprintf('issuer:%s', this.issuer.toString()));
			vals.push(sprintf('issuer_logo_url:%s', this.issuer_logo_url.toString()));
			vals.push(sprintf('contract_operator_included:%s', this.contract_operator_included.toString()));
			vals.push(sprintf('contract_operator:%s', this.contract_operator.toString()));
			vals.push(sprintf('contract_fee:%d', this.contract_fee));
			vals.push(sprintf('voting_systems:%s', this.voting_systems.toString()));
			vals.push(sprintf('contract_auth_flags:%s', this.contract_auth_flags.toString()));
			vals.push(sprintf('restricted_qty_instruments:%d', this.restricted_qty_instruments));
			vals.push(sprintf('administration_proposal:%s', this.administration_proposal.toString()));
			vals.push(sprintf('holder_proposal:%s', this.holder_proposal.toString()));
			vals.push(sprintf('oracles:%s', this.oracles.toString()));
			vals.push(sprintf('master_pkh:%s', this.master_pkh.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for ContractOffer.', e);
			throw e;
		}
		this.contract_name = parsed.contractName;
		
		this.body_of_agreement_type = parsed.bodyOfAgreementType;
		
		// this.body_of_agreement = parsed.bodyOfAgreement (varbin)
		
		this.contract_type = parsed.contractType;
		
		// this.supporting_docs (Document[])
		
		// this.governing_law = parsed.governingLaw (fixedchar)
		
		// this.jurisdiction = parsed.jurisdiction (fixedchar)
		
		this.contract_expiration = new Timestamp();
		this.contract_expiration.fromJSON(JSON.stringify(parsed.contractExpiration));
		
		this.contract_uri = parsed.contractURI;
		
		this.issuer = new Entity();
		this.issuer.fromJSON(JSON.stringify(parsed.issuer));
		
		this.issuer_logo_url = parsed.issuerLogoURL;
		
		// this.contract_operator_included = parsed.contractOperatorIncluded (bool)
		
		this.contract_operator = new Entity();
		this.contract_operator.fromJSON(JSON.stringify(parsed.contractOperator));
		
		this.contract_fee = parsed.contractFee;
		
		// this.voting_systems (VotingSystem[])
		
		// this.contract_auth_flags = parsed.contractAuthFlags (varbin)
		
		this.restricted_qty_instruments = parsed.restrictedQtyInstruments;
		
		// this.administration_proposal = parsed.administrationProposal (bool)
		
		// this.holder_proposal = parsed.holderProposal (bool)
		
		// this.oracles (Oracle[])
		
		this.master_pkh = new PublicKeyHash();
		this.master_pkh.fromJSON(JSON.stringify(parsed.masterPKH));
		
		return this.Validate();
	}
}

	// ContractFormation This txn is created by the contract (smart
// contract/off-chain agent/token contract) upon receipt of a valid
// Contract Offer Action from the administration. The smart contract will
// execute on a server controlled by the administration, or a smart
// contract operator on their behalf.
	export class ContractFormation extends OpReturnMessage {
		type = ActionCode.CodeContractFormation;
		typeStr = 'ContractFormation';

	
		// 	Can be any unique identifying string, including human readable names
	// for branding/vanity purposes. Contract identifier (instance) is the
	// bitcoin public key hash address. If the public address is lost, then
	// the administration will have to reissue the entire contract, instrument
	// definition and tokens with the new public address. Smart contracts can
	// be branded and specialized to suit any terms and conditions.
		contract_name;
	
		// 	1 - SHA-256 Hash, 2 - Tokenized Body of Agreement Format
		body_of_agreement_type;
	
		// 	SHA-256 hash of the body of the agreement (full contract in pdf format
	// or the like) or the full terms and conditions of an agreement in the
	// Tokenized Body of Agreement format. This is specific to the smart
	// contract and relevant Instruments. Legal and technical information.
		body_of_agreement;
	
		// 	Describes the purpose of the contract.
		contract_type;
	
		// 	Supporting documents that are important to the contract.
		supporting_docs;
	
		// 	5 Letter Code to identify which governing law the contract will adhere
	// to. Disputes are to be settled by this law in the jurisdiction
	// specified below. Private dispute resolution organizations can be used
	// as well. A custom code just needs to be defined.
		governing_law;
	
		// 	Legal proceedings/arbitration will take place using the specified
	// Governing Law in this location.
		jurisdiction;
	
		// 	All actions related to the contract will cease to work after this
	// timestamp. The smart contract will stop running. This will allow many
	// token use cases to be able to calculate smart contract running costs.
	// Eg. an issuer is creating tickets for an event on the 5th of June 2018.
	// The smart contract will facilitate exchange and send transactions up
	// until the 6th of June. Wallets can use this to forget tokens that are
	// no longer valid - or at least store them in an 'Expired' folder.
		contract_expiration;
	
		// 	Length 0-255 bytes. 0 is valid. Points to an information page that
	// also has a copy of the Contract. Anyone can go to the website to have a
	// look at the price/token, information about the Issuer (company),
	// information about the Instrument, legal information, etc. There will also be
	// a way for Token Owners to vote on this page and contact details with
	// the Issuer/tokenized companies. Could be a IPv6/IPv4, an IPFS address
	// (hash) or txn-id for on chain information or even a public address
	// (DNS).
		contract_uri;
	
		// 	The issuer of this contract.
		issuer;
	
		// 	The URL of the issuer's logo.
		issuer_logo_url;
	
		// 	If true, then the second input is a contract operator. If false, then
	// all additional inputs are just funding and "includes" fields are
	// skipped in serialization.
		contract_operator_included;
	
		// 	An additional entity with operator access to the contract.
		contract_operator;
	
		// 	Satoshis required to be paid to the contract for each instrument action.
		contract_fee;
	
		// 	A list voting systems.
		voting_systems;
	
		// 	A set of switches that define the authorization rules for this
	// contract. See the Authorization Flags documentation for more detail.
	// Other terms and conditions that are out of the smart contract's control
	// should be listed in the Body of Agreement
		contract_auth_flags;
	
		// 	Number of Instruments (non-fungible) permitted on this contract. 0 if
	// unlimited which will display an infinity symbol in UI
		restricted_qty_instruments;
	
		// 	Set to true if the administration is permitted to make proposals
	// outside of the smart contract scope.
		administration_proposal;
	
		// 	Set to true if a holder is permitted to make proposals outside of the
	// smart contract scope.
		holder_proposal;
	
		// 	A list of oracles that provide approval for all token transfers for
	// all instruments under the contract.
		oracles;
	
		// 	The public key hash of the contract's master key. This key has the
	// ability to change the active contract address in case of a security
	// failure with the active contract key.
		master_pkh;
	
		// 	A counter for the number of times this contract has been revised using
	// an amendment action.
		contract_revision;
	
		// 	Timestamp in nanoseconds of when the smart contract created the
	// action.
		timestamp;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeContractFormation;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// contract_name (string)
		{
			WriteVarChar(buf, this.contract_name, 8);
		}

		// body_of_agreement_type (uint8)
		{
			write(buf, this.body_of_agreement_type, 'uint8');
		}

		// body_of_agreement ([]byte)
		{
			WriteVarBin(buf, this.body_of_agreement, 32);
		}

		// contract_type (string)
		{
			WriteVarChar(buf, this.contract_type, 8);
		}

		// supporting_docs ([]Document)
		{
			WriteVariableSize(buf, this.supporting_docs.length, 8, 8);
			this.supporting_docs.forEach((value) => {
				buf.write(value.Serialize());
			});
		}

		// governing_law (string)
		{
			WriteFixedChar(buf, this.governing_law, 5);
		}

		// jurisdiction (string)
		{
			WriteFixedChar(buf, this.jurisdiction, 5);
		}

		// contract_expiration (Timestamp)
		{
			buf.write(this.contract_expiration.Serialize());
		}

		// contract_uri (string)
		{
			WriteVarChar(buf, this.contract_uri, 8);
		}

		// issuer (Entity)
		{
			buf.write(this.issuer.Serialize());
		}

		// issuer_logo_url (string)
		{
			WriteVarChar(buf, this.issuer_logo_url, 8);
		}

		// contract_operator_included (bool)
		{
			write(buf, this.contract_operator_included, 'bool');
		}

		// contract_operator (Entity)
		if (this.contract_operator_included) {
			buf.write(this.contract_operator.Serialize());
		}

		// contract_fee (uint64)
		{
			write(buf, this.contract_fee, 'uint64');
		}

		// voting_systems ([]VotingSystem)
		{
			WriteVariableSize(buf, this.voting_systems.length, 8, 8);
			this.voting_systems.forEach((value) => {
				buf.write(value.Serialize());
			});
		}

		// contract_auth_flags ([]byte)
		{
			WriteVarBin(buf, this.contract_auth_flags, 16);
		}

		// restricted_qty_instruments (uint64)
		{
			write(buf, this.restricted_qty_instruments, 'uint64');
		}

		// administration_proposal (bool)
		{
			write(buf, this.administration_proposal, 'bool');
		}

		// holder_proposal (bool)
		{
			write(buf, this.holder_proposal, 'bool');
		}

		// oracles ([]Oracle)
		{
			WriteVariableSize(buf, this.oracles.length, 8, 8);
			this.oracles.forEach((value) => {
				buf.write(value.Serialize());
			});
		}

		// master_pkh (PublicKeyHash)
		{
			buf.write(this.master_pkh.Serialize());
		}

		// contract_revision (uint32)
		{
			write(buf, this.contract_revision, 'uint32');
		}

		// timestamp (Timestamp)
		{
			buf.write(this.timestamp.Serialize());
		}

		return buf.buf;
	}

	// write populates the fields in ContractFormation from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// contract_name (string)
		{
			this.contract_name = ReadVarChar(buf, 8);
		}

		// body_of_agreement_type (uint8)
		{
			this.body_of_agreement_type = read(buf, 'uint8');
		}

		// body_of_agreement ([]byte)
		{
			this.body_of_agreement = ReadVarBin(buf, 32);
		}

		// contract_type (string)
		{
			this.contract_type = ReadVarChar(buf, 8);
		}

		// supporting_docs ([]Document)
		{
			// IsInternalTypeArray
			const size = ReadVariableSize(buf, 8, 8);
			this.supporting_docs = [];
			for (let i = 0; i < size; i++) {
				const newValue = new Document();
				newValue.Write(buf);

				this.supporting_docs.push(newValue);
			}
		}

		// governing_law (string)
		{
			this.governing_law = ReadFixedChar(buf, 5);
		}

		// jurisdiction (string)
		{
			this.jurisdiction = ReadFixedChar(buf, 5);
		}

		// contract_expiration (Timestamp)
		{
			this.contract_expiration = new Timestamp();
			this.contract_expiration.Write(buf);
		}

		// contract_uri (string)
		{
			this.contract_uri = ReadVarChar(buf, 8);
		}

		// issuer (Entity)
		{
			this.issuer = new Entity();
			this.issuer.Write(buf);
		}

		// issuer_logo_url (string)
		{
			this.issuer_logo_url = ReadVarChar(buf, 8);
		}

		// contract_operator_included (bool)
		{
			this.contract_operator_included = read(buf, 'bool');
		}

		// contract_operator (Entity)
		if (this.contract_operator_included) {
			this.contract_operator = new Entity();
			this.contract_operator.Write(buf);
		}

		// contract_fee (uint64)
		{
			this.contract_fee = read(buf, 'uint64');
		}

		// voting_systems ([]VotingSystem)
		{
			// IsInternalTypeArray
			const size = ReadVariableSize(buf, 8, 8);
			this.voting_systems = [];
			for (let i = 0; i < size; i++) {
				const newValue = new VotingSystem();
				newValue.Write(buf);

				this.voting_systems.push(newValue);
			}
		}

		// contract_auth_flags ([]byte)
		{
			this.contract_auth_flags = ReadVarBin(buf, 16);
		}

		// restricted_qty_instruments (uint64)
		{
			this.restricted_qty_instruments = read(buf, 'uint64');
		}

		// administration_proposal (bool)
		{
			this.administration_proposal = read(buf, 'bool');
		}

		// holder_proposal (bool)
		{
			this.holder_proposal = read(buf, 'bool');
		}

		// oracles ([]Oracle)
		{
			// IsInternalTypeArray
			const size = ReadVariableSize(buf, 8, 8);
			this.oracles = [];
			for (let i = 0; i < size; i++) {
				const newValue = new Oracle();
				newValue.Write(buf);

				this.oracles.push(newValue);
			}
		}

		// master_pkh (PublicKeyHash)
		{
			this.master_pkh = new PublicKeyHash();
			this.master_pkh.Write(buf);
		}

		// contract_revision (uint32)
		{
			this.contract_revision = read(buf, 'uint32');
		}

		// timestamp (Timestamp)
		{
			this.timestamp = new Timestamp();
			this.timestamp.Write(buf);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// ContractName (string) 
		{
			if (this.contract_name.length > (2 ** 8)) {
				return sprintf('varchar field contract_name too long %d/%d', this.contract_name.length, (2 ** 8) - 1);
			}
		}

		// BodyOfAgreementType (uint8) 
		{
			// $field.IntValues [1 2]
			if ( this.body_of_agreement_type !== 1 && this.body_of_agreement_type !== 2) {
				return sprintf('field body_of_agreement_type value is invalid : %d', this.body_of_agreement_type);
			}

		}

		// BodyOfAgreement ([]byte) 
		{
			if (this.body_of_agreement.length >= (2 ** 32)) {
				return sprintf('varbin field body_of_agreement too long %d/%d', this.body_of_agreement.length, (2 ** 32) - 1);
			}
		}

		// ContractType (string) 
		{
			if (this.contract_type.length > (2 ** 8)) {
				return sprintf('varchar field contract_type too long %d/%d', this.contract_type.length, (2 ** 8) - 1);
			}
		}

		// SupportingDocs ([]Document) 
		{
			if (this.supporting_docs.length > (2 ** 8) - 1) {
				return sprintf('list field supporting_docs has too many items %d/%d', this.supporting_docs.length, (2 ** 8) - 1);
			}

			const err = this.supporting_docs.find((value, i) => {
				const err2 = value.Validate();
				if (err2) return sprintf('list field supporting_docs[%d] is invalid : %s', i, err2);
			});
			if (err) return err;
		}

		// GoverningLaw (string) 
		{
			if (this.governing_law.length > 5) {
				return sprintf('fixedchar field governing_law too long %d/%d', this.governing_law.length, 5);
			}
		}

		// Jurisdiction (string) 
		{
			if (this.jurisdiction.length > 5) {
				return sprintf('fixedchar field jurisdiction too long %d/%d', this.jurisdiction.length, 5);
			}
		}

		// ContractExpiration (Timestamp) 
		{
			// IsInternalType
			const err = this.contract_expiration.Validate();
			if  (err) return sprintf('field contract_expiration is invalid : %s', err);

		}

		// ContractURI (string) 
		{
			if (this.contract_uri.length > (2 ** 8)) {
				return sprintf('varchar field contract_uri too long %d/%d', this.contract_uri.length, (2 ** 8) - 1);
			}
		}

		// Issuer (Entity) 
		{
			// IsInternalType
			const err = this.issuer.Validate();
			if  (err) return sprintf('field issuer is invalid : %s', err);

		}

		// IssuerLogoURL (string) 
		{
			if (this.issuer_logo_url.length > (2 ** 8)) {
				return sprintf('varchar field issuer_logo_url too long %d/%d', this.issuer_logo_url.length, (2 ** 8) - 1);
			}
		}

		// ContractOperatorIncluded (bool) 
		{
		}

		// ContractOperator (Entity)
		// IncludeIfTrue
		if (this.contract_operator_included) {
			// IsInternalType
			const err = this.contract_operator.Validate();
			if  (err) return sprintf('field contract_operator is invalid : %s', err);

		}

		// ContractFee (uint64) 
		{
		}

		// VotingSystems ([]VotingSystem) 
		{
			if (this.voting_systems.length > (2 ** 8) - 1) {
				return sprintf('list field voting_systems has too many items %d/%d', this.voting_systems.length, (2 ** 8) - 1);
			}

			const err = this.voting_systems.find((value, i) => {
				const err2 = value.Validate();
				if (err2) return sprintf('list field voting_systems[%d] is invalid : %s', i, err2);
			});
			if (err) return err;
		}

		// ContractAuthFlags ([]byte) 
		{
			if (this.contract_auth_flags.length >= (2 ** 16)) {
				return sprintf('varbin field contract_auth_flags too long %d/%d', this.contract_auth_flags.length, (2 ** 16) - 1);
			}
		}

		// RestrictedQtyInstruments (uint64)
		{
		}

		// AdministrationProposal (bool) 
		{
		}

		// HolderProposal (bool) 
		{
		}

		// Oracles ([]Oracle) 
		{
			if (this.oracles.length > (2 ** 8) - 1) {
				return sprintf('list field oracles has too many items %d/%d', this.oracles.length, (2 ** 8) - 1);
			}

			const err = this.oracles.find((value, i) => {
				const err2 = value.Validate();
				if (err2) return sprintf('list field oracles[%d] is invalid : %s', i, err2);
			});
			if (err) return err;
		}

		// MasterPKH (PublicKeyHash) 
		{
			// IsInternalType
			const err = this.master_pkh.Validate();
			if  (err) return sprintf('field master_pkh is invalid : %s', err);

		}

		// ContractRevision (uint32) 
		{
		}

		// Timestamp (Timestamp) 
		{
			// IsInternalType
			const err = this.timestamp.Validate();
			if  (err) return sprintf('field timestamp is invalid : %s', err);

		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('contract_name:%s', this.contract_name.toString()));
			vals.push(sprintf('body_of_agreement_type:%d', this.body_of_agreement_type));
			vals.push(sprintf('body_of_agreement:%s', this.body_of_agreement.toString()));
			vals.push(sprintf('contract_type:%s', this.contract_type.toString()));
			vals.push(sprintf('supporting_docs:%s', this.supporting_docs.toString()));
			vals.push(sprintf('governing_law:%s', this.governing_law.toString()));
			vals.push(sprintf('jurisdiction:%s', this.jurisdiction.toString()));
			vals.push(sprintf('contract_expiration:%s', this.contract_expiration.toString()));
			vals.push(sprintf('contract_uri:%s', this.contract_uri.toString()));
			vals.push(sprintf('issuer:%s', this.issuer.toString()));
			vals.push(sprintf('issuer_logo_url:%s', this.issuer_logo_url.toString()));
			vals.push(sprintf('contract_operator_included:%s', this.contract_operator_included.toString()));
			vals.push(sprintf('contract_operator:%s', this.contract_operator.toString()));
			vals.push(sprintf('contract_fee:%d', this.contract_fee));
			vals.push(sprintf('voting_systems:%s', this.voting_systems.toString()));
			vals.push(sprintf('contract_auth_flags:%s', this.contract_auth_flags.toString()));
			vals.push(sprintf('restricted_qty_instruments:%d', this.restricted_qty_instruments));
			vals.push(sprintf('administration_proposal:%s', this.administration_proposal.toString()));
			vals.push(sprintf('holder_proposal:%s', this.holder_proposal.toString()));
			vals.push(sprintf('oracles:%s', this.oracles.toString()));
			vals.push(sprintf('master_pkh:%s', this.master_pkh.toString()));
			vals.push(sprintf('contract_revision:%d', this.contract_revision));
			vals.push(sprintf('timestamp:%s', this.timestamp.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for ContractFormation.', e);
			throw e;
		}
		this.contract_name = parsed.contractName;
		
		this.body_of_agreement_type = parsed.bodyOfAgreementType;
		
		// this.body_of_agreement = parsed.bodyOfAgreement (varbin)
		
		this.contract_type = parsed.contractType;
		
		// this.supporting_docs (Document[])
		
		// this.governing_law = parsed.governingLaw (fixedchar)
		
		// this.jurisdiction = parsed.jurisdiction (fixedchar)
		
		this.contract_expiration = new Timestamp();
		this.contract_expiration.fromJSON(JSON.stringify(parsed.contractExpiration));
		
		this.contract_uri = parsed.contractURI;
		
		this.issuer = new Entity();
		this.issuer.fromJSON(JSON.stringify(parsed.issuer));
		
		this.issuer_logo_url = parsed.issuerLogoURL;
		
		// this.contract_operator_included = parsed.contractOperatorIncluded (bool)
		
		this.contract_operator = new Entity();
		this.contract_operator.fromJSON(JSON.stringify(parsed.contractOperator));
		
		this.contract_fee = parsed.contractFee;
		
		// this.voting_systems (VotingSystem[])
		
		// this.contract_auth_flags = parsed.contractAuthFlags (varbin)
		
		this.restricted_qty_instruments = parsed.restrictedQtyInstruments;
		
		// this.administration_proposal = parsed.administrationProposal (bool)
		
		// this.holder_proposal = parsed.holderProposal (bool)
		
		// this.oracles (Oracle[])
		
		this.master_pkh = new PublicKeyHash();
		this.master_pkh.fromJSON(JSON.stringify(parsed.masterPKH));
		
		this.contract_revision = parsed.contractRevision;
		
		this.timestamp = new Timestamp();
		this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
		
		return this.Validate();
	}
}

	// ContractAmendment The administration can initiate an amendment to the
// contract establishment metadata. The ability to make an amendment to the
// contract is restricted by the Authorization Flag set on the current
// revision of Contract Formation action.
	export class ContractAmendment extends OpReturnMessage {
		type = ActionCode.CodeContractAmendment;
		typeStr = 'ContractAmendment';

	
		// 	Used to change the administration address. The new administration
	// address must be in the input[1] position. A change of the
	// administration or operator address requires both the operator and the
	// administration address to be in the inputs (both signatures) of the
	// Contract Amendment action.
		change_administration_address;
	
		// 	Used to change the smart contract operator address. The new operator
	// address must be in the input[1] position, unless the administration is
	// being changed too, then it is in input[2]. A change of the
	// administration or operator address requires both the operator and the
	// administration address to be in the inputs (both signatures) of the
	// Contract Amendment action.
		change_operator_address;
	
		// 	Counter 0 to (2^32)-1
		contract_revision;
	
		// 	A collection of modifications to perform on this contract.
		amendments;
	
		// 	The Bitcoin transaction ID of the associated result action that
	// permitted the modifications. See Governance for more details.
		ref_tx_id;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeContractAmendment;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// change_administration_address (bool)
		{
			write(buf, this.change_administration_address, 'bool');
		}

		// change_operator_address (bool)
		{
			write(buf, this.change_operator_address, 'bool');
		}

		// contract_revision (uint32)
		{
			write(buf, this.contract_revision, 'uint32');
		}

		// amendments ([]Amendment)
		{
			WriteVariableSize(buf, this.amendments.length, 8, 8);
			this.amendments.forEach((value) => {
				buf.write(value.Serialize());
			});
		}

		// ref_tx_id (TxId)
		{
			buf.write(this.ref_tx_id.Serialize());
		}

		return buf.buf;
	}

	// write populates the fields in ContractAmendment from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// change_administration_address (bool)
		{
			this.change_administration_address = read(buf, 'bool');
		}

		// change_operator_address (bool)
		{
			this.change_operator_address = read(buf, 'bool');
		}

		// contract_revision (uint32)
		{
			this.contract_revision = read(buf, 'uint32');
		}

		// amendments ([]Amendment)
		{
			// IsInternalTypeArray
			const size = ReadVariableSize(buf, 8, 8);
			this.amendments = [];
			for (let i = 0; i < size; i++) {
				const newValue = new Amendment();
				newValue.Write(buf);

				this.amendments.push(newValue);
			}
		}

		// ref_tx_id (TxId)
		{
			this.ref_tx_id = new TxId();
			this.ref_tx_id.Write(buf);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// ChangeAdministrationAddress (bool) 
		{
		}

		// ChangeOperatorAddress (bool) 
		{
		}

		// ContractRevision (uint32) 
		{
		}

		// Amendments ([]Amendment) 
		{
			if (this.amendments.length > (2 ** 8) - 1) {
				return sprintf('list field amendments has too many items %d/%d', this.amendments.length, (2 ** 8) - 1);
			}

			const err = this.amendments.find((value, i) => {
				const err2 = value.Validate();
				if (err2) return sprintf('list field amendments[%d] is invalid : %s', i, err2);
			});
			if (err) return err;
		}

		// RefTxID (TxId) 
		{
			// IsInternalType
			const err = this.ref_tx_id.Validate();
			if  (err) return sprintf('field ref_tx_id is invalid : %s', err);

		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('change_administration_address:%s', this.change_administration_address.toString()));
			vals.push(sprintf('change_operator_address:%s', this.change_operator_address.toString()));
			vals.push(sprintf('contract_revision:%d', this.contract_revision));
			vals.push(sprintf('amendments:%s', this.amendments.toString()));
			vals.push(sprintf('ref_tx_id:%s', this.ref_tx_id.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for ContractAmendment.', e);
			throw e;
		}
		// this.change_administration_address = parsed.changeAdministrationAddress (bool)
		
		// this.change_operator_address = parsed.changeOperatorAddress (bool)
		
		this.contract_revision = parsed.contractRevision;
		
		// this.amendments (Amendment[])
		
		this.ref_tx_id = new TxId();
		this.ref_tx_id.fromJSON(JSON.stringify(parsed.refTxID));
		
		return this.Validate();
	}
}

	// StaticContractFormation Static Contract Formation Action
	export class StaticContractFormation extends OpReturnMessage {
		type = ActionCode.CodeStaticContractFormation;
		typeStr = 'StaticContractFormation';

	
		// 	Can be any unique identifying string, including human readable names
	// for branding/vanity purposes. Contract identifier (instance) is the
	// bitcoin public address. If the public address is lost, then the
	// administration will have to reissue the entire contract, Instrument
	// Definition and tokens with the new public address. Smart contracts can
	// be branded and specialized to suit any terms and conditions.
		contract_name;
	
		// 	32 randomly generated bytes. Each Contract Code should be unique. The
	// Contract ID will be human facing and will be the Contract Code, with a
	// checksum, encoded in base58 and prefixed by 'CON'. Contract ID = CON +
	// base58(ContractCode + checksum). Eg. Contract ID =
	// 'CON18RDoKK7Ed5zid2FkKVy7q3rULr4tgfjr4'
		contract_code;
	
		// 	1 - SHA-256 Hash, 2 - Tokenized Body of Agreement Format
		body_of_agreement_type;
	
		// 	SHA-256 hash of the body of the agreement (full contract in pdf format
	// or the like) or the full terms and conditions of an agreement in the
	// Tokenized Body of Agreement format. This is specific to the smart
	// contract and relevant Instruments. Legal and technical information.
		body_of_agreement;
	
		// 	Describes the purpose of the contract.
		contract_type;
	
		// 	Supporting documents that are important to the contract.
		supporting_docs;
	
		// 	Counter 0 to (2^32)-1
		contract_revision;
	
		// 	5 Letter Code to identify which governing law the contract will adhere
	// to. Disputes are to be settled by this law in the jurisdiction
	// specified below. Private dispute resolution organizations can be used
	// as well. A custom code just needs to be defined.
		governing_law;
	
		// 	Legal proceedings/arbitration will take place using the specified
	// Governing Law in this location.
		jurisdiction;
	
		// 	Start date of the contract.
		effective_date;
	
		// 	All actions related to the contract will cease to work after this
	// timestamp. The smart contract will stop running. This will allow many
	// token use cases to be able to calculate smart contract running costs.
	// Eg. an issuer is creating tickets for an event on the 5th of June 2018.
	// The smart contract will facilitate exchange and send transactions up
	// until the 6th of June. Wallets can use this to forget tokens that are
	// no longer valid - or at least store them in an 'Expired' folder.
		contract_expiration;
	
		// 	Length 0-255 bytes. Points to an information page that also has a copy
	// of the Contract. Anyone can go to the website to have a look at the
	// price/token, information about the issuer (company), information about
	// the Instrument, legal information, etc. There will also be a way for token
	// owners to vote on this page and contact details with the
	// issuer/tokenized companies. Could be a IPv6/IPv4, or txn-id for on
	// chain information or even a public address (DNS).
		contract_uri;
	
		// 	The Tx-ID of the previous contract revision.
		prev_rev_tx_id;
	
		// 	A list of legal entities associated with this contract.
		entities;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeStaticContractFormation;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// contract_name (string)
		{
			WriteVarChar(buf, this.contract_name, 8);
		}

		// contract_code (ContractCode)
		{
			buf.write(this.contract_code.Serialize());
		}

		// body_of_agreement_type (uint8)
		{
			write(buf, this.body_of_agreement_type, 'uint8');
		}

		// body_of_agreement ([]byte)
		{
			WriteVarBin(buf, this.body_of_agreement, 32);
		}

		// contract_type (string)
		{
			WriteVarChar(buf, this.contract_type, 8);
		}

		// supporting_docs ([]Document)
		{
			WriteVariableSize(buf, this.supporting_docs.length, 8, 8);
			this.supporting_docs.forEach((value) => {
				buf.write(value.Serialize());
			});
		}

		// contract_revision (uint32)
		{
			write(buf, this.contract_revision, 'uint32');
		}

		// governing_law (string)
		{
			WriteFixedChar(buf, this.governing_law, 5);
		}

		// jurisdiction (string)
		{
			WriteFixedChar(buf, this.jurisdiction, 5);
		}

		// effective_date (Timestamp)
		{
			buf.write(this.effective_date.Serialize());
		}

		// contract_expiration (Timestamp)
		{
			buf.write(this.contract_expiration.Serialize());
		}

		// contract_uri (string)
		{
			WriteVarChar(buf, this.contract_uri, 8);
		}

		// prev_rev_tx_id (TxId)
		{
			buf.write(this.prev_rev_tx_id.Serialize());
		}

		// entities ([]Entity)
		{
			WriteVariableSize(buf, this.entities.length, 8, 8);
			this.entities.forEach((value) => {
				buf.write(value.Serialize());
			});
		}

		return buf.buf;
	}

	// write populates the fields in StaticContractFormation from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// contract_name (string)
		{
			this.contract_name = ReadVarChar(buf, 8);
		}

		// contract_code (ContractCode)
		{
			this.contract_code = new ContractCode();
			this.contract_code.Write(buf);
		}

		// body_of_agreement_type (uint8)
		{
			this.body_of_agreement_type = read(buf, 'uint8');
		}

		// body_of_agreement ([]byte)
		{
			this.body_of_agreement = ReadVarBin(buf, 32);
		}

		// contract_type (string)
		{
			this.contract_type = ReadVarChar(buf, 8);
		}

		// supporting_docs ([]Document)
		{
			// IsInternalTypeArray
			const size = ReadVariableSize(buf, 8, 8);
			this.supporting_docs = [];
			for (let i = 0; i < size; i++) {
				const newValue = new Document();
				newValue.Write(buf);

				this.supporting_docs.push(newValue);
			}
		}

		// contract_revision (uint32)
		{
			this.contract_revision = read(buf, 'uint32');
		}

		// governing_law (string)
		{
			this.governing_law = ReadFixedChar(buf, 5);
		}

		// jurisdiction (string)
		{
			this.jurisdiction = ReadFixedChar(buf, 5);
		}

		// effective_date (Timestamp)
		{
			this.effective_date = new Timestamp();
			this.effective_date.Write(buf);
		}

		// contract_expiration (Timestamp)
		{
			this.contract_expiration = new Timestamp();
			this.contract_expiration.Write(buf);
		}

		// contract_uri (string)
		{
			this.contract_uri = ReadVarChar(buf, 8);
		}

		// prev_rev_tx_id (TxId)
		{
			this.prev_rev_tx_id = new TxId();
			this.prev_rev_tx_id.Write(buf);
		}

		// entities ([]Entity)
		{
			// IsInternalTypeArray
			const size = ReadVariableSize(buf, 8, 8);
			this.entities = [];
			for (let i = 0; i < size; i++) {
				const newValue = new Entity();
				newValue.Write(buf);

				this.entities.push(newValue);
			}
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// ContractName (string) 
		{
			if (this.contract_name.length > (2 ** 8)) {
				return sprintf('varchar field contract_name too long %d/%d', this.contract_name.length, (2 ** 8) - 1);
			}
		}

		// ContractCode (ContractCode) 
		{
			// IsInternalType
			const err = this.contract_code.Validate();
			if  (err) return sprintf('field contract_code is invalid : %s', err);

		}

		// BodyOfAgreementType (uint8) 
		{
			// $field.IntValues [1 2]
			if ( this.body_of_agreement_type !== 1 && this.body_of_agreement_type !== 2) {
				return sprintf('field body_of_agreement_type value is invalid : %d', this.body_of_agreement_type);
			}

		}

		// BodyOfAgreement ([]byte) 
		{
			if (this.body_of_agreement.length >= (2 ** 32)) {
				return sprintf('varbin field body_of_agreement too long %d/%d', this.body_of_agreement.length, (2 ** 32) - 1);
			}
		}

		// ContractType (string) 
		{
			if (this.contract_type.length > (2 ** 8)) {
				return sprintf('varchar field contract_type too long %d/%d', this.contract_type.length, (2 ** 8) - 1);
			}
		}

		// SupportingDocs ([]Document) 
		{
			if (this.supporting_docs.length > (2 ** 8) - 1) {
				return sprintf('list field supporting_docs has too many items %d/%d', this.supporting_docs.length, (2 ** 8) - 1);
			}

			const err = this.supporting_docs.find((value, i) => {
				const err2 = value.Validate();
				if (err2) return sprintf('list field supporting_docs[%d] is invalid : %s', i, err2);
			});
			if (err) return err;
		}

		// ContractRevision (uint32) 
		{
		}

		// GoverningLaw (string) 
		{
			if (this.governing_law.length > 5) {
				return sprintf('fixedchar field governing_law too long %d/%d', this.governing_law.length, 5);
			}
		}

		// Jurisdiction (string) 
		{
			if (this.jurisdiction.length > 5) {
				return sprintf('fixedchar field jurisdiction too long %d/%d', this.jurisdiction.length, 5);
			}
		}

		// EffectiveDate (Timestamp) 
		{
			// IsInternalType
			const err = this.effective_date.Validate();
			if  (err) return sprintf('field effective_date is invalid : %s', err);

		}

		// ContractExpiration (Timestamp) 
		{
			// IsInternalType
			const err = this.contract_expiration.Validate();
			if  (err) return sprintf('field contract_expiration is invalid : %s', err);

		}

		// ContractURI (string) 
		{
			if (this.contract_uri.length > (2 ** 8)) {
				return sprintf('varchar field contract_uri too long %d/%d', this.contract_uri.length, (2 ** 8) - 1);
			}
		}

		// PrevRevTxID (TxId) 
		{
			// IsInternalType
			const err = this.prev_rev_tx_id.Validate();
			if  (err) return sprintf('field prev_rev_tx_id is invalid : %s', err);

		}

		// Entities ([]Entity) 
		{
			if (this.entities.length > (2 ** 8) - 1) {
				return sprintf('list field entities has too many items %d/%d', this.entities.length, (2 ** 8) - 1);
			}

			const err = this.entities.find((value, i) => {
				const err2 = value.Validate();
				if (err2) return sprintf('list field entities[%d] is invalid : %s', i, err2);
			});
			if (err) return err;
		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('contract_name:%s', this.contract_name.toString()));
			vals.push(sprintf('contract_code:%s', this.contract_code.toString()));
			vals.push(sprintf('body_of_agreement_type:%d', this.body_of_agreement_type));
			vals.push(sprintf('body_of_agreement:%s', this.body_of_agreement.toString()));
			vals.push(sprintf('contract_type:%s', this.contract_type.toString()));
			vals.push(sprintf('supporting_docs:%s', this.supporting_docs.toString()));
			vals.push(sprintf('contract_revision:%d', this.contract_revision));
			vals.push(sprintf('governing_law:%s', this.governing_law.toString()));
			vals.push(sprintf('jurisdiction:%s', this.jurisdiction.toString()));
			vals.push(sprintf('effective_date:%s', this.effective_date.toString()));
			vals.push(sprintf('contract_expiration:%s', this.contract_expiration.toString()));
			vals.push(sprintf('contract_uri:%s', this.contract_uri.toString()));
			vals.push(sprintf('prev_rev_tx_id:%s', this.prev_rev_tx_id.toString()));
			vals.push(sprintf('entities:%s', this.entities.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for StaticContractFormation.', e);
			throw e;
		}
		this.contract_name = parsed.contractName;
		
		this.contract_code = new ContractCode();
		this.contract_code.fromJSON(JSON.stringify(parsed.contractCode));
		
		this.body_of_agreement_type = parsed.bodyOfAgreementType;
		
		// this.body_of_agreement = parsed.bodyOfAgreement (varbin)
		
		this.contract_type = parsed.contractType;
		
		// this.supporting_docs (Document[])
		
		this.contract_revision = parsed.contractRevision;
		
		// this.governing_law = parsed.governingLaw (fixedchar)
		
		// this.jurisdiction = parsed.jurisdiction (fixedchar)
		
		this.effective_date = new Timestamp();
		this.effective_date.fromJSON(JSON.stringify(parsed.effectiveDate));
		
		this.contract_expiration = new Timestamp();
		this.contract_expiration.fromJSON(JSON.stringify(parsed.contractExpiration));
		
		this.contract_uri = parsed.contractURI;
		
		this.prev_rev_tx_id = new TxId();
		this.prev_rev_tx_id.fromJSON(JSON.stringify(parsed.prevRevTxID));
		
		// this.entities (Entity[])
		
		return this.Validate();
	}
}

	// ContractAddressChange This txn is signed by the master contract key
// defined in the contract formation and changes the active contract
// address which the contract uses to receive and respond to requests. This
// is a worst case scenario fallback to only be used when the contract
// private key is believed to be exposed.
	export class ContractAddressChange extends OpReturnMessage {
		type = ActionCode.CodeContractAddressChange;
		typeStr = 'ContractAddressChange';

	
		// 	The address to be used by all future requests/responses for the
	// contract.
		new_contract_pkh;
	
		// 	Timestamp in nanoseconds of when the action was created.
		timestamp;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeContractAddressChange;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// new_contract_pkh (PublicKeyHash)
		{
			buf.write(this.new_contract_pkh.Serialize());
		}

		// timestamp (Timestamp)
		{
			buf.write(this.timestamp.Serialize());
		}

		return buf.buf;
	}

	// write populates the fields in ContractAddressChange from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// new_contract_pkh (PublicKeyHash)
		{
			this.new_contract_pkh = new PublicKeyHash();
			this.new_contract_pkh.Write(buf);
		}

		// timestamp (Timestamp)
		{
			this.timestamp = new Timestamp();
			this.timestamp.Write(buf);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// NewContractPKH (PublicKeyHash) 
		{
			// IsInternalType
			const err = this.new_contract_pkh.Validate();
			if  (err) return sprintf('field new_contract_pkh is invalid : %s', err);

		}

		// Timestamp (Timestamp) 
		{
			// IsInternalType
			const err = this.timestamp.Validate();
			if  (err) return sprintf('field timestamp is invalid : %s', err);

		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('new_contract_pkh:%s', this.new_contract_pkh.toString()));
			vals.push(sprintf('timestamp:%s', this.timestamp.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for ContractAddressChange.', e);
			throw e;
		}
		this.new_contract_pkh = new PublicKeyHash();
		this.new_contract_pkh.fromJSON(JSON.stringify(parsed.newContractPKH));
		
		this.timestamp = new Timestamp();
		this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
		
		return this.Validate();
	}
}

	// Order Used by the administration to signal to the smart contract that
// the tokens that a particular public address(es) owns are to be
// confiscated, frozen, thawed or reconciled.
	export class Order extends OpReturnMessage {
		type = ActionCode.CodeOrder;
		typeStr = 'Order';

	
		// 	Freeze (F), Thaw (T), Confiscate (C), Reconcile (R)
		compliance_action;
	
		// 	Three letter character that specifies the instrument type.
		instrument_type;
	
		// 	A unique code that is used to identify the instrument. It is generated by
	// hashing the contract public key hash and the instrument index.
	// SHA256(contract PKH + instrument index)
		instrument_code;
	
		// 	The holders and quantities that are effected by the order. For a
	// contract or instrument wide freeze only the contract address is specified.
	// Zero quantities are invalid unless it is for the contract address in an
	// instrument wide or contract wide freeze. In a thaw order this field is not
	// serialized, because the entire freeze from the FreezeTxId freeze action
	// will be thawed.
		target_addresses;
	
		// 	The tx id of the freeze action that is being thawed. Only serialized
	// for thaw orders.
		freeze_tx_id;
	
		// 	Used for a 'time out'. Tokens are automatically unfrozen after the
	// expiration timestamp without requiring a Thaw Action. Null value for
	// Thaw, Confiscation and Reconciallitaion orders.
		freeze_period;
	
		// 	The public address for confiscated tokens to be deposited in. Null for
	// Freeze, Thaw, actions.
		deposit_address;
	
		// 	Specifies if an authority signature is included. For Reconcialitaion
	// actions all authority signature related fields are skipped during
	// serialization.
		authority_included;
	
		// 	Length 0-255 bytes. Enforcement Authority Name (eg. Issuer, Queensland
	// Police Service, Tokenized, etc.)
		authority_name;
	
		// 	Length 0-255 bytes. Public Key associated with the Enforcement
	// Authority
		authority_public_key;
	
		// 	Algorithm used for order signature. Only valid value is currently 1 =
	// ECDSA+secp256k1
		signature_algorithm;
	
		// 	Length 0-255 bytes. Signature for a message that lists out the target
	// addresses and deposit address. Signature of (Contract PKH, Compliance
	// Action, Authority Name, Instrument Code, Supporting Evidence Hash,
	// FreezePeriod, TargetAddresses, and DepositAddress)
		order_signature;
	
		// 	SHA-256: warrant, court order, etc.
		supporting_evidence_hash;
	
		// 	The request/response actions that were dropped. The entire txn for
	// both actions is included as evidence that the actions were accepted
	// into the mempool at one point and that the senders (token/Bitcoin)
	// signed their intent to transfer. The management of this record keeping
	// is off-chain and managed by the administration or operator to preserve
	// the integrity of the state of the tokens. Only applicable for
	// reconcilliation actions. No subfield when F, T, R is selected as the
	// Compliance Action subfield.
		ref_txs;
	
		// 	Index of address in TargetAddresses and amount of bitcoin (in
	// satoshis) they are receiving in exchange for their tokens.
		bitcoin_dispersions;
	
		// 	A message to include with the enforcement order.
		message;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeOrder;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// compliance_action (byte)
		{
			write(buf, this.compliance_action, 'byte');
		}

		// instrument_type (string)
		if ( this.compliance_action === char('F') || this.compliance_action === char('C') || this.compliance_action === char('R')) {
			WriteFixedChar(buf, this.instrument_type, 3);
		}

		// instrument_code (InstrumentCode)
		if ( this.compliance_action === char('F') || this.compliance_action === char('C') || this.compliance_action === char('R')) {
			buf.write(this.instrument_code.Serialize());
		}

		// target_addresses ([]TargetAddress)
		if ( this.compliance_action === char('F') || this.compliance_action === char('C') || this.compliance_action === char('R')) {
			WriteVariableSize(buf, this.target_addresses.length, 16, 8);
			this.target_addresses.forEach((value) => {
				buf.write(value.Serialize());
			});
		}

		// freeze_tx_id (TxId)
		if ( this.compliance_action === char('T')) {
			buf.write(this.freeze_tx_id.Serialize());
		}

		// freeze_period (Timestamp)
		if ( this.compliance_action === char('F')) {
			buf.write(this.freeze_period.Serialize());
		}

		// deposit_address (PublicKeyHash)
		if ( this.compliance_action === char('C')) {
			buf.write(this.deposit_address.Serialize());
		}

		// authority_included (bool)
		if ( this.compliance_action === char('F') || this.compliance_action === char('T') || this.compliance_action === char('C')) {
			write(buf, this.authority_included, 'bool');
		}

		// authority_name (string)
		if (this.authority_included) {
			WriteVarChar(buf, this.authority_name, 8);
		}

		// authority_public_key ([]byte)
		if (this.authority_included) {
			WriteVarBin(buf, this.authority_public_key, 8);
		}

		// signature_algorithm (uint8)
		if (this.authority_included) {
			write(buf, this.signature_algorithm, 'uint8');
		}

		// order_signature ([]byte)
		if (this.authority_included) {
			WriteVarBin(buf, this.order_signature, 8);
		}

		// supporting_evidence_hash ([32]byte)
		{
			write(buf, this.supporting_evidence_hash, '[32]byte');
		}

		// ref_txs ([]byte)
		if ( this.compliance_action === char('R')) {
			WriteVarBin(buf, this.ref_txs, 32);
		}

		// bitcoin_dispersions ([]QuantityIndex)
		if ( this.compliance_action === char('R')) {
			WriteVariableSize(buf, this.bitcoin_dispersions.length, 16, 8);
			this.bitcoin_dispersions.forEach((value) => {
				buf.write(value.Serialize());
			});
		}

		// message (string)
		{
			WriteVarChar(buf, this.message, 32);
		}

		return buf.buf;
	}

	// write populates the fields in Order from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// compliance_action (byte)
		{
			this.compliance_action = read(buf, 'byte');
		}

		// instrument_type (string)
		if ( this.compliance_action === char('F') || this.compliance_action === char('C') || this.compliance_action === char('R')) {
			this.instrument_type = ReadFixedChar(buf, 3);
		}

		// instrument_code (InstrumentCode)
		if ( this.compliance_action === char('F') || this.compliance_action === char('C') || this.compliance_action === char('R')) {
			this.instrument_code = new InstrumentCode();
			this.instrument_code.Write(buf);
		}

		// target_addresses ([]TargetAddress)
		if ( this.compliance_action === char('F') || this.compliance_action === char('C') || this.compliance_action === char('R')) {
			// IsInternalTypeArray
			const size = ReadVariableSize(buf, 16, 8);
			this.target_addresses = [];
			for (let i = 0; i < size; i++) {
				const newValue = new TargetAddress();
				newValue.Write(buf);

				this.target_addresses.push(newValue);
			}
		}

		// freeze_tx_id (TxId)
		if ( this.compliance_action === char('T')) {
			this.freeze_tx_id = new TxId();
			this.freeze_tx_id.Write(buf);
		}

		// freeze_period (Timestamp)
		if ( this.compliance_action === char('F')) {
			this.freeze_period = new Timestamp();
			this.freeze_period.Write(buf);
		}

		// deposit_address (PublicKeyHash)
		if ( this.compliance_action === char('C')) {
			this.deposit_address = new PublicKeyHash();
			this.deposit_address.Write(buf);
		}

		// authority_included (bool)
		if ( this.compliance_action === char('F') || this.compliance_action === char('T') || this.compliance_action === char('C')) {
			this.authority_included = read(buf, 'bool');
		}

		// authority_name (string)
		if (this.authority_included) {
			this.authority_name = ReadVarChar(buf, 8);
		}

		// authority_public_key ([]byte)
		if (this.authority_included) {
			this.authority_public_key = ReadVarBin(buf, 8);
		}

		// signature_algorithm (uint8)
		if (this.authority_included) {
			this.signature_algorithm = read(buf, 'uint8');
		}

		// order_signature ([]byte)
		if (this.authority_included) {
			this.order_signature = ReadVarBin(buf, 8);
		}

		// supporting_evidence_hash ([32]byte)
		{
			this.supporting_evidence_hash = read(buf, '[32]byte');
		}

		// ref_txs ([]byte)
		if ( this.compliance_action === char('R')) {
			this.ref_txs = ReadVarBin(buf, 32);
		}

		// bitcoin_dispersions ([]QuantityIndex)
		if ( this.compliance_action === char('R')) {
			// IsInternalTypeArray
			const size = ReadVariableSize(buf, 16, 8);
			this.bitcoin_dispersions = [];
			for (let i = 0; i < size; i++) {
				const newValue = new QuantityIndex();
				newValue.Write(buf);

				this.bitcoin_dispersions.push(newValue);
			}
		}

		// message (string)
		{
			this.message = ReadVarChar(buf, 32);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// ComplianceAction (byte) 
		{
			if ( this.compliance_action !== 'F'
			&& this.compliance_action !== 'T'
			&& this.compliance_action !== 'C'
			&& this.compliance_action !== 'R'
			) {
				return sprintf('field compliance_action value is invalid : %d', this.compliance_action);
			}
		}

		// InstrumentType (string)
		// IncludeIf.Field
		if ( this.compliance_action === char('F') || this.compliance_action === char('C') || this.compliance_action === char('R')) {
			if (this.instrument_type.length > 3) {
				return sprintf('fixedchar field instrument_type too long %d/%d', this.instrument_type.length, 3);
			}
		}

		// InstrumentCode (InstrumentCode)
		// IncludeIf.Field
		if ( this.compliance_action === char('F') || this.compliance_action === char('C') || this.compliance_action === char('R')) {
			// IsInternalType
			const err = this.instrument_code.Validate();
			if  (err) return sprintf('field instrument_code is invalid : %s', err);

		}

		// TargetAddresses ([]TargetAddress)
		// IncludeIf.Field
		if ( this.compliance_action === char('F') || this.compliance_action === char('C') || this.compliance_action === char('R')) {
			if (this.target_addresses.length > (2 ** 16) - 1) {
				return sprintf('list field target_addresses has too many items %d/%d', this.target_addresses.length, (2 ** 16) - 1);
			}

			const err = this.target_addresses.find((value, i) => {
				const err2 = value.Validate();
				if (err2) return sprintf('list field target_addresses[%d] is invalid : %s', i, err2);
			});
			if (err) return err;
		}

		// FreezeTxId (TxId)
		// IncludeIf.Field
		if ( this.compliance_action === char('T')) {
			// IsInternalType
			const err = this.freeze_tx_id.Validate();
			if  (err) return sprintf('field freeze_tx_id is invalid : %s', err);

		}

		// FreezePeriod (Timestamp)
		// IncludeIf.Field
		if ( this.compliance_action === char('F')) {
			// IsInternalType
			const err = this.freeze_period.Validate();
			if  (err) return sprintf('field freeze_period is invalid : %s', err);

		}

		// DepositAddress (PublicKeyHash)
		// IncludeIf.Field
		if ( this.compliance_action === char('C')) {
			// IsInternalType
			const err = this.deposit_address.Validate();
			if  (err) return sprintf('field deposit_address is invalid : %s', err);

		}

		// AuthorityIncluded (bool)
		// IncludeIf.Field
		if ( this.compliance_action === char('F') || this.compliance_action === char('T') || this.compliance_action === char('C')) {
		}

		// AuthorityName (string)
		// IncludeIfTrue
		if (this.authority_included) {
			if (this.authority_name.length > (2 ** 8)) {
				return sprintf('varchar field authority_name too long %d/%d', this.authority_name.length, (2 ** 8) - 1);
			}
		}

		// AuthorityPublicKey ([]byte)
		// IncludeIfTrue
		if (this.authority_included) {
			if (this.authority_public_key.length >= (2 ** 8)) {
				return sprintf('varbin field authority_public_key too long %d/%d', this.authority_public_key.length, (2 ** 8) - 1);
			}
		}

		// SignatureAlgorithm (uint8)
		// IncludeIfTrue
		if (this.authority_included) {
			// $field.IntValues [1]
			if ( this.signature_algorithm !== 1) {
				return sprintf('field signature_algorithm value is invalid : %d', this.signature_algorithm);
			}

		}

		// OrderSignature ([]byte)
		// IncludeIfTrue
		if (this.authority_included) {
			if (this.order_signature.length >= (2 ** 8)) {
				return sprintf('varbin field order_signature too long %d/%d', this.order_signature.length, (2 ** 8) - 1);
			}
		}

		// SupportingEvidenceHash ([32]byte) 
		{
		}

		// RefTxs ([]byte)
		// IncludeIf.Field
		if ( this.compliance_action === char('R')) {
			if (this.ref_txs.length >= (2 ** 32)) {
				return sprintf('varbin field ref_txs too long %d/%d', this.ref_txs.length, (2 ** 32) - 1);
			}
		}

		// BitcoinDispersions ([]QuantityIndex)
		// IncludeIf.Field
		if ( this.compliance_action === char('R')) {
			if (this.bitcoin_dispersions.length > (2 ** 16) - 1) {
				return sprintf('list field bitcoin_dispersions has too many items %d/%d', this.bitcoin_dispersions.length, (2 ** 16) - 1);
			}

			const err = this.bitcoin_dispersions.find((value, i) => {
				const err2 = value.Validate();
				if (err2) return sprintf('list field bitcoin_dispersions[%d] is invalid : %s', i, err2);
			});
			if (err) return err;
		}

		// Message (string) 
		{
			if (this.message.length > (2 ** 32)) {
				return sprintf('varchar field message too long %d/%d', this.message.length, (2 ** 32) - 1);
			}
		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('compliance_action:%s', this.compliance_action.toString()));
			vals.push(sprintf('instrument_type:%s', this.instrument_type.toString()));
			vals.push(sprintf('instrument_code:%s', this.instrument_code.toString()));
			vals.push(sprintf('target_addresses:%s', this.target_addresses.toString()));
			vals.push(sprintf('freeze_tx_id:%s', this.freeze_tx_id.toString()));
			vals.push(sprintf('freeze_period:%s', this.freeze_period.toString()));
			vals.push(sprintf('deposit_address:%s', this.deposit_address.toString()));
			vals.push(sprintf('authority_included:%s', this.authority_included.toString()));
			vals.push(sprintf('authority_name:%s', this.authority_name.toString()));
			vals.push(sprintf('authority_public_key:%s', this.authority_public_key.toString()));
			vals.push(sprintf('signature_algorithm:%d', this.signature_algorithm));
			vals.push(sprintf('order_signature:%s', this.order_signature.toString()));
			vals.push(sprintf('supporting_evidence_hash:%s', this.supporting_evidence_hash.toString()));
			vals.push(sprintf('ref_txs:%s', this.ref_txs.toString()));
			vals.push(sprintf('bitcoin_dispersions:%s', this.bitcoin_dispersions.toString()));
			vals.push(sprintf('message:%s', this.message.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for Order.', e);
			throw e;
		}
		// this.compliance_action = parsed.complianceAction (fixedchar)
		
		// this.instrument_type = parsed.instrumentType (fixedchar)
		
		this.instrument_code = new InstrumentCode();
		this.instrument_code.fromJSON(JSON.stringify(parsed.instrumentCode));
		
		// this.target_addresses (TargetAddress[])
		
		this.freeze_tx_id = new TxId();
		this.freeze_tx_id.fromJSON(JSON.stringify(parsed.freezeTxId));
		
		this.freeze_period = new Timestamp();
		this.freeze_period.fromJSON(JSON.stringify(parsed.freezePeriod));
		
		this.deposit_address = new PublicKeyHash();
		this.deposit_address.fromJSON(JSON.stringify(parsed.depositAddress));
		
		// this.authority_included = parsed.authorityIncluded (bool)
		
		this.authority_name = parsed.authorityName;
		
		// this.authority_public_key = parsed.authorityPublicKey (varbin)
		
		this.signature_algorithm = parsed.signatureAlgorithm;
		
		// this.order_signature = parsed.orderSignature (varbin)
		
		// this.supporting_evidence_hash = parsed.supportingEvidenceHash (bin)
		
		// this.ref_txs = parsed.refTxs (varbin)
		
		// this.bitcoin_dispersions (QuantityIndex[])
		
		this.message = parsed.message;
		
		return this.Validate();
	}
}

	// Freeze The contract responding to an Order action to freeze instruments. To
// be used to comply with contractual/legal/issuer requirements. The target
// public address(es) will be marked as frozen. However the Freeze action
// publishes this fact to the public blockchain for transparency. The
// contract will not respond to any actions requested by the frozen
// address.
	export class Freeze extends OpReturnMessage {
		type = ActionCode.CodeFreeze;
		typeStr = 'Freeze';

	
		// 	Three letter character that specifies the instrument type.
		instrument_type;
	
		// 	A unique code that is used to identify the instrument. It is generated by
	// hashing the contract public key hash and the instrument index.
	// SHA256(contract PKH + instrument index)
		instrument_code;
	
		// 	Indices to addresses in outputs and the quantities being frozen. If
	// the only address is the contract address and the instrument code is zeros,
	// then it is a contract wide freeze. If the only address is the contract
	// address and the instrument code is specified, then it is an instrument wide
	// freeze.
		quantities;
	
		// 	Used for a 'time out'. Tokens are automatically unfrozen after the
	// expiration timestamp without requiring a Thaw Action.
		freeze_period;
	
		// 	Timestamp in nanoseconds of when the smart contract created the
	// action.
		timestamp;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeFreeze;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// instrument_type (string)
		{
			WriteFixedChar(buf, this.instrument_type, 3);
		}

		// instrument_code (InstrumentCode)
		{
			buf.write(this.instrument_code.Serialize());
		}

		// quantities ([]QuantityIndex)
		{
			WriteVariableSize(buf, this.quantities.length, 16, 8);
			this.quantities.forEach((value) => {
				buf.write(value.Serialize());
			});
		}

		// freeze_period (Timestamp)
		{
			buf.write(this.freeze_period.Serialize());
		}

		// timestamp (Timestamp)
		{
			buf.write(this.timestamp.Serialize());
		}

		return buf.buf;
	}

	// write populates the fields in Freeze from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// instrument_type (string)
		{
			this.instrument_type = ReadFixedChar(buf, 3);
		}

		// instrument_code (InstrumentCode)
		{
			this.instrument_code = new InstrumentCode();
			this.instrument_code.Write(buf);
		}

		// quantities ([]QuantityIndex)
		{
			// IsInternalTypeArray
			const size = ReadVariableSize(buf, 16, 8);
			this.quantities = [];
			for (let i = 0; i < size; i++) {
				const newValue = new QuantityIndex();
				newValue.Write(buf);

				this.quantities.push(newValue);
			}
		}

		// freeze_period (Timestamp)
		{
			this.freeze_period = new Timestamp();
			this.freeze_period.Write(buf);
		}

		// timestamp (Timestamp)
		{
			this.timestamp = new Timestamp();
			this.timestamp.Write(buf);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// InstrumentType (string)
		{
			if (this.instrument_type.length > 3) {
				return sprintf('fixedchar field instrument_type too long %d/%d', this.instrument_type.length, 3);
			}
		}

		// InstrumentCode (InstrumentCode)
		{
			// IsInternalType
			const err = this.instrument_code.Validate();
			if  (err) return sprintf('field instrument_code is invalid : %s', err);

		}

		// Quantities ([]QuantityIndex) 
		{
			if (this.quantities.length > (2 ** 16) - 1) {
				return sprintf('list field quantities has too many items %d/%d', this.quantities.length, (2 ** 16) - 1);
			}

			const err = this.quantities.find((value, i) => {
				const err2 = value.Validate();
				if (err2) return sprintf('list field quantities[%d] is invalid : %s', i, err2);
			});
			if (err) return err;
		}

		// FreezePeriod (Timestamp) 
		{
			// IsInternalType
			const err = this.freeze_period.Validate();
			if  (err) return sprintf('field freeze_period is invalid : %s', err);

		}

		// Timestamp (Timestamp) 
		{
			// IsInternalType
			const err = this.timestamp.Validate();
			if  (err) return sprintf('field timestamp is invalid : %s', err);

		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('instrument_type:%s', this.instrument_type.toString()));
			vals.push(sprintf('instrument_code:%s', this.instrument_code.toString()));
			vals.push(sprintf('quantities:%s', this.quantities.toString()));
			vals.push(sprintf('freeze_period:%s', this.freeze_period.toString()));
			vals.push(sprintf('timestamp:%s', this.timestamp.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for Freeze.', e);
			throw e;
		}
		// this.instrument_type = parsed.instrumentType (fixedchar)
		
		this.instrument_code = new InstrumentCode();
		this.instrument_code.fromJSON(JSON.stringify(parsed.instrumentCode));
		
		// this.quantities (QuantityIndex[])
		
		this.freeze_period = new Timestamp();
		this.freeze_period.fromJSON(JSON.stringify(parsed.freezePeriod));
		
		this.timestamp = new Timestamp();
		this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
		
		return this.Validate();
	}
}

	// Thaw The contract responding to an Order action to thaw instruments. To be
// used to comply with contractual obligations or legal requirements. The
// Alleged Offender's tokens will be unfrozen to allow them to resume
// normal exchange and governance activities.
	export class Thaw extends OpReturnMessage {
		type = ActionCode.CodeThaw;
		typeStr = 'Thaw';

	
		// 	The tx id of the freeze action that is being reversed.
		freeze_tx_id;
	
		// 	Timestamp in nanoseconds of when the smart contract created the
	// action.
		timestamp;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeThaw;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// freeze_tx_id (TxId)
		{
			buf.write(this.freeze_tx_id.Serialize());
		}

		// timestamp (Timestamp)
		{
			buf.write(this.timestamp.Serialize());
		}

		return buf.buf;
	}

	// write populates the fields in Thaw from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// freeze_tx_id (TxId)
		{
			this.freeze_tx_id = new TxId();
			this.freeze_tx_id.Write(buf);
		}

		// timestamp (Timestamp)
		{
			this.timestamp = new Timestamp();
			this.timestamp.Write(buf);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// FreezeTxId (TxId) 
		{
			// IsInternalType
			const err = this.freeze_tx_id.Validate();
			if  (err) return sprintf('field freeze_tx_id is invalid : %s', err);

		}

		// Timestamp (Timestamp) 
		{
			// IsInternalType
			const err = this.timestamp.Validate();
			if  (err) return sprintf('field timestamp is invalid : %s', err);

		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('freeze_tx_id:%s', this.freeze_tx_id.toString()));
			vals.push(sprintf('timestamp:%s', this.timestamp.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for Thaw.', e);
			throw e;
		}
		this.freeze_tx_id = new TxId();
		this.freeze_tx_id.fromJSON(JSON.stringify(parsed.freezeTxId));
		
		this.timestamp = new Timestamp();
		this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
		
		return this.Validate();
	}
}

	// Confiscation The contract responding to an Order action to confiscate
// instruments. To be used to comply with contractual obligations, legal and/or
// issuer requirements.
	export class Confiscation extends OpReturnMessage {
		type = ActionCode.CodeConfiscation;
		typeStr = 'Confiscation';

	
		// 	Three letter character that specifies the instrument type.
		instrument_type;
	
		// 	A unique code that is used to identify the instrument. It is generated by
	// hashing the contract public key hash and the instrument index.
	// SHA256(contract PKH + instrument index)
		instrument_code;
	
		// 	The holders effected by the confiscation and their balance remaining.
		quantities;
	
		// 	Deposit address's token balance after confiscation.
		deposit_qty;
	
		// 	Timestamp in nanoseconds of when the smart contract created the
	// action.
		timestamp;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeConfiscation;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// instrument_type (string)
		{
			WriteFixedChar(buf, this.instrument_type, 3);
		}

		// instrument_code (InstrumentCode)
		{
			buf.write(this.instrument_code.Serialize());
		}

		// quantities ([]QuantityIndex)
		{
			WriteVariableSize(buf, this.quantities.length, 16, 8);
			this.quantities.forEach((value) => {
				buf.write(value.Serialize());
			});
		}

		// deposit_qty (uint64)
		{
			write(buf, this.deposit_qty, 'uint64');
		}

		// timestamp (Timestamp)
		{
			buf.write(this.timestamp.Serialize());
		}

		return buf.buf;
	}

	// write populates the fields in Confiscation from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// instrument_type (string)
		{
			this.instrument_type = ReadFixedChar(buf, 3);
		}

		// instrument_code (InstrumentCode)
		{
			this.instrument_code = new InstrumentCode();
			this.instrument_code.Write(buf);
		}

		// quantities ([]QuantityIndex)
		{
			// IsInternalTypeArray
			const size = ReadVariableSize(buf, 16, 8);
			this.quantities = [];
			for (let i = 0; i < size; i++) {
				const newValue = new QuantityIndex();
				newValue.Write(buf);

				this.quantities.push(newValue);
			}
		}

		// deposit_qty (uint64)
		{
			this.deposit_qty = read(buf, 'uint64');
		}

		// timestamp (Timestamp)
		{
			this.timestamp = new Timestamp();
			this.timestamp.Write(buf);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// InstrumentType (string)
		{
			if (this.instrument_type.length > 3) {
				return sprintf('fixedchar field instrument_type too long %d/%d', this.instrument_type.length, 3);
			}
		}

		// InstrumentCode (InstrumentCode)
		{
			// IsInternalType
			const err = this.instrument_code.Validate();
			if  (err) return sprintf('field instrument_code is invalid : %s', err);

		}

		// Quantities ([]QuantityIndex) 
		{
			if (this.quantities.length > (2 ** 16) - 1) {
				return sprintf('list field quantities has too many items %d/%d', this.quantities.length, (2 ** 16) - 1);
			}

			const err = this.quantities.find((value, i) => {
				const err2 = value.Validate();
				if (err2) return sprintf('list field quantities[%d] is invalid : %s', i, err2);
			});
			if (err) return err;
		}

		// DepositQty (uint64) 
		{
		}

		// Timestamp (Timestamp) 
		{
			// IsInternalType
			const err = this.timestamp.Validate();
			if  (err) return sprintf('field timestamp is invalid : %s', err);

		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('instrument_type:%s', this.instrument_type.toString()));
			vals.push(sprintf('instrument_code:%s', this.instrument_code.toString()));
			vals.push(sprintf('quantities:%s', this.quantities.toString()));
			vals.push(sprintf('deposit_qty:%d', this.deposit_qty));
			vals.push(sprintf('timestamp:%s', this.timestamp.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for Confiscation.', e);
			throw e;
		}
		// this.instrument_type = parsed.instrumentType (fixedchar)
		
		this.instrument_code = new InstrumentCode();
		this.instrument_code.fromJSON(JSON.stringify(parsed.instrumentCode));
		
		// this.quantities (QuantityIndex[])
		
		this.deposit_qty = parsed.depositQty;
		
		this.timestamp = new Timestamp();
		this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
		
		return this.Validate();
	}
}

	// Reconciliation The contract responding to an Order action to reconcile
// instruments. To be used at the direction of the administration to fix record
// keeping errors with bitcoin and token balances.
	export class Reconciliation extends OpReturnMessage {
		type = ActionCode.CodeReconciliation;
		typeStr = 'Reconciliation';

	
		// 	Three letter character that specifies the instrument type.
		instrument_type;
	
		// 	A unique code that is used to identify the instrument. It is generated by
	// hashing the contract public key hash and the instrument index.
	// SHA256(contract PKH + instrument index)
		instrument_code;
	
		// 	The holders effected by the reconciliation and their balance
	// remaining.
		quantities;
	
		// 	Timestamp in nanoseconds of when the smart contract created the
	// action.
		timestamp;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeReconciliation;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// instrument_type (string)
		{
			WriteFixedChar(buf, this.instrument_type, 3);
		}

		// instrument_code (InstrumentCode)
		{
			buf.write(this.instrument_code.Serialize());
		}

		// quantities ([]QuantityIndex)
		{
			WriteVariableSize(buf, this.quantities.length, 16, 8);
			this.quantities.forEach((value) => {
				buf.write(value.Serialize());
			});
		}

		// timestamp (Timestamp)
		{
			buf.write(this.timestamp.Serialize());
		}

		return buf.buf;
	}

	// write populates the fields in Reconciliation from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// instrument_type (string)
		{
			this.instrument_type = ReadFixedChar(buf, 3);
		}

		// instrument_code (InstrumentCode)
		{
			this.instrument_code = new InstrumentCode();
			this.instrument_code.Write(buf);
		}

		// quantities ([]QuantityIndex)
		{
			// IsInternalTypeArray
			const size = ReadVariableSize(buf, 16, 8);
			this.quantities = [];
			for (let i = 0; i < size; i++) {
				const newValue = new QuantityIndex();
				newValue.Write(buf);

				this.quantities.push(newValue);
			}
		}

		// timestamp (Timestamp)
		{
			this.timestamp = new Timestamp();
			this.timestamp.Write(buf);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// InstrumentType (string)
		{
			if (this.instrument_type.length > 3) {
				return sprintf('fixedchar field instrument_type too long %d/%d', this.instrument_type.length, 3);
			}
		}

		// InstrumentCode (InstrumentCode)
		{
			// IsInternalType
			const err = this.instrument_code.Validate();
			if  (err) return sprintf('field instrument_code is invalid : %s', err);

		}

		// Quantities ([]QuantityIndex) 
		{
			if (this.quantities.length > (2 ** 16) - 1) {
				return sprintf('list field quantities has too many items %d/%d', this.quantities.length, (2 ** 16) - 1);
			}

			const err = this.quantities.find((value, i) => {
				const err2 = value.Validate();
				if (err2) return sprintf('list field quantities[%d] is invalid : %s', i, err2);
			});
			if (err) return err;
		}

		// Timestamp (Timestamp) 
		{
			// IsInternalType
			const err = this.timestamp.Validate();
			if  (err) return sprintf('field timestamp is invalid : %s', err);

		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('instrument_type:%s', this.instrument_type.toString()));
			vals.push(sprintf('instrument_code:%s', this.instrument_code.toString()));
			vals.push(sprintf('quantities:%s', this.quantities.toString()));
			vals.push(sprintf('timestamp:%s', this.timestamp.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for Reconciliation.', e);
			throw e;
		}
		// this.instrument_type = parsed.instrumentType (fixedchar)
		
		this.instrument_code = new InstrumentCode();
		this.instrument_code.fromJSON(JSON.stringify(parsed.instrumentCode));
		
		// this.quantities (QuantityIndex[])
		
		this.timestamp = new Timestamp();
		this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
		
		return this.Validate();
	}
}

	// Proposal Allows the Administration/Token Holders to propose a change
// (aka Initiative/Shareholder vote). A significant cost - specified in the
// Contract Formation - can be attached to this action when sent from Token
// Holders to reduce spam, as the resulting vote will be put to all token
// owners.
	export class Proposal extends OpReturnMessage {
		type = ActionCode.CodeProposal;
		typeStr = 'Proposal';

	
		// 	Who initiated the proposal. Supported values: 0 - Administration, 1 -
	// Holder
		initiator;
	
		// 	When true this proposal is specific to an instrument and the instrument type and
	// instrument code fields are serialized.
		instrument_specific_vote;
	
		// 	Three letter character that specifies the instrument type.
		instrument_type;
	
		// 	A unique code that is used to identify the instrument. It is generated by
	// hashing the contract public key hash and the instrument index.
	// SHA256(contract PKH + instrument index)
		instrument_code;
	
		// 	X for Vote System X. (1-255, 0 is not valid.)
		vote_system;
	
		// 	When true the ProposedAmendments field is included and specifies the
	// exact changes to make to the contract/instrument on chain. When false this
	// is just a general proposal like a strategy/direction and doesn't result
	// in any on chain update.
		specific;
	
		// 	Each element contains details of which fields to modify, or delete.
	// Because the number of fields in a Contract and Instrument is dynamic due to
	// some fields being able to be repeated, the index value of the field
	// needs to be calculated against the Contract or Instrument the changes are to
	// apply to. In the event of a Vote being created from this Initiative,
	// the changes will be applied to the version of the Contract or Instrument at
	// that time.
		proposed_amendments;
	
		// 	Length 1-255 bytes. 0 is not valid. Each byte allows for a different
	// vote option. Typical votes will likely be multiple choice or Y/N. Vote
	// instances are identified by the Tx-ID. AB would be used for Y/N
	// (binary) type votes. When Specific is true, only AB is a valid value.
		vote_options;
	
		// 	Range: 1-X. How many selections can a voter make in a Ballot Cast. 1
	// is selected for Y/N (binary). When Specific is true, only 1 is a valid
	// value.
		vote_max;
	
		// 	Length restricted by the Bitcoin protocol. 0 is valid. Description or
	// details of the vote
		proposal_description;
	
		// 	SHA256 Hash of the proposal document to be distributed to voters.
		proposal_document_hash;
	
		// 	Ballot casts after this timestamp will not be included. The vote has
	// finished.
		vote_cut_off_timestamp;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeProposal;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// initiator (uint8)
		{
			write(buf, this.initiator, 'uint8');
		}

		// instrument_specific_vote (bool)
		{
			write(buf, this.instrument_specific_vote, 'bool');
		}

		// instrument_type (string)
		if (this.instrument_specific_vote) {
			WriteFixedChar(buf, this.instrument_type, 3);
		}

		// instrument_code (InstrumentCode)
		if (this.instrument_specific_vote) {
			buf.write(this.instrument_code.Serialize());
		}

		// vote_system (uint8)
		{
			write(buf, this.vote_system, 'uint8');
		}

		// specific (bool)
		{
			write(buf, this.specific, 'bool');
		}

		// proposed_amendments ([]Amendment)
		if (this.specific) {
			WriteVariableSize(buf, this.proposed_amendments.length, 8, 8);
			this.proposed_amendments.forEach((value) => {
				buf.write(value.Serialize());
			});
		}

		// vote_options (string)
		{
			WriteVarChar(buf, this.vote_options, 8);
		}

		// vote_max (uint8)
		{
			write(buf, this.vote_max, 'uint8');
		}

		// proposal_description (string)
		{
			WriteVarChar(buf, this.proposal_description, 32);
		}

		// proposal_document_hash ([32]byte)
		{
			write(buf, this.proposal_document_hash, '[32]byte');
		}

		// vote_cut_off_timestamp (Timestamp)
		{
			buf.write(this.vote_cut_off_timestamp.Serialize());
		}

		return buf.buf;
	}

	// write populates the fields in Proposal from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// initiator (uint8)
		{
			this.initiator = read(buf, 'uint8');
		}

		// instrument_specific_vote (bool)
		{
			this.instrument_specific_vote = read(buf, 'bool');
		}

		// instrument_type (string)
		if (this.instrument_specific_vote) {
			this.instrument_type = ReadFixedChar(buf, 3);
		}

		// instrument_code (InstrumentCode)
		if (this.instrument_specific_vote) {
			this.instrument_code = new InstrumentCode();
			this.instrument_code.Write(buf);
		}

		// vote_system (uint8)
		{
			this.vote_system = read(buf, 'uint8');
		}

		// specific (bool)
		{
			this.specific = read(buf, 'bool');
		}

		// proposed_amendments ([]Amendment)
		if (this.specific) {
			// IsInternalTypeArray
			const size = ReadVariableSize(buf, 8, 8);
			this.proposed_amendments = [];
			for (let i = 0; i < size; i++) {
				const newValue = new Amendment();
				newValue.Write(buf);

				this.proposed_amendments.push(newValue);
			}
		}

		// vote_options (string)
		{
			this.vote_options = ReadVarChar(buf, 8);
		}

		// vote_max (uint8)
		{
			this.vote_max = read(buf, 'uint8');
		}

		// proposal_description (string)
		{
			this.proposal_description = ReadVarChar(buf, 32);
		}

		// proposal_document_hash ([32]byte)
		{
			this.proposal_document_hash = read(buf, '[32]byte');
		}

		// vote_cut_off_timestamp (Timestamp)
		{
			this.vote_cut_off_timestamp = new Timestamp();
			this.vote_cut_off_timestamp.Write(buf);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// Initiator (uint8) 
		{
			// $field.IntValues [0 1]
			if ( this.initiator !== 0 && this.initiator !== 1) {
				return sprintf('field initiator value is invalid : %d', this.initiator);
			}

		}

		// InstrumentSpecificVote (bool)
		{
		}

		// InstrumentType (string)
		// IncludeIfTrue
		if (this.instrument_specific_vote) {
			if (this.instrument_type.length > 3) {
				return sprintf('fixedchar field instrument_type too long %d/%d', this.instrument_type.length, 3);
			}
		}

		// InstrumentCode (InstrumentCode)
		// IncludeIfTrue
		if (this.instrument_specific_vote) {
			// IsInternalType
			const err = this.instrument_code.Validate();
			if  (err) return sprintf('field instrument_code is invalid : %s', err);

		}

		// VoteSystem (uint8) 
		{
		}

		// Specific (bool) 
		{
		}

		// ProposedAmendments ([]Amendment)
		// IncludeIfTrue
		if (this.specific) {
			if (this.proposed_amendments.length > (2 ** 8) - 1) {
				return sprintf('list field proposed_amendments has too many items %d/%d', this.proposed_amendments.length, (2 ** 8) - 1);
			}

			const err = this.proposed_amendments.find((value, i) => {
				const err2 = value.Validate();
				if (err2) return sprintf('list field proposed_amendments[%d] is invalid : %s', i, err2);
			});
			if (err) return err;
		}

		// VoteOptions (string) 
		{
			if (this.vote_options.length > (2 ** 8)) {
				return sprintf('varchar field vote_options too long %d/%d', this.vote_options.length, (2 ** 8) - 1);
			}
		}

		// VoteMax (uint8) 
		{
		}

		// ProposalDescription (string) 
		{
			if (this.proposal_description.length > (2 ** 32)) {
				return sprintf('varchar field proposal_description too long %d/%d', this.proposal_description.length, (2 ** 32) - 1);
			}
		}

		// ProposalDocumentHash ([32]byte) 
		{
		}

		// VoteCutOffTimestamp (Timestamp) 
		{
			// IsInternalType
			const err = this.vote_cut_off_timestamp.Validate();
			if  (err) return sprintf('field vote_cut_off_timestamp is invalid : %s', err);

		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('initiator:%d', this.initiator));
			vals.push(sprintf('instrument_specific_vote:%s', this.instrument_specific_vote.toString()));
			vals.push(sprintf('instrument_type:%s', this.instrument_type.toString()));
			vals.push(sprintf('instrument_code:%s', this.instrument_code.toString()));
			vals.push(sprintf('vote_system:%d', this.vote_system));
			vals.push(sprintf('specific:%s', this.specific.toString()));
			vals.push(sprintf('proposed_amendments:%s', this.proposed_amendments.toString()));
			vals.push(sprintf('vote_options:%s', this.vote_options.toString()));
			vals.push(sprintf('vote_max:%d', this.vote_max));
			vals.push(sprintf('proposal_description:%s', this.proposal_description.toString()));
			vals.push(sprintf('proposal_document_hash:%s', this.proposal_document_hash.toString()));
			vals.push(sprintf('vote_cut_off_timestamp:%s', this.vote_cut_off_timestamp.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for Proposal.', e);
			throw e;
		}
		this.initiator = parsed.initiator;
		
		// this.instrument_specific_vote = parsed.instrumentSpecificVote (bool)
		
		// this.instrument_type = parsed.instrumentType (fixedchar)
		
		this.instrument_code = new InstrumentCode();
		this.instrument_code.fromJSON(JSON.stringify(parsed.instrumentCode));
		
		this.vote_system = parsed.voteSystem;
		
		// this.specific = parsed.specific (bool)
		
		// this.proposed_amendments (Amendment[])
		
		this.vote_options = parsed.voteOptions;
		
		this.vote_max = parsed.voteMax;
		
		this.proposal_description = parsed.proposalDescription;
		
		// this.proposal_document_hash = parsed.proposalDocumentHash (bin)
		
		this.vote_cut_off_timestamp = new Timestamp();
		this.vote_cut_off_timestamp.fromJSON(JSON.stringify(parsed.voteCutOffTimestamp));
		
		return this.Validate();
	}
}

	// Vote A vote is created by the Contract in response to a valid Proposal
// Action.
	export class Vote extends OpReturnMessage {
		type = ActionCode.CodeVote;
		typeStr = 'Vote';

	
		// 	Timestamp in nanoseconds of when the smart contract created the
	// action.
		timestamp;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeVote;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// timestamp (Timestamp)
		{
			buf.write(this.timestamp.Serialize());
		}

		return buf.buf;
	}

	// write populates the fields in Vote from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// timestamp (Timestamp)
		{
			this.timestamp = new Timestamp();
			this.timestamp.Write(buf);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// Timestamp (Timestamp) 
		{
			// IsInternalType
			const err = this.timestamp.Validate();
			if  (err) return sprintf('field timestamp is invalid : %s', err);

		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('timestamp:%s', this.timestamp.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for Vote.', e);
			throw e;
		}
		this.timestamp = new Timestamp();
		this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
		
		return this.Validate();
	}
}

	// BallotCast Used by Token Owners to cast their ballot (vote) on
// proposals. 1 Vote per token unless a vote multiplier is specified in the
// relevant Instrument Definition action.
	export class BallotCast extends OpReturnMessage {
		type = ActionCode.CodeBallotCast;
		typeStr = 'BallotCast';

	
		// 	Tx ID of the Vote the Ballot Cast is being made for.
		vote_tx_id;
	
		// 	Length 1-255 bytes. 0 is not valid. Max length is the VoteMax value
	// from the Proposal action. Accept, Reject, Abstain, Spoiled, Multiple
	// Choice, or Preference List. 15 options total. Order of preference. 1st
	// position = 1st choice. 2nd position = 2nd choice, etc. A is always
	// Accept and B is always reject in a Y/N votes.
		vote;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeBallotCast;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// vote_tx_id (TxId)
		{
			buf.write(this.vote_tx_id.Serialize());
		}

		// vote (string)
		{
			WriteVarChar(buf, this.vote, 8);
		}

		return buf.buf;
	}

	// write populates the fields in BallotCast from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// vote_tx_id (TxId)
		{
			this.vote_tx_id = new TxId();
			this.vote_tx_id.Write(buf);
		}

		// vote (string)
		{
			this.vote = ReadVarChar(buf, 8);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// VoteTxId (TxId) 
		{
			// IsInternalType
			const err = this.vote_tx_id.Validate();
			if  (err) return sprintf('field vote_tx_id is invalid : %s', err);

		}

		// Vote (string) 
		{
			if (this.vote.length > (2 ** 8)) {
				return sprintf('varchar field vote too long %d/%d', this.vote.length, (2 ** 8) - 1);
			}
		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('vote_tx_id:%s', this.vote_tx_id.toString()));
			vals.push(sprintf('vote:%s', this.vote.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for BallotCast.', e);
			throw e;
		}
		this.vote_tx_id = new TxId();
		this.vote_tx_id.fromJSON(JSON.stringify(parsed.voteTxId));
		
		this.vote = parsed.vote;
		
		return this.Validate();
	}
}

	// BallotCounted The smart contract will respond to a Ballot Cast action
// with a Ballot Counted action if the Ballot Cast is valid. If the Ballot
// Cast is not valid, then the smart contract will respond with a Rejection
// Action.
	export class BallotCounted extends OpReturnMessage {
		type = ActionCode.CodeBallotCounted;
		typeStr = 'BallotCounted';

	
		// 	Tx ID of the Vote the Ballot Cast is being made for.
		vote_tx_id;
	
		// 	Length 1-255 bytes. 0 is not valid. Max length is the VoteMax value
	// from the Proposal action. Accept, Reject, Abstain, Spoiled, Multiple
	// Choice, or Preference List. 15 options total. Order of preference. 1st
	// position = 1st choice. 2nd position = 2nd choice, etc. A is always
	// Accept and B is always reject in a Y/N votes.
		vote;
	
		// 	Number of votes counted for this ballot. Factors in vote multipliers
	// if there are any allowed, otherwise it is just quantity of tokens held.
		quantity;
	
		// 	Timestamp in nanoseconds of when the smart contract created the
	// action.
		timestamp;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeBallotCounted;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// vote_tx_id (TxId)
		{
			buf.write(this.vote_tx_id.Serialize());
		}

		// vote (string)
		{
			WriteVarChar(buf, this.vote, 8);
		}

		// quantity (uint64)
		{
			write(buf, this.quantity, 'uint64');
		}

		// timestamp (Timestamp)
		{
			buf.write(this.timestamp.Serialize());
		}

		return buf.buf;
	}

	// write populates the fields in BallotCounted from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// vote_tx_id (TxId)
		{
			this.vote_tx_id = new TxId();
			this.vote_tx_id.Write(buf);
		}

		// vote (string)
		{
			this.vote = ReadVarChar(buf, 8);
		}

		// quantity (uint64)
		{
			this.quantity = read(buf, 'uint64');
		}

		// timestamp (Timestamp)
		{
			this.timestamp = new Timestamp();
			this.timestamp.Write(buf);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// VoteTxId (TxId) 
		{
			// IsInternalType
			const err = this.vote_tx_id.Validate();
			if  (err) return sprintf('field vote_tx_id is invalid : %s', err);

		}

		// Vote (string) 
		{
			if (this.vote.length > (2 ** 8)) {
				return sprintf('varchar field vote too long %d/%d', this.vote.length, (2 ** 8) - 1);
			}
		}

		// Quantity (uint64) 
		{
		}

		// Timestamp (Timestamp) 
		{
			// IsInternalType
			const err = this.timestamp.Validate();
			if  (err) return sprintf('field timestamp is invalid : %s', err);

		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('vote_tx_id:%s', this.vote_tx_id.toString()));
			vals.push(sprintf('vote:%s', this.vote.toString()));
			vals.push(sprintf('quantity:%d', this.quantity));
			vals.push(sprintf('timestamp:%s', this.timestamp.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for BallotCounted.', e);
			throw e;
		}
		this.vote_tx_id = new TxId();
		this.vote_tx_id.fromJSON(JSON.stringify(parsed.voteTxId));
		
		this.vote = parsed.vote;
		
		this.quantity = parsed.quantity;
		
		this.timestamp = new Timestamp();
		this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
		
		return this.Validate();
	}
}

	// Result Once a vote has been completed the results are published. After
// the result is posted, it is up to the administration to send a
// contract/instrument amendement if appropriate.
	export class Result extends OpReturnMessage {
		type = ActionCode.CodeResult;
		typeStr = 'Result';

	
		// 	When true this proposal is specific to an instrument and the instrument type and
	// instrument code fields are serialized.
		instrument_specific_vote;
	
		// 	Three letter character that specifies the instrument type.
		instrument_type;
	
		// 	A unique code that is used to identify the instrument. It is generated by
	// hashing the contract public key hash and the instrument index.
	// SHA256(contract PKH + instrument index)
		instrument_code;
	
		// 	When true the ProposedAmendments field is included and specifies the
	// exact changes to make to the Contract/Instrument on chain. When false this
	// is just a general proposal like a strategy/direction and doesn't result
	// in any on chain update.
		specific;
	
		// 	Each element contains details of which fields to modify, or delete.
	// Because the number of fields in a Contract and Instrument is dynamic due to
	// some fields being able to be repeated, the index value of the field
	// needs to be calculated against the Contract or Instrument the changes are to
	// apply to. In the event of a Vote being created from this Initiative,
	// the changes will be applied to the version of the Contract or Instrument at
	// that time.
		proposed_amendments;
	
		// 	Link to the Vote Action txn.
		vote_tx_id;
	
		// 	List of number of valid votes counted for each vote option. Length is
	// encoded like a regular list object, but must match the length of
	// VoteOptions from the Proposal action.
		option_tally;
	
		// 	Length 1-255 bytes. 0 is not valid. The Option with the most votes. In
	// the event of a draw for 1st place, all winning options are listed.
		result;
	
		// 	Timestamp in nanoseconds of when the smart contract created the
	// action.
		timestamp;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeResult;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// instrument_specific_vote (bool)
		{
			write(buf, this.instrument_specific_vote, 'bool');
		}

		// instrument_type (string)
		if (this.instrument_specific_vote) {
			WriteFixedChar(buf, this.instrument_type, 3);
		}

		// instrument_code (InstrumentCode)
		if (this.instrument_specific_vote) {
			buf.write(this.instrument_code.Serialize());
		}

		// specific (bool)
		{
			write(buf, this.specific, 'bool');
		}

		// proposed_amendments ([]Amendment)
		if (this.specific) {
			WriteVariableSize(buf, this.proposed_amendments.length, 8, 8);
			this.proposed_amendments.forEach((value) => {
				buf.write(value.Serialize());
			});
		}

		// vote_tx_id (TxId)
		{
			buf.write(this.vote_tx_id.Serialize());
		}

		// option_tally ([]uint64)
		{
			// IsNativeTypeArray []uint64
			const type = '[]uint64'.slice(2);
			WriteVariableSize(buf, this.option_tally.length, 8, 8);
			this.option_tally.forEach(value => {
				write(buf, value, type); // TODO
			});
		}

		// result (string)
		{
			WriteVarChar(buf, this.result, 8);
		}

		// timestamp (Timestamp)
		{
			buf.write(this.timestamp.Serialize());
		}

		return buf.buf;
	}

	// write populates the fields in Result from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// instrument_specific_vote (bool)
		{
			this.instrument_specific_vote = read(buf, 'bool');
		}

		// instrument_type (string)
		if (this.instrument_specific_vote) {
			this.instrument_type = ReadFixedChar(buf, 3);
		}

		// instrument_code (InstrumentCode)
		if (this.instrument_specific_vote) {
			this.instrument_code = new InstrumentCode();
			this.instrument_code.Write(buf);
		}

		// specific (bool)
		{
			this.specific = read(buf, 'bool');
		}

		// proposed_amendments ([]Amendment)
		if (this.specific) {
			// IsInternalTypeArray
			const size = ReadVariableSize(buf, 8, 8);
			this.proposed_amendments = [];
			for (let i = 0; i < size; i++) {
				const newValue = new Amendment();
				newValue.Write(buf);

				this.proposed_amendments.push(newValue);
			}
		}

		// vote_tx_id (TxId)
		{
			this.vote_tx_id = new TxId();
			this.vote_tx_id.Write(buf);
		}

		// option_tally ([]uint64)
		{
			const size = ReadVariableSize(buf, 8, 8);
			const type = '[]uint64'.slice(2);
			this.option_tally = [...Array(size)].map(() => read(buf, type));
		}

		// result (string)
		{
			this.result = ReadVarChar(buf, 8);
		}

		// timestamp (Timestamp)
		{
			this.timestamp = new Timestamp();
			this.timestamp.Write(buf);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// InstrumentSpecificVote (bool)
		{
		}

		// InstrumentType (string)
		// IncludeIfTrue
		if (this.instrument_specific_vote) {
			if (this.instrument_type.length > 3) {
				return sprintf('fixedchar field instrument_type too long %d/%d', this.instrument_type.length, 3);
			}
		}

		// InstrumentCode (InstrumentCode)
		// IncludeIfTrue
		if (this.instrument_specific_vote) {
			// IsInternalType
			const err = this.instrument_code.Validate();
			if  (err) return sprintf('field instrument_code is invalid : %s', err);

		}

		// Specific (bool) 
		{
		}

		// ProposedAmendments ([]Amendment)
		// IncludeIfTrue
		if (this.specific) {
			if (this.proposed_amendments.length > (2 ** 8) - 1) {
				return sprintf('list field proposed_amendments has too many items %d/%d', this.proposed_amendments.length, (2 ** 8) - 1);
			}

			const err = this.proposed_amendments.find((value, i) => {
				const err2 = value.Validate();
				if (err2) return sprintf('list field proposed_amendments[%d] is invalid : %s', i, err2);
			});
			if (err) return err;
		}

		// VoteTxId (TxId) 
		{
			// IsInternalType
			const err = this.vote_tx_id.Validate();
			if  (err) return sprintf('field vote_tx_id is invalid : %s', err);

		}

		// OptionTally ([]uint64) 
		{
			if (this.option_tally.length > (2 ** 8) - 1) {
				return sprintf('list field option_tally has too many items %d/%d', this.option_tally.length, (2 ** 8) - 1);
			}
		}

		// Result (string) 
		{
			if (this.result.length > (2 ** 8)) {
				return sprintf('varchar field result too long %d/%d', this.result.length, (2 ** 8) - 1);
			}
		}

		// Timestamp (Timestamp) 
		{
			// IsInternalType
			const err = this.timestamp.Validate();
			if  (err) return sprintf('field timestamp is invalid : %s', err);

		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('instrument_specific_vote:%s', this.instrument_specific_vote.toString()));
			vals.push(sprintf('instrument_type:%s', this.instrument_type.toString()));
			vals.push(sprintf('instrument_code:%s', this.instrument_code.toString()));
			vals.push(sprintf('specific:%s', this.specific.toString()));
			vals.push(sprintf('proposed_amendments:%s', this.proposed_amendments.toString()));
			vals.push(sprintf('vote_tx_id:%s', this.vote_tx_id.toString()));
			vals.push(sprintf('option_tally:%d', this.option_tally));
			vals.push(sprintf('result:%s', this.result.toString()));
			vals.push(sprintf('timestamp:%s', this.timestamp.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for Result.', e);
			throw e;
		}
		// this.instrument_specific_vote = parsed.instrumentSpecificVote (bool)
		
		// this.instrument_type = parsed.instrumentType (fixedchar)
		
		this.instrument_code = new InstrumentCode();
		this.instrument_code.fromJSON(JSON.stringify(parsed.instrumentCode));
		
		// this.specific = parsed.specific (bool)
		
		// this.proposed_amendments (Amendment[])
		
		this.vote_tx_id = new TxId();
		this.vote_tx_id.fromJSON(JSON.stringify(parsed.voteTxId));
		
		this.option_tally = parsed.optionTally;
		
		this.result = parsed.result;
		
		this.timestamp = new Timestamp();
		this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
		
		return this.Validate();
	}
}

	// Message The message action is a general purpose communication action.
// 'Twitter/SMS' for Issuers/Investors/Users. The message txn can also be
// used for passing partially signed txns on-chain, establishing private
// communication channels and EDI (receipting, invoices, PO, and private
// offers/bids). The messages are broken down by type for easy filtering in
// the a user's wallet. The Message Types are listed in the Message Types
// table.
	export class Message extends OpReturnMessage {
		type = ActionCode.CodeMessage;
		typeStr = 'Message';

	
		// 	Associates the message to a particular output by the index.
		address_indexes;
	
		// 	Potential for up to 65,535 different message types. Values from
	// resources/Messages.yaml
		message_type;
	
		// 	Public or private (RSA public key, Diffie-Hellman). Issuers/Contracts
	// can send the signifying amount of satoshis to themselves for public
	// announcements or private 'notes' if encrypted. See Message Types for a
	// full list of potential use cases.
		message_payload;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeMessage;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// address_indexes ([]uint16)
		{
			// IsNativeTypeArray []uint16
			const type = '[]uint16'.slice(2);
			WriteVariableSize(buf, this.address_indexes.length, 8, 8);
			this.address_indexes.forEach(value => {
				write(buf, value, type); // TODO
			});
		}

		// message_type (uint16)
		{
			write(buf, this.message_type, 'uint16');
		}

		// message_payload ([]byte)
		{
			WriteVarBin(buf, this.message_payload, 32);
		}

		return buf.buf;
	}

	// write populates the fields in Message from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// address_indexes ([]uint16)
		{
			const size = ReadVariableSize(buf, 8, 8);
			const type = '[]uint16'.slice(2);
			this.address_indexes = [...Array(size)].map(() => read(buf, type));
		}

		// message_type (uint16)
		{
			this.message_type = read(buf, 'uint16');
		}

		// message_payload ([]byte)
		{
			this.message_payload = ReadVarBin(buf, 32);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// AddressIndexes ([]uint16) 
		{
			if (this.address_indexes.length > (2 ** 8) - 1) {
				return sprintf('list field address_indexes has too many items %d/%d', this.address_indexes.length, (2 ** 8) - 1);
			}
		}

		// MessageType (uint16) 
		{
		}

		// MessagePayload ([]byte) 
		{
			if (this.message_payload.length >= (2 ** 32)) {
				return sprintf('varbin field message_payload too long %d/%d', this.message_payload.length, (2 ** 32) - 1);
			}
		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('address_indexes:%d', this.address_indexes));
			vals.push(sprintf('message_type:%d', this.message_type));
			vals.push(sprintf('message_payload:%s', this.message_payload.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for Message.', e);
			throw e;
		}
		this.address_indexes = parsed.addressIndexes;
		
		this.message_type = parsed.messageType;
		
		// this.message_payload = parsed.messagePayload (varbin)
		
		return this.Validate();
	}
}

	// Rejection Used to reject request actions that do not comply with the
// Contract. If money is to be returned to a User then it is used in lieu
// of the Settlement Action to properly account for token balances. All
// Administration/User request Actions must be responded to by the Contract
// with an Action. The only exception to this rule is when there is not
// enough fees in the first Action for the Contract response action to
// remain revenue neutral. If not enough fees are attached to pay for the
// Contract response then the Contract will not respond.
	export class Rejection extends OpReturnMessage {
		type = ActionCode.CodeRejection;
		typeStr = 'Rejection';

	
		// 	Associates the message to a particular output by the index.
		address_indexes;
	
		// 	The address which is believed to have caused the rejection.
		reject_address_index;
	
		// 	Classifies the rejection by a type.
		rejection_code;
	
		// 	Length 0-65,535 bytes. Message that explains the reasoning for a
	// rejection, if needed. Most rejection types will be captured by the
	// Rejection Type Subfield.
		message;
	
		// 	Timestamp in nanoseconds of when the smart contract created the
	// action.
		timestamp;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeRejection;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// address_indexes ([]uint16)
		{
			// IsNativeTypeArray []uint16
			const type = '[]uint16'.slice(2);
			WriteVariableSize(buf, this.address_indexes.length, 8, 8);
			this.address_indexes.forEach(value => {
				write(buf, value, type); // TODO
			});
		}

		// reject_address_index (uint16)
		{
			write(buf, this.reject_address_index, 'uint16');
		}

		// rejection_code (uint8)
		{
			write(buf, this.rejection_code, 'uint8');
		}

		// message (string)
		{
			WriteVarChar(buf, this.message, 16);
		}

		// timestamp (Timestamp)
		{
			buf.write(this.timestamp.Serialize());
		}

		return buf.buf;
	}

	// write populates the fields in Rejection from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// address_indexes ([]uint16)
		{
			const size = ReadVariableSize(buf, 8, 8);
			const type = '[]uint16'.slice(2);
			this.address_indexes = [...Array(size)].map(() => read(buf, type));
		}

		// reject_address_index (uint16)
		{
			this.reject_address_index = read(buf, 'uint16');
		}

		// rejection_code (uint8)
		{
			this.rejection_code = read(buf, 'uint8');
		}

		// message (string)
		{
			this.message = ReadVarChar(buf, 16);
		}

		// timestamp (Timestamp)
		{
			this.timestamp = new Timestamp();
			this.timestamp.Write(buf);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// AddressIndexes ([]uint16) 
		{
			if (this.address_indexes.length > (2 ** 8) - 1) {
				return sprintf('list field address_indexes has too many items %d/%d', this.address_indexes.length, (2 ** 8) - 1);
			}
		}

		// RejectAddressIndex (uint16) 
		{
		}

		// RejectionCode (uint8) 
		{
			if (Resources.GetRejectionCode(this.rejection_code) === null) {
				return sprintf('Invalid rejection code value : %d', this.rejection_code);
			}
		}

		// Message (string) 
		{
			if (this.message.length > (2 ** 16)) {
				return sprintf('varchar field message too long %d/%d', this.message.length, (2 ** 16) - 1);
			}
		}

		// Timestamp (Timestamp) 
		{
			// IsInternalType
			const err = this.timestamp.Validate();
			if  (err) return sprintf('field timestamp is invalid : %s', err);

		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('address_indexes:%d', this.address_indexes));
			vals.push(sprintf('reject_address_index:%d', this.reject_address_index));
			vals.push(sprintf('rejection_code:%s', this.rejection_code.toString()));
			vals.push(sprintf('message:%s', this.message.toString()));
			vals.push(sprintf('timestamp:%s', this.timestamp.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for Rejection.', e);
			throw e;
		}
		this.address_indexes = parsed.addressIndexes;
		
		this.reject_address_index = parsed.rejectAddressIndex;
		
		// this.rejection_code = parsed.rejectionCode (RejectionCode)
		
		this.message = parsed.message;
		
		this.timestamp = new Timestamp();
		this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
		
		return this.Validate();
	}
}

	// Establishment Establishes an on-chain register.
	export class Establishment extends OpReturnMessage {
		type = ActionCode.CodeEstablishment;
		typeStr = 'Establishment';

	
		// 	A custom message to include with this action.
		message;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeEstablishment;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// message (string)
		{
			WriteVarChar(buf, this.message, 32);
		}

		return buf.buf;
	}

	// write populates the fields in Establishment from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// message (string)
		{
			this.message = ReadVarChar(buf, 32);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// Message (string) 
		{
			if (this.message.length > (2 ** 32)) {
				return sprintf('varchar field message too long %d/%d', this.message.length, (2 ** 32) - 1);
			}
		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('message:%s', this.message.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for Establishment.', e);
			throw e;
		}
		this.message = parsed.message;
		
		return this.Validate();
	}
}

	// Addition Adds an entry to the Register.
	export class Addition extends OpReturnMessage {
		type = ActionCode.CodeAddition;
		typeStr = 'Addition';

	
		// 	A custom message to include with this action.
		message;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeAddition;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// message (string)
		{
			WriteVarChar(buf, this.message, 32);
		}

		return buf.buf;
	}

	// write populates the fields in Addition from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// message (string)
		{
			this.message = ReadVarChar(buf, 32);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// Message (string) 
		{
			if (this.message.length > (2 ** 32)) {
				return sprintf('varchar field message too long %d/%d', this.message.length, (2 ** 32) - 1);
			}
		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('message:%s', this.message.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for Addition.', e);
			throw e;
		}
		this.message = parsed.message;
		
		return this.Validate();
	}
}

	// Alteration A register entry/record can be altered.
	export class Alteration extends OpReturnMessage {
		type = ActionCode.CodeAlteration;
		typeStr = 'Alteration';

	
		// 	Transaction ID of the register entry to be altered.
		entry_tx_id;
	
		// 	A custom message to include with this action.
		message;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeAlteration;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// entry_tx_id (TxId)
		{
			buf.write(this.entry_tx_id.Serialize());
		}

		// message (string)
		{
			WriteVarChar(buf, this.message, 32);
		}

		return buf.buf;
	}

	// write populates the fields in Alteration from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// entry_tx_id (TxId)
		{
			this.entry_tx_id = new TxId();
			this.entry_tx_id.Write(buf);
		}

		// message (string)
		{
			this.message = ReadVarChar(buf, 32);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// EntryTxID (TxId) 
		{
			// IsInternalType
			const err = this.entry_tx_id.Validate();
			if  (err) return sprintf('field entry_tx_id is invalid : %s', err);

		}

		// Message (string) 
		{
			if (this.message.length > (2 ** 32)) {
				return sprintf('varchar field message too long %d/%d', this.message.length, (2 ** 32) - 1);
			}
		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('entry_tx_id:%s', this.entry_tx_id.toString()));
			vals.push(sprintf('message:%s', this.message.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for Alteration.', e);
			throw e;
		}
		this.entry_tx_id = new TxId();
		this.entry_tx_id.fromJSON(JSON.stringify(parsed.entryTxID));
		
		this.message = parsed.message;
		
		return this.Validate();
	}
}

	// Removal Removes an entry/record from the Register.
	export class Removal extends OpReturnMessage {
		type = ActionCode.CodeRemoval;
		typeStr = 'Removal';

	
		// 	Transaction ID of the register entry to be altered.
		entry_tx_id;
	
		// 	A custom message to include with this action.
		message;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeRemoval;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// entry_tx_id (TxId)
		{
			buf.write(this.entry_tx_id.Serialize());
		}

		// message (string)
		{
			WriteVarChar(buf, this.message, 32);
		}

		return buf.buf;
	}

	// write populates the fields in Removal from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// entry_tx_id (TxId)
		{
			this.entry_tx_id = new TxId();
			this.entry_tx_id.Write(buf);
		}

		// message (string)
		{
			this.message = ReadVarChar(buf, 32);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// EntryTxID (TxId) 
		{
			// IsInternalType
			const err = this.entry_tx_id.Validate();
			if  (err) return sprintf('field entry_tx_id is invalid : %s', err);

		}

		// Message (string) 
		{
			if (this.message.length > (2 ** 32)) {
				return sprintf('varchar field message too long %d/%d', this.message.length, (2 ** 32) - 1);
			}
		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('entry_tx_id:%s', this.entry_tx_id.toString()));
			vals.push(sprintf('message:%s', this.message.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for Removal.', e);
			throw e;
		}
		this.entry_tx_id = new TxId();
		this.entry_tx_id.fromJSON(JSON.stringify(parsed.entryTxID));
		
		this.message = parsed.message;
		
		return this.Validate();
	}
}

	// Transfer A Token Owner(s) Sends, Exchanges or Swaps a token(s) or
// Bitcoin for a token(s) or Bitcoin. Can be as simple as sending a single
// token to a receiver. Or can be as complex as many senders sending many
// different instruments - controlled by many different smart contracts - to a
// number of receivers. This action also supports atomic swaps (tokens for
// tokens). Since many parties and contracts can be involved in a transfer
// and the corresponding settlement action, the partially signed T1 and T2
// actions will need to be passed around on-chain with an M1 action, or
// off-chain.
	export class Transfer extends OpReturnMessage {
		type = ActionCode.CodeTransfer;
		typeStr = 'Transfer';

	
		// 	The Instruments involved in the Transfer Action.
		instruments;
	
		// 	This prevents any party from holding on to the partially signed
	// message as a form of an option. Eg. the exchange at this price is valid
	// for 30 mins.
		offer_expiry;
	
		// 	Fixed amount of bitcoin being paid to an exchange for facilitating a
	// transfer.
		exchange_fee;
	
		// 	Identifies the public address that the exchange fee should be paid to.
		exchange_fee_address;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeTransfer;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// instruments ([]InstrumentTransfer)
		{
			WriteVariableSize(buf, this.instruments.length, 8, 8);
			this.instruments.forEach((value) => {
				buf.write(value.Serialize());
			});
		}

		// offer_expiry (Timestamp)
		{
			buf.write(this.offer_expiry.Serialize());
		}

		// exchange_fee (uint64)
		{
			write(buf, this.exchange_fee, 'uint64');
		}

		// exchange_fee_address (PublicKeyHash)
		{
			buf.write(this.exchange_fee_address.Serialize());
		}

		return buf.buf;
	}

	// write populates the fields in Transfer from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// instruments ([]InstrumentTransfer)
		{
			// IsInternalTypeArray
			const size = ReadVariableSize(buf, 8, 8);
			this.instruments = [];
			for (let i = 0; i < size; i++) {
				const newValue = new InstrumentTransfer();
				newValue.Write(buf);

				this.instruments.push(newValue);
			}
		}

		// offer_expiry (Timestamp)
		{
			this.offer_expiry = new Timestamp();
			this.offer_expiry.Write(buf);
		}

		// exchange_fee (uint64)
		{
			this.exchange_fee = read(buf, 'uint64');
		}

		// exchange_fee_address (PublicKeyHash)
		{
			this.exchange_fee_address = new PublicKeyHash();
			this.exchange_fee_address.Write(buf);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// Instruments ([]InstrumentTransfer)
		{
			if (this.instruments.length > (2 ** 8) - 1) {
				return sprintf('list field instruments has too many items %d/%d', this.instruments.length, (2 ** 8) - 1);
			}

			const err = this.instruments.find((value, i) => {
				const err2 = value.Validate();
				if (err2) return sprintf('list field instruments[%d] is invalid : %s', i, err2);
			});
			if (err) return err;
		}

		// OfferExpiry (Timestamp) 
		{
			// IsInternalType
			const err = this.offer_expiry.Validate();
			if  (err) return sprintf('field offer_expiry is invalid : %s', err);

		}

		// ExchangeFee (uint64) 
		{
		}

		// ExchangeFeeAddress (PublicKeyHash) 
		{
			// IsInternalType
			const err = this.exchange_fee_address.Validate();
			if  (err) return sprintf('field exchange_fee_address is invalid : %s', err);

		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('instruments:%s', this.instruments.toString()));
			vals.push(sprintf('offer_expiry:%s', this.offer_expiry.toString()));
			vals.push(sprintf('exchange_fee:%d', this.exchange_fee));
			vals.push(sprintf('exchange_fee_address:%s', this.exchange_fee_address.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for Transfer.', e);
			throw e;
		}
		// this.instruments (InstrumentTransfer[])
		
		this.offer_expiry = new Timestamp();
		this.offer_expiry.fromJSON(JSON.stringify(parsed.offerExpiry));
		
		this.exchange_fee = parsed.exchangeFee;
		
		this.exchange_fee_address = new PublicKeyHash();
		this.exchange_fee_address.fromJSON(JSON.stringify(parsed.exchangeFeeAddress));
		
		return this.Validate();
	}
}

	// Settlement Settles the transfer request of bitcoins and tokens from
// transfer (T1) actions.
	export class Settlement extends OpReturnMessage {
		type = ActionCode.CodeSettlement;
		typeStr = 'Settlement';

	
		// 	The Instruments settled by the transfer action.
		instruments;
	
		// 	Timestamp in nanoseconds of when the smart contract created the
	// action.
		timestamp;
	

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.CodeSettlement;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
		// instruments ([]InstrumentSettlement)
		{
			WriteVariableSize(buf, this.instruments.length, 8, 8);
			this.instruments.forEach((value) => {
				buf.write(value.Serialize());
			});
		}

		// timestamp (Timestamp)
		{
			buf.write(this.timestamp.Serialize());
		}

		return buf.buf;
	}

	// write populates the fields in Settlement from the byte slice
	write(b: Buffer): number {
		const buf = new _.Reader(b);
		// instruments ([]InstrumentSettlement)
		{
			// IsInternalTypeArray
			const size = ReadVariableSize(buf, 8, 8);
			this.instruments = [];
			for (let i = 0; i < size; i++) {
				const newValue = new InstrumentSettlement();
				newValue.Write(buf);

				this.instruments.push(newValue);
			}
		}

		// timestamp (Timestamp)
		{
			this.timestamp = new Timestamp();
			this.timestamp.Write(buf);
		}

		return b.length - buf.length;
	}

	Validate(): string | null {

		// Instruments ([]InstrumentSettlement)
		{
			if (this.instruments.length > (2 ** 8) - 1) {
				return sprintf('list field instruments has too many items %d/%d', this.instruments.length, (2 ** 8) - 1);
			}

			const err = this.instruments.find((value, i) => {
				const err2 = value.Validate();
				if (err2) return sprintf('list field instruments[%d] is invalid : %s', i, err2);
			});
			if (err) return err;
		}

		// Timestamp (Timestamp) 
		{
			// IsInternalType
			const err = this.timestamp.Validate();
			if  (err) return sprintf('field timestamp is invalid : %s', err);

		}
		return null; 
	}

	toString(): string {
		const vals: string[] = [];
			vals.push(sprintf('instruments:%s', this.instruments.toString()));
			vals.push(sprintf('timestamp:%s', this.timestamp.toString()));
	
		return sprintf('{%s}', vals.join(' '));
	}

	parse(json: string): string | null {
		let parsed: any;
		try {
			parsed = JSON.parse(json);
		} catch (e) {
			console.error('Failed to parse JSON for Settlement.', e);
			throw e;
		}
		// this.instruments (InstrumentSettlement[])
		
		this.timestamp = new Timestamp();
		this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
		
		return this.Validate();
	}
}
