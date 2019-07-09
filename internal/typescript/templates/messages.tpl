import {sprintf} from 'sprintf-js';
import _ from '@keyring/util';
import {write, read, ReadVarChar, ReadVariableSize, ReadVarBin,
	WriteVarChar, WriteVariableSize, WriteVarBin} from './bytes';
import { TxId, Timestamp, } from './protocol_types';
import { TargetAddress, OutputTag, Output, Document } from './field_types';
import { Resources } from './resources';

export enum MsgCodes {
{{- range .}}
	{{.CodeNameComment}}
	{{.Name}} = {{.Code}},
{{ end }}
}

// MessagePayload is the interface for payloads within message actions.
export interface MessagePayload {
	Type(): number;
	Serialize(): Buffer;
	Write(Buffer): number;
	Validate(): string;
}

// MessageTypeMapping holds a mapping of message codes to message types.
export function MessageTypeMapping(code: number): MessagePayload {
	switch (code) {
{{- range . }}
	case MsgCodes.{{.Name}}:
		return new {{.Name}}();
{{- end }}
	default:
		return null;
	}
}

{{ range $action := . }}

{{comment (print .Name " " .Metadata.Description) "//"}}
export class {{.Name}} {
{{ range .Fields }}
	{{- comment (print .FieldDescription) "\t//"}}
	{{- if .IsInternalTypeArray }}
	{{ .SnakeCase }} = []; // {{ .FieldGoType }} {{ .GoTags }}
	{{- else if .IsInternalType }}
	{{ .SnakeCase }} = new {{.SingularType}}(); // {{ .FieldGoType }} {{ .GoTags }}
	{{- else }}
	{{ .SnakeCase }}; // {{ .FieldGoType }} {{ .GoTags }}
	{{- end }}
{{- end}}

	{{ if .Constructor -}}
	// New{{.Name}} returns a new {{.Name}} with specified values set.
	func New{{.Name}}({{ range $i, $c := .Constructor }}{{if $i}}, {{end}}{{ .ConstructorName }} {{ .ConstructorGoType }}{{ end -}}) *{{.Name}} {
		result := {{.Name}}{}
		{{ range .Constructor -}}
		result.{{ .ConstructorField -}}
		{{ if eq .ConstructorSetMethod "=" }} = {{ .ConstructorName }}{{ else }}.{{ .ConstructorSetMethod }}({{ .ConstructorName }}){{ end }}
		{{ end -}}
		return &result
	}
	{{ end -}}

	{{ range .Functions }}
	// Function {{.FunctionName}} {{ .FunctionDescription }}.
	{{.FunctionName}}({{ range $i, $c := .FunctionParams }}{{if ne .ParamName "hidden"}}{{if $i}}, {{end}}{{ .ParamName }} {{ .ParamGoType }}{{ end -}}{{ end -}}) {
	{{ if eq .FunctionType "set" -}}
		{{ range .FunctionParams -}}
		this.{{ .ParamField -}}
		{{ if eq .ParamSetMethod "=" }} = {{ .ParamValue }}{{ else }}.{{ .ParamSetMethod }}({{ .ParamValue }}){{ end }}
		{{ end -}}
	{{ else if eq .FunctionType "append" -}}
		this.{{ .FunctionAppendType }}s = append(this.{{ .FunctionAppendType }}s, *New{{ .FunctionAppendType }}({{ range $i, $param := .FunctionParams }}{{if $i}}, {{end}}{{ .ParamName }}{{ end }}))
	{{ end -}}
	}
	{{ end }}

