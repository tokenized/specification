
# Asset Creation Action

Asset Creation Action -  This action creates an Asset in response to the Issuer's instructions in the Definition Action.

The following breaks down the construction of a Asset Creation Action. The action is constructed by building a single string from each of the elements in order.

<div class="ritz grid-container" dir="ltr">
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
            <td class="s5" rowspan="18">Metadata (OP_RETURN Payload)</td>
            <td class="a6" colspan="7">
                <a href="javascript:;" data-popover="type-Header">
                   Header - Click to show content
                </a>
             </td>
        </tr>

        <tr>
            <td class="a10">Asset Type</td>
            <td class="a10">AssetType</td>
            <td class="a10">3</td>
            <td class="a10" style="word-break:break-all">
                SHC
            </td>
            <td class="a10">eg. Share - Common</td>
            <td class="a10">fixedchar</td>
            <td class="a11"></td>
        </tr>

        <tr>
            <td class="a10">Asset ID</td>
            <td class="a10">AssetID</td>
            <td class="a10">32</td>
            <td class="a10" style="word-break:break-all">
                apm2qsznhks23z8d83u41s8019hyri3i
            </td>
            <td class="a10">Randomly generated base58 string.  Each Asset ID should be unique.  However, a Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type + Asset ID = Asset Code.  An Asset Code is a human readable idenitfier that can be used in a similar way to a Bitcoin (BSV) address, a vanity identifying label.</td>
            <td class="a10">fixedchar</td>
            <td class="a11"></td>
        </tr>

        <tr>
            <td class="a10">Asset Authorization Flags</td>
            <td class="a10">AssetAuthFlags</td>
            <td class="a10">8</td>
            <td class="a10" style="word-break:break-all">
                0000000000000000000000000000000000000000000000000001000110111111
            </td>
            <td class="a10">Authorization Flags,  bitwise operation</td>
            <td class="a10">bin</td>
            <td class="a11"></td>
        </tr>

        <tr>
            <td class="a10">Transfers Permitted</td>
            <td class="a10">TransfersPermitted</td>
            <td class="a10">1</td>
            <td class="a10" style="word-break:break-all">
                1
            </td>
            <td class="a10">1 = Transfers are permitted.  0 = Transfers are not permitted.</td>
            <td class="a10">bool</td>
            <td class="a11"></td>
        </tr>

        <tr>
            <td class="a10">Trade Restrictions</td>
            <td class="a10">TradeRestrictions</td>
            <td class="a10">3</td>
            <td class="a10" style="word-break:break-all">
                GBR
            </td>
            <td class="a10">Asset can only be traded within the trade restrictions.  Eg. AUS - Australian residents only.  EU - European Union residents only.</td>
            <td class="a10">fixedchar</td>
            <td class="a11"></td>
        </tr>

        <tr>
            <td class="a10">Enforcement Orders Permitted</td>
            <td class="a10">EnforcementOrdersPermitted</td>
            <td class="a10">1</td>
            <td class="a10" style="word-break:break-all">
                1
            </td>
            <td class="a10">1 = Enforcement Orders are permitted. 0 = Enforcement Orders are not permitted.</td>
            <td class="a10">bool</td>
            <td class="a11"></td>
        </tr>

        <tr>
            <td class="a10">Vote Multiplier</td>
            <td class="a10">VoteMultiplier</td>
            <td class="a10">1</td>
            <td class="a10" style="word-break:break-all">
                3
            </td>
            <td class="a10">Multiplies the vote by the integer. 1 token = 1 vote with a 1 for vote multipler (normal).  1 token = 3 votes with a multiplier of 3, for example.</td>
            <td class="a10">uint</td>
            <td class="a11"></td>
        </tr>

        <tr>
            <td class="a10">Referendum Proposal</td>
            <td class="a10">ReferendumProposal</td>
            <td class="a10">1</td>
            <td class="a10" style="word-break:break-all">
                1
            </td>
            <td class="a10">A Referendum is permitted for Asset-Wide Proposals (outside of smart contract scope) if also permitted by the contract. If the contract has proposals by referendum restricted, then this flag is meaningless.</td>
            <td class="a10">bool</td>
            <td class="a11"></td>
        </tr>

        <tr>
            <td class="a10">Initiative Proposal</td>
            <td class="a10">InitiativeProposal</td>
            <td class="a10">1</td>
            <td class="a10" style="word-break:break-all">
                1
            </td>
            <td class="a10">An initiative is permitted for Asset-Wide Proposals (outside of smart contract scope) if also permitted by the contract. If the contract has proposals by initiative restricted, then this flag is meaningless.</td>
            <td class="a10">bool</td>
            <td class="a11"></td>
        </tr>

        <tr>
            <td class="a10">Asset Modification Governance</td>
            <td class="a10">AssetModificationGovernance</td>
            <td class="a10">1</td>
            <td class="a10" style="word-break:break-all">
                1
            </td>
            <td class="a10">1 - Contract-wide Asset Governance.  0 - Asset-wide Asset Governance.  If a referendum or initiative is used to propose a modification to a subfield controlled by the asset auth flags, then the vote will either be a contract-wide vote (all assets vote on the referendum/initiative) or an asset-wide vote (all assets vote on the referendum/initiative).  The voting system specifies the voting rules.</td>
            <td class="a10">bool</td>
            <td class="a11"></td>
        </tr>

        <tr>
            <td class="a10">Qty of Tokens</td>
            <td class="a10">TokenQty</td>
            <td class="a10">8</td>
            <td class="a10" style="word-break:break-all">
                1000000
            </td>
            <td class="a10">Quantity of token - 0 is valid. Fungible 'shares' of the Asset. 1 is used for non-fungible tokens.  Asset IDs become the non-fungible Asset ID and many Asset IDs can be associated with a particular Contract.</td>
            <td class="a10">uint</td>
            <td class="a11"></td>
        </tr>

        <tr>
            <td class="a10">Contract Fee Currency</td>
            <td class="a10">ContractFeeCurrency</td>
            <td class="a10">3</td>
            <td class="a10" style="word-break:break-all">
                AUD
            </td>
            <td class="a10">BSV, USD, AUD, EUR, etc.</td>
            <td class="a10">fixedchar</td>
            <td class="a11"></td>
        </tr>

        <tr>
            <td class="a10">Contract Fee Var</td>
            <td class="a10">ContractFeeVar</td>
            <td class="a10">4</td>
            <td class="a10" style="word-break:break-all">
                0.005
            </td>
            <td class="a10">Percent of the value of the transaction</td>
            <td class="a10">float</td>
            <td class="a11"></td>
        </tr>

        <tr>
            <td class="a10">Contract Fee Fixed</td>
            <td class="a10">ContractFeeFixed</td>
            <td class="a10">4</td>
            <td class="a10" style="word-break:break-all">
                0.01
            </td>
            <td class="a10">Fixed fee (payment made in BSV)</td>
            <td class="a10">float</td>
            <td class="a11"></td>
        </tr>

        <tr>
            <td class="a10">Asset Payload</td>
            <td class="a10">AssetPayload</td>
            <td class="a10">16</td>
            <td class="a10" style="word-break:break-all">
                some data
            </td>
            <td class="a10">Payload length is dependent on the asset type. Each asset is made up of a defined set of information pertaining to the specific asset type, and may contain fields of variable length type (nvarchar8, 16, 32)</td>
            <td class="a10">varbin</td>
            <td class="a11"></td>
        </tr>

        <tr>
            <td class="a10">Asset Revision</td>
            <td class="a10">Asset Revision</td>
            <td class="a10">8</td>
            <td class="a10" style="word-break:break-all">
                0000000000000000000000000000000000000000000000000001000110111111
            </td>
            <td class="a10">Counter 0 - 65,535</td>
            <td class="a10">uint</td>
            <td class="a11"></td>
        </tr>

        <tr>
            <td class="a10">Timestamp</td>
            <td class="a10">Timestamp</td>
            <td class="a10">8</td>
            <td class="a10" style="word-break:break-all">
                1551767413250187179
            </td>
            <td class="a10">Timestamp in nanoseconds of when the smart contract created the action.</td>
            <td class="a10">timestamp</td>
            <td class="a11">Cannot be changed by issuer, operator. Smart contract controls.</td>
        </tr>

    </table>
</div>


<div class="ui modal" id="type-Header">
    <i class="close icon"></i>
    <div class="content docs-content">
        <table class="ui table">
            <tr style='height:19px;'>
                <th style="width:9%" class="s1">Label</th>
                <th style="width:9%" class="s1">Name</th>
                <th style="width:2%" class="s1">Bytes</th>
                <th style="width:29%" class="s1">Example Values</th>
                <th style="width:26%" class="s1">Comments</th>
                <th style="width:5%" class="s1">Data Type</th>
                <th style="width:14%" class="s2">Amendment Restrictions</th>
            </tr>
        </table>
    </div>
</div>

