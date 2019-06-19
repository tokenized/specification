import {sprintf} from 'sprintf-js';
import _ from '@keyring/util';
import {write, read, ReadVarChar, ReadVariableSize, ReadVarBin, ReadFixedChar,
	WriteVarChar, WriteVariableSize, WriteFixedChar, WriteVarBin} from './bytes';
import { GetRoleType, GetEntityType, } from './resources';

{{range .}}
{{comment (print .Name " " .Metadata.Description) "//"}}
export class {{.Name}} {
{{range .Fields}}
	{{- comment (print .FieldDescription) "\t//"}}
	{{ .SnakeCase }}; // {{ .FieldGoType }} {{ .GoTags }}
{{ end -}}

	// Serialize returns the byte representation of the message.
	Serialize(): Buffer {
		const buf = new _.Writer();
	{{- range $i, $field := .Fields }}

		// {{.Name}} ({{.FieldGoType}})
	{{- if ne (len $field.IncludeIfTrue) 0 }}
		if (this.{{ snakecase $field.IncludeIfTrue }}) {
	{{- else if ne (len $field.IncludeIfFalse) 0 }}
		if (!this.{{ snakecase $field.IncludeIfFalse }}) {
	{{- else if ne (len $field.IncludeIf.Field) 0 }}
		if ({{ range $j, $include := $field.IncludeIf.Values }}{{ if (ne $j 0) }} ||{{ end }} this.{{snakecase $field.IncludeIf.Field}} === '{{ $include }}'{{ end }}) {
	{{- else if ne (len $field.IncludeIfInt.Field) 0 }}
		if ({{ range $j, $include := $field.IncludeIfInt.Values }}{{ if (ne $j 0) }} ||{{ end }} this.{{snakecase $field.IncludeIfInt.Field}} === {{ $include }}{{ end }}) {
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
			this.{{.SnakeCase}}.forEach(value => {
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
	{{- else if or .IsBytes .IsData }}
			WriteFixedBin(buf, this.{{.SnakeCase}}, {{.Length}});
	{{- else }}
			write(buf, this.{{.SnakeCase}}, '{{.FieldGoType}}');
	{{- end }}
		}
	{{- end }}
		return buf.buf;
	}

	Write(buf: _.Reader) {
	{{- range $i, $field := .Fields }}

		// {{.Name}} ({{.FieldGoType}})
	{{- if ne (len $field.IncludeIfTrue) 0 }}
		if (this.{{ snakecase $field.IncludeIfTrue }}) {
	{{- else if ne (len $field.IncludeIfFalse) 0 }}
		if (!action.{{ snakecase $field.IncludeIfFalse }}) {
	{{- else if ne (len $field.IncludeIf.Field) 0 }}
		if ({{ range $j, $include := $field.IncludeIf.Values }}{{ if (ne $j 0) }} ||{{ end }} this.{{snakecase $field.IncludeIf.Field}} === '{{ $include }}'{{ end }}) {
	{{- else if ne (len $field.IncludeIfInt.Field) 0 }}
		if ({{ range $j, $include := $field.IncludeIfInt.Values }}{{ if (ne $j 0) }} ||{{ end }} this.{{snakecase $field.IncludeIfInt.Field}} === {{ $include }}{{ end }}) {
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
			const size = ReadVariableSize(buf, {{.Length}}, 8);
			this.{{.SnakeCase}} = [Array(size)]
				.map(() => {
					const newValue = new {{.SingularType}}();
					newValue.Write(buf);
					return newValue;
				});
	{{- else if .IsNativeTypeArray }}
			const size = ReadVariableSize(buf, {{.Length}}, 8);
			const type = '{{.FieldGoType}}'.slice(2);
			this.{{.SnakeCase}} = [...Array(size)].map(() => read(buf, type));
	{{- else if .IsInternalType }}
			this.{{.SnakeCase}}.Write(buf);
	{{- else if or .IsBytes .IsData }}
			this.{{.SnakeCase}} = buf.read({{.Length}});
	{{- else }}
			this.{{.SnakeCase}} = read(buf, '{{.FieldGoType}}');
	{{- end }}
		}
	{{- end }}
	}

	Validate() {
	{{- range $i, $field := .Fields }}

		// {{.SnakeCase}} ({{.FieldGoType}})
	{{- if ne (len $field.IncludeIfTrue) 0 }}
		if (this.{{ snakecase $field.IncludeIfTrue }}) {
	{{- else if ne (len $field.IncludeIfFalse) 0 }}
		if (!action.{{ snakecase $field.IncludeIfFalse }}) {
	{{- else if ne (len $field.IncludeIf.Field) 0 }}
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

			this.{{.SnakeCase}}.forEach(value => {
		{{- if eq .Type "RejectionCode[]" }}
				if (!GetRejectionCode(value)) {
					return sprintf('Invalid rejection code value : %d', this.{{.SnakeCase}});
				}
		{{- else if eq .Type "Role[]" }}
				if (!GetRoleType(value)) {
					return sprintf('Invalid role value : %d', this.{{.SnakeCase}});
				}
		{{- else if eq .Type "CurrencyType[]" }}
				if (!GetCurrency(value)) {
					return sprintf('Invalid currency value : %d', this.{{.SnakeCase}});
				}
		{{- else if eq .Type "Polity[]" }}
				if (!GetPolityType(Buffer.from(value).toString('ascii'))) {
					return sprintf('Invalid polity value : %d', this.{{.SnakeCase}});
				}
		{{- else if eq .Type "EntityType[]" }}
				if (!GetEntityType(value)) {
					return sprintf('Invalid entity type value : %c', this.{{.SnakeCase}});
				}
		{{- else if eq .Type "Tag[]" }}
				if (!GetTagType(value)) {
					return sprintf('Invalid tag type value : %c', this.{{.SnakeCase}});
				}
		{{- end }}
			});
	{{- else if eq .Type "RejectionCode" }}
			if (!GetRejectionCode(this.{{.SnakeCase}})) {
				return sprintf('Invalid rejection code value : %d', this.{{.SnakeCase}});
			}
	{{- else if eq .Type "Role" }}
			if (!GetRoleType(this.{{.SnakeCase}})) {
				return sprintf('Invalid role value : %d', this.{{.SnakeCase}});
			}
	{{- else if eq .Type "CurrencyType" }}
			if (!GetCurrency(this.{{.SnakeCase}})) {
				return sprintf('Invalid currency value : %d', this.{{.SnakeCase}});
			}
	{{- else if eq .Type "Polity" }}
			if (!GetPolityType(string(this.{{.SnakeCase}}[:]))) {
				return sprintf('Invalid polity value : %d', this.{{.SnakeCase}});
			}
	{{- else if eq .Type "EntityType" }}
			if (!GetEntityType(this.{{.SnakeCase}})) {
				return sprintf('Invalid entity type value : %c', this.{{.SnakeCase}});
			}
	{{- else if eq .Type "Tag" }}
			if (!GetTagType(this.{{.SnakeCase}})) {
				return sprintf('Invalid tag type value : %c', this.{{.SnakeCase}});
			}
	{{- else if .IsInternalTypeArray }}
			if (this.{{.SnakeCase}}.length > (2 << {{.Length}}) - 1) {
				return sprintf('list field {{.SnakeCase}} has too many items %d/%d', this.{{.SnakeCase}}.length, (2 << {{.Length}}) - 1);
			}

			this.{{.SnakeCase}}.forEach((value, i) => {
				const err = value.Validate();
				if (err) {
					return sprintf('list field {{.SnakeCase}}[%d] is invalid : %s', i, err);
				}
			});
	{{- else if .IsNativeTypeArray }}
			if (this.{{.SnakeCase}}.length > (2 << {{.Length}}) - 1) {
				return sprintf('list field {{.SnakeCase}} has too many items %d/%d', this.{{.SnakeCase}}.length, (2 << {{.Length}}) - 1);
			}
	{{- else if .IsInternalType }}
			const err = this.{{.SnakeCase}}.Validate();
			if (err) {
				return sprintf('field {{.SnakeCase}} is invalid : %s', err);
			}
	{{- end }}
		}
	{{- end }}
		return;
	}

	Equal(other: {{.Name}}): boolean {
	{{- range .Fields }}

		// {{.Name}} ({{.FieldGoType}})
	{{- if .IsVarBin }}
		if (!Buffer.compare(this.{{.SnakeCase}}, other.{{.SnakeCase}})) {
			return false;
		}
	{{- else if or .IsBytes .IsData }}
		if (!Buffer.compare(this.{{.SnakeCase}}, other.{{.SnakeCase}})) {
			return false;
		}
	{{- else if .IsInternalTypeArray }}
		if (this.{{.SnakeCase}}.length !== 0 || other.{{.SnakeCase}}.length !== 0 ) {
			if (this.{{.SnakeCase}}.length !== other.{{.SnakeCase}}.length) {
				return false;
			}

			if (!this.{{.SnakeCase}}.every((value, i) => {
				if (!value.Equal(other.{{.SnakeCase}}[i])) {
					return false;
				}
				return true;
			})) return false;
		}
	{{- else if .IsInternalType }}
		if (!this.{{.SnakeCase}}.Equal(other.{{.SnakeCase}})) {
			return false;
		}
	{{- else }}
		if (this.{{.SnakeCase}} !== other.{{.SnakeCase}}) {
			return false;
		}
	{{- end }}
	{{- end }}
		return true;
	}

}

{{ end }}
