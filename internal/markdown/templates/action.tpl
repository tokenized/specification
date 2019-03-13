{{$letter := .TypeLetter}}
# {{.Label}}

{{.Metadata.Description}}

The following breaks down the construction of a {{.Label}}. The action is constructed by building a single string from each of the elements in order.

<div class="ritz grid-container" dir="ltr">
    <table class="waffle" cellspacing="0" cellpadding="0" table-layout=fixed width=100%>
         <tr style='height:19px;'>
            <th style="width:9%" class="s1">Label</th>
            <th style="width:9%" class="s1">Name</th>
            <th style="width:2%" class="s1">Bytes</th>
            <th style="width:29%" class="s1">Example Values</th>
            <th style="width:26%" class="s1">Comments</th>
            <th style="width:5%" class="s1">Data Type</th>
            <th style="width:14%" class="s1">Amendment Restrictions</th>
        </tr>

        <tr>
            <td class="{{$letter}}6" colspan="7">
                <a href="javascript:;" data-popover="type-Header">
                   Header - Click to show content
                </a>
             </td>
        </tr>
{{range .PayloadFields}}
        <tr>
{{- if .IsComplexType }}
            <td class="{{$letter}}6" colspan="7">
                <a href="javascript:;" data-popover="type-{{.SingularType}}">
                   {{.Label}} - Click to show content
                </a>
            </td>
{{- else}}
            <td class="{{$letter}}10">{{.Label}}</td>
            <td class="{{$letter}}10">{{.Name}}</td>
            <td class="{{$letter}}10">{{.FieldBytes}}</td>
            <td class="{{$letter}}10"><abbr {{.ExampleValue}}>Hover</abbr></td>
            <td class="{{$letter}}10">{{.Description}}</td>
            <td class="{{$letter}}10">{{.Type}}</td>
            <td class="{{$letter}}10">{{.Notes}}</td>
{{- end}}
        </tr>
{{end}}
    </table>
</div>

{{range .FieldTypes}}
<div class="ui modal" id="type-{{.Metadata.Name}}">
    <i class="close icon"></i>
    <div class="content docs-content">
        <table class="ui table">
            <tr style='height:19px;'>
                <th style="width:9%" class="s1">Label</th>
                <th style="width:9%" class="s1">Name</th>
                <th style="width:2%" class="s1">Bytes</th>
                <th style="width:29%" class="s1">Example Values</th>
                <th style="width:26%" class="s1">Comments</th>
                <th style="width:5%" class="s1">Data Type</th>
                <th style="width:14%" class="s2">Amendment Restrictions</th>
            </tr>{{range .Fields}}
            <tr>
                <td class="{{$letter}}10">{{.Label}}</td>
                <td class="{{$letter}}10">{{.Name}}</td>
                <td class="{{$letter}}10">{{.FieldBytes}}</td>
                <td class="{{$letter}}10" style="word-break:break-all">{{.ExampleValue}}</td>
                <td class="{{$letter}}10">{{.Description}}</td>
                <td class="{{$letter}}10">{{.Type}}</td>
                <td class="{{$letter}}10">{{.Notes}}</td>
            </tr>{{end}}
        </table>
    </div>
</div>
{{end}}
