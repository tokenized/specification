import {sprintf} from 'sprintf-js';
import _ from '@keyring/util';
import {write, read, ReadVarChar, ReadVariableSize, ReadVarBin,
	WriteVarChar, WriteVariableSize, WriteVarBin} from './bytes';
import { TxId, Timestamp, } from './protocol_types';
import { TargetAddress, OutputTag, Output, Document } from './field_types';
import { Resources } from './resources';

export enum MsgCodes {
	// CodePublicMessage identifies data as a PublicMessage message.
	PublicMessage = 2,

	// CodePrivateMessage identifies data as a PrivateMessage message.
	PrivateMessage = 3,

	// CodeRevertedTx identifies data as a RevertedTx message.
	RevertedTx = 4,

	// CodeOffer identifies data as a Offer message.
	Offer = 1001,

	// CodeSignatureRequest identifies data as a SignatureRequest message.
	SignatureRequest = 1002,

	// CodeSettlementRequest identifies data as a SettlementRequest message.
	SettlementRequest = 1003,

	// CodeOutputMetadata identifies data as a OutputMetadata message.
	OutputMetadata = 1004,

}

// MessagePayload is the interface for payloads within message actions.
export interface MessagePayload {
	Type(): number;
	Serialize(): Buffer;
	Write(Buffer): number;
	Validate(): string;
}

// MessageTypeMapping holds a mapping of message codes to message types.
export function MessageTypeMapping(code: number): MessagePayload {
	switch (code) {
	case MsgCodes.PublicMessage:
		return new PublicMessage();
	case MsgCodes.PrivateMessage:
		return new PrivateMessage();
	case MsgCodes.RevertedTx:
		return new RevertedTx();
	case MsgCodes.Offer:
		return new Offer();
	case MsgCodes.SignatureRequest:
		return new SignatureRequest();
	case MsgCodes.SettlementRequest:
		return new SettlementRequest();
	case MsgCodes.OutputMetadata:
		return new OutputMetadata();
	default:
		return null;
	}
}



// PublicMessage Generic public message or public announcement. Sent to an
// address(es). Can be used for an official issuer announcement.
export class PublicMessage {
// Payload Version
	version; // uint8 `json:"version,omitempty"`// Timestamp in nanoseconds for when the message sender creates the
	// transaction.
	timestamp = new Timestamp(); // Timestamp `json:"timestamp,omitempty"`// The subject / topic of the message.
	subject; // string `json:"subject,omitempty"`// The output of the message that this message is regarding (responding
	// to).
	regarding = new Output(); // Output `json:"regarding,omitempty"`// Tokenized Ltd. announces product launch.
	public_message = new Document(); // Document `json:"public_message,omitempty"`// Documents attached to the message.
	attachments = []; // []Document `json:"attachments,omitempty"`

	

	// Type returns the type identifer for this message.
	Type(): number {
		return MsgCodes.PublicMessage;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// Serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
	

		// Version (uint8)
		{
			write(buf, this.version, 'uint8');
		}
	

		// Timestamp (Timestamp)
		{
			buf.write(this.timestamp.Serialize());
		}
	

		// Subject (string)
		{
			WriteVarChar(buf, this.subject, 16);
		}
	

		// Regarding (Output)
		{
			buf.write(this.regarding.Serialize());
		}
	

		// PublicMessage (Document)
		{
			buf.write(this.public_message.Serialize());
		}
	

		// Attachments ([]Document)
		{
			WriteVariableSize(buf, this.attachments.length, 32, 8);
			this.attachments.forEach((value) => {
				buf.write(value.Serialize());
			});
		}
	
		return buf.buf;
	}

	// Write populates the fields in PublicMessage from the byte slice
	Write(b: Buffer): number {
		const buf = new _.Reader(b);

		// Version (uint8)
		{
			this.version = read(buf, 'uint8');
		}

	

		// Timestamp (Timestamp)
		{
			this.timestamp = new Timestamp();
			this.timestamp.Write(buf);
		}

	

		// Subject (string)
		{
			this.subject = ReadVarChar(buf, 16);
		}

	

		// Regarding (Output)
		{
			this.regarding = new Output();
			this.regarding.Write(buf);
		}

	

		// PublicMessage (Document)
		{
			this.public_message = new Document();
			this.public_message.Write(buf);
		}

	

		// Attachments ([]Document)
		{
			const size = ReadVariableSize(buf, 32, 8);
			this.attachments = [...Array(size)].map(() => {
				const v = new Document();
				v.Write(buf);
				return v;
			});
		}

	

		return 0;
	}

