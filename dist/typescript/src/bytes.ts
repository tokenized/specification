import {sprintf} from 'sprintf-js';
import _ from '@keyring/util';

export const write = (buf: _.Writer, value, type: string) => {
	// TODO handle arrays
	switch (type) {
		case 'bool': return buf.uint8(value? 1 : 0);
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
	// TODO handle arrays
	switch (type) {
		case 'bool': return buf.uint8();
		case 'uint8': return buf.uint8();
		case 'uint16': return buf.uint16le();
		case 'uint32': return buf.uint32le();
		case 'uint64': return buf.uint64le();
		default: throw new Error('Invalid type : ' + type);
	}
};

export const WriteFixedChar = (buf: _.Writer, value: string, size) => {
	const v = Buffer.from(value, 'ascii');
	if (v.length > size) {
		throw new Error(`FixedChar too long ${value.length} > ${size}`);
	}
	buf.Write(v);

	// Pad with zeroes
	if (v.length < size) {
		buf.Write(Buffer.alloc(size - v.length));
	}
	return;
};

// WriteVarBin writes a variable binary value
export const WriteVarBin = (buf: _.Writer, value, sizeBits: number) => {
	WriteVariableSize(buf, value.length, sizeBits, 0);
	buf.write(value);
	return;
};

// WriteVariableSize writes a size/length to the buffer using the specified size unsigned integer.
// defaultSizeBits is used if sizeBits is zero.
export const WriteVariableSize = (buf: _.Writer, size: number, sizeBits: number, defaultSizeBits: number) => {
	if (sizeBits === 0) {
		sizeBits = defaultSizeBits;
	}
	if (size >= 2<<sizeBits) {
		throw new Error(sprintf('Size beyond size bits limit (%d) : %d', (2<<sizeBits)-1, size));
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
	const v = Buffer.from(value, 'ascii');
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
