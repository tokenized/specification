# RevertedTx

A message that contains a bitcoin transaction that was accepted by the network or an agent and then invalidated by a reorg, or zero conf double spend. This serves as on chain evidence of the sending party's signatures and approval for the given transaction.



<div class="ritz grid-container" dir="ltr">
    <table class="waffle" cellspacing="0" cellpadding="0" table-layout=fixed width=100%>
         <tr style='height:19px;'>
            <th style="width:20%" class="s1">Field</th>
            <th style="width:10%" class="s1">Type</th>
            <th style="width:15%" class="s1">Description</th>
            <th style="width:20%" class="s1">Size</th>
            <th style="width:20%" class="s1">Example</th>
            <th class="s1">Notes</th>
        </tr>
        <tr>
            <td class="010">Version</td>
            <td class="010">uint</td>
            <td class="010">Payload Version</td>
            <td class="010">1</td>
            <td class="010"></td>
            <td class="010"></td>
        </tr>
        <tr>
            <td class="010">Timestamp</td>
            <td class="010">Timestamp</td>
            <td class="010">Timestamp in nanoseconds for when the message sender creates the transaction.</td>
            <td class="010">0</td>
            <td class="010">1551767413250187179</td>
            <td class="010"></td>
        </tr>
        <tr>
            <td class="010">Transaction</td>
            <td class="010">varbin</td>
            <td class="010">Serialized bitcoin transaction that was reverted/invalidated after being accepted.</td>
            <td class="010">32</td>
            <td class="010"></td>
            <td class="010"></td>
        </tr>
    </table>
</div>