	Validate(): string {

		// Version (uint8)
		{
		}
	

		// Timestamp (Timestamp)
		{
			const err = this.timestamp.Validate();
			if (err) return sprintf('field timestamp is invalid : %s', err);
	
		}
	

		// Subject (string)
		{
			if (this.subject.length > (2 ** 16) - 1) {
				return sprintf('varchar field subject too long %d/%d', this.subject.length, (2 ** 16) - 1);
			}
		}
	

		// Regarding (Output)
		{
			const err = this.regarding.Validate();
			if (err) return sprintf('field regarding is invalid : %s', err);
	
		}
	

		// PublicMessage (Document)
		{
			const err = this.public_message.Validate();
			if (err) return sprintf('field public_message is invalid : %s', err);
	
		}
	

		// Attachments ([]Document)
		{
			if (this.attachments.length > (2 ** 32) - 1) {
				return sprintf('list field attachments has too many items %d/%d', this.attachments.length, (2 ** 32) - 1);
			}

			const e = this.attachments.find((value, i) => {
				const err = value.Validate();
				if (err) return sprintf('list field attachments[%d] is invalid : %s', i, err);
			});
			if (e) return e;
		}
	
		return null;
	}

	toString(): string {
		const vals: string[] = [];
	
		vals.push(sprintf('version:%v', this.version));
		vals.push(sprintf('timestamp:%#+v', this.timestamp));
		vals.push(sprintf('subject:%#+v', this.subject));
		vals.push(sprintf('regarding:%#+v', this.regarding));
		vals.push(sprintf('public_message:%#+v', this.public_message));
		vals.push(sprintf('attachments:%#+v', this.attachments));

		return sprintf('{%s}', vals.join(' '));
	}
}

// PrivateMessage Generic private message. Sent to another address(es).
// Encryption is to be used.
export class PrivateMessage {
// Payload Version
	version; // uint8 `json:"version,omitempty"`// Timestamp in nanoseconds for when the message sender creates the
	// transaction.
	timestamp = new Timestamp(); // Timestamp `json:"timestamp,omitempty"`// The subject / topic of the message.
	subject; // string `json:"subject,omitempty"`// The output of the message that this message is regarding (responding
	// to).
	regarding = new Output(); // Output `json:"regarding,omitempty"`// Tokenized Ltd announces product launch.
	private_message = new Document(); // Document `json:"private_message,omitempty"`// Documents attached to the message.
	attachments = []; // []Document `json:"attachments,omitempty"`

	

	// Type returns the type identifer for this message.
	Type(): number {
		return MsgCodes.PrivateMessage;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// Serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
	

		// Version (uint8)
		{
			write(buf, this.version, 'uint8');
		}
	

		// Timestamp (Timestamp)
		{
			buf.write(this.timestamp.Serialize());
		}
	

		// Subject (string)
		{
			WriteVarChar(buf, this.subject, 16);
		}
	

		// Regarding (Output)
		{
			buf.write(this.regarding.Serialize());
		}
	

		// PrivateMessage (Document)
		{
			buf.write(this.private_message.Serialize());
		}
	

		// Attachments ([]Document)
		{
			WriteVariableSize(buf, this.attachments.length, 32, 8);
			this.attachments.forEach((value) => {
				buf.write(value.Serialize());
			});
		}
	
		return buf.buf;
	}

	// Write populates the fields in PrivateMessage from the byte slice
	Write(b: Buffer): number {
		const buf = new _.Reader(b);

		// Version (uint8)
		{
			this.version = read(buf, 'uint8');
		}

	

		// Timestamp (Timestamp)
		{
			this.timestamp = new Timestamp();
			this.timestamp.Write(buf);
		}

	

		// Subject (string)
		{
			this.subject = ReadVarChar(buf, 16);
		}

	

		// Regarding (Output)
		{
			this.regarding = new Output();
			this.regarding.Write(buf);
		}

	

		// PrivateMessage (Document)
		{
			this.private_message = new Document();
			this.private_message.Write(buf);
		}

	

		// Attachments ([]Document)
		{
			const size = ReadVariableSize(buf, 32, 8);
			this.attachments = [...Array(size)].map(() => {
				const v = new Document();
				v.Write(buf);
				return v;
			});
		}

	

		return 0;
	}

