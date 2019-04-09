# SettlementRequest

A message that contains a multi-contract settlement that needs settlement data added by another contract. Sent to another contract to request data be added.



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
            <td class="110">Version</td>
            <td class="110">uint</td>
            <td class="110">Payload Version</td>
            <td class="110">1</td>
            <td class="110"></td>
            <td class="110"></td>
        </tr>
        <tr>
            <td class="110">Timestamp</td>
            <td class="110">Timestamp</td>
            <td class="110">Timestamp in nanoseconds for when the message sender creates the transaction.</td>
            <td class="110">0</td>
            <td class="110">1551767413250187179</td>
            <td class="110"></td>
        </tr>
        <tr>
            <td class="110">Transfer Tx Id</td>
            <td class="110">TxId</td>
            <td class="110">Tx Id of the transfer request transaction that triggered this message.</td>
            <td class="110">0</td>
            <td class="110"></td>
            <td class="110"></td>
        </tr>
        <tr>
            <td class="110">Contract Fees</td>
            <td class="110">TargetAddress[]</td>
            <td class="110">Contract fees and addresses(PKHs) where fees should be paid. Added by each contract as settlement data is added.</td>
            <td class="110">8</td>
            <td class="110"></td>
            <td class="110"></td>
        </tr>
        <tr>
            <td class="110">Settlement</td>
            <td class="110">varbin</td>
            <td class="110">Serialized settlement OP_RETURN that needs data added by another contract.</td>
            <td class="110">32</td>
            <td class="110"></td>
            <td class="110"></td>
        </tr>
    </table>
</div>
