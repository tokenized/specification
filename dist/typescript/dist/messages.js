"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const sprintf_js_1 = require("sprintf-js");
const util_1 = __importDefault(require("@keyring/util"));
const bytes_1 = require("./bytes");
const protocol_types_1 = require("./protocol_types");
const field_types_1 = require("./field_types");
const resources_1 = require("./resources");
var MsgCodes;
(function (MsgCodes) {
    // CodePublicMessage identifies data as a PublicMessage message.
    MsgCodes[MsgCodes["PublicMessage"] = 2] = "PublicMessage";
    // CodePrivateMessage identifies data as a PrivateMessage message.
    MsgCodes[MsgCodes["PrivateMessage"] = 3] = "PrivateMessage";
    // CodeRevertedTx identifies data as a RevertedTx message.
    MsgCodes[MsgCodes["RevertedTx"] = 4] = "RevertedTx";
    // CodeOffer identifies data as a Offer message.
    MsgCodes[MsgCodes["Offer"] = 1001] = "Offer";
    // CodeSignatureRequest identifies data as a SignatureRequest message.
    MsgCodes[MsgCodes["SignatureRequest"] = 1002] = "SignatureRequest";
    // CodeSettlementRequest identifies data as a SettlementRequest message.
    MsgCodes[MsgCodes["SettlementRequest"] = 1003] = "SettlementRequest";
    // CodeOutputMetadata identifies data as a OutputMetadata message.
    MsgCodes[MsgCodes["OutputMetadata"] = 1004] = "OutputMetadata";
})(MsgCodes = exports.MsgCodes || (exports.MsgCodes = {}));
// MessageTypeMapping holds a mapping of message codes to message types.
function MessageTypeMapping(code) {
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
exports.MessageTypeMapping = MessageTypeMapping;
// PublicMessage Generic public message or public announcement. Sent to an
// address(es). Can be used for an official issuer announcement.
class PublicMessage {
    constructor() {
        // transaction.
        this.timestamp = new protocol_types_1.Timestamp(); // Timestamp `json:"timestamp,omitempty"`// Tokenized Ltd. announces product launch.
    }
    // Type returns the type identifer for this message.
    Type() {
        return MsgCodes.PublicMessage;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // Serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Version (uint8)
        {
            bytes_1.write(buf, this.version, 'uint8');
        }
        // Timestamp (Timestamp)
        {
            buf.write(this.timestamp.Serialize());
        }
        // PublicMessage (string)
        {
            bytes_1.WriteVarChar(buf, this.public_message, 32);
        }
        return buf.buf;
    }
    // Write populates the fields in PublicMessage from the byte slice
    Write(b) {
        const buf = new util_1.default.Reader(b);
        // Version (uint8)
        {
            this.version = bytes_1.read(buf, 'uint8');
        }
        // Timestamp (Timestamp)
        {
            this.timestamp = new protocol_types_1.Timestamp();
            this.timestamp.Write(buf);
        }
        // PublicMessage (string)
        {
            this.public_message = bytes_1.ReadVarChar(buf, 32);
        }
        return 0;
    }
    Validate() {
        // Version (uint8)
        {
        }
        // Timestamp (Timestamp)
        {
            const err = this.timestamp.Validate();
            if (err)
                return sprintf_js_1.sprintf('field timestamp is invalid : %s', err);
        }
        // PublicMessage (string)
        {
            if (this.public_message.length > (Math.pow(2, 32)) - 1) {
                return sprintf_js_1.sprintf('varchar field public_message too long %d/%d', this.public_message.length, (Math.pow(2, 32)) - 1);
            }
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('version:%v', this.version));
        vals.push(sprintf_js_1.sprintf('timestamp:%#+v', this.timestamp));
        vals.push(sprintf_js_1.sprintf('public_message:%#+v', this.public_message));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
}
exports.PublicMessage = PublicMessage;
// PrivateMessage Generic private message. Sent to another address(es).
// Encryption is to be used.
class PrivateMessage {
    constructor() {
        // transaction.
        this.timestamp = new protocol_types_1.Timestamp(); // Timestamp `json:"timestamp,omitempty"`// Tokenized Ltd announces product launch.
    }
    // Type returns the type identifer for this message.
    Type() {
        return MsgCodes.PrivateMessage;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // Serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Version (uint8)
        {
            bytes_1.write(buf, this.version, 'uint8');
        }
        // Timestamp (Timestamp)
        {
            buf.write(this.timestamp.Serialize());
        }
        // PrivateMessage ([]byte)
        {
            bytes_1.WriteVarBin(buf, this.private_message, 32);
        }
        return buf.buf;
    }
    // Write populates the fields in PrivateMessage from the byte slice
    Write(b) {
        const buf = new util_1.default.Reader(b);
        // Version (uint8)
        {
            this.version = bytes_1.read(buf, 'uint8');
        }
        // Timestamp (Timestamp)
        {
            this.timestamp = new protocol_types_1.Timestamp();
            this.timestamp.Write(buf);
        }
        // PrivateMessage ([]byte)
        {
            this.private_message = bytes_1.ReadVarBin(buf, 32);
        }
        return 0;
    }
    Validate() {
        // Version (uint8)
        {
        }
        // Timestamp (Timestamp)
        {
            const err = this.timestamp.Validate();
            if (err)
                return sprintf_js_1.sprintf('field timestamp is invalid : %s', err);
        }
        // PrivateMessage ([]byte)
        {
            if (this.private_message.length > (Math.pow(2, 32)) - 1) {
                return sprintf_js_1.sprintf('varbin field private_message too long %d/%d', this.private_message.length, (Math.pow(2, 32)) - 1);
            }
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('version:%v', this.version));
        vals.push(sprintf_js_1.sprintf('timestamp:%#+v', this.timestamp));
        vals.push(sprintf_js_1.sprintf('private_message:%#x', this.private_message));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
}
exports.PrivateMessage = PrivateMessage;
// RevertedTx A message that contains a bitcoin transaction that was
// accepted by the network or an agent and then invalidated by a reorg, or
// zero conf double spend. This serves as on chain evidence of the sending
// party's signatures and approval for the given transaction.
class RevertedTx {
    constructor() {
        // transaction.
        this.timestamp = new protocol_types_1.Timestamp(); // Timestamp `json:"timestamp,omitempty"`// Serialized bitcoin transaction that was reverted/invalidated after
    }
    // Type returns the type identifer for this message.
    Type() {
        return MsgCodes.RevertedTx;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // Serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Version (uint8)
        {
            bytes_1.write(buf, this.version, 'uint8');
        }
        // Timestamp (Timestamp)
        {
            buf.write(this.timestamp.Serialize());
        }
        // Transaction ([]byte)
        {
            bytes_1.WriteVarBin(buf, this.transaction, 32);
        }
        return buf.buf;
    }
    // Write populates the fields in RevertedTx from the byte slice
    Write(b) {
        const buf = new util_1.default.Reader(b);
        // Version (uint8)
        {
            this.version = bytes_1.read(buf, 'uint8');
        }
        // Timestamp (Timestamp)
        {
            this.timestamp = new protocol_types_1.Timestamp();
            this.timestamp.Write(buf);
        }
        // Transaction ([]byte)
        {
            this.transaction = bytes_1.ReadVarBin(buf, 32);
        }
        return 0;
    }
    Validate() {
        // Version (uint8)
        {
        }
        // Timestamp (Timestamp)
        {
            const err = this.timestamp.Validate();
            if (err)
                return sprintf_js_1.sprintf('field timestamp is invalid : %s', err);
        }
        // Transaction ([]byte)
        {
            if (this.transaction.length > (Math.pow(2, 32)) - 1) {
                return sprintf_js_1.sprintf('varbin field transaction too long %d/%d', this.transaction.length, (Math.pow(2, 32)) - 1);
            }
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('version:%v', this.version));
        vals.push(sprintf_js_1.sprintf('timestamp:%#+v', this.timestamp));
        vals.push(sprintf_js_1.sprintf('transaction:%#x', this.transaction));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
}
exports.RevertedTx = RevertedTx;
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
class Offer {
    constructor() {
        this.timestamp = new protocol_types_1.Timestamp(); // Timestamp `json:"timestamp,omitempty"`// Serialized Tokenized OP_RETURN message. The message needs data added by
    }
    // Type returns the type identifer for this message.
    Type() {
        return MsgCodes.Offer;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // Serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Version (uint8)
        {
            bytes_1.write(buf, this.version, 'uint8');
        }
        // Timestamp (Timestamp)
        {
            buf.write(this.timestamp.Serialize());
        }
        // Payload ([]byte)
        {
            bytes_1.WriteVarBin(buf, this.payload, 32);
        }
        return buf.buf;
    }
    // Write populates the fields in Offer from the byte slice
    Write(b) {
        const buf = new util_1.default.Reader(b);
        // Version (uint8)
        {
            this.version = bytes_1.read(buf, 'uint8');
        }
        // Timestamp (Timestamp)
        {
            this.timestamp = new protocol_types_1.Timestamp();
            this.timestamp.Write(buf);
        }
        // Payload ([]byte)
        {
            this.payload = bytes_1.ReadVarBin(buf, 32);
        }
        return 0;
    }
    Validate() {
        // Version (uint8)
        {
        }
        // Timestamp (Timestamp)
        {
            const err = this.timestamp.Validate();
            if (err)
                return sprintf_js_1.sprintf('field timestamp is invalid : %s', err);
        }
        // Payload ([]byte)
        {
            if (this.payload.length > (Math.pow(2, 32)) - 1) {
                return sprintf_js_1.sprintf('varbin field payload too long %d/%d', this.payload.length, (Math.pow(2, 32)) - 1);
            }
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('version:%v', this.version));
        vals.push(sprintf_js_1.sprintf('timestamp:%#+v', this.timestamp));
        vals.push(sprintf_js_1.sprintf('payload:%#x', this.payload));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
}
exports.Offer = Offer;
// SignatureRequest Partially-signed transactions (Tokenized actions,
// Bitcoin, Multisig, Threshold Signatures, etc.) can be passed around
// on-chain to the parties (including Smart Contracts) that still have to
// sign the transaction.
class SignatureRequest {
    constructor() {
        // transaction.
        this.timestamp = new protocol_types_1.Timestamp(); // Timestamp `json:"timestamp,omitempty"`// Full serialized bitcoin tx with multiple inputs from different
    }
    // Type returns the type identifer for this message.
    Type() {
        return MsgCodes.SignatureRequest;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // Serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Version (uint8)
        {
            bytes_1.write(buf, this.version, 'uint8');
        }
        // Timestamp (Timestamp)
        {
            buf.write(this.timestamp.Serialize());
        }
        // Payload ([]byte)
        {
            bytes_1.WriteVarBin(buf, this.payload, 32);
        }
        return buf.buf;
    }
    // Write populates the fields in SignatureRequest from the byte slice
    Write(b) {
        const buf = new util_1.default.Reader(b);
        // Version (uint8)
        {
            this.version = bytes_1.read(buf, 'uint8');
        }
        // Timestamp (Timestamp)
        {
            this.timestamp = new protocol_types_1.Timestamp();
            this.timestamp.Write(buf);
        }
        // Payload ([]byte)
        {
            this.payload = bytes_1.ReadVarBin(buf, 32);
        }
        return 0;
    }
    Validate() {
        // Version (uint8)
        {
        }
        // Timestamp (Timestamp)
        {
            const err = this.timestamp.Validate();
            if (err)
                return sprintf_js_1.sprintf('field timestamp is invalid : %s', err);
        }
        // Payload ([]byte)
        {
            if (this.payload.length > (Math.pow(2, 32)) - 1) {
                return sprintf_js_1.sprintf('varbin field payload too long %d/%d', this.payload.length, (Math.pow(2, 32)) - 1);
            }
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('version:%v', this.version));
        vals.push(sprintf_js_1.sprintf('timestamp:%#+v', this.timestamp));
        vals.push(sprintf_js_1.sprintf('payload:%#x', this.payload));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
}
exports.SignatureRequest = SignatureRequest;
// SettlementRequest A message that contains a multi-contract settlement
// that needs settlement data added by another contract. Sent to another
// contract to request data be added.
class SettlementRequest {
    constructor() {
        // transaction.
        this.timestamp = new protocol_types_1.Timestamp(); // Timestamp `json:"timestamp,omitempty"`// Tx Id of the transfer request transaction that triggered this message.
        this.transfer_tx_id = new protocol_types_1.TxId(); // TxId `json:"transfer_tx_id,omitempty"`// Contract fees (in bitcoin) and addresses(PKHs) where fees should be
        // paid. Added by each contract as settlement data is added.
        this.contract_fees = []; // []TargetAddress `json:"contract_fees,omitempty"`// Serialized settlement OP_RETURN that needs data added by another
    }
    // Type returns the type identifer for this message.
    Type() {
        return MsgCodes.SettlementRequest;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // Serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Version (uint8)
        {
            bytes_1.write(buf, this.version, 'uint8');
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
            bytes_1.WriteVariableSize(buf, this.contract_fees.length, 8, 8);
            this.contract_fees.forEach((value) => {
                buf.write(value.Serialize());
            });
        }
        // Settlement ([]byte)
        {
            bytes_1.WriteVarBin(buf, this.settlement, 32);
        }
        return buf.buf;
    }
    // Write populates the fields in SettlementRequest from the byte slice
    Write(b) {
        const buf = new util_1.default.Reader(b);
        // Version (uint8)
        {
            this.version = bytes_1.read(buf, 'uint8');
        }
        // Timestamp (Timestamp)
        {
            this.timestamp = new protocol_types_1.Timestamp();
            this.timestamp.Write(buf);
        }
        // TransferTxId (TxId)
        {
            this.transfer_tx_id = new protocol_types_1.TxId();
            this.transfer_tx_id.Write(buf);
        }
        // ContractFees ([]TargetAddress)
        {
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            this.contract_fees = [...Array(size)].map(() => {
                const v = new field_types_1.TargetAddress();
                v.Write(buf);
                return v;
            });
        }
        // Settlement ([]byte)
        {
            this.settlement = bytes_1.ReadVarBin(buf, 32);
        }
        return 0;
    }
    Validate() {
        // Version (uint8)
        {
        }
        // Timestamp (Timestamp)
        {
            const err = this.timestamp.Validate();
            if (err)
                return sprintf_js_1.sprintf('field timestamp is invalid : %s', err);
        }
        // TransferTxId (TxId)
        {
            const err = this.transfer_tx_id.Validate();
            if (err)
                return sprintf_js_1.sprintf('field transfer_tx_id is invalid : %s', err);
        }
        // ContractFees ([]TargetAddress)
        {
            if (this.contract_fees.length > (Math.pow(2, 8)) - 1) {
                return sprintf_js_1.sprintf('list field contract_fees has too many items %d/%d', this.contract_fees.length, (Math.pow(2, 8)) - 1);
            }
            const e = this.contract_fees.find((value, i) => {
                const err = value.Validate();
                if (err)
                    return sprintf_js_1.sprintf('list field contract_fees[%d] is invalid : %s', i, err);
            });
            if (e)
                return e;
        }
        // Settlement ([]byte)
        {
            if (this.settlement.length > (Math.pow(2, 32)) - 1) {
                return sprintf_js_1.sprintf('varbin field settlement too long %d/%d', this.settlement.length, (Math.pow(2, 32)) - 1);
            }
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('version:%v', this.version));
        vals.push(sprintf_js_1.sprintf('timestamp:%#+v', this.timestamp));
        vals.push(sprintf_js_1.sprintf('transfer_tx_id:%#+v', this.transfer_tx_id));
        vals.push(sprintf_js_1.sprintf('contract_fees:%#+v', this.contract_fees));
        vals.push(sprintf_js_1.sprintf('settlement:%#x', this.settlement));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
}
exports.SettlementRequest = SettlementRequest;
// OutputMetadata Metadata associated with the output. Aka Transaction
// details. It is used to describe the purpose of the transaction and add
// other relevant information. Often encrypted (DH, RSA) to make it private
// for one or more parties. DH for b2b where multiple parties can see the
// description. RSA or the like for descriptions only visible to one of the
// transacting parties. Optional
class OutputMetadata {
    constructor() {
        // Compressor Project, Cash Register 3, Fitness, Entertainment, Special,
        // VIP Section, North Carolina Store, Waitress: Cindy Smith, etc.
        this.custom_tags = []; // []OutputTag `json:"custom_tags,omitempty"`
    }
    // Type returns the type identifer for this message.
    Type() {
        return MsgCodes.OutputMetadata;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // Serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Version (uint8)
        {
            bytes_1.write(buf, this.version, 'uint8');
        }
        // OutputDescription (string)
        {
            bytes_1.WriteVarChar(buf, this.output_description, 32);
        }
        // Tags ([]uint8)
        {
            // IsNativeTypeArray []uint8
            const type = '[]uint8'.slice(2);
            bytes_1.WriteVariableSize(buf, this.tags.length, 8, 8);
            this.tags.forEach(value => {
                bytes_1.write(buf, value, type); // TODO
            });
        }
        // CustomTags ([]OutputTag)
        {
            bytes_1.WriteVariableSize(buf, this.custom_tags.length, 8, 8);
            this.custom_tags.forEach((value) => {
                buf.write(value.Serialize());
            });
        }
        return buf.buf;
    }
    // Write populates the fields in OutputMetadata from the byte slice
    Write(b) {
        const buf = new util_1.default.Reader(b);
        // Version (uint8)
        {
            this.version = bytes_1.read(buf, 'uint8');
        }
        // OutputDescription (string)
        {
            this.output_description = bytes_1.ReadVarChar(buf, 32);
        }
        // Tags ([]uint8)
        {
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            const type = '[]uint8'.slice(2);
            this.tags = [...Array(size)].map(() => bytes_1.read(buf, type));
        }
        // CustomTags ([]OutputTag)
        {
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            this.custom_tags = [...Array(size)].map(() => {
                const v = new field_types_1.OutputTag();
                v.Write(buf);
                return v;
            });
        }
        return 0;
    }
    Validate() {
        // Version (uint8)
        {
        }
        // OutputDescription (string)
        {
            if (this.output_description.length > (Math.pow(2, 32)) - 1) {
                return sprintf_js_1.sprintf('varchar field output_description too long %d/%d', this.output_description.length, (Math.pow(2, 32)) - 1);
            }
        }
        // Tags ([]uint8)
        {
            if (this.tags.length > (Math.pow(2, 8)) - 1) {
                return sprintf_js_1.sprintf('list field tags has too many items %d/%d', this.tags.length, (Math.pow(2, 8)) - 1);
            }
            let err = null;
            this.tags.find(value => {
                if (!resources_1.Resources.GetTagType(value)) {
                    err = sprintf_js_1.sprintf('tags: Invalid tag type value : \'%c\'/%d', value, value);
                    return true;
                }
            });
            if (err)
                return err;
        }
        // CustomTags ([]OutputTag)
        {
            if (this.custom_tags.length > (Math.pow(2, 8)) - 1) {
                return sprintf_js_1.sprintf('list field custom_tags has too many items %d/%d', this.custom_tags.length, (Math.pow(2, 8)) - 1);
            }
            const e = this.custom_tags.find((value, i) => {
                const err = value.Validate();
                if (err)
                    return sprintf_js_1.sprintf('list field custom_tags[%d] is invalid : %s', i, err);
            });
            if (e)
                return e;
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('version:%v', this.version));
        vals.push(sprintf_js_1.sprintf('output_description:%#+v', this.output_description));
        vals.push(sprintf_js_1.sprintf('tags:%#+v', this.tags));
        vals.push(sprintf_js_1.sprintf('custom_tags:%#+v', this.custom_tags));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
}
exports.OutputMetadata = OutputMetadata;