	Validate(): string {

		// Version (uint8)
		{
		}
	

		// Timestamp (Timestamp)
		{
			const err = this.timestamp.Validate();
			if (err) return sprintf('field timestamp is invalid : %s', err);
	
		}
	

		// Subject (string)
		{
			if (this.subject.length > (2 ** 16) - 1) {
				return sprintf('varchar field subject too long %d/%d', this.subject.length, (2 ** 16) - 1);
			}
		}
	

		// Regarding (Output)
		{
			const err = this.regarding.Validate();
			if (err) return sprintf('field regarding is invalid : %s', err);
	
		}
	

		// PrivateMessage (Document)
		{
			const err = this.private_message.Validate();
			if (err) return sprintf('field private_message is invalid : %s', err);
	
		}
	

		// Attachments ([]Document)
		{
			if (this.attachments.length > (2 ** 32) - 1) {
				return sprintf('list field attachments has too many items %d/%d', this.attachments.length, (2 ** 32) - 1);
			}

			const e = this.attachments.find((value, i) => {
				const err = value.Validate();
				if (err) return sprintf('list field attachments[%d] is invalid : %s', i, err);
			});
			if (e) return e;
		}
	
		return null;
	}

	toString(): string {
		const vals: string[] = [];
	
		vals.push(sprintf('version:%v', this.version));
		vals.push(sprintf('timestamp:%#+v', this.timestamp));
		vals.push(sprintf('subject:%#+v', this.subject));
		vals.push(sprintf('regarding:%#+v', this.regarding));
		vals.push(sprintf('private_message:%#+v', this.private_message));
		vals.push(sprintf('attachments:%#+v', this.attachments));

		return sprintf('{%s}', vals.join(' '));
	}
}

// RevertedTx A message that contains a bitcoin transaction that was
// accepted by the network or an agent and then invalidated by a reorg, or
// zero conf double spend. This serves as on chain evidence of the sending
// party's signatures and approval for the given transaction.
export class RevertedTx {
// Payload Version
	version; // uint8 `json:"version,omitempty"`// Timestamp in nanoseconds for when the message sender creates the
	// transaction.
	timestamp = new Timestamp(); // Timestamp `json:"timestamp,omitempty"`// Serialized bitcoin transaction that was reverted/invalidated after
	// being accepted.
	transaction; // []byte `json:"transaction,omitempty"`

	

	// Type returns the type identifer for this message.
	Type(): number {
		return MsgCodes.RevertedTx;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// Serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
	

		// Version (uint8)
		{
			write(buf, this.version, 'uint8');
		}
	

		// Timestamp (Timestamp)
		{
			buf.write(this.timestamp.Serialize());
		}
	

		// Transaction ([]byte)
		{
			WriteVarBin(buf, this.transaction, 32);
		}
	
		return buf.buf;
	}

	// Write populates the fields in RevertedTx from the byte slice
	Write(b: Buffer): number {
		const buf = new _.Reader(b);

		// Version (uint8)
		{
			this.version = read(buf, 'uint8');
		}

	

		// Timestamp (Timestamp)
		{
			this.timestamp = new Timestamp();
			this.timestamp.Write(buf);
		}

	

		// Transaction ([]byte)
		{
			this.transaction = ReadVarBin(buf, 32);
		}

	

		return 0;
	}

	Validate(): string {

		// Version (uint8)
		{
		}
	

		// Timestamp (Timestamp)
		{
			const err = this.timestamp.Validate();
			if (err) return sprintf('field timestamp is invalid : %s', err);
	
		}
	

		// Transaction ([]byte)
		{
			if (this.transaction.length > (2 ** 32) - 1) {
				return sprintf('varbin field transaction too long %d/%d', this.transaction.length, (2 ** 32) - 1);
			}
		}
	
		return null;
	}

	toString(): string {
		const vals: string[] = [];
	
		vals.push(sprintf('version:%v', this.version));
		vals.push(sprintf('timestamp:%#+v', this.timestamp));
		vals.push(sprintf('transaction:%#x', this.transaction));

		return sprintf('{%s}', vals.join(' '));
	}
}

