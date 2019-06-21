import {sprintf} from 'sprintf-js';
import _ from '@keyring/util';

export const write = (buf: _.Writer, value, type: string) => {
	const regex = /\[(\w+)\]/m;
	let m;
	if ((m = regex.exec(type)) !== null) {
		console.log('m:',  m[1]);
		const subtype = type.slice(m[0].length);
		return value.map((v) => {
			return write(buf, v, subtype);
		});
	}

	switch (type) {
		case 'bool': return buf.uint8(value? 1 : 0);
		case 'byte': return buf.uint8(value);
		case 'uint8': return buf.uint8(value);
		case 'uint16': return buf.uint16le(value);
		case 'uint32': return buf.uint32le(value);
		case 'uint64': return buf.uint64le(value);
		default: throw new Error('Invalid type : ' + type);
	}
};

// read fills the value with the appropriate number of bytes from the buffer.
//
// This is useful for fixed size types such as int, float etc.
export const read = (buf: _.Reader, type: string) => {
	const regex = /\[(\w+)\]/m;
	let m;
	if ((m = regex.exec(type)) !== null) {
		console.log('m:',  m[1]);
		const subtype = type.slice(m[0].length);
		return [...Array(parseInt(m[1], 10))].map(() => {
			return read(buf, subtype);
		});
	}

	switch (type) {
		case 'bool': return buf.uint8();
		case 'uint8': return buf.uint8();
		case 'byte': return buf.uint8();
		case 'uint16': return buf.uint16le();
		case 'uint32': return buf.uint32le();
		case 'uint64': return buf.uint64le();
		default: throw new Error('Invalid type : ' + type);
	}
};

export const WriteFixedChar = (buf: _.Writer, value: string, size) => {
	const v = value ? Buffer.from(value, 'ascii') : Buffer.from([]);
	if (v.length > size) {
		throw new Error(`FixedChar too long ${value.length} > ${size}`);
	}
	buf.write(v);

	// Pad with zeroes
	if (v.length < size) {
		buf.write(Buffer.alloc(size - v.length));
	}
	return;
};

// WriteVarBin writes a variable binary value
export const WriteVarBin = (buf: _.Writer, value, sizeBits: number) => {
	console.log('WriteVarBin ', value);
	const v = Buffer.from(value || []);
	WriteVariableSize(buf, v.length, sizeBits, 0);
	buf.write(v);
	return;
};

// WriteVariableSize writes a size/length to the buffer using the specified size unsigned integer.
// defaultSizeBits is used if sizeBits is zero.
export const WriteVariableSize = (buf: _.Writer, size: number, sizeBits: number, defaultSizeBits: number) => {
	if (sizeBits === 0) {
		sizeBits = defaultSizeBits;
	}
	console.log('WriteVariableSize sizeBits', sizeBits, (2**sizeBits)-1, size);
	if (size >= 2**sizeBits) {
		throw new Error(sprintf('Size beyond size bits limit (%d) : %d', (2**sizeBits)-1, size));
	}

	switch (sizeBits) {
	case 8:
		return write(buf, size, 'uint8');
	case 16:
		return write(buf, size, 'uint16');
	case 32:
		return write(buf, size, 'uint32');
	case 64:
		return write(buf, size, 'uint64');
	default:
		throw new Error('Invalid variable size bits : ' + sizeBits);
	}
};

// WriteVarChar writes a variable character string
export const WriteVarChar = (buf: _.Writer, value: string, sizeBits: number) => {
	const v = value ? Buffer.from(value, 'ascii') : Buffer.from([]);
	WriteVariableSize(buf, v.length, sizeBits, 0);
	buf.write(v);
};

// ReadFixedChar reads a fixed character string
export const ReadFixedChar = (buf: _.Reader, size: number): string => {
	return buf.read(size).toString('ascii');
};

// ReadVarBin reads a variable binary value
export const ReadVarBin = (buf: _.Reader, sizeBits: number): Buffer => {
	const size = ReadVariableSize(buf, sizeBits, 0);
	return buf.read(size);
};

// ReadVariableSize reads a size/length from the buffer using the specified size unsigned integer.
// defaultSizeBits is used if sizeBits is zero.
export const ReadVariableSize = (buf: _.Reader, sizeBits: number, defaultSizeBits: number): number => {
	if (sizeBits === 0) {
		sizeBits = defaultSizeBits;
	}

	switch (sizeBits) {
	case 8: return read(buf, 'uint8');
	case 16: return read(buf, 'uint16');
	case 32: return read(buf, 'uint32');
	case 64: return read(buf, 'uint64');
	default: throw new Error('Invalid variable size bits : ' + sizeBits);
	}
};

// ReadVarChar reads a variable character string
export const ReadVarChar = (buf: _.Reader, sizeBits: number): string => {
	const size = ReadVariableSize(buf, sizeBits, 0);
	return buf.read(size).toString('ascii');
};

// { write, read, ReadVarChar, ReadVariableSize, ReadVarBin, ReadFixedChar, WriteVarChar, WriteVariableSize, WriteFixedChar }
