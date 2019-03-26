{{$letter := .TypeLetter}}
{{$code := .Code}}
{{$description := .CodeComment}}
# {{.Metadata.Label}}

{{.Metadata.Description}}

The following breaks down the construction of a {{.Metadata.Label}}. The action is constructed by building a single string from each of the elements in order.

<div class="ritz grid-container" dir="ltr">
    <table class="waffle" cellspacing="0" cellpadding="0" table-layout=fixed width=100%>
         <tr style='height:19px;'>
            <th style="width:9%" class="s0">Label</th>
            <th style="width:9%" class="s1">Name</th>
            <th style="width:2%" class="s1">Bytes</th>
            <th style="width:25%" class="s1">Example Values</th>
            <th style="width:36%" class="s1">Comments</th>
            <th style="width:5%" class="s1">Data Type</th>
            <th class="s1">Amendment Restrictions</th>
        </tr>
        <tr>
            <td class="{{$letter}}5" colspan="7">
                <a href="javascript:;" data-popover="type-Header">
                   Header - Click to show content
                </a>
             </td>
        </tr>
{{- range .PayloadFields}}
        <tr>
{{- if .IsComplexType }}
            <td class="{{$letter}}5" colspan="7">
                <a href="javascript:;" data-popover="type-{{.SingularType}}">
                   {{.Label}} - Click to show content
                </a>
            </td>
{{- else}}
            <td class="{{$letter}}9">{{.Label}}</td>
            <td class="{{$letter}}10">{{.Name}}</td>
            <td class="{{$letter}}10">{{.Size}}</td>
            <td class="{{$letter}}10">{{if gt (len .ExampleValue) 32}}<abbr title="{{.ExampleValue}}">Hover for example</abbr>{{else}}{{.ExampleValue}}{{end}}</td>
            <td class="{{$letter}}10">{{if gt (len .Description) 96}}<abbr title="{{.Description}}">{{.DescriptionAbbr}} ...</abbr>{{else}}{{.Description}}{{end}}</td>
            <td class="{{$letter}}10">{{.Type}}</td>
            <td class="{{$letter}}10">{{.Notes}}</td>
{{- end}}
        </tr>
{{- end}}
    </table>
</div>

##{{.Metadata.Label}} Transaction Summary

<div class="ritz grid-container" dir="ltr">
    <table class="waffle" cellspacing="0" cellpadding="0" table-layout=fixed width=100%>
       <tr style='height:19px;'>
            <th class="s0" colspan="6">Smart Contract Operator Fee: {{.Rules.Fee}}</th>
       </tr>
       <tr style='height:19px;'>
            <th style="width:10%" class="s0">Index (input)</th>
            <th style="width:20%" class="s1">Txn inputs</th>
            <th style="width:20%" class="s1">Comments</th>
            <th style="width:10%" class="s1">Index (output)</th>
            <th style="width:20%" class="s1">Txn outputs</th>
            <th class="s1">Comments</th>
       </tr>

{{range .Rules.Rows}}
       <tr>
            <td class="{{$letter}}5">{{.InputIndex}}</td>
            <td class="{{$letter}}6">{{.Input.Label}}</td>
            <td class="{{$letter}}6">{{.Input.Comments}}</td>
            <td class="{{$letter}}10">{{.OutputIndex}}</td>
            <td class="{{$letter}}10">{{.Output.Label}}</td>
            <td class="{{$letter}}10">{{.Output.Comments}}</td>
        </tr>
{{end}}
    </table>
</div>

