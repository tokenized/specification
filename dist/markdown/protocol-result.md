
# Result Action

Result Action -  Once a vote has been completed the results are published.

The following breaks down the construction of a Result Action. The action is constructed by building a single string from each of the elements in order.

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
            <td class="s5" rowspan="10">Metadata (OP_RETURN Payload)</td>
            <td class="g6" colspan="7">
                <a href="javascript:;" data-popover="type-Header">
                   Header - Click to show content
                </a>
             </td>
        </tr>

        <tr>
            <td class="g10">Asset Type</td>
            <td class="g10">AssetType</td>
            <td class="g10">3</td>
            <td class="g10" style="word-break:break-all">
                SHC
            </td>
            <td class="g10">eg. Share, Bond, Ticket</td>
            <td class="g10">fixedchar</td>
            <td class="g11"></td>
        </tr>

        <tr>
            <td class="g10">Asset ID</td>
            <td class="g10">AssetID</td>
            <td class="g10">32</td>
            <td class="g10" style="word-break:break-all">
                apm2qsznhks23z8d83u41s8019hyri3i
            </td>
            <td class="g10">Randomly generated base58 string.  Each Asset ID should be unique.  However, a Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type can be the leading bytes - a convention - to make it easy to identify that it is a token by humans. If its a Contract vote then can be null.</td>
            <td class="g10">fixedchar</td>
            <td class="g11"></td>
        </tr>

        <tr>
            <td class="g10">Proposal</td>
            <td class="g10">Proposal</td>
            <td class="g10">1</td>
            <td class="g10" style="word-break:break-all">
                0
            </td>
            <td class="g10">1 for a Proposal, 0 for an initiative that is requesting changes to specific subfields for modification. If this field is true, the subfields should be empty.  The smart contract cannot interpret the results of a vote when Proposal = 1.  All meaning is interpreted by the token owners and smart contract simply facilates the record keeping.  When Proposal = 0, the smart contract always assumes the first choice is a 'yes', or 'pass', if the threshold is met, and will process the proposed changes accordingly.</td>
            <td class="g10">bool</td>
            <td class="g11"></td>
        </tr>

        <tr>
            <td class="g6" colspan="7">
                <a href="javascript:;" data-popover="type-Amendment">
                   Proposed Changes - Click to show content
                </a>
            </td>
        </tr>

        <tr>
            <td class="g10">Vote Txn ID</td>
            <td class="g10">VoteTxnID</td>
            <td class="g10">32</td>
            <td class="g10" style="word-break:break-all">
                f2318be9fb3f73e53a29868beae46b42911c2116f979a5d3284face90746cb37
            </td>
            <td class="g10">Link to the Vote Action txn.</td>
            <td class="g10">sha256</td>
            <td class="g11"></td>
        </tr>

        <tr>
            <td class="g10">VoteOptionsCount</td>
            <td class="g10">VoteOptionsCount</td>
            <td class="g10">1</td>
            <td class="g10" style="word-break:break-all">
                1
            </td>
            <td class="g10">Number of Vote Options to follow.</td>
            <td class="g10">uint</td>
            <td class="g11"></td>
        </tr>

        <tr>
            <td class="g10">Option X Tally</td>
            <td class="g10">OptionXTally</td>
            <td class="g10">8</td>
            <td class="g10" style="word-break:break-all">
                3000
            </td>
            <td class="g10">Number of valid votes counted for Option X</td>
            <td class="g10">uint</td>
            <td class="g11"></td>
        </tr>

        <tr>
            <td class="g10">Result</td>
            <td class="g10">Result</td>
            <td class="g10">8</td>
            <td class="g10" style="word-break:break-all">
                2
            </td>
            <td class="g10">Length 1-255 bytes. 0 is not valid. The Option with the most votes. In the event of a draw for 1st place, all winning options are listed. </td>
            <td class="g10">varchar</td>
            <td class="g11"></td>
        </tr>

        <tr>
            <td class="g10">Timestamp</td>
            <td class="g10">Timestamp</td>
            <td class="g10">8</td>
            <td class="g10" style="word-break:break-all">
                1551767413250187179
            </td>
            <td class="g10">Timestamp in nanoseconds of when the smart contract created the action.</td>
            <td class="g10">timestamp</td>
            <td class="g11">Cannot be changed by issuer, operator. Smart contract controls.</td>
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
                <td class="g10">Field Index</td>
                <td class="g10">FieldIndex</td>
                <td class="g10">1</td>
                <td class="g10" style="word-break:break-all">2</td>
                <td class="g10">Index of the field to be amended.</td>
                <td class="g10">uint</td>
                <td class="g11">A field with a complex array type uses the same FieldIndex value for all elements. For example, in C1 the VotingSystems field is FieldIndex 16. Indexes are zero based.</td>
            </tr>
            <tr>
                <td class="g10">Element</td>
                <td class="g10">Element</td>
                <td class="g10">2</td>
                <td class="g10" style="word-break:break-all">0</td>
                <td class="g10">Specifies the element of the complex array type to be amended. This only applies to array types, and has no meaning for a simple type such as uint64, string, byte or byte[]. Specifying a value > 0 for a simple type will result in a Rejection.</td>
                <td class="g10">uint</td>
                <td class="g11">To specify the 3rd VotingSystem of a Contract, the value 2 would be given. Indexes are zero based.</td>
            </tr>
            <tr>
                <td class="g10">Subfield Index</td>
                <td class="g10">SubfieldIndex</td>
                <td class="g10">1</td>
                <td class="g10" style="word-break:break-all">1</td>
                <td class="g10">Index of the subfield to be amended. This only applies to specific fields of an element in an array. This is used to specify which field of the array element the amendment applies to.</td>
                <td class="g10">uint</td>
                <td class="g11">For example to specify the 2nd field of a VotingSystem, value 1 would be given.</td>
            </tr>
            <tr>
                <td class="g10">Delete Element</td>
                <td class="g10">DeleteElement</td>
                <td class="g10">0</td>
                <td class="g10" style="word-break:break-all"></td>
                <td class="g10">1 = remove the element listed in the Element field, 0 means this is not a delete operation. The DeleteElement field only applies to a particilar element of a complex array type. For example, it could be used to remove a particular VotingSystem from a Contract.</td>
                <td class="g10">bool</td>
                <td class="g11"></td>
            </tr>
            <tr>
                <td class="g10">Data</td>
                <td class="g10">Data</td>
                <td class="g10">32</td>
                <td class="g10" style="word-break:break-all"></td>
                <td class="g10">New data for the amended subfield. Data type depends on the the type of the field being amended.</td>
                <td class="g10">varchar</td>
                <td class="g11">The bytes should be in an format appropriate for the field being modified.</td>
            </tr>
        </table>
    </div>
</div>

