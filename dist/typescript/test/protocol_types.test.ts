import * as chai from 'chai';
import { TxId } from '../dist/protocol_types';
const expect = chai.expect;

describe('TxId', () => {
	it('given a valid TxId, when toString() is called, then the output is a hex string of all 32 bytes.', () => {
    // GIVEN
    const txIdHex = '0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20';
    const txIdBuffer = Buffer.from(txIdHex, 'hex');
    const txId = TxId.fromBytes(txIdBuffer);
    
    // WHEN
    const outputString = txId.toString();

    // THEN
    expect(outputString).to.eql(txIdHex);
	});
});
