
# Asset Definition Action

Asset Definition Action -  This action is used by the issuer to define the properties/characteristics of the Asset (token) that it wants to create.

The following breaks down the construction of a Asset Definition Action. The action is constructed by building a single string from each of the elements in order.

| Field    | Label    | Name         | Example Values | Comments | Data Type          | Restrictions |
|----------|----------|--------------|----------------|----------|--------------------|--------------|
| Metadata | Header[] | Header Array | -              | -        | Common header data | Header       |
| Text Encoding | TextEncoding | 1 | 0 |  0 = ASCII, 1 = UTF-8, 2 = UTF-16, 3 = Unicode.  Encoding applies to all 'text' data types. All 'string' types will always be encoded with ASCII.  Where string is selected, all fields will be ASCII. | uint8 | Can be changed by issuer or operator at their discretion. |
| Asset Type | AssetType | 3 | COU | eg. Share - Common | string | Cannot be changed by issuer, operator or smart contract. |
| Asset ID | AssetID | 32 | apm2qsznhks23z8d83u41s8019hyri3i | Randomly generated base58 string.  Each Asset ID should be unique.  However, an Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type + Asset ID = Asset Code.  An Asset Code is a human readable idenitfier that can be used in a similar way to a Bitcoin (BSV) address, a vanity identifying label. | string | Cannot be changed by issuer, operator or smart contract. |
| Asset Authorization Flags | AssetAuthFlags | 8 | 0000000000000000000000000000000000000000000000000001000110111111 | Authorization Flags,  bitwise operation | bin |  |
| Transfers Permitted | TransfersPermitted | 1 | 1 | 1 = Transfers are permitted.  0 = Transfers are not permitted. | bool |  |
| Trade Restrictions | TradeRestrictions | 3 | GBR | Asset can only be traded within the trade restrictions.  Eg. AUS - Australian residents only.  EU - European Union residents only. | string |  |
| Enforcement Orders Permitted | EnforcementOrdersPermitted | 1 | 1 | 1 = Enforcement Orders are permitted. 0 = Enforcement Orders are not permitted. | bool |  |
| Vote Multiplier | VoteMultiplier | 1 | 3 | Multiplies the vote by the integer. 1 token = 1 vote with a 1 for vote multipler (normal).  1 token = 3 votes with a multiplier of 3, for example. | uint8 |  |
| Referendum Proposal | ReferendumProposal | 1 | 1 | A Referendum is permitted for Asset-Wide Proposals (outside of smart contract scope) if also permitted by the contract. If the contract has proposals by referendum restricted, then this flag is meaningless. | bool |  |
| Initiative Proposal | InitiativeProposal | 1 | 1 | An initiative is permitted for Asset-Wide Proposals (outside of smart contract scope) if also permitted by the contract. If the contract has proposals by initiative restricted, then this flag is meaningless. | bool |  |
| Asset Modification Governance | AssetModificationGovernance | 1 | 1 | 1 - Contract-wide Asset Governance.  0 - Asset-wide Asset Governance.  If a referendum or initiative is used to propose a modification to a subfield controlled by the asset auth flags, then the vote will either be a contract-wide vote (all assets vote on the referendum/initiative) or an asset-wide vote (all assets vote on the referendum/initiative) depending on the value in this subfield.  The voting system specifies the voting rules. | bool |  |
| Qty of Tokens | TokenQty | 8 | 1000000 | Quantity of token - 0 is valid. Fungible 'shares' of the Asset. 1 is used for non-fungible tokens.  Asset IDs become the non-fungible Asset ID and many Asset IDs can be associated with a particular Contract. | uint64 |  |
| Contract Fee Currency | ContractFeeCurrency | 3 | AUD | BSV, USD, AUD, EUR, etc. | string |  |
| Contract Fee Var | ContractFeeVar | 4 | 0.005 | Percent of the value of the transaction | float32 |  |
| Contract Fee Fixed | ContractFeeFixed | 4 | 0.01 | Fixed fee (payment made in BSV) | float32 |  |
| Asset Payload Length | AssetPayloadLen | 2 | 9 | Size of the asset payload in bytes. | uint16 |  |
| Asset Payload | AssetPayload | 0 | some data | Payload length is dependent on the asset type. Each asset is made up of a defined set of information pertaining to the specific asset type, and may contain fields of variable length type (nvarchar8, 16, 32) | byte[] |  |



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
        <td class="a10"> 0 = ASCII, 1 = UTF-8, 2 = UTF-16, 3 = Unicode.  Encoding applies to all 'text' data types. All 'string' types will always be encoded with ASCII.  Where string is selected, all fields will be ASCII.</td>
        <td class="a10">uint8</td>
        <td class="a11">Can be changed by issuer or operator at their discretion.</td>
    </tr>

    <tr>
        <td class="a10">Asset Type</td>
        <td class="a10">AssetType</td>
        <td class="a10">3</td>
        <td class="a10" style="word-break:break-all">COU</td>
        <td class="a10">eg. Share - Common</td>
        <td class="a10">string</td>
        <td class="a11">Cannot be changed by issuer, operator or smart contract.</td>
    </tr>

    <tr>
        <td class="a10">Asset ID</td>
        <td class="a10">AssetID</td>
        <td class="a10">32</td>
        <td class="a10" style="word-break:break-all">apm2qsznhks23z8d83u41s8019hyri3i</td>
        <td class="a10">Randomly generated base58 string.  Each Asset ID should be unique.  However, an Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type + Asset ID = Asset Code.  An Asset Code is a human readable idenitfier that can be used in a similar way to a Bitcoin (BSV) address, a vanity identifying label.</td>
        <td class="a10">string</td>
        <td class="a11">Cannot be changed by issuer, operator or smart contract.</td>
    </tr>

    <tr>
        <td class="a10">Asset Authorization Flags</td>
        <td class="a10">AssetAuthFlags</td>
        <td class="a10">8</td>
        <td class="a10" style="word-break:break-all">0000000000000000000000000000000000000000000000000001000110111111</td>
        <td class="a10">Authorization Flags,  bitwise operation</td>
        <td class="a10">bin</td>
        <td class="a11"></td>
    </tr>

    <tr>
        <td class="a10">Transfers Permitted</td>
        <td class="a10">TransfersPermitted</td>
        <td class="a10">1</td>
        <td class="a10" style="word-break:break-all">1</td>
        <td class="a10">1 = Transfers are permitted.  0 = Transfers are not permitted.</td>
        <td class="a10">bool</td>
        <td class="a11"></td>
    </tr>

    <tr>
        <td class="a10">Trade Restrictions</td>
        <td class="a10">TradeRestrictions</td>
        <td class="a10">3</td>
        <td class="a10" style="word-break:break-all">GBR</td>
        <td class="a10">Asset can only be traded within the trade restrictions.  Eg. AUS - Australian residents only.  EU - European Union residents only.</td>
        <td class="a10">string</td>
        <td class="a11"></td>
    </tr>

    <tr>
        <td class="a10">Enforcement Orders Permitted</td>
        <td class="a10">EnforcementOrdersPermitted</td>
        <td class="a10">1</td>
        <td class="a10" style="word-break:break-all">1</td>
        <td class="a10">1 = Enforcement Orders are permitted. 0 = Enforcement Orders are not permitted.</td>
        <td class="a10">bool</td>
        <td class="a11"></td>
    </tr>

    <tr>
        <td class="a10">Vote Multiplier</td>
        <td class="a10">VoteMultiplier</td>
        <td class="a10">1</td>
        <td class="a10" style="word-break:break-all">3</td>
        <td class="a10">Multiplies the vote by the integer. 1 token = 1 vote with a 1 for vote multipler (normal).  1 token = 3 votes with a multiplier of 3, for example.</td>
        <td class="a10">uint8</td>
        <td class="a11"></td>
    </tr>

    <tr>
        <td class="a10">Referendum Proposal</td>
        <td class="a10">ReferendumProposal</td>
        <td class="a10">1</td>
        <td class="a10" style="word-break:break-all">1</td>
        <td class="a10">A Referendum is permitted for Asset-Wide Proposals (outside of smart contract scope) if also permitted by the contract. If the contract has proposals by referendum restricted, then this flag is meaningless.</td>
        <td class="a10">bool</td>
        <td class="a11"></td>
    </tr>

    <tr>
        <td class="a10">Initiative Proposal</td>
        <td class="a10">InitiativeProposal</td>
        <td class="a10">1</td>
        <td class="a10" style="word-break:break-all">1</td>
        <td class="a10">An initiative is permitted for Asset-Wide Proposals (outside of smart contract scope) if also permitted by the contract. If the contract has proposals by initiative restricted, then this flag is meaningless.</td>
        <td class="a10">bool</td>
        <td class="a11"></td>
    </tr>

    <tr>
        <td class="a10">Asset Modification Governance</td>
        <td class="a10">AssetModificationGovernance</td>
        <td class="a10">1</td>
        <td class="a10" style="word-break:break-all">1</td>
        <td class="a10">1 - Contract-wide Asset Governance.  0 - Asset-wide Asset Governance.  If a referendum or initiative is used to propose a modification to a subfield controlled by the asset auth flags, then the vote will either be a contract-wide vote (all assets vote on the referendum/initiative) or an asset-wide vote (all assets vote on the referendum/initiative) depending on the value in this subfield.  The voting system specifies the voting rules.</td>
        <td class="a10">bool</td>
        <td class="a11"></td>
    </tr>

    <tr>
        <td class="a10">Qty of Tokens</td>
        <td class="a10">TokenQty</td>
        <td class="a10">8</td>
        <td class="a10" style="word-break:break-all">1000000</td>
        <td class="a10">Quantity of token - 0 is valid. Fungible 'shares' of the Asset. 1 is used for non-fungible tokens.  Asset IDs become the non-fungible Asset ID and many Asset IDs can be associated with a particular Contract.</td>
        <td class="a10">uint64</td>
        <td class="a11"></td>
    </tr>

    <tr>
        <td class="a10">Contract Fee Currency</td>
        <td class="a10">ContractFeeCurrency</td>
        <td class="a10">3</td>
        <td class="a10" style="word-break:break-all">AUD</td>
        <td class="a10">BSV, USD, AUD, EUR, etc.</td>
        <td class="a10">string</td>
        <td class="a11"></td>
    </tr>

    <tr>
        <td class="a10">Contract Fee Var</td>
        <td class="a10">ContractFeeVar</td>
        <td class="a10">4</td>
        <td class="a10" style="word-break:break-all">0.005</td>
        <td class="a10">Percent of the value of the transaction</td>
        <td class="a10">float32</td>
        <td class="a11"></td>
    </tr>

    <tr>
        <td class="a10">Contract Fee Fixed</td>
        <td class="a10">ContractFeeFixed</td>
        <td class="a10">4</td>
        <td class="a10" style="word-break:break-all">0.01</td>
        <td class="a10">Fixed fee (payment made in BSV)</td>
        <td class="a10">float32</td>
        <td class="a11"></td>
    </tr>

    <tr>
        <td class="a10">Asset Payload Length</td>
        <td class="a10">AssetPayloadLen</td>
        <td class="a10">2</td>
        <td class="a10" style="word-break:break-all">9</td>
        <td class="a10">Size of the asset payload in bytes.</td>
        <td class="a10">uint16</td>
        <td class="a11"></td>
    </tr>

    <tr>
        <td class="a10">Asset Payload</td>
        <td class="a10">AssetPayload</td>
        <td class="a10">0</td>
        <td class="a10" style="word-break:break-all">some data</td>
        <td class="a10">Payload length is dependent on the asset type. Each asset is made up of a defined set of information pertaining to the specific asset type, and may contain fields of variable length type (nvarchar8, 16, 32)</td>
        <td class="a10">byte[]</td>
        <td class="a11"></td>
    </tr>

</table>
!-->