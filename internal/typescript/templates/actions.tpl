// Package protocol provides base level structs and validation for
// the protocol.
//
// The code in this file is auto-generated. Do not edit it by hand as it will
// be overwritten when code is regenerated.

import {sprintf} from 'sprintf-js';
import _ from '@keyring/util';
import {write, read, ReadVarChar, ReadVariableSize, ReadVarBin, ReadFixedChar,
	WriteVarChar, WriteVariableSize, WriteFixedChar, WriteVarBin} from './bytes';
import { Document, Amendment, VotingSystem, Oracle, Entity, TargetAddress,
	QuantityIndex, AssetTransfer, AssetSettlement } from './field_types';
import { GetPolityType, GetRejectionCode, } from './resources';

enum ActionCode {
{{- range .}}
	{{.CodeNameComment}}
	{{.CodeName}} = '{{.Code}}',
{{ end }}

	// ComplianceActionFreeze identifies a freeze type
	ComplianceActionFreeze = 'F',

	// ComplianceActionThaw identifies a thaw type
	ComplianceActionThaw = 'T',

	// ComplianceActionConfiscation identifies a confiscation type
	ComplianceActionConfiscation = 'C',

	// ComplianceActionReconciliation identifies a reconcilation type
	ComplianceActionReconciliation = 'R',
}

class OpReturnMessage {}


// TypeMapping holds a mapping of action codes to action types.
export function TypeMapping(code: string): OpReturnMessage {
	switch (code) {
{{- range . }}
	case ActionCode.Code{{.Name}}:
		return new {{.Name}}();
{{- end }}
	default:
		return null;
	}
}

{{ range $action := . }}

{{comment (print .Name " " .Metadata.Description) "//"}}
class {{.Name}}  {
{{ range .Fields }}
	{{comment (print "\t" .FieldDescription) "\t//"}}
	{{ .SnakeCase }};{{ end -}}

{{ if .Constructor -}}
// New{{.Name}} returns a new {{.Name}} with specified values set.
constructor ({{ range $i, $c := .Constructor }}{{if $i}}, {{end}}{{ .ConstructorName }} {{ .ConstructorGoType }}{{ end -}}) *{{.Name}} {
	result := {{.Name}}{}
	{{ range .Constructor -}}
	result.{{ .ConstructorField -}}
	{{ if eq .ConstructorSetMethod "=" }} = {{ .ConstructorName }}{{ else }}.{{ .ConstructorSetMethod }}({{ .ConstructorName }}){{ end }}
	{{ end -}}
	return &result
}
{{ end -}}

{{ range .Functions }}// {{.FunctionName}} {{ .FunctionDescription }}.
func (action *{{ $action.Name }}) {{.FunctionName}}({{ range $i, $c := .FunctionParams }}{{if ne .ParamName "hidden"}}{{if $i}}, {{end}}{{ .ParamName }} {{ .ParamGoType }}{{ end -}}{{ end -}}) {
{{ if eq .FunctionType "set" -}}
	{{ range .FunctionParams -}}
	action.{{ .ParamField -}}
	{{ if eq .ParamSetMethod "=" }} = {{ .ParamValue }}{{ else }}.{{ .ParamSetMethod }}({{ .ParamValue }}){{ end }}
	{{ end -}}
{{ else if eq .FunctionType "append" -}}
	action.{{ .FunctionAppendType }}s = append(action.{{ .FunctionAppendType }}s, *New{{ .FunctionAppendType }}({{ range $i, $param := .FunctionParams }}{{if $i}}, {{end}}{{ .ParamName }}{{ end }}))
{{ end -}}
}
{{ end }}

	// Type returns the type identifer for this message.
	Type(): string {
		return ActionCode.Code{{.Name}};
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.serialize();

		b.fill(data);

		return data.length;
	}

