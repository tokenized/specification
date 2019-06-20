import * as mocha from 'mocha';
import * as chai from 'chai';
import {
{{- range .}}
	{{.Name}},
{{- end}}
} from '../src/actions';
import R from 'ramda';
const expect = chai.expect;
[mocha]

{{range .}}
describe('{{.Name}}', () => {
	it('should load codes', () => {
		// Create a randomized object
		let initialMessage = new {{.Name}}();

	});
});

{{end}}
