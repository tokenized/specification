{{$letter := .TypeLetter}}
<html>
	<head>
		<link rel="stylesheet" href="css/style.css">
		<H1>{{.MetaLabel}}</H1>
		<p>
		{{.MetaDescription}}<br><br>
		The following breaks down the construction of a {{.MetaLabel}}. The action is constructed by building a single string from each of the elements in order.
		</p>
	</head>
	<div class="ritz grid-container" dir="ltr">
		<body>
			<table class="waffle" cellspacing="0" cellpadding="0" table-layout=fixed width=100%>
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
					<td class="s5" rowspan="{{.FieldCount}}">Metadata (OP_RETURN Payload)</td>
			    	<td class="{{$letter}}6">Header[]</td>
			    	<td class="{{$letter}}6">Header Array</td>
			    	<td class="{{$letter}}6">-</td>
			    	<td class="{{$letter}}6">-</td>
			    	<td class="{{$letter}}6">Common header data for all messages</td>
			    	<td class="{{$letter}}6">Header</td>
			    	<td class="{{$letter}}7"></td>
			    </tr>
	{{range .PayloadFields}}				<tr>
			    	<td class="{{$letter}}10">{{.FieldLabel}}</td>
			    	<td class="{{$letter}}10">{{.FieldName}}</td>
			    	<td class="{{$letter}}10">{{.FieldBytes}}</td>
			    	<td class="{{$letter}}10" style="word-break:break-all">{{.FieldEG}}</td>
			    	<td class="{{$letter}}10">{{.FieldComments}}</td>
			    	<td class="{{$letter}}10">{{.FieldType}}</td>
			    	<td class="{{$letter}}11">{{.FieldRestrictions}}</td>
				</tr>{{end}}
			</table>
		</body>
	</div>
</html>