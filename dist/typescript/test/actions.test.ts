import {sprintf} from 'sprintf-js';
import BN from 'bn.js';
import * as mocha from 'mocha';
import * as chai from 'chai';
import {char} from '../src/bytes';
import { TxId, AssetCode, Timestamp, ContractCode, PublicKeyHash } from '../src/protocol_types';
import { Document, Amendment, VotingSystem, Oracle, Entity, TargetAddress,
	QuantityIndex, AssetTransfer, AssetSettlement } from '../src/field_types';
import {
	AssetDefinition,
	AssetCreation,
	AssetModification,
	ContractOffer,
	ContractFormation,
	ContractAmendment,
	StaticContractFormation,
	ContractAddressChange,
	Order,
	Freeze,
	Thaw,
	Confiscation,
	Reconciliation,
	Proposal,
	Vote,
	BallotCast,
	BallotCounted,
	Result,
	Message,
	Rejection,
	Establishment,
	Addition,
	Alteration,
	Removal,
	Transfer,
	Settlement,
} from '../src/actions';
import R from 'ramda';
const expect = chai.expect;
[mocha]

const getArrayOrType = (type: string) => {
	const regex = /\[(\w+)\]/m;
	let m;
	if ((m = regex.exec(type)) !== null) {
		console.log('m:',  m[1]);
		const subtype = type.slice(m[0].length);
		return [...Array(parseInt(m[1], 10))].map(() => 0);
	}
	if(type === 'uint64') return new BN(0);
	return 0;
}


