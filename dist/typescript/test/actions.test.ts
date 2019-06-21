import {sprintf} from 'sprintf-js';
import BN from 'bn.js';
import * as mocha from 'mocha';
import * as chai from 'chai';
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

const getArray = (type: string) => {
	const regex = /\[(\w+)\]/m;
	let m;
	if ((m = regex.exec(type)) !== null) {
		console.log('m:',  m[1]);
		const subtype = type.slice(m[0].length);
		return [...Array(parseInt(m[1], 10))].map(() => 0);
	}
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
			initialMessage.trade_restrictions = [...Array(5)].map(() => getArray('[3]byte'));
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
				initialMessage.vote_multiplier = getArray(type);
			else
				initialMessage.vote_multiplier = 1;
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
			// IsNativeType
				{
			let type = 'uint8';
			if(type.startsWith('['))
				initialMessage.asset_modification_governance = getArray(type);
			else
				initialMessage.asset_modification_governance = 1;
				}
		}
		
		// token_qty (uint)
		{
			// IsNativeType
				{
			let type = 'uint64';
			if(type.startsWith('['))
				initialMessage.token_qty = getArray(type);
			else
				initialMessage.token_qty = 1;
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
		if (!initialMessage.asset_auth_flags.equals(decodedMessage.asset_auth_flags)) {
			throw new Error(sprintf("asset_auth_flags doesn't match : %x != %x", initialMessage.asset_auth_flags, decodedMessage.asset_auth_flags));
		}
		// transfers_permitted (bool)
			// IsNativeType
		if (initialMessage.transfers_permitted != decodedMessage.transfers_permitted) {
			throw new Error(sprintf("transfers_permitted doesn't match : %v != %v", initialMessage.transfers_permitted, decodedMessage.transfers_permitted));
		}
		// trade_restrictions (Polity[])
			{
		let initialJson = JSON.stringify(initialMessage.trade_restrictions);
		let decodedJson = JSON.stringify(decodedMessage.trade_restrictions);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("trade_restrictions doesn't match : %s != %s", i, initialJson, decodedJson));
		}
			}
		// enforcement_orders_permitted (bool)
			// IsNativeType
		if (initialMessage.enforcement_orders_permitted != decodedMessage.enforcement_orders_permitted) {
			throw new Error(sprintf("enforcement_orders_permitted doesn't match : %v != %v", initialMessage.enforcement_orders_permitted, decodedMessage.enforcement_orders_permitted));
		}
		// voting_rights (bool)
			// IsNativeType
		if (initialMessage.voting_rights != decodedMessage.voting_rights) {
			throw new Error(sprintf("voting_rights doesn't match : %v != %v", initialMessage.voting_rights, decodedMessage.voting_rights));
		}
		// vote_multiplier (uint)
			// IsNativeType
		if (initialMessage.vote_multiplier != decodedMessage.vote_multiplier) {
			throw new Error(sprintf("vote_multiplier doesn't match : %v != %v", initialMessage.vote_multiplier, decodedMessage.vote_multiplier));
		}
		// administration_proposal (bool)
			// IsNativeType
		if (initialMessage.administration_proposal != decodedMessage.administration_proposal) {
			throw new Error(sprintf("administration_proposal doesn't match : %v != %v", initialMessage.administration_proposal, decodedMessage.administration_proposal));
		}
		// holder_proposal (bool)
			// IsNativeType
		if (initialMessage.holder_proposal != decodedMessage.holder_proposal) {
			throw new Error(sprintf("holder_proposal doesn't match : %v != %v", initialMessage.holder_proposal, decodedMessage.holder_proposal));
		}
		// asset_modification_governance (uint)
			// IsNativeType
		if (initialMessage.asset_modification_governance != decodedMessage.asset_modification_governance) {
			throw new Error(sprintf("asset_modification_governance doesn't match : %v != %v", initialMessage.asset_modification_governance, decodedMessage.asset_modification_governance));
		}
		// token_qty (uint)
			// IsNativeType
		if (initialMessage.token_qty != decodedMessage.token_qty) {
			throw new Error(sprintf("token_qty doesn't match : %v != %v", initialMessage.token_qty, decodedMessage.token_qty));
		}
		// asset_payload (varbin)
		if (!initialMessage.asset_payload.equals(decodedMessage.asset_payload)) {
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
				initialMessage.asset_index = getArray(type);
			else
				initialMessage.asset_index = 1;
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
			initialMessage.trade_restrictions = [...Array(5)].map(() => getArray('[3]byte'));
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
				initialMessage.vote_multiplier = getArray(type);
			else
				initialMessage.vote_multiplier = 1;
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
			// IsNativeType
				{
			let type = 'uint8';
			if(type.startsWith('['))
				initialMessage.asset_modification_governance = getArray(type);
			else
				initialMessage.asset_modification_governance = 1;
				}
		}
		
		// token_qty (uint)
		{
			// IsNativeType
				{
			let type = 'uint64';
			if(type.startsWith('['))
				initialMessage.token_qty = getArray(type);
			else
				initialMessage.token_qty = 1;
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
				initialMessage.asset_revision = getArray(type);
			else
				initialMessage.asset_revision = 1;
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
		if (initialMessage.asset_index != decodedMessage.asset_index) {
			throw new Error(sprintf("asset_index doesn't match : %v != %v", initialMessage.asset_index, decodedMessage.asset_index));
		}
		// asset_auth_flags (varbin)
		if (!initialMessage.asset_auth_flags.equals(decodedMessage.asset_auth_flags)) {
			throw new Error(sprintf("asset_auth_flags doesn't match : %x != %x", initialMessage.asset_auth_flags, decodedMessage.asset_auth_flags));
		}
		// transfers_permitted (bool)
			// IsNativeType
		if (initialMessage.transfers_permitted != decodedMessage.transfers_permitted) {
			throw new Error(sprintf("transfers_permitted doesn't match : %v != %v", initialMessage.transfers_permitted, decodedMessage.transfers_permitted));
		}
		// trade_restrictions (Polity[])
			{
		let initialJson = JSON.stringify(initialMessage.trade_restrictions);
		let decodedJson = JSON.stringify(decodedMessage.trade_restrictions);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("trade_restrictions doesn't match : %s != %s", i, initialJson, decodedJson));
		}
			}
		// enforcement_orders_permitted (bool)
			// IsNativeType
		if (initialMessage.enforcement_orders_permitted != decodedMessage.enforcement_orders_permitted) {
			throw new Error(sprintf("enforcement_orders_permitted doesn't match : %v != %v", initialMessage.enforcement_orders_permitted, decodedMessage.enforcement_orders_permitted));
		}
		// voting_rights (bool)
			// IsNativeType
		if (initialMessage.voting_rights != decodedMessage.voting_rights) {
			throw new Error(sprintf("voting_rights doesn't match : %v != %v", initialMessage.voting_rights, decodedMessage.voting_rights));
		}
		// vote_multiplier (uint)
			// IsNativeType
		if (initialMessage.vote_multiplier != decodedMessage.vote_multiplier) {
			throw new Error(sprintf("vote_multiplier doesn't match : %v != %v", initialMessage.vote_multiplier, decodedMessage.vote_multiplier));
		}
		// administration_proposal (bool)
			// IsNativeType
		if (initialMessage.administration_proposal != decodedMessage.administration_proposal) {
			throw new Error(sprintf("administration_proposal doesn't match : %v != %v", initialMessage.administration_proposal, decodedMessage.administration_proposal));
		}
		// holder_proposal (bool)
			// IsNativeType
		if (initialMessage.holder_proposal != decodedMessage.holder_proposal) {
			throw new Error(sprintf("holder_proposal doesn't match : %v != %v", initialMessage.holder_proposal, decodedMessage.holder_proposal));
		}
		// asset_modification_governance (uint)
			// IsNativeType
		if (initialMessage.asset_modification_governance != decodedMessage.asset_modification_governance) {
			throw new Error(sprintf("asset_modification_governance doesn't match : %v != %v", initialMessage.asset_modification_governance, decodedMessage.asset_modification_governance));
		}
		// token_qty (uint)
			// IsNativeType
		if (initialMessage.token_qty != decodedMessage.token_qty) {
			throw new Error(sprintf("token_qty doesn't match : %v != %v", initialMessage.token_qty, decodedMessage.token_qty));
		}
		// asset_payload (varbin)
		if (!initialMessage.asset_payload.equals(decodedMessage.asset_payload)) {
			throw new Error(sprintf("asset_payload doesn't match : %x != %x", initialMessage.asset_payload, decodedMessage.asset_payload));
		}
		// asset_revision (uint)
			// IsNativeType
		if (initialMessage.asset_revision != decodedMessage.asset_revision) {
			throw new Error(sprintf("asset_revision doesn't match : %v != %v", initialMessage.asset_revision, decodedMessage.asset_revision));
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
				initialMessage.asset_revision = getArray(type);
			else
				initialMessage.asset_revision = 1;
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
		if (initialMessage.asset_revision != decodedMessage.asset_revision) {
			throw new Error(sprintf("asset_revision doesn't match : %v != %v", initialMessage.asset_revision, decodedMessage.asset_revision));
		}
		// amendments (Amendment[])
		if (initialMessage.amendments.length != decodedMessage.amendments.length) {
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
			// IsNativeType
				{
			let type = 'uint8';
			if(type.startsWith('['))
				initialMessage.body_of_agreement_type = getArray(type);
			else
				initialMessage.body_of_agreement_type = 1;
				}
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
				initialMessage.contract_fee = getArray(type);
			else
				initialMessage.contract_fee = 1;
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
				initialMessage.restricted_qty_assets = getArray(type);
			else
				initialMessage.restricted_qty_assets = 1;
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

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// contract_name (varchar)
		if (!initialMessage.contract_name.equals(decodedMessage.contract_name)) {
			throw new Error(sprintf("contract_name doesn't match : %s != %s", initialMessage.contract_name, decodedMessage.contract_name));
		}
		// body_of_agreement_type (uint)
			// IsNativeType
		if (initialMessage.body_of_agreement_type != decodedMessage.body_of_agreement_type) {
			throw new Error(sprintf("body_of_agreement_type doesn't match : %v != %v", initialMessage.body_of_agreement_type, decodedMessage.body_of_agreement_type));
		}
		// body_of_agreement (varbin)
		if (!initialMessage.body_of_agreement.equals(decodedMessage.body_of_agreement)) {
			throw new Error(sprintf("body_of_agreement doesn't match : %x != %x", initialMessage.body_of_agreement, decodedMessage.body_of_agreement));
		}
		// contract_type (varchar)
		if (!initialMessage.contract_type.equals(decodedMessage.contract_type)) {
			throw new Error(sprintf("contract_type doesn't match : %s != %s", initialMessage.contract_type, decodedMessage.contract_type));
		}
		// supporting_docs (Document[])
		if (initialMessage.supporting_docs.length != decodedMessage.supporting_docs.length) {
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
		if (!initialMessage.contract_uri.equals(decodedMessage.contract_uri)) {
			throw new Error(sprintf("contract_uri doesn't match : %s != %s", initialMessage.contract_uri, decodedMessage.contract_uri));
		}
		// issuer (Entity)
		// Entity test compare not setup
		// issuer_logo_url (varchar)
		if (!initialMessage.issuer_logo_url.equals(decodedMessage.issuer_logo_url)) {
			throw new Error(sprintf("issuer_logo_url doesn't match : %s != %s", initialMessage.issuer_logo_url, decodedMessage.issuer_logo_url));
		}
		// contract_operator_included (bool)
			// IsNativeType
		if (initialMessage.contract_operator_included != decodedMessage.contract_operator_included) {
			throw new Error(sprintf("contract_operator_included doesn't match : %v != %v", initialMessage.contract_operator_included, decodedMessage.contract_operator_included));
		}
		// contract_operator (Entity)
		// Entity test compare not setup
		// contract_auth_flags (varbin)
		if (!initialMessage.contract_auth_flags.equals(decodedMessage.contract_auth_flags)) {
			throw new Error(sprintf("contract_auth_flags doesn't match : %x != %x", initialMessage.contract_auth_flags, decodedMessage.contract_auth_flags));
		}
		// contract_fee (uint)
			// IsNativeType
		if (initialMessage.contract_fee != decodedMessage.contract_fee) {
			throw new Error(sprintf("contract_fee doesn't match : %v != %v", initialMessage.contract_fee, decodedMessage.contract_fee));
		}
		// voting_systems (VotingSystem[])
		if (initialMessage.voting_systems.length != decodedMessage.voting_systems.length) {
			throw new Error(sprintf("voting_systems lengths don't match : %d != %d", initialMessage.voting_systems.length, decodedMessage.voting_systems.length));
		}
		// restricted_qty_assets (uint)
			// IsNativeType
		if (initialMessage.restricted_qty_assets != decodedMessage.restricted_qty_assets) {
			throw new Error(sprintf("restricted_qty_assets doesn't match : %v != %v", initialMessage.restricted_qty_assets, decodedMessage.restricted_qty_assets));
		}
		// administration_proposal (bool)
			// IsNativeType
		if (initialMessage.administration_proposal != decodedMessage.administration_proposal) {
			throw new Error(sprintf("administration_proposal doesn't match : %v != %v", initialMessage.administration_proposal, decodedMessage.administration_proposal));
		}
		// holder_proposal (bool)
			// IsNativeType
		if (initialMessage.holder_proposal != decodedMessage.holder_proposal) {
			throw new Error(sprintf("holder_proposal doesn't match : %v != %v", initialMessage.holder_proposal, decodedMessage.holder_proposal));
		}
		// oracles (Oracle[])
		if (initialMessage.oracles.length != decodedMessage.oracles.length) {
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
			// IsNativeType
				{
			let type = 'uint8';
			if(type.startsWith('['))
				initialMessage.body_of_agreement_type = getArray(type);
			else
				initialMessage.body_of_agreement_type = 1;
				}
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
				initialMessage.contract_fee = getArray(type);
			else
				initialMessage.contract_fee = 1;
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
				initialMessage.restricted_qty_assets = getArray(type);
			else
				initialMessage.restricted_qty_assets = 1;
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
				initialMessage.contract_revision = getArray(type);
			else
				initialMessage.contract_revision = 1;
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

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// contract_name (varchar)
		if (!initialMessage.contract_name.equals(decodedMessage.contract_name)) {
			throw new Error(sprintf("contract_name doesn't match : %s != %s", initialMessage.contract_name, decodedMessage.contract_name));
		}
		// body_of_agreement_type (uint)
			// IsNativeType
		if (initialMessage.body_of_agreement_type != decodedMessage.body_of_agreement_type) {
			throw new Error(sprintf("body_of_agreement_type doesn't match : %v != %v", initialMessage.body_of_agreement_type, decodedMessage.body_of_agreement_type));
		}
		// body_of_agreement (varbin)
		if (!initialMessage.body_of_agreement.equals(decodedMessage.body_of_agreement)) {
			throw new Error(sprintf("body_of_agreement doesn't match : %x != %x", initialMessage.body_of_agreement, decodedMessage.body_of_agreement));
		}
		// contract_type (varchar)
		if (!initialMessage.contract_type.equals(decodedMessage.contract_type)) {
			throw new Error(sprintf("contract_type doesn't match : %s != %s", initialMessage.contract_type, decodedMessage.contract_type));
		}
		// supporting_docs (Document[])
		if (initialMessage.supporting_docs.length != decodedMessage.supporting_docs.length) {
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
		if (!initialMessage.contract_uri.equals(decodedMessage.contract_uri)) {
			throw new Error(sprintf("contract_uri doesn't match : %s != %s", initialMessage.contract_uri, decodedMessage.contract_uri));
		}
		// issuer (Entity)
		// Entity test compare not setup
		// issuer_logo_url (varchar)
		if (!initialMessage.issuer_logo_url.equals(decodedMessage.issuer_logo_url)) {
			throw new Error(sprintf("issuer_logo_url doesn't match : %s != %s", initialMessage.issuer_logo_url, decodedMessage.issuer_logo_url));
		}
		// contract_operator_included (bool)
			// IsNativeType
		if (initialMessage.contract_operator_included != decodedMessage.contract_operator_included) {
			throw new Error(sprintf("contract_operator_included doesn't match : %v != %v", initialMessage.contract_operator_included, decodedMessage.contract_operator_included));
		}
		// contract_operator (Entity)
		// Entity test compare not setup
		// contract_auth_flags (varbin)
		if (!initialMessage.contract_auth_flags.equals(decodedMessage.contract_auth_flags)) {
			throw new Error(sprintf("contract_auth_flags doesn't match : %x != %x", initialMessage.contract_auth_flags, decodedMessage.contract_auth_flags));
		}
		// contract_fee (uint)
			// IsNativeType
		if (initialMessage.contract_fee != decodedMessage.contract_fee) {
			throw new Error(sprintf("contract_fee doesn't match : %v != %v", initialMessage.contract_fee, decodedMessage.contract_fee));
		}
		// voting_systems (VotingSystem[])
		if (initialMessage.voting_systems.length != decodedMessage.voting_systems.length) {
			throw new Error(sprintf("voting_systems lengths don't match : %d != %d", initialMessage.voting_systems.length, decodedMessage.voting_systems.length));
		}
		// restricted_qty_assets (uint)
			// IsNativeType
		if (initialMessage.restricted_qty_assets != decodedMessage.restricted_qty_assets) {
			throw new Error(sprintf("restricted_qty_assets doesn't match : %v != %v", initialMessage.restricted_qty_assets, decodedMessage.restricted_qty_assets));
		}
		// administration_proposal (bool)
			// IsNativeType
		if (initialMessage.administration_proposal != decodedMessage.administration_proposal) {
			throw new Error(sprintf("administration_proposal doesn't match : %v != %v", initialMessage.administration_proposal, decodedMessage.administration_proposal));
		}
		// holder_proposal (bool)
			// IsNativeType
		if (initialMessage.holder_proposal != decodedMessage.holder_proposal) {
			throw new Error(sprintf("holder_proposal doesn't match : %v != %v", initialMessage.holder_proposal, decodedMessage.holder_proposal));
		}
		// oracles (Oracle[])
		if (initialMessage.oracles.length != decodedMessage.oracles.length) {
			throw new Error(sprintf("oracles lengths don't match : %d != %d", initialMessage.oracles.length, decodedMessage.oracles.length));
		}
		// master_pkh (PublicKeyHash)
		// PublicKeyHash test compare not setup
		// contract_revision (uint)
			// IsNativeType
		if (initialMessage.contract_revision != decodedMessage.contract_revision) {
			throw new Error(sprintf("contract_revision doesn't match : %v != %v", initialMessage.contract_revision, decodedMessage.contract_revision));
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
				initialMessage.contract_revision = getArray(type);
			else
				initialMessage.contract_revision = 1;
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
		if (initialMessage.change_administration_address != decodedMessage.change_administration_address) {
			throw new Error(sprintf("change_administration_address doesn't match : %v != %v", initialMessage.change_administration_address, decodedMessage.change_administration_address));
		}
		// change_operator_address (bool)
			// IsNativeType
		if (initialMessage.change_operator_address != decodedMessage.change_operator_address) {
			throw new Error(sprintf("change_operator_address doesn't match : %v != %v", initialMessage.change_operator_address, decodedMessage.change_operator_address));
		}
		// contract_revision (uint)
			// IsNativeType
		if (initialMessage.contract_revision != decodedMessage.contract_revision) {
			throw new Error(sprintf("contract_revision doesn't match : %v != %v", initialMessage.contract_revision, decodedMessage.contract_revision));
		}
		// amendments (Amendment[])
		if (initialMessage.amendments.length != decodedMessage.amendments.length) {
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
			// IsNativeType
				{
			let type = 'uint8';
			if(type.startsWith('['))
				initialMessage.body_of_agreement_type = getArray(type);
			else
				initialMessage.body_of_agreement_type = 1;
				}
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
				initialMessage.contract_revision = getArray(type);
			else
				initialMessage.contract_revision = 1;
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

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// contract_name (varchar)
		if (!initialMessage.contract_name.equals(decodedMessage.contract_name)) {
			throw new Error(sprintf("contract_name doesn't match : %s != %s", initialMessage.contract_name, decodedMessage.contract_name));
		}
		// contract_code (ContractCode)
		// ContractCode test compare not setup
		// body_of_agreement_type (uint)
			// IsNativeType
		if (initialMessage.body_of_agreement_type != decodedMessage.body_of_agreement_type) {
			throw new Error(sprintf("body_of_agreement_type doesn't match : %v != %v", initialMessage.body_of_agreement_type, decodedMessage.body_of_agreement_type));
		}
		// body_of_agreement (varbin)
		if (!initialMessage.body_of_agreement.equals(decodedMessage.body_of_agreement)) {
			throw new Error(sprintf("body_of_agreement doesn't match : %x != %x", initialMessage.body_of_agreement, decodedMessage.body_of_agreement));
		}
		// contract_type (varchar)
		if (!initialMessage.contract_type.equals(decodedMessage.contract_type)) {
			throw new Error(sprintf("contract_type doesn't match : %s != %s", initialMessage.contract_type, decodedMessage.contract_type));
		}
		// supporting_docs (Document[])
		if (initialMessage.supporting_docs.length != decodedMessage.supporting_docs.length) {
			throw new Error(sprintf("supporting_docs lengths don't match : %d != %d", initialMessage.supporting_docs.length, decodedMessage.supporting_docs.length));
		}
		// contract_revision (uint)
			// IsNativeType
		if (initialMessage.contract_revision != decodedMessage.contract_revision) {
			throw new Error(sprintf("contract_revision doesn't match : %v != %v", initialMessage.contract_revision, decodedMessage.contract_revision));
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
		if (!initialMessage.contract_uri.equals(decodedMessage.contract_uri)) {
			throw new Error(sprintf("contract_uri doesn't match : %s != %s", initialMessage.contract_uri, decodedMessage.contract_uri));
		}
		// prev_rev_tx_id (TxId)
		// TxId test compare not setup
		// entities (Entity[])
		if (initialMessage.entities.length != decodedMessage.entities.length) {
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
				initialMessage.compliance_action = getArray(type);
			else
				initialMessage.compliance_action = 1;
				}
		}
		
		// asset_type (fixedchar)
		if ( initialMessage.compliance_action == 'F' || initialMessage.compliance_action == 'C' || initialMessage.compliance_action == 'R') {
			initialMessage.asset_type = [...Array(3).keys()]
				.map(i => String.fromCharCode(i + 65 + 1))
				.join('');
		}
		
		// asset_code (AssetCode)
		if ( initialMessage.compliance_action == 'F' || initialMessage.compliance_action == 'C' || initialMessage.compliance_action == 'R') {
			initialMessage.asset_code = new AssetCode();
		}
		
		// target_addresses (TargetAddress[])
		if ( initialMessage.compliance_action == 'F' || initialMessage.compliance_action == 'C' || initialMessage.compliance_action == 'R') {
			// IsInternalTypeArray
			initialMessage.target_addresses = [...Array(2)].map(() => new TargetAddress());
		}
		
		// freeze_tx_id (TxId)
		if ( initialMessage.compliance_action == 'T') {
			initialMessage.freeze_tx_id = new TxId();
		}
		
		// freeze_period (Timestamp)
		if ( initialMessage.compliance_action == 'F') {
			initialMessage.freeze_period = new Timestamp();
		}
		
		// deposit_address (PublicKeyHash)
		if ( initialMessage.compliance_action == 'C') {
			initialMessage.deposit_address = new PublicKeyHash();
		}
		
		// authority_included (bool)
		if ( initialMessage.compliance_action == 'F' || initialMessage.compliance_action == 'T' || initialMessage.compliance_action == 'C') {
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
			// IsNativeType
				{
			let type = 'uint8';
			if(type.startsWith('['))
				initialMessage.signature_algorithm = getArray(type);
			else
				initialMessage.signature_algorithm = 1;
				}
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
				initialMessage.supporting_evidence_hash = getArray(type);
			else
				initialMessage.supporting_evidence_hash = 1;
				}
		}
		
		// ref_txs (varbin)
		if ( initialMessage.compliance_action == 'R') {
			initialMessage.ref_txs = Buffer.from([...Array(32).keys()]
				.map(i => i + 65 + 13));
		}
		
		// bitcoin_dispersions (QuantityIndex[])
		if ( initialMessage.compliance_action == 'R') {
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
		if (initialMessage.compliance_action != decodedMessage.compliance_action) {
			throw new Error(sprintf("compliance_action doesn't match : %v != %v", initialMessage.compliance_action, decodedMessage.compliance_action));
		}
		// asset_type (fixedchar)
			// IsFixedChar
		if (initialMessage.asset_type != decodedMessage.asset_type) {
			throw new Error(sprintf("asset_type doesn't match : %s != %s", initialMessage.asset_type, decodedMessage.asset_type));
		}
		// asset_code (AssetCode)
		// AssetCode test compare not setup
		// target_addresses (TargetAddress[])
		if (initialMessage.target_addresses.length != decodedMessage.target_addresses.length) {
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
		if (initialMessage.authority_included != decodedMessage.authority_included) {
			throw new Error(sprintf("authority_included doesn't match : %v != %v", initialMessage.authority_included, decodedMessage.authority_included));
		}
		// authority_name (varchar)
		if (!initialMessage.authority_name.equals(decodedMessage.authority_name)) {
			throw new Error(sprintf("authority_name doesn't match : %s != %s", initialMessage.authority_name, decodedMessage.authority_name));
		}
		// authority_public_key (varbin)
		if (!initialMessage.authority_public_key.equals(decodedMessage.authority_public_key)) {
			throw new Error(sprintf("authority_public_key doesn't match : %x != %x", initialMessage.authority_public_key, decodedMessage.authority_public_key));
		}
		// signature_algorithm (uint)
			// IsNativeType
		if (initialMessage.signature_algorithm != decodedMessage.signature_algorithm) {
			throw new Error(sprintf("signature_algorithm doesn't match : %v != %v", initialMessage.signature_algorithm, decodedMessage.signature_algorithm));
		}
		// order_signature (varbin)
		if (!initialMessage.order_signature.equals(decodedMessage.order_signature)) {
			throw new Error(sprintf("order_signature doesn't match : %x != %x", initialMessage.order_signature, decodedMessage.order_signature));
		}
		// supporting_evidence_hash (bin)
			// IsNativeType
		if (initialMessage.supporting_evidence_hash != decodedMessage.supporting_evidence_hash) {
			throw new Error(sprintf("supporting_evidence_hash doesn't match : %v != %v", initialMessage.supporting_evidence_hash, decodedMessage.supporting_evidence_hash));
		}
		// ref_txs (varbin)
		if (!initialMessage.ref_txs.equals(decodedMessage.ref_txs)) {
			throw new Error(sprintf("ref_txs doesn't match : %x != %x", initialMessage.ref_txs, decodedMessage.ref_txs));
		}
		// bitcoin_dispersions (QuantityIndex[])
		if (initialMessage.bitcoin_dispersions.length != decodedMessage.bitcoin_dispersions.length) {
			throw new Error(sprintf("bitcoin_dispersions lengths don't match : %d != %d", initialMessage.bitcoin_dispersions.length, decodedMessage.bitcoin_dispersions.length));
		}
		// message (varchar)
		if (!initialMessage.message.equals(decodedMessage.message)) {
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
		if (initialMessage.quantities.length != decodedMessage.quantities.length) {
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
				initialMessage.deposit_qty = getArray(type);
			else
				initialMessage.deposit_qty = 1;
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
		if (initialMessage.quantities.length != decodedMessage.quantities.length) {
			throw new Error(sprintf("quantities lengths don't match : %d != %d", initialMessage.quantities.length, decodedMessage.quantities.length));
		}
		// deposit_qty (uint)
			// IsNativeType
		if (initialMessage.deposit_qty != decodedMessage.deposit_qty) {
			throw new Error(sprintf("deposit_qty doesn't match : %v != %v", initialMessage.deposit_qty, decodedMessage.deposit_qty));
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
		if (initialMessage.quantities.length != decodedMessage.quantities.length) {
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
			// IsNativeType
				{
			let type = 'uint8';
			if(type.startsWith('['))
				initialMessage.initiator = getArray(type);
			else
				initialMessage.initiator = 1;
				}
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
				initialMessage.vote_system = getArray(type);
			else
				initialMessage.vote_system = 1;
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
				initialMessage.vote_max = getArray(type);
			else
				initialMessage.vote_max = 1;
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
				initialMessage.proposal_document_hash = getArray(type);
			else
				initialMessage.proposal_document_hash = 1;
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
		if (initialMessage.initiator != decodedMessage.initiator) {
			throw new Error(sprintf("initiator doesn't match : %v != %v", initialMessage.initiator, decodedMessage.initiator));
		}
		// asset_specific_vote (bool)
			// IsNativeType
		if (initialMessage.asset_specific_vote != decodedMessage.asset_specific_vote) {
			throw new Error(sprintf("asset_specific_vote doesn't match : %v != %v", initialMessage.asset_specific_vote, decodedMessage.asset_specific_vote));
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
		if (initialMessage.vote_system != decodedMessage.vote_system) {
			throw new Error(sprintf("vote_system doesn't match : %v != %v", initialMessage.vote_system, decodedMessage.vote_system));
		}
		// specific (bool)
			// IsNativeType
		if (initialMessage.specific != decodedMessage.specific) {
			throw new Error(sprintf("specific doesn't match : %v != %v", initialMessage.specific, decodedMessage.specific));
		}
		// proposed_amendments (Amendment[])
		if (initialMessage.proposed_amendments.length != decodedMessage.proposed_amendments.length) {
			throw new Error(sprintf("proposed_amendments lengths don't match : %d != %d", initialMessage.proposed_amendments.length, decodedMessage.proposed_amendments.length));
		}
		// vote_options (varchar)
		if (!initialMessage.vote_options.equals(decodedMessage.vote_options)) {
			throw new Error(sprintf("vote_options doesn't match : %s != %s", initialMessage.vote_options, decodedMessage.vote_options));
		}
		// vote_max (uint)
			// IsNativeType
		if (initialMessage.vote_max != decodedMessage.vote_max) {
			throw new Error(sprintf("vote_max doesn't match : %v != %v", initialMessage.vote_max, decodedMessage.vote_max));
		}
		// proposal_description (varchar)
		if (!initialMessage.proposal_description.equals(decodedMessage.proposal_description)) {
			throw new Error(sprintf("proposal_description doesn't match : %s != %s", initialMessage.proposal_description, decodedMessage.proposal_description));
		}
		// proposal_document_hash (bin)
			// IsNativeType
		if (initialMessage.proposal_document_hash != decodedMessage.proposal_document_hash) {
			throw new Error(sprintf("proposal_document_hash doesn't match : %v != %v", initialMessage.proposal_document_hash, decodedMessage.proposal_document_hash));
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
		if (!initialMessage.vote.equals(decodedMessage.vote)) {
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
				initialMessage.quantity = getArray(type);
			else
				initialMessage.quantity = 1;
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
		if (!initialMessage.vote.equals(decodedMessage.vote)) {
			throw new Error(sprintf("vote doesn't match : %s != %s", initialMessage.vote, decodedMessage.vote));
		}
		// quantity (uint)
			// IsNativeType
		if (initialMessage.quantity != decodedMessage.quantity) {
			throw new Error(sprintf("quantity doesn't match : %v != %v", initialMessage.quantity, decodedMessage.quantity));
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
			initialMessage.option_tally = [...Array(5)].map(() => getArray('uint64'));
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
		if (initialMessage.asset_specific_vote != decodedMessage.asset_specific_vote) {
			throw new Error(sprintf("asset_specific_vote doesn't match : %v != %v", initialMessage.asset_specific_vote, decodedMessage.asset_specific_vote));
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
		if (initialMessage.specific != decodedMessage.specific) {
			throw new Error(sprintf("specific doesn't match : %v != %v", initialMessage.specific, decodedMessage.specific));
		}
		// proposed_amendments (Amendment[])
		if (initialMessage.proposed_amendments.length != decodedMessage.proposed_amendments.length) {
			throw new Error(sprintf("proposed_amendments lengths don't match : %d != %d", initialMessage.proposed_amendments.length, decodedMessage.proposed_amendments.length));
		}
		// vote_tx_id (TxId)
		// TxId test compare not setup
		// option_tally (uint64[])
			{
		let initialJson = JSON.stringify(initialMessage.option_tally);
		let decodedJson = JSON.stringify(decodedMessage.option_tally);
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("option_tally doesn't match : %s != %s", i, initialJson, decodedJson));
		}
			}
		// result (varchar)
		if (!initialMessage.result.equals(decodedMessage.result)) {
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
			initialMessage.address_indexes = [...Array(5)].map(() => getArray('uint16'));
		}
		
		// message_type (uint16)
		{
			// IsNativeType
				{
			let type = 'uint16';
			if(type.startsWith('['))
				initialMessage.message_type = getArray(type);
			else
				initialMessage.message_type = 1;
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
			throw new Error(sprintf("address_indexes doesn't match : %s != %s", i, initialJson, decodedJson));
		}
			}
		// message_type (uint16)
			// IsNativeType
		if (initialMessage.message_type != decodedMessage.message_type) {
			throw new Error(sprintf("message_type doesn't match : %v != %v", initialMessage.message_type, decodedMessage.message_type));
		}
		// message_payload (varbin)
		if (!initialMessage.message_payload.equals(decodedMessage.message_payload)) {
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
			initialMessage.address_indexes = [...Array(5)].map(() => getArray('uint16'));
		}
		
		// reject_address_index (uint)
		{
			// IsNativeType
				{
			let type = 'uint16';
			if(type.startsWith('['))
				initialMessage.reject_address_index = getArray(type);
			else
				initialMessage.reject_address_index = 1;
				}
		}
		
		// rejection_code (RejectionCode)
		{
			// IsNativeType
				{
			let type = 'uint8';
			if(type.startsWith('['))
				initialMessage.rejection_code = getArray(type);
			else
				initialMessage.rejection_code = 1;
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
			throw new Error(sprintf("address_indexes doesn't match : %s != %s", i, initialJson, decodedJson));
		}
			}
		// reject_address_index (uint)
			// IsNativeType
		if (initialMessage.reject_address_index != decodedMessage.reject_address_index) {
			throw new Error(sprintf("reject_address_index doesn't match : %v != %v", initialMessage.reject_address_index, decodedMessage.reject_address_index));
		}
		// rejection_code (RejectionCode)
			// IsNativeType
		if (initialMessage.rejection_code != decodedMessage.rejection_code) {
			throw new Error(sprintf("rejection_code doesn't match : %v != %v", initialMessage.rejection_code, decodedMessage.rejection_code));
		}
		// message (varchar)
		if (!initialMessage.message.equals(decodedMessage.message)) {
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

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// message (varchar)
		if (!initialMessage.message.equals(decodedMessage.message)) {
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

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// message (varchar)
		if (!initialMessage.message.equals(decodedMessage.message)) {
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
		if (!initialMessage.message.equals(decodedMessage.message)) {
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
		if (!initialMessage.message.equals(decodedMessage.message)) {
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
				initialMessage.exchange_fee = getArray(type);
			else
				initialMessage.exchange_fee = 1;
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

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// assets (AssetTransfer[])
		if (initialMessage.assets.length != decodedMessage.assets.length) {
			throw new Error(sprintf("assets lengths don't match : %d != %d", initialMessage.assets.length, decodedMessage.assets.length));
		}
		// offer_expiry (Timestamp)
		// Timestamp test compare not setup
		// exchange_fee (uint)
			// IsNativeType
		if (initialMessage.exchange_fee != decodedMessage.exchange_fee) {
			throw new Error(sprintf("exchange_fee doesn't match : %v != %v", initialMessage.exchange_fee, decodedMessage.exchange_fee));
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

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		// assets (AssetSettlement[])
		if (initialMessage.assets.length != decodedMessage.assets.length) {
			throw new Error(sprintf("assets lengths don't match : %d != %d", initialMessage.assets.length, decodedMessage.assets.length));
		}
		// timestamp (Timestamp)
		// Timestamp test compare not setup


	});
});


