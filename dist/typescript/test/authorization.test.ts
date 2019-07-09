import {
	ContractOffer, OpReturnMessage, Permission, Timestamp, Entity, PublicKeyHash
} from '../index';
import * as mocha from 'mocha';
import * as chai from 'chai';
const expect = chai.expect;
[mocha];

const getRandomPerms = () => {
	const rbool = () => Math.random() >= .5;
	let voteSystemsAllowed = [ rbool(), rbool() ];

	let permission = new Permission();
	permission.permitted = rbool();
	permission.administration_proposal = rbool();
	permission.holder_proposal = rbool();
	if (permission.administration_proposal || permission.holder_proposal)
		permission.voting_systems_allowed = voteSystemsAllowed;
	return permission;
}
describe('Permissions', () => {
	it('should encode and decode', () => {

		// Note: Permissions can be different for each field.
		let permissions = [...Array(20)].map(_ => getRandomPerms());

		// Serialize auth flags
		let authFlags = Permission.WriteAuthFlags(permissions);

		// Generate a new contract offer action
		let action = new ContractOffer();
		action.contract_name = 'Test';
		action.body_of_agreement_type = 2;
		action.body_of_agreement = [...Buffer.from('<contract agreement>', 'ascii')];
		action.contract_type = 'Test Type';
		action.contract_auth_flags = authFlags;
		action.supporting_docs = [];
		action.contract_expiration = new Timestamp();
		action.issuer = new Entity();
		action.voting_systems = [];
		action.oracles = [];
		action.master_pkh = new PublicKeyHash();
		// Specify any other fields necessary
		// ...

		// Serialize action
		const isTest = true // use "test.tokenized" OP_RETURN signature
		let actionData = OpReturnMessage.Serialize(action, isTest);

		// Convert to hex
		console.log("Contract Offer Hex :\n", actionData);

		// Output as human readable string
		console.log("Contract Offer text : \n", JSON.stringify(action));

		let msg = <ContractOffer>OpReturnMessage.Deserialize(actionData, isTest);
		console.log("orig contract_auth_flags: \n", msg.contract_auth_flags);
		console.log("read contract_auth_flags: \n", msg.contract_auth_flags);
		let recoded = OpReturnMessage.Serialize(action, isTest);
		expect(recoded).to.eql(actionData);

		let readPerms = Permission.ReadAuthFlags(msg.contract_auth_flags, 20, 2);
		console.log("permissions: \n", permissions);
		console.log("readPerms: \n", readPerms);
		expect(readPerms).to.eql(permissions);
	});
});

