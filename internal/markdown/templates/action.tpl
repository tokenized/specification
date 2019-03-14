{{$letter := .TypeLetter}}
{{$code := .ActionCode}}
{{$description := .CodeComment}}
# {{.Label}}

{{.Metadata.Description}}

The following breaks down the construction of a {{.Label}}. The action is constructed by building a single string from each of the elements in order.

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
            <td class="{{$letter}}10">{{.FieldBytes}}</td>
            <td class="{{$letter}}10">{{if gt (len .ExampleValue) 32}}<abbr title="{{.ExampleValue}}">Hover for example</abbr>{{else}}{{.ExampleValue}}{{end}}</td>
            <td class="{{$letter}}10">{{if gt (len .Description) 96}}<abbr title="{{.Description}}">{{.DescriptionAbbr}} ...</abbr>{{else}}{{.Description}}{{end}}</td>
            <td class="{{$letter}}10">{{.Type}}</td>
            <td class="{{$letter}}10">{{.Notes}}</td>
{{- end}}
        </tr>
{{- end}}
    </table>
</div>

##{{.Label}} Transaction Summary

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
       <tr>
            <td class="{{$letter}}5">{{.Rules.Inputs}}</td>
            <td class="{{$letter}}6">.</td>
            <td class="{{$letter}}6">.</td>
            <td class="{{$letter}}10">.</td>
            <td class="{{$letter}}10">.</td>
            <td class="{{$letter}}10">.</td>
        </tr>
    </table>
</div>

{{$header := 1}}
{{range .FieldTypes}}
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
                <td class="{{$letter}}10">{{.FieldBytes}}</td>
                <td class="{{$letter}}10" style="word-break:break-all">{{if and (eq $header 1) (eq .Name "ActionPrefix")}}{{$code}}{{else}}{{.ExampleValue}}{{end}}</td>
                <td class="{{$letter}}10">{{if and (eq $header 1) (eq .Name "ActionPrefix")}}{{$description}}{{$header = 0}}{{else}}{{.Description}}{{end}}</td>
                <td class="{{$letter}}10">{{.Type}}</td>
                <td class="{{$letter}}10">{{.Notes}}</td>
            </tr>{{end}}
        </table>
    </div>
</div>
{{end}}
