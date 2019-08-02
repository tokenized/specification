"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (Object.hasOwnProperty.call(mod, k)) result[k] = mod[k];
    result["default"] = mod;
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
const sprintf_js_1 = require("sprintf-js");
const util_1 = __importDefault(require("@keyring/util"));
// import BN from 'bn.js';
const actions_1 = require("./actions");
const R = __importStar(require("ramda"));
const filterData = R.filter((e) => R.is(Array, e) && e.length > 1);
// OpReturn (OP_RETURN) is a script opcode is used to mark a transaction
// output as invalid, and can be used to add data to a TX.
const OP_RETURN = 0x6a;
// const OP_DUP          = 0x76;
// const OP_HASH160      = 0xa9;
// const OP_PUSH_DATA_20 = 0x14;
// const OP_EQUALVERIFY  = 0x88;
// const OP_CHECKSIG     = 0xac;
// OP_MAX_SINGLE_BYTE_PUSH_DATA represents the max length for a single byte push
const OP_MAX_SINGLE_BYTE_PUSH_DATA = 0x4b;
// OP_PUSH_DATA_1 represent the OP_PUSHDATA1 opcode.
const OP_PUSH_DATA_1 = 0x4c;
// OP_PUSH_DATA_2 represents the OP_PUSHDATA2 opcode.
const OP_PUSH_DATA_2 = 0x4d;
// OP_PUSH_DATA_4 represents the OP_PUSHDATA4 opcode.
const OP_PUSH_DATA_4 = 0x4e;
// OP_PUSH_DATA_1_MAX is the maximum number of bytes that can be used in the
// OP_PUSHDATA1 opcode.
const OP_PUSH_DATA_1_MAX = 255;
// OP_PUSH_DATA_2_MAX is the maximum number of bytes that can be used in the
// OP_PUSHDATA2 opcode.
const OP_PUSH_DATA_2_MAX = 65535;
function ParsePushDataScript(buf) {
    const opCode = buf.uint8();
    if (opCode <= OP_MAX_SINGLE_BYTE_PUSH_DATA) {
        return opCode;
    }
    switch (opCode) {
        case OP_PUSH_DATA_1:
            return buf.uint8();
        case OP_PUSH_DATA_2:
            return buf.uint16le();
        case OP_PUSH_DATA_4:
            return buf.uint32le();
        default: throw new Error(sprintf_js_1.sprintf('Invalid push data op code : 0x%02x', opCode));
    }
}
// PushDataScript prepares a push data script based on the given size
function PushDataScript(size) {
    const buf = new util_1.default.Writer();
    if (size <= OP_MAX_SINGLE_BYTE_PUSH_DATA) {
        return buf.uint8(size).buf; // Single byte push
    }
    else if (size < OP_PUSH_DATA_1_MAX) {
        return buf.write(Buffer.from([OP_PUSH_DATA_1, size])).buf;
    }
    else if (size < OP_PUSH_DATA_2_MAX) {
        buf.uint8(OP_PUSH_DATA_2);
        buf.uint16le(size);
        return buf.buf;
    }
    buf.uint8(OP_PUSH_DATA_4);
    buf.uint32le(size);
    return buf.buf;
}
// OpReturnMessage implements a base interface for all message types.
class OpReturnMessage {
    // Deserialize returns a message, as an OpReturnMessage, from the OP_RETURN script.
    static Deserialize(b, isTest = false) {
        const buf = new util_1.default.Reader(b);
        // Parse OP_RETURN op code
        const opCode = buf.uint8();
        if (opCode !== OP_RETURN) {
            throw new Error(sprintf_js_1.sprintf('Not an op return output : %02x', opCode));
        }
        // Parse protocol ID
        const protocolId = isTest ? OpReturnMessage.TestProtocolID : OpReturnMessage.ProtocolID;
        let size = ParsePushDataScript(buf);
        if (size !== protocolId.length) {
            throw new Error(sprintf_js_1.sprintf('Push not correct size for protocol ID : %d != %d', size, protocolId.length));
        }
        const readProtocolId = buf.read(size).toString('ascii');
        if (readProtocolId !== protocolId) {
            throw new Error('Invalid protocol ID:' + readProtocolId);
        }
        // Parse push op code for payload length + 3 for version and message type code
        size = ParsePushDataScript(buf);
        // if uint64(buf.Len()) < payloadSize {
        // 	return nil, sprintf('Payload push op code is too large for message : %d', payloadSize)
        // }
        // Parse version
        const version = buf.uint8();
        if (version !== OpReturnMessage.Version) {
            throw new Error(sprintf_js_1.sprintf('Unsupported version : %02x', version));
        }
        // Parse message type code
        const code = buf.read(2).toString('ascii');
        const msg = actions_1.TypeMapping(code);
        if (!msg) {
            throw new Error(sprintf_js_1.sprintf('Unknown code : %s', code));
        }
        msg.write(buf.read(size - 3));
        return msg;
    }
    // Serialize returns a complete op return script including the specified payload.
    static Serialize(msg, isTest = false) {
        const buf = new util_1.default.Writer();
        // Write OP_RETURN op code
        buf.uint8(OP_RETURN);
        const protocolId = isTest ? OpReturnMessage.TestProtocolID : OpReturnMessage.ProtocolID;
        // Write protocol Id push op code
        buf.write(PushDataScript(protocolId.length));
        // Write protocol Id
        buf.write(Buffer.from(protocolId, 'ascii'));
        const payload = msg.Serialize();
        // Write push op code for payload length + 3 for version and message type code
        buf.write(PushDataScript(payload.length + 3));
        // Write version
        buf.uint8(OpReturnMessage.Version);
        // Write message type code
        buf.write(Buffer.from(msg.Type(), 'ascii'));
        // Write payload
        buf.write(payload);
        return buf.buf;
    }
    // Code returns the identifying code from the OP_RETURN payload.
    static Code(b, isTest = false) {
        const buf = new util_1.default.Reader(b);
        // Parse OP_RETURN op code
        const opCode = buf.uint8();
        if (opCode !== OP_RETURN) {
            throw new Error(sprintf_js_1.sprintf('Not an op return output : %02x', opCode));
        }
        // Parse protocol ID
        const protocolId = isTest ? OpReturnMessage.TestProtocolID : OpReturnMessage.ProtocolID;
        let size = ParsePushDataScript(buf);
        if (size !== protocolId.length) {
            throw new Error(sprintf_js_1.sprintf('Push not correct size for protocol ID : %d != %d', size, protocolId.length));
        }
        const readProtocolId = buf.read(size).toString('ascii');
        if (readProtocolId !== protocolId) {
            throw new Error('Invalid protocol ID:' + readProtocolId);
        }
        // Parse push op code for payload length + 3 for version and message type code
        size = ParsePushDataScript(buf);
        // if uint64(buf.Len()) < payloadSize {
        // 	return nil, sprintf('Payload push op code is too large for message : %d', payloadSize)
        // }
        // Parse version
        const version = buf.uint8();
        if (version !== OpReturnMessage.Version) {
            throw new Error(sprintf_js_1.sprintf('Unsupported version : %02x', version));
        }
        // Parse message type code
        return buf.read(2).toString('ascii');
    }
    static getMsg(b) {
        const buf = new util_1.default.Reader(b);
        const version = buf.uint8();
        if (version !== OpReturnMessage.Version) {
            throw new Error(sprintf_js_1.sprintf('Unsupported version : %02x', version));
        }
        // Parse message type code
        const code = buf.read(2).toString('ascii');
        const msg = actions_1.TypeMapping(code);
        if (!msg) {
            throw new Error('Unknown code: ' + code);
        }
        msg.write(buf.read(b.length - 3));
        return msg;
    }
    static keyringTxTokMsgs(tx) {
        return filterData(tx.data()).map(arr => [arr[0].toString('ascii'), OpReturnMessage.getMsg(arr[1])]);
    }
    static bitdbToTokMsgs(tx) {
        return tx.out.filter((out) => out.str.includes('OP_RETURN') && out.str.includes(OpReturnMessage.ProtocolIdHex))
            .map((out) => {
            const parts = out.str.split(' ');
            return [
                Buffer.from(parts[1], 'hex').toString('ascii'),
                OpReturnMessage.getMsg(Buffer.from(parts[2], 'hex')),
            ];
        });
    }
    static addTokMsgToKeyringTx(tx, msg, isTest = false) {
        return tx.data(Buffer.from(isTest ? OpReturnMessage.TestProtocolID : OpReturnMessage.ProtocolID, 'ascii'), Buffer.concat([Buffer.from([OpReturnMessage.Version]), Buffer.from(msg.Type(), 'ascii'), msg.Serialize()]));
    }
    Type() { return null; }
    String() { return null; }
    Serialize() { return null; }
    write(_) { return 0; }
    Validate() { return null; }
}
// ProtocolID is the current protocol ID
OpReturnMessage.ProtocolID = 'tokenized';
// TestProtocolID is the current protocol ID for testing.
OpReturnMessage.TestProtocolID = 'test.tokenized';
// Version of the Tokenized protocol.
OpReturnMessage.Version = 0;
OpReturnMessage.ProtocolIdHex = Buffer.from(OpReturnMessage.ProtocolID, 'ascii').toString('hex');
exports.OpReturnMessage = OpReturnMessage;
