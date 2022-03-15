"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var RejectCodes;
(function (RejectCodes) {
    // Success - No failure. This code should not be used in a reject message.
    RejectCodes[RejectCodes["RejectSuccess"] = 0] = "RejectSuccess";
    // MsgMalformed - The OP_RETURN message is malformed. It doesn't pass data validation or something about it isn't according to protocol.
    RejectCodes[RejectCodes["RejectMsgMalformed"] = 1] = "RejectMsgMalformed";
    // TxMalformed - The Bitcoin tx is malformed. Incorrect inputs/outputs or something similar.
    RejectCodes[RejectCodes["RejectTxMalformed"] = 2] = "RejectTxMalformed";
    // Timeout - A dependency, other contract/service, has failed to complete before the smart contract's timeout.
    RejectCodes[RejectCodes["RejectTimeout"] = 3] = "RejectTimeout";
    // ContractMoved - The contract has been moved to a different address. Please find the addres change message and send requests to new address.
    RejectCodes[RejectCodes["RejectContractMoved"] = 4] = "RejectContractMoved";
    // DoubleSpend - A double spend attempt has been detected.
    RejectCodes[RejectCodes["RejectDoubleSpend"] = 5] = "RejectDoubleSpend";
    // ContractExists - The contract already exists and can't be recreated.
    RejectCodes[RejectCodes["RejectContractExists"] = 10] = "RejectContractExists";
    // ContractInstrumentQtyReduction - Sent when a CA tries to reduce the number of allowed instruments below the number of instruments that already exist for this contract.
    RejectCodes[RejectCodes["RejectContractInstrumentQtyReduction"] = 12] = "RejectContractInstrumentQtyReduction";
    // ContractFixedQuantity - Sent when the administration attempted to increase the quantity of instruments in a contract beyond the maximum number allowed.
    RejectCodes[RejectCodes["RejectContractFixedQuantity"] = 13] = "RejectContractFixedQuantity";
    // ContractAuthFlags - The contract auth flags don't permit the action requested.
    RejectCodes[RejectCodes["RejectContractAuthFlags"] = 14] = "RejectContractAuthFlags";
    // ContractExpired - The contract is expired so can no longer accept requests.
    RejectCodes[RejectCodes["RejectContractExpired"] = 15] = "RejectContractExpired";
    // ContractFrozen - The contract is frozen and the request is not permitted while frozen.
    RejectCodes[RejectCodes["RejectContractFrozen"] = 16] = "RejectContractFrozen";
    // ContractRevision - The revision in a contract amendment is incorrect.
    RejectCodes[RejectCodes["RejectContractRevision"] = 17] = "RejectContractRevision";
    // ContractNotPermitted - Action not permitted by contract.
    RejectCodes[RejectCodes["RejectContractNotPermitted"] = 18] = "RejectContractNotPermitted";
    // ContractBothOperatorsRequired - Both operators signatures are required to perform this action.
    RejectCodes[RejectCodes["RejectContractBothOperatorsRequired"] = 19] = "RejectContractBothOperatorsRequired";
    // InstrumentCodeExists - The instrument code specified already exists and can't be reused.
    RejectCodes[RejectCodes["RejectInstrumentCodeExists"] = 20] = "RejectInstrumentCodeExists";
    // InstrumentNotFound - The instrument code is not found.
    RejectCodes[RejectCodes["RejectInstrumentNotFound"] = 21] = "RejectInstrumentNotFound";
    // InstrumentAuthFlags - The instrument auth flags don't permit the action requested.
    RejectCodes[RejectCodes["RejectInstrumentAuthFlags"] = 22] = "RejectInstrumentAuthFlags";
    // InstrumentFrozen - The instrument is frozen and the request is not permitted while frozen.
    RejectCodes[RejectCodes["RejectInstrumentFrozen"] = 23] = "RejectInstrumentFrozen";
    // InstrumentRevision - The revision in an instrument amendment is incorrect.
    RejectCodes[RejectCodes["RejectInstrumentRevision"] = 24] = "RejectInstrumentRevision";
    // InstrumentNotPermitted - Action not permitted by instrument.
    RejectCodes[RejectCodes["RejectInstrumentNotPermitted"] = 25] = "RejectInstrumentNotPermitted";
    // TransferSelf - Transfers with the sender and receiver addresses the same are not permitted.
    RejectCodes[RejectCodes["RejectTransferSelf"] = 30] = "RejectTransferSelf";
    // TransferExpired - The transfer has expired.
    RejectCodes[RejectCodes["RejectTransferExpired"] = 31] = "RejectTransferExpired";
    // HoldingsFrozen - Holdings are frozen, so the request can't be completed.
    RejectCodes[RejectCodes["RejectHoldingsFrozen"] = 32] = "RejectHoldingsFrozen";
    // HolderProposalProhibited - Holders are not permitted to make proposals.
    RejectCodes[RejectCodes["RejectHolderProposalProhibited"] = 40] = "RejectHolderProposalProhibited";
    // ProposalConflicts - The proposal conflicts with an unapplied proposal.
    RejectCodes[RejectCodes["RejectProposalConflicts"] = 41] = "RejectProposalConflicts";
    // VoteNotFound - The vote ID referenced is not found.
    RejectCodes[RejectCodes["RejectVoteNotFound"] = 42] = "RejectVoteNotFound";
    // VoteClosed - The vote has closed and ballots are no longer permitted.
    RejectCodes[RejectCodes["RejectVoteClosed"] = 43] = "RejectVoteClosed";
    // BallotAlreadyCounted - The ballot has already been counted for this address.
    RejectCodes[RejectCodes["RejectBallotAlreadyCounted"] = 44] = "RejectBallotAlreadyCounted";
    // VoteSystemNotPermitted - The voting system isn't permitted for this request.
    RejectCodes[RejectCodes["RejectVoteSystemNotPermitted"] = 45] = "RejectVoteSystemNotPermitted";
    // InsufficientTxFeeFunding - Insufficient bitcoin quantities for response transaction fees.
    RejectCodes[RejectCodes["RejectInsufficientTxFeeFunding"] = 60] = "RejectInsufficientTxFeeFunding";
    // InsufficientValue - Insufficient bitcoin quantity in inputs to fund request.
    RejectCodes[RejectCodes["RejectInsufficientValue"] = 61] = "RejectInsufficientValue";
    // InsufficientQuantity - Insufficient token holdings to for request.
    RejectCodes[RejectCodes["RejectInsufficientQuantity"] = 62] = "RejectInsufficientQuantity";
    // NotAdministration - The requestor is not the administration and is required to be for this request.
    RejectCodes[RejectCodes["RejectNotAdministration"] = 70] = "RejectNotAdministration";
    // NotOperator - The requestor is not the operator and is required to be for this request.
    RejectCodes[RejectCodes["RejectNotOperator"] = 71] = "RejectNotOperator";
    // UnauthorizedAddress - The address specified is not permitted for this request.
    RejectCodes[RejectCodes["RejectUnauthorizedAddress"] = 72] = "RejectUnauthorizedAddress";
    // InvalidSignature - The signature provided is not valid. This is for signatures included within OP_RETURN data. Not bitcoin transaction signature scripts.
    RejectCodes[RejectCodes["RejectInvalidSignature"] = 80] = "RejectInvalidSignature";
})(RejectCodes = exports.RejectCodes || (exports.RejectCodes = {}));
