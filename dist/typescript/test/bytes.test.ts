import * as mocha from 'mocha';
import * as chai from 'chai';
import { read } from '../src/bytes';
import R from 'ramda';
import _ from '@keyring/util';
const expect = chai.expect;
[mocha]

describe('RejectionCodes', () => {
	it('should load codes', () => {
		let size = 1;
		let buf = new _.Reader();
		const type = '[][3]byte'.slice(2);
		let trade_restrictions = [...Array(size)].map(() => read(buf, type));
	});
});
