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
        'AssetID':                         [0, DAT_fixedchar, 32],
        'AssetAuthFlags':                  [1, DAT_bin, 8],
        'TransfersPermitted':              [2, DAT_bool, 1],
        'TradeRestrictions':               [3, DAT_fixedchar, 3],
        'EnforcementOrdersPermitted':      [4, DAT_bool, 1],
        'VoteMultiplier':                  [5, DAT_uint, 1],
        'ReferendumProposal':              [6, DAT_bool, 1],
        'InitiativeProposal':              [7, DAT_bool, 1],
        'AssetModificationGovernance':     [8, DAT_bool, 1],
        'TokenQty':                        [9, DAT_uint, 8],
        'ContractFeeCurrency':             [10, DAT_fixedchar, 3],
        'ContractFeeVar':                  [11, DAT_float, 4],
        'ContractFeeFixed':                [12, DAT_float, 4],
        'AssetPayload':                    [13, DAT_varbin, 16]
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
        self.ReferendumProposal = None
        self.InitiativeProposal = None
        self.AssetModificationGovernance = None
        self.TokenQty = None
        self.ContractFeeCurrency = None
        self.ContractFeeVar = None
        self.ContractFeeFixed = None
        self.AssetPayload = None


# Asset Creation Action - This action creates an Asset in response to the
# Issuer's instructions in the Definition Action.