// Offer A message that contains all of the details required for an
// agreement to be formed. Sent to an address(es). The Offer should have
// all, or nearly all, of the details required for the receiving party to
// accept the offer. The Offer shall be in the form of a partially formed
// Bitcoin transaction with all of the relevent details (offer,
// consideration, offeror's payment/receipt details, etc.). The Offer
// message is different to a Signature Request message in that it is
// missing the offeree's payment/receipt details (eg. UTXOs). If the Offer
// message is well received by the offeree, then the offeree can add their
// relevent details (eg. inputs/outputs) and sign the transaction. If an
// additional signature is required from the offeror at this point, then
// the partially-signed transaction can be sent to the offeror by way of a
// Signature Request message.
export class Offer {
// Payload Version
	version; // uint8 `json:"version,omitempty"`// Timestamp in nanoseconds for when the message sender created the offer.
	timestamp = new Timestamp(); // Timestamp `json:"timestamp,omitempty"`// Serialized Tokenized OP_RETURN message. The message needs data added by
	// another party upon acceptance of offer.
	payload; // []byte `json:"payload,omitempty"`

	

	// Type returns the type identifer for this message.
	Type(): number {
		return MsgCodes.Offer;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// Serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
	

		// Version (uint8)
		{
			write(buf, this.version, 'uint8');
		}
	

		// Timestamp (Timestamp)
		{
			buf.write(this.timestamp.Serialize());
		}
	

		// Payload ([]byte)
		{
			WriteVarBin(buf, this.payload, 32);
		}
	
		return buf.buf;
	}

	// Write populates the fields in Offer from the byte slice
	Write(b: Buffer): number {
		const buf = new _.Reader(b);

		// Version (uint8)
		{
			this.version = read(buf, 'uint8');
		}

	

		// Timestamp (Timestamp)
		{
			this.timestamp = new Timestamp();
			this.timestamp.Write(buf);
		}

	

		// Payload ([]byte)
		{
			this.payload = ReadVarBin(buf, 32);
		}

	

		return 0;
	}

	Validate(): string {

		// Version (uint8)
		{
		}
	

		// Timestamp (Timestamp)
		{
			const err = this.timestamp.Validate();
			if (err) return sprintf('field timestamp is invalid : %s', err);
	
		}
	

		// Payload ([]byte)
		{
			if (this.payload.length > (2 ** 32) - 1) {
				return sprintf('varbin field payload too long %d/%d', this.payload.length, (2 ** 32) - 1);
			}
		}
	
		return null;
	}

	toString(): string {
		const vals: string[] = [];
	
		vals.push(sprintf('version:%v', this.version));
		vals.push(sprintf('timestamp:%#+v', this.timestamp));
		vals.push(sprintf('payload:%#x', this.payload));

		return sprintf('{%s}', vals.join(' '));
	}
}

// SignatureRequest Partially-signed transactions (Tokenized actions,
// Bitcoin, Multisig, Threshold Signatures, etc.) can be passed around
// on-chain to the parties (including Smart Contracts) that still have to
// sign the transaction.
export class SignatureRequest {
// Payload Version
	version; // uint8 `json:"version,omitempty"`// Timestamp in nanoseconds for when the message sender creates the
	// transaction.
	timestamp = new Timestamp(); // Timestamp `json:"timestamp,omitempty"`// Full serialized bitcoin tx with multiple inputs from different
	// wallets/users.
	payload; // []byte `json:"payload,omitempty"`

	

	// Type returns the type identifer for this message.
	Type(): number {
		return MsgCodes.SignatureRequest;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// Serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
	

		// Version (uint8)
		{
			write(buf, this.version, 'uint8');
		}
	

		// Timestamp (Timestamp)
		{
			buf.write(this.timestamp.Serialize());
		}
	

		// Payload ([]byte)
		{
			WriteVarBin(buf, this.payload, 32);
		}
	
		return buf.buf;
	}

	// Write populates the fields in SignatureRequest from the byte slice
	Write(b: Buffer): number {
		const buf = new _.Reader(b);

		// Version (uint8)
		{
			this.version = read(buf, 'uint8');
		}

	

		// Timestamp (Timestamp)
		{
			this.timestamp = new Timestamp();
			this.timestamp.Write(buf);
		}

	

		// Payload ([]byte)
		{
			this.payload = ReadVarBin(buf, 32);
		}

	

		return 0;
	}

	Validate(): string {

		// Version (uint8)
		{
		}
	

		// Timestamp (Timestamp)
		{
			const err = this.timestamp.Validate();
			if (err) return sprintf('field timestamp is invalid : %s', err);
	
		}
	

		// Payload ([]byte)
		{
			if (this.payload.length > (2 ** 32) - 1) {
				return sprintf('varbin field payload too long %d/%d', this.payload.length, (2 ** 32) - 1);
			}
		}
	
		return null;
	}

