
# Settlement Action

Settlement Action -  Settles the transfer request of bitcoins and tokens from transfer (T1) actions.

The following breaks down the construction of a Settlement Action. The action is constructed by building a single string from each of the elements in order.

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
            <td class="s5" rowspan="8">Metadata (OP_RETURN Payload)</td>
            <td class="t6" colspan="7">
                <a href="javascript:;" data-popover="type-Header">
                   Header - Click to show content
                </a>
             </td>
        </tr>

        <tr>
            <td class="t10">Text Encoding</td>
            <td class="t10">TextEncoding</td>
            <td class="t10">1</td>
            <td class="t10" style="word-break:break-all">
                0
            </td>
            <td class="t10"> 0 = ASCII, 1 = UTF-8, 2 = UTF-16, 3 = Unicode.  Encoding applies to all 'text' data types. All 'string' types will always be encoded with ASCII.  Where string is selected, all fields will be ASCII.</td>
            <td class="t10">uint8</td>
            <td class="t11">Can be changed by Issuer or Operator at their discretion.</td>
        </tr>

        <tr>
            <td class="t10">Asset Count</td>
            <td class="t10">AssetCount</td>
            <td class="t10">1</td>
            <td class="t10" style="word-break:break-all">
                4
            </td>
            <td class="t10">The number of Assets specified by the Transfer action to be settled.</td>
            <td class="t10">uint8</td>
            <td class="t11"></td>
        </tr>

        <tr>
            <td class="t10">Asset Type X</td>
            <td class="t10">AssetTypeX</td>
            <td class="t10">3</td>
            <td class="t10" style="word-break:break-all">
                RRE
            </td>
            <td class="t10">eg. Share, Bond, Ticket</td>
            <td class="t10">string</td>
            <td class="t11"></td>
        </tr>

        <tr>
            <td class="t10">Asset ID X</td>
            <td class="t10">AssetIDX</td>
            <td class="t10">32</td>
            <td class="t10" style="word-break:break-all">
                apm2qsznhks23z8d83u41s8019hyri3i
            </td>
            <td class="t10">Randomly generated base58 string.  Each Asset ID should be unique.  However, a Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type can be the leading bytes - a convention - to make it easy to identify that it is a token by humans. </td>
            <td class="t10">string</td>
            <td class="t11"></td>
        </tr>

        <tr>
            <td class="t10">Asset X Settlements Count</td>
            <td class="t10">AssetXSettlementsCount</td>
            <td class="t10">1</td>
            <td class="t10" style="word-break:break-all">
                0
            </td>
            <td class="t10">Number of settlements for Asset X.</td>
            <td class="t10">uint8</td>
            <td class="t11"></td>
        </tr>

        <tr>
            <td class="t6" colspan="7">
                <a href="javascript:;" data-popover="type-QuantityIndex">
                   Asset X Address X Qty - Click to show content
                </a>
            </td>
        </tr>

        <tr>
            <td class="t10">Timestamp</td>
            <td class="t10">Timestamp</td>
            <td class="t10">8</td>
            <td class="t10" style="word-break:break-all">
                1551767413250187179
            </td>
            <td class="t10">Timestamp in nanoseconds of when the smart contract created the action.</td>
            <td class="t10">timestamp</td>
            <td class="t11">Cannot be changed by issuer, operator. Smart contract controls.</td>
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
            <tr>
                <td class="t10">Protocol Identifier</td>
                <td class="t10">ProtocolID</td>
                <td class="t10">13</td>
                <td class="t10" style="word-break:break-all">tokenized.com</td>
                <td class="t10">Tokenized ID Prefix.  tokenized.com</td>
                <td class="t10">string</td>
                <td class="t11"></td>
            </tr>
            <tr>
                <td class="t10">Push Data</td>
                <td class="t10">OpPushdata</td>
                <td class="t10">1</td>
                <td class="t10" style="word-break:break-all">77</td>
                <td class="t10">PACKET LENGTH, PUSHDATA1 (76), PUSHDATA2 (77), or PUSHDATA4 (78) depending on total size of action payload.</td>
                <td class="t10">opcode</td>
                <td class="t11">Cannot be changed by issuer, operator or smart contract.</td>
            </tr>
            <tr>
                <td class="t10">Length of Action Payload</td>
                <td class="t10">LenActionPayload</td>
                <td class="t10">2</td>
                <td class="t10" style="word-break:break-all">409</td>
                <td class="t10">Length of the action message (0 - 65,535 bytes). 0 if pushdata length <76B, 1 byte if PUSHDATA1 is used, 2 bytes if PUSHDATA2 and 4 bytes if PUSHDATA4.</td>
                <td class="t10">pushdata_length</td>
                <td class="t11">Depends on Action Payload</td>
            </tr>
            <tr>
                <td class="t10">Version</td>
                <td class="t10">Version</td>
                <td class="t10">1</td>
                <td class="t10" style="word-break:break-all">0</td>
                <td class="t10">255 reserved for additional versions. Tokenized protocol versioning.</td>
                <td class="t10">uint8</td>
                <td class="t11">Can be changed by Issuer or Operator at their discretion.  Smart Contract will reject if it hasn't been updated to interpret the specified version.</td>
            </tr>
            <tr>
                <td class="t10">Action Prefix</td>
                <td class="t10">ActionPrefix</td>
                <td class="t10">2</td>
                <td class="t10" style="word-break:break-all">C1</td>
                <td class="t10">Contract Offer: The Contract Offer Action allows the Issuer to initialize a smart contract by providing all the necessary information, including T&C's.  The Contract Offer Action can also be used to signal to a market actor that they want to buy/form a contract.</td>
                <td class="t10">string</td>
                <td class="t11">Cannot be changed by issuer, operator or smart contract.</td>
            </tr>
        </table>
    </div>
</div>

<div class="ui modal" id="type-QuantityIndex">
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
                <td class="t10">Index</td>
                <td class="t10">Index</td>
                <td class="t10">2</td>
                <td class="t10" style="word-break:break-all">0</td>
                <td class="t10">The index of the input sending the tokens</td>
                <td class="t10">uint16</td>
                <td class="t11"></td>
            </tr>
            <tr>
                <td class="t10">Quantity</td>
                <td class="t10">Quantity</td>
                <td class="t10">8</td>
                <td class="t10" style="word-break:break-all">100</td>
                <td class="t10">Number of tokens being sent</td>
                <td class="t10">uint64</td>
                <td class="t11"></td>
            </tr>
        </table>
    </div>
</div>

