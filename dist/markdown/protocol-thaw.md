
# Thaw Action

Thaw Action -  to be used to comply with contractual obligations or legal requirements.  The Alleged Offender's tokens will be unfrozen to allow them to resume normal exchange and governance activities.

The following breaks down the construction of a Thaw Action. The action is constructed by building a single string from each of the elements in order.

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
            <td class="s5" rowspan="4">Metadata (OP_RETURN Payload)</td>
            <td class="e6" colspan="7">
                <a href="javascript:;" data-popover="type-Header">
                   Header - Click to show content
                </a>
             </td>
        </tr>

        <tr>
            <td class="e6" colspan="7">
                <a href="javascript:;" data-popover="type-Address">
                   Addresses - Click to show content
                </a>
            </td>
        </tr>

        <tr>
            <td class="e10">Ref Txn ID</td>
            <td class="e10">RefTxnID</td>
            <td class="e10">32</td>
            <td class="e10" style="word-break:break-all">
                f3318be9fb3f73e53b29868beae46b42911c2116f979a5d3284face90746cb37
            </td>
            <td class="e10">The related freeze action.</td>
            <td class="e10">sha256</td>
            <td class="e11"></td>
        </tr>

        <tr>
            <td class="e10">Timestamp</td>
            <td class="e10">Timestamp</td>
            <td class="e10">8</td>
            <td class="e10" style="word-break:break-all">
                1551767413250187179
            </td>
            <td class="e10">Timestamp in nanoseconds of when the smart contract created the action.</td>
            <td class="e10">timestamp</td>
            <td class="e11">Cannot be changed by issuer, operator. Smart contract controls.</td>
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
                <td class="e10">Address</td>
                <td class="e10">Address</td>
                <td class="e10">20</td>
                <td class="e10" style="word-break:break-all">1HQ2ULuD7T5ykaucZ3KmTo4i29925Qa6ic</td>
                <td class="e10">Public address where the token balance will be changed.</td>
                <td class="e10">bin</td>
                <td class="e11"></td>
            </tr>
        </table>
    </div>
</div>

