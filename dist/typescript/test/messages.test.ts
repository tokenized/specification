import {sprintf} from 'sprintf-js';
import BN from 'bn.js';
import * as mocha from 'mocha';
import * as chai from 'chai';
import { TxId, Timestamp } from '../src/protocol_types';
import { TargetAddress, OutputTag } from '../src/field_types';
import {
	PublicMessage,
	PrivateMessage,
	RevertedTx,
	Offer,
	SignatureRequest,
	SettlementRequest,
	OutputMetadata,
} from '../src/messages';
const expect = chai.expect;
[mocha]

const getArrayOrType = (type: string) => {
	const regex = /\[(\w+)\]/m;
	let m;
	if ((m = regex.exec(type)) !== null) {
		console.log('m:',  m[1]);
		// const subtype = 
		type.slice(m[0].length);
		return [...Array(parseInt(m[1], 10))].map(() => 0);
	}
	if(type === 'uint64') return new BN(0);
	return 0;
}


describe('PublicMessage', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new PublicMessage();
		// version (uint)
		{
			// IsNativeType
				{
			let type = 'uint8';
			if(type.startsWith('['))
				initialMessage.version = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.version = new BN(65 + 0);
			else
				initialMessage.version = 65 + 0;
				}
		}
		
		// timestamp (Timestamp)
		{
			initialMessage.timestamp = new Timestamp();
		}
		
		// public_message (varchar)
		{
			initialMessage.public_message = "Text 2"
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new PublicMessage();

		decodedMessage.Write(initialEncoding);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// version (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.version);
		let decodedJson = JSON.stringify(decodedMessage.version);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("version doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// timestamp (Timestamp)
		// Timestamp test compare not setup
		// public_message (varchar)
		if (initialMessage.public_message !== (decodedMessage.public_message)) {
			throw new Error(sprintf("public_message doesn't match : %s != %s", initialMessage.public_message, decodedMessage.public_message));
		}


	});
});


describe('PrivateMessage', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new PrivateMessage();
		// version (uint)
		{
			// IsNativeType
				{
			let type = 'uint8';
			if(type.startsWith('['))
				initialMessage.version = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.version = new BN(65 + 0);
			else
				initialMessage.version = 65 + 0;
				}
		}
		
		// timestamp (Timestamp)
		{
			initialMessage.timestamp = new Timestamp();
		}
		
		// private_message (varbin)
		{
			initialMessage.private_message = Buffer.from([...Array(32).keys()]
				.map(i => i + 65 + 2));
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new PrivateMessage();

		decodedMessage.Write(initialEncoding);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// version (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.version);
		let decodedJson = JSON.stringify(decodedMessage.version);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("version doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// timestamp (Timestamp)
		// Timestamp test compare not setup
		// private_message (varbin)
		if ((initialMessage.private_message && decodedMessage.private_message) 
				&& !initialMessage.private_message.equals(decodedMessage.private_message)) {
			throw new Error(sprintf("private_message doesn't match : %x != %x", initialMessage.private_message, decodedMessage.private_message));
		}


	});
});


describe('RevertedTx', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new RevertedTx();
		// version (uint)
		{
			// IsNativeType
				{
			let type = 'uint8';
			if(type.startsWith('['))
				initialMessage.version = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.version = new BN(65 + 0);
			else
				initialMessage.version = 65 + 0;
				}
		}
		
		// timestamp (Timestamp)
		{
			initialMessage.timestamp = new Timestamp();
		}
		
		// transaction (varbin)
		{
			initialMessage.transaction = Buffer.from([...Array(32).keys()]
				.map(i => i + 65 + 2));
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new RevertedTx();

		decodedMessage.Write(initialEncoding);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// version (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.version);
		let decodedJson = JSON.stringify(decodedMessage.version);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("version doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// timestamp (Timestamp)
		// Timestamp test compare not setup
		// transaction (varbin)
		if ((initialMessage.transaction && decodedMessage.transaction) 
				&& !initialMessage.transaction.equals(decodedMessage.transaction)) {
			throw new Error(sprintf("transaction doesn't match : %x != %x", initialMessage.transaction, decodedMessage.transaction));
		}


	});
});


describe('Offer', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new Offer();
		// version (uint)
		{
			// IsNativeType
				{
			let type = 'uint8';
			if(type.startsWith('['))
				initialMessage.version = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.version = new BN(65 + 0);
			else
				initialMessage.version = 65 + 0;
				}
		}
		
		// timestamp (Timestamp)
		{
			initialMessage.timestamp = new Timestamp();
		}
		
		// payload (varbin)
		{
			initialMessage.payload = Buffer.from([...Array(32).keys()]
				.map(i => i + 65 + 2));
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new Offer();

		decodedMessage.Write(initialEncoding);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// version (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.version);
		let decodedJson = JSON.stringify(decodedMessage.version);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("version doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// timestamp (Timestamp)
		// Timestamp test compare not setup
		// payload (varbin)
		if ((initialMessage.payload && decodedMessage.payload) 
				&& !initialMessage.payload.equals(decodedMessage.payload)) {
			throw new Error(sprintf("payload doesn't match : %x != %x", initialMessage.payload, decodedMessage.payload));
		}


	});
});


