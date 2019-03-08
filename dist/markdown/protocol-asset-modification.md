
# Asset Modification Action

Asset Modification Action -  Token Dilutions, Call Backs/Revocations, burning etc.

The following breaks down the construction of a Asset Modification Action. The action is constructed by building a single string from each of the elements in order.

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
		<td class="a6">Header[]</td>
		<td class="a6">Header Array</td>
		<td class="a6">-</td>
		<td class="a6">-</td>
		<td class="a6">Common header data for all actions</td>
		<td class="a6">Header</td>
		<td class="a7"></td>
	</tr>

	<tr>
		<td class="a10">Text Encoding</td>
		<td class="a10">TextEncoding</td>
		<td class="a10">1</td>
		<td class="a10" style="word-break:break-all">0</td>
		<td class="a10"> 0 = ASCII, 1 = UTF-8, 2 = UTF-16, 3 = Unicode.  Encoding applies to all &#39;text&#39; data types. All &#39;string&#39; types will always be encoded with ASCII.  Where string is selected, all fields will be ASCII.</td>
		<td class="a10">uint8</td>
		<td class="a11">Can be changed by Issuer or Operator at their discretion.</td>
	</tr>

	<tr>
		<td class="a10">Asset Type</td>
		<td class="a10">AssetType</td>
		<td class="a10">3</td>
		<td class="a10" style="word-break:break-all">SHC</td>
		<td class="a10">eg. Share - Common</td>
		<td class="a10">string</td>
		<td class="a11"></td>
	</tr>

	<tr>
		<td class="a10">Asset ID</td>
		<td class="a10">AssetID</td>
		<td class="a10">32</td>
		<td class="a10" style="word-break:break-all">apm2qsznhks23z8d83u41s8019hyri3i</td>
		<td class="a10">Randomly generated base58 string.  Each Asset ID should be unique.  However, a Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type &#43; Asset ID = Asset Code.  An Asset Code is a human readable idenitfier that can be used in a similar way to a Bitcoin (BSV) address, a vanity identifying label.</td>
		<td class="a10">string</td>
		<td class="a11"></td>
	</tr>

	<tr>
		<td class="a10">Asset Revision</td>
		<td class="a10">AssetRevision</td>
		<td class="a10">8</td>
		<td class="a10" style="word-break:break-all">0</td>
		<td class="a10">Counter. (Subfield cannot be manually changed by Asset Modification Action.  Only SC can increment by 1 with each AC action. SC will reject AM actions where the wrong asset revision has been selected. </td>
		<td class="a10">uint64</td>
		<td class="a11">Cannot be Amended</td>
	</tr>

	<tr>
		<td class="a10">ModificationCount</td>
		<td class="a10">ModificationCount</td>
		<td class="a10">1</td>
		<td class="a10" style="word-break:break-all">0</td>
		<td class="a10">Number of Modifications. Must be less than the max Subfield Index of CF.</td>
		<td class="a10">uint8</td>
		<td class="a11"></td>
	</tr>

	<tr>
		<td class="a10">Modifications</td>
		<td class="a10">Modifications</td>
		<td class="a10">0</td>
		<td class="a10" style="word-break:break-all"></td>
		<td class="a10"></td>
		<td class="a10">Amendment[]</td>
		<td class="a11"></td>
	</tr>

	<tr>
		<td class="a10">Ref Tx-ID</td>
		<td class="a10">RefTxID</td>
		<td class="a10">32</td>
		<td class="a10" style="word-break:break-all">a8700385d4cc62628cc34629862121f84e6237689de8e45e151dcbc8cf30b33d</td>
		<td class="a10">Tx-ID of the associated Result action (governance) that permitted the modifications.</td>
		<td class="a10">sha256</td>
		<td class="a11"></td>
	</tr>

</table>
