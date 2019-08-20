package actions

const (
	// Success - No failure. This code should not be used in a reject message.
	RejectSuccess = 0

	// MsgMalformed - The OP_RETURN message is malformed. It doesn&#39;t pass data validation or something about it isn&#39;t according to protocol.
	RejectMsgMalformed = 1

	// TxMalformed - The Bitcoin tx is malformed. Incorrect inputs/outputs or something similar.
	RejectTxMalformed = 2

	// Timeout - A dependency, other contract/service, has failed to complete before the smart contract&#39;s timeout.
	RejectTimeout = 3

	// ContractMoved - The contract has been moved to a different address. Please find the addres change message and send requests to new address.
	RejectContractMoved = 4

	// DoubleSpend - A double spend attempt has been detected.
	RejectDoubleSpend = 5

	// ContractExists - The contract already exists and can&#39;t be recreated.
	RejectContractExists = 10

	// ContractAssetQtyReduction - Sent when a CA tries to reduce the number of allowed assets below the number of assets that already exist for this contract.
	RejectContractAssetQtyReduction = 12

	// ContractFixedQuantity - Sent when the administration attempted to increase the quantity of assets in a contract beyond the maximum number allowed.
	RejectContractFixedQuantity = 13

	// ContractAuthFlags - The contract auth flags don&#39;t permit the action requested.
	RejectContractAuthFlags = 14

	// ContractExpired - The contract is expired so can no longer accept requests.
	RejectContractExpired = 15

	// ContractFrozen - The contract is frozen and the request is not permitted while frozen.
	RejectContractFrozen = 16

	// ContractRevision - The revision in a contract amendment is incorrect.
	RejectContractRevision = 17

	// ContractNotPermitted - Action not permitted by contract.
	RejectContractNotPermitted = 18

	// ContractBothOperatorsRequired - Both operators signatures are required to perform this action.
	RejectContractBothOperatorsRequired = 19

	// AssetCodeExists - The asset code specified already exists and can&#39;t be reused.
	RejectAssetCodeExists = 20

	// AssetNotFound - The asset code is not found.
	RejectAssetNotFound = 21

	// AssetAuthFlags - The asset auth flags don&#39;t permit the action requested.
	RejectAssetAuthFlags = 22

	// AssetFrozen - The asset is frozen and the request is not permitted while frozen.
	RejectAssetFrozen = 23

	// AssetRevision - The revision in an asset amendment is incorrect.
	RejectAssetRevision = 24

	// AssetNotPermitted - Action not permitted by asset.
	RejectAssetNotPermitted = 25

	// TransferSelf - Transfers with the sender and receiver addresses the same are not permitted.
	RejectTransferSelf = 30

	// TransferExpired - The transfer has expired.
	RejectTransferExpired = 31

	// HoldingsFrozen - Holdings are frozen, so the request can&#39;t be completed.
	RejectHoldingsFrozen = 32

	// HolderProposalProhibited - Holders are not permitted to make proposals.
	RejectHolderProposalProhibited = 40

	// ProposalConflicts - The proposal conflicts with an unapplied proposal.
	RejectProposalConflicts = 41

	// VoteNotFound - The vote ID referenced is not found.
	RejectVoteNotFound = 42

	// VoteClosed - The vote has closed and ballots are no longer permitted.
	RejectVoteClosed = 43

	// BallotAlreadyCounted - The ballot has already been counted for this address.
	RejectBallotAlreadyCounted = 44

	// VoteSystemNotPermitted - The voting system isn&#39;t permitted for this request.
	RejectVoteSystemNotPermitted = 45

	// InsufficientTxFeeFunding - Insufficient bitcoin quantities for response transaction fees.
	RejectInsufficientTxFeeFunding = 60

	// InsufficientValue - Insufficient bitcoin quantity in inputs to fund request.
	RejectInsufficientValue = 61

	// InsufficientQuantity - Insufficient token holdings to for request.
	RejectInsufficientQuantity = 62

	// NotAdministration - The requestor is not the administration and is required to be for this request.
	RejectNotAdministration = 70

	// NotOperator - The requestor is not the operator and is required to be for this request.
	RejectNotOperator = 71

	// UnauthorizedAddress - The address specified is not permitted for this request.
	RejectUnauthorizedAddress = 72

	// InvalidSignature - The signature provided is not valid. This is for signatures included within OP_RETURN data. Not bitcoin transaction signature scripts.
	RejectInvalidSignature = 80
)

