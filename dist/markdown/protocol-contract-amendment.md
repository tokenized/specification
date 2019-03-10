
# Contract Amendment Action

Contract Amendment Action - the issuer can initiate an amendment to the contract establishment metadata. The ability to make an amendment to the contract is restricted by the Authorization Flag set on the current revision of Contract Formation action.

The following breaks down the construction of a Contract Amendment Action. The action is constructed by building a single string from each of the elements in order.

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
		<td class="c6">Header[]</td>
		<td class="c6">Header Array</td>
		<td class="c6">-</td>
		<td class="c6">-</td>
		<td class="c6">Common header data for all actions</td>
		<td class="c6">Header</td>
		<td class="c7"></td>
	</tr>

	<tr>
		<td class="c10">Text Encoding</td>
		<td class="c10">TextEncoding</td>
		<td class="c10">1</td>
		<td class="c10" style="word-break:break-all">0</td>
		<td class="c10"> 0 = ASCII, 1 = UTF-8, 2 = UTF-16, 3 = Unicode.  Encoding applies to all 'text' data types. All 'string' types will always be encoded with ASCII.  Where string is selected, all fields will be ASCII.</td>
		<td class="c10">uint8</td>
		<td class="c11">Can be changed by Issuer or Operator at their discretion.</td>
	</tr>

	<tr>
		<td class="c10">Change Issuer Address</td>
		<td class="c10">ChangeIssuerAddress</td>
		<td class="c10">1</td>
		<td class="c10" style="word-break:break-all">1</td>
		<td class="c10">1 - Yes, 0 - No.  Used to change the issuer address.  The new issuer address must be in the input[1] position.</td>
		<td class="c10">bool</td>
		<td class="c11"></td>
	</tr>

	<tr>
		<td class="c10">Change Operator Address</td>
		<td class="c10">ChangeOperatorAddress</td>
		<td class="c10">1</td>
		<td class="c10" style="word-break:break-all">1</td>
		<td class="c10">1 - Yes, 0 - No.  Used to change the smart contract operator address.  The new operator address must be in the input[1] position.</td>
		<td class="c10">bool</td>
		<td class="c11"></td>
	</tr>

	<tr>
		<td class="c10">Contract Revision</td>
		<td class="c10">ContractRevision</td>
		<td class="c10">2</td>
		<td class="c10" style="word-break:break-all">42</td>
		<td class="c10">Counter 0 - 65,535</td>
		<td class="c10">uint16</td>
		<td class="c11"></td>
	</tr>

	<tr>
		<td class="c10">AmendmentsCount</td>
		<td class="c10">AmendmentsCount</td>
		<td class="c10">1</td>
		<td class="c10" style="word-break:break-all">0</td>
		<td class="c10">Number of Amendments. Must be less than the max Subfield Index of CF.</td>
		<td class="c10">uint8</td>
		<td class="c11"></td>
	</tr>

	<tr>
		<td class="c10">Amendments</td>
		<td class="c10">Amendments</td>
		<td class="c10">0</td>
		<td class="c10" style="word-break:break-all"></td>
		<td class="c10"></td>
		<td class="c10">Amendment[]</td>
		<td class="c11"></td>
	</tr>

	<tr>
		<td class="c10">Ref Tx-ID</td>
		<td class="c10">RefTxID</td>
		<td class="c10">32</td>
		<td class="c10" style="word-break:break-all">a8700385d4cc62628cc34629862121f84e6237689de8e45e151dcbc8cf30b33d</td>
		<td class="c10">Tx-ID of the associated Result action (governance) that permitted the modifications.</td>
		<td class="c10">SHA256</td>
		<td class="c11"></td>
	</tr>

</table>
