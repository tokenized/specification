import {sprintf} from 'sprintf-js';
import _ from '@keyring/util';
import {write, read, ReadVarChar, ReadFixedChar,
	WriteVarChar, WriteFixedChar} from './bytes';
import { Resources } from './resources';
import { Timestamp } from '../src/protocol_types';
import { AgeRestriction } from '../src/field_types';


// AssetTypeLen is the size in bytes of all asset type variants.
const AssetTypeLen = 152;
{{range . }}
// Code{{.Name}} identifies data as a {{.Name}} message.
const Code{{.Name}} = '{{.Code}}';
{{end}}

// AssetPayload is the interface for payloads within asset actions.
export interface AssetPayload {
	Type(): string;
	Serialize(): Buffer;
	Write(Buffer): number;
	Validate(): string;
}

// AssetTypeMapping holds a mapping of asset codes to asset types.
export function AssetTypeMapping(code: string): AssetPayload {
	switch (code) {
{{- range . }}
	case Code{{.Name}}:
		return new {{.Name}}();
{{- end }}
	default:
		return null;
	}
}

{{range .}}
// {{.Name}} asset type.
export class {{.Name}} implements AssetPayload {

{{range .Fields }}// {{ .FieldDescription }}
	{{.SnakeCase}}; // {{.FieldGoType}} {{ .GoTags }}
{{ end }}

	// Type returns the type identifer for this message.
	Type(): string {
		return Code{{.Name}};
	}

