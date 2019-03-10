
# Message Action

Message Action - the message action is a general purpose communication action. 'Twitter/SMS' for Issuers/Investors/Users. The message txn can also be used for passing partially signed txns on-chain, establishing private communication channels and EDI (receipting, invoices, PO, and private offers/bids). The messages are broken down by type for easy filtering in the a user's wallet. The Message Types are listed in the Message Types table.


The following breaks down the construction of a Message Action. The action is constructed by building a single string from each of the elements in order.

| Field    | Label    | Name         | Example Values | Comments | Data Type          | Restrictions |
|----------|----------|--------------|----------------|----------|--------------------|--------------|
| Metadata | Header[] | Header Array | -              | -        | Common header data | Header       |
| Text Encoding | TextEncoding | 1 | 0 | 0 = ASCII, 1 = UTF-8, 2 = UTF-16, 3 = Unicode. Encoding applies to all 'text' data types. All 'string' types will always be encoded with ASCII. Where string is selected, all fields will be ASCII.
 | uint8 | Can be changed by Issuer or Operator at their discretion. |
| Qty Receiving Addresses | QtyReceivingAddresses | 1 | 2 | 0-255 Message Receiving Addresses | uint8 |  |
| Address Indexes | AddressIndexes | 0 |  | Associates the message to a particular output by the index. | uint16[] |  |
| Message Type | MessageType | 2 | 6 | Potential for up to 65,535 different message types | string |  |
| Message Payload | MessagePayload | 0 | Hello world! | Length only limited by the Bitcoin protocol. Public or private (RSA public key, Diffie-Hellman). Issuers/Contracts can send the signifying amount of satoshis to themselves for public announcements or private 'notes' if encrypted. See Message Types for a full list of potential use cases.
 | nvarchar64 |  |



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
        <td class="m10">0 = ASCII, 1 = UTF-8, 2 = UTF-16, 3 = Unicode. Encoding applies to all 'text' data types. All 'string' types will always be encoded with ASCII. Where string is selected, all fields will be ASCII.
</td>
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
        <td class="m10">Message Type</td>
        <td class="m10">MessageType</td>
        <td class="m10">2</td>
        <td class="m10" style="word-break:break-all">6</td>
        <td class="m10">Potential for up to 65,535 different message types</td>
        <td class="m10">string</td>
        <td class="m11"></td>
    </tr>

    <tr>
        <td class="m10">Message Payload</td>
        <td class="m10">MessagePayload</td>
        <td class="m10">0</td>
        <td class="m10" style="word-break:break-all">Hello world!</td>
        <td class="m10">Length only limited by the Bitcoin protocol. Public or private (RSA public key, Diffie-Hellman). Issuers/Contracts can send the signifying amount of satoshis to themselves for public announcements or private 'notes' if encrypted. See Message Types for a full list of potential use cases.
</td>
        <td class="m10">nvarchar64</td>
        <td class="m11"></td>
    </tr>

</table>
!-->