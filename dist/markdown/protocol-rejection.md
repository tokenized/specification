
# Rejection Action

Rejection Action - used to reject request actions that do not comply with the Contract. If money is to be returned to a User then it is used in lieu of the Settlement Action to properly account for token balances. All Issuer/User request Actions must be responded to by the Contract with an Action.  The only exception to this rule is when there is not enough fees in the first Action for the Contract response action to remain revenue neutral.  If not enough fees are attached to pay for the Contract response then the Contract will not respond.

The following breaks down the construction of a Rejection Action. The action is constructed by building a single string from each of the elements in order.

| Field    | Label    | Name         | Example Values | Comments | Data Type          | Restrictions |
|----------|----------|--------------|----------------|----------|--------------------|--------------|
| Metadata | Header[] | Header Array | -              | -        | Common header data | Header       |
| Text Encoding | TextEncoding | 1 | 0 |  0 = ASCII, 1 = UTF-8, 2 = UTF-16, 3 = Unicode.  Encoding applies to all 'text' data types. All 'string' types will always be encoded with ASCII.  Where string is selected, all fields will be ASCII. | uint8 | Can be changed by Issuer or Operator at their discretion. |
| Qty Receiving Addresses | QtyReceivingAddresses | 1 | 2 | 0-255 Message Receiving Addresses | uint8 |  |
| Address Indexes | AddressIndexes | 0 |  | Associates the message to a particular output by the index. | uint16[] |  |
| Rejection Type | RejectionType | 1 | 1 | Classifies the rejection by a type. | uint8 |  |
| Message Payload | MessagePayload | 0 | Sorry, you don't have enough tokens. | Length 0-65,535 bytes. Message that explains the reasoning for a rejection, if needed.  Most rejection types will be captured by the Rejection Type Subfield. | nvarchar16 |  |
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
        <td class="m6">Header[]</td>
        <td class="m6">Header Array</td>
        <td class="m6">-</td>
        <td class="m6">-</td>
        <td class="m6">Common header data for all actions</td>
        <td class="m6">Header</td>
        <td class="m7"></td>
    </tr>

    <tr>
        <td class="m10">Text Encoding</td>
        <td class="m10">TextEncoding</td>
        <td class="m10">1</td>
        <td class="m10" style="word-break:break-all">0</td>
        <td class="m10"> 0 = ASCII, 1 = UTF-8, 2 = UTF-16, 3 = Unicode.  Encoding applies to all 'text' data types. All 'string' types will always be encoded with ASCII.  Where string is selected, all fields will be ASCII.</td>
        <td class="m10">uint8</td>
        <td class="m11">Can be changed by Issuer or Operator at their discretion.</td>
    </tr>

    <tr>
        <td class="m10">Qty Receiving Addresses</td>
        <td class="m10">QtyReceivingAddresses</td>
        <td class="m10">1</td>
        <td class="m10" style="word-break:break-all">2</td>
        <td class="m10">0-255 Message Receiving Addresses</td>
        <td class="m10">uint8</td>
        <td class="m11"></td>
    </tr>

    <tr>
        <td class="m10">Address Indexes</td>
        <td class="m10">AddressIndexes</td>
        <td class="m10">0</td>
        <td class="m10" style="word-break:break-all"></td>
        <td class="m10">Associates the message to a particular output by the index.</td>
        <td class="m10">uint16[]</td>
        <td class="m11"></td>
    </tr>

    <tr>
        <td class="m10">Rejection Type</td>
        <td class="m10">RejectionType</td>
        <td class="m10">1</td>
        <td class="m10" style="word-break:break-all">1</td>
        <td class="m10">Classifies the rejection by a type.</td>
        <td class="m10">uint8</td>
        <td class="m11"></td>
    </tr>

    <tr>
        <td class="m10">Message Payload</td>
        <td class="m10">MessagePayload</td>
        <td class="m10">0</td>
        <td class="m10" style="word-break:break-all">Sorry, you don't have enough tokens.</td>
        <td class="m10">Length 0-65,535 bytes. Message that explains the reasoning for a rejection, if needed.  Most rejection types will be captured by the Rejection Type Subfield.</td>
        <td class="m10">nvarchar16</td>
        <td class="m11"></td>
    </tr>

    <tr>
        <td class="m10">Timestamp</td>
        <td class="m10">Timestamp</td>
        <td class="m10">8</td>
        <td class="m10" style="word-break:break-all">1551767413250187179</td>
        <td class="m10">Timestamp in nanoseconds of when the smart contract created the action.</td>
        <td class="m10">timestamp</td>
        <td class="m11">Cannot be changed by issuer, operator. Smart contract controls.</td>
    </tr>

</table>
!-->