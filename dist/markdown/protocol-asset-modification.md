
# Asset Modification Action

Asset Modification Action -  Token Dilutions, Call Backs/Revocations, burning etc.

The following breaks down the construction of a Asset Modification Action. The action is constructed by building a single string from each of the elements in order.

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
            <td class="s5" rowspan="6">Metadata (OP_RETURN Payload)</td>
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
            <td class="a10">Asset Revision</td>
            <td class="a10">AssetRevision</td>
            <td class="a10">8</td>
            <td class="a10" style="word-break:break-all">
                0
            </td>
            <td class="a10">Counter. (Subfield cannot be manually changed by Asset Modification Action.  Only SC can increment by 1 with each AC action. SC will reject AM actions where the wrong asset revision has been selected. </td>
            <td class="a10">uint</td>
            <td class="a11">Cannot be Amended</td>
        </tr>

        <tr>
            <td class="a6" colspan="7">
                <a href="javascript:;" data-popover="type-Amendment">
                   Modifications - Click to show content
                </a>
            </td>
        </tr>

        <tr>
            <td class="a10">Ref Tx-ID</td>
            <td class="a10">RefTxID</td>
            <td class="a10">32</td>
            <td class="a10" style="word-break:break-all">
                a8700385d4cc62628cc34629862121f84e6237689de8e45e151dcbc8cf30b33d
            </td>
            <td class="a10">Tx-ID of the associated Result action (governance) that permitted the modifications.</td>
            <td class="a10">sha256</td>
            <td class="a11"></td>
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

<div class="ui modal" id="type-Amendment">
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
                <td class="a10">Field Index</td>
                <td class="a10">FieldIndex</td>
                <td class="a10">1</td>
                <td class="a10" style="word-break:break-all">2</td>
                <td class="a10">Index of the field to be amended.</td>
                <td class="a10">uint</td>
                <td class="a11">A field with a complex array type uses the same FieldIndex value for all elements. For example, in C1 the VotingSystems field is FieldIndex 16. Indexes are zero based.</td>
            </tr>
            <tr>
                <td class="a10">Element</td>
                <td class="a10">Element</td>
                <td class="a10">2</td>
                <td class="a10" style="word-break:break-all">0</td>
                <td class="a10">Specifies the element of the complex array type to be amended. This only applies to array types, and has no meaning for a simple type such as uint64, string, byte or byte[]. Specifying a value > 0 for a simple type will result in a Rejection.</td>
                <td class="a10">uint</td>
                <td class="a11">To specify the 3rd VotingSystem of a Contract, the value 2 would be given. Indexes are zero based.</td>
            </tr>
            <tr>
                <td class="a10">Subfield Index</td>
                <td class="a10">SubfieldIndex</td>
                <td class="a10">1</td>
                <td class="a10" style="word-break:break-all">1</td>
                <td class="a10">Index of the subfield to be amended. This only applies to specific fields of an element in an array. This is used to specify which field of the array element the amendment applies to.</td>
                <td class="a10">uint</td>
                <td class="a11">For example to specify the 2nd field of a VotingSystem, value 1 would be given.</td>
            </tr>
            <tr>
                <td class="a10">Delete Element</td>
                <td class="a10">DeleteElement</td>
                <td class="a10">0</td>
                <td class="a10" style="word-break:break-all"></td>
                <td class="a10">1 = remove the element listed in the Element field, 0 means this is not a delete operation. The DeleteElement field only applies to a particilar element of a complex array type. For example, it could be used to remove a particular VotingSystem from a Contract.</td>
                <td class="a10">bool</td>
                <td class="a11"></td>
            </tr>
            <tr>
                <td class="a10">Data</td>
                <td class="a10">Data</td>
                <td class="a10">32</td>
                <td class="a10" style="word-break:break-all"></td>
                <td class="a10">New data for the amended subfield. Data type depends on the the type of the field being amended.</td>
                <td class="a10">varchar</td>
                <td class="a11">The bytes should be in an format appropriate for the field being modified.</td>
            </tr>
        </table>
    </div>
</div>