class Action_AssetCreation(ActionBase):
    ActionPrefix = 'A2'

    schema = {
        'AssetID':                         [0, DAT_fixedchar, 32],
        'AssetAuthFlags':                  [1, DAT_bin, 8],
        'TransfersPermitted':              [2, DAT_bool, 1],
        'TradeRestrictions':               [3, DAT_fixedchar, 3],
        'EnforcementOrdersPermitted':      [4, DAT_bool, 1],
        'VoteMultiplier':                  [5, DAT_uint, 1],
        'ReferendumProposal':              [6, DAT_bool, 1],
        'InitiativeProposal':              [7, DAT_bool, 1],
        'AssetModificationGovernance':     [8, DAT_bool, 1],
        'TokenQty':                        [9, DAT_uint, 8],
        'ContractFeeCurrency':             [10, DAT_fixedchar, 3],
        'ContractFeeVar':                  [11, DAT_float, 4],
        'ContractFeeFixed':                [12, DAT_float, 4],
        'AssetPayload':                    [13, DAT_varbin, 16],
        'Asset Revision':                  [14, DAT_uint, 8],
        'Timestamp':                       [15, DAT_timestamp, 8]
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
        self.ReferendumProposal = None
        self.InitiativeProposal = None
        self.AssetModificationGovernance = None
        self.TokenQty = None
        self.ContractFeeCurrency = None
        self.ContractFeeVar = None
        self.ContractFeeFixed = None
        self.AssetPayload = None
        self.Asset Revision = None
        self.Timestamp = None


# Asset Modification Action - Token Dilutions, Call Backs/Revocations,
# burning etc.

class Action_AssetModification(ActionBase):
    ActionPrefix = 'A3'

    schema = {
        'AssetID':                         [0, DAT_fixedchar, 32],
        'AssetRevision':                   [1, DAT_uint, 8],
        'Modifications':                   [2, DAT_Amendment[], 0],
        'RefTxID':                         [3, DAT_sha256, 32]
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
        'ContractFileType':                [0, DAT_uint, 1],
        'ContractFile':                    [1, DAT_varbin, 32],
        'GoverningLaw':                    [2, DAT_fixedchar, 5],
        'Jurisdiction':                    [3, DAT_fixedchar, 5],
        'ContractExpiration':              [4, DAT_time, 8],
        'ContractURI':                     [5, DAT_varchar, 8],
        'IssuerName':                      [6, DAT_varchar, 8],
        'IssuerType':                      [7, DAT_fixedchar, 1],
        'IssuerLogoURL':                   [8, DAT_varchar, 8],
        'ContractOperatorID':              [9, DAT_varchar, 8],
        'ContractAuthFlags':               [10, DAT_bin, 16],
        'VotingSystems':                   [11, DAT_VotingSystem[], 0],
        'RestrictedQtyAssets':             [12, DAT_uint, 8],
        'ReferendumProposal':              [13, DAT_bool, 1],
        'InitiativeProposal':              [14, DAT_bool, 1],
        'Registries':                      [15, DAT_Registry[], 0],
        'IssuerAddress':                   [16, DAT_bool, 1],
        'UnitNumber':                      [17, DAT_varchar, 8],
        'BuildingNumber':                  [18, DAT_varchar, 8],
        'Street':                          [19, DAT_varchar, 16],
        'SuburbCity':                      [20, DAT_varchar, 8],
        'TerritoryStateProvinceCode':      [21, DAT_fixedchar, 5],
        'CountryCode':                     [22, DAT_fixedchar, 3],
        'PostalZIPCode':                   [23, DAT_varchar, 8],
        'EmailAddress':                    [24, DAT_varchar, 8],
        'PhoneNumber':                     [25, DAT_varchar, 8],
        'KeyRoles':                        [26, DAT_KeyRole[], 0],
        'NotableRoles':                    [27, DAT_NotableRole[], 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.ContractFile = None
        self.GoverningLaw = None
        self.Jurisdiction = None
        self.ContractExpiration = None
        self.ContractURI = None
        self.IssuerName = None
        self.IssuerType = None
        self.IssuerLogoURL = None
        self.ContractOperatorID = None
        self.ContractAuthFlags = None
        self.VotingSystems = None
        self.RestrictedQtyAssets = None
        self.ReferendumProposal = None
        self.InitiativeProposal = None
        self.Registries = None
        self.IssuerAddress = None
        self.UnitNumber = None
        self.BuildingNumber = None
        self.Street = None
        self.SuburbCity = None
        self.TerritoryStateProvinceCode = None
        self.CountryCode = None
        self.PostalZIPCode = None
        self.EmailAddress = None
        self.PhoneNumber = None
        self.KeyRoles = None
        self.NotableRoles = None


# This txn is created by the Contract (smart contract/off-chain agent/token
# contract) upon receipt of a valid Contract Offer Action from the issuer.
# The Smart Contract will execute on a server controlled by the Issuer. or
# a Smart Contract Operator on their behalf.

class Action_ContractFormation(ActionBase):
    ActionPrefix = 'C2'

    schema = {
        'ContractFileType':                [0, DAT_uint, 1],
        'ContractFile':                    [1, DAT_varbin, 32],
        'GoverningLaw':                    [2, DAT_fixedchar, 5],
        'Jurisdiction':                    [3, DAT_fixedchar, 5],
        'ContractExpiration':              [4, DAT_time, 8],
        'ContractURI':                     [5, DAT_varchar, 8],
        'IssuerName':                      [6, DAT_varchar, 8],
        'IssuerType':                      [7, DAT_fixedchar, 1],
        'IssuerLogoURL':                   [8, DAT_varchar, 8],
        'ContractOperatorID':              [9, DAT_varchar, 8],
        'ContractAuthFlags':               [10, DAT_bin, 16],
        'VotingSystems':                   [11, DAT_VotingSystem[], 0],
        'RestrictedQtyAssets':             [12, DAT_uint, 8],
        'ReferendumProposal':              [13, DAT_bool, 1],
        'InitiativeProposal':              [14, DAT_bool, 1],
        'Registries':                      [15, DAT_Registry[], 0],
        'IssuerAddress':                   [16, DAT_bool, 1],
        'UnitNumber':                      [17, DAT_varchar, 8],
        'BuildingNumber':                  [18, DAT_varchar, 8],
        'Street':                          [19, DAT_varchar, 16],
        'SuburbCity':                      [20, DAT_varchar, 8],
        'TerritoryStateProvinceCode':      [21, DAT_fixedchar, 5],
        'CountryCode':                     [22, DAT_fixedchar, 3],
        'PostalZIPCode':                   [23, DAT_varchar, 8],
        'EmailAddress':                    [24, DAT_varchar, 8],
        'PhoneNumber':                     [25, DAT_varchar, 8],
        'KeyRoles':                        [26, DAT_KeyRole[], 0],
        'NotableRoles':                    [27, DAT_NotableRole[], 0],
        'ContractRevision':                [28, DAT_uint, 8],
        'Timestamp':                       [29, DAT_timestamp, 8]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.ContractFile = None
        self.GoverningLaw = None
        self.Jurisdiction = None
        self.ContractExpiration = None
        self.ContractURI = None
        self.IssuerName = None
        self.IssuerType = None
        self.IssuerLogoURL = None
        self.ContractOperatorID = None
        self.ContractAuthFlags = None
        self.VotingSystems = None
        self.RestrictedQtyAssets = None
        self.ReferendumProposal = None
        self.InitiativeProposal = None
        self.Registries = None
        self.IssuerAddress = None
        self.UnitNumber = None
        self.BuildingNumber = None
        self.Street = None
        self.SuburbCity = None
        self.TerritoryStateProvinceCode = None
        self.CountryCode = None
        self.PostalZIPCode = None
        self.EmailAddress = None
        self.PhoneNumber = None
        self.KeyRoles = None
        self.NotableRoles = None
        self.ContractRevision = None
        self.Timestamp = None


# Contract Amendment Action - the issuer can initiate an amendment to the
# contract establishment metadata. The ability to make an amendment to the
# contract is restricted by the Authorization Flag set on the current
# revision of Contract Formation action.

class Action_ContractAmendment(ActionBase):
    ActionPrefix = 'C3'

    schema = {
        'ChangeOperatorAddress':           [0, DAT_bool, 1],
        'ContractRevision':                [1, DAT_uint, 2],
        'Amendments':                      [2, DAT_Amendment[], 0],
        'RefTxID':                         [3, DAT_SHA256, 32]
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
        'ContractType':                    [0, DAT_varchar, 8],
        'ContractFileType':                [1, DAT_uint, 1],
        'ContractFile':                    [2, DAT_varbin, 32],
        'ContractRevision':                [3, DAT_uint, 2],
        'GoverningLaw':                    [4, DAT_fixedchar, 5],
        'Jurisdiction':                    [5, DAT_fixedchar, 5],
        'EffectiveDate':                   [6, DAT_time, 8],
        'ContractExpiration':              [7, DAT_time, 8],
        'ContractURI':                     [8, DAT_varchar, 8],
        'PrevRevTxID':                     [9, DAT_sha256, 32],
        'Entities':                        [10, DAT_Entity[], 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.ContractFileType = None
        self.ContractFile = None
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
        'AssetID':                         [0, DAT_fixedchar, 32],
        'ComplianceAction':                [1, DAT_fixedchar, 1],
        'TargetAddresses':                 [2, DAT_TargetAddress[], 16],
        'DepositAddress':                  [3, DAT_Address, 0],
        'AuthorityName':                   [4, DAT_varchar, 8],
        'SigAlgoAddressList':              [5, DAT_uint, 1],
        'AuthorityPublicKey':              [6, DAT_varchar, 8],
        'OrderSignature':                  [7, DAT_varchar, 8],
        'SupportingEvidenceHash':          [8, DAT_sha256, 32],
        'RefTxnID':                        [9, DAT_sha256, 32],
        'FreezePeriod':                    [10, DAT_time, 8],
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
        self.SupportingEvidenceHash = None
        self.RefTxnID = None
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
        'Timestamp':                       [0, DAT_timestamp, 8]
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
        'RefTxnID':                        [0, DAT_sha256, 32],
        'Timestamp':                       [1, DAT_timestamp, 8]
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
        'Timestamp':                       [1, DAT_timestamp, 8]
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
        'Timestamp':                       [0, DAT_timestamp, 8]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):


# Initiative Action - Allows Token Owners to propose an Initiative (aka
# Initiative/Shareholder vote). A significant cost - specified in the
# Contract Formation - can be attached to this action to reduce spam, as
# the resulting vote will be put to all token owners.

class Action_Initiative(ActionBase):
    ActionPrefix = 'G1'

    schema = {
        'AssetID':                         [0, DAT_fixedchar, 32],
        'VoteSystem':                      [1, DAT_uint, 1],
        'Proposal':                        [2, DAT_bool, 1],
        'ProposedChanges':                 [3, DAT_Amendment[], 0],
        'VoteOptions':                     [4, DAT_varchar, 8],
        'VoteMax':                         [5, DAT_uint, 1],
        'ProposalDescription':             [6, DAT_varchar, 32],
        'ProposalDocumentHash':            [7, DAT_sha256, 32],
        'VoteCutOffTimestamp':             [8, DAT_time, 8]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.VoteSystem = None
        self.Proposal = None
        self.ProposedChanges = None
        self.VoteOptions = None
        self.VoteMax = None
        self.ProposalDescription = None
        self.ProposalDocumentHash = None
        self.VoteCutOffTimestamp = None


# Referendum Action - Issuer instructs the Contract to Initiate a Token
# Owner Vote. Usually used for contract amendments, organizational
# governance, etc.

class Action_Referendum(ActionBase):
    ActionPrefix = 'G2'

    schema = {
        'AssetType':                       [0, DAT_fixedchar, 3],
        'AssetID':                         [1, DAT_fixedchar, 32],
        'VoteSystem':                      [2, DAT_uint, 1],
        'Proposal':                        [3, DAT_bool, 1],
        'ProposedChanges':                 [4, DAT_Amendment[], 0],
        'VoteOptions':                     [5, DAT_varchar, 8],
        'VoteMax':                         [6, DAT_uint, 1],
        'ProposalDescription':             [7, DAT_varchar, 32],
        'ProposalDocumentHash':            [8, DAT_sha256, 32],
        'VoteCutOffTimestamp':             [9, DAT_time, 8]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.AssetID = None
        self.VoteSystem = None
        self.Proposal = None
        self.ProposedChanges = None
        self.VoteOptions = None
        self.VoteMax = None
        self.ProposalDescription = None
        self.ProposalDocumentHash = None
        self.VoteCutOffTimestamp = None


# Vote Action - A vote is created by the Contract in response to a valid
# Referendum (Issuer) or Initiative (User) Action.

class Action_Vote(ActionBase):
    ActionPrefix = 'G3'

    schema = {
        
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):


# Ballot Cast Action - Used by Token Owners to cast their ballot (vote) on
# proposals raised by the Issuer (Referendum) or other token holders
# (Initiative). 1 Vote per token unless a vote multiplier is specified in
# the relevant Asset Definition action.

class Action_BallotCast(ActionBase):
    ActionPrefix = 'G4'

    schema = {
        'AssetID':                         [0, DAT_fixedchar, 32],
        'VoteTxnID':                       [1, DAT_sha256, 32],
        'Vote':                            [2, DAT_varchar, 8]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.VoteTxnID = None
        self.Vote = None


# Ballot Counted Action - The smart contract will respond to a Ballot Cast
# action with a Ballot Counted action if the Ballot Cast is valid. If the
# Ballot Cast is not valid, then the smart contract will respond with a
# Rejection Action.

class Action_BallotCounted(ActionBase):
    ActionPrefix = 'G5'

    schema = {
        
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):


# Result Action - Once a vote has been completed the results are published.

class Action_Result(ActionBase):
    ActionPrefix = 'G6'

    schema = {
        'AssetID':                         [0, DAT_fixedchar, 32],
        'Proposal':                        [1, DAT_bool, 1],
        'ProposedChanges':                 [2, DAT_Amendment[], 0],
        'VoteTxnID':                       [3, DAT_sha256, 32],
        'VoteOptionsCount':                [4, DAT_uint, 1],
        'OptionXTally':                    [5, DAT_uint, 8],
        'Result':                          [6, DAT_varchar, 8],
        'Timestamp':                       [7, DAT_timestamp, 8]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.Proposal = None
        self.ProposedChanges = None
        self.VoteTxnID = None
        self.VoteOptionsCount = None
        self.OptionXTally = None
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
        'MessageType':                     [0, DAT_fixedchar, 2],
        'MessagePayload':                  [1, DAT_varchar, 32]
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
        'RejectionType':                   [0, DAT_uint, 1],
        'MessagePayload':                  [1, DAT_varchar, 32],
        'Timestamp':                       [2, DAT_timestamp, 8]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
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


# A Token Owner(s) Sends, Excahnges or Swaps a token(s) or Bitcoin for a
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
        'OfferExpiry':                     [0, DAT_time, 8],
        'ExchangeFeeCurrency':             [1, DAT_fixedchar, 3],
        'ExchangeFeeVar':                  [2, DAT_float, 4],
        'ExchangeFeeFixed':                [3, DAT_float, 4],
        'ExchangeFeeAddress':              [4, DAT_Address, 0]
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
        'Timestamp':                       [0, DAT_timestamp, 8]
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
    'G1': Action_Initiative,
    'G2': Action_Referendum,
    'G3': Action_Vote,
    'G4': Action_BallotCast,
    'G5': Action_BallotCounted,
    'G6': Action_Result,
    'M1': Action_Message,
    'M2': Action_Rejection,
    'R1': Action_Establishment,
    'R2': Action_Addition,
    'R3': Action_Alteration,
    'R4': Action_Removal,
    'T1': Action_Transfer,
    'T4': Action_Settlement
}