describe('SignatureRequest', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new SignatureRequest();
		// version (uint)
		{
			// IsNativeType
				{
			let type = 'uint8';
			if(type.startsWith('['))
				initialMessage.version = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.version = new BN(65 + 0);
			else
				initialMessage.version = 65 + 0;
				}
		}
		
		// timestamp (Timestamp)
		{
			initialMessage.timestamp = new Timestamp();
		}
		
		// payload (varbin)
		{
			initialMessage.payload = Buffer.from([...Array(32).keys()]
				.map(i => i + 65 + 2));
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new SignatureRequest();

		decodedMessage.Write(initialEncoding);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// version (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.version);
		let decodedJson = JSON.stringify(decodedMessage.version);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("version doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// timestamp (Timestamp)
		// Timestamp test compare not setup
		// payload (varbin)
		if ((initialMessage.payload && decodedMessage.payload) 
				&& !initialMessage.payload.equals(decodedMessage.payload)) {
			throw new Error(sprintf("payload doesn't match : %x != %x", initialMessage.payload, decodedMessage.payload));
		}


	});
});


describe('SettlementRequest', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new SettlementRequest();
		// version (uint)
		{
			// IsNativeType
				{
			let type = 'uint8';
			if(type.startsWith('['))
				initialMessage.version = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.version = new BN(65 + 0);
			else
				initialMessage.version = 65 + 0;
				}
		}
		
		// timestamp (Timestamp)
		{
			initialMessage.timestamp = new Timestamp();
		}
		
		// transfer_tx_id (TxId)
		{
			initialMessage.transfer_tx_id = new TxId();
		}
		
		// contract_fees (TargetAddress[])
		{
			// IsInternalTypeArray
			initialMessage.contract_fees = [...Array(2)].map(() => new TargetAddress());
		}
		
		// settlement (varbin)
		{
			initialMessage.settlement = Buffer.from([...Array(32).keys()]
				.map(i => i + 65 + 4));
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new SettlementRequest();

		decodedMessage.Write(initialEncoding);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// version (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.version);
		let decodedJson = JSON.stringify(decodedMessage.version);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("version doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// timestamp (Timestamp)
		// Timestamp test compare not setup
		// transfer_tx_id (TxId)
		// TxId test compare not setup
		// contract_fees (TargetAddress[])
		if ((initialMessage.contract_fees && decodedMessage.contract_fees) 
				&& initialMessage.contract_fees.length != decodedMessage.contract_fees.length) {
			throw new Error(sprintf("contract_fees lengths don't match : %d != %d", initialMessage.contract_fees.length, decodedMessage.contract_fees.length));
		}
		// settlement (varbin)
		if ((initialMessage.settlement && decodedMessage.settlement) 
				&& !initialMessage.settlement.equals(decodedMessage.settlement)) {
			throw new Error(sprintf("settlement doesn't match : %x != %x", initialMessage.settlement, decodedMessage.settlement));
		}


	});
});


describe('OutputMetadata', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new OutputMetadata();
		// version (uint)
		{
			// IsNativeType
				{
			let type = 'uint8';
			if(type.startsWith('['))
				initialMessage.version = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.version = new BN(65 + 0);
			else
				initialMessage.version = 65 + 0;
				}
		}
		
		// output_description (varchar)
		{
			initialMessage.output_description = "Text 1"
		}
		
		// tags (Tag[])
		{
			// IsNativeTypeArray
			initialMessage.tags = [...Array(5)].map(() => 121);
		}
		
		// custom_tags (OutputTag[])
		{
			// IsInternalTypeArray
			initialMessage.custom_tags = [...Array(2)].map(() => new OutputTag());
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new OutputMetadata();

		decodedMessage.Write(initialEncoding);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// version (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.version);
		let decodedJson = JSON.stringify(decodedMessage.version);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("version doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// output_description (varchar)
		if (initialMessage.output_description !== (decodedMessage.output_description)) {
			throw new Error(sprintf("output_description doesn't match : %s != %s", initialMessage.output_description, decodedMessage.output_description));
		}
		// tags (Tag[])
			{
		let initialJson = JSON.stringify(initialMessage.tags);
		let decodedJson = JSON.stringify(decodedMessage.tags);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("tags doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// custom_tags (OutputTag[])
		if ((initialMessage.custom_tags && decodedMessage.custom_tags) 
				&& initialMessage.custom_tags.length != decodedMessage.custom_tags.length) {
			throw new Error(sprintf("custom_tags lengths don't match : %d != %d", initialMessage.custom_tags.length, decodedMessage.custom_tags.length));
		}


	});
});


