import * as mocha from 'mocha';
import * as chai from 'chai';
import { Resources } from '../src/resources';
import R from 'ramda';
const expect = chai.expect;
[mocha]

describe('RejectionCodes', () => {
	it('should load codes', () => {
		const rejectionCodes = Resources.GetRejectionCodes();
		console.log('Loaded %d rejection codes.', R.keys(rejectionCodes).length);

		const msgMalformed = Resources.GetRejectionCode(1);
		console.log('MsgMalformed :', msgMalformed);
		expect(msgMalformed.name).to.eql('MsgMalformed');
	});
});

describe('GetCurrencies', () => {
	it('should load codes', () => {
		const currencies = Resources.GetCurrencies();
		console.log('Loaded %d currencies.', R.keys(currencies).length);

		const usd = Resources.GetCurrency('USD');
		console.log('usd :', usd);
		expect(usd.name).to.eql('UnitedStatesDollar');
	});
});

describe('Entities', () => {
	it('should load codes', () => {
		const entities = Resources.GetEntityTypes();
		console.log('Loaded %d entities.', R.keys(entities).length);

		const lp = Resources.GetEntityType('L'.charCodeAt(0));
		console.log('lp :', lp);
		expect(lp.name).to.eql('LimitedPartnership');
	});
});

describe('PolityTypes', () => {
	it('should load codes', () => {
		const polityTypes = Resources.GetPolityTypes();
		console.log('Loaded %d polityTypes.', R.keys(polityTypes).length);
		// console.log(polityTypes);

		const aus = Resources.GetPolityType('AUS');
		console.log('aus :', aus);
		expect(aus.name).to.eql('Australia');
	});
});

describe('RoleTypes', () => {
	it('should load codes', () => {
		const roleTypes = Resources.GetRoleTypes();
		console.log('Loaded %d roleTypes.', R.keys(roleTypes).length);
		// console.log(roleTypes);

		const ceo = Resources.GetRoleType(5);
		console.log('ceo :', ceo);
		expect(ceo.label).to.eql('CEO');
	});
});

describe('GetTagTypes', () => {
	it('should load codes', () => {
		const tagTypes = Resources.GetTagTypes();
		console.log('Loaded %d roleTypes.', R.keys(tagTypes).length);
		// console.log(tagTypes);

		const promotional = Resources.GetTagType(121);
		console.log('promotional :', promotional);
		expect(promotional.label).to.eql('Promotional');
	});
});

