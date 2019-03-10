{{$letter := .TypeLetter}}
# {{.Label}}

{{.Description}}

The following breaks down the construction of a {{.Label}}. The action is constructed by building a single string from each of the elements in order.

| Field    | Label    | Name         | Example Values | Comments | Data Type          | Restrictions |
|----------|----------|--------------|----------------|----------|--------------------|--------------|
| Metadata | Header[] | Header Array | -              | -        | Common header data | Header       |
{{ range .PayloadFields -}}
| {{.Label}} | {{.FieldName}} | {{.FieldBytes}} | {{.ExampleValue}} | {{.Description}} | {{.Type}} | {{.Notes}} |
{{ end }}


<!--
<table class="waffle">
    <tr style='height:19px;'>
        <th style="width:6%" class="s0">Field</th>
        <th style="width:9%" class="s1">Label</th>
        <th style="width:9%" class="s1">Name</th>
        <th style="width:2%" class="s1">Bytes</th>
        <th style="width:29%" class="s1">Example Values</th>
        <th style="width:26%" class="s1">Comments</th>
        <th style="width:5%" class="s1">Data Type</th>
        <th style="width:14%" class="s2">Amendment Restrictions</th>
    </tr>
    <tr>
        <td class="s5" rowspan="100">Metadata (OP_RETURN Payload)</td>
        <td class="{{$letter}}6">Header[]</td>
        <td class="{{$letter}}6">Header Array</td>
        <td class="{{$letter}}6">-</td>
        <td class="{{$letter}}6">-</td>
        <td class="{{$letter}}6">Common header data for all actions</td>
        <td class="{{$letter}}6">Header</td>
        <td class="{{$letter}}7"></td>
    </tr>
{{range .PayloadFields}}
    <tr>
        <td class="{{$letter}}10">{{.Label}}</td>
        <td class="{{$letter}}10">{{.FieldName}}</td>
        <td class="{{$letter}}10">{{.FieldBytes}}</td>
        <td class="{{$letter}}10" style="word-break:break-all">{{.ExampleValue}}</td>
        <td class="{{$letter}}10">{{.Description}}</td>
        <td class="{{$letter}}10">{{.Type}}</td>
        <td class="{{$letter}}11">{{.Notes}}</td>
    </tr>
{{end}}
</table>
!-->