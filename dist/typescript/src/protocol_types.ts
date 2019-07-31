import {sprintf} from 'sprintf-js';
import _ from '@keyring/util';
import BN from 'bn.js';
import { read, write } from './bytes';


// ------------------------------------------------------------------------------------------------
// TxId represents a Bitcoin transaction ID. (Double SHA256 of tx data)
export class TxId {
	data = Buffer.alloc(32, 0);

	static fromBytes(data: Buffer): TxId {
		const txid = new TxId();
		txid.data = Buffer.from(data);
		return txid;
	}

	// Validate returns an error if the value is invalid
	Validate(): string {
		return null;
	}

	// IsZero returns true if the tx id is all zeros.
	IsZero(): boolean {
		return this.data.equals(zeroTxId.data);
	}

	// Equal returns true if the specified values are the same.
	Equal(other: TxId): boolean {
		return this.data.equals(other.data);
	}

	// Bytes returns the byte slice for the TxId.
	Bytes(): Buffer {
		return this.data;
	}

	// String converts to a string
	toString(): string {
		return sprintf('%x', this.data);
	}

	// Serialize returns a byte slice with the TxId in it.
	Serialize(): Buffer {
		return this.data;
	}

	// Write reads a TxId from a bytes.Buffer
	Write(buf: _.Reader) {
		return this.data = buf.read(32);
	}

	// MarshalJSON converts to json.
	MarshalJSON(): string {
		return `"${this.data.toString('hex')}"`;
	}

	// UnmarshalJSON converts from json.
	UnmarshalJSON(data: string) {
		this.data = Buffer.from(data, 'hex');
	}

	// Set sets the value specified
	Set(value: Buffer) {
		this.data = Buffer.from(value);
	}
}

const zeroTxId = new TxId();

export class AssetCode {
	data = Buffer.alloc(32, 0);

	static fromBytes(data: Buffer): AssetCode {
		const code = new AssetCode();
		code.data = Buffer.from(data);
		return code;
	}

	static fromContract(contractPKH: Buffer, index: number): AssetCode {
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
	Validate(): string {
		return null;
	}

	// IsZero returns true if the tx id is all zeros.
	IsZero(): boolean {
		return this.data.equals(zeroTxId.data);
	}

	// Equal returns true if the specified values are the same.
	Equal(other: TxId): boolean {
		return this.data.equals(other.data);
	}

	// Bytes returns the byte slice for the TxId.
	Bytes(): Buffer {
		return this.data;
	}

	// String converts to a string
	toString(): string {
		return sprintf('%x', this.data);
	}

	// Serialize returns a byte slice with the TxId in it.
	Serialize(): Buffer {
		return this.data;
	}

	// Write reads a TxId from a bytes.Buffer
	Write(buf: _.Reader) {
		return this.data = buf.read(32);
	}

	// MarshalJSON converts to json.
	MarshalJSON(): string {
		return `"${this.data.toString('hex')}"`;
	}

	// UnmarshalJSON converts from json.
	UnmarshalJSON(data: string) {
		this.data = Buffer.from(data, 'hex');
	}

	// MarshalJSON converts to json.
	toJSON(): string {
		return `${this.data.toString('hex')}`;
	}

	// UnmarshalJSON converts from json.
	fromJSON(data: string) {
		this.data = Buffer.from(data, 'hex');
	}

	// Set sets the value specified
	Set(value: Buffer) {
		this.data = Buffer.from(value);
	}
}

// ------------------------------------------------------------------------------------------------
// Timestamp represents a time in the Tokenized protocol.
export class Timestamp {
	// nanoseconds; // uint64 // nanoseconds since Unix epoch
	milliseconds: number;

	constructor(value = Date.now()) {
		this.milliseconds = value;
	}
	// NewTimestamp returns a new timestamp from nanoseconds.
	static NewTimestamp(value): Timestamp {
		return new Timestamp(value);
	}

	// CurrentTimestamp returns a Timestamp containing the current time.
	static CurrentTimestamp(): Timestamp {
		return new Timestamp(Date.now());
	}

	// Validate returns an error if the value is invalid
	Validate(): string {
		return null;
	}

	// Equal returns true if the specified values are the same.
	Equal(other: Timestamp): boolean {
		return (this.milliseconds === other.milliseconds);
	}

	// Nano returns the nanoseconds since the Unix epoch for the Timestamp.
	// Nano() uint64 {
	// 	return time.nanoseconds
	// }

