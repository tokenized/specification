# {{.Metadata.Name}}

{{.Metadata.Description}}

{{$letter := .TypeLetter}}

<div class="ritz grid-container" dir="ltr">
    <table class="waffle" cellspacing="0" cellpadding="0" table-layout=fixed width=100%>
         <tr style='height:19px;'>
            <th style="width:20%" class="s1">Field</th>
            <th style="width:10%" class="s1">Type</th>
            <th style="width:15%" class="s1">Description</th>
            <th style="width:20%" class="s1">Size</th>
            <th style="width:20%" class="s1">Example</th>
            <th class="s1">Notes</th>
        </tr>
{{- range .Fields}}
        <tr>
            <td class="{{$letter}}10">{{.Label}}</td>
            <td class="{{$letter}}10">{{.Type}}</td>
            <td class="{{$letter}}10">{{.Description}}</td>
            <td class="{{$letter}}10">{{.Size}}</td>
            <td class="{{$letter}}10">{{.ExampleValue}}</td>
            <td class="{{$letter}}10">{{.Notes}}</td>
        </tr>
{{- end}}
    </table>
</div>
