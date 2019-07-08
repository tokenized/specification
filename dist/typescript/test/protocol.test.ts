import * as mocha from 'mocha';
import * as chai from 'chai';
import { OpReturnMessage } from '../src/protocol';
import { AssetDefinition } from '../src/actions';
import { AssetTypeMapping, ShareCommon } from '../src/assets';
const expect = chai.expect;
[mocha]

const sample = '6a0e746573742e746f6b656e697a65644c5700413153484305924924924001000000000000000040420f0000000000380000000000000000000041434d45000000000000000000000000001c0041434d4520436f6d6d6f6e20536861726573202854657374696e6729';

describe('deserialize', () => {
	it('should read opreturn', () => {
		const msg = <AssetDefinition>OpReturnMessage.Deserialize(Buffer.from(sample, 'hex'), true);
		console.log(msg);
		const asset = <ShareCommon>AssetTypeMapping(msg.asset_type);
		asset.Write(msg.asset_payload);
		console.log(asset);

		expect(asset.description).to.eql('ACME Common Shares (Testing)');

		let restr = OpReturnMessage.Serialize(msg, true).toString('hex');

		expect(restr).to.eql(sample);
	});
});
