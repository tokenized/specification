"use strict";
// Package protocol provides base level structs and validation for
// the protocol.
//
// The code in this file is auto-generated. Do not edit it by hand as it will
// be overwritten when code is regenerated.
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const sprintf_js_1 = require("sprintf-js");
const util_1 = __importDefault(require("@keyring/util"));
const bytes_1 = require("./bytes");
const protocol_types_1 = require("./protocol_types");
const field_types_1 = require("./field_types");
const resources_1 = require("./resources");
const protocol_1 = require("./protocol");
var ActionCode;
(function (ActionCode) {
    // CodeAssetDefinition identifies data as a AssetDefinition message.
    ActionCode["CodeAssetDefinition"] = "A1";
    // CodeAssetCreation identifies data as a AssetCreation message.
    ActionCode["CodeAssetCreation"] = "A2";
    // CodeAssetModification identifies data as a AssetModification message.
    ActionCode["CodeAssetModification"] = "A3";
    // CodeContractOffer identifies data as a ContractOffer message.
    ActionCode["CodeContractOffer"] = "C1";
    // CodeContractFormation identifies data as a ContractFormation message.
    ActionCode["CodeContractFormation"] = "C2";
    // CodeContractAmendment identifies data as a ContractAmendment message.
    ActionCode["CodeContractAmendment"] = "C3";
    // CodeStaticContractFormation identifies data as a
    // StaticContractFormation message.
    ActionCode["CodeStaticContractFormation"] = "C4";
    // CodeContractAddressChange identifies data as a ContractAddressChange
    // message.
    ActionCode["CodeContractAddressChange"] = "C5";
    // CodeOrder identifies data as a Order message.
    ActionCode["CodeOrder"] = "E1";
    // CodeFreeze identifies data as a Freeze message.
    ActionCode["CodeFreeze"] = "E2";
    // CodeThaw identifies data as a Thaw message.
    ActionCode["CodeThaw"] = "E3";
    // CodeConfiscation identifies data as a Confiscation message.
    ActionCode["CodeConfiscation"] = "E4";
    // CodeReconciliation identifies data as a Reconciliation message.
    ActionCode["CodeReconciliation"] = "E5";
    // CodeProposal identifies data as a Proposal message.
    ActionCode["CodeProposal"] = "G1";
    // CodeVote identifies data as a Vote message.
    ActionCode["CodeVote"] = "G2";
    // CodeBallotCast identifies data as a BallotCast message.
    ActionCode["CodeBallotCast"] = "G3";
    // CodeBallotCounted identifies data as a BallotCounted message.
    ActionCode["CodeBallotCounted"] = "G4";
    // CodeResult identifies data as a Result message.
    ActionCode["CodeResult"] = "G5";
    // CodeMessage identifies data as a Message message.
    ActionCode["CodeMessage"] = "M1";
    // CodeRejection identifies data as a Rejection message.
    ActionCode["CodeRejection"] = "M2";
    // CodeEstablishment identifies data as a Establishment message.
    ActionCode["CodeEstablishment"] = "R1";
    // CodeAddition identifies data as a Addition message.
    ActionCode["CodeAddition"] = "R2";
    // CodeAlteration identifies data as a Alteration message.
    ActionCode["CodeAlteration"] = "R3";
    // CodeRemoval identifies data as a Removal message.
    ActionCode["CodeRemoval"] = "R4";
    // CodeTransfer identifies data as a Transfer message.
    ActionCode["CodeTransfer"] = "T1";
    // CodeSettlement identifies data as a Settlement message.
    ActionCode["CodeSettlement"] = "T2";
    // ComplianceActionFreeze identifies a freeze type
    ActionCode["ComplianceActionFreeze"] = "F";
    // ComplianceActionThaw identifies a thaw type
    ActionCode["ComplianceActionThaw"] = "T";
    // ComplianceActionConfiscation identifies a confiscation type
    ActionCode["ComplianceActionConfiscation"] = "C";
    // ComplianceActionReconciliation identifies a reconcilation type
    ActionCode["ComplianceActionReconciliation"] = "R";
})(ActionCode = exports.ActionCode || (exports.ActionCode = {}));
// TypeMapping holds a mapping of action codes to action types.
function TypeMapping(code) {
    switch (code) {
        case ActionCode.CodeAssetDefinition:
            return new AssetDefinition();
        case ActionCode.CodeAssetCreation:
            return new AssetCreation();
        case ActionCode.CodeAssetModification:
            return new AssetModification();
        case ActionCode.CodeContractOffer:
            return new ContractOffer();
        case ActionCode.CodeContractFormation:
            return new ContractFormation();
        case ActionCode.CodeContractAmendment:
            return new ContractAmendment();
        case ActionCode.CodeStaticContractFormation:
            return new StaticContractFormation();
        case ActionCode.CodeContractAddressChange:
            return new ContractAddressChange();
        case ActionCode.CodeOrder:
            return new Order();
        case ActionCode.CodeFreeze:
            return new Freeze();
        case ActionCode.CodeThaw:
            return new Thaw();
        case ActionCode.CodeConfiscation:
            return new Confiscation();
        case ActionCode.CodeReconciliation:
            return new Reconciliation();
        case ActionCode.CodeProposal:
            return new Proposal();
        case ActionCode.CodeVote:
            return new Vote();
        case ActionCode.CodeBallotCast:
            return new BallotCast();
        case ActionCode.CodeBallotCounted:
            return new BallotCounted();
        case ActionCode.CodeResult:
            return new Result();
        case ActionCode.CodeMessage:
            return new Message();
        case ActionCode.CodeRejection:
            return new Rejection();
        case ActionCode.CodeEstablishment:
            return new Establishment();
        case ActionCode.CodeAddition:
            return new Addition();
        case ActionCode.CodeAlteration:
            return new Alteration();
        case ActionCode.CodeRemoval:
            return new Removal();
        case ActionCode.CodeTransfer:
            return new Transfer();
        case ActionCode.CodeSettlement:
            return new Settlement();
        default:
            return null;
    }
}
exports.TypeMapping = TypeMapping;
// AssetDefinition This action is used by the administration to define the
// properties/characteristics of the asset (token) that it wants to create.
// An asset has a unique identifier called AssetID = AssetType +
// base58(AssetCode + checksum). An asset is always linked to a contract
// that is identified by the public address of the Contract wallet.
class AssetDefinition extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeAssetDefinition;
        this.typeStr = 'AssetDefinition';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeAssetDefinition;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // asset_type (string)
        {
            bytes_1.WriteFixedChar(buf, this.asset_type, 3);
        }
        // asset_auth_flags ([]byte)
        {
            bytes_1.WriteVarBin(buf, this.asset_auth_flags, 8);
        }
        // transfers_permitted (bool)
        {
            bytes_1.write(buf, this.transfers_permitted, 'bool');
        }
        // trade_restrictions ([][3]byte)
        {
            // IsNativeTypeArray [][3]byte
            const type = '[][3]byte'.slice(2);
            bytes_1.WriteVariableSize(buf, this.trade_restrictions.length, 16, 8);
            this.trade_restrictions.forEach(value => {
                bytes_1.write(buf, value, type); // TODO
            });
        }
        // enforcement_orders_permitted (bool)
        {
            bytes_1.write(buf, this.enforcement_orders_permitted, 'bool');
        }
        // voting_rights (bool)
        {
            bytes_1.write(buf, this.voting_rights, 'bool');
        }
        // vote_multiplier (uint8)
        {
            bytes_1.write(buf, this.vote_multiplier, 'uint8');
        }
        // administration_proposal (bool)
        {
            bytes_1.write(buf, this.administration_proposal, 'bool');
        }
        // holder_proposal (bool)
        {
            bytes_1.write(buf, this.holder_proposal, 'bool');
        }
        // asset_modification_governance (uint8)
        {
            bytes_1.write(buf, this.asset_modification_governance, 'uint8');
        }
        // token_qty (uint64)
        {
            bytes_1.write(buf, this.token_qty, 'uint64');
        }
        // asset_payload ([]byte)
        {
            bytes_1.WriteVarBin(buf, this.asset_payload, 16);
        }
        return buf.buf;
    }
    // write populates the fields in AssetDefinition from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // asset_type (string)
        {
            this.asset_type = bytes_1.ReadFixedChar(buf, 3);
        }
        // asset_auth_flags ([]byte)
        {
            this.asset_auth_flags = bytes_1.ReadVarBin(buf, 8);
        }
        // transfers_permitted (bool)
        {
            this.transfers_permitted = bytes_1.read(buf, 'bool');
        }
        // trade_restrictions ([][3]byte)
        {
            const size = bytes_1.ReadVariableSize(buf, 16, 8);
            const type = '[][3]byte'.slice(2);
            this.trade_restrictions = [...Array(size)].map(() => bytes_1.read(buf, type));
        }
        // enforcement_orders_permitted (bool)
        {
            this.enforcement_orders_permitted = bytes_1.read(buf, 'bool');
        }
        // voting_rights (bool)
        {
            this.voting_rights = bytes_1.read(buf, 'bool');
        }
        // vote_multiplier (uint8)
        {
            this.vote_multiplier = bytes_1.read(buf, 'uint8');
        }
        // administration_proposal (bool)
        {
            this.administration_proposal = bytes_1.read(buf, 'bool');
        }
        // holder_proposal (bool)
        {
            this.holder_proposal = bytes_1.read(buf, 'bool');
        }
        // asset_modification_governance (uint8)
        {
            this.asset_modification_governance = bytes_1.read(buf, 'uint8');
        }
        // token_qty (uint64)
        {
            this.token_qty = bytes_1.read(buf, 'uint64');
        }
        // asset_payload ([]byte)
        {
            this.asset_payload = bytes_1.ReadVarBin(buf, 16);
        }
        return b.length - buf.length;
    }
    Validate() {
        // AssetType (string) 
        {
            if (this.asset_type.length > 3) {
                return sprintf_js_1.sprintf('fixedchar field asset_type too long %d/%d', this.asset_type.length, 3);
            }
        }
        // AssetAuthFlags ([]byte) 
        {
            if (this.asset_auth_flags.length >= (Math.pow(2, 8))) {
                return sprintf_js_1.sprintf('varbin field asset_auth_flags too long %d/%d', this.asset_auth_flags.length, (Math.pow(2, 8)) - 1);
            }
        }
        // TransfersPermitted (bool) 
        {
        }
        // TradeRestrictions ([][3]byte) 
        {
            if (this.trade_restrictions.length >= (Math.pow(2, 16))) {
                return sprintf_js_1.sprintf('list field trade_restrictions has too many items %d/%d', this.trade_restrictions.length, (Math.pow(2, 16)) - 1);
            }
            let err = null;
            this.trade_restrictions.find((value) => {
                const str = Buffer.from(value).toString('ascii');
                if (!resources_1.Resources.GetPolityType(str)) {
                    err = sprintf_js_1.sprintf('trade_restrictions: Invalid polity value : %s', str);
                    return true;
                }
            });
            if (err)
                return err;
        }
        // EnforcementOrdersPermitted (bool) 
        {
        }
        // VotingRights (bool) 
        {
        }
        // VoteMultiplier (uint8) 
        {
        }
        // AdministrationProposal (bool) 
        {
        }
        // HolderProposal (bool) 
        {
        }
        // AssetModificationGovernance (uint8) 
        {
            // $field.IntValues [0 1]
            if (this.asset_modification_governance !== 0 && this.asset_modification_governance !== 1) {
                return sprintf_js_1.sprintf('field asset_modification_governance value is invalid : %d', this.asset_modification_governance);
            }
        }
        // TokenQty (uint64) 
        {
        }
        // AssetPayload ([]byte) 
        {
            if (this.asset_payload.length >= (Math.pow(2, 16))) {
                return sprintf_js_1.sprintf('varbin field asset_payload too long %d/%d', this.asset_payload.length, (Math.pow(2, 16)) - 1);
            }
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('asset_type:%s', this.asset_type.toString()));
        vals.push(sprintf_js_1.sprintf('asset_auth_flags:%s', this.asset_auth_flags.toString()));
        vals.push(sprintf_js_1.sprintf('transfers_permitted:%s', this.transfers_permitted.toString()));
        vals.push(sprintf_js_1.sprintf('trade_restrictions:%s', this.trade_restrictions.toString()));
        vals.push(sprintf_js_1.sprintf('enforcement_orders_permitted:%s', this.enforcement_orders_permitted.toString()));
        vals.push(sprintf_js_1.sprintf('voting_rights:%s', this.voting_rights.toString()));
        vals.push(sprintf_js_1.sprintf('vote_multiplier:%d', this.vote_multiplier));
        vals.push(sprintf_js_1.sprintf('administration_proposal:%s', this.administration_proposal.toString()));
        vals.push(sprintf_js_1.sprintf('holder_proposal:%s', this.holder_proposal.toString()));
        vals.push(sprintf_js_1.sprintf('asset_modification_governance:%d', this.asset_modification_governance));
        vals.push(sprintf_js_1.sprintf('token_qty:%d', this.token_qty));
        vals.push(sprintf_js_1.sprintf('asset_payload:%s', this.asset_payload.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for AssetDefinition.', e);
            throw e;
        }
        // this.asset_type = parsed.assetType (fixedchar)
        // this.asset_auth_flags = parsed.assetAuthFlags (varbin)
        // this.transfers_permitted = parsed.transfersPermitted (bool)
        // this.trade_restrictions = parsed.tradeRestrictions (Polity[])
        // this.enforcement_orders_permitted = parsed.enforcementOrdersPermitted (bool)
        // this.voting_rights = parsed.votingRights (bool)
        this.vote_multiplier = parsed.voteMultiplier;
        // this.administration_proposal = parsed.administrationProposal (bool)
        // this.holder_proposal = parsed.holderProposal (bool)
        this.asset_modification_governance = parsed.assetModificationGovernance;
        this.token_qty = parsed.tokenQty;
        // this.asset_payload = parsed.assetPayload (varbin)
        return this.Validate();
    }
}
exports.AssetDefinition = AssetDefinition;
// AssetCreation This action creates an asset in response to the
// administration's instructions in the Definition Action.
class AssetCreation extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeAssetCreation;
        this.typeStr = 'AssetCreation';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeAssetCreation;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // asset_type (string)
        {
            bytes_1.WriteFixedChar(buf, this.asset_type, 3);
        }
        // asset_code (AssetCode)
        {
            buf.write(this.asset_code.Serialize());
        }
        // asset_index (uint64)
        {
            bytes_1.write(buf, this.asset_index, 'uint64');
        }
        // asset_auth_flags ([]byte)
        {
            bytes_1.WriteVarBin(buf, this.asset_auth_flags, 8);
        }
        // transfers_permitted (bool)
        {
            bytes_1.write(buf, this.transfers_permitted, 'bool');
        }
        // trade_restrictions ([][3]byte)
        {
            // IsNativeTypeArray [][3]byte
            const type = '[][3]byte'.slice(2);
            bytes_1.WriteVariableSize(buf, this.trade_restrictions.length, 16, 8);
            this.trade_restrictions.forEach(value => {
                bytes_1.write(buf, value, type); // TODO
            });
        }
        // enforcement_orders_permitted (bool)
        {
            bytes_1.write(buf, this.enforcement_orders_permitted, 'bool');
        }
        // voting_rights (bool)
        {
            bytes_1.write(buf, this.voting_rights, 'bool');
        }
        // vote_multiplier (uint8)
        {
            bytes_1.write(buf, this.vote_multiplier, 'uint8');
        }
        // administration_proposal (bool)
        {
            bytes_1.write(buf, this.administration_proposal, 'bool');
        }
        // holder_proposal (bool)
        {
            bytes_1.write(buf, this.holder_proposal, 'bool');
        }
        // asset_modification_governance (uint8)
        {
            bytes_1.write(buf, this.asset_modification_governance, 'uint8');
        }
        // token_qty (uint64)
        {
            bytes_1.write(buf, this.token_qty, 'uint64');
        }
        // asset_payload ([]byte)
        {
            bytes_1.WriteVarBin(buf, this.asset_payload, 16);
        }
        // asset_revision (uint32)
        {
            bytes_1.write(buf, this.asset_revision, 'uint32');
        }
        // timestamp (Timestamp)
        {
            buf.write(this.timestamp.Serialize());
        }
        return buf.buf;
    }
    // write populates the fields in AssetCreation from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // asset_type (string)
        {
            this.asset_type = bytes_1.ReadFixedChar(buf, 3);
        }
        // asset_code (AssetCode)
        {
            this.asset_code = new protocol_types_1.AssetCode();
            this.asset_code.Write(buf);
        }
        // asset_index (uint64)
        {
            this.asset_index = bytes_1.read(buf, 'uint64');
        }
        // asset_auth_flags ([]byte)
        {
            this.asset_auth_flags = bytes_1.ReadVarBin(buf, 8);
        }
        // transfers_permitted (bool)
        {
            this.transfers_permitted = bytes_1.read(buf, 'bool');
        }
        // trade_restrictions ([][3]byte)
        {
            const size = bytes_1.ReadVariableSize(buf, 16, 8);
            const type = '[][3]byte'.slice(2);
            this.trade_restrictions = [...Array(size)].map(() => bytes_1.read(buf, type));
        }
        // enforcement_orders_permitted (bool)
        {
            this.enforcement_orders_permitted = bytes_1.read(buf, 'bool');
        }
        // voting_rights (bool)
        {
            this.voting_rights = bytes_1.read(buf, 'bool');
        }
        // vote_multiplier (uint8)
        {
            this.vote_multiplier = bytes_1.read(buf, 'uint8');
        }
        // administration_proposal (bool)
        {
            this.administration_proposal = bytes_1.read(buf, 'bool');
        }
        // holder_proposal (bool)
        {
            this.holder_proposal = bytes_1.read(buf, 'bool');
        }
        // asset_modification_governance (uint8)
        {
            this.asset_modification_governance = bytes_1.read(buf, 'uint8');
        }
        // token_qty (uint64)
        {
            this.token_qty = bytes_1.read(buf, 'uint64');
        }
        // asset_payload ([]byte)
        {
            this.asset_payload = bytes_1.ReadVarBin(buf, 16);
        }
        // asset_revision (uint32)
        {
            this.asset_revision = bytes_1.read(buf, 'uint32');
        }
        // timestamp (Timestamp)
        {
            this.timestamp = new protocol_types_1.Timestamp();
            this.timestamp.Write(buf);
        }
        return b.length - buf.length;
    }
    Validate() {
        // AssetType (string) 
        {
            if (this.asset_type.length > 3) {
                return sprintf_js_1.sprintf('fixedchar field asset_type too long %d/%d', this.asset_type.length, 3);
            }
        }
        // AssetCode (AssetCode) 
        {
            // IsInternalType
            const err = this.asset_code.Validate();
            if (err)
                return sprintf_js_1.sprintf('field asset_code is invalid : %s', err);
        }
        // AssetIndex (uint64) 
        {
        }
        // AssetAuthFlags ([]byte) 
        {
            if (this.asset_auth_flags.length >= (Math.pow(2, 8))) {
                return sprintf_js_1.sprintf('varbin field asset_auth_flags too long %d/%d', this.asset_auth_flags.length, (Math.pow(2, 8)) - 1);
            }
        }
        // TransfersPermitted (bool) 
        {
        }
        // TradeRestrictions ([][3]byte) 
        {
            if (this.trade_restrictions.length >= (Math.pow(2, 16))) {
                return sprintf_js_1.sprintf('list field trade_restrictions has too many items %d/%d', this.trade_restrictions.length, (Math.pow(2, 16)) - 1);
            }
            let err = null;
            this.trade_restrictions.find((value) => {
                const str = Buffer.from(value).toString('ascii');
                if (!resources_1.Resources.GetPolityType(str)) {
                    err = sprintf_js_1.sprintf('trade_restrictions: Invalid polity value : %s', str);
                    return true;
                }
            });
            if (err)
                return err;
        }
        // EnforcementOrdersPermitted (bool) 
        {
        }
        // VotingRights (bool) 
        {
        }
        // VoteMultiplier (uint8) 
        {
        }
        // AdministrationProposal (bool) 
        {
        }
        // HolderProposal (bool) 
        {
        }
        // AssetModificationGovernance (uint8) 
        {
            // $field.IntValues [0 1]
            if (this.asset_modification_governance !== 0 && this.asset_modification_governance !== 1) {
                return sprintf_js_1.sprintf('field asset_modification_governance value is invalid : %d', this.asset_modification_governance);
            }
        }
        // TokenQty (uint64) 
        {
        }
        // AssetPayload ([]byte) 
        {
            if (this.asset_payload.length >= (Math.pow(2, 16))) {
                return sprintf_js_1.sprintf('varbin field asset_payload too long %d/%d', this.asset_payload.length, (Math.pow(2, 16)) - 1);
            }
        }
        // AssetRevision (uint32) 
        {
        }
        // Timestamp (Timestamp) 
        {
            // IsInternalType
            const err = this.timestamp.Validate();
            if (err)
                return sprintf_js_1.sprintf('field timestamp is invalid : %s', err);
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('asset_type:%s', this.asset_type.toString()));
        vals.push(sprintf_js_1.sprintf('asset_code:%s', this.asset_code.toString()));
        vals.push(sprintf_js_1.sprintf('asset_index:%d', this.asset_index));
        vals.push(sprintf_js_1.sprintf('asset_auth_flags:%s', this.asset_auth_flags.toString()));
        vals.push(sprintf_js_1.sprintf('transfers_permitted:%s', this.transfers_permitted.toString()));
        vals.push(sprintf_js_1.sprintf('trade_restrictions:%s', this.trade_restrictions.toString()));
        vals.push(sprintf_js_1.sprintf('enforcement_orders_permitted:%s', this.enforcement_orders_permitted.toString()));
        vals.push(sprintf_js_1.sprintf('voting_rights:%s', this.voting_rights.toString()));
        vals.push(sprintf_js_1.sprintf('vote_multiplier:%d', this.vote_multiplier));
        vals.push(sprintf_js_1.sprintf('administration_proposal:%s', this.administration_proposal.toString()));
        vals.push(sprintf_js_1.sprintf('holder_proposal:%s', this.holder_proposal.toString()));
        vals.push(sprintf_js_1.sprintf('asset_modification_governance:%d', this.asset_modification_governance));
        vals.push(sprintf_js_1.sprintf('token_qty:%d', this.token_qty));
        vals.push(sprintf_js_1.sprintf('asset_payload:%s', this.asset_payload.toString()));
        vals.push(sprintf_js_1.sprintf('asset_revision:%d', this.asset_revision));
        vals.push(sprintf_js_1.sprintf('timestamp:%s', this.timestamp.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for AssetCreation.', e);
            throw e;
        }
        // this.asset_type = parsed.assetType (fixedchar)
        this.asset_code = new protocol_types_1.AssetCode();
        this.asset_code.fromJSON(JSON.stringify(parsed.assetCode));
        this.asset_index = parsed.assetIndex;
        // this.asset_auth_flags = parsed.assetAuthFlags (varbin)
        // this.transfers_permitted = parsed.transfersPermitted (bool)
        // this.trade_restrictions = parsed.tradeRestrictions (Polity[])
        // this.enforcement_orders_permitted = parsed.enforcementOrdersPermitted (bool)
        // this.voting_rights = parsed.votingRights (bool)
        this.vote_multiplier = parsed.voteMultiplier;
        // this.administration_proposal = parsed.administrationProposal (bool)
        // this.holder_proposal = parsed.holderProposal (bool)
        this.asset_modification_governance = parsed.assetModificationGovernance;
        this.token_qty = parsed.tokenQty;
        // this.asset_payload = parsed.assetPayload (varbin)
        this.asset_revision = parsed.assetRevision;
        this.timestamp = new protocol_types_1.Timestamp();
        this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
        return this.Validate();
    }
}
exports.AssetCreation = AssetCreation;
// AssetModification Token Dilutions, Call Backs/Revocations, burning etc.
class AssetModification extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeAssetModification;
        this.typeStr = 'AssetModification';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeAssetModification;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // asset_type (string)
        {
            bytes_1.WriteFixedChar(buf, this.asset_type, 3);
        }
        // asset_code (AssetCode)
        {
            buf.write(this.asset_code.Serialize());
        }
        // asset_revision (uint32)
        {
            bytes_1.write(buf, this.asset_revision, 'uint32');
        }
        // amendments ([]Amendment)
        {
            bytes_1.WriteVariableSize(buf, this.amendments.length, 8, 8);
            this.amendments.forEach((value) => {
                buf.write(value.Serialize());
            });
        }
        // ref_tx_id (TxId)
        {
            buf.write(this.ref_tx_id.Serialize());
        }
        return buf.buf;
    }
    // write populates the fields in AssetModification from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // asset_type (string)
        {
            this.asset_type = bytes_1.ReadFixedChar(buf, 3);
        }
        // asset_code (AssetCode)
        {
            this.asset_code = new protocol_types_1.AssetCode();
            this.asset_code.Write(buf);
        }
        // asset_revision (uint32)
        {
            this.asset_revision = bytes_1.read(buf, 'uint32');
        }
        // amendments ([]Amendment)
        {
            // IsInternalTypeArray
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            this.amendments = [];
            for (let i = 0; i < size; i++) {
                const newValue = new field_types_1.Amendment();
                newValue.Write(buf);
                this.amendments.push(newValue);
            }
        }
        // ref_tx_id (TxId)
        {
            this.ref_tx_id = new protocol_types_1.TxId();
            this.ref_tx_id.Write(buf);
        }
        return b.length - buf.length;
    }
    Validate() {
        // AssetType (string) 
        {
            if (this.asset_type.length > 3) {
                return sprintf_js_1.sprintf('fixedchar field asset_type too long %d/%d', this.asset_type.length, 3);
            }
        }
        // AssetCode (AssetCode) 
        {
            // IsInternalType
            const err = this.asset_code.Validate();
            if (err)
                return sprintf_js_1.sprintf('field asset_code is invalid : %s', err);
        }
        // AssetRevision (uint32) 
        {
        }
        // Amendments ([]Amendment) 
        {
            if (this.amendments.length > (Math.pow(2, 8)) - 1) {
                return sprintf_js_1.sprintf('list field amendments has too many items %d/%d', this.amendments.length, (Math.pow(2, 8)) - 1);
            }
            const err = this.amendments.find((value, i) => {
                const err2 = value.Validate();
                if (err2)
                    return sprintf_js_1.sprintf('list field amendments[%d] is invalid : %s', i, err2);
            });
            if (err)
                return err;
        }
        // RefTxID (TxId) 
        {
            // IsInternalType
            const err = this.ref_tx_id.Validate();
            if (err)
                return sprintf_js_1.sprintf('field ref_tx_id is invalid : %s', err);
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('asset_type:%s', this.asset_type.toString()));
        vals.push(sprintf_js_1.sprintf('asset_code:%s', this.asset_code.toString()));
        vals.push(sprintf_js_1.sprintf('asset_revision:%d', this.asset_revision));
        vals.push(sprintf_js_1.sprintf('amendments:%s', this.amendments.toString()));
        vals.push(sprintf_js_1.sprintf('ref_tx_id:%s', this.ref_tx_id.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for AssetModification.', e);
            throw e;
        }
        // this.asset_type = parsed.assetType (fixedchar)
        this.asset_code = new protocol_types_1.AssetCode();
        this.asset_code.fromJSON(JSON.stringify(parsed.assetCode));
        this.asset_revision = parsed.assetRevision;
        // this.amendments (Amendment[])
        this.ref_tx_id = new protocol_types_1.TxId();
        this.ref_tx_id.fromJSON(JSON.stringify(parsed.refTxID));
        return this.Validate();
    }
}
exports.AssetModification = AssetModification;
// ContractOffer Allows the administration to tell the smart contract what
// they want the details (labels, data, T&C's, etc.) of the Contract to be
// on-chain in a public and immutable way. The Contract Offer action
// 'initializes' a generic smart contract that has been spun up by either
// the smart contract operator or the administration. This on-chain action
// allows for the positive response from the smart contract with either a
// Contract Formation Action or a Rejection Action.
class ContractOffer extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeContractOffer;
        this.typeStr = 'ContractOffer';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeContractOffer;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // contract_name (string)
        {
            bytes_1.WriteVarChar(buf, this.contract_name, 8);
        }
        // body_of_agreement_type (uint8)
        {
            bytes_1.write(buf, this.body_of_agreement_type, 'uint8');
        }
        // body_of_agreement ([]byte)
        {
            bytes_1.WriteVarBin(buf, this.body_of_agreement, 32);
        }
        // contract_type (string)
        {
            bytes_1.WriteVarChar(buf, this.contract_type, 8);
        }
        // supporting_docs ([]Document)
        {
            bytes_1.WriteVariableSize(buf, this.supporting_docs.length, 8, 8);
            this.supporting_docs.forEach((value) => {
                buf.write(value.Serialize());
            });
        }
        // governing_law (string)
        {
            bytes_1.WriteFixedChar(buf, this.governing_law, 5);
        }
        // jurisdiction (string)
        {
            bytes_1.WriteFixedChar(buf, this.jurisdiction, 5);
        }
        // contract_expiration (Timestamp)
        {
            buf.write(this.contract_expiration.Serialize());
        }
        // contract_uri (string)
        {
            bytes_1.WriteVarChar(buf, this.contract_uri, 8);
        }
        // issuer (Entity)
        {
            buf.write(this.issuer.Serialize());
        }
        // issuer_logo_url (string)
        {
            bytes_1.WriteVarChar(buf, this.issuer_logo_url, 8);
        }
        // contract_operator_included (bool)
        {
            bytes_1.write(buf, this.contract_operator_included, 'bool');
        }
        // contract_operator (Entity)
        if (this.contract_operator_included) {
            buf.write(this.contract_operator.Serialize());
        }
        // contract_fee (uint64)
        {
            bytes_1.write(buf, this.contract_fee, 'uint64');
        }
        // voting_systems ([]VotingSystem)
        {
            bytes_1.WriteVariableSize(buf, this.voting_systems.length, 8, 8);
            this.voting_systems.forEach((value) => {
                buf.write(value.Serialize());
            });
        }
        // contract_auth_flags ([]byte)
        {
            bytes_1.WriteVarBin(buf, this.contract_auth_flags, 16);
        }
        // restricted_qty_assets (uint64)
        {
            bytes_1.write(buf, this.restricted_qty_assets, 'uint64');
        }
        // administration_proposal (bool)
        {
            bytes_1.write(buf, this.administration_proposal, 'bool');
        }
        // holder_proposal (bool)
        {
            bytes_1.write(buf, this.holder_proposal, 'bool');
        }
        // oracles ([]Oracle)
        {
            bytes_1.WriteVariableSize(buf, this.oracles.length, 8, 8);
            this.oracles.forEach((value) => {
                buf.write(value.Serialize());
            });
        }
        // master_pkh (PublicKeyHash)
        {
            buf.write(this.master_pkh.Serialize());
        }
        return buf.buf;
    }
    // write populates the fields in ContractOffer from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // contract_name (string)
        {
            this.contract_name = bytes_1.ReadVarChar(buf, 8);
        }
        // body_of_agreement_type (uint8)
        {
            this.body_of_agreement_type = bytes_1.read(buf, 'uint8');
        }
        // body_of_agreement ([]byte)
        {
            this.body_of_agreement = bytes_1.ReadVarBin(buf, 32);
        }
        // contract_type (string)
        {
            this.contract_type = bytes_1.ReadVarChar(buf, 8);
        }
        // supporting_docs ([]Document)
        {
            // IsInternalTypeArray
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            this.supporting_docs = [];
            for (let i = 0; i < size; i++) {
                const newValue = new field_types_1.Document();
                newValue.Write(buf);
                this.supporting_docs.push(newValue);
            }
        }
        // governing_law (string)
        {
            this.governing_law = bytes_1.ReadFixedChar(buf, 5);
        }
        // jurisdiction (string)
        {
            this.jurisdiction = bytes_1.ReadFixedChar(buf, 5);
        }
        // contract_expiration (Timestamp)
        {
            this.contract_expiration = new protocol_types_1.Timestamp();
            this.contract_expiration.Write(buf);
        }
        // contract_uri (string)
        {
            this.contract_uri = bytes_1.ReadVarChar(buf, 8);
        }
        // issuer (Entity)
        {
            this.issuer = new field_types_1.Entity();
            this.issuer.Write(buf);
        }
        // issuer_logo_url (string)
        {
            this.issuer_logo_url = bytes_1.ReadVarChar(buf, 8);
        }
        // contract_operator_included (bool)
        {
            this.contract_operator_included = bytes_1.read(buf, 'bool');
        }
        // contract_operator (Entity)
        if (this.contract_operator_included) {
            this.contract_operator = new field_types_1.Entity();
            this.contract_operator.Write(buf);
        }
        // contract_fee (uint64)
        {
            this.contract_fee = bytes_1.read(buf, 'uint64');
        }
        // voting_systems ([]VotingSystem)
        {
            // IsInternalTypeArray
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            this.voting_systems = [];
            for (let i = 0; i < size; i++) {
                const newValue = new field_types_1.VotingSystem();
                newValue.Write(buf);
                this.voting_systems.push(newValue);
            }
        }
        // contract_auth_flags ([]byte)
        {
            this.contract_auth_flags = bytes_1.ReadVarBin(buf, 16);
        }
        // restricted_qty_assets (uint64)
        {
            this.restricted_qty_assets = bytes_1.read(buf, 'uint64');
        }
        // administration_proposal (bool)
        {
            this.administration_proposal = bytes_1.read(buf, 'bool');
        }
        // holder_proposal (bool)
        {
            this.holder_proposal = bytes_1.read(buf, 'bool');
        }
        // oracles ([]Oracle)
        {
            // IsInternalTypeArray
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            this.oracles = [];
            for (let i = 0; i < size; i++) {
                const newValue = new field_types_1.Oracle();
                newValue.Write(buf);
                this.oracles.push(newValue);
            }
        }
        // master_pkh (PublicKeyHash)
        {
            this.master_pkh = new protocol_types_1.PublicKeyHash();
            this.master_pkh.Write(buf);
        }
        return b.length - buf.length;
    }
    Validate() {
        // ContractName (string) 
        {
            if (this.contract_name.length > (Math.pow(2, 8))) {
                return sprintf_js_1.sprintf('varchar field contract_name too long %d/%d', this.contract_name.length, (Math.pow(2, 8)) - 1);
            }
        }
        // BodyOfAgreementType (uint8) 
        {
            // $field.IntValues [1 2]
            if (this.body_of_agreement_type !== 1 && this.body_of_agreement_type !== 2) {
                return sprintf_js_1.sprintf('field body_of_agreement_type value is invalid : %d', this.body_of_agreement_type);
            }
        }
        // BodyOfAgreement ([]byte) 
        {
            if (this.body_of_agreement.length >= (Math.pow(2, 32))) {
                return sprintf_js_1.sprintf('varbin field body_of_agreement too long %d/%d', this.body_of_agreement.length, (Math.pow(2, 32)) - 1);
            }
        }
        // ContractType (string) 
        {
            if (this.contract_type.length > (Math.pow(2, 8))) {
                return sprintf_js_1.sprintf('varchar field contract_type too long %d/%d', this.contract_type.length, (Math.pow(2, 8)) - 1);
            }
        }
        // SupportingDocs ([]Document) 
        {
            if (this.supporting_docs.length > (Math.pow(2, 8)) - 1) {
                return sprintf_js_1.sprintf('list field supporting_docs has too many items %d/%d', this.supporting_docs.length, (Math.pow(2, 8)) - 1);
            }
            const err = this.supporting_docs.find((value, i) => {
                const err2 = value.Validate();
                if (err2)
                    return sprintf_js_1.sprintf('list field supporting_docs[%d] is invalid : %s', i, err2);
            });
            if (err)
                return err;
        }
        // GoverningLaw (string) 
        {
            if (this.governing_law.length > 5) {
                return sprintf_js_1.sprintf('fixedchar field governing_law too long %d/%d', this.governing_law.length, 5);
            }
        }
        // Jurisdiction (string) 
        {
            if (this.jurisdiction.length > 5) {
                return sprintf_js_1.sprintf('fixedchar field jurisdiction too long %d/%d', this.jurisdiction.length, 5);
            }
        }
        // ContractExpiration (Timestamp) 
        {
            // IsInternalType
            const err = this.contract_expiration.Validate();
            if (err)
                return sprintf_js_1.sprintf('field contract_expiration is invalid : %s', err);
        }
        // ContractURI (string) 
        {
            if (this.contract_uri.length > (Math.pow(2, 8))) {
                return sprintf_js_1.sprintf('varchar field contract_uri too long %d/%d', this.contract_uri.length, (Math.pow(2, 8)) - 1);
            }
        }
        // Issuer (Entity) 
        {
            // IsInternalType
            const err = this.issuer.Validate();
            if (err)
                return sprintf_js_1.sprintf('field issuer is invalid : %s', err);
        }
        // IssuerLogoURL (string) 
        {
            if (this.issuer_logo_url.length > (Math.pow(2, 8))) {
                return sprintf_js_1.sprintf('varchar field issuer_logo_url too long %d/%d', this.issuer_logo_url.length, (Math.pow(2, 8)) - 1);
            }
        }
        // ContractOperatorIncluded (bool) 
        {
        }
        // ContractOperator (Entity)
        // IncludeIfTrue
        if (this.contract_operator_included) {
            // IsInternalType
            const err = this.contract_operator.Validate();
            if (err)
                return sprintf_js_1.sprintf('field contract_operator is invalid : %s', err);
        }
        // ContractFee (uint64) 
        {
        }
        // VotingSystems ([]VotingSystem) 
        {
            if (this.voting_systems.length > (Math.pow(2, 8)) - 1) {
                return sprintf_js_1.sprintf('list field voting_systems has too many items %d/%d', this.voting_systems.length, (Math.pow(2, 8)) - 1);
            }
            const err = this.voting_systems.find((value, i) => {
                const err2 = value.Validate();
                if (err2)
                    return sprintf_js_1.sprintf('list field voting_systems[%d] is invalid : %s', i, err2);
            });
            if (err)
                return err;
        }
        // ContractAuthFlags ([]byte) 
        {
            if (this.contract_auth_flags.length >= (Math.pow(2, 16))) {
                return sprintf_js_1.sprintf('varbin field contract_auth_flags too long %d/%d', this.contract_auth_flags.length, (Math.pow(2, 16)) - 1);
            }
        }
        // RestrictedQtyAssets (uint64) 
        {
        }
        // AdministrationProposal (bool) 
        {
        }
        // HolderProposal (bool) 
        {
        }
        // Oracles ([]Oracle) 
        {
            if (this.oracles.length > (Math.pow(2, 8)) - 1) {
                return sprintf_js_1.sprintf('list field oracles has too many items %d/%d', this.oracles.length, (Math.pow(2, 8)) - 1);
            }
            const err = this.oracles.find((value, i) => {
                const err2 = value.Validate();
                if (err2)
                    return sprintf_js_1.sprintf('list field oracles[%d] is invalid : %s', i, err2);
            });
            if (err)
                return err;
        }
        // MasterPKH (PublicKeyHash) 
        {
            // IsInternalType
            const err = this.master_pkh.Validate();
            if (err)
                return sprintf_js_1.sprintf('field master_pkh is invalid : %s', err);
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('contract_name:%s', this.contract_name.toString()));
        vals.push(sprintf_js_1.sprintf('body_of_agreement_type:%d', this.body_of_agreement_type));
        vals.push(sprintf_js_1.sprintf('body_of_agreement:%s', this.body_of_agreement.toString()));
        vals.push(sprintf_js_1.sprintf('contract_type:%s', this.contract_type.toString()));
        vals.push(sprintf_js_1.sprintf('supporting_docs:%s', this.supporting_docs.toString()));
        vals.push(sprintf_js_1.sprintf('governing_law:%s', this.governing_law.toString()));
        vals.push(sprintf_js_1.sprintf('jurisdiction:%s', this.jurisdiction.toString()));
        vals.push(sprintf_js_1.sprintf('contract_expiration:%s', this.contract_expiration.toString()));
        vals.push(sprintf_js_1.sprintf('contract_uri:%s', this.contract_uri.toString()));
        vals.push(sprintf_js_1.sprintf('issuer:%s', this.issuer.toString()));
        vals.push(sprintf_js_1.sprintf('issuer_logo_url:%s', this.issuer_logo_url.toString()));
        vals.push(sprintf_js_1.sprintf('contract_operator_included:%s', this.contract_operator_included.toString()));
        vals.push(sprintf_js_1.sprintf('contract_operator:%s', this.contract_operator.toString()));
        vals.push(sprintf_js_1.sprintf('contract_fee:%d', this.contract_fee));
        vals.push(sprintf_js_1.sprintf('voting_systems:%s', this.voting_systems.toString()));
        vals.push(sprintf_js_1.sprintf('contract_auth_flags:%s', this.contract_auth_flags.toString()));
        vals.push(sprintf_js_1.sprintf('restricted_qty_assets:%d', this.restricted_qty_assets));
        vals.push(sprintf_js_1.sprintf('administration_proposal:%s', this.administration_proposal.toString()));
        vals.push(sprintf_js_1.sprintf('holder_proposal:%s', this.holder_proposal.toString()));
        vals.push(sprintf_js_1.sprintf('oracles:%s', this.oracles.toString()));
        vals.push(sprintf_js_1.sprintf('master_pkh:%s', this.master_pkh.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for ContractOffer.', e);
            throw e;
        }
        this.contract_name = parsed.contractName;
        this.body_of_agreement_type = parsed.bodyOfAgreementType;
        // this.body_of_agreement = parsed.bodyOfAgreement (varbin)
        this.contract_type = parsed.contractType;
        // this.supporting_docs (Document[])
        // this.governing_law = parsed.governingLaw (fixedchar)
        // this.jurisdiction = parsed.jurisdiction (fixedchar)
        this.contract_expiration = new protocol_types_1.Timestamp();
        this.contract_expiration.fromJSON(JSON.stringify(parsed.contractExpiration));
        this.contract_uri = parsed.contractURI;
        this.issuer = new field_types_1.Entity();
        this.issuer.fromJSON(JSON.stringify(parsed.issuer));
        this.issuer_logo_url = parsed.issuerLogoURL;
        // this.contract_operator_included = parsed.contractOperatorIncluded (bool)
        this.contract_operator = new field_types_1.Entity();
        this.contract_operator.fromJSON(JSON.stringify(parsed.contractOperator));
        this.contract_fee = parsed.contractFee;
        // this.voting_systems (VotingSystem[])
        // this.contract_auth_flags = parsed.contractAuthFlags (varbin)
        this.restricted_qty_assets = parsed.restrictedQtyAssets;
        // this.administration_proposal = parsed.administrationProposal (bool)
        // this.holder_proposal = parsed.holderProposal (bool)
        // this.oracles (Oracle[])
        this.master_pkh = new protocol_types_1.PublicKeyHash();
        this.master_pkh.fromJSON(JSON.stringify(parsed.masterPKH));
        return this.Validate();
    }
}
exports.ContractOffer = ContractOffer;
// ContractFormation This txn is created by the contract (smart
// contract/off-chain agent/token contract) upon receipt of a valid
// Contract Offer Action from the administration. The smart contract will
// execute on a server controlled by the administration, or a smart
// contract operator on their behalf.
class ContractFormation extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeContractFormation;
        this.typeStr = 'ContractFormation';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeContractFormation;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // contract_name (string)
        {
            bytes_1.WriteVarChar(buf, this.contract_name, 8);
        }
        // body_of_agreement_type (uint8)
        {
            bytes_1.write(buf, this.body_of_agreement_type, 'uint8');
        }
        // body_of_agreement ([]byte)
        {
            bytes_1.WriteVarBin(buf, this.body_of_agreement, 32);
        }
        // contract_type (string)
        {
            bytes_1.WriteVarChar(buf, this.contract_type, 8);
        }
        // supporting_docs ([]Document)
        {
            bytes_1.WriteVariableSize(buf, this.supporting_docs.length, 8, 8);
            this.supporting_docs.forEach((value) => {
                buf.write(value.Serialize());
            });
        }
        // governing_law (string)
        {
            bytes_1.WriteFixedChar(buf, this.governing_law, 5);
        }
        // jurisdiction (string)
        {
            bytes_1.WriteFixedChar(buf, this.jurisdiction, 5);
        }
        // contract_expiration (Timestamp)
        {
            buf.write(this.contract_expiration.Serialize());
        }
        // contract_uri (string)
        {
            bytes_1.WriteVarChar(buf, this.contract_uri, 8);
        }
        // issuer (Entity)
        {
            buf.write(this.issuer.Serialize());
        }
        // issuer_logo_url (string)
        {
            bytes_1.WriteVarChar(buf, this.issuer_logo_url, 8);
        }
        // contract_operator_included (bool)
        {
            bytes_1.write(buf, this.contract_operator_included, 'bool');
        }
        // contract_operator (Entity)
        if (this.contract_operator_included) {
            buf.write(this.contract_operator.Serialize());
        }
        // contract_fee (uint64)
        {
            bytes_1.write(buf, this.contract_fee, 'uint64');
        }
        // voting_systems ([]VotingSystem)
        {
            bytes_1.WriteVariableSize(buf, this.voting_systems.length, 8, 8);
            this.voting_systems.forEach((value) => {
                buf.write(value.Serialize());
            });
        }
        // contract_auth_flags ([]byte)
        {
            bytes_1.WriteVarBin(buf, this.contract_auth_flags, 16);
        }
        // restricted_qty_assets (uint64)
        {
            bytes_1.write(buf, this.restricted_qty_assets, 'uint64');
        }
        // administration_proposal (bool)
        {
            bytes_1.write(buf, this.administration_proposal, 'bool');
        }
        // holder_proposal (bool)
        {
            bytes_1.write(buf, this.holder_proposal, 'bool');
        }
        // oracles ([]Oracle)
        {
            bytes_1.WriteVariableSize(buf, this.oracles.length, 8, 8);
            this.oracles.forEach((value) => {
                buf.write(value.Serialize());
            });
        }
        // master_pkh (PublicKeyHash)
        {
            buf.write(this.master_pkh.Serialize());
        }
        // contract_revision (uint32)
        {
            bytes_1.write(buf, this.contract_revision, 'uint32');
        }
        // timestamp (Timestamp)
        {
            buf.write(this.timestamp.Serialize());
        }
        return buf.buf;
    }
    // write populates the fields in ContractFormation from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // contract_name (string)
        {
            this.contract_name = bytes_1.ReadVarChar(buf, 8);
        }
        // body_of_agreement_type (uint8)
        {
            this.body_of_agreement_type = bytes_1.read(buf, 'uint8');
        }
        // body_of_agreement ([]byte)
        {
            this.body_of_agreement = bytes_1.ReadVarBin(buf, 32);
        }
        // contract_type (string)
        {
            this.contract_type = bytes_1.ReadVarChar(buf, 8);
        }
        // supporting_docs ([]Document)
        {
            // IsInternalTypeArray
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            this.supporting_docs = [];
            for (let i = 0; i < size; i++) {
                const newValue = new field_types_1.Document();
                newValue.Write(buf);
                this.supporting_docs.push(newValue);
            }
        }
        // governing_law (string)
        {
            this.governing_law = bytes_1.ReadFixedChar(buf, 5);
        }
        // jurisdiction (string)
        {
            this.jurisdiction = bytes_1.ReadFixedChar(buf, 5);
        }
        // contract_expiration (Timestamp)
        {
            this.contract_expiration = new protocol_types_1.Timestamp();
            this.contract_expiration.Write(buf);
        }
        // contract_uri (string)
        {
            this.contract_uri = bytes_1.ReadVarChar(buf, 8);
        }
        // issuer (Entity)
        {
            this.issuer = new field_types_1.Entity();
            this.issuer.Write(buf);
        }
        // issuer_logo_url (string)
        {
            this.issuer_logo_url = bytes_1.ReadVarChar(buf, 8);
        }
        // contract_operator_included (bool)
        {
            this.contract_operator_included = bytes_1.read(buf, 'bool');
        }
        // contract_operator (Entity)
        if (this.contract_operator_included) {
            this.contract_operator = new field_types_1.Entity();
            this.contract_operator.Write(buf);
        }
        // contract_fee (uint64)
        {
            this.contract_fee = bytes_1.read(buf, 'uint64');
        }
        // voting_systems ([]VotingSystem)
        {
            // IsInternalTypeArray
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            this.voting_systems = [];
            for (let i = 0; i < size; i++) {
                const newValue = new field_types_1.VotingSystem();
                newValue.Write(buf);
                this.voting_systems.push(newValue);
            }
        }
        // contract_auth_flags ([]byte)
        {
            this.contract_auth_flags = bytes_1.ReadVarBin(buf, 16);
        }
        // restricted_qty_assets (uint64)
        {
            this.restricted_qty_assets = bytes_1.read(buf, 'uint64');
        }
        // administration_proposal (bool)
        {
            this.administration_proposal = bytes_1.read(buf, 'bool');
        }
        // holder_proposal (bool)
        {
            this.holder_proposal = bytes_1.read(buf, 'bool');
        }
        // oracles ([]Oracle)
        {
            // IsInternalTypeArray
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            this.oracles = [];
            for (let i = 0; i < size; i++) {
                const newValue = new field_types_1.Oracle();
                newValue.Write(buf);
                this.oracles.push(newValue);
            }
        }
        // master_pkh (PublicKeyHash)
        {
            this.master_pkh = new protocol_types_1.PublicKeyHash();
            this.master_pkh.Write(buf);
        }
        // contract_revision (uint32)
        {
            this.contract_revision = bytes_1.read(buf, 'uint32');
        }
        // timestamp (Timestamp)
        {
            this.timestamp = new protocol_types_1.Timestamp();
            this.timestamp.Write(buf);
        }
        return b.length - buf.length;
    }
    Validate() {
        // ContractName (string) 
        {
            if (this.contract_name.length > (Math.pow(2, 8))) {
                return sprintf_js_1.sprintf('varchar field contract_name too long %d/%d', this.contract_name.length, (Math.pow(2, 8)) - 1);
            }
        }
        // BodyOfAgreementType (uint8) 
        {
            // $field.IntValues [1 2]
            if (this.body_of_agreement_type !== 1 && this.body_of_agreement_type !== 2) {
                return sprintf_js_1.sprintf('field body_of_agreement_type value is invalid : %d', this.body_of_agreement_type);
            }
        }
        // BodyOfAgreement ([]byte) 
        {
            if (this.body_of_agreement.length >= (Math.pow(2, 32))) {
                return sprintf_js_1.sprintf('varbin field body_of_agreement too long %d/%d', this.body_of_agreement.length, (Math.pow(2, 32)) - 1);
            }
        }
        // ContractType (string) 
        {
            if (this.contract_type.length > (Math.pow(2, 8))) {
                return sprintf_js_1.sprintf('varchar field contract_type too long %d/%d', this.contract_type.length, (Math.pow(2, 8)) - 1);
            }
        }
        // SupportingDocs ([]Document) 
        {
            if (this.supporting_docs.length > (Math.pow(2, 8)) - 1) {
                return sprintf_js_1.sprintf('list field supporting_docs has too many items %d/%d', this.supporting_docs.length, (Math.pow(2, 8)) - 1);
            }
            const err = this.supporting_docs.find((value, i) => {
                const err2 = value.Validate();
                if (err2)
                    return sprintf_js_1.sprintf('list field supporting_docs[%d] is invalid : %s', i, err2);
            });
            if (err)
                return err;
        }
        // GoverningLaw (string) 
        {
            if (this.governing_law.length > 5) {
                return sprintf_js_1.sprintf('fixedchar field governing_law too long %d/%d', this.governing_law.length, 5);
            }
        }
        // Jurisdiction (string) 
        {
            if (this.jurisdiction.length > 5) {
                return sprintf_js_1.sprintf('fixedchar field jurisdiction too long %d/%d', this.jurisdiction.length, 5);
            }
        }
        // ContractExpiration (Timestamp) 
        {
            // IsInternalType
            const err = this.contract_expiration.Validate();
            if (err)
                return sprintf_js_1.sprintf('field contract_expiration is invalid : %s', err);
        }
        // ContractURI (string) 
        {
            if (this.contract_uri.length > (Math.pow(2, 8))) {
                return sprintf_js_1.sprintf('varchar field contract_uri too long %d/%d', this.contract_uri.length, (Math.pow(2, 8)) - 1);
            }
        }
        // Issuer (Entity) 
        {
            // IsInternalType
            const err = this.issuer.Validate();
            if (err)
                return sprintf_js_1.sprintf('field issuer is invalid : %s', err);
        }
        // IssuerLogoURL (string) 
        {
            if (this.issuer_logo_url.length > (Math.pow(2, 8))) {
                return sprintf_js_1.sprintf('varchar field issuer_logo_url too long %d/%d', this.issuer_logo_url.length, (Math.pow(2, 8)) - 1);
            }
        }
        // ContractOperatorIncluded (bool) 
        {
        }
        // ContractOperator (Entity)
        // IncludeIfTrue
        if (this.contract_operator_included) {
            // IsInternalType
            const err = this.contract_operator.Validate();
            if (err)
                return sprintf_js_1.sprintf('field contract_operator is invalid : %s', err);
        }
        // ContractFee (uint64) 
        {
        }
        // VotingSystems ([]VotingSystem) 
        {
            if (this.voting_systems.length > (Math.pow(2, 8)) - 1) {
                return sprintf_js_1.sprintf('list field voting_systems has too many items %d/%d', this.voting_systems.length, (Math.pow(2, 8)) - 1);
            }
            const err = this.voting_systems.find((value, i) => {
                const err2 = value.Validate();
                if (err2)
                    return sprintf_js_1.sprintf('list field voting_systems[%d] is invalid : %s', i, err2);
            });
            if (err)
                return err;
        }
        // ContractAuthFlags ([]byte) 
        {
            if (this.contract_auth_flags.length >= (Math.pow(2, 16))) {
                return sprintf_js_1.sprintf('varbin field contract_auth_flags too long %d/%d', this.contract_auth_flags.length, (Math.pow(2, 16)) - 1);
            }
        }
        // RestrictedQtyAssets (uint64) 
        {
        }
        // AdministrationProposal (bool) 
        {
        }
        // HolderProposal (bool) 
        {
        }
        // Oracles ([]Oracle) 
        {
            if (this.oracles.length > (Math.pow(2, 8)) - 1) {
                return sprintf_js_1.sprintf('list field oracles has too many items %d/%d', this.oracles.length, (Math.pow(2, 8)) - 1);
            }
            const err = this.oracles.find((value, i) => {
                const err2 = value.Validate();
                if (err2)
                    return sprintf_js_1.sprintf('list field oracles[%d] is invalid : %s', i, err2);
            });
            if (err)
                return err;
        }
        // MasterPKH (PublicKeyHash) 
        {
            // IsInternalType
            const err = this.master_pkh.Validate();
            if (err)
                return sprintf_js_1.sprintf('field master_pkh is invalid : %s', err);
        }
        // ContractRevision (uint32) 
        {
        }
        // Timestamp (Timestamp) 
        {
            // IsInternalType
            const err = this.timestamp.Validate();
            if (err)
                return sprintf_js_1.sprintf('field timestamp is invalid : %s', err);
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('contract_name:%s', this.contract_name.toString()));
        vals.push(sprintf_js_1.sprintf('body_of_agreement_type:%d', this.body_of_agreement_type));
        vals.push(sprintf_js_1.sprintf('body_of_agreement:%s', this.body_of_agreement.toString()));
        vals.push(sprintf_js_1.sprintf('contract_type:%s', this.contract_type.toString()));
        vals.push(sprintf_js_1.sprintf('supporting_docs:%s', this.supporting_docs.toString()));
        vals.push(sprintf_js_1.sprintf('governing_law:%s', this.governing_law.toString()));
        vals.push(sprintf_js_1.sprintf('jurisdiction:%s', this.jurisdiction.toString()));
        vals.push(sprintf_js_1.sprintf('contract_expiration:%s', this.contract_expiration.toString()));
        vals.push(sprintf_js_1.sprintf('contract_uri:%s', this.contract_uri.toString()));
        vals.push(sprintf_js_1.sprintf('issuer:%s', this.issuer.toString()));
        vals.push(sprintf_js_1.sprintf('issuer_logo_url:%s', this.issuer_logo_url.toString()));
        vals.push(sprintf_js_1.sprintf('contract_operator_included:%s', this.contract_operator_included.toString()));
        vals.push(sprintf_js_1.sprintf('contract_operator:%s', this.contract_operator.toString()));
        vals.push(sprintf_js_1.sprintf('contract_fee:%d', this.contract_fee));
        vals.push(sprintf_js_1.sprintf('voting_systems:%s', this.voting_systems.toString()));
        vals.push(sprintf_js_1.sprintf('contract_auth_flags:%s', this.contract_auth_flags.toString()));
        vals.push(sprintf_js_1.sprintf('restricted_qty_assets:%d', this.restricted_qty_assets));
        vals.push(sprintf_js_1.sprintf('administration_proposal:%s', this.administration_proposal.toString()));
        vals.push(sprintf_js_1.sprintf('holder_proposal:%s', this.holder_proposal.toString()));
        vals.push(sprintf_js_1.sprintf('oracles:%s', this.oracles.toString()));
        vals.push(sprintf_js_1.sprintf('master_pkh:%s', this.master_pkh.toString()));
        vals.push(sprintf_js_1.sprintf('contract_revision:%d', this.contract_revision));
        vals.push(sprintf_js_1.sprintf('timestamp:%s', this.timestamp.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for ContractFormation.', e);
            throw e;
        }
        this.contract_name = parsed.contractName;
        this.body_of_agreement_type = parsed.bodyOfAgreementType;
        // this.body_of_agreement = parsed.bodyOfAgreement (varbin)
        this.contract_type = parsed.contractType;
        // this.supporting_docs (Document[])
        // this.governing_law = parsed.governingLaw (fixedchar)
        // this.jurisdiction = parsed.jurisdiction (fixedchar)
        this.contract_expiration = new protocol_types_1.Timestamp();
        this.contract_expiration.fromJSON(JSON.stringify(parsed.contractExpiration));
        this.contract_uri = parsed.contractURI;
        this.issuer = new field_types_1.Entity();
        this.issuer.fromJSON(JSON.stringify(parsed.issuer));
        this.issuer_logo_url = parsed.issuerLogoURL;
        // this.contract_operator_included = parsed.contractOperatorIncluded (bool)
        this.contract_operator = new field_types_1.Entity();
        this.contract_operator.fromJSON(JSON.stringify(parsed.contractOperator));
        this.contract_fee = parsed.contractFee;
        // this.voting_systems (VotingSystem[])
        // this.contract_auth_flags = parsed.contractAuthFlags (varbin)
        this.restricted_qty_assets = parsed.restrictedQtyAssets;
        // this.administration_proposal = parsed.administrationProposal (bool)
        // this.holder_proposal = parsed.holderProposal (bool)
        // this.oracles (Oracle[])
        this.master_pkh = new protocol_types_1.PublicKeyHash();
        this.master_pkh.fromJSON(JSON.stringify(parsed.masterPKH));
        this.contract_revision = parsed.contractRevision;
        this.timestamp = new protocol_types_1.Timestamp();
        this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
        return this.Validate();
    }
}
exports.ContractFormation = ContractFormation;
// ContractAmendment The administration can initiate an amendment to the
// contract establishment metadata. The ability to make an amendment to the
// contract is restricted by the Authorization Flag set on the current
// revision of Contract Formation action.
class ContractAmendment extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeContractAmendment;
        this.typeStr = 'ContractAmendment';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeContractAmendment;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // change_administration_address (bool)
        {
            bytes_1.write(buf, this.change_administration_address, 'bool');
        }
        // change_operator_address (bool)
        {
            bytes_1.write(buf, this.change_operator_address, 'bool');
        }
        // contract_revision (uint32)
        {
            bytes_1.write(buf, this.contract_revision, 'uint32');
        }
        // amendments ([]Amendment)
        {
            bytes_1.WriteVariableSize(buf, this.amendments.length, 8, 8);
            this.amendments.forEach((value) => {
                buf.write(value.Serialize());
            });
        }
        // ref_tx_id (TxId)
        {
            buf.write(this.ref_tx_id.Serialize());
        }
        return buf.buf;
    }
    // write populates the fields in ContractAmendment from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // change_administration_address (bool)
        {
            this.change_administration_address = bytes_1.read(buf, 'bool');
        }
        // change_operator_address (bool)
        {
            this.change_operator_address = bytes_1.read(buf, 'bool');
        }
        // contract_revision (uint32)
        {
            this.contract_revision = bytes_1.read(buf, 'uint32');
        }
        // amendments ([]Amendment)
        {
            // IsInternalTypeArray
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            this.amendments = [];
            for (let i = 0; i < size; i++) {
                const newValue = new field_types_1.Amendment();
                newValue.Write(buf);
                this.amendments.push(newValue);
            }
        }
        // ref_tx_id (TxId)
        {
            this.ref_tx_id = new protocol_types_1.TxId();
            this.ref_tx_id.Write(buf);
        }
        return b.length - buf.length;
    }
    Validate() {
        // ChangeAdministrationAddress (bool) 
        {
        }
        // ChangeOperatorAddress (bool) 
        {
        }
        // ContractRevision (uint32) 
        {
        }
        // Amendments ([]Amendment) 
        {
            if (this.amendments.length > (Math.pow(2, 8)) - 1) {
                return sprintf_js_1.sprintf('list field amendments has too many items %d/%d', this.amendments.length, (Math.pow(2, 8)) - 1);
            }
            const err = this.amendments.find((value, i) => {
                const err2 = value.Validate();
                if (err2)
                    return sprintf_js_1.sprintf('list field amendments[%d] is invalid : %s', i, err2);
            });
            if (err)
                return err;
        }
        // RefTxID (TxId) 
        {
            // IsInternalType
            const err = this.ref_tx_id.Validate();
            if (err)
                return sprintf_js_1.sprintf('field ref_tx_id is invalid : %s', err);
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('change_administration_address:%s', this.change_administration_address.toString()));
        vals.push(sprintf_js_1.sprintf('change_operator_address:%s', this.change_operator_address.toString()));
        vals.push(sprintf_js_1.sprintf('contract_revision:%d', this.contract_revision));
        vals.push(sprintf_js_1.sprintf('amendments:%s', this.amendments.toString()));
        vals.push(sprintf_js_1.sprintf('ref_tx_id:%s', this.ref_tx_id.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for ContractAmendment.', e);
            throw e;
        }
        // this.change_administration_address = parsed.changeAdministrationAddress (bool)
        // this.change_operator_address = parsed.changeOperatorAddress (bool)
        this.contract_revision = parsed.contractRevision;
        // this.amendments (Amendment[])
        this.ref_tx_id = new protocol_types_1.TxId();
        this.ref_tx_id.fromJSON(JSON.stringify(parsed.refTxID));
        return this.Validate();
    }
}
exports.ContractAmendment = ContractAmendment;
// StaticContractFormation Static Contract Formation Action
class StaticContractFormation extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeStaticContractFormation;
        this.typeStr = 'StaticContractFormation';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeStaticContractFormation;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // contract_name (string)
        {
            bytes_1.WriteVarChar(buf, this.contract_name, 8);
        }
        // contract_code (ContractCode)
        {
            buf.write(this.contract_code.Serialize());
        }
        // body_of_agreement_type (uint8)
        {
            bytes_1.write(buf, this.body_of_agreement_type, 'uint8');
        }
        // body_of_agreement ([]byte)
        {
            bytes_1.WriteVarBin(buf, this.body_of_agreement, 32);
        }
        // contract_type (string)
        {
            bytes_1.WriteVarChar(buf, this.contract_type, 8);
        }
        // supporting_docs ([]Document)
        {
            bytes_1.WriteVariableSize(buf, this.supporting_docs.length, 8, 8);
            this.supporting_docs.forEach((value) => {
                buf.write(value.Serialize());
            });
        }
        // contract_revision (uint32)
        {
            bytes_1.write(buf, this.contract_revision, 'uint32');
        }
        // governing_law (string)
        {
            bytes_1.WriteFixedChar(buf, this.governing_law, 5);
        }
        // jurisdiction (string)
        {
            bytes_1.WriteFixedChar(buf, this.jurisdiction, 5);
        }
        // effective_date (Timestamp)
        {
            buf.write(this.effective_date.Serialize());
        }
        // contract_expiration (Timestamp)
        {
            buf.write(this.contract_expiration.Serialize());
        }
        // contract_uri (string)
        {
            bytes_1.WriteVarChar(buf, this.contract_uri, 8);
        }
        // prev_rev_tx_id (TxId)
        {
            buf.write(this.prev_rev_tx_id.Serialize());
        }
        // entities ([]Entity)
        {
            bytes_1.WriteVariableSize(buf, this.entities.length, 8, 8);
            this.entities.forEach((value) => {
                buf.write(value.Serialize());
            });
        }
        return buf.buf;
    }
    // write populates the fields in StaticContractFormation from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // contract_name (string)
        {
            this.contract_name = bytes_1.ReadVarChar(buf, 8);
        }
        // contract_code (ContractCode)
        {
            this.contract_code = new protocol_types_1.ContractCode();
            this.contract_code.Write(buf);
        }
        // body_of_agreement_type (uint8)
        {
            this.body_of_agreement_type = bytes_1.read(buf, 'uint8');
        }
        // body_of_agreement ([]byte)
        {
            this.body_of_agreement = bytes_1.ReadVarBin(buf, 32);
        }
        // contract_type (string)
        {
            this.contract_type = bytes_1.ReadVarChar(buf, 8);
        }
        // supporting_docs ([]Document)
        {
            // IsInternalTypeArray
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            this.supporting_docs = [];
            for (let i = 0; i < size; i++) {
                const newValue = new field_types_1.Document();
                newValue.Write(buf);
                this.supporting_docs.push(newValue);
            }
        }
        // contract_revision (uint32)
        {
            this.contract_revision = bytes_1.read(buf, 'uint32');
        }
        // governing_law (string)
        {
            this.governing_law = bytes_1.ReadFixedChar(buf, 5);
        }
        // jurisdiction (string)
        {
            this.jurisdiction = bytes_1.ReadFixedChar(buf, 5);
        }
        // effective_date (Timestamp)
        {
            this.effective_date = new protocol_types_1.Timestamp();
            this.effective_date.Write(buf);
        }
        // contract_expiration (Timestamp)
        {
            this.contract_expiration = new protocol_types_1.Timestamp();
            this.contract_expiration.Write(buf);
        }
        // contract_uri (string)
        {
            this.contract_uri = bytes_1.ReadVarChar(buf, 8);
        }
        // prev_rev_tx_id (TxId)
        {
            this.prev_rev_tx_id = new protocol_types_1.TxId();
            this.prev_rev_tx_id.Write(buf);
        }
        // entities ([]Entity)
        {
            // IsInternalTypeArray
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            this.entities = [];
            for (let i = 0; i < size; i++) {
                const newValue = new field_types_1.Entity();
                newValue.Write(buf);
                this.entities.push(newValue);
            }
        }
        return b.length - buf.length;
    }
    Validate() {
        // ContractName (string) 
        {
            if (this.contract_name.length > (Math.pow(2, 8))) {
                return sprintf_js_1.sprintf('varchar field contract_name too long %d/%d', this.contract_name.length, (Math.pow(2, 8)) - 1);
            }
        }
        // ContractCode (ContractCode) 
        {
            // IsInternalType
            const err = this.contract_code.Validate();
            if (err)
                return sprintf_js_1.sprintf('field contract_code is invalid : %s', err);
        }
        // BodyOfAgreementType (uint8) 
        {
            // $field.IntValues [1 2]
            if (this.body_of_agreement_type !== 1 && this.body_of_agreement_type !== 2) {
                return sprintf_js_1.sprintf('field body_of_agreement_type value is invalid : %d', this.body_of_agreement_type);
            }
        }
        // BodyOfAgreement ([]byte) 
        {
            if (this.body_of_agreement.length >= (Math.pow(2, 32))) {
                return sprintf_js_1.sprintf('varbin field body_of_agreement too long %d/%d', this.body_of_agreement.length, (Math.pow(2, 32)) - 1);
            }
        }
        // ContractType (string) 
        {
            if (this.contract_type.length > (Math.pow(2, 8))) {
                return sprintf_js_1.sprintf('varchar field contract_type too long %d/%d', this.contract_type.length, (Math.pow(2, 8)) - 1);
            }
        }
        // SupportingDocs ([]Document) 
        {
            if (this.supporting_docs.length > (Math.pow(2, 8)) - 1) {
                return sprintf_js_1.sprintf('list field supporting_docs has too many items %d/%d', this.supporting_docs.length, (Math.pow(2, 8)) - 1);
            }
            const err = this.supporting_docs.find((value, i) => {
                const err2 = value.Validate();
                if (err2)
                    return sprintf_js_1.sprintf('list field supporting_docs[%d] is invalid : %s', i, err2);
            });
            if (err)
                return err;
        }
        // ContractRevision (uint32) 
        {
        }
        // GoverningLaw (string) 
        {
            if (this.governing_law.length > 5) {
                return sprintf_js_1.sprintf('fixedchar field governing_law too long %d/%d', this.governing_law.length, 5);
            }
        }
        // Jurisdiction (string) 
        {
            if (this.jurisdiction.length > 5) {
                return sprintf_js_1.sprintf('fixedchar field jurisdiction too long %d/%d', this.jurisdiction.length, 5);
            }
        }
        // EffectiveDate (Timestamp) 
        {
            // IsInternalType
            const err = this.effective_date.Validate();
            if (err)
                return sprintf_js_1.sprintf('field effective_date is invalid : %s', err);
        }
        // ContractExpiration (Timestamp) 
        {
            // IsInternalType
            const err = this.contract_expiration.Validate();
            if (err)
                return sprintf_js_1.sprintf('field contract_expiration is invalid : %s', err);
        }
        // ContractURI (string) 
        {
            if (this.contract_uri.length > (Math.pow(2, 8))) {
                return sprintf_js_1.sprintf('varchar field contract_uri too long %d/%d', this.contract_uri.length, (Math.pow(2, 8)) - 1);
            }
        }
        // PrevRevTxID (TxId) 
        {
            // IsInternalType
            const err = this.prev_rev_tx_id.Validate();
            if (err)
                return sprintf_js_1.sprintf('field prev_rev_tx_id is invalid : %s', err);
        }
        // Entities ([]Entity) 
        {
            if (this.entities.length > (Math.pow(2, 8)) - 1) {
                return sprintf_js_1.sprintf('list field entities has too many items %d/%d', this.entities.length, (Math.pow(2, 8)) - 1);
            }
            const err = this.entities.find((value, i) => {
                const err2 = value.Validate();
                if (err2)
                    return sprintf_js_1.sprintf('list field entities[%d] is invalid : %s', i, err2);
            });
            if (err)
                return err;
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('contract_name:%s', this.contract_name.toString()));
        vals.push(sprintf_js_1.sprintf('contract_code:%s', this.contract_code.toString()));
        vals.push(sprintf_js_1.sprintf('body_of_agreement_type:%d', this.body_of_agreement_type));
        vals.push(sprintf_js_1.sprintf('body_of_agreement:%s', this.body_of_agreement.toString()));
        vals.push(sprintf_js_1.sprintf('contract_type:%s', this.contract_type.toString()));
        vals.push(sprintf_js_1.sprintf('supporting_docs:%s', this.supporting_docs.toString()));
        vals.push(sprintf_js_1.sprintf('contract_revision:%d', this.contract_revision));
        vals.push(sprintf_js_1.sprintf('governing_law:%s', this.governing_law.toString()));
        vals.push(sprintf_js_1.sprintf('jurisdiction:%s', this.jurisdiction.toString()));
        vals.push(sprintf_js_1.sprintf('effective_date:%s', this.effective_date.toString()));
        vals.push(sprintf_js_1.sprintf('contract_expiration:%s', this.contract_expiration.toString()));
        vals.push(sprintf_js_1.sprintf('contract_uri:%s', this.contract_uri.toString()));
        vals.push(sprintf_js_1.sprintf('prev_rev_tx_id:%s', this.prev_rev_tx_id.toString()));
        vals.push(sprintf_js_1.sprintf('entities:%s', this.entities.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for StaticContractFormation.', e);
            throw e;
        }
        this.contract_name = parsed.contractName;
        this.contract_code = new protocol_types_1.ContractCode();
        this.contract_code.fromJSON(JSON.stringify(parsed.contractCode));
        this.body_of_agreement_type = parsed.bodyOfAgreementType;
        // this.body_of_agreement = parsed.bodyOfAgreement (varbin)
        this.contract_type = parsed.contractType;
        // this.supporting_docs (Document[])
        this.contract_revision = parsed.contractRevision;
        // this.governing_law = parsed.governingLaw (fixedchar)
        // this.jurisdiction = parsed.jurisdiction (fixedchar)
        this.effective_date = new protocol_types_1.Timestamp();
        this.effective_date.fromJSON(JSON.stringify(parsed.effectiveDate));
        this.contract_expiration = new protocol_types_1.Timestamp();
        this.contract_expiration.fromJSON(JSON.stringify(parsed.contractExpiration));
        this.contract_uri = parsed.contractURI;
        this.prev_rev_tx_id = new protocol_types_1.TxId();
        this.prev_rev_tx_id.fromJSON(JSON.stringify(parsed.prevRevTxID));
        // this.entities (Entity[])
        return this.Validate();
    }
}
exports.StaticContractFormation = StaticContractFormation;
// ContractAddressChange This txn is signed by the master contract key
// defined in the contract formation and changes the active contract
// address which the contract uses to receive and respond to requests. This
// is a worst case scenario fallback to only be used when the contract
// private key is believed to be exposed.
class ContractAddressChange extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeContractAddressChange;
        this.typeStr = 'ContractAddressChange';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeContractAddressChange;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // new_contract_pkh (PublicKeyHash)
        {
            buf.write(this.new_contract_pkh.Serialize());
        }
        // timestamp (Timestamp)
        {
            buf.write(this.timestamp.Serialize());
        }
        return buf.buf;
    }
    // write populates the fields in ContractAddressChange from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // new_contract_pkh (PublicKeyHash)
        {
            this.new_contract_pkh = new protocol_types_1.PublicKeyHash();
            this.new_contract_pkh.Write(buf);
        }
        // timestamp (Timestamp)
        {
            this.timestamp = new protocol_types_1.Timestamp();
            this.timestamp.Write(buf);
        }
        return b.length - buf.length;
    }
    Validate() {
        // NewContractPKH (PublicKeyHash) 
        {
            // IsInternalType
            const err = this.new_contract_pkh.Validate();
            if (err)
                return sprintf_js_1.sprintf('field new_contract_pkh is invalid : %s', err);
        }
        // Timestamp (Timestamp) 
        {
            // IsInternalType
            const err = this.timestamp.Validate();
            if (err)
                return sprintf_js_1.sprintf('field timestamp is invalid : %s', err);
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('new_contract_pkh:%s', this.new_contract_pkh.toString()));
        vals.push(sprintf_js_1.sprintf('timestamp:%s', this.timestamp.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for ContractAddressChange.', e);
            throw e;
        }
        this.new_contract_pkh = new protocol_types_1.PublicKeyHash();
        this.new_contract_pkh.fromJSON(JSON.stringify(parsed.newContractPKH));
        this.timestamp = new protocol_types_1.Timestamp();
        this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
        return this.Validate();
    }
}
exports.ContractAddressChange = ContractAddressChange;
// Order Used by the administration to signal to the smart contract that
// the tokens that a particular public address(es) owns are to be
// confiscated, frozen, thawed or reconciled.
class Order extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeOrder;
        this.typeStr = 'Order';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeOrder;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // compliance_action (byte)
        {
            bytes_1.write(buf, this.compliance_action, 'byte');
        }
        // asset_type (string)
        if (this.compliance_action === bytes_1.char('F') || this.compliance_action === bytes_1.char('C') || this.compliance_action === bytes_1.char('R')) {
            bytes_1.WriteFixedChar(buf, this.asset_type, 3);
        }
        // asset_code (AssetCode)
        if (this.compliance_action === bytes_1.char('F') || this.compliance_action === bytes_1.char('C') || this.compliance_action === bytes_1.char('R')) {
            buf.write(this.asset_code.Serialize());
        }
        // target_addresses ([]TargetAddress)
        if (this.compliance_action === bytes_1.char('F') || this.compliance_action === bytes_1.char('C') || this.compliance_action === bytes_1.char('R')) {
            bytes_1.WriteVariableSize(buf, this.target_addresses.length, 16, 8);
            this.target_addresses.forEach((value) => {
                buf.write(value.Serialize());
            });
        }
        // freeze_tx_id (TxId)
        if (this.compliance_action === bytes_1.char('T')) {
            buf.write(this.freeze_tx_id.Serialize());
        }
        // freeze_period (Timestamp)
        if (this.compliance_action === bytes_1.char('F')) {
            buf.write(this.freeze_period.Serialize());
        }
        // deposit_address (PublicKeyHash)
        if (this.compliance_action === bytes_1.char('C')) {
            buf.write(this.deposit_address.Serialize());
        }
        // authority_included (bool)
        if (this.compliance_action === bytes_1.char('F') || this.compliance_action === bytes_1.char('T') || this.compliance_action === bytes_1.char('C')) {
            bytes_1.write(buf, this.authority_included, 'bool');
        }
        // authority_name (string)
        if (this.authority_included) {
            bytes_1.WriteVarChar(buf, this.authority_name, 8);
        }
        // authority_public_key ([]byte)
        if (this.authority_included) {
            bytes_1.WriteVarBin(buf, this.authority_public_key, 8);
        }
        // signature_algorithm (uint8)
        if (this.authority_included) {
            bytes_1.write(buf, this.signature_algorithm, 'uint8');
        }
        // order_signature ([]byte)
        if (this.authority_included) {
            bytes_1.WriteVarBin(buf, this.order_signature, 8);
        }
        // supporting_evidence_hash ([32]byte)
        {
            bytes_1.write(buf, this.supporting_evidence_hash, '[32]byte');
        }
        // ref_txs ([]byte)
        if (this.compliance_action === bytes_1.char('R')) {
            bytes_1.WriteVarBin(buf, this.ref_txs, 32);
        }
        // bitcoin_dispersions ([]QuantityIndex)
        if (this.compliance_action === bytes_1.char('R')) {
            bytes_1.WriteVariableSize(buf, this.bitcoin_dispersions.length, 16, 8);
            this.bitcoin_dispersions.forEach((value) => {
                buf.write(value.Serialize());
            });
        }
        // message (string)
        {
            bytes_1.WriteVarChar(buf, this.message, 32);
        }
        return buf.buf;
    }
    // write populates the fields in Order from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // compliance_action (byte)
        {
            this.compliance_action = bytes_1.read(buf, 'byte');
        }
        // asset_type (string)
        if (this.compliance_action === bytes_1.char('F') || this.compliance_action === bytes_1.char('C') || this.compliance_action === bytes_1.char('R')) {
            this.asset_type = bytes_1.ReadFixedChar(buf, 3);
        }
        // asset_code (AssetCode)
        if (this.compliance_action === bytes_1.char('F') || this.compliance_action === bytes_1.char('C') || this.compliance_action === bytes_1.char('R')) {
            this.asset_code = new protocol_types_1.AssetCode();
            this.asset_code.Write(buf);
        }
        // target_addresses ([]TargetAddress)
        if (this.compliance_action === bytes_1.char('F') || this.compliance_action === bytes_1.char('C') || this.compliance_action === bytes_1.char('R')) {
            // IsInternalTypeArray
            const size = bytes_1.ReadVariableSize(buf, 16, 8);
            this.target_addresses = [];
            for (let i = 0; i < size; i++) {
                const newValue = new field_types_1.TargetAddress();
                newValue.Write(buf);
                this.target_addresses.push(newValue);
            }
        }
        // freeze_tx_id (TxId)
        if (this.compliance_action === bytes_1.char('T')) {
            this.freeze_tx_id = new protocol_types_1.TxId();
            this.freeze_tx_id.Write(buf);
        }
        // freeze_period (Timestamp)
        if (this.compliance_action === bytes_1.char('F')) {
            this.freeze_period = new protocol_types_1.Timestamp();
            this.freeze_period.Write(buf);
        }
        // deposit_address (PublicKeyHash)
        if (this.compliance_action === bytes_1.char('C')) {
            this.deposit_address = new protocol_types_1.PublicKeyHash();
            this.deposit_address.Write(buf);
        }
        // authority_included (bool)
        if (this.compliance_action === bytes_1.char('F') || this.compliance_action === bytes_1.char('T') || this.compliance_action === bytes_1.char('C')) {
            this.authority_included = bytes_1.read(buf, 'bool');
        }
        // authority_name (string)
        if (this.authority_included) {
            this.authority_name = bytes_1.ReadVarChar(buf, 8);
        }
        // authority_public_key ([]byte)
        if (this.authority_included) {
            this.authority_public_key = bytes_1.ReadVarBin(buf, 8);
        }
        // signature_algorithm (uint8)
        if (this.authority_included) {
            this.signature_algorithm = bytes_1.read(buf, 'uint8');
        }
        // order_signature ([]byte)
        if (this.authority_included) {
            this.order_signature = bytes_1.ReadVarBin(buf, 8);
        }
        // supporting_evidence_hash ([32]byte)
        {
            this.supporting_evidence_hash = bytes_1.read(buf, '[32]byte');
        }
        // ref_txs ([]byte)
        if (this.compliance_action === bytes_1.char('R')) {
            this.ref_txs = bytes_1.ReadVarBin(buf, 32);
        }
        // bitcoin_dispersions ([]QuantityIndex)
        if (this.compliance_action === bytes_1.char('R')) {
            // IsInternalTypeArray
            const size = bytes_1.ReadVariableSize(buf, 16, 8);
            this.bitcoin_dispersions = [];
            for (let i = 0; i < size; i++) {
                const newValue = new field_types_1.QuantityIndex();
                newValue.Write(buf);
                this.bitcoin_dispersions.push(newValue);
            }
        }
        // message (string)
        {
            this.message = bytes_1.ReadVarChar(buf, 32);
        }
        return b.length - buf.length;
    }
    Validate() {
        // ComplianceAction (byte) 
        {
            if (this.compliance_action !== 'F'
                && this.compliance_action !== 'T'
                && this.compliance_action !== 'C'
                && this.compliance_action !== 'R') {
                return sprintf_js_1.sprintf('field compliance_action value is invalid : %d', this.compliance_action);
            }
        }
        // AssetType (string)
        // IncludeIf.Field
        if (this.compliance_action === bytes_1.char('F') || this.compliance_action === bytes_1.char('C') || this.compliance_action === bytes_1.char('R')) {
            if (this.asset_type.length > 3) {
                return sprintf_js_1.sprintf('fixedchar field asset_type too long %d/%d', this.asset_type.length, 3);
            }
        }
        // AssetCode (AssetCode)
        // IncludeIf.Field
        if (this.compliance_action === bytes_1.char('F') || this.compliance_action === bytes_1.char('C') || this.compliance_action === bytes_1.char('R')) {
            // IsInternalType
            const err = this.asset_code.Validate();
            if (err)
                return sprintf_js_1.sprintf('field asset_code is invalid : %s', err);
        }
        // TargetAddresses ([]TargetAddress)
        // IncludeIf.Field
        if (this.compliance_action === bytes_1.char('F') || this.compliance_action === bytes_1.char('C') || this.compliance_action === bytes_1.char('R')) {
            if (this.target_addresses.length > (Math.pow(2, 16)) - 1) {
                return sprintf_js_1.sprintf('list field target_addresses has too many items %d/%d', this.target_addresses.length, (Math.pow(2, 16)) - 1);
            }
            const err = this.target_addresses.find((value, i) => {
                const err2 = value.Validate();
                if (err2)
                    return sprintf_js_1.sprintf('list field target_addresses[%d] is invalid : %s', i, err2);
            });
            if (err)
                return err;
        }
        // FreezeTxId (TxId)
        // IncludeIf.Field
        if (this.compliance_action === bytes_1.char('T')) {
            // IsInternalType
            const err = this.freeze_tx_id.Validate();
            if (err)
                return sprintf_js_1.sprintf('field freeze_tx_id is invalid : %s', err);
        }
        // FreezePeriod (Timestamp)
        // IncludeIf.Field
        if (this.compliance_action === bytes_1.char('F')) {
            // IsInternalType
            const err = this.freeze_period.Validate();
            if (err)
                return sprintf_js_1.sprintf('field freeze_period is invalid : %s', err);
        }
        // DepositAddress (PublicKeyHash)
        // IncludeIf.Field
        if (this.compliance_action === bytes_1.char('C')) {
            // IsInternalType
            const err = this.deposit_address.Validate();
            if (err)
                return sprintf_js_1.sprintf('field deposit_address is invalid : %s', err);
        }
        // AuthorityIncluded (bool)
        // IncludeIf.Field
        if (this.compliance_action === bytes_1.char('F') || this.compliance_action === bytes_1.char('T') || this.compliance_action === bytes_1.char('C')) {
        }
        // AuthorityName (string)
        // IncludeIfTrue
        if (this.authority_included) {
            if (this.authority_name.length > (Math.pow(2, 8))) {
                return sprintf_js_1.sprintf('varchar field authority_name too long %d/%d', this.authority_name.length, (Math.pow(2, 8)) - 1);
            }
        }
        // AuthorityPublicKey ([]byte)
        // IncludeIfTrue
        if (this.authority_included) {
            if (this.authority_public_key.length >= (Math.pow(2, 8))) {
                return sprintf_js_1.sprintf('varbin field authority_public_key too long %d/%d', this.authority_public_key.length, (Math.pow(2, 8)) - 1);
            }
        }
        // SignatureAlgorithm (uint8)
        // IncludeIfTrue
        if (this.authority_included) {
            // $field.IntValues [1]
            if (this.signature_algorithm !== 1) {
                return sprintf_js_1.sprintf('field signature_algorithm value is invalid : %d', this.signature_algorithm);
            }
        }
        // OrderSignature ([]byte)
        // IncludeIfTrue
        if (this.authority_included) {
            if (this.order_signature.length >= (Math.pow(2, 8))) {
                return sprintf_js_1.sprintf('varbin field order_signature too long %d/%d', this.order_signature.length, (Math.pow(2, 8)) - 1);
            }
        }
        // SupportingEvidenceHash ([32]byte) 
        {
        }
        // RefTxs ([]byte)
        // IncludeIf.Field
        if (this.compliance_action === bytes_1.char('R')) {
            if (this.ref_txs.length >= (Math.pow(2, 32))) {
                return sprintf_js_1.sprintf('varbin field ref_txs too long %d/%d', this.ref_txs.length, (Math.pow(2, 32)) - 1);
            }
        }
        // BitcoinDispersions ([]QuantityIndex)
        // IncludeIf.Field
        if (this.compliance_action === bytes_1.char('R')) {
            if (this.bitcoin_dispersions.length > (Math.pow(2, 16)) - 1) {
                return sprintf_js_1.sprintf('list field bitcoin_dispersions has too many items %d/%d', this.bitcoin_dispersions.length, (Math.pow(2, 16)) - 1);
            }
            const err = this.bitcoin_dispersions.find((value, i) => {
                const err2 = value.Validate();
                if (err2)
                    return sprintf_js_1.sprintf('list field bitcoin_dispersions[%d] is invalid : %s', i, err2);
            });
            if (err)
                return err;
        }
        // Message (string) 
        {
            if (this.message.length > (Math.pow(2, 32))) {
                return sprintf_js_1.sprintf('varchar field message too long %d/%d', this.message.length, (Math.pow(2, 32)) - 1);
            }
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('compliance_action:%s', this.compliance_action.toString()));
        vals.push(sprintf_js_1.sprintf('asset_type:%s', this.asset_type.toString()));
        vals.push(sprintf_js_1.sprintf('asset_code:%s', this.asset_code.toString()));
        vals.push(sprintf_js_1.sprintf('target_addresses:%s', this.target_addresses.toString()));
        vals.push(sprintf_js_1.sprintf('freeze_tx_id:%s', this.freeze_tx_id.toString()));
        vals.push(sprintf_js_1.sprintf('freeze_period:%s', this.freeze_period.toString()));
        vals.push(sprintf_js_1.sprintf('deposit_address:%s', this.deposit_address.toString()));
        vals.push(sprintf_js_1.sprintf('authority_included:%s', this.authority_included.toString()));
        vals.push(sprintf_js_1.sprintf('authority_name:%s', this.authority_name.toString()));
        vals.push(sprintf_js_1.sprintf('authority_public_key:%s', this.authority_public_key.toString()));
        vals.push(sprintf_js_1.sprintf('signature_algorithm:%d', this.signature_algorithm));
        vals.push(sprintf_js_1.sprintf('order_signature:%s', this.order_signature.toString()));
        vals.push(sprintf_js_1.sprintf('supporting_evidence_hash:%s', this.supporting_evidence_hash.toString()));
        vals.push(sprintf_js_1.sprintf('ref_txs:%s', this.ref_txs.toString()));
        vals.push(sprintf_js_1.sprintf('bitcoin_dispersions:%s', this.bitcoin_dispersions.toString()));
        vals.push(sprintf_js_1.sprintf('message:%s', this.message.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for Order.', e);
            throw e;
        }
        // this.compliance_action = parsed.complianceAction (fixedchar)
        // this.asset_type = parsed.assetType (fixedchar)
        this.asset_code = new protocol_types_1.AssetCode();
        this.asset_code.fromJSON(JSON.stringify(parsed.assetCode));
        // this.target_addresses (TargetAddress[])
        this.freeze_tx_id = new protocol_types_1.TxId();
        this.freeze_tx_id.fromJSON(JSON.stringify(parsed.freezeTxId));
        this.freeze_period = new protocol_types_1.Timestamp();
        this.freeze_period.fromJSON(JSON.stringify(parsed.freezePeriod));
        this.deposit_address = new protocol_types_1.PublicKeyHash();
        this.deposit_address.fromJSON(JSON.stringify(parsed.depositAddress));
        // this.authority_included = parsed.authorityIncluded (bool)
        this.authority_name = parsed.authorityName;
        // this.authority_public_key = parsed.authorityPublicKey (varbin)
        this.signature_algorithm = parsed.signatureAlgorithm;
        // this.order_signature = parsed.orderSignature (varbin)
        // this.supporting_evidence_hash = parsed.supportingEvidenceHash (bin)
        // this.ref_txs = parsed.refTxs (varbin)
        // this.bitcoin_dispersions (QuantityIndex[])
        this.message = parsed.message;
        return this.Validate();
    }
}
exports.Order = Order;
// Freeze The contract responding to an Order action to freeze assets. To
// be used to comply with contractual/legal/issuer requirements. The target
// public address(es) will be marked as frozen. However the Freeze action
// publishes this fact to the public blockchain for transparency. The
// contract will not respond to any actions requested by the frozen
// address.
class Freeze extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeFreeze;
        this.typeStr = 'Freeze';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeFreeze;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // asset_type (string)
        {
            bytes_1.WriteFixedChar(buf, this.asset_type, 3);
        }
        // asset_code (AssetCode)
        {
            buf.write(this.asset_code.Serialize());
        }
        // quantities ([]QuantityIndex)
        {
            bytes_1.WriteVariableSize(buf, this.quantities.length, 16, 8);
            this.quantities.forEach((value) => {
                buf.write(value.Serialize());
            });
        }
        // freeze_period (Timestamp)
        {
            buf.write(this.freeze_period.Serialize());
        }
        // timestamp (Timestamp)
        {
            buf.write(this.timestamp.Serialize());
        }
        return buf.buf;
    }
    // write populates the fields in Freeze from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // asset_type (string)
        {
            this.asset_type = bytes_1.ReadFixedChar(buf, 3);
        }
        // asset_code (AssetCode)
        {
            this.asset_code = new protocol_types_1.AssetCode();
            this.asset_code.Write(buf);
        }
        // quantities ([]QuantityIndex)
        {
            // IsInternalTypeArray
            const size = bytes_1.ReadVariableSize(buf, 16, 8);
            this.quantities = [];
            for (let i = 0; i < size; i++) {
                const newValue = new field_types_1.QuantityIndex();
                newValue.Write(buf);
                this.quantities.push(newValue);
            }
        }
        // freeze_period (Timestamp)
        {
            this.freeze_period = new protocol_types_1.Timestamp();
            this.freeze_period.Write(buf);
        }
        // timestamp (Timestamp)
        {
            this.timestamp = new protocol_types_1.Timestamp();
            this.timestamp.Write(buf);
        }
        return b.length - buf.length;
    }
    Validate() {
        // AssetType (string) 
        {
            if (this.asset_type.length > 3) {
                return sprintf_js_1.sprintf('fixedchar field asset_type too long %d/%d', this.asset_type.length, 3);
            }
        }
        // AssetCode (AssetCode) 
        {
            // IsInternalType
            const err = this.asset_code.Validate();
            if (err)
                return sprintf_js_1.sprintf('field asset_code is invalid : %s', err);
        }
        // Quantities ([]QuantityIndex) 
        {
            if (this.quantities.length > (Math.pow(2, 16)) - 1) {
                return sprintf_js_1.sprintf('list field quantities has too many items %d/%d', this.quantities.length, (Math.pow(2, 16)) - 1);
            }
            const err = this.quantities.find((value, i) => {
                const err2 = value.Validate();
                if (err2)
                    return sprintf_js_1.sprintf('list field quantities[%d] is invalid : %s', i, err2);
            });
            if (err)
                return err;
        }
        // FreezePeriod (Timestamp) 
        {
            // IsInternalType
            const err = this.freeze_period.Validate();
            if (err)
                return sprintf_js_1.sprintf('field freeze_period is invalid : %s', err);
        }
        // Timestamp (Timestamp) 
        {
            // IsInternalType
            const err = this.timestamp.Validate();
            if (err)
                return sprintf_js_1.sprintf('field timestamp is invalid : %s', err);
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('asset_type:%s', this.asset_type.toString()));
        vals.push(sprintf_js_1.sprintf('asset_code:%s', this.asset_code.toString()));
        vals.push(sprintf_js_1.sprintf('quantities:%s', this.quantities.toString()));
        vals.push(sprintf_js_1.sprintf('freeze_period:%s', this.freeze_period.toString()));
        vals.push(sprintf_js_1.sprintf('timestamp:%s', this.timestamp.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for Freeze.', e);
            throw e;
        }
        // this.asset_type = parsed.assetType (fixedchar)
        this.asset_code = new protocol_types_1.AssetCode();
        this.asset_code.fromJSON(JSON.stringify(parsed.assetCode));
        // this.quantities (QuantityIndex[])
        this.freeze_period = new protocol_types_1.Timestamp();
        this.freeze_period.fromJSON(JSON.stringify(parsed.freezePeriod));
        this.timestamp = new protocol_types_1.Timestamp();
        this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
        return this.Validate();
    }
}
exports.Freeze = Freeze;
// Thaw The contract responding to an Order action to thaw assets. To be
// used to comply with contractual obligations or legal requirements. The
// Alleged Offender's tokens will be unfrozen to allow them to resume
// normal exchange and governance activities.
class Thaw extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeThaw;
        this.typeStr = 'Thaw';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeThaw;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // freeze_tx_id (TxId)
        {
            buf.write(this.freeze_tx_id.Serialize());
        }
        // timestamp (Timestamp)
        {
            buf.write(this.timestamp.Serialize());
        }
        return buf.buf;
    }
    // write populates the fields in Thaw from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // freeze_tx_id (TxId)
        {
            this.freeze_tx_id = new protocol_types_1.TxId();
            this.freeze_tx_id.Write(buf);
        }
        // timestamp (Timestamp)
        {
            this.timestamp = new protocol_types_1.Timestamp();
            this.timestamp.Write(buf);
        }
        return b.length - buf.length;
    }
    Validate() {
        // FreezeTxId (TxId) 
        {
            // IsInternalType
            const err = this.freeze_tx_id.Validate();
            if (err)
                return sprintf_js_1.sprintf('field freeze_tx_id is invalid : %s', err);
        }
        // Timestamp (Timestamp) 
        {
            // IsInternalType
            const err = this.timestamp.Validate();
            if (err)
                return sprintf_js_1.sprintf('field timestamp is invalid : %s', err);
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('freeze_tx_id:%s', this.freeze_tx_id.toString()));
        vals.push(sprintf_js_1.sprintf('timestamp:%s', this.timestamp.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for Thaw.', e);
            throw e;
        }
        this.freeze_tx_id = new protocol_types_1.TxId();
        this.freeze_tx_id.fromJSON(JSON.stringify(parsed.freezeTxId));
        this.timestamp = new protocol_types_1.Timestamp();
        this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
        return this.Validate();
    }
}
exports.Thaw = Thaw;
// Confiscation The contract responding to an Order action to confiscate
// assets. To be used to comply with contractual obligations, legal and/or
// issuer requirements.
class Confiscation extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeConfiscation;
        this.typeStr = 'Confiscation';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeConfiscation;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // asset_type (string)
        {
            bytes_1.WriteFixedChar(buf, this.asset_type, 3);
        }
        // asset_code (AssetCode)
        {
            buf.write(this.asset_code.Serialize());
        }
        // quantities ([]QuantityIndex)
        {
            bytes_1.WriteVariableSize(buf, this.quantities.length, 16, 8);
            this.quantities.forEach((value) => {
                buf.write(value.Serialize());
            });
        }
        // deposit_qty (uint64)
        {
            bytes_1.write(buf, this.deposit_qty, 'uint64');
        }
        // timestamp (Timestamp)
        {
            buf.write(this.timestamp.Serialize());
        }
        return buf.buf;
    }
    // write populates the fields in Confiscation from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // asset_type (string)
        {
            this.asset_type = bytes_1.ReadFixedChar(buf, 3);
        }
        // asset_code (AssetCode)
        {
            this.asset_code = new protocol_types_1.AssetCode();
            this.asset_code.Write(buf);
        }
        // quantities ([]QuantityIndex)
        {
            // IsInternalTypeArray
            const size = bytes_1.ReadVariableSize(buf, 16, 8);
            this.quantities = [];
            for (let i = 0; i < size; i++) {
                const newValue = new field_types_1.QuantityIndex();
                newValue.Write(buf);
                this.quantities.push(newValue);
            }
        }
        // deposit_qty (uint64)
        {
            this.deposit_qty = bytes_1.read(buf, 'uint64');
        }
        // timestamp (Timestamp)
        {
            this.timestamp = new protocol_types_1.Timestamp();
            this.timestamp.Write(buf);
        }
        return b.length - buf.length;
    }
    Validate() {
        // AssetType (string) 
        {
            if (this.asset_type.length > 3) {
                return sprintf_js_1.sprintf('fixedchar field asset_type too long %d/%d', this.asset_type.length, 3);
            }
        }
        // AssetCode (AssetCode) 
        {
            // IsInternalType
            const err = this.asset_code.Validate();
            if (err)
                return sprintf_js_1.sprintf('field asset_code is invalid : %s', err);
        }
        // Quantities ([]QuantityIndex) 
        {
            if (this.quantities.length > (Math.pow(2, 16)) - 1) {
                return sprintf_js_1.sprintf('list field quantities has too many items %d/%d', this.quantities.length, (Math.pow(2, 16)) - 1);
            }
            const err = this.quantities.find((value, i) => {
                const err2 = value.Validate();
                if (err2)
                    return sprintf_js_1.sprintf('list field quantities[%d] is invalid : %s', i, err2);
            });
            if (err)
                return err;
        }
        // DepositQty (uint64) 
        {
        }
        // Timestamp (Timestamp) 
        {
            // IsInternalType
            const err = this.timestamp.Validate();
            if (err)
                return sprintf_js_1.sprintf('field timestamp is invalid : %s', err);
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('asset_type:%s', this.asset_type.toString()));
        vals.push(sprintf_js_1.sprintf('asset_code:%s', this.asset_code.toString()));
        vals.push(sprintf_js_1.sprintf('quantities:%s', this.quantities.toString()));
        vals.push(sprintf_js_1.sprintf('deposit_qty:%d', this.deposit_qty));
        vals.push(sprintf_js_1.sprintf('timestamp:%s', this.timestamp.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for Confiscation.', e);
            throw e;
        }
        // this.asset_type = parsed.assetType (fixedchar)
        this.asset_code = new protocol_types_1.AssetCode();
        this.asset_code.fromJSON(JSON.stringify(parsed.assetCode));
        // this.quantities (QuantityIndex[])
        this.deposit_qty = parsed.depositQty;
        this.timestamp = new protocol_types_1.Timestamp();
        this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
        return this.Validate();
    }
}
exports.Confiscation = Confiscation;
// Reconciliation The contract responding to an Order action to reconcile
// assets. To be used at the direction of the administration to fix record
// keeping errors with bitcoin and token balances.
class Reconciliation extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeReconciliation;
        this.typeStr = 'Reconciliation';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeReconciliation;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // asset_type (string)
        {
            bytes_1.WriteFixedChar(buf, this.asset_type, 3);
        }
        // asset_code (AssetCode)
        {
            buf.write(this.asset_code.Serialize());
        }
        // quantities ([]QuantityIndex)
        {
            bytes_1.WriteVariableSize(buf, this.quantities.length, 16, 8);
            this.quantities.forEach((value) => {
                buf.write(value.Serialize());
            });
        }
        // timestamp (Timestamp)
        {
            buf.write(this.timestamp.Serialize());
        }
        return buf.buf;
    }
    // write populates the fields in Reconciliation from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // asset_type (string)
        {
            this.asset_type = bytes_1.ReadFixedChar(buf, 3);
        }
        // asset_code (AssetCode)
        {
            this.asset_code = new protocol_types_1.AssetCode();
            this.asset_code.Write(buf);
        }
        // quantities ([]QuantityIndex)
        {
            // IsInternalTypeArray
            const size = bytes_1.ReadVariableSize(buf, 16, 8);
            this.quantities = [];
            for (let i = 0; i < size; i++) {
                const newValue = new field_types_1.QuantityIndex();
                newValue.Write(buf);
                this.quantities.push(newValue);
            }
        }
        // timestamp (Timestamp)
        {
            this.timestamp = new protocol_types_1.Timestamp();
            this.timestamp.Write(buf);
        }
        return b.length - buf.length;
    }
    Validate() {
        // AssetType (string) 
        {
            if (this.asset_type.length > 3) {
                return sprintf_js_1.sprintf('fixedchar field asset_type too long %d/%d', this.asset_type.length, 3);
            }
        }
        // AssetCode (AssetCode) 
        {
            // IsInternalType
            const err = this.asset_code.Validate();
            if (err)
                return sprintf_js_1.sprintf('field asset_code is invalid : %s', err);
        }
        // Quantities ([]QuantityIndex) 
        {
            if (this.quantities.length > (Math.pow(2, 16)) - 1) {
                return sprintf_js_1.sprintf('list field quantities has too many items %d/%d', this.quantities.length, (Math.pow(2, 16)) - 1);
            }
            const err = this.quantities.find((value, i) => {
                const err2 = value.Validate();
                if (err2)
                    return sprintf_js_1.sprintf('list field quantities[%d] is invalid : %s', i, err2);
            });
            if (err)
                return err;
        }
        // Timestamp (Timestamp) 
        {
            // IsInternalType
            const err = this.timestamp.Validate();
            if (err)
                return sprintf_js_1.sprintf('field timestamp is invalid : %s', err);
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('asset_type:%s', this.asset_type.toString()));
        vals.push(sprintf_js_1.sprintf('asset_code:%s', this.asset_code.toString()));
        vals.push(sprintf_js_1.sprintf('quantities:%s', this.quantities.toString()));
        vals.push(sprintf_js_1.sprintf('timestamp:%s', this.timestamp.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for Reconciliation.', e);
            throw e;
        }
        // this.asset_type = parsed.assetType (fixedchar)
        this.asset_code = new protocol_types_1.AssetCode();
        this.asset_code.fromJSON(JSON.stringify(parsed.assetCode));
        // this.quantities (QuantityIndex[])
        this.timestamp = new protocol_types_1.Timestamp();
        this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
        return this.Validate();
    }
}
exports.Reconciliation = Reconciliation;
// Proposal Allows the Administration/Token Holders to propose a change
// (aka Initiative/Shareholder vote). A significant cost - specified in the
// Contract Formation - can be attached to this action when sent from Token
// Holders to reduce spam, as the resulting vote will be put to all token
// owners.
class Proposal extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeProposal;
        this.typeStr = 'Proposal';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeProposal;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // initiator (uint8)
        {
            bytes_1.write(buf, this.initiator, 'uint8');
        }
        // asset_specific_vote (bool)
        {
            bytes_1.write(buf, this.asset_specific_vote, 'bool');
        }
        // asset_type (string)
        if (this.asset_specific_vote) {
            bytes_1.WriteFixedChar(buf, this.asset_type, 3);
        }
        // asset_code (AssetCode)
        if (this.asset_specific_vote) {
            buf.write(this.asset_code.Serialize());
        }
        // vote_system (uint8)
        {
            bytes_1.write(buf, this.vote_system, 'uint8');
        }
        // specific (bool)
        {
            bytes_1.write(buf, this.specific, 'bool');
        }
        // proposed_amendments ([]Amendment)
        if (this.specific) {
            bytes_1.WriteVariableSize(buf, this.proposed_amendments.length, 8, 8);
            this.proposed_amendments.forEach((value) => {
                buf.write(value.Serialize());
            });
        }
        // vote_options (string)
        {
            bytes_1.WriteVarChar(buf, this.vote_options, 8);
        }
        // vote_max (uint8)
        {
            bytes_1.write(buf, this.vote_max, 'uint8');
        }
        // proposal_description (string)
        {
            bytes_1.WriteVarChar(buf, this.proposal_description, 32);
        }
        // proposal_document_hash ([32]byte)
        {
            bytes_1.write(buf, this.proposal_document_hash, '[32]byte');
        }
        // vote_cut_off_timestamp (Timestamp)
        {
            buf.write(this.vote_cut_off_timestamp.Serialize());
        }
        return buf.buf;
    }
    // write populates the fields in Proposal from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // initiator (uint8)
        {
            this.initiator = bytes_1.read(buf, 'uint8');
        }
        // asset_specific_vote (bool)
        {
            this.asset_specific_vote = bytes_1.read(buf, 'bool');
        }
        // asset_type (string)
        if (this.asset_specific_vote) {
            this.asset_type = bytes_1.ReadFixedChar(buf, 3);
        }
        // asset_code (AssetCode)
        if (this.asset_specific_vote) {
            this.asset_code = new protocol_types_1.AssetCode();
            this.asset_code.Write(buf);
        }
        // vote_system (uint8)
        {
            this.vote_system = bytes_1.read(buf, 'uint8');
        }
        // specific (bool)
        {
            this.specific = bytes_1.read(buf, 'bool');
        }
        // proposed_amendments ([]Amendment)
        if (this.specific) {
            // IsInternalTypeArray
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            this.proposed_amendments = [];
            for (let i = 0; i < size; i++) {
                const newValue = new field_types_1.Amendment();
                newValue.Write(buf);
                this.proposed_amendments.push(newValue);
            }
        }
        // vote_options (string)
        {
            this.vote_options = bytes_1.ReadVarChar(buf, 8);
        }
        // vote_max (uint8)
        {
            this.vote_max = bytes_1.read(buf, 'uint8');
        }
        // proposal_description (string)
        {
            this.proposal_description = bytes_1.ReadVarChar(buf, 32);
        }
        // proposal_document_hash ([32]byte)
        {
            this.proposal_document_hash = bytes_1.read(buf, '[32]byte');
        }
        // vote_cut_off_timestamp (Timestamp)
        {
            this.vote_cut_off_timestamp = new protocol_types_1.Timestamp();
            this.vote_cut_off_timestamp.Write(buf);
        }
        return b.length - buf.length;
    }
    Validate() {
        // Initiator (uint8) 
        {
            // $field.IntValues [0 1]
            if (this.initiator !== 0 && this.initiator !== 1) {
                return sprintf_js_1.sprintf('field initiator value is invalid : %d', this.initiator);
            }
        }
        // AssetSpecificVote (bool) 
        {
        }
        // AssetType (string)
        // IncludeIfTrue
        if (this.asset_specific_vote) {
            if (this.asset_type.length > 3) {
                return sprintf_js_1.sprintf('fixedchar field asset_type too long %d/%d', this.asset_type.length, 3);
            }
        }
        // AssetCode (AssetCode)
        // IncludeIfTrue
        if (this.asset_specific_vote) {
            // IsInternalType
            const err = this.asset_code.Validate();
            if (err)
                return sprintf_js_1.sprintf('field asset_code is invalid : %s', err);
        }
        // VoteSystem (uint8) 
        {
        }
        // Specific (bool) 
        {
        }
        // ProposedAmendments ([]Amendment)
        // IncludeIfTrue
        if (this.specific) {
            if (this.proposed_amendments.length > (Math.pow(2, 8)) - 1) {
                return sprintf_js_1.sprintf('list field proposed_amendments has too many items %d/%d', this.proposed_amendments.length, (Math.pow(2, 8)) - 1);
            }
            const err = this.proposed_amendments.find((value, i) => {
                const err2 = value.Validate();
                if (err2)
                    return sprintf_js_1.sprintf('list field proposed_amendments[%d] is invalid : %s', i, err2);
            });
            if (err)
                return err;
        }
        // VoteOptions (string) 
        {
            if (this.vote_options.length > (Math.pow(2, 8))) {
                return sprintf_js_1.sprintf('varchar field vote_options too long %d/%d', this.vote_options.length, (Math.pow(2, 8)) - 1);
            }
        }
        // VoteMax (uint8) 
        {
        }
        // ProposalDescription (string) 
        {
            if (this.proposal_description.length > (Math.pow(2, 32))) {
                return sprintf_js_1.sprintf('varchar field proposal_description too long %d/%d', this.proposal_description.length, (Math.pow(2, 32)) - 1);
            }
        }
        // ProposalDocumentHash ([32]byte) 
        {
        }
        // VoteCutOffTimestamp (Timestamp) 
        {
            // IsInternalType
            const err = this.vote_cut_off_timestamp.Validate();
            if (err)
                return sprintf_js_1.sprintf('field vote_cut_off_timestamp is invalid : %s', err);
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('initiator:%d', this.initiator));
        vals.push(sprintf_js_1.sprintf('asset_specific_vote:%s', this.asset_specific_vote.toString()));
        vals.push(sprintf_js_1.sprintf('asset_type:%s', this.asset_type.toString()));
        vals.push(sprintf_js_1.sprintf('asset_code:%s', this.asset_code.toString()));
        vals.push(sprintf_js_1.sprintf('vote_system:%d', this.vote_system));
        vals.push(sprintf_js_1.sprintf('specific:%s', this.specific.toString()));
        vals.push(sprintf_js_1.sprintf('proposed_amendments:%s', this.proposed_amendments.toString()));
        vals.push(sprintf_js_1.sprintf('vote_options:%s', this.vote_options.toString()));
        vals.push(sprintf_js_1.sprintf('vote_max:%d', this.vote_max));
        vals.push(sprintf_js_1.sprintf('proposal_description:%s', this.proposal_description.toString()));
        vals.push(sprintf_js_1.sprintf('proposal_document_hash:%s', this.proposal_document_hash.toString()));
        vals.push(sprintf_js_1.sprintf('vote_cut_off_timestamp:%s', this.vote_cut_off_timestamp.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for Proposal.', e);
            throw e;
        }
        this.initiator = parsed.initiator;
        // this.asset_specific_vote = parsed.assetSpecificVote (bool)
        // this.asset_type = parsed.assetType (fixedchar)
        this.asset_code = new protocol_types_1.AssetCode();
        this.asset_code.fromJSON(JSON.stringify(parsed.assetCode));
        this.vote_system = parsed.voteSystem;
        // this.specific = parsed.specific (bool)
        // this.proposed_amendments (Amendment[])
        this.vote_options = parsed.voteOptions;
        this.vote_max = parsed.voteMax;
        this.proposal_description = parsed.proposalDescription;
        // this.proposal_document_hash = parsed.proposalDocumentHash (bin)
        this.vote_cut_off_timestamp = new protocol_types_1.Timestamp();
        this.vote_cut_off_timestamp.fromJSON(JSON.stringify(parsed.voteCutOffTimestamp));
        return this.Validate();
    }
}
exports.Proposal = Proposal;
// Vote A vote is created by the Contract in response to a valid Proposal
// Action.
class Vote extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeVote;
        this.typeStr = 'Vote';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeVote;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // timestamp (Timestamp)
        {
            buf.write(this.timestamp.Serialize());
        }
        return buf.buf;
    }
    // write populates the fields in Vote from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // timestamp (Timestamp)
        {
            this.timestamp = new protocol_types_1.Timestamp();
            this.timestamp.Write(buf);
        }
        return b.length - buf.length;
    }
    Validate() {
        // Timestamp (Timestamp) 
        {
            // IsInternalType
            const err = this.timestamp.Validate();
            if (err)
                return sprintf_js_1.sprintf('field timestamp is invalid : %s', err);
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('timestamp:%s', this.timestamp.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for Vote.', e);
            throw e;
        }
        this.timestamp = new protocol_types_1.Timestamp();
        this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
        return this.Validate();
    }
}
exports.Vote = Vote;
// BallotCast Used by Token Owners to cast their ballot (vote) on
// proposals. 1 Vote per token unless a vote multiplier is specified in the
// relevant Asset Definition action.
class BallotCast extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeBallotCast;
        this.typeStr = 'BallotCast';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeBallotCast;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // vote_tx_id (TxId)
        {
            buf.write(this.vote_tx_id.Serialize());
        }
        // vote (string)
        {
            bytes_1.WriteVarChar(buf, this.vote, 8);
        }
        return buf.buf;
    }
    // write populates the fields in BallotCast from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // vote_tx_id (TxId)
        {
            this.vote_tx_id = new protocol_types_1.TxId();
            this.vote_tx_id.Write(buf);
        }
        // vote (string)
        {
            this.vote = bytes_1.ReadVarChar(buf, 8);
        }
        return b.length - buf.length;
    }
    Validate() {
        // VoteTxId (TxId) 
        {
            // IsInternalType
            const err = this.vote_tx_id.Validate();
            if (err)
                return sprintf_js_1.sprintf('field vote_tx_id is invalid : %s', err);
        }
        // Vote (string) 
        {
            if (this.vote.length > (Math.pow(2, 8))) {
                return sprintf_js_1.sprintf('varchar field vote too long %d/%d', this.vote.length, (Math.pow(2, 8)) - 1);
            }
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('vote_tx_id:%s', this.vote_tx_id.toString()));
        vals.push(sprintf_js_1.sprintf('vote:%s', this.vote.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for BallotCast.', e);
            throw e;
        }
        this.vote_tx_id = new protocol_types_1.TxId();
        this.vote_tx_id.fromJSON(JSON.stringify(parsed.voteTxId));
        this.vote = parsed.vote;
        return this.Validate();
    }
}
exports.BallotCast = BallotCast;
// BallotCounted The smart contract will respond to a Ballot Cast action
// with a Ballot Counted action if the Ballot Cast is valid. If the Ballot
// Cast is not valid, then the smart contract will respond with a Rejection
// Action.
class BallotCounted extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeBallotCounted;
        this.typeStr = 'BallotCounted';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeBallotCounted;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // vote_tx_id (TxId)
        {
            buf.write(this.vote_tx_id.Serialize());
        }
        // vote (string)
        {
            bytes_1.WriteVarChar(buf, this.vote, 8);
        }
        // quantity (uint64)
        {
            bytes_1.write(buf, this.quantity, 'uint64');
        }
        // timestamp (Timestamp)
        {
            buf.write(this.timestamp.Serialize());
        }
        return buf.buf;
    }
    // write populates the fields in BallotCounted from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // vote_tx_id (TxId)
        {
            this.vote_tx_id = new protocol_types_1.TxId();
            this.vote_tx_id.Write(buf);
        }
        // vote (string)
        {
            this.vote = bytes_1.ReadVarChar(buf, 8);
        }
        // quantity (uint64)
        {
            this.quantity = bytes_1.read(buf, 'uint64');
        }
        // timestamp (Timestamp)
        {
            this.timestamp = new protocol_types_1.Timestamp();
            this.timestamp.Write(buf);
        }
        return b.length - buf.length;
    }
    Validate() {
        // VoteTxId (TxId) 
        {
            // IsInternalType
            const err = this.vote_tx_id.Validate();
            if (err)
                return sprintf_js_1.sprintf('field vote_tx_id is invalid : %s', err);
        }
        // Vote (string) 
        {
            if (this.vote.length > (Math.pow(2, 8))) {
                return sprintf_js_1.sprintf('varchar field vote too long %d/%d', this.vote.length, (Math.pow(2, 8)) - 1);
            }
        }
        // Quantity (uint64) 
        {
        }
        // Timestamp (Timestamp) 
        {
            // IsInternalType
            const err = this.timestamp.Validate();
            if (err)
                return sprintf_js_1.sprintf('field timestamp is invalid : %s', err);
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('vote_tx_id:%s', this.vote_tx_id.toString()));
        vals.push(sprintf_js_1.sprintf('vote:%s', this.vote.toString()));
        vals.push(sprintf_js_1.sprintf('quantity:%d', this.quantity));
        vals.push(sprintf_js_1.sprintf('timestamp:%s', this.timestamp.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for BallotCounted.', e);
            throw e;
        }
        this.vote_tx_id = new protocol_types_1.TxId();
        this.vote_tx_id.fromJSON(JSON.stringify(parsed.voteTxId));
        this.vote = parsed.vote;
        this.quantity = parsed.quantity;
        this.timestamp = new protocol_types_1.Timestamp();
        this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
        return this.Validate();
    }
}
exports.BallotCounted = BallotCounted;
// Result Once a vote has been completed the results are published. After
// the result is posted, it is up to the administration to send a
// contract/asset amendement if appropriate.
class Result extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeResult;
        this.typeStr = 'Result';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeResult;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // asset_specific_vote (bool)
        {
            bytes_1.write(buf, this.asset_specific_vote, 'bool');
        }
        // asset_type (string)
        if (this.asset_specific_vote) {
            bytes_1.WriteFixedChar(buf, this.asset_type, 3);
        }
        // asset_code (AssetCode)
        if (this.asset_specific_vote) {
            buf.write(this.asset_code.Serialize());
        }
        // specific (bool)
        {
            bytes_1.write(buf, this.specific, 'bool');
        }
        // proposed_amendments ([]Amendment)
        if (this.specific) {
            bytes_1.WriteVariableSize(buf, this.proposed_amendments.length, 8, 8);
            this.proposed_amendments.forEach((value) => {
                buf.write(value.Serialize());
            });
        }
        // vote_tx_id (TxId)
        {
            buf.write(this.vote_tx_id.Serialize());
        }
        // option_tally ([]uint64)
        {
            // IsNativeTypeArray []uint64
            const type = '[]uint64'.slice(2);
            bytes_1.WriteVariableSize(buf, this.option_tally.length, 8, 8);
            this.option_tally.forEach(value => {
                bytes_1.write(buf, value, type); // TODO
            });
        }
        // result (string)
        {
            bytes_1.WriteVarChar(buf, this.result, 8);
        }
        // timestamp (Timestamp)
        {
            buf.write(this.timestamp.Serialize());
        }
        return buf.buf;
    }
    // write populates the fields in Result from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // asset_specific_vote (bool)
        {
            this.asset_specific_vote = bytes_1.read(buf, 'bool');
        }
        // asset_type (string)
        if (this.asset_specific_vote) {
            this.asset_type = bytes_1.ReadFixedChar(buf, 3);
        }
        // asset_code (AssetCode)
        if (this.asset_specific_vote) {
            this.asset_code = new protocol_types_1.AssetCode();
            this.asset_code.Write(buf);
        }
        // specific (bool)
        {
            this.specific = bytes_1.read(buf, 'bool');
        }
        // proposed_amendments ([]Amendment)
        if (this.specific) {
            // IsInternalTypeArray
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            this.proposed_amendments = [];
            for (let i = 0; i < size; i++) {
                const newValue = new field_types_1.Amendment();
                newValue.Write(buf);
                this.proposed_amendments.push(newValue);
            }
        }
        // vote_tx_id (TxId)
        {
            this.vote_tx_id = new protocol_types_1.TxId();
            this.vote_tx_id.Write(buf);
        }
        // option_tally ([]uint64)
        {
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            const type = '[]uint64'.slice(2);
            this.option_tally = [...Array(size)].map(() => bytes_1.read(buf, type));
        }
        // result (string)
        {
            this.result = bytes_1.ReadVarChar(buf, 8);
        }
        // timestamp (Timestamp)
        {
            this.timestamp = new protocol_types_1.Timestamp();
            this.timestamp.Write(buf);
        }
        return b.length - buf.length;
    }
    Validate() {
        // AssetSpecificVote (bool) 
        {
        }
        // AssetType (string)
        // IncludeIfTrue
        if (this.asset_specific_vote) {
            if (this.asset_type.length > 3) {
                return sprintf_js_1.sprintf('fixedchar field asset_type too long %d/%d', this.asset_type.length, 3);
            }
        }
        // AssetCode (AssetCode)
        // IncludeIfTrue
        if (this.asset_specific_vote) {
            // IsInternalType
            const err = this.asset_code.Validate();
            if (err)
                return sprintf_js_1.sprintf('field asset_code is invalid : %s', err);
        }
        // Specific (bool) 
        {
        }
        // ProposedAmendments ([]Amendment)
        // IncludeIfTrue
        if (this.specific) {
            if (this.proposed_amendments.length > (Math.pow(2, 8)) - 1) {
                return sprintf_js_1.sprintf('list field proposed_amendments has too many items %d/%d', this.proposed_amendments.length, (Math.pow(2, 8)) - 1);
            }
            const err = this.proposed_amendments.find((value, i) => {
                const err2 = value.Validate();
                if (err2)
                    return sprintf_js_1.sprintf('list field proposed_amendments[%d] is invalid : %s', i, err2);
            });
            if (err)
                return err;
        }
        // VoteTxId (TxId) 
        {
            // IsInternalType
            const err = this.vote_tx_id.Validate();
            if (err)
                return sprintf_js_1.sprintf('field vote_tx_id is invalid : %s', err);
        }
        // OptionTally ([]uint64) 
        {
            if (this.option_tally.length > (Math.pow(2, 8)) - 1) {
                return sprintf_js_1.sprintf('list field option_tally has too many items %d/%d', this.option_tally.length, (Math.pow(2, 8)) - 1);
            }
        }
        // Result (string) 
        {
            if (this.result.length > (Math.pow(2, 8))) {
                return sprintf_js_1.sprintf('varchar field result too long %d/%d', this.result.length, (Math.pow(2, 8)) - 1);
            }
        }
        // Timestamp (Timestamp) 
        {
            // IsInternalType
            const err = this.timestamp.Validate();
            if (err)
                return sprintf_js_1.sprintf('field timestamp is invalid : %s', err);
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('asset_specific_vote:%s', this.asset_specific_vote.toString()));
        vals.push(sprintf_js_1.sprintf('asset_type:%s', this.asset_type.toString()));
        vals.push(sprintf_js_1.sprintf('asset_code:%s', this.asset_code.toString()));
        vals.push(sprintf_js_1.sprintf('specific:%s', this.specific.toString()));
        vals.push(sprintf_js_1.sprintf('proposed_amendments:%s', this.proposed_amendments.toString()));
        vals.push(sprintf_js_1.sprintf('vote_tx_id:%s', this.vote_tx_id.toString()));
        vals.push(sprintf_js_1.sprintf('option_tally:%d', this.option_tally));
        vals.push(sprintf_js_1.sprintf('result:%s', this.result.toString()));
        vals.push(sprintf_js_1.sprintf('timestamp:%s', this.timestamp.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for Result.', e);
            throw e;
        }
        // this.asset_specific_vote = parsed.assetSpecificVote (bool)
        // this.asset_type = parsed.assetType (fixedchar)
        this.asset_code = new protocol_types_1.AssetCode();
        this.asset_code.fromJSON(JSON.stringify(parsed.assetCode));
        // this.specific = parsed.specific (bool)
        // this.proposed_amendments (Amendment[])
        this.vote_tx_id = new protocol_types_1.TxId();
        this.vote_tx_id.fromJSON(JSON.stringify(parsed.voteTxId));
        this.option_tally = parsed.optionTally;
        this.result = parsed.result;
        this.timestamp = new protocol_types_1.Timestamp();
        this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
        return this.Validate();
    }
}
exports.Result = Result;
// Message The message action is a general purpose communication action.
// 'Twitter/SMS' for Issuers/Investors/Users. The message txn can also be
// used for passing partially signed txns on-chain, establishing private
// communication channels and EDI (receipting, invoices, PO, and private
// offers/bids). The messages are broken down by type for easy filtering in
// the a user's wallet. The Message Types are listed in the Message Types
// table.
class Message extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeMessage;
        this.typeStr = 'Message';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeMessage;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // address_indexes ([]uint16)
        {
            // IsNativeTypeArray []uint16
            const type = '[]uint16'.slice(2);
            bytes_1.WriteVariableSize(buf, this.address_indexes.length, 8, 8);
            this.address_indexes.forEach(value => {
                bytes_1.write(buf, value, type); // TODO
            });
        }
        // message_type (uint16)
        {
            bytes_1.write(buf, this.message_type, 'uint16');
        }
        // message_payload ([]byte)
        {
            bytes_1.WriteVarBin(buf, this.message_payload, 32);
        }
        return buf.buf;
    }
    // write populates the fields in Message from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // address_indexes ([]uint16)
        {
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            const type = '[]uint16'.slice(2);
            this.address_indexes = [...Array(size)].map(() => bytes_1.read(buf, type));
        }
        // message_type (uint16)
        {
            this.message_type = bytes_1.read(buf, 'uint16');
        }
        // message_payload ([]byte)
        {
            this.message_payload = bytes_1.ReadVarBin(buf, 32);
        }
        return b.length - buf.length;
    }
    Validate() {
        // AddressIndexes ([]uint16) 
        {
            if (this.address_indexes.length > (Math.pow(2, 8)) - 1) {
                return sprintf_js_1.sprintf('list field address_indexes has too many items %d/%d', this.address_indexes.length, (Math.pow(2, 8)) - 1);
            }
        }
        // MessageType (uint16) 
        {
        }
        // MessagePayload ([]byte) 
        {
            if (this.message_payload.length >= (Math.pow(2, 32))) {
                return sprintf_js_1.sprintf('varbin field message_payload too long %d/%d', this.message_payload.length, (Math.pow(2, 32)) - 1);
            }
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('address_indexes:%d', this.address_indexes));
        vals.push(sprintf_js_1.sprintf('message_type:%d', this.message_type));
        vals.push(sprintf_js_1.sprintf('message_payload:%s', this.message_payload.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for Message.', e);
            throw e;
        }
        this.address_indexes = parsed.addressIndexes;
        this.message_type = parsed.messageType;
        // this.message_payload = parsed.messagePayload (varbin)
        return this.Validate();
    }
}
exports.Message = Message;
// Rejection Used to reject request actions that do not comply with the
// Contract. If money is to be returned to a User then it is used in lieu
// of the Settlement Action to properly account for token balances. All
// Administration/User request Actions must be responded to by the Contract
// with an Action. The only exception to this rule is when there is not
// enough fees in the first Action for the Contract response action to
// remain revenue neutral. If not enough fees are attached to pay for the
// Contract response then the Contract will not respond.
class Rejection extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeRejection;
        this.typeStr = 'Rejection';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeRejection;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // address_indexes ([]uint16)
        {
            // IsNativeTypeArray []uint16
            const type = '[]uint16'.slice(2);
            bytes_1.WriteVariableSize(buf, this.address_indexes.length, 8, 8);
            this.address_indexes.forEach(value => {
                bytes_1.write(buf, value, type); // TODO
            });
        }
        // reject_address_index (uint16)
        {
            bytes_1.write(buf, this.reject_address_index, 'uint16');
        }
        // rejection_code (uint8)
        {
            bytes_1.write(buf, this.rejection_code, 'uint8');
        }
        // message (string)
        {
            bytes_1.WriteVarChar(buf, this.message, 16);
        }
        // timestamp (Timestamp)
        {
            buf.write(this.timestamp.Serialize());
        }
        return buf.buf;
    }
    // write populates the fields in Rejection from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // address_indexes ([]uint16)
        {
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            const type = '[]uint16'.slice(2);
            this.address_indexes = [...Array(size)].map(() => bytes_1.read(buf, type));
        }
        // reject_address_index (uint16)
        {
            this.reject_address_index = bytes_1.read(buf, 'uint16');
        }
        // rejection_code (uint8)
        {
            this.rejection_code = bytes_1.read(buf, 'uint8');
        }
        // message (string)
        {
            this.message = bytes_1.ReadVarChar(buf, 16);
        }
        // timestamp (Timestamp)
        {
            this.timestamp = new protocol_types_1.Timestamp();
            this.timestamp.Write(buf);
        }
        return b.length - buf.length;
    }
    Validate() {
        // AddressIndexes ([]uint16) 
        {
            if (this.address_indexes.length > (Math.pow(2, 8)) - 1) {
                return sprintf_js_1.sprintf('list field address_indexes has too many items %d/%d', this.address_indexes.length, (Math.pow(2, 8)) - 1);
            }
        }
        // RejectAddressIndex (uint16) 
        {
        }
        // RejectionCode (uint8) 
        {
            if (resources_1.Resources.GetRejectionCode(this.rejection_code) === null) {
                return sprintf_js_1.sprintf('Invalid rejection code value : %d', this.rejection_code);
            }
        }
        // Message (string) 
        {
            if (this.message.length > (Math.pow(2, 16))) {
                return sprintf_js_1.sprintf('varchar field message too long %d/%d', this.message.length, (Math.pow(2, 16)) - 1);
            }
        }
        // Timestamp (Timestamp) 
        {
            // IsInternalType
            const err = this.timestamp.Validate();
            if (err)
                return sprintf_js_1.sprintf('field timestamp is invalid : %s', err);
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('address_indexes:%d', this.address_indexes));
        vals.push(sprintf_js_1.sprintf('reject_address_index:%d', this.reject_address_index));
        vals.push(sprintf_js_1.sprintf('rejection_code:%s', this.rejection_code.toString()));
        vals.push(sprintf_js_1.sprintf('message:%s', this.message.toString()));
        vals.push(sprintf_js_1.sprintf('timestamp:%s', this.timestamp.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for Rejection.', e);
            throw e;
        }
        this.address_indexes = parsed.addressIndexes;
        this.reject_address_index = parsed.rejectAddressIndex;
        // this.rejection_code = parsed.rejectionCode (RejectionCode)
        this.message = parsed.message;
        this.timestamp = new protocol_types_1.Timestamp();
        this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
        return this.Validate();
    }
}
exports.Rejection = Rejection;
// Establishment Establishes an on-chain register.
class Establishment extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeEstablishment;
        this.typeStr = 'Establishment';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeEstablishment;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // message (string)
        {
            bytes_1.WriteVarChar(buf, this.message, 32);
        }
        return buf.buf;
    }
    // write populates the fields in Establishment from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // message (string)
        {
            this.message = bytes_1.ReadVarChar(buf, 32);
        }
        return b.length - buf.length;
    }
    Validate() {
        // Message (string) 
        {
            if (this.message.length > (Math.pow(2, 32))) {
                return sprintf_js_1.sprintf('varchar field message too long %d/%d', this.message.length, (Math.pow(2, 32)) - 1);
            }
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('message:%s', this.message.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for Establishment.', e);
            throw e;
        }
        this.message = parsed.message;
        return this.Validate();
    }
}
exports.Establishment = Establishment;
// Addition Adds an entry to the Register.
class Addition extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeAddition;
        this.typeStr = 'Addition';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeAddition;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // message (string)
        {
            bytes_1.WriteVarChar(buf, this.message, 32);
        }
        return buf.buf;
    }
    // write populates the fields in Addition from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // message (string)
        {
            this.message = bytes_1.ReadVarChar(buf, 32);
        }
        return b.length - buf.length;
    }
    Validate() {
        // Message (string) 
        {
            if (this.message.length > (Math.pow(2, 32))) {
                return sprintf_js_1.sprintf('varchar field message too long %d/%d', this.message.length, (Math.pow(2, 32)) - 1);
            }
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('message:%s', this.message.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for Addition.', e);
            throw e;
        }
        this.message = parsed.message;
        return this.Validate();
    }
}
exports.Addition = Addition;
// Alteration A register entry/record can be altered.
class Alteration extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeAlteration;
        this.typeStr = 'Alteration';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeAlteration;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // entry_tx_id (TxId)
        {
            buf.write(this.entry_tx_id.Serialize());
        }
        // message (string)
        {
            bytes_1.WriteVarChar(buf, this.message, 32);
        }
        return buf.buf;
    }
    // write populates the fields in Alteration from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // entry_tx_id (TxId)
        {
            this.entry_tx_id = new protocol_types_1.TxId();
            this.entry_tx_id.Write(buf);
        }
        // message (string)
        {
            this.message = bytes_1.ReadVarChar(buf, 32);
        }
        return b.length - buf.length;
    }
    Validate() {
        // EntryTxID (TxId) 
        {
            // IsInternalType
            const err = this.entry_tx_id.Validate();
            if (err)
                return sprintf_js_1.sprintf('field entry_tx_id is invalid : %s', err);
        }
        // Message (string) 
        {
            if (this.message.length > (Math.pow(2, 32))) {
                return sprintf_js_1.sprintf('varchar field message too long %d/%d', this.message.length, (Math.pow(2, 32)) - 1);
            }
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('entry_tx_id:%s', this.entry_tx_id.toString()));
        vals.push(sprintf_js_1.sprintf('message:%s', this.message.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for Alteration.', e);
            throw e;
        }
        this.entry_tx_id = new protocol_types_1.TxId();
        this.entry_tx_id.fromJSON(JSON.stringify(parsed.entryTxID));
        this.message = parsed.message;
        return this.Validate();
    }
}
exports.Alteration = Alteration;
// Removal Removes an entry/record from the Register.
class Removal extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeRemoval;
        this.typeStr = 'Removal';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeRemoval;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // entry_tx_id (TxId)
        {
            buf.write(this.entry_tx_id.Serialize());
        }
        // message (string)
        {
            bytes_1.WriteVarChar(buf, this.message, 32);
        }
        return buf.buf;
    }
    // write populates the fields in Removal from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // entry_tx_id (TxId)
        {
            this.entry_tx_id = new protocol_types_1.TxId();
            this.entry_tx_id.Write(buf);
        }
        // message (string)
        {
            this.message = bytes_1.ReadVarChar(buf, 32);
        }
        return b.length - buf.length;
    }
    Validate() {
        // EntryTxID (TxId) 
        {
            // IsInternalType
            const err = this.entry_tx_id.Validate();
            if (err)
                return sprintf_js_1.sprintf('field entry_tx_id is invalid : %s', err);
        }
        // Message (string) 
        {
            if (this.message.length > (Math.pow(2, 32))) {
                return sprintf_js_1.sprintf('varchar field message too long %d/%d', this.message.length, (Math.pow(2, 32)) - 1);
            }
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('entry_tx_id:%s', this.entry_tx_id.toString()));
        vals.push(sprintf_js_1.sprintf('message:%s', this.message.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for Removal.', e);
            throw e;
        }
        this.entry_tx_id = new protocol_types_1.TxId();
        this.entry_tx_id.fromJSON(JSON.stringify(parsed.entryTxID));
        this.message = parsed.message;
        return this.Validate();
    }
}
exports.Removal = Removal;
// Transfer A Token Owner(s) Sends, Exchanges or Swaps a token(s) or
// Bitcoin for a token(s) or Bitcoin. Can be as simple as sending a single
// token to a receiver. Or can be as complex as many senders sending many
// different assets - controlled by many different smart contracts - to a
// number of receivers. This action also supports atomic swaps (tokens for
// tokens). Since many parties and contracts can be involved in a transfer
// and the corresponding settlement action, the partially signed T1 and T2
// actions will need to be passed around on-chain with an M1 action, or
// off-chain.
class Transfer extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeTransfer;
        this.typeStr = 'Transfer';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeTransfer;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // assets ([]AssetTransfer)
        {
            bytes_1.WriteVariableSize(buf, this.assets.length, 8, 8);
            this.assets.forEach((value) => {
                buf.write(value.Serialize());
            });
        }
        // offer_expiry (Timestamp)
        {
            buf.write(this.offer_expiry.Serialize());
        }
        // exchange_fee (uint64)
        {
            bytes_1.write(buf, this.exchange_fee, 'uint64');
        }
        // exchange_fee_address (PublicKeyHash)
        {
            buf.write(this.exchange_fee_address.Serialize());
        }
        return buf.buf;
    }
    // write populates the fields in Transfer from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // assets ([]AssetTransfer)
        {
            // IsInternalTypeArray
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            this.assets = [];
            for (let i = 0; i < size; i++) {
                const newValue = new field_types_1.AssetTransfer();
                newValue.Write(buf);
                this.assets.push(newValue);
            }
        }
        // offer_expiry (Timestamp)
        {
            this.offer_expiry = new protocol_types_1.Timestamp();
            this.offer_expiry.Write(buf);
        }
        // exchange_fee (uint64)
        {
            this.exchange_fee = bytes_1.read(buf, 'uint64');
        }
        // exchange_fee_address (PublicKeyHash)
        {
            this.exchange_fee_address = new protocol_types_1.PublicKeyHash();
            this.exchange_fee_address.Write(buf);
        }
        return b.length - buf.length;
    }
    Validate() {
        // Assets ([]AssetTransfer) 
        {
            if (this.assets.length > (Math.pow(2, 8)) - 1) {
                return sprintf_js_1.sprintf('list field assets has too many items %d/%d', this.assets.length, (Math.pow(2, 8)) - 1);
            }
            const err = this.assets.find((value, i) => {
                const err2 = value.Validate();
                if (err2)
                    return sprintf_js_1.sprintf('list field assets[%d] is invalid : %s', i, err2);
            });
            if (err)
                return err;
        }
        // OfferExpiry (Timestamp) 
        {
            // IsInternalType
            const err = this.offer_expiry.Validate();
            if (err)
                return sprintf_js_1.sprintf('field offer_expiry is invalid : %s', err);
        }
        // ExchangeFee (uint64) 
        {
        }
        // ExchangeFeeAddress (PublicKeyHash) 
        {
            // IsInternalType
            const err = this.exchange_fee_address.Validate();
            if (err)
                return sprintf_js_1.sprintf('field exchange_fee_address is invalid : %s', err);
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('assets:%s', this.assets.toString()));
        vals.push(sprintf_js_1.sprintf('offer_expiry:%s', this.offer_expiry.toString()));
        vals.push(sprintf_js_1.sprintf('exchange_fee:%d', this.exchange_fee));
        vals.push(sprintf_js_1.sprintf('exchange_fee_address:%s', this.exchange_fee_address.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for Transfer.', e);
            throw e;
        }
        // this.assets (AssetTransfer[])
        this.offer_expiry = new protocol_types_1.Timestamp();
        this.offer_expiry.fromJSON(JSON.stringify(parsed.offerExpiry));
        this.exchange_fee = parsed.exchangeFee;
        this.exchange_fee_address = new protocol_types_1.PublicKeyHash();
        this.exchange_fee_address.fromJSON(JSON.stringify(parsed.exchangeFeeAddress));
        return this.Validate();
    }
}
exports.Transfer = Transfer;
// Settlement Settles the transfer request of bitcoins and tokens from
// transfer (T1) actions.
class Settlement extends protocol_1.OpReturnMessage {
    constructor() {
        super(...arguments);
        this.type = ActionCode.CodeSettlement;
        this.typeStr = 'Settlement';
    }
    // Type returns the type identifer for this message.
    Type() {
        return ActionCode.CodeSettlement;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // assets ([]AssetSettlement)
        {
            bytes_1.WriteVariableSize(buf, this.assets.length, 8, 8);
            this.assets.forEach((value) => {
                buf.write(value.Serialize());
            });
        }
        // timestamp (Timestamp)
        {
            buf.write(this.timestamp.Serialize());
        }
        return buf.buf;
    }
    // write populates the fields in Settlement from the byte slice
    write(b) {
        const buf = new util_1.default.Reader(b);
        // assets ([]AssetSettlement)
        {
            // IsInternalTypeArray
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            this.assets = [];
            for (let i = 0; i < size; i++) {
                const newValue = new field_types_1.AssetSettlement();
                newValue.Write(buf);
                this.assets.push(newValue);
            }
        }
        // timestamp (Timestamp)
        {
            this.timestamp = new protocol_types_1.Timestamp();
            this.timestamp.Write(buf);
        }
        return b.length - buf.length;
    }
    Validate() {
        // Assets ([]AssetSettlement) 
        {
            if (this.assets.length > (Math.pow(2, 8)) - 1) {
                return sprintf_js_1.sprintf('list field assets has too many items %d/%d', this.assets.length, (Math.pow(2, 8)) - 1);
            }
            const err = this.assets.find((value, i) => {
                const err2 = value.Validate();
                if (err2)
                    return sprintf_js_1.sprintf('list field assets[%d] is invalid : %s', i, err2);
            });
            if (err)
                return err;
        }
        // Timestamp (Timestamp) 
        {
            // IsInternalType
            const err = this.timestamp.Validate();
            if (err)
                return sprintf_js_1.sprintf('field timestamp is invalid : %s', err);
        }
        return null;
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('assets:%s', this.assets.toString()));
        vals.push(sprintf_js_1.sprintf('timestamp:%s', this.timestamp.toString()));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
    parse(json) {
        let parsed;
        try {
            parsed = JSON.parse(json);
        }
        catch (e) {
            console.error('Failed to parse JSON for Settlement.', e);
            throw e;
        }
        // this.assets (AssetSettlement[])
        this.timestamp = new protocol_types_1.Timestamp();
        this.timestamp.fromJSON(JSON.stringify(parsed.timestamp));
        return this.Validate();
    }
}
exports.Settlement = Settlement;
