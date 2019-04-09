# Protocol Assets

- [Introduction](#introduction)
- [Smart Contracts](#smart-contracts)
    - [Contract Operator](#contract-operator)
- [Static Contracts](#static-contracts)

<a name="introduction"></a>
## Introduction

Bla bla bla bla bla bla bla bla blabla bla bla blabla bla bla blabla bla bla 
blabla bla bla blabla bla bla blabla bla bla blabla bla bla blabla bla bla blabla bla bla
blabla bla bla blabla bla bla blabla bla bla blabla bla bla blabla bla bla blabla bla bla
blabla bla bla blabla bla bla blabla bla bla blabla bla bla b

Blabla bla bla blabla bla bla blabla bla bla blabla bla bla blabla bla bla blabla bla bl
a blabla bla bla blabla bla bla bla

{{- range .}}

# {{.Metadata.Name}}

{{.Metadata.Description}}

{{$letter := "a"}}

<div class="ritz grid-container" dir="ltr">
    <table class="waffle" cellspacing="0" cellpadding="0" table-layout=fixed width=100%>
         <tr style='height:19px;'>
            <th style="width:18%" class="s1">Field</th>
            <th style="width:9%" class="s1">Type</th>
            <th style="width:15%" class="s1">Description</th>
            <th style="width:20%" class="s1">Size</th>
            <th class="s1">Example</th>
        </tr>
{{- range .Fields}}
        <tr>
            <td class="{{$letter}}10">{{.Label}}</td>
            <td class="{{$letter}}10">{{.Type}}</td>
            <td class="{{$letter}}10">{{.Description}}</td>
            <td class="{{$letter}}10">{{.Size}}</td>
            <td class="{{$letter}}10">{{.ExampleValue}}</td>
        </tr>
{{- end}}
    </table>
</div>

{{ end }}
