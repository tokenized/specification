import * as mocha from 'mocha';
import * as chai from 'chai';
import { OpReturnMessage } from '../src/protocol';
import { PublicKeyHash } from '../src/protocol_types';
import { AssetDefinition } from '../src/actions';
import { AssetTypeMapping, ShareCommon } from '../src/assets';
const expect = chai.expect;
[mocha]

describe('deserialize', () => {
	it('should read opreturn', () => {
		const sample = '6a0e746573742e746f6b656e697a65644c5700413153484305924924924001000000000000000040420f0000000000380000000000000000000041434d45000000000000000000000000001c0041434d4520436f6d6d6f6e20536861726573202854657374696e6729';
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

describe('bitdb', () => {
	it('should read opreturn', () => {
		const sample = '005432010000435552cc9eda8da0c23cbf1ee0851422e5c0e71eada8ce6246fefd07e602963504bd1d020000f0b9f5050000000001001027000000000000c9c5ad979b1cae15';
		const msg = OpReturnMessage.getMsg(Buffer.from(sample, 'hex'));
		console.log(JSON.stringify(msg, null, 2));
	});
});

describe('PublicKeyHash', () => {
	it('should decode and encode', () => {
		const sample = '17JXjVqzQgZfFZASp3FAY3wcikLp6CAY7h';
		const hash = PublicKeyHash.fromStr(sample);
		console.log(JSON.stringify(hash, null, 2));
		expect(hash.format()).to.eql(sample);
	});
});
