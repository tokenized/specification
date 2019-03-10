
# Settlement Action

Settlement Action -  Finalizes the transfer of bitcoins and tokens from send, exchange, and swap actions.

The following breaks down the construction of a Settlement Action. The action is constructed by building a single string from each of the elements in order.

| Field    | Label    | Name         | Example Values | Comments | Data Type          | Restrictions |
|----------|----------|--------------|----------------|----------|--------------------|--------------|
| Metadata | Header[] | Header Array | -              | -        | Common header data | Header       |
| Text Encoding | TextEncoding | 1 | 0 |  0 = ASCII, 1 = UTF-8, 2 = UTF-16, 3 = Unicode.  Encoding applies to all 'text' data types. All 'string' types will always be encoded with ASCII.  Where string is selected, all fields will be ASCII. | uint8 | Can be changed by Issuer or Operator at their discretion. |
| Transfer Type | TransferType | 1 | S | S - Send, E - Exchange, X - Swap | string |  |
| Asset Type 1 | AssetType1 | 3 | RRE | eg. Share, Bond, Ticket | string |  |
| Asset ID 1 | AssetID1 | 32 | apm2qsznhks23z8d83u41s8019hyri3i | Randomly generated base58 string.  Each Asset ID should be unique.  However, a Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type can be the leading bytes - a convention - to make it easy to identify that it is a token by humans.  | string |  |
| Asset Type 2 | AssetType2 | 3 | SHC | eg. Share, Bond, Ticket. NULL for Send and Exchange Response Type. | string |  |
| Asset ID 2 | AssetID2 | 32 | apm2qsznhks23z8d83u41s8019hyri3i | Randomly generated base58 string.  Each Asset ID should be unique.  However, a Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type can be the leading bytes - a convention - to make it easy to identify that it is a token by humans.  NULL for Send and Exchange Response Type. | string |  |
| Asset 1 Settlements Count | Asset1SettlementsCount | 1 | 0 | Number of settlements for Asset 1. | uint8 |  |
| Asset 1 Address X Qty | Asset1AddressesXQty | 0 |  | Each element contains the resulting token balance of Asset 1 for the output Address, which is referred to by the index. | QuantityIndex[] |  |
| Asset 2 Settlements Count | Asset2SettlementsCount | 1 | 0 | Number of settlements for Asset 2. 0 for Send and Exchange Response Type. | uint8 |  |
| Asset 2 Address X Qty | Asset2AddressXQty | 0 |  | Each element contains the resulting token balance of Asset 2 for the output address, which is referred to by the index. Will not be present for Send and Exchange Response Type. | QuantityIndex[] |  |
| Timestamp | Timestamp | 8 | 1551767413250187179 | Timestamp in nanoseconds of when the smart contract created the action. | timestamp | Cannot be changed by issuer, operator. Smart contract controls. |



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
        <td class="t6">Header[]</td>
        <td class="t6">Header Array</td>
        <td class="t6">-</td>
        <td class="t6">-</td>
        <td class="t6">Common header data for all actions</td>
        <td class="t6">Header</td>
        <td class="t7"></td>
    </tr>

    <tr>
        <td class="t10">Text Encoding</td>
        <td class="t10">TextEncoding</td>
        <td class="t10">1</td>
        <td class="t10" style="word-break:break-all">0</td>
        <td class="t10"> 0 = ASCII, 1 = UTF-8, 2 = UTF-16, 3 = Unicode.  Encoding applies to all 'text' data types. All 'string' types will always be encoded with ASCII.  Where string is selected, all fields will be ASCII.</td>
        <td class="t10">uint8</td>
        <td class="t11">Can be changed by Issuer or Operator at their discretion.</td>
    </tr>

    <tr>
        <td class="t10">Transfer Type</td>
        <td class="t10">TransferType</td>
        <td class="t10">1</td>
        <td class="t10" style="word-break:break-all">S</td>
        <td class="t10">S - Send, E - Exchange, X - Swap</td>
        <td class="t10">string</td>
        <td class="t11"></td>
    </tr>

    <tr>
        <td class="t10">Asset Type 1</td>
        <td class="t10">AssetType1</td>
        <td class="t10">3</td>
        <td class="t10" style="word-break:break-all">RRE</td>
        <td class="t10">eg. Share, Bond, Ticket</td>
        <td class="t10">string</td>
        <td class="t11"></td>
    </tr>

    <tr>
        <td class="t10">Asset ID 1</td>
        <td class="t10">AssetID1</td>
        <td class="t10">32</td>
        <td class="t10" style="word-break:break-all">apm2qsznhks23z8d83u41s8019hyri3i</td>
        <td class="t10">Randomly generated base58 string.  Each Asset ID should be unique.  However, a Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type can be the leading bytes - a convention - to make it easy to identify that it is a token by humans. </td>
        <td class="t10">string</td>
        <td class="t11"></td>
    </tr>

    <tr>
        <td class="t10">Asset Type 2</td>
        <td class="t10">AssetType2</td>
        <td class="t10">3</td>
        <td class="t10" style="word-break:break-all">SHC</td>
        <td class="t10">eg. Share, Bond, Ticket. NULL for Send and Exchange Response Type.</td>
        <td class="t10">string</td>
        <td class="t11"></td>
    </tr>

    <tr>
        <td class="t10">Asset ID 2</td>
        <td class="t10">AssetID2</td>
        <td class="t10">32</td>
        <td class="t10" style="word-break:break-all">apm2qsznhks23z8d83u41s8019hyri3i</td>
        <td class="t10">Randomly generated base58 string.  Each Asset ID should be unique.  However, a Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type can be the leading bytes - a convention - to make it easy to identify that it is a token by humans.  NULL for Send and Exchange Response Type.</td>
        <td class="t10">string</td>
        <td class="t11"></td>
    </tr>

    <tr>
        <td class="t10">Asset 1 Settlements Count</td>
        <td class="t10">Asset1SettlementsCount</td>
        <td class="t10">1</td>
        <td class="t10" style="word-break:break-all">0</td>
        <td class="t10">Number of settlements for Asset 1.</td>
        <td class="t10">uint8</td>
        <td class="t11"></td>
    </tr>

    <tr>
        <td class="t10">Asset 1 Address X Qty</td>
        <td class="t10">Asset1AddressesXQty</td>
        <td class="t10">0</td>
        <td class="t10" style="word-break:break-all"></td>
        <td class="t10">Each element contains the resulting token balance of Asset 1 for the output Address, which is referred to by the index.</td>
        <td class="t10">QuantityIndex[]</td>
        <td class="t11"></td>
    </tr>

    <tr>
        <td class="t10">Asset 2 Settlements Count</td>
        <td class="t10">Asset2SettlementsCount</td>
        <td class="t10">1</td>
        <td class="t10" style="word-break:break-all">0</td>
        <td class="t10">Number of settlements for Asset 2. 0 for Send and Exchange Response Type.</td>
        <td class="t10">uint8</td>
        <td class="t11"></td>
    </tr>

    <tr>
        <td class="t10">Asset 2 Address X Qty</td>
        <td class="t10">Asset2AddressXQty</td>
        <td class="t10">0</td>
        <td class="t10" style="word-break:break-all"></td>
        <td class="t10">Each element contains the resulting token balance of Asset 2 for the output address, which is referred to by the index. Will not be present for Send and Exchange Response Type.</td>
        <td class="t10">QuantityIndex[]</td>
        <td class="t11"></td>
    </tr>

    <tr>
        <td class="t10">Timestamp</td>
        <td class="t10">Timestamp</td>
        <td class="t10">8</td>
        <td class="t10" style="word-break:break-all">1551767413250187179</td>
        <td class="t10">Timestamp in nanoseconds of when the smart contract created the action.</td>
        <td class="t10">timestamp</td>
        <td class="t11">Cannot be changed by issuer, operator. Smart contract controls.</td>
    </tr>

</table>
!-->