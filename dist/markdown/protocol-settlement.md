


# Settlement Action

Settlement Action -  Settles the transfer request of bitcoins and tokens from transfer (T1) actions.

The following breaks down the construction of a Settlement Action. The action is constructed by building a single string from each of the elements in order.

<div class="ritz grid-container" dir="ltr">
    <table class="waffle" cellspacing="0" cellpadding="0" table-layout=fixed width=100%>
         <tr style='height:19px;'>
            <th style="width:9%" class="s0">Label</th>
            <th style="width:9%" class="s1">Name</th>
            <th style="width:2%" class="s1">Bytes</th>
            <th style="width:25%" class="s1">Example Values</th>
            <th style="width:36%" class="s1">Comments</th>
            <th style="width:5%" class="s1">Data Type</th>
            <th class="s1">Amendment Restrictions</th>
        </tr>
        <tr>
            <td class="t5" colspan="7">
                <a href="javascript:;" data-popover="type-Header">
                   Header - Click to show content
                </a>
             </td>
        </tr>
        <tr>
            <td class="t5" colspan="7">
                <a href="javascript:;" data-popover="type-AssetSettlement">
                   Assets - Click to show content
                </a>
            </td>
        </tr>
        <tr>
            <td class="t9">Timestamp</td>
            <td class="t10">Timestamp</td>
            <td class="t10">8</td>
            <td class="t10">1551767413250187179</td>
            <td class="t10">Timestamp in nanoseconds of when the smart contract created the action.</td>
            <td class="t10">timestamp</td>
            <td class="t10">Cannot be changed by issuer, operator. Smart contract controls.</td>
        </tr>
    </table>
</div>

##Settlement Action Transaction Summary

<div class="ritz grid-container" dir="ltr">
    <table class="waffle" cellspacing="0" cellpadding="0" table-layout=fixed width=100%>
         <tr style='height:19px;'>
            <th class="s0" colspan="6">Smart Contract Operator Fee: 0</th>
       </tr>
         <tr style='height:19px;'>
            <th style="width:10%" class="s0">Index (input)</th>
            <th style="width:20%" class="s1">Txn inputs</th>
            <th style="width:20%" class="s1">Comments</th>
            <th style="width:10%" class="s1">Index (output)</th>
            <th style="width:20%" class="s1">Txn outputs</th>
            <th class="s1">Comments</th>
       </tr>
       <tr>
            <td class="t5">[{AssetXContract Contract Public Address (Asset X) }]</td>
            <td class="t6">.</td>
            <td class="t6">.</td>
            <td class="t10">.</td>
            <td class="t10">.</td>
            <td class="t10">.</td>
        </tr>
    </table>
</div>



<div class="ui modal" id="type-Header">
    <i class="close icon"></i>
    <div class="content docs-content">
        <table class="ui table">
            <tr style='height:19px;'>
                <th style="width:5%" class="s1">Label</th>
                <th style="width:9%" class="s1">Name</th>
                <th style="width:3%" class="s1">Bytes</th>
                <th style="width:33%" class="s1">Example Values</th>
                <th style="width:26%" class="s1">Comments</th>
                <th style="width:5%" class="s1">Data Type</th>
                <th class="s2">Amendment Restrictions</th>
            </tr>
        </table>
    </div>
</div>

<div class="ui modal" id="type-AssetSettlement">
    <i class="close icon"></i>
    <div class="content docs-content">
        <table class="ui table">
            <tr style='height:19px;'>
                <th style="width:5%" class="s1">Label</th>
                <th style="width:9%" class="s1">Name</th>
                <th style="width:3%" class="s1">Bytes</th>
                <th style="width:33%" class="s1">Example Values</th>
                <th style="width:26%" class="s1">Comments</th>
                <th style="width:5%" class="s1">Data Type</th>
                <th class="s2">Amendment Restrictions</th>
            </tr>
            <tr>
                <td class="t10">Asset Type</td>
                <td class="t10">AssetType</td>
                <td class="t10">3</td>
                <td class="t10" style="word-break:break-all">SHC</td>
                <td class="t10">eg. Share, Bond, Ticket. All characters must be capitalised.</td>
                <td class="t10">fixedchar</td>
                <td class="t10"></td>
            </tr>
            <tr>
                <td class="t10">Asset ID</td>
                <td class="t10">AssetID</td>
                <td class="t10">32</td>
                <td class="t10" style="word-break:break-all">apm2qsznhks23z8d83u41s8019hyri3i</td>
                <td class="t10">Randomly generated base58 string.  Each Asset ID should be unique.  However, a Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type can be the leading bytes - a convention - to make it easy to identify that it is a token by humans.</td>
                <td class="t10">fixedchar</td>
                <td class="t10"></td>
            </tr>
            <tr>
                <td class="t10">Settlements[]</td>
                <td class="t10">Settlements</td>
                <td class="t10">0</td>
                <td class="t10" style="word-break:break-all"></td>
                <td class="t10">Each element contains the resulting token balance of Asset X for the output Address, which is referred to by the index.</td>
                <td class="t10">QuantityIndex[]</td>
                <td class="t10"></td>
            </tr>
        </table>
    </div>
</div>

