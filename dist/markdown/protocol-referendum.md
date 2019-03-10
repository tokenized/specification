
# Referendum Action

Referendum Action -  Issuer instructs the Contract to Initiate a Token Owner Vote. Usually used for contract amendments, organizational governance, etc.

The following breaks down the construction of a Referendum Action. The action is constructed by building a single string from each of the elements in order.

| Field    | Label    | Name         | Example Values | Comments | Data Type          | Restrictions |
|----------|----------|--------------|----------------|----------|--------------------|--------------|
| Metadata | Header[] | Header Array | -              | -        | Common header data | Header       |
| Text Encoding | TextEncoding | 1 | 0 |  0 = ASCII, 1 = UTF-8, 2 = UTF-16, 3 = Unicode.  Encoding applies to all 'text' data types. All 'string' types will always be encoded with ASCII.  Where string is selected, all fields will be ASCII. | uint8 | Can be changed by Issuer or Operator at their discretion. |
| Asset Specific Vote | AssetSpecificVote | 1 | 1 | 1 - Yes, 0 - No.  No Asset Type/AssetID subfields for N - No. | bool |  |
| Asset Type | AssetType | 3 | RRE | eg. Share, Bond, Ticket | string |  |
| Asset ID | AssetID | 32 | apm2qsznhks23z8d83u41s8019hyri3i | Randomly generated base58 string.  Each Asset ID should be unique.  However, an Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type can be the leading bytes - a convention - to make it easy to identify that it is a token by humans. | string |  |
| Vote System | VoteSystem | 1 | 1 | X for Vote System X. (1-255, 0 is not valid.) | uint8 |  |
| Proposal | Proposal | 1 | 0 | 1 for a Proposal, 0 for an initiative that is requesting changes to specific subfields for modification. If this field is true, the subfields should be empty.  The smart contract cannot interpret the results of a vote when Proposal = 1.  All meaning is interpreted by the token owners and smart contract simply facilates the record keeping.  When Proposal = 0, the smart contract always assumes the first choice is a 'yes', or 'pass', if the threshold is met, and will process the proposed changes accordingly. | bool |  |
|  | ProposedChangesCount | 1 | 0 |  | uint8 |  |
| Proposed Changes | ProposedChanges | 0 |  | Each element contains details of which fields to modify, or delete. Because the number of fields in a Contract and Asset is dynamic due to some fields being able to be repeated, the index value of the field needs to be calculated against the Contract or Asset the changes are to apply to. In the event of a Vote being created from this Initiative, the changes will be applied to the version of the Contract or Asset at that time. | Amendment[] |  |
| Vote Options | VoteOptions | 0 | ABCDEFGHIJKLMNO | Length 1-255 bytes. 0 is not valid. Each byte allows for a different vote option.  Typical votes will likely be multiple choice or Y/N. Vote instances are identified by the Tx-ID. AB000000000 would be chosen for Y/N (binary) type votes. Only applicable if Proposal Type is set to P for Proposal.  All other Proposal Types will be binary.  Pass/Fail. | nvarchar8 |  |
| Vote Max | VoteMax | 1 | 15 | Range: 1-15. How many selections can a voter make in a Ballot Cast.  1 is selected for Y/N (binary) | uint8 |  |
| Proposal Description | ProposalDescription | 0 | Change the name of the Contract. | Length 0 to 65,535 bytes. 0 is valid. Description of the vote. | nvarchar16 |  |
| Proposal Document Hash | ProposalDocumentHash | 32 | 77201b0094f50df309f0343e4f44dae64d0de503c91038faf2c6b039f9f18aec | Hash of the proposal document to be distributed to voters | sha256 |  |
| Vote Cut-Off Timestamp | VoteCutOffTimestamp | 8 | 10/07/2018 00:00:00 | Ballot casts after this timestamp will not be included. The vote has finished. | time |  |



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
        <td class="g10">Text Encoding</td>
        <td class="g10">TextEncoding</td>
        <td class="g10">1</td>
        <td class="g10" style="word-break:break-all">0</td>
        <td class="g10"> 0 = ASCII, 1 = UTF-8, 2 = UTF-16, 3 = Unicode.  Encoding applies to all 'text' data types. All 'string' types will always be encoded with ASCII.  Where string is selected, all fields will be ASCII.</td>
        <td class="g10">uint8</td>
        <td class="g11">Can be changed by Issuer or Operator at their discretion.</td>
    </tr>

    <tr>
        <td class="g10">Asset Specific Vote</td>
        <td class="g10">AssetSpecificVote</td>
        <td class="g10">1</td>
        <td class="g10" style="word-break:break-all">1</td>
        <td class="g10">1 - Yes, 0 - No.  No Asset Type/AssetID subfields for N - No.</td>
        <td class="g10">bool</td>
        <td class="g11"></td>
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
        <td class="g10">Randomly generated base58 string.  Each Asset ID should be unique.  However, an Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type can be the leading bytes - a convention - to make it easy to identify that it is a token by humans.</td>
        <td class="g10">string</td>
        <td class="g11"></td>
    </tr>

    <tr>
        <td class="g10">Vote System</td>
        <td class="g10">VoteSystem</td>
        <td class="g10">1</td>
        <td class="g10" style="word-break:break-all">1</td>
        <td class="g10">X for Vote System X. (1-255, 0 is not valid.)</td>
        <td class="g10">uint8</td>
        <td class="g11"></td>
    </tr>

    <tr>
        <td class="g10">Proposal</td>
        <td class="g10">Proposal</td>
        <td class="g10">1</td>
        <td class="g10" style="word-break:break-all">0</td>
        <td class="g10">1 for a Proposal, 0 for an initiative that is requesting changes to specific subfields for modification. If this field is true, the subfields should be empty.  The smart contract cannot interpret the results of a vote when Proposal = 1.  All meaning is interpreted by the token owners and smart contract simply facilates the record keeping.  When Proposal = 0, the smart contract always assumes the first choice is a 'yes', or 'pass', if the threshold is met, and will process the proposed changes accordingly.</td>
        <td class="g10">bool</td>
        <td class="g11"></td>
    </tr>

    <tr>
        <td class="g10"></td>
        <td class="g10">ProposedChangesCount</td>
        <td class="g10">1</td>
        <td class="g10" style="word-break:break-all">0</td>
        <td class="g10"></td>
        <td class="g10">uint8</td>
        <td class="g11"></td>
    </tr>

    <tr>
        <td class="g10">Proposed Changes</td>
        <td class="g10">ProposedChanges</td>
        <td class="g10">0</td>
        <td class="g10" style="word-break:break-all"></td>
        <td class="g10">Each element contains details of which fields to modify, or delete. Because the number of fields in a Contract and Asset is dynamic due to some fields being able to be repeated, the index value of the field needs to be calculated against the Contract or Asset the changes are to apply to. In the event of a Vote being created from this Initiative, the changes will be applied to the version of the Contract or Asset at that time.</td>
        <td class="g10">Amendment[]</td>
        <td class="g11"></td>
    </tr>

    <tr>
        <td class="g10">Vote Options</td>
        <td class="g10">VoteOptions</td>
        <td class="g10">0</td>
        <td class="g10" style="word-break:break-all">ABCDEFGHIJKLMNO</td>
        <td class="g10">Length 1-255 bytes. 0 is not valid. Each byte allows for a different vote option.  Typical votes will likely be multiple choice or Y/N. Vote instances are identified by the Tx-ID. AB000000000 would be chosen for Y/N (binary) type votes. Only applicable if Proposal Type is set to P for Proposal.  All other Proposal Types will be binary.  Pass/Fail.</td>
        <td class="g10">nvarchar8</td>
        <td class="g11"></td>
    </tr>

    <tr>
        <td class="g10">Vote Max</td>
        <td class="g10">VoteMax</td>
        <td class="g10">1</td>
        <td class="g10" style="word-break:break-all">15</td>
        <td class="g10">Range: 1-15. How many selections can a voter make in a Ballot Cast.  1 is selected for Y/N (binary)</td>
        <td class="g10">uint8</td>
        <td class="g11"></td>
    </tr>

    <tr>
        <td class="g10">Proposal Description</td>
        <td class="g10">ProposalDescription</td>
        <td class="g10">0</td>
        <td class="g10" style="word-break:break-all">Change the name of the Contract.</td>
        <td class="g10">Length 0 to 65,535 bytes. 0 is valid. Description of the vote.</td>
        <td class="g10">nvarchar16</td>
        <td class="g11"></td>
    </tr>

    <tr>
        <td class="g10">Proposal Document Hash</td>
        <td class="g10">ProposalDocumentHash</td>
        <td class="g10">32</td>
        <td class="g10" style="word-break:break-all">77201b0094f50df309f0343e4f44dae64d0de503c91038faf2c6b039f9f18aec</td>
        <td class="g10">Hash of the proposal document to be distributed to voters</td>
        <td class="g10">sha256</td>
        <td class="g11"></td>
    </tr>

    <tr>
        <td class="g10">Vote Cut-Off Timestamp</td>
        <td class="g10">VoteCutOffTimestamp</td>
        <td class="g10">8</td>
        <td class="g10" style="word-break:break-all">10/07/2018 00:00:00</td>
        <td class="g10">Ballot casts after this timestamp will not be included. The vote has finished.</td>
        <td class="g10">time</td>
        <td class="g11"></td>
    </tr>

</table>
!-->