	// Len returns the byte size of this message.
	Len(): number {
		return AssetTypeLen;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	Read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// Serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
	{{range $f, $field := .Fields}}

		// {{.FieldName}} ({{.FieldGoType}})
	{{- if ne (len $field.IncludeIfTrue) 0 }}
		if (this.{{ snakecase $field.IncludeIfTrue }}) {
	{{- else if ne (len $field.IncludeIfFalse) 0 }}
		if (!this.{{ snakecase $field.IncludeIfFalse }}) {
	{{- else if ne (len $field.IncludeIf.Field) 0 }}
		if ({{ range $j, $include := $field.IncludeIf.Values }}{{ if (ne $j 0) }} ||{{ end }} this.{{snakecase $field.IncludeIf.Field}} === '{{ $include }}'{{ end }}) {
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
	{{- else if .IsBytes }}
			write(buf, pad(this.{{.SnakeCase}}, {{.Length}}), '{{.FieldGoType}}');
	{{- else}}
			write(buf, this.{{.SnakeCase}}, '{{.FieldGoType}}');
	{{- end }}
		}
	{{ end }}
		return buf.buf;
	}

	// Write populates the fields in {{.Name}} from the byte slice
	Write(b: Buffer): number {
		const buf = new _.Reader(b);
	{{- range $f, $field := .Fields}}

		// {{.Name}} ({{.FieldGoType}})
	{{- if ne (len $field.IncludeIfTrue) 0 }}
		if (this.{{snakecase $field.IncludeIfTrue }}) {
	{{- else if ne (len $field.IncludeIfFalse) 0 }}
		if (!this.{{snakecase $field.IncludeIfFalse }}) {
	{{- else if ne (len $field.IncludeIf.Field) 0 }}
		if ({{ range $j, $include := $field.IncludeIf.Values }}{{ if (ne $j 0) }} ||{{ end }} this.{{snakecase $field.IncludeIf.Field}} === '{{ $include }}'{{ end }}) {
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
	{{ else if .IsInternalType }}
			this.{{ $field.SnakeCase }} = new {{.FieldGoType}}();
			this.{{.SnakeCase}}.Write(buf);
	{{ else if or .IsBytes .IsData }}
			this.{{.SnakeCase}} = buf.read({{.Length}}, '{{.FieldGoType}}');
	{{ else if .Trimmable }}
			this.{{.SnakeCase}} = bytes.Trim(this.{{.SnakeCase}}, "\x00"); // TODO {{.FieldName}}
	{{ else }}
			this.{{.SnakeCase}} = read(buf, '{{.FieldGoType}}');
	{{ end }}
		}
	{{ end }}
		return b.length - buf.length;
	}

	Validate(): string {
	{{- range $i, $field := .Fields }}

		// {{.Name}} ({{.FieldGoType}})
	{{- if ne (len $field.IncludeIfTrue) 0 }}
		if (this.{{ $field.IncludeIfTrue }}) {
	{{- else if ne (len $field.IncludeIfFalse) 0 }}
		if (!this.{{ $field.IncludeIfFalse }}) {
	{{- else if ne (len $field.IncludeIf.Field) 0 }}
		if ({{ range $j, $include := $field.IncludeIf.Values }}{{ if (ne $j 0) }} ||{{ end }} this.{{$field.IncludeIf.Field}} === '{{ $include }}'{{ end }}) {
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

			for _, value := range this.{{.SnakeCase}} {
		{{- if eq .Type "RejectionCode[]" }}
				if (!Resources.GetRejectionCode(value)) {
					return sprintf('Invalid rejection code value : %d', this.{{.SnakeCase}});
				}
		{{- else if eq .Type "Role[]" }}
				if (!Resources.GetRoleType(value)) {
					return sprintf('Invalid role value : %d', this.{{.SnakeCase}});
				}
		{{- else if eq .Type "CurrencyType[]" }}
				if (!Resources.GetCurrency(value)) {
					return sprintf('Invalid currency value : %d', this.{{.SnakeCase}});
				}
		{{- else if eq .Type "Polity[]" }}
				if (!Resources.GetPolityType(string(value[:]))) {
					return sprintf('Invalid polity value : %d', this.{{.SnakeCase}});
				}
		{{- else if eq .Type "EntityType[]" }}
				if (!Resources.GetEntityType(value)) {
					return sprintf('Invalid entity type value : %c', this.{{.SnakeCase}});
				}
		{{- else if eq .Type "Tag[]" }}
				if (!Resources.GetTagType(value)) {
					return sprintf('Invalid tag type value : %c', this.{{.SnakeCase}});
				}
		{{- end }}
			}
	{{- else if eq .Type "RejectionCode" }}
			if (!Resources.GetRejectionCode(this.{{.SnakeCase}})) {
				return sprintf('Invalid rejection code value : %d', this.{{.SnakeCase}});
			}
	{{- else if eq .Type "Role" }}
			if (!Resources.GetRoleType(this.{{.SnakeCase}})) {
				return sprintf('Invalid role value : %d', this.{{.SnakeCase}});
			}
	{{- else if eq .Type "CurrencyType" }}
			if (!Resources.GetCurrency(this.{{.SnakeCase}})) {
				return sprintf('Invalid currency value : %d', this.{{.SnakeCase}});
			}
	{{- else if eq .Type "Polity" }}
			if (!Resources.GetPolityType(string(this.{{.SnakeCase}}[:]))) {
				return sprintf('Invalid polity value : %d', this.{{.SnakeCase}});
			}
	{{- else if eq .Type "EntityType" }}
			if (!Resources.GetEntityType(this.{{.SnakeCase}})) {
				return sprintf('Invalid entity type value : %c', this.{{.SnakeCase}});
			}
	{{- else if eq .Type "Tag" }}
			if (!Resources.GetTagType(this.{{.SnakeCase}})) {
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
			if (err) {
				return sprintf('field {{.SnakeCase}} is invalid : %s', err);
			}
	{{ else if ne (len $field.IntValues) 0 }}
			if ({{ range $j, $value := $field.IntValues }}{{ if (ne $j 0) }} &&{{ end }} this.{{$field.SnakeCase}} !== {{ $value }}{{ end }}) {
				return sprintf('field {{$field.SnakeCase}} value is invalid : %d', this.{{$field.SnakeCase}});
			}
	{{ else if ne (len $field.CharValues) 0 }}
			if ({{ range $j, $value := $field.CharValues }}{{ if (ne $j 0) }} &&{{ end }} this.{{$field.SnakeCase}} !== '{{ $value }}'{{ end }}) {
				return sprintf('field {{$field.SnakeCase}} value is invalid : %d', this.{{$field.SnakeCase}});
			}
	{{- end }}
		}
	{{ end }}
	}

	toString(): string {
		const vals = [];
	{{ range .Fields -}}
		{{- if eq .Type "STRING" }}
		vals.push(sprintf('{{.SnakeCase}}:\'%v\'', string(this.{{.SnakeCase}})));
		{{- else if .IsNumeric }}
		vals.push(sprintf('{{.SnakeCase}}:%v', this.{{.SnakeCase}}));
		{{- else if eq .Type "SHA" }}
		vals.push(sprintf('{{.SnakeCase}}:\'%x\'', this.{{.SnakeCase}}));
		{{- else if eq .FieldGoType "[]byte" }}
		vals.push(sprintf('{{.SnakeCase}}:%#x', this.{{.SnakeCase}}));
		{{- else }}
		vals.push(sprintf('{{.SnakeCase}}:%#+v', this.{{.SnakeCase}}));
		{{- end }}{{ end }}

		return sprintf('{%s}', vals.join(' '));
	}
}

{{ end -}}