describe('AssetDefinition', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new AssetDefinition();
		// asset_type (fixedchar)
		{
			initialMessage.asset_type = [...Array(3).keys()]
				.map(i => String.fromCharCode(i + 65 + 0))
				.join('');
		}
		
		// asset_auth_flags (varbin)
		{
			initialMessage.asset_auth_flags = Buffer.from([...Array(8).keys()]
				.map(i => i + 65 + 1));
		}
		
		// transfers_permitted (bool)
		{
			initialMessage.transfers_permitted = true
		}
		
		// trade_restrictions (Polity[])
		{
			// IsNativeTypeArray
			initialMessage.trade_restrictions = [...Array(5)].map(() => Buffer.from('AUS', 'ascii'));
		}
		
		// enforcement_orders_permitted (bool)
		{
			initialMessage.enforcement_orders_permitted = true
		}
		
		// voting_rights (bool)
		{
			initialMessage.voting_rights = true
		}
		
		// vote_multiplier (uint)
		{
			// IsNativeType
				{
			let type = 'uint8';
			if(type.startsWith('['))
				initialMessage.vote_multiplier = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.vote_multiplier = new BN(65 + 6);
			else
				initialMessage.vote_multiplier = 65 + 6;
				}
		}
		
		// administration_proposal (bool)
		{
			initialMessage.administration_proposal = true
		}
		
		// holder_proposal (bool)
		{
			initialMessage.holder_proposal = true
		}
		
		// asset_modification_governance (uint)
		{
				initialMessage.asset_modification_governance = 1;
		}
		
		// token_qty (uint)
		{
			// IsNativeType
				{
			let type = 'uint64';
			if(type.startsWith('['))
				initialMessage.token_qty = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.token_qty = new BN(65 + 10);
			else
				initialMessage.token_qty = 65 + 10;
				}
		}
		
		// asset_payload (varbin)
		{
			initialMessage.asset_payload = Buffer.from([...Array(16).keys()]
				.map(i => i + 65 + 11));
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new AssetDefinition();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// asset_type (fixedchar)
			// IsFixedChar
		if (initialMessage.asset_type != decodedMessage.asset_type) {
			throw new Error(sprintf("asset_type doesn't match : %s != %s", initialMessage.asset_type, decodedMessage.asset_type));
		}
		// asset_auth_flags (varbin)
		if ((initialMessage.asset_auth_flags && decodedMessage.asset_auth_flags) 
				&& !initialMessage.asset_auth_flags.equals(decodedMessage.asset_auth_flags)) {
			throw new Error(sprintf("asset_auth_flags doesn't match : %x != %x", initialMessage.asset_auth_flags, decodedMessage.asset_auth_flags));
		}
		// transfers_permitted (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.transfers_permitted);
		let decodedJson = JSON.stringify(decodedMessage.transfers_permitted);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("transfers_permitted doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// trade_restrictions (Polity[])
			{
		let initialJson = JSON.stringify(initialMessage.trade_restrictions);
		let decodedJson = JSON.stringify(decodedMessage.trade_restrictions);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("trade_restrictions doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// enforcement_orders_permitted (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.enforcement_orders_permitted);
		let decodedJson = JSON.stringify(decodedMessage.enforcement_orders_permitted);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("enforcement_orders_permitted doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// voting_rights (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.voting_rights);
		let decodedJson = JSON.stringify(decodedMessage.voting_rights);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("voting_rights doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// vote_multiplier (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.vote_multiplier);
		let decodedJson = JSON.stringify(decodedMessage.vote_multiplier);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("vote_multiplier doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// administration_proposal (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.administration_proposal);
		let decodedJson = JSON.stringify(decodedMessage.administration_proposal);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("administration_proposal doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// holder_proposal (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.holder_proposal);
		let decodedJson = JSON.stringify(decodedMessage.holder_proposal);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("holder_proposal doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// asset_modification_governance (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.asset_modification_governance);
		let decodedJson = JSON.stringify(decodedMessage.asset_modification_governance);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("asset_modification_governance doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// token_qty (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.token_qty);
		let decodedJson = JSON.stringify(decodedMessage.token_qty);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("token_qty doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// asset_payload (varbin)
		if ((initialMessage.asset_payload && decodedMessage.asset_payload) 
				&& !initialMessage.asset_payload.equals(decodedMessage.asset_payload)) {
			throw new Error(sprintf("asset_payload doesn't match : %x != %x", initialMessage.asset_payload, decodedMessage.asset_payload));
		}


	});
});


describe('AssetCreation', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new AssetCreation();
		// asset_type (fixedchar)
		{
			initialMessage.asset_type = [...Array(3).keys()]
				.map(i => String.fromCharCode(i + 65 + 0))
				.join('');
		}
		
		// asset_code (AssetCode)
		{
			initialMessage.asset_code = new AssetCode();
			
		}
		
		// asset_index (uint)
		{
			// IsNativeType
				{
			let type = 'uint64';
			if(type.startsWith('['))
				initialMessage.asset_index = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.asset_index = new BN(65 + 2);
			else
				initialMessage.asset_index = 65 + 2;
				}
		}
		
		// asset_auth_flags (varbin)
		{
			initialMessage.asset_auth_flags = Buffer.from([...Array(8).keys()]
				.map(i => i + 65 + 3));
		}
		
		// transfers_permitted (bool)
		{
			initialMessage.transfers_permitted = true
		}
		
		// trade_restrictions (Polity[])
		{
			// IsNativeTypeArray
			initialMessage.trade_restrictions = [...Array(5)].map(() => Buffer.from('AUS', 'ascii'));
		}
		
		// enforcement_orders_permitted (bool)
		{
			initialMessage.enforcement_orders_permitted = true
		}
		
		// voting_rights (bool)
		{
			initialMessage.voting_rights = true
		}
		
		// vote_multiplier (uint)
		{
			// IsNativeType
				{
			let type = 'uint8';
			if(type.startsWith('['))
				initialMessage.vote_multiplier = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.vote_multiplier = new BN(65 + 8);
			else
				initialMessage.vote_multiplier = 65 + 8;
				}
		}
		
		// administration_proposal (bool)
		{
			initialMessage.administration_proposal = true
		}
		
		// holder_proposal (bool)
		{
			initialMessage.holder_proposal = true
		}
		
		// asset_modification_governance (uint)
		{
				initialMessage.asset_modification_governance = 1;
		}
		
		// token_qty (uint)
		{
			// IsNativeType
				{
			let type = 'uint64';
			if(type.startsWith('['))
				initialMessage.token_qty = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.token_qty = new BN(65 + 12);
			else
				initialMessage.token_qty = 65 + 12;
				}
		}
		
		// asset_payload (varbin)
		{
			initialMessage.asset_payload = Buffer.from([...Array(16).keys()]
				.map(i => i + 65 + 13));
		}
		
		// asset_revision (uint)
		{
			// IsNativeType
				{
			let type = 'uint32';
			if(type.startsWith('['))
				initialMessage.asset_revision = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.asset_revision = new BN(65 + 14);
			else
				initialMessage.asset_revision = 65 + 14;
				}
		}
		
		// timestamp (Timestamp)
		{
			initialMessage.timestamp = new Timestamp();
			
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new AssetCreation();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// asset_type (fixedchar)
			// IsFixedChar
		if (initialMessage.asset_type != decodedMessage.asset_type) {
			throw new Error(sprintf("asset_type doesn't match : %s != %s", initialMessage.asset_type, decodedMessage.asset_type));
		}
		// asset_code (AssetCode)
		// AssetCode test compare not setup
		// asset_index (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.asset_index);
		let decodedJson = JSON.stringify(decodedMessage.asset_index);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("asset_index doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// asset_auth_flags (varbin)
		if ((initialMessage.asset_auth_flags && decodedMessage.asset_auth_flags) 
				&& !initialMessage.asset_auth_flags.equals(decodedMessage.asset_auth_flags)) {
			throw new Error(sprintf("asset_auth_flags doesn't match : %x != %x", initialMessage.asset_auth_flags, decodedMessage.asset_auth_flags));
		}
		// transfers_permitted (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.transfers_permitted);
		let decodedJson = JSON.stringify(decodedMessage.transfers_permitted);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("transfers_permitted doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// trade_restrictions (Polity[])
			{
		let initialJson = JSON.stringify(initialMessage.trade_restrictions);
		let decodedJson = JSON.stringify(decodedMessage.trade_restrictions);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("trade_restrictions doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// enforcement_orders_permitted (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.enforcement_orders_permitted);
		let decodedJson = JSON.stringify(decodedMessage.enforcement_orders_permitted);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("enforcement_orders_permitted doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// voting_rights (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.voting_rights);
		let decodedJson = JSON.stringify(decodedMessage.voting_rights);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("voting_rights doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// vote_multiplier (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.vote_multiplier);
		let decodedJson = JSON.stringify(decodedMessage.vote_multiplier);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("vote_multiplier doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// administration_proposal (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.administration_proposal);
		let decodedJson = JSON.stringify(decodedMessage.administration_proposal);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("administration_proposal doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// holder_proposal (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.holder_proposal);
		let decodedJson = JSON.stringify(decodedMessage.holder_proposal);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("holder_proposal doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// asset_modification_governance (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.asset_modification_governance);
		let decodedJson = JSON.stringify(decodedMessage.asset_modification_governance);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("asset_modification_governance doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// token_qty (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.token_qty);
		let decodedJson = JSON.stringify(decodedMessage.token_qty);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("token_qty doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// asset_payload (varbin)
		if ((initialMessage.asset_payload && decodedMessage.asset_payload) 
				&& !initialMessage.asset_payload.equals(decodedMessage.asset_payload)) {
			throw new Error(sprintf("asset_payload doesn't match : %x != %x", initialMessage.asset_payload, decodedMessage.asset_payload));
		}
		// asset_revision (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.asset_revision);
		let decodedJson = JSON.stringify(decodedMessage.asset_revision);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("asset_revision doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// timestamp (Timestamp)
		// Timestamp test compare not setup


	});
});


describe('AssetModification', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new AssetModification();
		// asset_type (fixedchar)
		{
			initialMessage.asset_type = [...Array(3).keys()]
				.map(i => String.fromCharCode(i + 65 + 0))
				.join('');
		}
		
		// asset_code (AssetCode)
		{
			initialMessage.asset_code = new AssetCode();
			
		}
		
		// asset_revision (uint)
		{
			// IsNativeType
				{
			let type = 'uint32';
			if(type.startsWith('['))
				initialMessage.asset_revision = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.asset_revision = new BN(65 + 2);
			else
				initialMessage.asset_revision = 65 + 2;
				}
		}
		
		// amendments (Amendment[])
		{
			// IsInternalTypeArray
			initialMessage.amendments = [...Array(2)].map(() => new Amendment());
		}
		
		// ref_tx_id (TxId)
		{
			initialMessage.ref_tx_id = new TxId();
			
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new AssetModification();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// asset_type (fixedchar)
			// IsFixedChar
		if (initialMessage.asset_type != decodedMessage.asset_type) {
			throw new Error(sprintf("asset_type doesn't match : %s != %s", initialMessage.asset_type, decodedMessage.asset_type));
		}
		// asset_code (AssetCode)
		// AssetCode test compare not setup
		// asset_revision (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.asset_revision);
		let decodedJson = JSON.stringify(decodedMessage.asset_revision);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("asset_revision doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// amendments (Amendment[])
		if ((initialMessage.amendments && decodedMessage.amendments) 
				&& initialMessage.amendments.length != decodedMessage.amendments.length) {
			throw new Error(sprintf("amendments lengths don't match : %d != %d", initialMessage.amendments.length, decodedMessage.amendments.length));
		}
		// ref_tx_id (TxId)
		// TxId test compare not setup


	});
});


describe('ContractOffer', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new ContractOffer();
		// contract_name (varchar)
		{
			initialMessage.contract_name = "Text 0"
		}
		
		// body_of_agreement_type (uint)
		{
				initialMessage.body_of_agreement_type = 1;
		}
		
		// body_of_agreement (varbin)
		{
			initialMessage.body_of_agreement = Buffer.from([...Array(32).keys()]
				.map(i => i + 65 + 2));
		}
		
		// contract_type (varchar)
		{
			initialMessage.contract_type = "Text 3"
		}
		
		// supporting_docs (Document[])
		{
			// IsInternalTypeArray
			initialMessage.supporting_docs = [...Array(2)].map(() => new Document());
		}
		
		// governing_law (fixedchar)
		{
			initialMessage.governing_law = [...Array(5).keys()]
				.map(i => String.fromCharCode(i + 65 + 5))
				.join('');
		}
		
		// jurisdiction (fixedchar)
		{
			initialMessage.jurisdiction = [...Array(5).keys()]
				.map(i => String.fromCharCode(i + 65 + 6))
				.join('');
		}
		
		// contract_expiration (Timestamp)
		{
			initialMessage.contract_expiration = new Timestamp();
			
		}
		
		// contract_uri (varchar)
		{
			initialMessage.contract_uri = "Text 8"
		}
		
		// issuer (Entity)
		{
			initialMessage.issuer = new Entity();
			
		}
		
		// issuer_logo_url (varchar)
		{
			initialMessage.issuer_logo_url = "Text 10"
		}
		
		// contract_operator_included (bool)
		{
			initialMessage.contract_operator_included = true
		}
		
		// contract_operator (Entity)
		if (initialMessage.contract_operator_included) {
			initialMessage.contract_operator = new Entity();
			
		}
		
		// contract_auth_flags (varbin)
		{
			initialMessage.contract_auth_flags = Buffer.from([...Array(16).keys()]
				.map(i => i + 65 + 13));
		}
		
		// contract_fee (uint)
		{
			// IsNativeType
				{
			let type = 'uint64';
			if(type.startsWith('['))
				initialMessage.contract_fee = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.contract_fee = new BN(65 + 14);
			else
				initialMessage.contract_fee = 65 + 14;
				}
		}
		
		// voting_systems (VotingSystem[])
		{
			// IsInternalTypeArray
			initialMessage.voting_systems = [...Array(2)].map(() => new VotingSystem());
		}
		
		// restricted_qty_assets (uint)
		{
			// IsNativeType
				{
			let type = 'uint64';
			if(type.startsWith('['))
				initialMessage.restricted_qty_assets = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.restricted_qty_assets = new BN(65 + 16);
			else
				initialMessage.restricted_qty_assets = 65 + 16;
				}
		}
		
		// administration_proposal (bool)
		{
			initialMessage.administration_proposal = true
		}
		
		// holder_proposal (bool)
		{
			initialMessage.holder_proposal = true
		}
		
		// oracles (Oracle[])
		{
			// IsInternalTypeArray
			initialMessage.oracles = [...Array(2)].map(() => new Oracle());
		}
		
		// master_pkh (PublicKeyHash)
		{
			initialMessage.master_pkh = new PublicKeyHash();
			
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new ContractOffer();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// contract_name (varchar)
		if (initialMessage.contract_name !== (decodedMessage.contract_name)) {
			throw new Error(sprintf("contract_name doesn't match : %s != %s", initialMessage.contract_name, decodedMessage.contract_name));
		}
		// body_of_agreement_type (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.body_of_agreement_type);
		let decodedJson = JSON.stringify(decodedMessage.body_of_agreement_type);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("body_of_agreement_type doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// body_of_agreement (varbin)
		if ((initialMessage.body_of_agreement && decodedMessage.body_of_agreement) 
				&& !initialMessage.body_of_agreement.equals(decodedMessage.body_of_agreement)) {
			throw new Error(sprintf("body_of_agreement doesn't match : %x != %x", initialMessage.body_of_agreement, decodedMessage.body_of_agreement));
		}
		// contract_type (varchar)
		if (initialMessage.contract_type !== (decodedMessage.contract_type)) {
			throw new Error(sprintf("contract_type doesn't match : %s != %s", initialMessage.contract_type, decodedMessage.contract_type));
		}
		// supporting_docs (Document[])
		if ((initialMessage.supporting_docs && decodedMessage.supporting_docs) 
				&& initialMessage.supporting_docs.length != decodedMessage.supporting_docs.length) {
			throw new Error(sprintf("supporting_docs lengths don't match : %d != %d", initialMessage.supporting_docs.length, decodedMessage.supporting_docs.length));
		}
		// governing_law (fixedchar)
			// IsFixedChar
		if (initialMessage.governing_law != decodedMessage.governing_law) {
			throw new Error(sprintf("governing_law doesn't match : %s != %s", initialMessage.governing_law, decodedMessage.governing_law));
		}
		// jurisdiction (fixedchar)
			// IsFixedChar
		if (initialMessage.jurisdiction != decodedMessage.jurisdiction) {
			throw new Error(sprintf("jurisdiction doesn't match : %s != %s", initialMessage.jurisdiction, decodedMessage.jurisdiction));
		}
		// contract_expiration (Timestamp)
		// Timestamp test compare not setup
		// contract_uri (varchar)
		if (initialMessage.contract_uri !== (decodedMessage.contract_uri)) {
			throw new Error(sprintf("contract_uri doesn't match : %s != %s", initialMessage.contract_uri, decodedMessage.contract_uri));
		}
		// issuer (Entity)
		// Entity test compare not setup
		// issuer_logo_url (varchar)
		if (initialMessage.issuer_logo_url !== (decodedMessage.issuer_logo_url)) {
			throw new Error(sprintf("issuer_logo_url doesn't match : %s != %s", initialMessage.issuer_logo_url, decodedMessage.issuer_logo_url));
		}
		// contract_operator_included (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.contract_operator_included);
		let decodedJson = JSON.stringify(decodedMessage.contract_operator_included);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("contract_operator_included doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// contract_operator (Entity)
		// Entity test compare not setup
		// contract_auth_flags (varbin)
		if ((initialMessage.contract_auth_flags && decodedMessage.contract_auth_flags) 
				&& !initialMessage.contract_auth_flags.equals(decodedMessage.contract_auth_flags)) {
			throw new Error(sprintf("contract_auth_flags doesn't match : %x != %x", initialMessage.contract_auth_flags, decodedMessage.contract_auth_flags));
		}
		// contract_fee (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.contract_fee);
		let decodedJson = JSON.stringify(decodedMessage.contract_fee);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("contract_fee doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// voting_systems (VotingSystem[])
		if ((initialMessage.voting_systems && decodedMessage.voting_systems) 
				&& initialMessage.voting_systems.length != decodedMessage.voting_systems.length) {
			throw new Error(sprintf("voting_systems lengths don't match : %d != %d", initialMessage.voting_systems.length, decodedMessage.voting_systems.length));
		}
		// restricted_qty_assets (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.restricted_qty_assets);
		let decodedJson = JSON.stringify(decodedMessage.restricted_qty_assets);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("restricted_qty_assets doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// administration_proposal (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.administration_proposal);
		let decodedJson = JSON.stringify(decodedMessage.administration_proposal);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("administration_proposal doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// holder_proposal (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.holder_proposal);
		let decodedJson = JSON.stringify(decodedMessage.holder_proposal);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("holder_proposal doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// oracles (Oracle[])
		if ((initialMessage.oracles && decodedMessage.oracles) 
				&& initialMessage.oracles.length != decodedMessage.oracles.length) {
			throw new Error(sprintf("oracles lengths don't match : %d != %d", initialMessage.oracles.length, decodedMessage.oracles.length));
		}
		// master_pkh (PublicKeyHash)
		// PublicKeyHash test compare not setup


	});
});


describe('ContractFormation', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new ContractFormation();
		// contract_name (varchar)
		{
			initialMessage.contract_name = "Text 0"
		}
		
		// body_of_agreement_type (uint)
		{
				initialMessage.body_of_agreement_type = 1;
		}
		
		// body_of_agreement (varbin)
		{
			initialMessage.body_of_agreement = Buffer.from([...Array(32).keys()]
				.map(i => i + 65 + 2));
		}
		
		// contract_type (varchar)
		{
			initialMessage.contract_type = "Text 3"
		}
		
		// supporting_docs (Document[])
		{
			// IsInternalTypeArray
			initialMessage.supporting_docs = [...Array(2)].map(() => new Document());
		}
		
		// governing_law (fixedchar)
		{
			initialMessage.governing_law = [...Array(5).keys()]
				.map(i => String.fromCharCode(i + 65 + 5))
				.join('');
		}
		
		// jurisdiction (fixedchar)
		{
			initialMessage.jurisdiction = [...Array(5).keys()]
				.map(i => String.fromCharCode(i + 65 + 6))
				.join('');
		}
		
		// contract_expiration (Timestamp)
		{
			initialMessage.contract_expiration = new Timestamp();
			
		}
		
		// contract_uri (varchar)
		{
			initialMessage.contract_uri = "Text 8"
		}
		
		// issuer (Entity)
		{
			initialMessage.issuer = new Entity();
			
		}
		
		// issuer_logo_url (varchar)
		{
			initialMessage.issuer_logo_url = "Text 10"
		}
		
		// contract_operator_included (bool)
		{
			initialMessage.contract_operator_included = true
		}
		
		// contract_operator (Entity)
		if (initialMessage.contract_operator_included) {
			initialMessage.contract_operator = new Entity();
			
		}
		
		// contract_auth_flags (varbin)
		{
			initialMessage.contract_auth_flags = Buffer.from([...Array(16).keys()]
				.map(i => i + 65 + 13));
		}
		
		// contract_fee (uint)
		{
			// IsNativeType
				{
			let type = 'uint64';
			if(type.startsWith('['))
				initialMessage.contract_fee = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.contract_fee = new BN(65 + 14);
			else
				initialMessage.contract_fee = 65 + 14;
				}
		}
		
		// voting_systems (VotingSystem[])
		{
			// IsInternalTypeArray
			initialMessage.voting_systems = [...Array(2)].map(() => new VotingSystem());
		}
		
		// restricted_qty_assets (uint)
		{
			// IsNativeType
				{
			let type = 'uint64';
			if(type.startsWith('['))
				initialMessage.restricted_qty_assets = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.restricted_qty_assets = new BN(65 + 16);
			else
				initialMessage.restricted_qty_assets = 65 + 16;
				}
		}
		
		// administration_proposal (bool)
		{
			initialMessage.administration_proposal = true
		}
		
		// holder_proposal (bool)
		{
			initialMessage.holder_proposal = true
		}
		
		// oracles (Oracle[])
		{
			// IsInternalTypeArray
			initialMessage.oracles = [...Array(2)].map(() => new Oracle());
		}
		
		// master_pkh (PublicKeyHash)
		{
			initialMessage.master_pkh = new PublicKeyHash();
			
		}
		
		// contract_revision (uint)
		{
			// IsNativeType
				{
			let type = 'uint32';
			if(type.startsWith('['))
				initialMessage.contract_revision = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.contract_revision = new BN(65 + 21);
			else
				initialMessage.contract_revision = 65 + 21;
				}
		}
		
		// timestamp (Timestamp)
		{
			initialMessage.timestamp = new Timestamp();
			
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new ContractFormation();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// contract_name (varchar)
		if (initialMessage.contract_name !== (decodedMessage.contract_name)) {
			throw new Error(sprintf("contract_name doesn't match : %s != %s", initialMessage.contract_name, decodedMessage.contract_name));
		}
		// body_of_agreement_type (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.body_of_agreement_type);
		let decodedJson = JSON.stringify(decodedMessage.body_of_agreement_type);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("body_of_agreement_type doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// body_of_agreement (varbin)
		if ((initialMessage.body_of_agreement && decodedMessage.body_of_agreement) 
				&& !initialMessage.body_of_agreement.equals(decodedMessage.body_of_agreement)) {
			throw new Error(sprintf("body_of_agreement doesn't match : %x != %x", initialMessage.body_of_agreement, decodedMessage.body_of_agreement));
		}
		// contract_type (varchar)
		if (initialMessage.contract_type !== (decodedMessage.contract_type)) {
			throw new Error(sprintf("contract_type doesn't match : %s != %s", initialMessage.contract_type, decodedMessage.contract_type));
		}
		// supporting_docs (Document[])
		if ((initialMessage.supporting_docs && decodedMessage.supporting_docs) 
				&& initialMessage.supporting_docs.length != decodedMessage.supporting_docs.length) {
			throw new Error(sprintf("supporting_docs lengths don't match : %d != %d", initialMessage.supporting_docs.length, decodedMessage.supporting_docs.length));
		}
		// governing_law (fixedchar)
			// IsFixedChar
		if (initialMessage.governing_law != decodedMessage.governing_law) {
			throw new Error(sprintf("governing_law doesn't match : %s != %s", initialMessage.governing_law, decodedMessage.governing_law));
		}
		// jurisdiction (fixedchar)
			// IsFixedChar
		if (initialMessage.jurisdiction != decodedMessage.jurisdiction) {
			throw new Error(sprintf("jurisdiction doesn't match : %s != %s", initialMessage.jurisdiction, decodedMessage.jurisdiction));
		}
		// contract_expiration (Timestamp)
		// Timestamp test compare not setup
		// contract_uri (varchar)
		if (initialMessage.contract_uri !== (decodedMessage.contract_uri)) {
			throw new Error(sprintf("contract_uri doesn't match : %s != %s", initialMessage.contract_uri, decodedMessage.contract_uri));
		}
		// issuer (Entity)
		// Entity test compare not setup
		// issuer_logo_url (varchar)
		if (initialMessage.issuer_logo_url !== (decodedMessage.issuer_logo_url)) {
			throw new Error(sprintf("issuer_logo_url doesn't match : %s != %s", initialMessage.issuer_logo_url, decodedMessage.issuer_logo_url));
		}
		// contract_operator_included (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.contract_operator_included);
		let decodedJson = JSON.stringify(decodedMessage.contract_operator_included);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("contract_operator_included doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// contract_operator (Entity)
		// Entity test compare not setup
		// contract_auth_flags (varbin)
		if ((initialMessage.contract_auth_flags && decodedMessage.contract_auth_flags) 
				&& !initialMessage.contract_auth_flags.equals(decodedMessage.contract_auth_flags)) {
			throw new Error(sprintf("contract_auth_flags doesn't match : %x != %x", initialMessage.contract_auth_flags, decodedMessage.contract_auth_flags));
		}
		// contract_fee (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.contract_fee);
		let decodedJson = JSON.stringify(decodedMessage.contract_fee);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("contract_fee doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// voting_systems (VotingSystem[])
		if ((initialMessage.voting_systems && decodedMessage.voting_systems) 
				&& initialMessage.voting_systems.length != decodedMessage.voting_systems.length) {
			throw new Error(sprintf("voting_systems lengths don't match : %d != %d", initialMessage.voting_systems.length, decodedMessage.voting_systems.length));
		}
		// restricted_qty_assets (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.restricted_qty_assets);
		let decodedJson = JSON.stringify(decodedMessage.restricted_qty_assets);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("restricted_qty_assets doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// administration_proposal (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.administration_proposal);
		let decodedJson = JSON.stringify(decodedMessage.administration_proposal);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("administration_proposal doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// holder_proposal (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.holder_proposal);
		let decodedJson = JSON.stringify(decodedMessage.holder_proposal);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("holder_proposal doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// oracles (Oracle[])
		if ((initialMessage.oracles && decodedMessage.oracles) 
				&& initialMessage.oracles.length != decodedMessage.oracles.length) {
			throw new Error(sprintf("oracles lengths don't match : %d != %d", initialMessage.oracles.length, decodedMessage.oracles.length));
		}
		// master_pkh (PublicKeyHash)
		// PublicKeyHash test compare not setup
		// contract_revision (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.contract_revision);
		let decodedJson = JSON.stringify(decodedMessage.contract_revision);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("contract_revision doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// timestamp (Timestamp)
		// Timestamp test compare not setup


	});
});


describe('ContractAmendment', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new ContractAmendment();
		// change_administration_address (bool)
		{
			initialMessage.change_administration_address = true
		}
		
		// change_operator_address (bool)
		{
			initialMessage.change_operator_address = true
		}
		
		// contract_revision (uint)
		{
			// IsNativeType
				{
			let type = 'uint32';
			if(type.startsWith('['))
				initialMessage.contract_revision = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.contract_revision = new BN(65 + 2);
			else
				initialMessage.contract_revision = 65 + 2;
				}
		}
		
		// amendments (Amendment[])
		{
			// IsInternalTypeArray
			initialMessage.amendments = [...Array(2)].map(() => new Amendment());
		}
		
		// ref_tx_id (TxId)
		{
			initialMessage.ref_tx_id = new TxId();
			
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new ContractAmendment();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// change_administration_address (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.change_administration_address);
		let decodedJson = JSON.stringify(decodedMessage.change_administration_address);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("change_administration_address doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// change_operator_address (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.change_operator_address);
		let decodedJson = JSON.stringify(decodedMessage.change_operator_address);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("change_operator_address doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// contract_revision (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.contract_revision);
		let decodedJson = JSON.stringify(decodedMessage.contract_revision);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("contract_revision doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// amendments (Amendment[])
		if ((initialMessage.amendments && decodedMessage.amendments) 
				&& initialMessage.amendments.length != decodedMessage.amendments.length) {
			throw new Error(sprintf("amendments lengths don't match : %d != %d", initialMessage.amendments.length, decodedMessage.amendments.length));
		}
		// ref_tx_id (TxId)
		// TxId test compare not setup


	});
});


describe('StaticContractFormation', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new StaticContractFormation();
		// contract_name (varchar)
		{
			initialMessage.contract_name = "Text 0"
		}
		
		// contract_code (ContractCode)
		{
			initialMessage.contract_code = new ContractCode();
			
		}
		
		// body_of_agreement_type (uint)
		{
				initialMessage.body_of_agreement_type = 1;
		}
		
		// body_of_agreement (varbin)
		{
			initialMessage.body_of_agreement = Buffer.from([...Array(32).keys()]
				.map(i => i + 65 + 3));
		}
		
		// contract_type (varchar)
		{
			initialMessage.contract_type = "Text 4"
		}
		
		// supporting_docs (Document[])
		{
			// IsInternalTypeArray
			initialMessage.supporting_docs = [...Array(2)].map(() => new Document());
		}
		
		// contract_revision (uint)
		{
			// IsNativeType
				{
			let type = 'uint32';
			if(type.startsWith('['))
				initialMessage.contract_revision = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.contract_revision = new BN(65 + 6);
			else
				initialMessage.contract_revision = 65 + 6;
				}
		}
		
		// governing_law (fixedchar)
		{
			initialMessage.governing_law = [...Array(5).keys()]
				.map(i => String.fromCharCode(i + 65 + 7))
				.join('');
		}
		
		// jurisdiction (fixedchar)
		{
			initialMessage.jurisdiction = [...Array(5).keys()]
				.map(i => String.fromCharCode(i + 65 + 8))
				.join('');
		}
		
		// effective_date (Timestamp)
		{
			initialMessage.effective_date = new Timestamp();
			
		}
		
		// contract_expiration (Timestamp)
		{
			initialMessage.contract_expiration = new Timestamp();
			
		}
		
		// contract_uri (varchar)
		{
			initialMessage.contract_uri = "Text 11"
		}
		
		// prev_rev_tx_id (TxId)
		{
			initialMessage.prev_rev_tx_id = new TxId();
			
		}
		
		// entities (Entity[])
		{
			// IsInternalTypeArray
			initialMessage.entities = [...Array(2)].map(() => new Entity());
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new StaticContractFormation();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// contract_name (varchar)
		if (initialMessage.contract_name !== (decodedMessage.contract_name)) {
			throw new Error(sprintf("contract_name doesn't match : %s != %s", initialMessage.contract_name, decodedMessage.contract_name));
		}
		// contract_code (ContractCode)
		// ContractCode test compare not setup
		// body_of_agreement_type (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.body_of_agreement_type);
		let decodedJson = JSON.stringify(decodedMessage.body_of_agreement_type);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("body_of_agreement_type doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// body_of_agreement (varbin)
		if ((initialMessage.body_of_agreement && decodedMessage.body_of_agreement) 
				&& !initialMessage.body_of_agreement.equals(decodedMessage.body_of_agreement)) {
			throw new Error(sprintf("body_of_agreement doesn't match : %x != %x", initialMessage.body_of_agreement, decodedMessage.body_of_agreement));
		}
		// contract_type (varchar)
		if (initialMessage.contract_type !== (decodedMessage.contract_type)) {
			throw new Error(sprintf("contract_type doesn't match : %s != %s", initialMessage.contract_type, decodedMessage.contract_type));
		}
		// supporting_docs (Document[])
		if ((initialMessage.supporting_docs && decodedMessage.supporting_docs) 
				&& initialMessage.supporting_docs.length != decodedMessage.supporting_docs.length) {
			throw new Error(sprintf("supporting_docs lengths don't match : %d != %d", initialMessage.supporting_docs.length, decodedMessage.supporting_docs.length));
		}
		// contract_revision (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.contract_revision);
		let decodedJson = JSON.stringify(decodedMessage.contract_revision);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("contract_revision doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// governing_law (fixedchar)
			// IsFixedChar
		if (initialMessage.governing_law != decodedMessage.governing_law) {
			throw new Error(sprintf("governing_law doesn't match : %s != %s", initialMessage.governing_law, decodedMessage.governing_law));
		}
		// jurisdiction (fixedchar)
			// IsFixedChar
		if (initialMessage.jurisdiction != decodedMessage.jurisdiction) {
			throw new Error(sprintf("jurisdiction doesn't match : %s != %s", initialMessage.jurisdiction, decodedMessage.jurisdiction));
		}
		// effective_date (Timestamp)
		// Timestamp test compare not setup
		// contract_expiration (Timestamp)
		// Timestamp test compare not setup
		// contract_uri (varchar)
		if (initialMessage.contract_uri !== (decodedMessage.contract_uri)) {
			throw new Error(sprintf("contract_uri doesn't match : %s != %s", initialMessage.contract_uri, decodedMessage.contract_uri));
		}
		// prev_rev_tx_id (TxId)
		// TxId test compare not setup
		// entities (Entity[])
		if ((initialMessage.entities && decodedMessage.entities) 
				&& initialMessage.entities.length != decodedMessage.entities.length) {
			throw new Error(sprintf("entities lengths don't match : %d != %d", initialMessage.entities.length, decodedMessage.entities.length));
		}


	});
});


describe('ContractAddressChange', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new ContractAddressChange();
		// new_contract_pkh (PublicKeyHash)
		{
			initialMessage.new_contract_pkh = new PublicKeyHash();
			
		}
		
		// timestamp (Timestamp)
		{
			initialMessage.timestamp = new Timestamp();
			
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new ContractAddressChange();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// new_contract_pkh (PublicKeyHash)
		// PublicKeyHash test compare not setup
		// timestamp (Timestamp)
		// Timestamp test compare not setup


	});
});


describe('Order', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new Order();
		// compliance_action (fixedchar)
		{
			// IsNativeType
				{
			let type = 'byte';
			if(type.startsWith('['))
				initialMessage.compliance_action = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.compliance_action = new BN(65 + 0);
			else
				initialMessage.compliance_action = 65 + 0;
				}
		}
		
		// asset_type (fixedchar)
		if ( initialMessage.compliance_action == char('F') || initialMessage.compliance_action == char('C') || initialMessage.compliance_action == char('R')) {
			initialMessage.asset_type = [...Array(3).keys()]
				.map(i => String.fromCharCode(i + 65 + 1))
				.join('');
		}
		
		// asset_code (AssetCode)
		if ( initialMessage.compliance_action == char('F') || initialMessage.compliance_action == char('C') || initialMessage.compliance_action == char('R')) {
			initialMessage.asset_code = new AssetCode();
			
		}
		
		// target_addresses (TargetAddress[])
		if ( initialMessage.compliance_action == char('F') || initialMessage.compliance_action == char('C') || initialMessage.compliance_action == char('R')) {
			// IsInternalTypeArray
			initialMessage.target_addresses = [...Array(2)].map(() => new TargetAddress());
		}
		
		// freeze_tx_id (TxId)
		if ( initialMessage.compliance_action == char('T')) {
			initialMessage.freeze_tx_id = new TxId();
			
		}
		
		// freeze_period (Timestamp)
		if ( initialMessage.compliance_action == char('F')) {
			initialMessage.freeze_period = new Timestamp();
			
		}
		
		// deposit_address (PublicKeyHash)
		if ( initialMessage.compliance_action == char('C')) {
			initialMessage.deposit_address = new PublicKeyHash();
			
		}
		
		// authority_included (bool)
		if ( initialMessage.compliance_action == char('F') || initialMessage.compliance_action == char('T') || initialMessage.compliance_action == char('C')) {
			initialMessage.authority_included = true
		}
		
		// authority_name (varchar)
		if (initialMessage.authority_included) {
			initialMessage.authority_name = "Text 8"
		}
		
		// authority_public_key (varbin)
		if (initialMessage.authority_included) {
			initialMessage.authority_public_key = Buffer.from([...Array(8).keys()]
				.map(i => i + 65 + 9));
		}
		
		// signature_algorithm (uint)
		if (initialMessage.authority_included) {
				initialMessage.signature_algorithm = 1;
		}
		
		// order_signature (varbin)
		if (initialMessage.authority_included) {
			initialMessage.order_signature = Buffer.from([...Array(8).keys()]
				.map(i => i + 65 + 11));
		}
		
		// supporting_evidence_hash (bin)
		{
			// IsNativeType
				{
			let type = '[32]byte';
			if(type.startsWith('['))
				initialMessage.supporting_evidence_hash = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.supporting_evidence_hash = new BN(65 + 12);
			else
				initialMessage.supporting_evidence_hash = 65 + 12;
				}
		}
		
		// ref_txs (varbin)
		if ( initialMessage.compliance_action == char('R')) {
			initialMessage.ref_txs = Buffer.from([...Array(32).keys()]
				.map(i => i + 65 + 13));
		}
		
		// bitcoin_dispersions (QuantityIndex[])
		if ( initialMessage.compliance_action == char('R')) {
			// IsInternalTypeArray
			initialMessage.bitcoin_dispersions = [...Array(2)].map(() => new QuantityIndex());
		}
		
		// message (varchar)
		{
			initialMessage.message = "Text 15"
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new Order();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// compliance_action (fixedchar)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.compliance_action);
		let decodedJson = JSON.stringify(decodedMessage.compliance_action);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("compliance_action doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// asset_type (fixedchar)
			// IsFixedChar
		if (initialMessage.asset_type != decodedMessage.asset_type) {
			throw new Error(sprintf("asset_type doesn't match : %s != %s", initialMessage.asset_type, decodedMessage.asset_type));
		}
		// asset_code (AssetCode)
		// AssetCode test compare not setup
		// target_addresses (TargetAddress[])
		if ((initialMessage.target_addresses && decodedMessage.target_addresses) 
				&& initialMessage.target_addresses.length != decodedMessage.target_addresses.length) {
			throw new Error(sprintf("target_addresses lengths don't match : %d != %d", initialMessage.target_addresses.length, decodedMessage.target_addresses.length));
		}
		// freeze_tx_id (TxId)
		// TxId test compare not setup
		// freeze_period (Timestamp)
		// Timestamp test compare not setup
		// deposit_address (PublicKeyHash)
		// PublicKeyHash test compare not setup
		// authority_included (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.authority_included);
		let decodedJson = JSON.stringify(decodedMessage.authority_included);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("authority_included doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// authority_name (varchar)
		if (initialMessage.authority_name !== (decodedMessage.authority_name)) {
			throw new Error(sprintf("authority_name doesn't match : %s != %s", initialMessage.authority_name, decodedMessage.authority_name));
		}
		// authority_public_key (varbin)
		if ((initialMessage.authority_public_key && decodedMessage.authority_public_key) 
				&& !initialMessage.authority_public_key.equals(decodedMessage.authority_public_key)) {
			throw new Error(sprintf("authority_public_key doesn't match : %x != %x", initialMessage.authority_public_key, decodedMessage.authority_public_key));
		}
		// signature_algorithm (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.signature_algorithm);
		let decodedJson = JSON.stringify(decodedMessage.signature_algorithm);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("signature_algorithm doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// order_signature (varbin)
		if ((initialMessage.order_signature && decodedMessage.order_signature) 
				&& !initialMessage.order_signature.equals(decodedMessage.order_signature)) {
			throw new Error(sprintf("order_signature doesn't match : %x != %x", initialMessage.order_signature, decodedMessage.order_signature));
		}
		// supporting_evidence_hash (bin)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.supporting_evidence_hash);
		let decodedJson = JSON.stringify(decodedMessage.supporting_evidence_hash);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("supporting_evidence_hash doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// ref_txs (varbin)
		if ((initialMessage.ref_txs && decodedMessage.ref_txs) 
				&& !initialMessage.ref_txs.equals(decodedMessage.ref_txs)) {
			throw new Error(sprintf("ref_txs doesn't match : %x != %x", initialMessage.ref_txs, decodedMessage.ref_txs));
		}
		// bitcoin_dispersions (QuantityIndex[])
		if ((initialMessage.bitcoin_dispersions && decodedMessage.bitcoin_dispersions) 
				&& initialMessage.bitcoin_dispersions.length != decodedMessage.bitcoin_dispersions.length) {
			throw new Error(sprintf("bitcoin_dispersions lengths don't match : %d != %d", initialMessage.bitcoin_dispersions.length, decodedMessage.bitcoin_dispersions.length));
		}
		// message (varchar)
		if (initialMessage.message !== (decodedMessage.message)) {
			throw new Error(sprintf("message doesn't match : %s != %s", initialMessage.message, decodedMessage.message));
		}


	});
});


describe('Freeze', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new Freeze();
		// asset_type (fixedchar)
		{
			initialMessage.asset_type = [...Array(3).keys()]
				.map(i => String.fromCharCode(i + 65 + 0))
				.join('');
		}
		
		// asset_code (AssetCode)
		{
			initialMessage.asset_code = new AssetCode();
			
		}
		
		// quantities (QuantityIndex[])
		{
			// IsInternalTypeArray
			initialMessage.quantities = [...Array(2)].map(() => new QuantityIndex());
		}
		
		// freeze_period (Timestamp)
		{
			initialMessage.freeze_period = new Timestamp();
			
		}
		
		// timestamp (Timestamp)
		{
			initialMessage.timestamp = new Timestamp();
			
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new Freeze();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// asset_type (fixedchar)
			// IsFixedChar
		if (initialMessage.asset_type != decodedMessage.asset_type) {
			throw new Error(sprintf("asset_type doesn't match : %s != %s", initialMessage.asset_type, decodedMessage.asset_type));
		}
		// asset_code (AssetCode)
		// AssetCode test compare not setup
		// quantities (QuantityIndex[])
		if ((initialMessage.quantities && decodedMessage.quantities) 
				&& initialMessage.quantities.length != decodedMessage.quantities.length) {
			throw new Error(sprintf("quantities lengths don't match : %d != %d", initialMessage.quantities.length, decodedMessage.quantities.length));
		}
		// freeze_period (Timestamp)
		// Timestamp test compare not setup
		// timestamp (Timestamp)
		// Timestamp test compare not setup


	});
});


describe('Thaw', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new Thaw();
		// freeze_tx_id (TxId)
		{
			initialMessage.freeze_tx_id = new TxId();
			
		}
		
		// timestamp (Timestamp)
		{
			initialMessage.timestamp = new Timestamp();
			
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new Thaw();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// freeze_tx_id (TxId)
		// TxId test compare not setup
		// timestamp (Timestamp)
		// Timestamp test compare not setup


	});
});


describe('Confiscation', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new Confiscation();
		// asset_type (fixedchar)
		{
			initialMessage.asset_type = [...Array(3).keys()]
				.map(i => String.fromCharCode(i + 65 + 0))
				.join('');
		}
		
		// asset_code (AssetCode)
		{
			initialMessage.asset_code = new AssetCode();
			
		}
		
		// quantities (QuantityIndex[])
		{
			// IsInternalTypeArray
			initialMessage.quantities = [...Array(2)].map(() => new QuantityIndex());
		}
		
		// deposit_qty (uint)
		{
			// IsNativeType
				{
			let type = 'uint64';
			if(type.startsWith('['))
				initialMessage.deposit_qty = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.deposit_qty = new BN(65 + 3);
			else
				initialMessage.deposit_qty = 65 + 3;
				}
		}
		
		// timestamp (Timestamp)
		{
			initialMessage.timestamp = new Timestamp();
			
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new Confiscation();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// asset_type (fixedchar)
			// IsFixedChar
		if (initialMessage.asset_type != decodedMessage.asset_type) {
			throw new Error(sprintf("asset_type doesn't match : %s != %s", initialMessage.asset_type, decodedMessage.asset_type));
		}
		// asset_code (AssetCode)
		// AssetCode test compare not setup
		// quantities (QuantityIndex[])
		if ((initialMessage.quantities && decodedMessage.quantities) 
				&& initialMessage.quantities.length != decodedMessage.quantities.length) {
			throw new Error(sprintf("quantities lengths don't match : %d != %d", initialMessage.quantities.length, decodedMessage.quantities.length));
		}
		// deposit_qty (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.deposit_qty);
		let decodedJson = JSON.stringify(decodedMessage.deposit_qty);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("deposit_qty doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// timestamp (Timestamp)
		// Timestamp test compare not setup


	});
});


describe('Reconciliation', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new Reconciliation();
		// asset_type (fixedchar)
		{
			initialMessage.asset_type = [...Array(3).keys()]
				.map(i => String.fromCharCode(i + 65 + 0))
				.join('');
		}
		
		// asset_code (AssetCode)
		{
			initialMessage.asset_code = new AssetCode();
			
		}
		
		// quantities (QuantityIndex[])
		{
			// IsInternalTypeArray
			initialMessage.quantities = [...Array(2)].map(() => new QuantityIndex());
		}
		
		// timestamp (Timestamp)
		{
			initialMessage.timestamp = new Timestamp();
			
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new Reconciliation();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// asset_type (fixedchar)
			// IsFixedChar
		if (initialMessage.asset_type != decodedMessage.asset_type) {
			throw new Error(sprintf("asset_type doesn't match : %s != %s", initialMessage.asset_type, decodedMessage.asset_type));
		}
		// asset_code (AssetCode)
		// AssetCode test compare not setup
		// quantities (QuantityIndex[])
		if ((initialMessage.quantities && decodedMessage.quantities) 
				&& initialMessage.quantities.length != decodedMessage.quantities.length) {
			throw new Error(sprintf("quantities lengths don't match : %d != %d", initialMessage.quantities.length, decodedMessage.quantities.length));
		}
		// timestamp (Timestamp)
		// Timestamp test compare not setup


	});
});


describe('Proposal', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new Proposal();
		// initiator (uint)
		{
				initialMessage.initiator = 1;
		}
		
		// asset_specific_vote (bool)
		{
			initialMessage.asset_specific_vote = true
		}
		
		// asset_type (fixedchar)
		if (initialMessage.asset_specific_vote) {
			initialMessage.asset_type = [...Array(3).keys()]
				.map(i => String.fromCharCode(i + 65 + 2))
				.join('');
		}
		
		// asset_code (AssetCode)
		if (initialMessage.asset_specific_vote) {
			initialMessage.asset_code = new AssetCode();
			
		}
		
		// vote_system (uint)
		{
			// IsNativeType
				{
			let type = 'uint8';
			if(type.startsWith('['))
				initialMessage.vote_system = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.vote_system = new BN(65 + 4);
			else
				initialMessage.vote_system = 65 + 4;
				}
		}
		
		// specific (bool)
		{
			initialMessage.specific = true
		}
		
		// proposed_amendments (Amendment[])
		if (initialMessage.specific) {
			// IsInternalTypeArray
			initialMessage.proposed_amendments = [...Array(2)].map(() => new Amendment());
		}
		
		// vote_options (varchar)
		{
			initialMessage.vote_options = "Text 7"
		}
		
		// vote_max (uint)
		{
			// IsNativeType
				{
			let type = 'uint8';
			if(type.startsWith('['))
				initialMessage.vote_max = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.vote_max = new BN(65 + 8);
			else
				initialMessage.vote_max = 65 + 8;
				}
		}
		
		// proposal_description (varchar)
		{
			initialMessage.proposal_description = "Text 9"
		}
		
		// proposal_document_hash (bin)
		{
			// IsNativeType
				{
			let type = '[32]byte';
			if(type.startsWith('['))
				initialMessage.proposal_document_hash = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.proposal_document_hash = new BN(65 + 10);
			else
				initialMessage.proposal_document_hash = 65 + 10;
				}
		}
		
		// vote_cut_off_timestamp (Timestamp)
		{
			initialMessage.vote_cut_off_timestamp = new Timestamp();
			
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new Proposal();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// initiator (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.initiator);
		let decodedJson = JSON.stringify(decodedMessage.initiator);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("initiator doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// asset_specific_vote (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.asset_specific_vote);
		let decodedJson = JSON.stringify(decodedMessage.asset_specific_vote);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("asset_specific_vote doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// asset_type (fixedchar)
			// IsFixedChar
		if (initialMessage.asset_type != decodedMessage.asset_type) {
			throw new Error(sprintf("asset_type doesn't match : %s != %s", initialMessage.asset_type, decodedMessage.asset_type));
		}
		// asset_code (AssetCode)
		// AssetCode test compare not setup
		// vote_system (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.vote_system);
		let decodedJson = JSON.stringify(decodedMessage.vote_system);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("vote_system doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// specific (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.specific);
		let decodedJson = JSON.stringify(decodedMessage.specific);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("specific doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// proposed_amendments (Amendment[])
		if ((initialMessage.proposed_amendments && decodedMessage.proposed_amendments) 
				&& initialMessage.proposed_amendments.length != decodedMessage.proposed_amendments.length) {
			throw new Error(sprintf("proposed_amendments lengths don't match : %d != %d", initialMessage.proposed_amendments.length, decodedMessage.proposed_amendments.length));
		}
		// vote_options (varchar)
		if (initialMessage.vote_options !== (decodedMessage.vote_options)) {
			throw new Error(sprintf("vote_options doesn't match : %s != %s", initialMessage.vote_options, decodedMessage.vote_options));
		}
		// vote_max (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.vote_max);
		let decodedJson = JSON.stringify(decodedMessage.vote_max);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("vote_max doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// proposal_description (varchar)
		if (initialMessage.proposal_description !== (decodedMessage.proposal_description)) {
			throw new Error(sprintf("proposal_description doesn't match : %s != %s", initialMessage.proposal_description, decodedMessage.proposal_description));
		}
		// proposal_document_hash (bin)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.proposal_document_hash);
		let decodedJson = JSON.stringify(decodedMessage.proposal_document_hash);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("proposal_document_hash doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// vote_cut_off_timestamp (Timestamp)
		// Timestamp test compare not setup


	});
});


describe('Vote', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new Vote();
		// timestamp (Timestamp)
		{
			initialMessage.timestamp = new Timestamp();
			
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new Vote();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// timestamp (Timestamp)
		// Timestamp test compare not setup


	});
});


describe('BallotCast', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new BallotCast();
		// vote_tx_id (TxId)
		{
			initialMessage.vote_tx_id = new TxId();
			
		}
		
		// vote (varchar)
		{
			initialMessage.vote = "Text 1"
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new BallotCast();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// vote_tx_id (TxId)
		// TxId test compare not setup
		// vote (varchar)
		if (initialMessage.vote !== (decodedMessage.vote)) {
			throw new Error(sprintf("vote doesn't match : %s != %s", initialMessage.vote, decodedMessage.vote));
		}


	});
});


describe('BallotCounted', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new BallotCounted();
		// vote_tx_id (TxId)
		{
			initialMessage.vote_tx_id = new TxId();
			
		}
		
		// vote (varchar)
		{
			initialMessage.vote = "Text 1"
		}
		
		// quantity (uint)
		{
			// IsNativeType
				{
			let type = 'uint64';
			if(type.startsWith('['))
				initialMessage.quantity = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.quantity = new BN(65 + 2);
			else
				initialMessage.quantity = 65 + 2;
				}
		}
		
		// timestamp (Timestamp)
		{
			initialMessage.timestamp = new Timestamp();
			
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new BallotCounted();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// vote_tx_id (TxId)
		// TxId test compare not setup
		// vote (varchar)
		if (initialMessage.vote !== (decodedMessage.vote)) {
			throw new Error(sprintf("vote doesn't match : %s != %s", initialMessage.vote, decodedMessage.vote));
		}
		// quantity (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.quantity);
		let decodedJson = JSON.stringify(decodedMessage.quantity);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("quantity doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// timestamp (Timestamp)
		// Timestamp test compare not setup


	});
});


describe('Result', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new Result();
		// asset_specific_vote (bool)
		{
			initialMessage.asset_specific_vote = true
		}
		
		// asset_type (fixedchar)
		if (initialMessage.asset_specific_vote) {
			initialMessage.asset_type = [...Array(3).keys()]
				.map(i => String.fromCharCode(i + 65 + 1))
				.join('');
		}
		
		// asset_code (AssetCode)
		if (initialMessage.asset_specific_vote) {
			initialMessage.asset_code = new AssetCode();
			
		}
		
		// specific (bool)
		{
			initialMessage.specific = true
		}
		
		// proposed_amendments (Amendment[])
		if (initialMessage.specific) {
			// IsInternalTypeArray
			initialMessage.proposed_amendments = [...Array(2)].map(() => new Amendment());
		}
		
		// vote_tx_id (TxId)
		{
			initialMessage.vote_tx_id = new TxId();
			
		}
		
		// option_tally (uint64[])
		{
			// IsNativeTypeArray
			initialMessage.option_tally = [...Array(5)].map(() => getArrayOrType('uint64'));
		}
		
		// result (varchar)
		{
			initialMessage.result = "Text 7"
		}
		
		// timestamp (Timestamp)
		{
			initialMessage.timestamp = new Timestamp();
			
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new Result();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// asset_specific_vote (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.asset_specific_vote);
		let decodedJson = JSON.stringify(decodedMessage.asset_specific_vote);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("asset_specific_vote doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// asset_type (fixedchar)
			// IsFixedChar
		if (initialMessage.asset_type != decodedMessage.asset_type) {
			throw new Error(sprintf("asset_type doesn't match : %s != %s", initialMessage.asset_type, decodedMessage.asset_type));
		}
		// asset_code (AssetCode)
		// AssetCode test compare not setup
		// specific (bool)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.specific);
		let decodedJson = JSON.stringify(decodedMessage.specific);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("specific doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// proposed_amendments (Amendment[])
		if ((initialMessage.proposed_amendments && decodedMessage.proposed_amendments) 
				&& initialMessage.proposed_amendments.length != decodedMessage.proposed_amendments.length) {
			throw new Error(sprintf("proposed_amendments lengths don't match : %d != %d", initialMessage.proposed_amendments.length, decodedMessage.proposed_amendments.length));
		}
		// vote_tx_id (TxId)
		// TxId test compare not setup
		// option_tally (uint64[])
			{
		let initialJson = JSON.stringify(initialMessage.option_tally);
		let decodedJson = JSON.stringify(decodedMessage.option_tally);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("option_tally doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// result (varchar)
		if (initialMessage.result !== (decodedMessage.result)) {
			throw new Error(sprintf("result doesn't match : %s != %s", initialMessage.result, decodedMessage.result));
		}
		// timestamp (Timestamp)
		// Timestamp test compare not setup


	});
});


describe('Message', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new Message();
		// address_indexes (uint16[])
		{
			// IsNativeTypeArray
			initialMessage.address_indexes = [...Array(5)].map(() => getArrayOrType('uint16'));
		}
		
		// message_type (uint16)
		{
			// IsNativeType
				{
			let type = 'uint16';
			if(type.startsWith('['))
				initialMessage.message_type = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.message_type = new BN(65 + 1);
			else
				initialMessage.message_type = 65 + 1;
				}
		}
		
		// message_payload (varbin)
		{
			initialMessage.message_payload = Buffer.from([...Array(32).keys()]
				.map(i => i + 65 + 2));
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new Message();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// address_indexes (uint16[])
			{
		let initialJson = JSON.stringify(initialMessage.address_indexes);
		let decodedJson = JSON.stringify(decodedMessage.address_indexes);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("address_indexes doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// message_type (uint16)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.message_type);
		let decodedJson = JSON.stringify(decodedMessage.message_type);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("message_type doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// message_payload (varbin)
		if ((initialMessage.message_payload && decodedMessage.message_payload) 
				&& !initialMessage.message_payload.equals(decodedMessage.message_payload)) {
			throw new Error(sprintf("message_payload doesn't match : %x != %x", initialMessage.message_payload, decodedMessage.message_payload));
		}


	});
});


describe('Rejection', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new Rejection();
		// address_indexes (uint16[])
		{
			// IsNativeTypeArray
			initialMessage.address_indexes = [...Array(5)].map(() => getArrayOrType('uint16'));
		}
		
		// reject_address_index (uint)
		{
			// IsNativeType
				{
			let type = 'uint16';
			if(type.startsWith('['))
				initialMessage.reject_address_index = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.reject_address_index = new BN(65 + 1);
			else
				initialMessage.reject_address_index = 65 + 1;
				}
		}
		
		// rejection_code (RejectionCode)
		{
			// IsNativeType
				{
			let type = 'uint8';
			if(type.startsWith('['))
				initialMessage.rejection_code = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.rejection_code = new BN(65 + 2);
			else
				initialMessage.rejection_code = 65 + 2;
				}
		}
		
		// message (varchar)
		{
			initialMessage.message = "Text 3"
		}
		
		// timestamp (Timestamp)
		{
			initialMessage.timestamp = new Timestamp();
			
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new Rejection();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// address_indexes (uint16[])
			{
		let initialJson = JSON.stringify(initialMessage.address_indexes);
		let decodedJson = JSON.stringify(decodedMessage.address_indexes);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("address_indexes doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// reject_address_index (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.reject_address_index);
		let decodedJson = JSON.stringify(decodedMessage.reject_address_index);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("reject_address_index doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// rejection_code (RejectionCode)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.rejection_code);
		let decodedJson = JSON.stringify(decodedMessage.rejection_code);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("rejection_code doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// message (varchar)
		if (initialMessage.message !== (decodedMessage.message)) {
			throw new Error(sprintf("message doesn't match : %s != %s", initialMessage.message, decodedMessage.message));
		}
		// timestamp (Timestamp)
		// Timestamp test compare not setup


	});
});


describe('Establishment', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new Establishment();
		// message (varchar)
		{
			initialMessage.message = "Text 0"
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new Establishment();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// message (varchar)
		if (initialMessage.message !== (decodedMessage.message)) {
			throw new Error(sprintf("message doesn't match : %s != %s", initialMessage.message, decodedMessage.message));
		}


	});
});


describe('Addition', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new Addition();
		// message (varchar)
		{
			initialMessage.message = "Text 0"
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new Addition();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// message (varchar)
		if (initialMessage.message !== (decodedMessage.message)) {
			throw new Error(sprintf("message doesn't match : %s != %s", initialMessage.message, decodedMessage.message));
		}


	});
});


describe('Alteration', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new Alteration();
		// entry_tx_id (TxId)
		{
			initialMessage.entry_tx_id = new TxId();
			
		}
		
		// message (varchar)
		{
			initialMessage.message = "Text 1"
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new Alteration();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// entry_tx_id (TxId)
		// TxId test compare not setup
		// message (varchar)
		if (initialMessage.message !== (decodedMessage.message)) {
			throw new Error(sprintf("message doesn't match : %s != %s", initialMessage.message, decodedMessage.message));
		}


	});
});


describe('Removal', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new Removal();
		// entry_tx_id (TxId)
		{
			initialMessage.entry_tx_id = new TxId();
			
		}
		
		// message (varchar)
		{
			initialMessage.message = "Text 1"
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new Removal();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// entry_tx_id (TxId)
		// TxId test compare not setup
		// message (varchar)
		if (initialMessage.message !== (decodedMessage.message)) {
			throw new Error(sprintf("message doesn't match : %s != %s", initialMessage.message, decodedMessage.message));
		}


	});
});


describe('Transfer', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new Transfer();
		// assets (AssetTransfer[])
		{
			// IsInternalTypeArray
			initialMessage.assets = [...Array(2)].map(() => new AssetTransfer());
		}
		
		// offer_expiry (Timestamp)
		{
			initialMessage.offer_expiry = new Timestamp();
			
		}
		
		// exchange_fee (uint)
		{
			// IsNativeType
				{
			let type = 'uint64';
			if(type.startsWith('['))
				initialMessage.exchange_fee = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.exchange_fee = new BN(65 + 2);
			else
				initialMessage.exchange_fee = 65 + 2;
				}
		}
		
		// exchange_fee_address (PublicKeyHash)
		{
			initialMessage.exchange_fee_address = new PublicKeyHash();
			
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new Transfer();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// assets (AssetTransfer[])
		if ((initialMessage.assets && decodedMessage.assets) 
				&& initialMessage.assets.length != decodedMessage.assets.length) {
			throw new Error(sprintf("assets lengths don't match : %d != %d", initialMessage.assets.length, decodedMessage.assets.length));
		}
		// offer_expiry (Timestamp)
		// Timestamp test compare not setup
		// exchange_fee (uint)
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.exchange_fee);
		let decodedJson = JSON.stringify(decodedMessage.exchange_fee);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("exchange_fee doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
		// exchange_fee_address (PublicKeyHash)
		// PublicKeyHash test compare not setup


	});
});


describe('Settlement', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new Settlement();
		// assets (AssetSettlement[])
		{
			// IsInternalTypeArray
			initialMessage.assets = [...Array(2)].map(() => new AssetSettlement());
		}
		
		// timestamp (Timestamp)
		{
			initialMessage.timestamp = new Timestamp();
			
		}
		

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new Settlement();

		let n = decodedMessage.write(initialEncoding);
		let err = decodedMessage.Validate();
		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// assets (AssetSettlement[])
		if ((initialMessage.assets && decodedMessage.assets) 
				&& initialMessage.assets.length != decodedMessage.assets.length) {
			throw new Error(sprintf("assets lengths don't match : %d != %d", initialMessage.assets.length, decodedMessage.assets.length));
		}
		// timestamp (Timestamp)
		// Timestamp test compare not setup


	});
});


