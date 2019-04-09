{{$actions := index . 0 -}}
{{$fieldTypes := index . 1 -}}

# Protocol Actions

- [Introduction](#introduction)
- [Available Actions](#available-actions)
- [Field Types](#field-types)
    - [Primitive Types](#primitive-types)
    - [Basic Types](#basic-types)
    - [Compound Types](#compound-types)

<a name="introduction"></a>
## Introduction

The following actions break down the construction of a Tokenized protocol message. The action is constructed by building a single string from each of the elements in order. Each field within the action is given a specific type, including standard types and compound types.

Each message should be prefixed with a common header data. This can be changed by issuer or operator at their discretion. See the article on [how to construct a header](#) for more details.

<a name="available-actions"></a>
## Available Actions

<div class="content-list collection-method-list" markdown="1">
{{- range $actions }}
- [{{.Metadata.Label}}](#{{.URLCode}})
{{- end }}
</div>

{{- range $actions}}

<a name="{{.URLCode}}"></a>
#### {{.Metadata.Label}}

{{.Metadata.Description}}

<table>
    <tr>
        <th style="width:15%">Field</th>
        <th style="width:15%">Type</th>
        <th>Description</th>
    </tr>
    {{- range .Fields}}
    <tr>
        <td>{{.Name}}</td>
        <td>
        {{- if .IsComplexType }}
            <a href="#{{.TypeURLCode}}">{{.Type}}</a>
        {{- else}}
            {{.Type}}{{ if ne .Size 0 }}({{.Size}}){{ end }}
        {{- end}}
        </td>
        <td>{{.Label}} {{.Description}} {{.Notes}}</td>
    </tr>
    {{- end}}
</table>

##### Transaction Summary

{{if .Rules.Inputs}}
<table>
   <tr>
        <th style="width:5%" class="text-center">Index</th>
        <th style="width:30%">Input</th>
        <th>Description</th>
   </tr>
    {{- range $index, $val := .Rules.Inputs}}
   <tr>
        <td class="text-center">{{$index}}</td>
        <td>{{$val.Label}}</td>
        <td>{{$val.Comments}}</td>
    </tr>
    {{- end}}
</table>
{{end}}

{{if .Rules.Outputs}}
<table>
   <tr>
        <th style="width:5%" class="text-center">Index</th>
        <th style="width:30%">Output</th>
        <th>Description</th>
   </tr>
    {{- range $index, $val := .Rules.Outputs}}
   <tr>
        <td class="text-center">{{$index}}</td>
        <td>{{$val.Label}}</td>
        <td>{{$val.Comments}}</td>
    </tr>
    {{- end}}
</table>
{{end}}

{{- if ne .Rules.Fee 0 }}
> **Note**: This action requires a Smart Contract Operator Fee of at least <strong>{{.Rules.Fee}}</strong>
{{- end }}

<hr />

{{end}}

<a name="field-types"></a>
## Field Types

Each field in a protocol action si assigned with a data type. Standard scalar types have a single value, these include primitive and basic types. While compound types are made up of fields that have their own types, including nested compound types.

<a name="primitive-types"></a>
### Primitive Types

<table>
   <tr>
        <th style="width:15%">Type</th>
        <th>Description</th>
   </tr>
    <tr><td>int</td><td>is a signed integer. <code>size</code> is the number of bits for the integer. Valid values for <code>size</code> are 8, 16, 32, 64.</td></tr>
    <tr><td>uint</td><td>is an unsigned integer. <code>size</code> is the number of bits for the integer. Valid values for <code>size</code> are 8, 16, 32, 64.</td></tr>
    <tr><td>float</td><td>is a floating point number. <code>size</code> is the number of bits for the float. Valid values for <code>size</code> are 32 and 64.</td></tr>
    <tr><td>bool</td><td>is a boolean stored as 1 byte.</td></tr>
    <tr><td>bin</td><td>is fixed length binary data. <code>size</code> is the length in bytes of the data.</td></tr>
    <tr>
        <td>varbin</td>
        <td>
            is variable length binary data.
            The data is preceded by an integer that specifies the actual length in bytes.
            <code>size</code> is the number of bits used to serialize the length in bytes of the data.
            Valid values for <code>size</code> are 8, 16, 32, and 64.
            For example, a <code>varbin</code> value with a <code>size</code> of 8 will be able to contain up to 255 bytes and will be preceded by a 1 byte unsigned integer specifying the length in bytes.
        </td>
    </tr>
    <tr>
        <td>fixedchar</td>
        <td>
            is fixed length text data.
            The data is assumed to be UTF-8 unless the first two bytes are a valid UTF-16 BOM (Byte Order Method).
            <code>size</code> is the length in bytes of the data.
        </td>
    </tr>
    <tr>
        <td>varchar</td>
        <td>
            is variable length text data.
            The data is assumed to be UTF-8 unless the first two bytes are a valid UTF-16 BOM (Byte Order Method).
            The data is preceded by an integer that specifies the actual length in bytes.
            <code>size</code> is the number of bits used to serialize the length in bytes of the data.
            Valid values for <code>size</code> are 8, 16, 32, and 64.
            For example, a `varchar` value with a <code>size</code> of 8 will be able to contain up to 255 bytes and will be preceded by a 1 byte unsigned integer specifying the length in bytes.
        </td>
    </tr>
</table>

<a name="basic-types"></a>
### Basic Types

<a name="type-op-return-message"></a>
##### Op Return Message

`OpReturnMessage` implements a base interface for all message types.

Provides the ability to serialize the data as a Bitcoin OP_RETURN script and to request the variable payload data.

<a name="type-payload-message"></a>
##### Payload Message

`PayloadMessage` is the interface for messages that are derived from payloads, such as asset types and message types.

Provides the ability to serialize the data.

<a name="type-tx-id"></a>
##### Tx Id

`TxId` represents a Bitcoin transaction ID.

It is the double SHA256 hash of the serialized transaction.

`size` does not need to be specified and is always 32 bytes.

<a name="type-public-key-hash"></a>
##### Public Key Hash

`PublicKeyHash` represents a Bitcoin Public Key Hash.

It is the RIPEMD160 hash of the SHA256 of the serialized public key.

Public key hashes are used as an "address" to send/receive transactions, tokens, and messages.

`size` does not need to be specified and is always 20 bytes.

<a name="type-asset-code"></a>
##### Asset Code

`AssetCode` represents a unique identifier for an asset/token.

`size` does not need to be specified and is always 32 bytes.

<a name="type-contract-code"></a>
##### Contract Code

`ContractCode` represents a unique identifier for a static contract.

The `size` does not need to be specified and is always 32 bytes.

<a name="type-timestamp"></a>
##### Timestamp

`Timestamp` represents a time. `size` does not need to be specified and is always 8 bytes.

<a name="compound-types"></a>
### Compound Types

<div class="content-list collection-method-list" markdown="1">
{{- range $fieldTypes }}
- [{{.Metadata.Label}}](#{{.URLCode}})
{{- end }}
</div>

{{range $fieldTypes}}

<a name="{{.URLCode}}"></a>
#### {{.Metadata.Label}}

{{.Metadata.Description}}

<table>
    <tr>
        <th style="width:15%">Field</th>
        <th style="width:15%">Type</th>
        <th>Description</th>
    </tr>
    {{- range .Fields}}
    <tr>
        <td>{{.Name}}</td>
        <td>
        {{- if .IsComplexType }}
            <a href="#{{.TypeURLCode}}">{{.Type}}</a>
        {{- else}}
            {{.Type}}{{ if ne .Size 0 }}({{.Size}}){{ end }}
        {{- end}}
        </td>
        <td>{{.Label}} {{.Description}} {{.Notes}}</td>
    </tr>
    {{- end}}
</table>

{{end}}
