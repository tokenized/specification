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
        'AssetType':                       [0, DAT_string, 3],
        'AssetID':                         [1, DAT_string, 32],
        'AssetAuthFlags':                  [2, DAT_bin, 8],
        'TransfersPermitted':              [3, DAT_bool, 1],
        'TradeRestrictions':               [4, DAT_string, 3],
        'EnforcementOrdersPermitted':      [5, DAT_bool, 1],
        'VoteMultiplier':                  [6, DAT_uint8, 1],
        'ReferendumProposal':              [7, DAT_bool, 1],
        'InitiativeProposal':              [8, DAT_bool, 1],
        'AssetModificationGovernance':     [9, DAT_bool, 1],
        'TokenQty':                        [10, DAT_uint64, 8],
        'ContractFeeCurrency':             [11, DAT_string, 3],
        'ContractFeeVar':                  [12, DAT_float32, 4],
        'ContractFeeFixed':                [13, DAT_float32, 4],
        'AssetPayloadLen':                 [14, DAT_uint16, 2],
        'AssetPayload':                    [15, DAT_byte[], 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.AssetID = None
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
        self.AssetPayloadLen = None
        self.AssetPayload = None


# Asset Creation Action - This action creates an Asset in response to the
# Issuer's instructions in the Definition Action.

class Action_AssetCreation(ActionBase):
    ActionPrefix = 'A2'

    schema = {
        'AssetType':                       [0, DAT_string, 3],
        'AssetID':                         [1, DAT_string, 32],
        'AssetAuthFlags':                  [2, DAT_bin, 8],
        'TransfersPermitted':              [3, DAT_bool, 1],
        'TradeRestrictions':               [4, DAT_string, 3],
        'EnforcementOrdersPermitted':      [5, DAT_bool, 1],
        'VoteMultiplier':                  [6, DAT_uint8, 1],
        'ReferendumProposal':              [7, DAT_bool, 1],
        'InitiativeProposal':              [8, DAT_bool, 1],
        'AssetModificationGovernance':     [9, DAT_bool, 1],
        'TokenQty':                        [10, DAT_uint64, 8],
        'ContractFeeCurrency':             [11, DAT_string, 3],
        'ContractFeeVar':                  [12, DAT_float32, 4],
        'ContractFeeFixed':                [13, DAT_float32, 4],
        'AssetPayloadLen':                 [14, DAT_uint16, 2],
        'AssetPayload':                    [15, DAT_byte[], 0],
        'Asset Revision':                  [16, DAT_uint64, 8],
        'Timestamp':                       [17, DAT_timestamp, 8]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.AssetID = None
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
        self.AssetPayloadLen = None
        self.AssetPayload = None
        self.Asset Revision = None
        self.Timestamp = None


# Asset Modification Action - Token Dilutions, Call Backs/Revocations,
# burning etc.

class Action_AssetModification(ActionBase):
    ActionPrefix = 'A3'

    schema = {
        'AssetType':                       [0, DAT_string, 3],
        'AssetID':                         [1, DAT_string, 32],
        'AssetRevision':                   [2, DAT_uint64, 8],
        'ModificationCount':               [3, DAT_uint8, 1],
        'Modifications':                   [4, DAT_Amendment[], 0],
        'RefTxID':                         [5, DAT_sha256, 32]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.AssetID = None
        self.AssetRevision = None
        self.ModificationCount = None
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
        'ContractName':                    [0, DAT_nvarchar8, 0],
        'ContractFileType':                [1, DAT_uint8, 1],
        'LenContractFile':                 [2, DAT_uint32, 4],
        'ContractFile':                    [3, DAT_string, 32],
        'GoverningLaw':                    [4, DAT_string, 5],
        'Jurisdiction':                    [5, DAT_string, 5],
        'ContractExpiration':              [6, DAT_time, 8],
        'ContractURI':                     [7, DAT_nvarchar8, 0],
        'IssuerName':                      [8, DAT_nvarchar8, 0],
        'IssuerType':                      [9, DAT_string, 1],
        'IssuerLogoURL':                   [10, DAT_nvarchar8, 0],
        'ContractOperatorID':              [11, DAT_nvarchar8, 0],
        'ContractAuthFlags':               [12, DAT_bin, 16],
        'VotingSystemCount':               [13, DAT_uint8, 1],
        'VotingSystems':                   [14, DAT_VotingSystem[], 0],
        'RestrictedQtyAssets':             [15, DAT_uint64, 8],
        'ReferendumProposal':              [16, DAT_bool, 1],
        'InitiativeProposal':              [17, DAT_bool, 1],
        'RegistryCount':                   [18, DAT_uint8, 1],
        'Registries':                      [19, DAT_Registry[], 0],
        'IssuerAddress':                   [20, DAT_bool, 1],
        'UnitNumber':                      [21, DAT_nvarchar8, 0],
        'BuildingNumber':                  [22, DAT_nvarchar8, 0],
        'Street':                          [23, DAT_nvarchar16, 0],
        'SuburbCity':                      [24, DAT_nvarchar8, 0],
        'TerritoryStateProvinceCode':      [25, DAT_string, 5],
        'CountryCode':                     [26, DAT_string, 3],
        'PostalZIPCode':                   [27, DAT_nvarchar8, 0],
        'EmailAddress':                    [28, DAT_nvarchar8, 0],
        'PhoneNumber':                     [29, DAT_nvarchar8, 0],
        'KeyRolesCount':                   [30, DAT_uint8, 1],
        'KeyRoles':                        [31, DAT_KeyRole[], 0],
        'NotableRolesCount':               [32, DAT_uint8, 1],
        'NotableRoles':                    [33, DAT_NotableRole[], 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.ContractFileType = None
        self.LenContractFile = None
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
        self.VotingSystemCount = None
        self.VotingSystems = None
        self.RestrictedQtyAssets = None
        self.ReferendumProposal = None
        self.InitiativeProposal = None
        self.RegistryCount = None
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
        self.KeyRolesCount = None
        self.KeyRoles = None
        self.NotableRolesCount = None
        self.NotableRoles = None


# This txn is created by the Contract (smart contract/off-chain agent/token
# contract) upon receipt of a valid Contract Offer Action from the issuer.
# The Smart Contract will execute on a server controlled by the Issuer. or
# a Smart Contract Operator on their behalf.

class Action_ContractFormation(ActionBase):
    ActionPrefix = 'C2'

    schema = {
        'ContractName':                    [0, DAT_nvarchar8, 0],
        'ContractFileType':                [1, DAT_uint8, 1],
        'LenContractFile':                 [2, DAT_uint32, 4],
        'ContractFile':                    [3, DAT_string, 32],
        'GoverningLaw':                    [4, DAT_string, 5],
        'Jurisdiction':                    [5, DAT_string, 5],
        'ContractExpiration':              [6, DAT_time, 8],
        'ContractURI':                     [7, DAT_nvarchar8, 0],
        'IssuerName':                      [8, DAT_nvarchar8, 0],
        'IssuerType':                      [9, DAT_string, 1],
        'IssuerLogoURL':                   [10, DAT_nvarchar8, 0],
        'ContractOperatorID':              [11, DAT_nvarchar8, 0],
        'ContractAuthFlags':               [12, DAT_bin, 16],
        'VotingSystemCount':               [13, DAT_uint8, 1],
        'VotingSystems':                   [14, DAT_VotingSystem[], 0],
        'RestrictedQtyAssets':             [15, DAT_uint64, 8],
        'ReferendumProposal':              [16, DAT_bool, 1],
        'InitiativeProposal':              [17, DAT_bool, 1],
        'RegistryCount':                   [18, DAT_uint8, 1],
        'Registries':                      [19, DAT_Registry[], 0],
        'IssuerAddress':                   [20, DAT_bool, 1],
        'UnitNumber':                      [21, DAT_nvarchar8, 0],
        'BuildingNumber':                  [22, DAT_nvarchar8, 0],
        'Street':                          [23, DAT_nvarchar16, 0],
        'SuburbCity':                      [24, DAT_nvarchar8, 0],
        'TerritoryStateProvinceCode':      [25, DAT_string, 5],
        'CountryCode':                     [26, DAT_string, 3],
        'PostalZIPCode':                   [27, DAT_nvarchar8, 0],
        'EmailAddress':                    [28, DAT_nvarchar8, 0],
        'PhoneNumber':                     [29, DAT_nvarchar8, 0],
        'KeyRolesCount':                   [30, DAT_uint8, 1],
        'KeyRoles':                        [31, DAT_KeyRole[], 0],
        'NotableRolesCount':               [32, DAT_uint8, 1],
        'NotableRoles':                    [33, DAT_NotableRole[], 0],
        'ContractRevision':                [34, DAT_uint64, 8],
        'Timestamp':                       [35, DAT_timestamp, 8]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.ContractFileType = None
        self.LenContractFile = None
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
        self.VotingSystemCount = None
        self.VotingSystems = None
        self.RestrictedQtyAssets = None
        self.ReferendumProposal = None
        self.InitiativeProposal = None
        self.RegistryCount = None
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
        self.KeyRolesCount = None
        self.KeyRoles = None
        self.NotableRolesCount = None
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
        'ChangeIssuerAddress':             [0, DAT_bool, 1],
        'ChangeOperatorAddress':           [1, DAT_bool, 1],
        'ContractRevision':                [2, DAT_uint16, 2],
        'AmendmentsCount':                 [3, DAT_uint8, 1],
        'Amendments':                      [4, DAT_Amendment[], 0],
        'RefTxID':                         [5, DAT_SHA256, 32]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.ChangeOperatorAddress = None
        self.ContractRevision = None
        self.AmendmentsCount = None
        self.Amendments = None
        self.RefTxID = None


# Static Contract Formation Action

class Action_StaticContractFormation(ActionBase):
    ActionPrefix = 'C4'

    schema = {
        'ContractName':                    [0, DAT_nvarchar8, 30],
        'ContractType':                    [1, DAT_nvarchar8, 30],
        'ContractFileType':                [2, DAT_uint8, 1],
        'LenContractFile':                 [3, DAT_uint32, 4],
        'ContractFile':                    [4, DAT_string, 32],
        'ContractRevision':                [5, DAT_uint16, 2],
        'GoverningLaw':                    [6, DAT_string, 5],
        'Jurisdiction':                    [7, DAT_string, 5],
        'EffectiveDate':                   [8, DAT_time, 8],
        'ContractExpiration':              [9, DAT_time, 8],
        'ContractURI':                     [10, DAT_nvarchar8, 53],
        'PrevRevTxID':                     [11, DAT_string, 32],
        'EntityCount':                     [12, DAT_uint8, 1],
        'Entities':                        [13, DAT_Entity[], 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.ContractType = None
        self.ContractFileType = None
        self.LenContractFile = None
        self.ContractFile = None
        self.ContractRevision = None
        self.GoverningLaw = None
        self.Jurisdiction = None
        self.EffectiveDate = None
        self.ContractExpiration = None
        self.ContractURI = None
        self.PrevRevTxID = None
        self.EntityCount = None
        self.Entities = None


# Order Action - Issuer to signal to the smart contract that the tokens
# that a particular public address(es) owns are to be confiscated, frozen,
# thawed or reconciled.

class Action_Order(ActionBase):
    ActionPrefix = 'E1'

    schema = {
        'AssetType':                       [0, DAT_string, 3],
        'AssetID':                         [1, DAT_string, 32],
        'ComplianceAction':                [2, DAT_string, 1],
        'TargetAddressCount':              [3, DAT_uint16, 2],
        'TargetAddresses':                 [4, DAT_TargetAddress[], 0],
        'DepositAddress':                  [5, DAT_nvarchar8, 0],
        'AuthorityName':                   [6, DAT_nvarchar8, 0],
        'SigAlgoAddressList':              [7, DAT_uint8, 1],
        'AuthorityPublicKey':              [8, DAT_nvarchar8, 0],
        'OrderSignature':                  [9, DAT_nvarchar8, 0],
        'SupportingEvidenceHash':          [10, DAT_sha256, 32],
        'RefTxnID':                        [11, DAT_sha256, 32],
        'FreezePeriod':                    [12, DAT_time, 8],
        'Message':                         [13, DAT_nvarchar64, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.AssetID = None
        self.ComplianceAction = None
        self.TargetAddressCount = None
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
        'Addresses':                       [0, DAT_Address[], 0],
        'Timestamp':                       [1, DAT_timestamp, 8]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.Timestamp = None


# Thaw Action - to be used to comply with contractual obligations or legal
# requirements. The Alleged Offender's tokens will be unfrozen to allow
# them to resume normal exchange and governance activities.

class Action_Thaw(ActionBase):
    ActionPrefix = 'E3'

    schema = {
        'Addresses':                       [0, DAT_Address[], 0],
        'RefTxnID':                        [1, DAT_sha256, 32],
        'Timestamp':                       [2, DAT_timestamp, 8]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.RefTxnID = None
        self.Timestamp = None


# Confiscation Action - to be used to comply with contractual obligations,
# legal and/or issuer requirements.

class Action_Confiscation(ActionBase):
    ActionPrefix = 'E4'

    schema = {
        'Addresses':                       [0, DAT_Address[], 0],
        'DepositQty':                      [1, DAT_uint64, 8],
        'Timestamp':                       [2, DAT_timestamp, 8]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.DepositQty = None
        self.Timestamp = None


# Reconciliation Action - to be used at the direction of the issuer to fix
# record keeping errors with bitcoin and token balances.

class Action_Reconciliation(ActionBase):
    ActionPrefix = 'E5'

    schema = {
        'Addresses':                       [0, DAT_Address[], 0],
        'Timestamp':                       [1, DAT_timestamp, 8]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.Timestamp = None


# Initiative Action - Allows Token Owners to propose an Initiative (aka
# Initiative/Shareholder vote). A significant cost - specified in the
# Contract Formation - can be attached to this action to reduce spam, as
# the resulting vote will be put to all token owners.

class Action_Initiative(ActionBase):
    ActionPrefix = 'G1'

    schema = {
        'AssetType':                       [0, DAT_string, 3],
        'AssetID':                         [1, DAT_string, 32],
        'VoteSystem':                      [2, DAT_uint8, 1],
        'Proposal':                        [3, DAT_bool, 1],
        'ProposedChangesCount':            [4, DAT_uint8, 1],
        'ProposedChanges':                 [5, DAT_Amendment[], 0],
        'VoteOptions':                     [6, DAT_nvarchar8, 0],
        'VoteMax':                         [7, DAT_uint8, 1],
        'ProposalDescription':             [8, DAT_nvarchar16, 0],
        'ProposalDocumentHash':            [9, DAT_sha256, 32],
        'VoteCutOffTimestamp':             [10, DAT_time, 8]
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
        self.ProposedChangesCount = None
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
        'AssetSpecificVote':               [0, DAT_bool, 1],
        'AssetType':                       [1, DAT_string, 3],
        'AssetID':                         [2, DAT_string, 32],
        'VoteSystem':                      [3, DAT_uint8, 1],
        'Proposal':                        [4, DAT_bool, 1],
        'ProposedChangesCount':            [5, DAT_uint8, 1],
        'ProposedChanges':                 [6, DAT_Amendment[], 0],
        'VoteOptions':                     [7, DAT_nvarchar8, 0],
        'VoteMax':                         [8, DAT_uint8, 1],
        'ProposalDescription':             [9, DAT_nvarchar16, 0],
        'ProposalDocumentHash':            [10, DAT_sha256, 32],
        'VoteCutOffTimestamp':             [11, DAT_time, 8]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.AssetType = None
        self.AssetID = None
        self.VoteSystem = None
        self.Proposal = None
        self.ProposedChangesCount = None
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
        'AssetID':                         [0, DAT_string, 32],
        'VoteTxnID':                       [1, DAT_sha256, 32],
        'Vote':                            [2, DAT_nvarchar8, 0]
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
        'AssetType':                       [0, DAT_string, 3],
        'AssetID':                         [1, DAT_string, 32],
        'Proposal':                        [2, DAT_bool, 1],
        'ProposedChangesCount':            [3, DAT_uint8, 1],
        'ProposedChanges':                 [4, DAT_Amendment[], 0],
        'VoteTxnID':                       [5, DAT_sha256, 32],
        'VoteOptionsCount':                [6, DAT_uint8, 1],
        'OptionXTally':                    [7, DAT_uint64, 8],
        'Result':                          [8, DAT_nvarchar8, 0],
        'Timestamp':                       [9, DAT_timestamp, 8]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.AssetID = None
        self.Proposal = None
        self.ProposedChangesCount = None
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
        'QtyReceivingAddresses':           [0, DAT_uint8, 1],
        'AddressIndexes':                  [1, DAT_uint16[], 0],
        'MessageType':                     [2, DAT_string, 2],
        'MessagePayload':                  [3, DAT_nvarchar64, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.AddressIndexes = None
        self.MessageType = None
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
        'QtyReceivingAddresses':           [0, DAT_uint8, 1],
        'AddressIndexes':                  [1, DAT_uint16[], 0],
        'RejectionType':                   [2, DAT_uint8, 1],
        'MessagePayload':                  [3, DAT_nvarchar16, 0],
        'Timestamp':                       [4, DAT_timestamp, 8]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.AddressIndexes = None
        self.RejectionType = None
        self.MessagePayload = None
        self.Timestamp = None


# Establishment Action - Establishes an on-chain register.

class Action_Establishment(ActionBase):
    ActionPrefix = 'R1'

    schema = {
        'Message':                         [0, DAT_nvarchar64, 25]
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
        'Message':                         [0, DAT_nvarchar64, 0]
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
        'Message':                         [0, DAT_nvarchar64, 0]
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
        'Message':                         [0, DAT_nvarchar64, 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):


# Send Action - A Token Owner Sends a Token to a Receiver. The Send Action
# requires no sign-off by the Token Receiving Party and does not provide
# any on-chain consideration to the Token Sending Party. Can be used for
# redeeming a ticket, coupon, points, etc.

class Action_Send(ActionBase):
    ActionPrefix = 'T1'

    schema = {
        'AssetType':                       [0, DAT_string, 3],
        'AssetID':                         [1, DAT_string, 32],
        'TokenSenderCount':                [2, DAT_uint8, 1],
        'TokenSenders':                    [3, DAT_QuantityIndex[], 0],
        'TokenReceiverCount':              [4, DAT_uint8, 1],
        'TokenReceivers':                  [5, DAT_TokenReceiver[], 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.AssetID = None
        self.TokenSenderCount = None
        self.TokenSenders = None
        self.TokenReceiverCount = None
        self.TokenReceivers = None


# Exchange Action - Tokens exchanged for Bitcoin (BSV). Example: Bob (Token
# Sender) to sell 21,000 tokens to Alice (Token Receiver) for 7 BSV. Both
# parties must sign the transaction for it to be valid.

class Action_Exchange(ActionBase):
    ActionPrefix = 'T2'

    schema = {
        'AssetType':                       [0, DAT_string, 3],
        'AssetID':                         [1, DAT_string, 32],
        'OfferExpiry':                     [2, DAT_time, 8],
        'ExchangeFeeCurrency':             [3, DAT_string, 3],
        'ExchangeFeeVar':                  [4, DAT_float32, 4],
        'ExchangeFeeFixed':                [5, DAT_float32, 4],
        'ExchangeFeeAddress':              [6, DAT_string, 34],
        'TokenSenderCount':                [7, DAT_uint8, 1],
        'TokenSenders':                    [8, DAT_QuantityIndex[], 0],
        'TokenReceiverCount':              [9, DAT_uint8, 1],
        'TokenReceivers':                  [10, DAT_TokenReceiver[], 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.AssetID = None
        self.OfferExpiry = None
        self.ExchangeFeeCurrency = None
        self.ExchangeFeeVar = None
        self.ExchangeFeeFixed = None
        self.ExchangeFeeAddress = None
        self.TokenSenderCount = None
        self.TokenSenders = None
        self.TokenReceiverCount = None
        self.TokenReceivers = None


# Swap Action - Two parties (or more) want to swap a token (Atomic Swap)
# directly for another token. At a minimum, Bitcoin is used in the txn for
# paying the necessary network/transaction fees.

class Action_Swap(ActionBase):
    ActionPrefix = 'T3'

    schema = {
        'AssetType1':                      [0, DAT_string, 3],
        'AssetID1':                        [1, DAT_string, 32],
        'AssetType2':                      [2, DAT_string, 3],
        'AssetID2':                        [3, DAT_string, 32],
        'OfferExpiry':                     [4, DAT_time, 8],
        'ExchangeFeeCurrency':             [5, DAT_string, 3],
        'ExchangeFeeVar':                  [6, DAT_float32, 4],
        'ExchangeFeeFixed':                [7, DAT_float32, 4],
        'ExchangeFeeAddress':              [8, DAT_string, 34],
        'Asset1SenderCount':               [9, DAT_uint8, 1],
        'Asset1Senders':                   [10, DAT_QuantityIndex[], 0],
        'Asset1ReceiverCount':             [11, DAT_uint8, 1],
        'Asset1Receivers':                 [12, DAT_TokenReceiver[], 0],
        'Asset2SenderCount':               [13, DAT_uint8, 1],
        'Asset2Senders':                   [14, DAT_QuantityIndex[], 0],
        'Asset2ReceiverCount':             [15, DAT_uint8, 1],
        'Asset2Receivers':                 [16, DAT_TokenReceiver[], 0]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.AssetID1 = None
        self.AssetType2 = None
        self.AssetID2 = None
        self.OfferExpiry = None
        self.ExchangeFeeCurrency = None
        self.ExchangeFeeVar = None
        self.ExchangeFeeFixed = None
        self.ExchangeFeeAddress = None
        self.Asset1SenderCount = None
        self.Asset1Senders = None
        self.Asset1ReceiverCount = None
        self.Asset1Receivers = None
        self.Asset2SenderCount = None
        self.Asset2Senders = None
        self.Asset2ReceiverCount = None
        self.Asset2Receivers = None


# Settlement Action - Finalizes the transfer of bitcoins and tokens from
# send, exchange, and swap actions.

class Action_Settlement(ActionBase):
    ActionPrefix = 'T4'

    schema = {
        'TransferType':                    [0, DAT_string, 1],
        'AssetType1':                      [1, DAT_string, 3],
        'AssetID1':                        [2, DAT_string, 32],
        'AssetType2':                      [3, DAT_string, 3],
        'AssetID2':                        [4, DAT_string, 32],
        'Asset1SettlementsCount':          [5, DAT_uint8, 1],
        'Asset1AddressesXQty':             [6, DAT_QuantityIndex[], 0],
        'Asset2SettlementsCount':          [7, DAT_uint8, 1],
        'Asset2AddressXQty':               [8, DAT_QuantityIndex[], 0],
        'Timestamp':                       [9, DAT_timestamp, 8]
    }

    rules = {
        'contractFee': 0,
        'inputs': [ACT_CONTRACT],
        'outputs': [ACT_USER, ACT_CONTRACT]
    }

    def init_attributes(self):
        self.AssetType1 = None
        self.AssetID1 = None
        self.AssetType2 = None
        self.AssetID2 = None
        self.Asset1SettlementsCount = None
        self.Asset1AddressesXQty = None
        self.Asset2SettlementsCount = None
        self.Asset2AddressXQty = None
        self.Timestamp = None



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
    'T1': Action_Send,
    'T2': Action_Exchange,
    'T3': Action_Swap,
    'T4': Action_Settlement
}
