import {sprintf} from 'sprintf-js';
import BN from 'bn.js';
import * as mocha from 'mocha';
import * as chai from 'chai';
import {char} from '../src/bytes';
import { TxId, AssetCode, Timestamp, ContractCode, PublicKeyHash } from '../src/protocol_types';
import { Document, Amendment, VotingSystem, Oracle, Entity, TargetAddress,
	QuantityIndex, AssetTransfer, AssetSettlement } from '../src/field_types';
import {
{{- range .}}
	{{.Name}},
{{- end}}
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

{{range .}}
describe('{{.Name}}', () => {
	it('should encode and decode', () => {
		// Create a randomized object
		let initialMessage = new {{.Name}}();


		{{- range $i, $field := .PayloadFields }}
		// {{ $field.SnakeCase }} ({{ $field.Type }})
		{{- if ne (len .IncludeIfTrue) 0 }}
		if (initialMessage.{{snakecase .IncludeIfTrue }}) {
		{{- else if ne (len $field.IncludeIf.Field) 0 }}
		if ({{ range $j, $include := .IncludeIf.Values }}{{ if (ne $j 0) }} ||{{ end }} initialMessage.{{snakecase $field.IncludeIf.Field}} == char('{{ $include }}'){{ end }}) {
		{{- else }}
		{
		{{- end }}
			{{- if eq $field.Type "bool" }}
			initialMessage.{{ $field.SnakeCase }} = true
			{{- else if $field.IsVarChar }}
			initialMessage.{{ $field.SnakeCase }} = "Text {{ $i }}"
			{{- else if $field.IsFixedChar }}
			initialMessage.{{ $field.SnakeCase }} = [...Array({{ $field.Length }}).keys()]
				.map(i => String.fromCharCode(i + 65 + {{ $i }}))
				.join('');
			{{- else if $field.IsBytes }}
			initialMessage.{{ $field.SnakeCase }} = Buffer.from([...Array({{ $field.Length }}).keys()]
				.map(i => i + 65 + {{ $i }}));
			{{- else if eq $field.Type "timestamp" }}
			initialMessage.{{ $field.SnakeCase }} = new BN(Date.now()).mul(1000000);
			{{- else if .IsInternalTypeArray }}
			// IsInternalTypeArray
			initialMessage.{{ $field.SnakeCase }} = [...Array(2)].map(() => new {{.SingularType}}());
			{{- else if .IsNativeTypeArray }}
			// IsNativeTypeArray
				{{- if eq .Type "Polity[]" }}
			initialMessage.{{ $field.SnakeCase }} = [...Array(5)].map(() => [...Buffer.from('AUS', 'ascii')]);
				{{- else}}
			initialMessage.{{ $field.SnakeCase }} = [...Array(5)].map(() => getArrayOrType('{{.SingularType}}'));
				{{- end }}
			{{- else if .IsInternalType }}
			initialMessage.{{ $field.SnakeCase }} = new {{.FieldGoType}}();
			{{ else if ne (len $field.IntValues) 0 }}
				initialMessage.{{ $field.SnakeCase }} = 1;
			{{- else if .IsNativeType }}
			// IsNativeType
				{
			let type = '{{.FieldGoType}}';
			if(type.startsWith('['))
				initialMessage.{{ $field.SnakeCase }} = getArrayOrType(type);
			else if(type === 'uint64')
				initialMessage.{{ $field.SnakeCase }} = new BN(65 + {{ $i }});
			else
				initialMessage.{{ $field.SnakeCase }} = 65 + {{ $i }};
				}
			{{- else }}
			// {{ $field.Type }} test not setup
			{{- end }}
		}
		{{ end }}

//		let err = initialMessage.Validate();
//		if(err) throw new Error('Error validating initial message: ' + err);

		// Encode message
		let initialEncoding = initialMessage.Serialize();

		// Decode message
		let decodedMessage = new {{.Name}}();

		let n = decodedMessage.write(initialEncoding);
//		err = decodedMessage.Validate();
//		if(err) throw new Error('Error validating decoded message: ' + err);

		// expect(n).to.eql(initialEncoding.length);

		// Serializing the message should give us the same bytes
		let secondEncoding = decodedMessage.Serialize();

		expect(initialEncoding).to.eql(secondEncoding);

		// if !cmp.Equal(&initialMessage, &decodedMessage) {
		// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
		// }

		// Compare re-serialized values
		{{- range $i, $field := .PayloadFields }}
		// {{ $field.SnakeCase }} ({{ $field.Type }})
			{{- if $field.IsVarChar }}
		if (initialMessage.{{ $field.SnakeCase }} !== (decodedMessage.{{ $field.SnakeCase }})) {
			throw new Error(sprintf("{{ $field.SnakeCase }} doesn't match : %s != %s", initialMessage.{{ $field.SnakeCase }}, decodedMessage.{{ $field.SnakeCase }}));
		}
			{{- else if $field.IsFixedChar }}
			// IsFixedChar
		if (initialMessage.{{ $field.SnakeCase }} != decodedMessage.{{ $field.SnakeCase }}) {
			throw new Error(sprintf("{{ $field.SnakeCase }} doesn't match : %s != %s", initialMessage.{{ $field.SnakeCase }}, decodedMessage.{{ $field.SnakeCase }}));
		}
			{{- else if $field.IsBytes }}
		if ((initialMessage.{{ $field.SnakeCase }} && decodedMessage.{{ $field.SnakeCase }}) 
				&& !initialMessage.{{ $field.SnakeCase }}.equals(decodedMessage.{{ $field.SnakeCase }})) {
			throw new Error(sprintf("{{ $field.SnakeCase }} doesn't match : %x != %x", initialMessage.{{ $field.SnakeCase }}, decodedMessage.{{ $field.SnakeCase }}));
		}
			{{- else if eq $field.Type "Polity" }}
			{
		let initialJson = JSON.stringify(initialMessage.{{ $field.SnakeCase }});
		let decodedJson = JSON.stringify(decodedMessage.{{ $field.SnakeCase }});
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("{{ $field.SnakeCase }} doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
			{{- else if .IsInternalTypeArray }}
		if ((initialMessage.{{ $field.SnakeCase }} && decodedMessage.{{ $field.SnakeCase }}) 
				&& initialMessage.{{ $field.SnakeCase }}.length != decodedMessage.{{ $field.SnakeCase }}.length) {
			throw new Error(sprintf("{{ $field.SnakeCase }} lengths don't match : %d != %d", initialMessage.{{ $field.SnakeCase }}.length, decodedMessage.{{ $field.SnakeCase }}.length));
		}
			{{- else if .IsNativeTypeArray }}
			{
		let initialJson = JSON.stringify(initialMessage.{{ $field.SnakeCase }});
		let decodedJson = JSON.stringify(decodedMessage.{{ $field.SnakeCase }});
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("{{ $field.SnakeCase }} doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
			{{- else if .IsNativeType }}
			// IsNativeType
			{
		let initialJson = JSON.stringify(initialMessage.{{ $field.SnakeCase }});
		let decodedJson = JSON.stringify(decodedMessage.{{ $field.SnakeCase }});
		if (initialJson !== decodedJson) {
			throw new Error(sprintf("{{ $field.SnakeCase }} doesn't match : %s != %s", initialJson, decodedJson));
		}
			}
			{{- else }}
		// {{ $field.Type }} test compare not setup
			{{- end }}
		{{- end }}


	});
});

{{end}}