	toString(): string {
		const vals: string[] = [];
	
		vals.push(sprintf('version:%v', this.version));
		vals.push(sprintf('timestamp:%#+v', this.timestamp));
		vals.push(sprintf('payload:%#x', this.payload));

		return sprintf('{%s}', vals.join(' '));
	}
}

// SettlementRequest A message that contains a multi-contract settlement
// that needs settlement data added by another contract. Sent to another
// contract to request data be added.
export class SettlementRequest {
// Payload Version
	version; // uint8 `json:"version,omitempty"`// Timestamp in nanoseconds for when the message sender creates the
	// transaction.
	timestamp = new Timestamp(); // Timestamp `json:"timestamp,omitempty"`// Tx Id of the transfer request transaction that triggered this message.
	transfer_tx_id = new TxId(); // TxId `json:"transfer_tx_id,omitempty"`// Contract fees (in bitcoin) and addresses(PKHs) where fees should be
	// paid. Added by each contract as settlement data is added.
	contract_fees = []; // []TargetAddress `json:"contract_fees,omitempty"`// Serialized settlement OP_RETURN that needs data added by another
	// contract.
	settlement; // []byte `json:"settlement,omitempty"`

	

	// Type returns the type identifer for this message.
	Type(): number {
		return MsgCodes.SettlementRequest;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// Serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
	

		// Version (uint8)
		{
			write(buf, this.version, 'uint8');
		}
	

		// Timestamp (Timestamp)
		{
			buf.write(this.timestamp.Serialize());
		}
	

		// TransferTxId (TxId)
		{
			buf.write(this.transfer_tx_id.Serialize());
		}
	

		// ContractFees ([]TargetAddress)
		{
			WriteVariableSize(buf, this.contract_fees.length, 8, 8);
			this.contract_fees.forEach((value) => {
				buf.write(value.Serialize());
			});
		}
	

		// Settlement ([]byte)
		{
			WriteVarBin(buf, this.settlement, 32);
		}
	
		return buf.buf;
	}

	// Write populates the fields in SettlementRequest from the byte slice
	Write(b: Buffer): number {
		const buf = new _.Reader(b);

		// Version (uint8)
		{
			this.version = read(buf, 'uint8');
		}

	

		// Timestamp (Timestamp)
		{
			this.timestamp = new Timestamp();
			this.timestamp.Write(buf);
		}

	

		// TransferTxId (TxId)
		{
			this.transfer_tx_id = new TxId();
			this.transfer_tx_id.Write(buf);
		}

	

		// ContractFees ([]TargetAddress)
		{
			const size = ReadVariableSize(buf, 8, 8);
			this.contract_fees = [...Array(size)].map(() => {
				const v = new TargetAddress();
				v.Write(buf);
				return v;
			});
		}

	

		// Settlement ([]byte)
		{
			this.settlement = ReadVarBin(buf, 32);
		}

	

		return 0;
	}

	Validate(): string {

		// Version (uint8)
		{
		}
	

		// Timestamp (Timestamp)
		{
			const err = this.timestamp.Validate();
			if (err) return sprintf('field timestamp is invalid : %s', err);
	
		}
	

		// TransferTxId (TxId)
		{
			const err = this.transfer_tx_id.Validate();
			if (err) return sprintf('field transfer_tx_id is invalid : %s', err);
	
		}
	

		// ContractFees ([]TargetAddress)
		{
			if (this.contract_fees.length > (2 ** 8) - 1) {
				return sprintf('list field contract_fees has too many items %d/%d', this.contract_fees.length, (2 ** 8) - 1);
			}

			const e = this.contract_fees.find((value, i) => {
				const err = value.Validate();
				if (err) return sprintf('list field contract_fees[%d] is invalid : %s', i, err);
			});
			if (e) return e;
		}
	

		// Settlement ([]byte)
		{
			if (this.settlement.length > (2 ** 32) - 1) {
				return sprintf('varbin field settlement too long %d/%d', this.settlement.length, (2 ** 32) - 1);
			}
		}
	
		return null;
	}

	toString(): string {
		const vals: string[] = [];
	
		vals.push(sprintf('version:%v', this.version));
		vals.push(sprintf('timestamp:%#+v', this.timestamp));
		vals.push(sprintf('transfer_tx_id:%#+v', this.transfer_tx_id));
		vals.push(sprintf('contract_fees:%#+v', this.contract_fees));
		vals.push(sprintf('settlement:%#x', this.settlement));

		return sprintf('{%s}', vals.join(' '));
	}
}