	// serialize returns the full OP_RETURN payload bytes.
	serialize(): Buffer {
		const buf = new _.Writer();


{{- range $f, $field := .PayloadFields }}
{{- $fieldName := .FieldName }}
		// {{.SnakeCase}} ({{.FieldGoType}})
{{- if ne (len $field.IncludeIfTrue) 0 }}
		if (this.{{ snakecase $field.IncludeIfTrue }}) {
{{- else if ne (len $field.IncludeIfFalse) 0 }}
		if (!this.{{ snakecase $field.IncludeIfFalse }}) {
{{- else if ne (len $field.IncludeIf.Field) 0 }}
		if ({{ range $j, $include := $field.IncludeIf.Values }}{{ if (ne $j 0) }} ||{{ end }} this.{{ snakecase $field.IncludeIf.Field}} === '{{ $include }}'{{ end }}) {
{{- else if ne (len $field.IncludeIfInt.Field) 0 }}
		if ({{ range $j, $include := $field.IncludeIfInt.Values }}{{ if (ne $j 0) }} ||{{ end }} this.{{ snakecase $field.IncludeIfInt.Field}} === {{ $include }}{{ end }}) {
{{- else }}
		{
{{- end }}
{{- if .IsVarChar }}
			WriteVarChar(buf, this.{{.SnakeCase}}, {{.Length}});
{{- else if .IsFixedChar }}
			WriteFixedChar(buf, this.{{.SnakeCase}}, {{.Length}});
{{- else if .IsVarBin }}
			WriteVarBin(buf, this.{{.SnakeCase}}, {{.Length}});
{{- else if .IsInternalTypeArray }}
			WriteVariableSize(buf, this.{{.SnakeCase}}.length, {{.Length}}, 8);
			this.{{.SnakeCase}}.forEach((value) => {
				buf.write(value.Serialize());
			});
{{- else if .IsNativeTypeArray }}
			// IsNativeTypeArray {{.FieldGoType}}
			const type = '{{.FieldGoType}}'.slice(2);
			WriteVariableSize(buf, this.{{.SnakeCase}}.length, {{.Length}}, 8);
			this.{{.SnakeCase}}.forEach(value => {
				write(buf, value, type); // TODO
			});
{{- else if .IsInternalType }}
			buf.write(this.{{.SnakeCase}}.Serialize()); 
{{- else if .IsBytes }}
			write(buf, pad(this.{{.SnakeCase}}, {{.Length}}));
{{- else}}
			write(buf, this.{{.SnakeCase}}, '{{.FieldGoType}}');
{{- end }}
		}
{{ end }}
		return buf.buf;
	}

	// write populates the fields in {{.Name}} from the byte slice
	write(b: _.Writer): number {
		const buf = new Buffer(b);

{{- range $f, $field := .PayloadFields }}
		// {{.SnakeCase}} ({{.FieldGoType}})
{{- if ne (len $field.IncludeIfTrue) 0 }}
		if (this.{{ snakecase $field.IncludeIfTrue }}) {
{{- else if ne (len $field.IncludeIfFalse) 0 }}
		if (!this.{{ snakecase $field.IncludeIfFalse }}) {
{{- else if ne (len $field.IncludeIf.Field) 0 }}
		if ({{ range $j, $include := $field.IncludeIf.Values }}{{ if (ne $j 0) }} ||{{ end }} this.{{ snakecase $field.IncludeIf.Field}} === '{{ $include }}'{{ end }}) {
{{- else if ne (len $field.IncludeIfInt.Field) 0 }}
		if ({{ range $j, $include := $field.IncludeIfInt.Values }}{{ if (ne $j 0) }} ||{{ end }} this.{{ snakecase $field.IncludeIfInt.Field}} === {{ $include }}{{ end }}) {
{{- else }}
		{
{{- end }}
{{- if .IsVarChar }}
			this.{{.SnakeCase}} = ReadVarChar(buf, {{.Length}});
{{- else if .IsFixedChar }}
			this.{{.SnakeCase}} = ReadFixedChar(buf, {{.Length}});
{{- else if .IsVarBin }}
			this.{{.SnakeCase}} = ReadVarBin(buf, {{.Length}});
{{- else if .IsInternalTypeArray }}
			// IsInternalTypeArray
			const size = ReadVariableSize(buf, {{.Length}}, 8);
			this.{{.SnakeCase}} = [];
			for (let i = 0; i < size; i++) {
				const newValue = new {{.SingularType}}();
				newValue.Write(buf);

				this.{{.SnakeCase}}.push(newValue);
			}
{{- else if .IsNativeTypeArray }}
			const size = ReadVariableSize(buf, {{.Length}}, 8);
			const type = '{{.FieldGoType}}'.slice(2);
			this.{{.SnakeCase}} = [...Array(size)].map(() => read(buf, type));
{{- else if .IsInternalType }}
			this.{{.SnakeCase}}.Write(buf);
{{- else if or .IsBytes .IsData }}
			this.{{.SnakeCase}} = buf.read({{.Length}});
{{- else if .Trimmable }}
			this.{{.SnakeCase}} = bytes.Trim(this.{{.SnakeCase}}, "\x00");
{{- else }}
			read(buf, this.{{.SnakeCase}});
{{- end }}
		}
{{ end }}
		return b.length - buf.length;
	}

	Validate(): string {
{{- range $i, $field := .PayloadFields }}

		// {{.Name}} ({{.FieldGoType}})
{{- if ne (len $field.IncludeIfTrue) 0 }}
		// IncludeIfTrue
		if (this.{{ snakecase $field.IncludeIfTrue }}) {
{{- else if ne (len $field.IncludeIfFalse) 0 }}
		// IncludeIfFalse
		if (!action.{{ snakecase $field.IncludeIfFalse }}) {
{{- else if ne (len $field.IncludeIf.Field) 0 }}
		// IncludeIf.Field
		if ({{ range $j, $include := $field.IncludeIf.Values }}{{ if (ne $j 0) }} ||{{ end }} this.{{snakecase $field.IncludeIf.Field}} === '{{ $include }}'{{ end }}) {
{{- else }}
		{
{{- end }}
{{- if .IsVarChar }}
			if (this.{{.SnakeCase}}.length > (2 << {{.Length}}) - 1) {
				return sprintf('varchar field {{.SnakeCase}} too long %d/%d', this.{{.SnakeCase}}.length, (2 << {{.Length}}) - 1);
			}
{{- else if .IsFixedChar }}
			if (this.{{.SnakeCase}}.length > {{.Length}}) {
				return sprintf('fixedchar field {{.SnakeCase}} too long %d/%d', this.{{.SnakeCase}}.length, {{.Length}});
			}
{{- else if .IsVarBin }}
			if (this.{{.SnakeCase}}.length > (2 << {{.Length}}) - 1) {
				return sprintf('varbin field {{.SnakeCase}} too long %d/%d', this.{{.SnakeCase}}.length, (2 << {{.Length}}) - 1);
			}
{{- else if .IsResourceTypeArray }}
			if (this.{{.SnakeCase}}.length > (2 << {{.Length}}) - 1) {
				return sprintf('list field {{.SnakeCase}} has too many items %d/%d', this.{{.SnakeCase}}.length, (2 << {{.Length}}) - 1);
			}

			const err = this.{{.SnakeCase}}.find((value) => {
{{- if eq .Type "RejectionCode[]" }}
				if (GetRejectionCode(value) === null) {
					return sprintf('Invalid rejection code value : %d', this.{{.SnakeCase}});
				};
{{- else if eq .Type "Role[]" }}
				if (GetRoleType(value) === null) {
					return sprintf('Invalid role value : %d', this.{{.SnakeCase}});
				}
{{- else if eq .Type "CurrencyType[]" }}
				if (GetCurrency(value) === null) {
					return sprintf('Invalid currency value : %d', this.{{.SnakeCase}});
				}
{{- else if eq .Type "Polity[]" }}
				if (GetPolityType(Buffer.from(value).toString('ascii')) === null) {
					return sprintf('Invalid polity value : %d', this.{{.SnakeCase}});
				}
{{- else if eq .Type "EntityType[]" }}
				if (GetEntityType(value) === null) {
					return sprintf('Invalid entity type value : %c', this.{{.SnakeCase}});
				}
{{- else if eq .Type "Tag[]" }}
				if (GetTagType(value) === null) {
					return sprintf('Invalid tag type value : %c', this.{{.SnakeCase}});
				}
{{- end }}
			});
			if (err) return err;
{{- else if eq .Type "RejectionCode" }}
			if (GetRejectionCode(this.{{.SnakeCase}}) === null) {
				return sprintf('Invalid rejection code value : %d', this.{{.SnakeCase}});
			}
{{- else if eq .Type "Role" }}
			if (GetRoleType(this.{{.SnakeCase}}) === null) {
				return sprintf('Invalid role value : %d', this.{{.SnakeCase}});
			}
{{- else if eq .Type "CurrencyType" }}
			if (GetCurrency(this.{{.SnakeCase}}) === null) {
				return sprintf('Invalid currency value : %d', this.{{.SnakeCase}});
			}
{{- else if eq .Type "Polity" }}
			if (GetPolityType(Buffer.from(this.{{.SnakeCase}}).toString('ascii') === null) {
				return sprintf('Invalid polity value : %d', this.{{.SnakeCase}});
			}
{{- else if eq .Type "EntityType" }}
			if (GetEntityType(this.{{.SnakeCase}}) === null) {
				return sprintf('Invalid entity type value : %c', this.{{.SnakeCase}});
			}
{{- else if eq .Type "Tag" }}
			if (GetTagType(this.{{.SnakeCase}}) === null) {
				return sprintf('Invalid tag type value : %c', this.{{.SnakeCase}});
			}
{{- else if .IsInternalTypeArray }}
			if (this.{{.SnakeCase}}.length > (2 << {{.Length}}) - 1) {
				return sprintf('list field {{.SnakeCase}} has too many items %d/%d', this.{{.SnakeCase}}.length, (2 << {{.Length}}) - 1);
			}

			const err = this.{{.SnakeCase}}.find((value, i) => {
				const err2 = value.Validate();
				if (err2) return sprintf('list field {{.SnakeCase}}[%d] is invalid : %s', i, err2);
			});
			if (err) return err;
{{- else if .IsNativeTypeArray }}
			if (this.{{.SnakeCase}}.length > (2 << {{.Length}}) - 1) {
				return sprintf('list field {{.SnakeCase}} has too many items %d/%d', this.{{.SnakeCase}}.length, (2 << {{.Length}}) - 1);
			}
{{- else if .IsInternalType }}
			const err = this.{{.SnakeCase}}.Validate();
			if  (err) return sprintf('field {{.SnakeCase}} is invalid : %s', err);
{{ else if ne (len $field.IntValues) 0 }}
			if ({{ range $j, $value := $field.IntValues }}{{ if (ne $j 0) }} &&{{ end }} this.{{$field.SnakeCase}} !== {{ $value }}{{ end }}) {
				return sprintf('field {{$field.SnakeCase}} value is invalid : %d', this.{{$field.SnakeCase}});
			}
{{ else if ne (len $field.CharValues) 0 }}
			if ({{ range $j, $value := $field.CharValues }}{{ if (ne $j 0) }}&&{{ end }} this.{{$field.SnakeCase}} !== '{{ $value }}'
			{{ end }}) {
				return sprintf('field {{$field.SnakeCase}} value is invalid : %d', this.{{$field.SnakeCase}});
			}
{{- end }}
		}
	{{- end }}
		return null;
	}

	toString(): string {
		const vals: string[] = [];
	{{- range .Fields -}}
		{{- if eq .Type "STRING" }}
		vals.push(sprintf('{{.SnakeCase}}:\"%v\"', string(this.{{.SnakeCase}})));
		{{- else if .IsNumeric }}
		vals.push(sprintf('{{.SnakeCase}}:%v', this.{{.SnakeCase}}));
		{{- else if eq .Type "SHA" }}
		vals.push(sprintf('{{.SnakeCase}}:\"%x\"', this.{{.SnakeCase}}));
		{{- else if eq .FieldGoType "[]byte" }}
		vals.push(sprintf('{{.SnakeCase}}:%#x', this.{{.SnakeCase}}));
		{{- else }}
		vals.push(sprintf('{{.SnakeCase}}:%#+v', this.{{.SnakeCase}}));
		{{- end }}{{ end }}

		return sprintf('{%s}', vals.join(' '));
	}
}
{{- end }}

