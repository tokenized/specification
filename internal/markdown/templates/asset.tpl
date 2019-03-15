# {{.Metadata.Name}}

{{.Metadata.Description}}

{{$letter := "a"}}

<div class="ritz grid-container" dir="ltr">
    <table class="waffle" cellspacing="0" cellpadding="0" table-layout=fixed width=100%>
         <tr style='height:19px;'>
            <th style="width:9%" class="s0">Name</th>
            <th style="width:9%" class="s0">Label</th>
            <th style="width:9%" class="s1">Type</th>
            <th style="width:2%" class="s1">Description</th>
            <th style="width:25%" class="s1">Size</th>
            <th style="width:36%" class="s1">Example</th>
            <th style="width:5%" class="s1">Computed</th>
            <th style="width:5%" class="s1">Display Order</th>
            <th style="width:5%" class="s1">Required</th>
        </tr>
{{- range .Fields}}
        <tr>
            <td class="{{$letter}}10">{{.Name}}</td>
            <td class="{{$letter}}9">{{.Label}}</td>
            <td class="{{$letter}}10">{{.Type}}</td>
            <td class="{{$letter}}10">{{.Description}}</td>
            <td class="{{$letter}}10">{{.Size}}</td>
            <td class="{{$letter}}10">{{.ExampleValue}}</td>
            <td class="{{$letter}}10">{{.Computed}}</td>
            <td class="{{$letter}}10">{{.DisplayOrder}}</td>
            <td class="{{$letter}}10">{{.Required}}</td>
        </tr>
{{- end}}
    </table>
</div>