	// Nano returns the seconds since the Unix epoch for the Timestamp.
	Seconds(): number {
		return this.milliseconds / 1000;
	}

	// String converts to a string
	toString(): string {
		return new Date(this.milliseconds).toString();
	}

	// Serialize returns a byte slice with the Timestamp in it.
	Serialize(): Buffer {
		const buf = new _.Writer();
		write(buf, new BN(this.milliseconds).mul(new BN(1000)), 'uint64');
		return buf.buf;
	}

	// Write reads a Timestamp from a bytes.Buffer
	Write(buf: _.Reader) {
		const bn = read(buf, 'uint64');
		this.milliseconds = bn.div(new BN(1000)).toNumber();
		// console.log('\n\nTimestamp.Write', bn);
		// console.log('\n\nTimestamp.Write', this.milliseconds);
	}

	// // MarshalJSON converts to json.
	// MarshalJSON() ([]byte, error) {
	// 	return []byte(strconv.FormatUint(time.nanoseconds, 10)), nil
	// }

	// // UnmarshalJSON converts from json.
	// UnmarshalJSON(data []byte) error {
	// 	if len(data) == 2 && data[0] == '{' && data[1] == '}' {
	// 		time.nanoseconds = 0
	// 		return nil
	// 	}

	// 	var err error
	// 	time.nanoseconds, err = strconv.ParseUint(string(data), 10, 64)
	// 	return err
	// }

}

export class ContractCode {
	data = Buffer.alloc(32, 0);

	static fromBytes(data: Buffer): ContractCode {
		const code = new ContractCode();
		code.data = Buffer.from(data);
		return code;
	}

	// Validate returns an error if the value is invalid
	Validate(): string {
		return null;
	}

	// IsZero returns true if the tx id is all zeros.
	IsZero(): boolean {
		return this.data.equals(zeroTxId.data);
	}

	// Equal returns true if the specified values are the same.
	Equal(other: TxId): boolean {
		return this.data.equals(other.data);
	}

	// Bytes returns the byte slice for the TxId.
	Bytes(): Buffer {
		return this.data;
	}

	// String converts to a string
	toString(): string {
		return sprintf('%x', this.data);
	}

	// Serialize returns a byte slice with the TxId in it.
	Serialize(): Buffer {
		return this.data;
	}

	// Write reads a TxId from a bytes.Buffer
	Write(buf: _.Reader) {
		return this.data = buf.read(32);
	}

	// MarshalJSON converts to json.
	MarshalJSON(): string {
		return `"${this.data.toString('hex')}"`;
	}

	// UnmarshalJSON converts from json.
	UnmarshalJSON(data: string) {
		this.data = Buffer.from(data, 'hex');
	}

	// Set sets the value specified
	Set(value: Buffer) {
		this.data = Buffer.from(value);
	}
}

export class PublicKeyHash {
	data = Buffer.alloc(20, 0);

	static fromBytes(data: Buffer): PublicKeyHash {
		const pkh = new PublicKeyHash();
		pkh.data = Buffer.from(data);
		return pkh;
	}
	static fromStr(addr: string): PublicKeyHash {
		const pkh = new PublicKeyHash();
		pkh.data = Buffer.from(_.addr.from(addr).hash);
		return pkh;
	}

	// Validate returns an error if the value is invalid
	Validate(): string {
		return null;
	}

	// IsZero returns true if the tx id is all zeros.
	IsZero(): boolean {
		return this.data.equals(Buffer.alloc(20, 0));
	}

	// Equal returns true if the specified values are the same.
	Equal(other: TxId): boolean {
		return this.data.equals(other.data);
	}

	// Bytes returns the byte slice for the TxId.
	Bytes(): Buffer {
		return this.data;
	}

	toJSON() {
		return this.format();
	}

	format() {
		return _.addr.format(this.data);
	}

	// String converts to a string
	toString(): string {
		return sprintf('%x', this.data);
	}

	// Serialize returns a byte slice with the TxId in it.
	Serialize(): Buffer {
		return this.data;
	}

	// Write reads a TxId from a bytes.Buffer
	Write(buf: _.Reader) {
		return this.data = buf.read(20);
	}

	// MarshalJSON converts to json.
	MarshalJSON(): string {
		return `"${this.data.toString('hex')}"`;
	}

	// UnmarshalJSON converts from json.
	UnmarshalJSON(data: string) {
		this.data = Buffer.from(data, 'hex');
	}

	// Set sets the value specified
	Set(value: Buffer) {
		this.data = Buffer.from(value);
	}
}
