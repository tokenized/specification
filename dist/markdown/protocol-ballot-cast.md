
# Ballot Cast Action

Ballot Cast Action -  Used by Token Owners to cast their ballot (vote) on proposals raised by the Issuer (Referendum) or other token holders (Initiative). 1 Vote per token unless a vote multiplier is specified in the relevant Asset Definition action.

The following breaks down the construction of a Ballot Cast Action. The action is constructed by building a single string from each of the elements in order.

| Field    | Label    | Name         | Example Values | Comments | Data Type          | Restrictions |
|----------|----------|--------------|----------------|----------|--------------------|--------------|
| Metadata | Header[] | Header Array | -              | -        | Common header data | Header       |
| Asset Type | AssetType | 3 | RRE | eg. Share, Bond, Ticket | string |  |
| Asset ID | AssetID | 32 | apm2qsznhks23z8d83u41s8019hyri3i | Randomly generated base58 string.  Each Asset ID should be unique.  However, a Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type can be the leading bytes - a convention - to make it easy to identify that it is a token by humans. | string |  |
| Vote Txn ID | VoteTxnID | 32 | f3318be9fb3f73e53b29868beae46b42911c2116f979a5d3284face90746cb37 | Tx-ID of the Vote the Ballot Cast is being made for. | sha256 |  |
| Vote | Vote | 0 | A | Length 1-255 bytes. 0 is not valid. Accept, Reject, Abstain, Spoiled, Multiple Choice, or Preference List. 15 options total. Order of preference.  1st position = 1st choice. 2nd position = 2nd choice, etc.  A is always Accept and B is always reject in a Y/N votes. | nvarchar8 |  |



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
        <td class="g6">Header[]</td>
        <td class="g6">Header Array</td>
        <td class="g6">-</td>
        <td class="g6">-</td>
        <td class="g6">Common header data for all actions</td>
        <td class="g6">Header</td>
        <td class="g7"></td>
    </tr>

    <tr>
        <td class="g10">Asset Type</td>
        <td class="g10">AssetType</td>
        <td class="g10">3</td>
        <td class="g10" style="word-break:break-all">RRE</td>
        <td class="g10">eg. Share, Bond, Ticket</td>
        <td class="g10">string</td>
        <td class="g11"></td>
    </tr>

    <tr>
        <td class="g10">Asset ID</td>
        <td class="g10">AssetID</td>
        <td class="g10">32</td>
        <td class="g10" style="word-break:break-all">apm2qsznhks23z8d83u41s8019hyri3i</td>
        <td class="g10">Randomly generated base58 string.  Each Asset ID should be unique.  However, a Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type can be the leading bytes - a convention - to make it easy to identify that it is a token by humans.</td>
        <td class="g10">string</td>
        <td class="g11"></td>
    </tr>

    <tr>
        <td class="g10">Vote Txn ID</td>
        <td class="g10">VoteTxnID</td>
        <td class="g10">32</td>
        <td class="g10" style="word-break:break-all">f3318be9fb3f73e53b29868beae46b42911c2116f979a5d3284face90746cb37</td>
        <td class="g10">Tx-ID of the Vote the Ballot Cast is being made for.</td>
        <td class="g10">sha256</td>
        <td class="g11"></td>
    </tr>

    <tr>
        <td class="g10">Vote</td>
        <td class="g10">Vote</td>
        <td class="g10">0</td>
        <td class="g10" style="word-break:break-all">A</td>
        <td class="g10">Length 1-255 bytes. 0 is not valid. Accept, Reject, Abstain, Spoiled, Multiple Choice, or Preference List. 15 options total. Order of preference.  1st position = 1st choice. 2nd position = 2nd choice, etc.  A is always Accept and B is always reject in a Y/N votes.</td>
        <td class="g10">nvarchar8</td>
        <td class="g11"></td>
    </tr>

</table>
!-->