// OutputMetadata Metadata associated with the output. Aka Transaction
// details. It is used to describe the purpose of the transaction and add
// other relevant information. Often encrypted (DH, RSA) to make it private
// for one or more parties. DH for b2b where multiple parties can see the
// description. RSA or the like for descriptions only visible to one of the
// transacting parties. Optional
export class OutputMetadata {
// Payload Version
	version; // uint8 `json:"version,omitempty"`// A Description that accompanies the output. A transaction description.
	output_description; // string `json:"output_description,omitempty"`// Predefined values for describing the output.
	tags; // []uint8 `json:"tags,omitempty"`// Free form text fields for describing the output. Groceries, Moomba Gas
	// Compressor Project, Cash Register 3, Fitness, Entertainment, Special,
	// VIP Section, North Carolina Store, Waitress: Cindy Smith, etc.
	custom_tags = []; // []OutputTag `json:"custom_tags,omitempty"`

	

	// Type returns the type identifer for this message.
	Type(): number {
		return MsgCodes.OutputMetadata;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// Serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
	

		// Version (uint8)
		{
			write(buf, this.version, 'uint8');
		}
	

		// OutputDescription (string)
		{
			WriteVarChar(buf, this.output_description, 32);
		}
	

		// Tags ([]uint8)
		{
			// IsNativeTypeArray []uint8
			const type = '[]uint8'.slice(2);
			WriteVariableSize(buf, this.tags.length, 8, 8);
			this.tags.forEach(value => {
				write(buf, value, type); // TODO
			});
		}
	

		// CustomTags ([]OutputTag)
		{
			WriteVariableSize(buf, this.custom_tags.length, 8, 8);
			this.custom_tags.forEach((value) => {
				buf.write(value.Serialize());
			});
		}
	
		return buf.buf;
	}

	// Write populates the fields in OutputMetadata from the byte slice
	Write(b: Buffer): number {
		const buf = new _.Reader(b);

		// Version (uint8)
		{
			this.version = read(buf, 'uint8');
		}

	

		// OutputDescription (string)
		{
			this.output_description = ReadVarChar(buf, 32);
		}

	

		// Tags ([]uint8)
		{
			const size = ReadVariableSize(buf, 8, 8);
			const type = '[]uint8'.slice(2);
			this.tags = [...Array(size)].map(() => read(buf, type));
		}

	

		// CustomTags ([]OutputTag)
		{
			const size = ReadVariableSize(buf, 8, 8);
			this.custom_tags = [...Array(size)].map(() => {
				const v = new OutputTag();
				v.Write(buf);
				return v;
			});
		}

	

		return 0;
	}

	Validate(): string {

		// Version (uint8)
		{
		}
	

		// OutputDescription (string)
		{
			if (this.output_description.length > (2 ** 32) - 1) {
				return sprintf('varchar field output_description too long %d/%d', this.output_description.length, (2 ** 32) - 1);
			}
		}
	

		// Tags ([]uint8)
		{
			if (this.tags.length > (2 ** 8) - 1) {
				return sprintf('list field tags has too many items %d/%d', this.tags.length, (2 ** 8) - 1);
			}

			let err = null;
			this.tags.find(value => {
				if (!Resources.GetTagType(value)) {
					err = sprintf('tags: Invalid tag type value : \'%c\'/%d', value, value);
					return true;
				}
			});
			if (err) return err;
		}
	

		// CustomTags ([]OutputTag)
		{
			if (this.custom_tags.length > (2 ** 8) - 1) {
				return sprintf('list field custom_tags has too many items %d/%d', this.custom_tags.length, (2 ** 8) - 1);
			}

			const e = this.custom_tags.find((value, i) => {
				const err = value.Validate();
				if (err) return sprintf('list field custom_tags[%d] is invalid : %s', i, err);
			});
			if (e) return e;
		}
	
		return null;
	}

	toString(): string {
		const vals: string[] = [];
	
		vals.push(sprintf('version:%v', this.version));
		vals.push(sprintf('output_description:%#+v', this.output_description));
		vals.push(sprintf('tags:%#+v', this.tags));
		vals.push(sprintf('custom_tags:%#+v', this.custom_tags));

		return sprintf('{%s}', vals.join(' '));
	}
}