<div class="ui modal" id="type-Header">
    <i class="close icon"></i>
    <div class="content docs-content">
        <table class="ui table">
            <tr style='height:19px;'>
                <th style="width:5%" class="s1">Label</th>
                <th style="width:9%" class="s1">Name</th>
                <th style="width:3%" class="s1">Bytes</th>
                <th style="width:33%" class="s1">Example Values</th>
                <th style="width:26%" class="s1">Comments</th>
                <th style="width:5%" class="s1">Data Type</th>
                <th class="s2">Amendment Restrictions</th>
            </tr>
            <tr>
                <td class="{{$letter}}10">Protocol Identifier</td>
                <td class="{{$letter}}10">ProtocolID</td>
                <td class="{{$letter}}10">13</td>
                <td class="{{$letter}}10">tokenized.com</td>
                <td class="{{$letter}}10" style="word-break:break-all">Tokenized ID Prefix. tokenized.com</td>
                <td class="{{$letter}}10">byte</td>
                <td class="{{$letter}}10"></td>
            </tr>
            <tr>
                <td class="{{$letter}}10">Push Data</td>
                <td class="{{$letter}}10">OpPushdata</td>
                <td class="{{$letter}}10">1</td>
                <td class="{{$letter}}10">0x4d</td>
                <td class="{{$letter}}10" style="word-break:break-all">PACKET LENGTH, PUSHDATA1 (76), PUSHDATA2 (77), or PUSHDATA4 (78) depending on total size of action payload.</td>
                <td class="{{$letter}}10">byte</td>
                <td class="{{$letter}}10"></td>
            </tr>
            <tr>
                <td class="{{$letter}}10">Length of Action Payload</td>
                <td class="{{$letter}}10">LenActionPayload</td>
                <td class="{{$letter}}10"></td>
                <td class="{{$letter}}10">tokenized.com</td>
                <td class="{{$letter}}10" style="word-break:break-all">Length of the action message (0 - 65,535 bytes). 0 if pushdata length <76B, 1 byte if PUSHDATA1 is used, 2 bytes if PUSHDATA2 and 4 bytes if PUSHDATA4.</td>
                <td class="{{$letter}}10">byte</td>
                <td class="{{$letter}}10">Size depends on Action Payload.</td>
            </tr>
            <tr>
                <td class="{{$letter}}10">Version</td>
                <td class="{{$letter}}10">Version</td>
                <td class="{{$letter}}10">1</td>
                <td class="{{$letter}}10">0</td>
                <td class="{{$letter}}10" style="word-break:break-all">255 reserved for additional versions. Tokenized protocol versioning.</td>
                <td class="{{$letter}}10">uint8</td>
                <td class="{{$letter}}10">Can be changed by Issuer or Operator at their discretion.  Smart Contract will reject if it hasn't been updated to interpret the specified version.</td>
            </tr>
            <tr>
                <td class="{{$letter}}10">Action Prefix</td>
                <td class="{{$letter}}10">ActionPrefix</td>
                <td class="{{$letter}}10">2</td>
                <td class="{{$letter}}10">{{$code}}</td>
                <td class="{{$letter}}10" style="word-break:break-all">{{$description}}</td>
                <td class="{{$letter}}10">string</td>
                <td class="{{$letter}}10">Cannot be changed by issuer, operator or smart contract..</td>
            </tr>
        </table>
    </div>
</div>

{{- $header := 1 -}}
{{range .FieldTypesWithoutHeader}}
<div class="ui modal" id="type-{{.Metadata.Name}}">
    <i class="close icon"></i>
    <div class="content docs-content">
        <table class="ui table">
            <tr style='height:19px;'>
                <th style="width:5%" class="s1">Label</th>
                <th style="width:9%" class="s1">Name</th>
                <th style="width:3%" class="s1">Bytes</th>
                <th style="width:33%" class="s1">Example Values</th>
                <th style="width:26%" class="s1">Comments</th>
                <th style="width:5%" class="s1">Data Type</th>
                <th class="s2">Amendment Restrictions</th>
            </tr>{{range .Fields}}
            <tr>
                <td class="{{$letter}}10">{{.Label}}</td>
                <td class="{{$letter}}10">{{.Name}}</td>
                <td class="{{$letter}}10">{{.Size}}</td>
                <td class="{{$letter}}10" style="word-break:break-all">{{if and (eq $header 1) (eq .Name "ActionPrefix")}}{{$code}}{{else}}{{.ExampleValue}}{{end}}</td>
                <td class="{{$letter}}10">{{if and (eq $header 1) (eq .Name "ActionPrefix")}}{{$description}}{{$header = 0}}{{else}}{{.Description}}{{end}}</td>
                <td class="{{$letter}}10">{{.Type}}</td>
                <td class="{{$letter}}10">{{.Notes}}</td>
            </tr>{{end}}
        </table>
    </div>
</div>
{{end}}
