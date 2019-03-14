
# Transfer Action

A Token Owner(s) Sends, Excahnges or Swaps a token(s) or Bitcoin for a token(s) or Bitcoin.  Can be as simple as sending a single token to a receiver.  Or can be as complex as many senders sending many different assets - controlled by many different smart contracts - to a number of receivers.  This action also supports atomic swaps (tokens for tokens).  Since many parties and contracts can be involved in a transfer and the corresponding settlement action, the partially signed T1 and T2 actions will need to be passed around on-chain with an M1 action, or off-chain.

The following breaks down the construction of a Transfer Action. The action is constructed by building a single string from each of the elements in order.

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
            <td class="s5" rowspan="7">Metadata (OP_RETURN Payload)</td>
            <td class="t6" colspan="7">
                <a href="javascript:;" data-popover="type-Header">
                   Header - Click to show content
                </a>
             </td>
        </tr>

        <tr>
            <td class="t6" colspan="7">
                <a href="javascript:;" data-popover="type-AssetTransfer">
                   Assets - Click to show content
                </a>
            </td>
        </tr>

        <tr>
            <td class="t10">Offer Expiry</td>
            <td class="t10">OfferExpiry</td>
            <td class="t10">8</td>
            <td class="t10" style="word-break:break-all">
                Sun May 06 2018 06:00:00 GMT+1000 (AEST)
            </td>
            <td class="t10">This prevents any party from holding on to the partially signed message as a form of an option.  Eg. the exchange at this price is valid for 30 mins.</td>
            <td class="t10">time</td>
            <td class="t11"></td>
        </tr>

        <tr>
            <td class="t10">Exchange Fee Currency</td>
            <td class="t10">ExchangeFeeCurrency</td>
            <td class="t10">3</td>
            <td class="t10" style="word-break:break-all">
                AUD
            </td>
            <td class="t10">BSV, USD, AUD, EUR, etc.</td>
            <td class="t10">fixedchar</td>
            <td class="t11"></td>
        </tr>

        <tr>
            <td class="t10">Exchange Fee Variable</td>
            <td class="t10">ExchangeFeeVar</td>
            <td class="t10">4</td>
            <td class="t10" style="word-break:break-all">
                0.005
            </td>
            <td class="t10">Percent of the value of the transaction</td>
            <td class="t10">float</td>
            <td class="t11"></td>
        </tr>

        <tr>
            <td class="t10">Exchange Fee Fixed</td>
            <td class="t10">ExchangeFeeFixed</td>
            <td class="t10">4</td>
            <td class="t10" style="word-break:break-all">
                0.01
            </td>
            <td class="t10">Fixed fee</td>
            <td class="t10">float</td>
            <td class="t11"></td>
        </tr>

        <tr>
            <td class="t6" colspan="7">
                <a href="javascript:;" data-popover="type-Address">
                   Exchange Fee Address - Click to show content
                </a>
            </td>
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

<div class="ui modal" id="type-AssetTransfer">
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
            <tr>
                <td class="t10">Asset Type</td>
                <td class="t10">AssetType</td>
                <td class="t10">3</td>
                <td class="t10" style="word-break:break-all">SHC</td>
                <td class="t10">eg. Share, Bond, Ticket. All characters must be capitalised.</td>
                <td class="t10">fixedchar</td>
                <td class="t11"></td>
            </tr>
            <tr>
                <td class="t10">Asset ID</td>
                <td class="t10">AssetID</td>
                <td class="t10">32</td>
                <td class="t10" style="word-break:break-all">apm2qsznhks23z8d83u41s8019hyri3i</td>
                <td class="t10">Randomly generated base58 string.  Each Asset ID should be unique.  However, a Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type can be the leading bytes - a convention - to make it easy to identify that it is a token by humans.</td>
                <td class="t10">fixedchar</td>
                <td class="t11"></td>
            </tr>
            <tr>
                <td class="t10">Asset Senders</td>
                <td class="t10">AssetSenders</td>
                <td class="t10">0</td>
                <td class="t10" style="word-break:break-all"></td>
                <td class="t10">Each element has the value of tokens to be spent from the input address, which is referred to by the index.</td>
                <td class="t10">QuantityIndex[]</td>
                <td class="t11"></td>
            </tr>
            <tr>
                <td class="t10">Token Receivers</td>
                <td class="t10">AssetReceivers</td>
                <td class="t10">0</td>
                <td class="t10" style="word-break:break-all"></td>
                <td class="t10">Each element has the value of tokens to be received by the output address, which is referred to by the index.</td>
                <td class="t10">TokenReceiver[]</td>
                <td class="t11"></td>
            </tr>
        </table>
    </div>
</div>

<div class="ui modal" id="type-Address">
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
            <tr>
                <td class="t10">Address</td>
                <td class="t10">Address</td>
                <td class="t10">20</td>
                <td class="t10" style="word-break:break-all">1HQ2ULuD7T5ykaucZ3KmTo4i29925Qa6ic</td>
                <td class="t10">Public address where the token balance will be changed.</td>
                <td class="t10">bin</td>
                <td class="t11"></td>
            </tr>
        </table>
    </div>
</div>

