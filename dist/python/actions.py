# Package protocol provides base level structs and validation for
# the protocol.
#
# The code in this file is auto-generated. Do not edit it by hand as it will
# be overwritten when code is regenerated.


# Asset Definition Action - This action is used by the issuer to define the
# properties/characteristics of the Asset (token) that it wants to create.

class Action_AssetDefinition(ActionBase):
    ActionPrefix = 'A1'

    schema = {
        'AssetCode':                       [0, DAT_AssetCode, 0],
        'AssetAuthFlags':                  [1, DAT_varbin, 8],
        'TransfersPermitted':              [2, DAT_bool, 0],
        'TradeRestrictions':               [3, DAT_Polity, 0],
        'EnforcementOrdersPermitted':      [4, DAT_bool, 0],
        'VotingRights':                    [5, DAT_bool, 0],
        'VoteMultiplier':                  [6, DAT_uint, 1],
        'IssuerProposal':                  [7, DAT_bool, 0],
        'HolderProposal':                  [8, DAT_bool, 0],
        'AssetModificationGovernance':     [9, DAT_bool, 0],
        'TokenQty':                        [10, DAT_uint, 8],
        'AssetPayload':                    [11, DAT_varbin, 16]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.AssetAuthFlags = None
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


# Asset Creation Action - This action creates an Asset in response to the
# Issuer's instructions in the Definition Action.

class Action_AssetCreation(ActionBase):
    ActionPrefix = 'A2'

    schema = {
        'AssetCode':                       [0, DAT_AssetCode, 0],
        'AssetAuthFlags':                  [1, DAT_varbin, 8],
        'TransfersPermitted':              [2, DAT_bool, 0],
        'TradeRestrictions':               [3, DAT_Polity, 0],
        'EnforcementOrdersPermitted':      [4, DAT_bool, 0],
        'VoteMultiplier':                  [5, DAT_uint, 1],
        'IssuerProposal':                  [6, DAT_bool, 0],
        'HolderProposal':                  [7, DAT_bool, 0],
        'AssetModificationGovernance':     [8, DAT_bool, 0],
        'TokenQty':                        [9, DAT_uint, 8],
        'AssetPayload':                    [10, DAT_varbin, 16],
        'Asset Revision':                  [11, DAT_uint, 4],
        'Timestamp':                       [12, DAT_Timestamp, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.AssetAuthFlags = None
        self.TransfersPermitted = None
        self.TradeRestrictions = None
        self.EnforcementOrdersPermitted = None
        self.VoteMultiplier = None
        self.IssuerProposal = None
        self.HolderProposal = None
        self.AssetModificationGovernance = None
        self.TokenQty = None
        self.AssetPayload = None
        self.Asset Revision = None
        self.Timestamp = None


# Asset Modification Action - Token Dilutions, Call Backs/Revocations,
# burning etc.

class Action_AssetModification(ActionBase):
    ActionPrefix = 'A3'

    schema = {
        'AssetCode':                       [0, DAT_AssetCode, 0],
        'AssetRevision':                   [1, DAT_uint, 4],
        'Modifications':                   [2, DAT_Amendment[], 0],
        'RefTxID':                         [3, DAT_TxId, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.AssetRevision = None
        self.Modifications = None
        self.RefTxID = None


# The Contract Offer action allows the Issuer to tell the smart contract
# what they want the details (labels, data, T&C's, etc.) of the Contract to
# be on-chain in a public and immutable way. The Contract Offer action
# 'initializes' a generic smart contract that has been spun up by either
# the Smart Contract Operator or the Issuer. This on-chain action allows
# for the positive response from the smart contract with either a Contract
# Formation Action or a Rejection Action.

class Action_ContractOffer(ActionBase):
    ActionPrefix = 'C1'

    schema = {
        'BodyOfAgreementType':             [0, DAT_uint, 1],
        'BodyOfAgreement':                 [1, DAT_varbin, 32],
        'ContractType':                    [2, DAT_varchar, 8],
        'SupportingDocsFileType':          [3, DAT_uint, 1],
        'SupportingDocs':                  [4, DAT_varbin, 32],
        'GoverningLaw':                    [5, DAT_fixedchar, 5],
        'Jurisdiction':                    [6, DAT_fixedchar, 5],
        'ContractExpiration':              [7, DAT_Timestamp, 0],
        'ContractURI':                     [8, DAT_varchar, 8],
        'Issuer':                          [9, DAT_Entity, 0],
        'IssuerLogoURL':                   [10, DAT_varchar, 8],
        'ContractOperatorIncluded':        [11, DAT_bool, 0],
        'ContractOperator':                [12, DAT_Entity, 0],
        'ContractAuthFlags':               [13, DAT_varbin, 16],
        'ContractFee':                     [14, DAT_uint, 8],
        'VotingSystems':                   [15, DAT_VotingSystem[], 0],
        'RestrictedQtyAssets':             [16, DAT_uint, 8],
        'IssuerProposal':                  [17, DAT_bool, 0],
        'HolderProposal':                  [18, DAT_bool, 0],
        'Registries':                      [19, DAT_Registry[], 0]
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
        self.Registries = None


# This txn is created by the Contract (smart contract/off-chain agent/token
# contract) upon receipt of a valid Contract Offer Action from the issuer.
# The Smart Contract will execute on a server controlled by the Issuer. or
# a Smart Contract Operator on their behalf.

class Action_ContractFormation(ActionBase):
    ActionPrefix = 'C2'

    schema = {
        'BodyOfAgreementType':             [0, DAT_uint, 1],
        'BodyOfAgreement':                 [1, DAT_varbin, 32],
        'ContractType':                    [2, DAT_varchar, 8],
        'SupportingDocsFileType':          [3, DAT_uint, 1],
        'SupportingDocs':                  [4, DAT_varbin, 32],
        'GoverningLaw':                    [5, DAT_fixedchar, 5],
        'Jurisdiction':                    [6, DAT_fixedchar, 5],
        'ContractExpiration':              [7, DAT_Timestamp, 0],
        'ContractURI':                     [8, DAT_varchar, 8],
        'Issuer':                          [9, DAT_Entity, 0],
        'IssuerLogoURL':                   [10, DAT_varchar, 8],
        'ContractOperatorIncluded':        [11, DAT_bool, 0],
        'ContractOperator':                [12, DAT_Entity, 0],
        'ContractAuthFlags':               [13, DAT_varbin, 16],
        'ContractFee':                     [14, DAT_uint, 8],
        'VotingSystems':                   [15, DAT_VotingSystem[], 0],
        'RestrictedQtyAssets':             [16, DAT_uint, 8],
        'IssuerProposal':                  [17, DAT_bool, 0],
        'HolderProposal':                  [18, DAT_bool, 0],
        'Registries':                      [19, DAT_Registry[], 0],
        'ContractRevision':                [20, DAT_uint, 4],
        'Timestamp':                       [21, DAT_Timestamp, 0]
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
        self.Registries = None
        self.ContractRevision = None
        self.Timestamp = None


# Contract Amendment Action - the issuer can initiate an amendment to the
# contract establishment metadata. The ability to make an amendment to the
# contract is restricted by the Authorization Flag set on the current
# revision of Contract Formation action.

class Action_ContractAmendment(ActionBase):
    ActionPrefix = 'C3'

    schema = {
        'ChangeOperatorAddress':           [0, DAT_bool, 0],
        'ContractRevision':                [1, DAT_uint, 4],
        'Amendments':                      [2, DAT_Amendment[], 0],
        'RefTxID':                         [3, DAT_TxId, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.ContractRevision = None
        self.Amendments = None
        self.RefTxID = None


# Static Contract Formation Action

class Action_StaticContractFormation(ActionBase):
    ActionPrefix = 'C4'

    schema = {
        'ContractCode':                    [0, DAT_ContractCode, 0],
        'BodyOfAgreementType':             [1, DAT_uint, 1],
        'BodyOfAgreement':                 [2, DAT_varbin, 32],
        'ContractType':                    [3, DAT_varchar, 8],
        'SupportingDocsFileType':          [4, DAT_uint, 1],
        'SupportingDocs':                  [5, DAT_varchar, 32],
        'ContractRevision':                [6, DAT_uint, 4],
        'GoverningLaw':                    [7, DAT_fixedchar, 5],
        'Jurisdiction':                    [8, DAT_fixedchar, 5],
        'EffectiveDate':                   [9, DAT_Timestamp, 0],
        'ContractExpiration':              [10, DAT_Timestamp, 0],
        'ContractURI':                     [11, DAT_varchar, 8],
        'PrevRevTxID':                     [12, DAT_TxId, 0],
        'Entities':                        [13, DAT_Entity[], 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.BodyOfAgreementType = None
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


# Order Action - Issuer to signal to the smart contract that the tokens
# that a particular public address(es) owns are to be confiscated, frozen,
# thawed or reconciled.

class Action_Order(ActionBase):
    ActionPrefix = 'E1'

    schema = {
        'AssetCode':                       [0, DAT_AssetCode, 0],
        'ComplianceAction':                [1, DAT_fixedchar, 1],
        'TargetAddresses':                 [2, DAT_TargetAddress[], 16],
        'DepositAddress':                  [3, DAT_PublicKeyHash, 0],
        'AuthorityName':                   [4, DAT_varchar, 8],
        'SigAlgoAddressList':              [5, DAT_uint, 1],
        'AuthorityPublicKey':              [6, DAT_varchar, 8],
        'OrderSignature':                  [7, DAT_varchar, 8],
        'SupportingEvidenceTxId':          [8, DAT_TxId, 0],
        'RefTxns':                         [9, DAT_TxId, 0],
        'FreezePeriod':                    [10, DAT_Timestamp, 0],
        'Message':                         [11, DAT_varchar, 32]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.ComplianceAction = None
        self.TargetAddresses = None
        self.DepositAddress = None
        self.AuthorityName = None
        self.SigAlgoAddressList = None
        self.AuthorityPublicKey = None
        self.OrderSignature = None
        self.SupportingEvidenceTxId = None
        self.RefTxns = None
        self.FreezePeriod = None
        self.Message = None


# Freeze Action - To be used to comply with contractual/legal/issuer
# requirements. The target public address(es) will be marked as frozen.
# However the Freeze action publishes this fact to the public blockchain
# for transparency. The Contract will not respond to any actions requested
# by the frozen address.

class Action_Freeze(ActionBase):
    ActionPrefix = 'E2'

    schema = {
        'Timestamp':                       [0, DAT_Timestamp, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):


# Thaw Action - to be used to comply with contractual obligations or legal
# requirements. The Alleged Offender's tokens will be unfrozen to allow
# them to resume normal exchange and governance activities.

class Action_Thaw(ActionBase):
    ActionPrefix = 'E3'

    schema = {
        'RefTxID':                         [0, DAT_TxId, 0],
        'Timestamp':                       [1, DAT_Timestamp, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.Timestamp = None


# Confiscation Action - to be used to comply with contractual obligations,
# legal and/or issuer requirements.

class Action_Confiscation(ActionBase):
    ActionPrefix = 'E4'

    schema = {
        'DepositQty':                      [0, DAT_uint, 8],
        'Timestamp':                       [1, DAT_Timestamp, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.Timestamp = None


# Reconciliation Action - to be used at the direction of the issuer to fix
# record keeping errors with bitcoin and token balances.

class Action_Reconciliation(ActionBase):
    ActionPrefix = 'E5'

    schema = {
        'Timestamp':                       [0, DAT_Timestamp, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):


# Proposal Action - Allows Issuers/Token Holders to propose a change (aka
# Initiative/Shareholder vote). A significant cost - specified in the
# Contract Formation - can be attached to this action when sent from Token
# Holders to reduce spam, as the resulting vote will be put to all token
# owners.

class Action_Proposal(ActionBase):
    ActionPrefix = 'G1'

    schema = {
        'AssetSpecificVote':               [0, DAT_bool, 0],
        'AssetType':                       [1, DAT_fixedchar, 3],
        'AssetCode':                       [2, DAT_AssetCode, 0],
        'VoteSystem':                      [3, DAT_uint, 1],
        'Specific':                        [4, DAT_bool, 0],
        'ProposedAmendments':              [5, DAT_Amendment[], 0],
        'VoteOptions':                     [6, DAT_varchar, 8],
        'VoteMax':                         [7, DAT_uint, 1],
        'ProposalDescription':             [8, DAT_varchar, 32],
        'ProposalDocumentHash':            [9, DAT_bin, 32],
        'VoteCutOffTimestamp':             [10, DAT_Timestamp, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.AssetType = None
        self.AssetCode = None
        self.VoteSystem = None
        self.Specific = None
        self.ProposedAmendments = None
        self.VoteOptions = None
        self.VoteMax = None
        self.ProposalDescription = None
        self.ProposalDocumentHash = None
        self.VoteCutOffTimestamp = None


# Vote Action - A vote is created by the Contract in response to a valid
# Proposal Action.

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


# Ballot Cast Action - Used by Token Owners to cast their ballot (vote) on
# proposals. 1 Vote per token unless a vote multiplier is specified in the
# relevant Asset Definition action.

class Action_BallotCast(ActionBase):
    ActionPrefix = 'G3'

    schema = {
        'Vote':                            [0, DAT_varchar, 8]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):


# Ballot Counted Action - The smart contract will respond to a Ballot Cast
# action with a Ballot Counted action if the Ballot Cast is valid. If the
# Ballot Cast is not valid, then the smart contract will respond with a
# Rejection Action.

class Action_BallotCounted(ActionBase):
    ActionPrefix = 'G4'

    schema = {
        'Timestamp':                       [0, DAT_Timestamp, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):


# Result Action - Once a vote has been completed the results are published.
# After the result is posted, it is up to the issuer to send a
# contract/asset amendement if appropriate.

class Action_Result(ActionBase):
    ActionPrefix = 'G5'

    schema = {
        'AssetType':                       [0, DAT_fixedchar, 3],
        'AssetCode':                       [1, DAT_AssetCode, 0],
        'Specific':                        [2, DAT_bool, 0],
        'ProposedAmendments':              [3, DAT_Amendment[], 0],
        'VoteTxID':                        [4, DAT_TxId, 0],
        'OptionTally':                     [5, DAT_uint64[], 8],
        'Result':                          [6, DAT_varchar, 8],
        'Timestamp':                       [7, DAT_Timestamp, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.AssetCode = None
        self.Specific = None
        self.ProposedAmendments = None
        self.VoteTxID = None
        self.OptionTally = None
        self.Result = None
        self.Timestamp = None


# Message Action - the message action is a general purpose communication
# action. 'Twitter/SMS' for Issuers/Investors/Users. The message txn can
# also be used for passing partially signed txns on-chain, establishing
# private communication channels and EDI (receipting, invoices, PO, and
# private offers/bids). The messages are broken down by type for easy
# filtering in the a user's wallet. The Message Types are listed in the
# Message Types table.

class Action_Message(ActionBase):
    ActionPrefix = 'M1'

    schema = {
        'MessageType':                     [0, DAT_fixedchar, 4],
        'MessagePayload':                  [1, DAT_varbin, 32]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.MessagePayload = None


# Rejection Action - used to reject request actions that do not comply with
# the Contract. If money is to be returned to a User then it is used in
# lieu of the Settlement Action to properly account for token balances. All
# Issuer/User request Actions must be responded to by the Contract with an
# Action. The only exception to this rule is when there is not enough fees
# in the first Action for the Contract response action to remain revenue
# neutral. If not enough fees are attached to pay for the Contract response
# then the Contract will not respond.

class Action_Rejection(ActionBase):
    ActionPrefix = 'M2'

    schema = {
        'AddressIndexes':                  [0, DAT_uint16[], 0],
        'RejectionType':                   [1, DAT_uint, 1],
        'MessagePayload':                  [2, DAT_varchar, 32],
        'Timestamp':                       [3, DAT_Timestamp, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.RejectionType = None
        self.MessagePayload = None
        self.Timestamp = None


# Establishment Action - Establishes an on-chain register.

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


# Addition Action - Adds an entry to the Register.

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


# Alteration Action - A register entry/record can be altered.

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


# Removal Action - Removes an entry/record from the Register.

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
        'OfferExpiry':                     [0, DAT_Timestamp, 0],
        'ExchangeFeeCurrency':             [1, DAT_fixedchar, 3],
        'ExchangeFeeVar':                  [2, DAT_float, 4],
        'ExchangeFeeFixed':                [3, DAT_float, 4],
        'ExchangeFeeAddress':              [4, DAT_PublicKeyHash, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.ExchangeFeeCurrency = None
        self.ExchangeFeeVar = None
        self.ExchangeFeeFixed = None
        self.ExchangeFeeAddress = None


# Settlement Action - Settles the transfer request of bitcoins and tokens
# from transfer (T1) actions.

class Action_Settlement(ActionBase):
    ActionPrefix = 'T4'

    schema = {
        'Timestamp':                       [0, DAT_Timestamp, 0]
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
    'T4': Action_Settlement
}