	// Type returns the type identifer for this message.
	Type(): number {
		return MsgCodes.{{.Name}};
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// Serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
	{{ range $f, $field := .PayloadFields }}

		// {{.FieldName}} ({{.FieldGoType}})
	{{- if ne (len $field.IncludeIfTrue) 0 }}
		if (this.{{ $field.IncludeIfTrue }}) {
	{{- else if ne (len $field.IncludeIfFalse) 0 }}
		if (!this.{{ $field.IncludeIfFalse }}) {
	{{- else if ne (len $field.IncludeIf.Field) 0 }}
		if ({{ range $j, $include := $field.IncludeIf.Values }}{{ if (ne $j 0) }} ||{{ end }} this.{{$field.IncludeIf.Field}} == '{{ $include }}'{{ end }}) {
	{{- else if ne (len $field.IncludeIfInt.Field) 0 }}
		if ({{ range $j, $include := $field.IncludeIfInt.Values }}{{ if (ne $j 0) }} ||{{ end }} this.{{$field.IncludeIfInt.Field}} === {{ $include }}{{ end }}) {
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

	// Write populates the fields in {{.Name}} from the byte slice
	Write(b: Buffer): number {
		const buf = new _.Reader(b);

	{{- range $f, $field := .Fields }}

		// {{.FieldName}} ({{.FieldGoType}})
	{{- if ne (len $field.IncludeIfTrue) 0 }}
		if (this.{{ $field.IncludeIfTrue }}) {
	{{- else if ne (len $field.IncludeIfFalse) 0 }}
		if (!this.{{ $field.IncludeIfFalse }}) {
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
			this.{{.SnakeCase}} = [...Array(size)].map(() => {
				const v = new {{.SingularType}}();
				v.Write(buf);
				return v;
			});
	{{- else if .IsNativeTypeArray }}
			const size = ReadVariableSize(buf, {{.Length}}, 8);
			const type = '{{.FieldGoType}}'.slice(2);
			this.{{.SnakeCase}} = [...Array(size)].map(() => read(buf, type));
	{{- else if .IsInternalType }}
			this.{{ $field.SnakeCase }} = new {{.FieldGoType}}();
			this.{{.SnakeCase}}.Write(buf);
	{{- else if or .IsBytes .IsData }}
			this.{{.SnakeCase}} = buf.read({{.Length}}, '{{.FieldGoType}}');
	{{- else if .Trimmable }}
			this.{{.SnakeCase}} = bytes.Trim(this.{{.SnakeCase}}, '\x00');
	{{- else }}
			this.{{.SnakeCase}} = read(buf, '{{.FieldGoType}}');
	{{- end }}
		}

	{{ end }}

		return 0;
	}

	Validate(): string {
	{{- range $i, $field := .Fields }}

		// {{.Name}} ({{.FieldGoType}})
	{{- if ne (len $field.IncludeIfTrue) 0 }}
		if (this.{{ $field.IncludeIfTrue }}) {
	{{- else if ne (len $field.IncludeIfFalse) 0 }}
		if (!this.{{ $field.IncludeIfFalse }}) {
	{{- else if ne (len $field.IncludeIf.Field) 0 }}
		if ({{ range $j, $include := $field.IncludeIf.Values }}{{ if (ne $j 0) }} ||{{ end }} this.{{$field.IncludeIf.Field}} == '{{ $include }}'{{ end }}) {
	{{- else }}
		{
	{{- end }}
	{{- if .IsVarChar }}
			if (this.{{.SnakeCase}}.length > (2 ** {{.Length}}) - 1) {
				return sprintf('varchar field {{.SnakeCase}} too long %d/%d', this.{{.SnakeCase}}.length, (2 ** {{.Length}}) - 1);
			}
	{{- else if .IsFixedChar }}
			if (this.{{.SnakeCase}}.length > {{.Length}}) {
				return sprintf('fixedchar field {{.SnakeCase}} too long %d/%d', this.{{.SnakeCase}}.length, {{.Length}});
			}
	{{- else if .IsVarBin }}
			if (this.{{.SnakeCase}}.length > (2 ** {{.Length}}) - 1) {
				return sprintf('varbin field {{.SnakeCase}} too long %d/%d', this.{{.SnakeCase}}.length, (2 ** {{.Length}}) - 1);
			}
	{{- else if .IsResourceTypeArray }}
			if (this.{{.SnakeCase}}.length > (2 ** {{.Length}}) - 1) {
				return sprintf('list field {{.SnakeCase}} has too many items %d/%d', this.{{.SnakeCase}}.length, (2 ** {{.Length}}) - 1);
			}

			let err = null;
			this.{{.SnakeCase}}.find(value => {
		{{- if eq .Type "RejectionCode[]" }}
				if (!Resources.GetRejectionCode(value)) {
					err = sprintf('{{.SnakeCase}}: Invalid rejection code value : %d', value);
					return true;
				}
		{{- else if eq .Type "Role[]" }}
				if (!Resources.GetRoleType(value)) {
					err = sprintf('{{.SnakeCase}}: Invalid role value : %d', value);
					return true;
				}
		{{- else if eq .Type "CurrencyType[]" }}
				if (!Resources.GetCurrency(value)) {
					err = sprintf('{{.SnakeCase}}: Invalid currency value : %d', value);
					return true;
				}
		{{- else if eq .Type "Polity[]" }}
				if (!Resources.GetPolityType(string(value[:]))) {
					err = sprintf('{{.SnakeCase}}: Invalid polity value : %d', value);
					return true;
				}
		{{- else if eq .Type "EntityType[]" }}
				if (!Resources.GetEntityType(value)) {
					err = sprintf('{{.SnakeCase}}: Invalid entity type value : \'%c\'/%d', value, value);
					return true;
				}
		{{- else if eq .Type "Tag[]" }}
				if (!Resources.GetTagType(value)) {
					err = sprintf('{{.SnakeCase}}: Invalid tag type value : \'%c\'/%d', value, value);
					return true;
				}
		{{- end }}
			});
			if (err) return err;
	{{- else if eq .Type "RejectionCode" }}
			if (!Resources.GetRejectionCode(this.{{.SnakeCase}})) {
				return sprintf('Invalid rejection code value : %d', this.{{.SnakeCase}});
			}
	{{- else if eq .Type "Role" }}
			if (!Resources.GetRoleType(this.{{.SnakeCase}})) {
				return sprintf('Invalid role value : %d', this.{{.SnakeCase}});
			}
	{{- else if eq .Type "MessageType" }}
			if (!Resources.GetMessageType(this.{{.SnakeCase}})) {
				return sprintf('Invalid message value : %d', this.{{.SnakeCase}});
			}
	{{- else if eq .Type "Currency" }}
			if (!Resources.GetCurrency(this.{{.SnakeCase}})) {
				return sprintf('Invalid currency value : %d', this.{{.SnakeCase}});
			}
	{{- else if eq .Type "Polity" }}
			if (!Resources.GetPolityType(this.{{.SnakeCase}})) {
				return sprintf('Invalid polity value : %d', this.{{.SnakeCase}});
			}
	{{- else if eq .Type "EntityType" }}
			if (!Resources.GetEntityType(this.{{.SnakeCase}})) {
				return sprintf('Invalid entity type value : \'%c\'/%d', this.{{.SnakeCase}});
			}
	{{- else if eq .Type "Tag" }}
			if (!Resources.GetTagType(this.{{.SnakeCase}})) {
				return sprintf('Invalid tag type value : \'%c\'/%d', this.{{.SnakeCase}});
			}
	{{- else if .IsInternalTypeArray }}
			if (this.{{.SnakeCase}}.length > (2 ** {{.Length}}) - 1) {
				return sprintf('list field {{.SnakeCase}} has too many items %d/%d', this.{{.SnakeCase}}.length, (2 ** {{.Length}}) - 1);
			}

			const e = this.{{.SnakeCase}}.find((value, i) => {
				const err = value.Validate();
				if (err) return sprintf('list field {{.SnakeCase}}[%d] is invalid : %s', i, err);
			});
			if (e) return e;
	{{- else if .IsNativeTypeArray }}
			if (this.{{.SnakeCase}}.length > (2 ** {{.Length}}) - 1) {
				return sprintf('list field {{.SnakeCase}} has too many items %d/%d', this.{{.SnakeCase}}.length, (2 ** {{.Length}}) - 1);
			}
	{{- else if .IsInternalType }}
			const err = this.{{.SnakeCase}}.Validate();
			if (err) return sprintf('field {{.SnakeCase}} is invalid : %s', err);
	{{ else if ne (len $field.IntValues) 0 }}
			if ({{ range $j, $value := $field.IntValues }}{{ if (ne $j 0) }} &&{{ end }} this.{{snakecase $field.Name}} != {{ $value }}{{ end }}) {
				return sprintf('field {{snakecase $field.Name}} value is invalid : %d', this.{{snakecase $field.Name}});
			}
	{{ else if ne (len $field.CharValues) 0 }}
			if ({{ range $j, $value := $field.CharValues }}{{ if (ne $j 0) }} &&{{ end }} this.{{snakecase $field.Name}} != '{{ $value }}'{{ end }}) {
				return sprintf('field {{snakecase $field.Name}} value is invalid : %d', this.{{snakecase $field.Name}});
			}
	{{- end }}
		}
	{{ end }}
		return null;
	}

	toString(): string {
		const vals: string[] = [];
	{{ range .Fields -}}
		{{- if eq .Type "STRING" }}
		vals.push(sprintf('{{.SnakeCase}}:\'%v\'', 'this.{{.SnakeCase}}'));
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
	{{- end}}


