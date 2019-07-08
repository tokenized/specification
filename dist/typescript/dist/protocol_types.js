"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const sprintf_js_1 = require("sprintf-js");
const util_1 = __importDefault(require("@keyring/util"));
const bn_js_1 = __importDefault(require("bn.js"));
const bytes_1 = require("./bytes");
// ------------------------------------------------------------------------------------------------
// TxId represents a Bitcoin transaction ID. (Double SHA256 of tx data)
class TxId {
    constructor() {
        this.data = Buffer.alloc(32, 0);
    }
    static fromBytes(data) {
        const txid = new TxId();
        txid.data = Buffer.from(data);
        return txid;
    }
    // Validate returns an error if the value is invalid
    Validate() {
        return null;
    }
    // IsZero returns true if the tx id is all zeros.
    IsZero() {
        return this.data.equals(zeroTxId.data);
    }
    // Equal returns true if the specified values are the same.
    Equal(other) {
        return this.data.equals(other.data);
    }
    // Bytes returns the byte slice for the TxId.
    Bytes() {
        return this.data;
    }
    // String converts to a string
    toString() {
        return sprintf_js_1.sprintf('%x', this.data);
    }
    // Serialize returns a byte slice with the TxId in it.
    Serialize() {
        return this.data;
    }
    // Write reads a TxId from a bytes.Buffer
    Write(buf) {
        return this.data = buf.read(32);
    }
    // MarshalJSON converts to json.
    MarshalJSON() {
        return `"${this.data.toString('hex')}"`;
    }
    // UnmarshalJSON converts from json.
    UnmarshalJSON(data) {
        this.data = Buffer.from(data, 'hex');
    }
    // Set sets the value specified
    Set(value) {
        this.data = Buffer.from(value);
    }
}
exports.TxId = TxId;
const zeroTxId = new TxId();
class AssetCode {
    constructor() {
        this.data = Buffer.alloc(32, 0);
    }
    static fromBytes(data) {
        const txid = new TxId();
        txid.data = Buffer.from(data);
        return txid;
    }
    static fromContract(contractPKH, index) {
        // TODO
        // hash256 := sha256.New()
        // hash256.Write(contractPKH)
        // binary.Write(hash256, DefaultEndian, &index)
        // hash := hash256.Sum(nil)
        console.log(contractPKH, index);
        const result = new AssetCode();
        // result.Set(hash);
        return result;
    }
    // Validate returns an error if the value is invalid
    Validate() {
        return null;
    }
    // IsZero returns true if the tx id is all zeros.
    IsZero() {
        return this.data.equals(zeroTxId.data);
    }
    // Equal returns true if the specified values are the same.
    Equal(other) {
        return this.data.equals(other.data);
    }
    // Bytes returns the byte slice for the TxId.
    Bytes() {
        return this.data;
    }
    // String converts to a string
    toString() {
        return sprintf_js_1.sprintf('%x', this.data);
    }
    // Serialize returns a byte slice with the TxId in it.
    Serialize() {
        return this.data;
    }
    // Write reads a TxId from a bytes.Buffer
    Write(buf) {
        return this.data = buf.read(32);
    }
    // MarshalJSON converts to json.
    MarshalJSON() {
        return `"${this.data.toString('hex')}"`;
    }
    // UnmarshalJSON converts from json.
    UnmarshalJSON(data) {
        this.data = Buffer.from(data, 'hex');
    }
    // Set sets the value specified
    Set(value) {
        this.data = Buffer.from(value);
    }
}
exports.AssetCode = AssetCode;
// ------------------------------------------------------------------------------------------------
// Timestamp represents a time in the Tokenized protocol.
class Timestamp {
    constructor(value = Date.now()) {
        this.milliseconds = value;
    }
    // NewTimestamp returns a new timestamp from nanoseconds.
    static NewTimestamp(value) {
        return new Timestamp(value);
    }
    // CurrentTimestamp returns a Timestamp containing the current time.
    static CurrentTimestamp() {
        return new Timestamp(Date.now());
    }
    // Validate returns an error if the value is invalid
    Validate() {
        return null;
    }
    // Equal returns true if the specified values are the same.
    Equal(other) {
        return (this.milliseconds === other.milliseconds);
    }
    // Nano returns the nanoseconds since the Unix epoch for the Timestamp.
    // Nano() uint64 {
    // 	return time.nanoseconds
    // }
    // Nano returns the seconds since the Unix epoch for the Timestamp.
    Seconds() {
        return this.milliseconds / 1000;
    }
    // String converts to a string
    toString() {
        return new Date(this.milliseconds).toString();
    }
    // Serialize returns a byte slice with the Timestamp in it.
    Serialize() {
        const buf = new util_1.default.Writer();
        bytes_1.write(buf, new bn_js_1.default(this.milliseconds).mul(new bn_js_1.default(1000000)), 'uint64');
        return buf.buf;
    }
    // Write reads a Timestamp from a bytes.Buffer
    Write(buf) {
        const bn = bytes_1.read(buf, 'uint64');
        this.milliseconds = bn.div(new bn_js_1.default(1000000)).toNumber();
        console.log('\n\nTimestamp.Write', bn);
        console.log('\n\nTimestamp.Write', this.milliseconds);
    }
}
exports.Timestamp = Timestamp;
class ContractCode {
    constructor() {
        this.data = Buffer.alloc(32, 0);
    }
    static fromBytes(data) {
        const txid = new TxId();
        txid.data = Buffer.from(data);
        return txid;
    }
    // Validate returns an error if the value is invalid
    Validate() {
        return null;
    }
    // IsZero returns true if the tx id is all zeros.
    IsZero() {
        return this.data.equals(zeroTxId.data);
    }
    // Equal returns true if the specified values are the same.
    Equal(other) {
        return this.data.equals(other.data);
    }
    // Bytes returns the byte slice for the TxId.
    Bytes() {
        return this.data;
    }
    // String converts to a string
    toString() {
        return sprintf_js_1.sprintf('%x', this.data);
    }
    // Serialize returns a byte slice with the TxId in it.
    Serialize() {
        return this.data;
    }
    // Write reads a TxId from a bytes.Buffer
    Write(buf) {
        return this.data = buf.read(32);
    }
    // MarshalJSON converts to json.
    MarshalJSON() {
        return `"${this.data.toString('hex')}"`;
    }
    // UnmarshalJSON converts from json.
    UnmarshalJSON(data) {
        this.data = Buffer.from(data, 'hex');
    }
    // Set sets the value specified
    Set(value) {
        this.data = Buffer.from(value);
    }
}
exports.ContractCode = ContractCode;
class PublicKeyHash {
    constructor() {
        this.data = Buffer.alloc(20, 0);
    }
    static fromBytes(data) {
        const txid = new TxId();
        txid.data = Buffer.from(data);
        return txid;
    }
    // Validate returns an error if the value is invalid
    Validate() {
        return null;
    }
    // IsZero returns true if the tx id is all zeros.
    IsZero() {
        return this.data.equals(Buffer.alloc(20, 0));
    }
    // Equal returns true if the specified values are the same.
    Equal(other) {
        return this.data.equals(other.data);
    }
    // Bytes returns the byte slice for the TxId.
    Bytes() {
        return this.data;
    }
    // String converts to a string
    toString() {
        return sprintf_js_1.sprintf('%x', this.data);
    }
    // Serialize returns a byte slice with the TxId in it.
    Serialize() {
        return this.data;
    }
    // Write reads a TxId from a bytes.Buffer
    Write(buf) {
        return this.data = buf.read(20);
    }
    // MarshalJSON converts to json.
    MarshalJSON() {
        return `"${this.data.toString('hex')}"`;
    }
    // UnmarshalJSON converts from json.
    UnmarshalJSON(data) {
        this.data = Buffer.from(data, 'hex');
    }
    // Set sets the value specified
    Set(value) {
        this.data = Buffer.from(value);
    }
}
exports.PublicKeyHash = PublicKeyHash;
