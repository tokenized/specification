# Package protocol provides base level structs and validation for
# the protocol.
#
# The code in this file is auto-generated. Do not edit it by hand as it will
# be overwritten when code is regenerated.


# This action is used by the issuer to define the
# properties/characteristics of the Asset (token) that it wants to create.
# An asset has a unique identifier called AssetID = AssetType +
# base58(AssetCode + checksum). An asset is always linked to a Contract
# that is identified by the public address of the Contract wallet.

class Action_AssetDefinition(ActionBase):
    ActionPrefix = 'A1'

    schema = {
        'AssetAuthFlags':                  [0, DAT_varbin, 8],
        'TransfersPermitted':              [1, DAT_bool, 0],
        'TradeRestrictions':               [2, DAT_Polity[], 16],
        'EnforcementOrdersPermitted':      [3, DAT_bool, 0],
        'VotingRights':                    [4, DAT_bool, 0],
        'VoteMultiplier':                  [5, DAT_uint, 1],
        'IssuerProposal':                  [6, DAT_bool, 0],
        'HolderProposal':                  [7, DAT_bool, 0],
        'AssetModificationGovernance':     [8, DAT_uint, 1],
        'TokenQty':                        [9, DAT_uint, 8],
        'AssetPayload':                    [10, DAT_varbin, 16]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.TransfersPermitted = None
        self.TradeRestrictions = None
        self.EnforcementOrdersPermitted = None
        self.VotingRights = None
        self.VoteMultiplier = None
        self.IssuerProposal = None
        self.HolderProposal = None
        self.AssetModificationGovernance = None
        self.TokenQty = None
        self.AssetPayload = None


# This action creates an Asset in response to the Issuer's instructions in
# the Definition Action.

class Action_AssetCreation(ActionBase):
    ActionPrefix = 'A2'

    schema = {
        'AssetAuthFlags':                  [0, DAT_varbin, 8],
        'TransfersPermitted':              [1, DAT_bool, 0],
        'TradeRestrictions':               [2, DAT_Polity[], 16],
        'EnforcementOrdersPermitted':      [3, DAT_bool, 0],
        'VotingRights':                    [4, DAT_bool, 0],
        'VoteMultiplier':                  [5, DAT_uint, 1],
        'IssuerProposal':                  [6, DAT_bool, 0],
        'HolderProposal':                  [7, DAT_bool, 0],
        'AssetModificationGovernance':     [8, DAT_uint, 1],
        'TokenQty':                        [9, DAT_uint, 8],
        'AssetPayload':                    [10, DAT_varbin, 16],
        'AssetRevision':                   [11, DAT_uint, 4],
        'Timestamp':                       [12, DAT_Timestamp, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.TransfersPermitted = None
        self.TradeRestrictions = None
        self.EnforcementOrdersPermitted = None
        self.VotingRights = None
        self.VoteMultiplier = None
        self.IssuerProposal = None
        self.HolderProposal = None
        self.AssetModificationGovernance = None
        self.TokenQty = None
        self.AssetPayload = None
        self.AssetRevision = None
        self.Timestamp = None


# Token Dilutions, Call Backs/Revocations, burning etc.

class Action_AssetModification(ActionBase):
    ActionPrefix = 'A3'

    schema = {
        'AssetRevision':                   [0, DAT_uint, 4],
        'Amendments':                      [1, DAT_Amendment[], 0],
        'RefTxID':                         [2, DAT_TxId, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.Amendments = None
        self.RefTxID = None


# Allows the Issuer to tell the smart contract what they want the details
# (labels, data, T&C's, etc.) of the Contract to be on-chain in a public
# and immutable way. The Contract Offer action 'initializes' a generic
# smart contract that has been spun up by either the Smart Contract
# Operator or the Issuer. This on-chain action allows for the positive
# response from the smart contract with either a Contract Formation Action
# or a Rejection Action.

class Action_ContractOffer(ActionBase):
    ActionPrefix = 'C1'

    schema = {
        'BodyOfAgreement':                 [0, DAT_varbin, 32],
        'ContractType':                    [1, DAT_varchar, 8],
        'SupportingDocsFileType':          [2, DAT_uint, 1],
        'SupportingDocs':                  [3, DAT_varbin, 32],
        'GoverningLaw':                    [4, DAT_fixedchar, 5],
        'Jurisdiction':                    [5, DAT_fixedchar, 5],
        'ContractExpiration':              [6, DAT_Timestamp, 0],
        'ContractURI':                     [7, DAT_varchar, 8],
        'Issuer':                          [8, DAT_Entity, 0],
        'IssuerLogoURL':                   [9, DAT_varchar, 8],
        'ContractOperatorIncluded':        [10, DAT_bool, 0],
        'ContractOperator':                [11, DAT_Entity, 0],
        'ContractAuthFlags':               [12, DAT_varbin, 16],
        'ContractFee':                     [13, DAT_uint, 8],
        'VotingSystems':                   [14, DAT_VotingSystem[], 0],
        'RestrictedQtyAssets':             [15, DAT_uint, 8],
        'IssuerProposal':                  [16, DAT_bool, 0],
        'HolderProposal':                  [17, DAT_bool, 0],
        'Oracles':                         [18, DAT_Oracle[], 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.ContractType = None
        self.SupportingDocsFileType = None
        self.SupportingDocs = None
        self.GoverningLaw = None
        self.Jurisdiction = None
        self.ContractExpiration = None
        self.ContractURI = None
        self.Issuer = None
        self.IssuerLogoURL = None
        self.ContractOperatorIncluded = None
        self.ContractOperator = None
        self.ContractAuthFlags = None
        self.ContractFee = None
        self.VotingSystems = None
        self.RestrictedQtyAssets = None
        self.IssuerProposal = None
        self.HolderProposal = None
        self.Oracles = None


# This txn is created by the Contract (smart contract/off-chain agent/token
# contract) upon receipt of a valid Contract Offer Action from the issuer.
# The Smart Contract will execute on a server controlled by the Issuer. or
# a Smart Contract Operator on their behalf.

class Action_ContractFormation(ActionBase):
    ActionPrefix = 'C2'

    schema = {
        'BodyOfAgreement':                 [0, DAT_varbin, 32],
        'ContractType':                    [1, DAT_varchar, 8],
        'SupportingDocsFileType':          [2, DAT_uint, 1],
        'SupportingDocs':                  [3, DAT_varbin, 32],
        'GoverningLaw':                    [4, DAT_fixedchar, 5],
        'Jurisdiction':                    [5, DAT_fixedchar, 5],
        'ContractExpiration':              [6, DAT_Timestamp, 0],
        'ContractURI':                     [7, DAT_varchar, 8],
        'Issuer':                          [8, DAT_Entity, 0],
        'IssuerLogoURL':                   [9, DAT_varchar, 8],
        'ContractOperatorIncluded':        [10, DAT_bool, 0],
        'ContractOperator':                [11, DAT_Entity, 0],
        'ContractAuthFlags':               [12, DAT_varbin, 16],
        'ContractFee':                     [13, DAT_uint, 8],
        'VotingSystems':                   [14, DAT_VotingSystem[], 0],
        'RestrictedQtyAssets':             [15, DAT_uint, 8],
        'IssuerProposal':                  [16, DAT_bool, 0],
        'HolderProposal':                  [17, DAT_bool, 0],
        'Oracles':                         [18, DAT_Oracle[], 0],
        'ContractRevision':                [19, DAT_uint, 4],
        'Timestamp':                       [20, DAT_Timestamp, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.ContractType = None
        self.SupportingDocsFileType = None
        self.SupportingDocs = None
        self.GoverningLaw = None
        self.Jurisdiction = None
        self.ContractExpiration = None
        self.ContractURI = None
        self.Issuer = None
        self.IssuerLogoURL = None
        self.ContractOperatorIncluded = None
        self.ContractOperator = None
        self.ContractAuthFlags = None
        self.ContractFee = None
        self.VotingSystems = None
        self.RestrictedQtyAssets = None
        self.IssuerProposal = None
        self.HolderProposal = None
        self.Oracles = None
        self.ContractRevision = None
        self.Timestamp = None


# The issuer can initiate an amendment to the contract establishment
# metadata. The ability to make an amendment to the contract is restricted
# by the Authorization Flag set on the current revision of Contract
# Formation action.

class Action_ContractAmendment(ActionBase):
    ActionPrefix = 'C3'

    schema = {
        'ContractRevision':                [0, DAT_uint, 4],
        'Amendments':                      [1, DAT_Amendment[], 0],
        'RefTxID':                         [2, DAT_TxId, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.Amendments = None
        self.RefTxID = None


# Static Contract Formation Action

class Action_StaticContractFormation(ActionBase):
    ActionPrefix = 'C4'

    schema = {
        'BodyOfAgreementType':             [0, DAT_uint, 1],
        'BodyOfAgreement':                 [1, DAT_varbin, 32],
        'ContractType':                    [2, DAT_varchar, 8],
        'SupportingDocsFileType':          [3, DAT_uint, 1],
        'SupportingDocs':                  [4, DAT_varchar, 32],
        'ContractRevision':                [5, DAT_uint, 4],
        'GoverningLaw':                    [6, DAT_fixedchar, 5],
        'Jurisdiction':                    [7, DAT_fixedchar, 5],
        'EffectiveDate':                   [8, DAT_Timestamp, 0],
        'ContractExpiration':              [9, DAT_Timestamp, 0],
        'ContractURI':                     [10, DAT_varchar, 8],
        'PrevRevTxID':                     [11, DAT_TxId, 0],
        'Entities':                        [12, DAT_Entity[], 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.BodyOfAgreement = None
        self.ContractType = None
        self.SupportingDocsFileType = None
        self.SupportingDocs = None
        self.ContractRevision = None
        self.GoverningLaw = None
        self.Jurisdiction = None
        self.EffectiveDate = None
        self.ContractExpiration = None
        self.ContractURI = None
        self.PrevRevTxID = None
        self.Entities = None


# Used by the issuer to signal to the smart contract that the tokens that a
# particular public address(es) owns are to be confiscated, frozen, thawed
# or reconciled.

class Action_Order(ActionBase):
    ActionPrefix = 'E1'

    schema = {
        'AssetCode':                       [0, DAT_AssetCode, 0],
        'TargetAddresses':                 [1, DAT_TargetAddress[], 16],
        'FreezeTxId':                      [2, DAT_TxId, 0],
        'FreezePeriod':                    [3, DAT_Timestamp, 0],
        'DepositAddress':                  [4, DAT_PublicKeyHash, 0],
        'AuthorityIncluded':               [5, DAT_bool, 0],
        'AuthorityName':                   [6, DAT_varchar, 8],
        'AuthorityPublicKey':              [7, DAT_varbin, 8],
        'SignatureAlgorithm':              [8, DAT_uint, 1],
        'OrderSignature':                  [9, DAT_varbin, 8],
        'SupportingEvidenceHash':          [10, DAT_bin, 32],
        'RefTxs':                          [11, DAT_varbin, 32],
        'BitcoinDispersions':              [12, DAT_QuantityIndex[], 16],
        'Message':                         [13, DAT_varchar, 32]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.TargetAddresses = None
        self.FreezeTxId = None
        self.FreezePeriod = None
        self.DepositAddress = None
        self.AuthorityIncluded = None
        self.AuthorityName = None
        self.AuthorityPublicKey = None
        self.SignatureAlgorithm = None
        self.OrderSignature = None
        self.SupportingEvidenceHash = None
        self.RefTxs = None
        self.BitcoinDispersions = None
        self.Message = None


# The contract responding to an Order action to freeze assets. To be used
# to comply with contractual/legal/issuer requirements. The target public
# address(es) will be marked as frozen. However the Freeze action publishes
# this fact to the public blockchain for transparency. The Contract will
# not respond to any actions requested by the frozen address.

class Action_Freeze(ActionBase):
    ActionPrefix = 'E2'

    schema = {
        'Quantities':                      [0, DAT_QuantityIndex[], 16],
        'FreezePeriod':                    [1, DAT_Timestamp, 0],
        'Timestamp':                       [2, DAT_Timestamp, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.FreezePeriod = None
        self.Timestamp = None


# The contract responding to an Order action to thaw assets. To be used to
# comply with contractual obligations or legal requirements. The Alleged
# Offender's tokens will be unfrozen to allow them to resume normal
# exchange and governance activities.

class Action_Thaw(ActionBase):
    ActionPrefix = 'E3'

    schema = {
        
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):


# The contract responding to an Order action to confiscate assets. To be
# used to comply with contractual obligations, legal and/or issuer
# requirements.

class Action_Confiscation(ActionBase):
    ActionPrefix = 'E4'

    schema = {
        'Quantities':                      [0, DAT_QuantityIndex[], 16],
        'DepositQty':                      [1, DAT_uint, 8],
        'Timestamp':                       [2, DAT_Timestamp, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.DepositQty = None
        self.Timestamp = None


# The contract responding to an Order action to reconcile assets. To be
# used at the direction of the issuer to fix record keeping errors with
# bitcoin and token balances.

class Action_Reconciliation(ActionBase):
    ActionPrefix = 'E5'

    schema = {
        'Quantities':                      [0, DAT_QuantityIndex[], 16],
        'Timestamp':                       [1, DAT_Timestamp, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.Timestamp = None


# Allows Issuers/Token Holders to propose a change (aka
# Initiative/Shareholder vote). A significant cost - specified in the
# Contract Formation - can be attached to this action when sent from Token
# Holders to reduce spam, as the resulting vote will be put to all token
# owners.

class Action_Proposal(ActionBase):
    ActionPrefix = 'G1'

    schema = {
        'AssetType':                       [0, DAT_fixedchar, 3],
        'AssetCode':                       [1, DAT_AssetCode, 0],
        'VoteSystem':                      [2, DAT_uint, 1],
        'Specific':                        [3, DAT_bool, 0],
        'ProposedAmendments':              [4, DAT_Amendment[], 0],
        'VoteOptions':                     [5, DAT_varchar, 8],
        'VoteMax':                         [6, DAT_uint, 1],
        'ProposalDescription':             [7, DAT_varchar, 32],
        'ProposalDocumentHash':            [8, DAT_bin, 32],
        'VoteCutOffTimestamp':             [9, DAT_Timestamp, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.AssetCode = None
        self.VoteSystem = None
        self.Specific = None
        self.ProposedAmendments = None
        self.VoteOptions = None
        self.VoteMax = None
        self.ProposalDescription = None
        self.ProposalDocumentHash = None
        self.VoteCutOffTimestamp = None


# A vote is created by the Contract in response to a valid Proposal Action.

class Action_Vote(ActionBase):
    ActionPrefix = 'G2'

    schema = {
        
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):


# Used by Token Owners to cast their ballot (vote) on proposals. 1 Vote per
# token unless a vote multiplier is specified in the relevant Asset
# Definition action.

class Action_BallotCast(ActionBase):
    ActionPrefix = 'G3'

    schema = {
        
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):


# The smart contract will respond to a Ballot Cast action with a Ballot
# Counted action if the Ballot Cast is valid. If the Ballot Cast is not
# valid, then the smart contract will respond with a Rejection Action.

class Action_BallotCounted(ActionBase):
    ActionPrefix = 'G4'

    schema = {
        'Quantity':                        [0, DAT_uint, 8],
        'Timestamp':                       [1, DAT_Timestamp, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.Timestamp = None


# Once a vote has been completed the results are published. After the
# result is posted, it is up to the issuer to send a contract/asset
# amendement if appropriate.

class Action_Result(ActionBase):
    ActionPrefix = 'G5'

    schema = {
        'AssetCode':                       [0, DAT_AssetCode, 0],
        'Specific':                        [1, DAT_bool, 0],
        'ProposedAmendments':              [2, DAT_Amendment[], 0],
        'VoteTxId':                        [3, DAT_TxId, 0],
        'OptionTally':                     [4, DAT_uint64[], 8],
        'Result':                          [5, DAT_varchar, 8],
        'Timestamp':                       [6, DAT_Timestamp, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.Specific = None
        self.ProposedAmendments = None
        self.VoteTxId = None
        self.OptionTally = None
        self.Result = None
        self.Timestamp = None


# The message action is a general purpose communication action.
# 'Twitter/SMS' for Issuers/Investors/Users. The message txn can also be
# used for passing partially signed txns on-chain, establishing private
# communication channels and EDI (receipting, invoices, PO, and private
# offers/bids). The messages are broken down by type for easy filtering in
# the a user's wallet. The Message Types are listed in the Message Types
# table.

class Action_Message(ActionBase):
    ActionPrefix = 'M1'

    schema = {
        'MessagePayload':                  [0, DAT_varbin, 32]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):


# Used to reject request actions that do not comply with the Contract. If
# money is to be returned to a User then it is used in lieu of the
# Settlement Action to properly account for token balances. All Issuer/User
# request Actions must be responded to by the Contract with an Action. The
# only exception to this rule is when there is not enough fees in the first
# Action for the Contract response action to remain revenue neutral. If not
# enough fees are attached to pay for the Contract response then the
# Contract will not respond.

class Action_Rejection(ActionBase):
    ActionPrefix = 'M2'

    schema = {
        'RejectionCode':                   [0, DAT_RejectionCode, 0],
        'Message':                         [1, DAT_varchar, 16],
        'Timestamp':                       [2, DAT_Timestamp, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.Message = None
        self.Timestamp = None


# Establishes an on-chain register.

class Action_Establishment(ActionBase):
    ActionPrefix = 'R1'

    schema = {
        
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):


# Adds an entry to the Register.

class Action_Addition(ActionBase):
    ActionPrefix = 'R2'

    schema = {
        
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):


# A register entry/record can be altered.

class Action_Alteration(ActionBase):
    ActionPrefix = 'R3'

    schema = {
        
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):


# Removes an entry/record from the Register.

class Action_Removal(ActionBase):
    ActionPrefix = 'R4'

    schema = {
        
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):


# A Token Owner(s) Sends, Exchanges or Swaps a token(s) or Bitcoin for a
# token(s) or Bitcoin. Can be as simple as sending a single token to a
# receiver. Or can be as complex as many senders sending many different
# assets - controlled by many different smart contracts - to a number of
# receivers. This action also supports atomic swaps (tokens for tokens).
# Since many parties and contracts can be involved in a transfer and the
# corresponding settlement action, the partially signed T1 and T2 actions
# will need to be passed around on-chain with an M1 action, or off-chain.

class Action_Transfer(ActionBase):
    ActionPrefix = 'T1'

    schema = {
        'ExchangeFee':                     [0, DAT_uint, 8],
        'ExchangeFeeAddress':              [1, DAT_PublicKeyHash, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.ExchangeFeeAddress = None


# Settles the transfer request of bitcoins and tokens from transfer (T1)
# actions.

class Action_Settlement(ActionBase):
    ActionPrefix = 'T2'

    schema = {
        
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):



ActionClassMap = {
    'A1': Action_AssetDefinition,
    'A2': Action_AssetCreation,
    'A3': Action_AssetModification,
    'C1': Action_ContractOffer,
    'C2': Action_ContractFormation,
    'C3': Action_ContractAmendment,
    'C4': Action_StaticContractFormation,
    'E1': Action_Order,
    'E2': Action_Freeze,
    'E3': Action_Thaw,
    'E4': Action_Confiscation,
    'E5': Action_Reconciliation,
    'G1': Action_Proposal,
    'G2': Action_Vote,
    'G3': Action_BallotCast,
    'G4': Action_BallotCounted,
    'G5': Action_Result,
    'M1': Action_Message,
    'M2': Action_Rejection,
    'R1': Action_Establishment,
    'R2': Action_Addition,
    'R3': Action_Alteration,
    'R4': Action_Removal,
    'T1': Action_Transfer,
    'T2': Action_Settlement
}
