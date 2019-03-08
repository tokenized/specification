
# Establishment Action

Establishment Action -  Establishes an on-chain register.

The following breaks down the construction of a Establishment Action. The action is constructed by building a single string from each of the elements in order.

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
		<td class="r6">Header[]</td>
		<td class="r6">Header Array</td>
		<td class="r6">-</td>
		<td class="r6">-</td>
		<td class="r6">Common header data for all actions</td>
		<td class="r6">Header</td>
		<td class="r7"></td>
	</tr>

	<tr>
		<td class="r10">Text Encoding</td>
		<td class="r10">TextEncoding</td>
		<td class="r10">1</td>
		<td class="r10" style="word-break:break-all">0</td>
		<td class="r10"> 0 = ASCII, 1 = UTF-8, 2 = UTF-16, 3 = Unicode.  Encoding applies to all &#39;text&#39; data types. All &#39;string&#39; types will always be encoded with ASCII.  Where string is selected, all fields will be ASCII.</td>
		<td class="r10">uint8</td>
		<td class="r11"></td>
	</tr>

	<tr>
		<td class="r10">Message</td>
		<td class="r10">Message</td>
		<td class="r10">25</td>
		<td class="r10" style="word-break:break-all">North America Whitelist</td>
		<td class="r10">Length only limited by Bitcoin protocol.</td>
		<td class="r10">nvarchar64</td>
		<td class="r11"></td>
	</tr>

</table>