type RejectionCode struct {
	Name        string
	Label       string
	Description string
}

// TypeMapping holds a mapping of action codes to action types.
func RejectionData(code uint32) *RejectionCode {
	switch code {
	case RejectSuccess:
		return &RejectionCode{
			Name:        "Success",
			Label:       "Success",
			Description: "No failure. This code should not be used in a reject message.",
		}
	case RejectMsgMalformed:
		return &RejectionCode{
			Name:        "MsgMalformed",
			Label:       "Message Malformed",
			Description: "The OP_RETURN message is malformed. It doesn&#39;t pass data validation or something about it isn&#39;t according to protocol.",
		}
	case RejectTxMalformed:
		return &RejectionCode{
			Name:        "TxMalformed",
			Label:       "Transaction Malformed",
			Description: "The Bitcoin tx is malformed. Incorrect inputs/outputs or something similar.",
		}
	case RejectTimeout:
		return &RejectionCode{
			Name:        "Timeout",
			Label:       "Time Out",
			Description: "A dependency, other contract/service, has failed to complete before the smart contract&#39;s timeout.",
		}
	case RejectContractMoved:
		return &RejectionCode{
			Name:        "ContractMoved",
			Label:       "Contract Moved",
			Description: "The contract has been moved to a different address. Please find the addres change message and send requests to new address.",
		}
	case RejectDoubleSpend:
		return &RejectionCode{
			Name:        "DoubleSpend",
			Label:       "Double Spend",
			Description: "A double spend attempt has been detected.",
		}
	case RejectContractExists:
		return &RejectionCode{
			Name:        "ContractExists",
			Label:       "Contract Already Exists",
			Description: "The contract already exists and can&#39;t be recreated.",
		}
	case RejectContractAssetQtyReduction:
		return &RejectionCode{
			Name:        "ContractAssetQtyReduction",
			Label:       "Contract Asset Quantity Reduction",
			Description: "Sent when a CA tries to reduce the number of allowed assets below the number of assets that already exist for this contract.",
		}
	case RejectContractFixedQuantity:
		return &RejectionCode{
			Name:        "ContractFixedQuantity",
			Label:       "Contract Fixed Quantity",
			Description: "Sent when the administration attempted to increase the quantity of assets in a contract beyond the maximum number allowed.",
		}
	case RejectContractAuthFlags:
		return &RejectionCode{
			Name:        "ContractAuthFlags",
			Label:       "Contract Auth Flags Prohibit",
			Description: "The contract auth flags don&#39;t permit the action requested.",
		}
	case RejectContractExpired:
		return &RejectionCode{
			Name:        "ContractExpired",
			Label:       "Contract Expired",
			Description: "The contract is expired so can no longer accept requests.",
		}
	case RejectContractFrozen:
		return &RejectionCode{
			Name:        "ContractFrozen",
			Label:       "Contract Frozen",
			Description: "The contract is frozen and the request is not permitted while frozen.",
		}
	case RejectContractRevision:
		return &RejectionCode{
			Name:        "ContractRevision",
			Label:       "Contract Revision Incorrect",
			Description: "The revision in a contract amendment is incorrect.",
		}
	case RejectContractNotPermitted:
		return &RejectionCode{
			Name:        "ContractNotPermitted",
			Label:       "Contract Not Permitted",
			Description: "Action not permitted by contract.",
		}
	case RejectContractBothOperatorsRequired:
		return &RejectionCode{
			Name:        "ContractBothOperatorsRequired",
			Label:       "Contract BothOperatorsRequired",
			Description: "Both operators signatures are required to perform this action.",
		}
	case RejectAssetCodeExists:
		return &RejectionCode{
			Name:        "AssetCodeExists",
			Label:       "Asset Code Already Exists",
			Description: "The asset code specified already exists and can&#39;t be reused.",
		}
	case RejectAssetNotFound:
		return &RejectionCode{
			Name:        "AssetNotFound",
			Label:       "Asset Not Found",
			Description: "The asset code is not found.",
		}
	case RejectAssetAuthFlags:
		return &RejectionCode{
			Name:        "AssetAuthFlags",
			Label:       "Asset Auth Flags Prohibit",
			Description: "The asset auth flags don&#39;t permit the action requested.",
		}
	case RejectAssetFrozen:
		return &RejectionCode{
			Name:        "AssetFrozen",
			Label:       "Asset Frozen",
			Description: "The asset is frozen and the request is not permitted while frozen.",
		}
	case RejectAssetRevision:
		return &RejectionCode{
			Name:        "AssetRevision",
			Label:       "Asset Revision Incorrect",
			Description: "The revision in an asset amendment is incorrect.",
		}
	case RejectAssetNotPermitted:
		return &RejectionCode{
			Name:        "AssetNotPermitted",
			Label:       "Asset Not Permitted",
			Description: "Action not permitted by asset.",
		}
	case RejectTransferSelf:
		return &RejectionCode{
			Name:        "TransferSelf",
			Label:       "Transfer To Self Prohibited",
			Description: "Transfers with the sender and receiver addresses the same are not permitted.",
		}
	case RejectTransferExpired:
		return &RejectionCode{
			Name:        "TransferExpired",
			Label:       "Transfer Expired",
			Description: "The transfer has expired.",
		}
	case RejectHoldingsFrozen:
		return &RejectionCode{
			Name:        "HoldingsFrozen",
			Label:       "Holdings Frozen",
			Description: "Holdings are frozen, so the request can&#39;t be completed.",
		}
	case RejectHolderProposalProhibited:
		return &RejectionCode{
			Name:        "HolderProposalProhibited",
			Label:       "Holder Proposal Prohibited",
			Description: "Holders are not permitted to make proposals.",
		}
	case RejectProposalConflicts:
		return &RejectionCode{
			Name:        "ProposalConflicts",
			Label:       "Proposal Conflicts",
			Description: "The proposal conflicts with an unapplied proposal.",
		}
	case RejectVoteNotFound:
		return &RejectionCode{
			Name:        "VoteNotFound",
			Label:       "Vote Not Found",
			Description: "The vote ID referenced is not found.",
		}
	case RejectVoteClosed:
		return &RejectionCode{
			Name:        "VoteClosed",
			Label:       "Vote Closed",
			Description: "The vote has closed and ballots are no longer permitted.",
		}
	case RejectBallotAlreadyCounted:
		return &RejectionCode{
			Name:        "BallotAlreadyCounted",
			Label:       "Ballot Already Counted",
			Description: "The ballot has already been counted for this address.",
		}
	case RejectVoteSystemNotPermitted:
		return &RejectionCode{
			Name:        "VoteSystemNotPermitted",
			Label:       "Vote System Not Permitted",
			Description: "The voting system isn&#39;t permitted for this request.",
		}
	case RejectInsufficientTxFeeFunding:
		return &RejectionCode{
			Name:        "InsufficientTxFeeFunding",
			Label:       "Insufficient Transaction Fee Funding",
			Description: "Insufficient bitcoin quantities for response transaction fees.",
		}
	case RejectInsufficientValue:
		return &RejectionCode{
			Name:        "InsufficientValue",
			Label:       "Insufficient Value",
			Description: "Insufficient bitcoin quantity in inputs to fund request.",
		}
	case RejectInsufficientQuantity:
		return &RejectionCode{
			Name:        "InsufficientQuantity",
			Label:       "Insufficient Quantity",
			Description: "Insufficient token holdings to for request.",
		}
	case RejectNotAdministration:
		return &RejectionCode{
			Name:        "NotAdministration",
			Label:       "Requestor Is Not Administration",
			Description: "The requestor is not the administration and is required to be for this request.",
		}
	case RejectNotOperator:
		return &RejectionCode{
			Name:        "NotOperator",
			Label:       "Requestor Is Not Operator",
			Description: "The requestor is not the operator and is required to be for this request.",
		}
	case RejectUnauthorizedAddress:
		return &RejectionCode{
			Name:        "UnauthorizedAddress",
			Label:       "Unauthorized Address",
			Description: "The address specified is not permitted for this request.",
		}
	case RejectInvalidSignature:
		return &RejectionCode{
			Name:        "InvalidSignature",
			Label:       "Invalid Signature",
			Description: "The signature provided is not valid. This is for signatures included within OP_RETURN data. Not bitcoin transaction signature scripts.",
		}
	default:
		return nil
	}
}
