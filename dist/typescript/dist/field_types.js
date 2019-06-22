"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const sprintf_js_1 = require("sprintf-js");
const util_1 = __importDefault(require("@keyring/util"));
const bytes_1 = require("./bytes");
const resources_1 = require("./resources");
const protocol_types_1 = require("./protocol_types");
// Administrator Administrator is used to refer to a Administration role in
// an Entity.
class Administrator {
    // Serialize returns the byte representation of the message.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Type (uint8)
        {
            bytes_1.write(buf, this.type, 'uint8');
        }
        // Name (string)
        {
            bytes_1.WriteVarChar(buf, this.name, 8);
        }
        return buf.buf;
    }
    Write(buf) {
        // Type (uint8)
        {
            this.type = bytes_1.read(buf, 'uint8');
        }
        // Name (string)
        {
            this.name = bytes_1.ReadVarChar(buf, 8);
        }
    }
    Validate() {
        // type (uint8)
        {
            if (!resources_1.Resources.GetRoleType(this.type)) {
                return sprintf_js_1.sprintf('Invalid role value : %d', this.type);
            }
        }
        // name (string)
        {
            if (this.name.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field name too long %d/%d', this.name.length, (2 << 8) - 1);
            }
        }
        return;
    }
    Equal(other) {
        // Type (uint8)
        if (this.type !== other.type) {
            return false;
        }
        // Name (string)
        if (this.name !== other.name) {
            return false;
        }
        return true;
    }
}
exports.Administrator = Administrator;
// AgeRestriction Age restriction is used to specify required ages for
// asset ownership.
class AgeRestriction {
    // Serialize returns the byte representation of the message.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Lower (uint8)
        {
            bytes_1.write(buf, this.lower, 'uint8');
        }
        // Upper (uint8)
        {
            bytes_1.write(buf, this.upper, 'uint8');
        }
        return buf.buf;
    }
    Write(buf) {
        // Lower (uint8)
        {
            this.lower = bytes_1.read(buf, 'uint8');
        }
        // Upper (uint8)
        {
            this.upper = bytes_1.read(buf, 'uint8');
        }
    }
    Validate() {
        // lower (uint8)
        {
        }
        // upper (uint8)
        {
        }
        return;
    }
    Equal(other) {
        // Lower (uint8)
        if (this.lower !== other.lower) {
            return false;
        }
        // Upper (uint8)
        if (this.upper !== other.upper) {
            return false;
        }
        return true;
    }
}
exports.AgeRestriction = AgeRestriction;
// Amendment An Amendment is used to describe the modification of a single
// field in a Contract or Asset, as defined in the ContractFormation and
// AssetCreation messages.
class Amendment {
    // Serialize returns the byte representation of the message.
    Serialize() {
        const buf = new util_1.default.Writer();
        // FieldIndex (uint8)
        {
            bytes_1.write(buf, this.field_index, 'uint8');
        }
        // Element (uint16)
        {
            bytes_1.write(buf, this.element, 'uint16');
        }
        // SubfieldIndex (uint8)
        {
            bytes_1.write(buf, this.subfield_index, 'uint8');
        }
        // SubfieldElement (uint16)
        {
            bytes_1.write(buf, this.subfield_element, 'uint16');
        }
        // Operation (uint8)
        {
            bytes_1.write(buf, this.operation, 'uint8');
        }
        // Data ([]byte)
        {
            bytes_1.WriteVarBin(buf, this.data, 32);
        }
        return buf.buf;
    }
    Write(buf) {
        // FieldIndex (uint8)
        {
            this.field_index = bytes_1.read(buf, 'uint8');
        }
        // Element (uint16)
        {
            this.element = bytes_1.read(buf, 'uint16');
        }
        // SubfieldIndex (uint8)
        {
            this.subfield_index = bytes_1.read(buf, 'uint8');
        }
        // SubfieldElement (uint16)
        {
            this.subfield_element = bytes_1.read(buf, 'uint16');
        }
        // Operation (uint8)
        {
            this.operation = bytes_1.read(buf, 'uint8');
        }
        // Data ([]byte)
        {
            this.data = bytes_1.ReadVarBin(buf, 32);
        }
    }
    Validate() {
        // field_index (uint8)
        {
        }
        // element (uint16)
        {
        }
        // subfield_index (uint8)
        {
        }
        // subfield_element (uint16)
        {
        }
        // operation (uint8)
        {
        }
        // data ([]byte)
        {
            if (this.data.length > (2 << 32) - 1) {
                return sprintf_js_1.sprintf('varbin field data too long %d/%d', this.data.length, (2 << 32) - 1);
            }
        }
        return;
    }
    Equal(other) {
        // FieldIndex (uint8)
        if (this.field_index !== other.field_index) {
            return false;
        }
        // Element (uint16)
        if (this.element !== other.element) {
            return false;
        }
        // SubfieldIndex (uint8)
        if (this.subfield_index !== other.subfield_index) {
            return false;
        }
        // SubfieldElement (uint16)
        if (this.subfield_element !== other.subfield_element) {
            return false;
        }
        // Operation (uint8)
        if (this.operation !== other.operation) {
            return false;
        }
        // Data ([]byte)
        if (!Buffer.compare(this.data, other.data)) {
            return false;
        }
        return true;
    }
}
exports.Amendment = Amendment;
// AssetReceiver An AssetReceiver is a quantity, address, and oracle
// signature. The quantity could be used to describe a number of tokens, or
// a value. The address is where the asset will be sent.
class AssetReceiver {
    constructor() {
        // The address receiving the tokens
        this.address = new protocol_types_1.PublicKeyHash(); // PublicKeyHash `json:"address,omitempty"`
    }
    // Serialize returns the byte representation of the message.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Address (PublicKeyHash)
        {
            buf.write(this.address.Serialize());
        }
        // Quantity (uint64)
        {
            bytes_1.write(buf, this.quantity, 'uint64');
        }
        // OracleSigAlgorithm (uint8)
        {
            bytes_1.write(buf, this.oracle_sig_algorithm, 'uint8');
        }
        // OracleIndex (uint8)
        if (this.oracle_sig_algorithm === 1) {
            bytes_1.write(buf, this.oracle_index, 'uint8');
        }
        // OracleConfirmationSig ([]byte)
        if (this.oracle_sig_algorithm === 1) {
            bytes_1.WriteVarBin(buf, this.oracle_confirmation_sig, 8);
        }
        // OracleSigBlockHeight (uint32)
        if (this.oracle_sig_algorithm === 1) {
            bytes_1.write(buf, this.oracle_sig_block_height, 'uint32');
        }
        return buf.buf;
    }
    Write(buf) {
        // Address (PublicKeyHash)
        {
            this.address.Write(buf);
        }
        // Quantity (uint64)
        {
            this.quantity = bytes_1.read(buf, 'uint64');
        }
        // OracleSigAlgorithm (uint8)
        {
            this.oracle_sig_algorithm = bytes_1.read(buf, 'uint8');
        }
        // OracleIndex (uint8)
        if (this.oracle_sig_algorithm === 1) {
            this.oracle_index = bytes_1.read(buf, 'uint8');
        }
        // OracleConfirmationSig ([]byte)
        if (this.oracle_sig_algorithm === 1) {
            this.oracle_confirmation_sig = bytes_1.ReadVarBin(buf, 8);
        }
        // OracleSigBlockHeight (uint32)
        if (this.oracle_sig_algorithm === 1) {
            this.oracle_sig_block_height = bytes_1.read(buf, 'uint32');
        }
    }
    Validate() {
        // address (PublicKeyHash)
        {
            const err = this.address.Validate();
            if (err) {
                return sprintf_js_1.sprintf('field address is invalid : %s', err);
            }
        }
        // quantity (uint64)
        {
        }
        // oracle_sig_algorithm (uint8)
        {
        }
        // oracle_index (uint8)
        {
        }
        // oracle_confirmation_sig ([]byte)
        {
            if (this.oracle_confirmation_sig.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varbin field oracle_confirmation_sig too long %d/%d', this.oracle_confirmation_sig.length, (2 << 8) - 1);
            }
        }
        // oracle_sig_block_height (uint32)
        {
        }
        return;
    }
    Equal(other) {
        // Address (PublicKeyHash)
        if (!this.address.Equal(other.address)) {
            return false;
        }
        // Quantity (uint64)
        if (this.quantity !== other.quantity) {
            return false;
        }
        // OracleSigAlgorithm (uint8)
        if (this.oracle_sig_algorithm !== other.oracle_sig_algorithm) {
            return false;
        }
        // OracleIndex (uint8)
        if (this.oracle_index !== other.oracle_index) {
            return false;
        }
        // OracleConfirmationSig ([]byte)
        if (!Buffer.compare(this.oracle_confirmation_sig, other.oracle_confirmation_sig)) {
            return false;
        }
        // OracleSigBlockHeight (uint32)
        if (this.oracle_sig_block_height !== other.oracle_sig_block_height) {
            return false;
        }
        return true;
    }
}
exports.AssetReceiver = AssetReceiver;
// AssetSettlement AssetSettlement is the data required to settle an asset
// transfer.
class AssetSettlement {
    constructor() {
        // A unique code that is used to identify the asset. It is generated by
        // hashing the contract public key hash and the asset index.
        // SHA256(contract PKH + asset index)
        this.asset_code = new protocol_types_1.AssetCode(); // AssetCode `json:"asset_code,omitempty"`
        // Each element contains the resulting token balance of Asset X for the
        // output Address, which is referred to by the index.
        this.settlements = []; // []QuantityIndex `json:"settlements,omitempty"`
    }
    // Serialize returns the byte representation of the message.
    Serialize() {
        const buf = new util_1.default.Writer();
        // ContractIndex (uint16)
        {
            bytes_1.write(buf, this.contract_index, 'uint16');
        }
        // AssetType (string)
        {
            bytes_1.WriteFixedChar(buf, this.asset_type, 3);
        }
        // AssetCode (AssetCode)
        {
            buf.write(this.asset_code.Serialize());
        }
        // Settlements ([]QuantityIndex)
        {
            if (!this.settlements)
                this.settlements = [];
            bytes_1.WriteVariableSize(buf, this.settlements.length, 8, 8);
            this.settlements.forEach(value => {
                buf.write(value.Serialize());
            });
        }
        return buf.buf;
    }
    Write(buf) {
        // ContractIndex (uint16)
        {
            this.contract_index = bytes_1.read(buf, 'uint16');
        }
        // AssetType (string)
        {
            this.asset_type = bytes_1.ReadFixedChar(buf, 3);
        }
        // AssetCode (AssetCode)
        {
            this.asset_code.Write(buf);
        }
        // Settlements ([]QuantityIndex)
        {
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            this.settlements = [...Array(size)]
                .map(() => {
                const newValue = new QuantityIndex();
                newValue.Write(buf);
                return newValue;
            });
        }
    }
    Validate() {
        // contract_index (uint16)
        {
        }
        // asset_type (string)
        {
            if (this.asset_type.length > 3) {
                return sprintf_js_1.sprintf('fixedchar field asset_type too long %d/%d', this.asset_type.length, 3);
            }
        }
        // asset_code (AssetCode)
        {
            const err = this.asset_code.Validate();
            if (err) {
                return sprintf_js_1.sprintf('field asset_code is invalid : %s', err);
            }
        }
        // settlements ([]QuantityIndex)
        {
            if (this.settlements.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('list field settlements has too many items %d/%d', this.settlements.length, (2 << 8) - 1);
            }
            this.settlements.forEach((value, i) => {
                const err = value.Validate();
                if (err) {
                    return sprintf_js_1.sprintf('list field settlements[%d] is invalid : %s', i, err);
                }
            });
        }
        return;
    }
    Equal(other) {
        // ContractIndex (uint16)
        if (this.contract_index !== other.contract_index) {
            return false;
        }
        // AssetType (string)
        if (this.asset_type !== other.asset_type) {
            return false;
        }
        // AssetCode (AssetCode)
        if (!this.asset_code.Equal(other.asset_code)) {
            return false;
        }
        // Settlements ([]QuantityIndex)
        if (this.settlements.length !== 0 || other.settlements.length !== 0) {
            if (this.settlements.length !== other.settlements.length) {
                return false;
            }
            if (!this.settlements.every((value, i) => {
                if (!value.Equal(other.settlements[i])) {
                    return false;
                }
                return true;
            }))
                return false;
        }
        return true;
    }
}
exports.AssetSettlement = AssetSettlement;
// AssetTransfer AssetTransfer is the data required to transfer an asset.
class AssetTransfer {
    constructor() {
        // A unique code that is used to identify the asset. It is generated by
        // hashing the contract public key hash and the asset index.
        // SHA256(contract PKH + asset index)
        this.asset_code = new protocol_types_1.AssetCode(); // AssetCode `json:"asset_code,omitempty"`
        // Each element has the value of tokens to be spent from the input
        // address, which is referred to by the index.
        this.asset_senders = []; // []QuantityIndex `json:"asset_senders,omitempty"`
        // Each element has the value of tokens to be received, the address, and
        // an oracle signature if required.
        this.asset_receivers = []; // []AssetReceiver `json:"asset_receivers,omitempty"`
    }
    // Serialize returns the byte representation of the message.
    Serialize() {
        const buf = new util_1.default.Writer();
        // ContractIndex (uint16)
        {
            bytes_1.write(buf, this.contract_index, 'uint16');
        }
        // AssetType (string)
        {
            bytes_1.WriteFixedChar(buf, this.asset_type, 3);
        }
        // AssetCode (AssetCode)
        {
            buf.write(this.asset_code.Serialize());
        }
        // AssetSenders ([]QuantityIndex)
        {
            if (!this.asset_senders)
                this.asset_senders = [];
            bytes_1.WriteVariableSize(buf, this.asset_senders.length, 8, 8);
            this.asset_senders.forEach(value => {
                buf.write(value.Serialize());
            });
        }
        // AssetReceivers ([]AssetReceiver)
        {
            if (!this.asset_receivers)
                this.asset_receivers = [];
            bytes_1.WriteVariableSize(buf, this.asset_receivers.length, 8, 8);
            this.asset_receivers.forEach(value => {
                buf.write(value.Serialize());
            });
        }
        return buf.buf;
    }
    Write(buf) {
        // ContractIndex (uint16)
        {
            this.contract_index = bytes_1.read(buf, 'uint16');
        }
        // AssetType (string)
        {
            this.asset_type = bytes_1.ReadFixedChar(buf, 3);
        }
        // AssetCode (AssetCode)
        {
            this.asset_code.Write(buf);
        }
        // AssetSenders ([]QuantityIndex)
        {
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            this.asset_senders = [...Array(size)]
                .map(() => {
                const newValue = new QuantityIndex();
                newValue.Write(buf);
                return newValue;
            });
        }
        // AssetReceivers ([]AssetReceiver)
        {
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            this.asset_receivers = [...Array(size)]
                .map(() => {
                const newValue = new AssetReceiver();
                newValue.Write(buf);
                return newValue;
            });
        }
    }
    Validate() {
        // contract_index (uint16)
        {
        }
        // asset_type (string)
        {
            if (this.asset_type.length > 3) {
                return sprintf_js_1.sprintf('fixedchar field asset_type too long %d/%d', this.asset_type.length, 3);
            }
        }
        // asset_code (AssetCode)
        {
            const err = this.asset_code.Validate();
            if (err) {
                return sprintf_js_1.sprintf('field asset_code is invalid : %s', err);
            }
        }
        // asset_senders ([]QuantityIndex)
        {
            if (this.asset_senders.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('list field asset_senders has too many items %d/%d', this.asset_senders.length, (2 << 8) - 1);
            }
            this.asset_senders.forEach((value, i) => {
                const err = value.Validate();
                if (err) {
                    return sprintf_js_1.sprintf('list field asset_senders[%d] is invalid : %s', i, err);
                }
            });
        }
        // asset_receivers ([]AssetReceiver)
        {
            if (this.asset_receivers.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('list field asset_receivers has too many items %d/%d', this.asset_receivers.length, (2 << 8) - 1);
            }
            this.asset_receivers.forEach((value, i) => {
                const err = value.Validate();
                if (err) {
                    return sprintf_js_1.sprintf('list field asset_receivers[%d] is invalid : %s', i, err);
                }
            });
        }
        return;
    }
    Equal(other) {
        // ContractIndex (uint16)
        if (this.contract_index !== other.contract_index) {
            return false;
        }
        // AssetType (string)
        if (this.asset_type !== other.asset_type) {
            return false;
        }
        // AssetCode (AssetCode)
        if (!this.asset_code.Equal(other.asset_code)) {
            return false;
        }
        // AssetSenders ([]QuantityIndex)
        if (this.asset_senders.length !== 0 || other.asset_senders.length !== 0) {
            if (this.asset_senders.length !== other.asset_senders.length) {
                return false;
            }
            if (!this.asset_senders.every((value, i) => {
                if (!value.Equal(other.asset_senders[i])) {
                    return false;
                }
                return true;
            }))
                return false;
        }
        // AssetReceivers ([]AssetReceiver)
        if (this.asset_receivers.length !== 0 || other.asset_receivers.length !== 0) {
            if (this.asset_receivers.length !== other.asset_receivers.length) {
                return false;
            }
            if (!this.asset_receivers.every((value, i) => {
                if (!value.Equal(other.asset_receivers[i])) {
                    return false;
                }
                return true;
            }))
                return false;
        }
        return true;
    }
}
exports.AssetTransfer = AssetTransfer;
// Document A file containing data.
class Document {
    // Serialize returns the byte representation of the message.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Name (string)
        {
            bytes_1.WriteVarChar(buf, this.name, 8);
        }
        // Type (string)
        {
            bytes_1.WriteVarChar(buf, this.type, 8);
        }
        // Algorithm (uint8)
        {
            bytes_1.write(buf, this.algorithm, 'uint8');
        }
        // Contents ([]byte)
        {
            bytes_1.WriteVarBin(buf, this.contents, 32);
        }
        return buf.buf;
    }
    Write(buf) {
        // Name (string)
        {
            this.name = bytes_1.ReadVarChar(buf, 8);
        }
        // Type (string)
        {
            this.type = bytes_1.ReadVarChar(buf, 8);
        }
        // Algorithm (uint8)
        {
            this.algorithm = bytes_1.read(buf, 'uint8');
        }
        // Contents ([]byte)
        {
            this.contents = bytes_1.ReadVarBin(buf, 32);
        }
    }
    Validate() {
        // name (string)
        {
            if (this.name.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field name too long %d/%d', this.name.length, (2 << 8) - 1);
            }
        }
        // type (string)
        {
            if (this.type.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field type too long %d/%d', this.type.length, (2 << 8) - 1);
            }
        }
        // algorithm (uint8)
        {
        }
        // contents ([]byte)
        {
            if (this.contents.length > (2 << 32) - 1) {
                return sprintf_js_1.sprintf('varbin field contents too long %d/%d', this.contents.length, (2 << 32) - 1);
            }
        }
        return;
    }
    Equal(other) {
        // Name (string)
        if (this.name !== other.name) {
            return false;
        }
        // Type (string)
        if (this.type !== other.type) {
            return false;
        }
        // Algorithm (uint8)
        if (this.algorithm !== other.algorithm) {
            return false;
        }
        // Contents ([]byte)
        if (!Buffer.compare(this.contents, other.contents)) {
            return false;
        }
        return true;
    }
}
exports.Document = Document;
// Entity Entity represents the details of a legal Entity, such as a public
// or private company, government agency, or and individual.
class Entity {
    constructor() {
        // A list of people that are in Administrative Roles for the Entity. eg.
        // Chair, Director, Managing Partner, etc.
        this.administration = []; // []Administrator `json:"administration,omitempty"`
        // A list of people in Management Roles for the Entity. e.g CEO, COO, CTO,
        // CFO, Secretary, Executive, etc.
        this.management = []; // []Manager `json:"management,omitempty"`
    }
    // Serialize returns the byte representation of the message.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Name (string)
        {
            bytes_1.WriteVarChar(buf, this.name, 8);
        }
        // Type (byte)
        {
            bytes_1.write(buf, this.type, 'byte');
        }
        // LEI (string)
        {
            bytes_1.WriteFixedChar(buf, this.lei, 20);
        }
        // AddressIncluded (bool)
        {
            bytes_1.write(buf, this.address_included, 'bool');
        }
        // UnitNumber (string)
        if (this.address_included) {
            bytes_1.WriteVarChar(buf, this.unit_number, 8);
        }
        // BuildingNumber (string)
        if (this.address_included) {
            bytes_1.WriteVarChar(buf, this.building_number, 8);
        }
        // Street (string)
        if (this.address_included) {
            bytes_1.WriteVarChar(buf, this.street, 16);
        }
        // SuburbCity (string)
        if (this.address_included) {
            bytes_1.WriteVarChar(buf, this.suburb_city, 8);
        }
        // TerritoryStateProvinceCode (string)
        if (this.address_included) {
            bytes_1.WriteFixedChar(buf, this.territory_state_province_code, 5);
        }
        // CountryCode (string)
        if (this.address_included) {
            bytes_1.WriteFixedChar(buf, this.country_code, 3);
        }
        // PostalZIPCode (string)
        if (this.address_included) {
            bytes_1.WriteFixedChar(buf, this.postal_zip_code, 12);
        }
        // EmailAddress (string)
        {
            bytes_1.WriteVarChar(buf, this.email_address, 8);
        }
        // PhoneNumber (string)
        {
            bytes_1.WriteVarChar(buf, this.phone_number, 8);
        }
        // Administration ([]Administrator)
        {
            if (!this.administration)
                this.administration = [];
            bytes_1.WriteVariableSize(buf, this.administration.length, 8, 8);
            this.administration.forEach(value => {
                buf.write(value.Serialize());
            });
        }
        // Management ([]Manager)
        {
            if (!this.management)
                this.management = [];
            bytes_1.WriteVariableSize(buf, this.management.length, 8, 8);
            this.management.forEach(value => {
                buf.write(value.Serialize());
            });
        }
        return buf.buf;
    }
    Write(buf) {
        // Name (string)
        {
            this.name = bytes_1.ReadVarChar(buf, 8);
        }
        // Type (byte)
        {
            this.type = bytes_1.read(buf, 'byte');
        }
        // LEI (string)
        {
            this.lei = bytes_1.ReadFixedChar(buf, 20);
        }
        // AddressIncluded (bool)
        {
            this.address_included = bytes_1.read(buf, 'bool');
        }
        // UnitNumber (string)
        if (this.address_included) {
            this.unit_number = bytes_1.ReadVarChar(buf, 8);
        }
        // BuildingNumber (string)
        if (this.address_included) {
            this.building_number = bytes_1.ReadVarChar(buf, 8);
        }
        // Street (string)
        if (this.address_included) {
            this.street = bytes_1.ReadVarChar(buf, 16);
        }
        // SuburbCity (string)
        if (this.address_included) {
            this.suburb_city = bytes_1.ReadVarChar(buf, 8);
        }
        // TerritoryStateProvinceCode (string)
        if (this.address_included) {
            this.territory_state_province_code = bytes_1.ReadFixedChar(buf, 5);
        }
        // CountryCode (string)
        if (this.address_included) {
            this.country_code = bytes_1.ReadFixedChar(buf, 3);
        }
        // PostalZIPCode (string)
        if (this.address_included) {
            this.postal_zip_code = bytes_1.ReadFixedChar(buf, 12);
        }
        // EmailAddress (string)
        {
            this.email_address = bytes_1.ReadVarChar(buf, 8);
        }
        // PhoneNumber (string)
        {
            this.phone_number = bytes_1.ReadVarChar(buf, 8);
        }
        // Administration ([]Administrator)
        {
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            this.administration = [...Array(size)]
                .map(() => {
                const newValue = new Administrator();
                newValue.Write(buf);
                return newValue;
            });
        }
        // Management ([]Manager)
        {
            const size = bytes_1.ReadVariableSize(buf, 8, 8);
            this.management = [...Array(size)]
                .map(() => {
                const newValue = new Manager();
                newValue.Write(buf);
                return newValue;
            });
        }
    }
    Validate() {
        // name (string)
        {
            if (this.name.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field name too long %d/%d', this.name.length, (2 << 8) - 1);
            }
        }
        // type (byte)
        {
            if (!resources_1.Resources.GetEntityType(this.type)) {
                return sprintf_js_1.sprintf('Invalid entity type value : %c', this.type);
            }
        }
        // lei (string)
        {
            if (this.lei.length > 20) {
                return sprintf_js_1.sprintf('fixedchar field lei too long %d/%d', this.lei.length, 20);
            }
        }
        // address_included (bool)
        {
        }
        // unit_number (string)
        if (this.address_included) {
            if (this.unit_number.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field unit_number too long %d/%d', this.unit_number.length, (2 << 8) - 1);
            }
        }
        // building_number (string)
        if (this.address_included) {
            if (this.building_number.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field building_number too long %d/%d', this.building_number.length, (2 << 8) - 1);
            }
        }
        // street (string)
        if (this.address_included) {
            if (this.street.length > (2 << 16) - 1) {
                return sprintf_js_1.sprintf('varchar field street too long %d/%d', this.street.length, (2 << 16) - 1);
            }
        }
        // suburb_city (string)
        if (this.address_included) {
            if (this.suburb_city.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field suburb_city too long %d/%d', this.suburb_city.length, (2 << 8) - 1);
            }
        }
        // territory_state_province_code (string)
        if (this.address_included) {
            if (this.territory_state_province_code.length > 5) {
                return sprintf_js_1.sprintf('fixedchar field territory_state_province_code too long %d/%d', this.territory_state_province_code.length, 5);
            }
        }
        // country_code (string)
        if (this.address_included) {
            if (this.country_code.length > 3) {
                return sprintf_js_1.sprintf('fixedchar field country_code too long %d/%d', this.country_code.length, 3);
            }
        }
        // postal_zip_code (string)
        if (this.address_included) {
            if (this.postal_zip_code.length > 12) {
                return sprintf_js_1.sprintf('fixedchar field postal_zip_code too long %d/%d', this.postal_zip_code.length, 12);
            }
        }
        // email_address (string)
        {
            if (this.email_address.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field email_address too long %d/%d', this.email_address.length, (2 << 8) - 1);
            }
        }
        // phone_number (string)
        {
            if (this.phone_number.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field phone_number too long %d/%d', this.phone_number.length, (2 << 8) - 1);
            }
        }
        // administration ([]Administrator)
        {
            if (this.administration.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('list field administration has too many items %d/%d', this.administration.length, (2 << 8) - 1);
            }
            this.administration.forEach((value, i) => {
                const err = value.Validate();
                if (err) {
                    return sprintf_js_1.sprintf('list field administration[%d] is invalid : %s', i, err);
                }
            });
        }
        // management ([]Manager)
        {
            if (this.management.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('list field management has too many items %d/%d', this.management.length, (2 << 8) - 1);
            }
            this.management.forEach((value, i) => {
                const err = value.Validate();
                if (err) {
                    return sprintf_js_1.sprintf('list field management[%d] is invalid : %s', i, err);
                }
            });
        }
        return;
    }
    Equal(other) {
        // Name (string)
        if (this.name !== other.name) {
            return false;
        }
        // Type (byte)
        if (this.type !== other.type) {
            return false;
        }
        // LEI (string)
        if (this.lei !== other.lei) {
            return false;
        }
        // AddressIncluded (bool)
        if (this.address_included !== other.address_included) {
            return false;
        }
        // UnitNumber (string)
        if (this.unit_number !== other.unit_number) {
            return false;
        }
        // BuildingNumber (string)
        if (this.building_number !== other.building_number) {
            return false;
        }
        // Street (string)
        if (this.street !== other.street) {
            return false;
        }
        // SuburbCity (string)
        if (this.suburb_city !== other.suburb_city) {
            return false;
        }
        // TerritoryStateProvinceCode (string)
        if (this.territory_state_province_code !== other.territory_state_province_code) {
            return false;
        }
        // CountryCode (string)
        if (this.country_code !== other.country_code) {
            return false;
        }
        // PostalZIPCode (string)
        if (this.postal_zip_code !== other.postal_zip_code) {
            return false;
        }
        // EmailAddress (string)
        if (this.email_address !== other.email_address) {
            return false;
        }
        // PhoneNumber (string)
        if (this.phone_number !== other.phone_number) {
            return false;
        }
        // Administration ([]Administrator)
        if (this.administration.length !== 0 || other.administration.length !== 0) {
            if (this.administration.length !== other.administration.length) {
                return false;
            }
            if (!this.administration.every((value, i) => {
                if (!value.Equal(other.administration[i])) {
                    return false;
                }
                return true;
            }))
                return false;
        }
        // Management ([]Manager)
        if (this.management.length !== 0 || other.management.length !== 0) {
            if (this.management.length !== other.management.length) {
                return false;
            }
            if (!this.management.every((value, i) => {
                if (!value.Equal(other.management[i])) {
                    return false;
                }
                return true;
            }))
                return false;
        }
        return true;
    }
}
exports.Entity = Entity;
// Manager Manager is used to refer to a role that is responsible for the
// Management of an Entity.
class Manager {
    // Serialize returns the byte representation of the message.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Type (uint8)
        {
            bytes_1.write(buf, this.type, 'uint8');
        }
        // Name (string)
        {
            bytes_1.WriteVarChar(buf, this.name, 8);
        }
        return buf.buf;
    }
    Write(buf) {
        // Type (uint8)
        {
            this.type = bytes_1.read(buf, 'uint8');
        }
        // Name (string)
        {
            this.name = bytes_1.ReadVarChar(buf, 8);
        }
    }
    Validate() {
        // type (uint8)
        {
            if (!resources_1.Resources.GetRoleType(this.type)) {
                return sprintf_js_1.sprintf('Invalid role value : %d', this.type);
            }
        }
        // name (string)
        {
            if (this.name.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field name too long %d/%d', this.name.length, (2 << 8) - 1);
            }
        }
        return;
    }
    Equal(other) {
        // Type (uint8)
        if (this.type !== other.type) {
            return false;
        }
        // Name (string)
        if (this.name !== other.name) {
            return false;
        }
        return true;
    }
}
exports.Manager = Manager;
// Oracle A Oracle defines the details of a public Oracle.
class Oracle {
    // Serialize returns the byte representation of the message.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Name (string)
        {
            bytes_1.WriteVarChar(buf, this.name, 8);
        }
        // URL (string)
        {
            bytes_1.WriteVarChar(buf, this.url, 8);
        }
        // PublicKey ([]byte)
        {
            bytes_1.WriteVarBin(buf, this.public_key, 8);
        }
        return buf.buf;
    }
    Write(buf) {
        // Name (string)
        {
            this.name = bytes_1.ReadVarChar(buf, 8);
        }
        // URL (string)
        {
            this.url = bytes_1.ReadVarChar(buf, 8);
        }
        // PublicKey ([]byte)
        {
            this.public_key = bytes_1.ReadVarBin(buf, 8);
        }
    }
    Validate() {
        // name (string)
        {
            if (this.name.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field name too long %d/%d', this.name.length, (2 << 8) - 1);
            }
        }
        // url (string)
        {
            if (this.url.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field url too long %d/%d', this.url.length, (2 << 8) - 1);
            }
        }
        // public_key ([]byte)
        {
            if (this.public_key.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varbin field public_key too long %d/%d', this.public_key.length, (2 << 8) - 1);
            }
        }
        return;
    }
    Equal(other) {
        // Name (string)
        if (this.name !== other.name) {
            return false;
        }
        // URL (string)
        if (this.url !== other.url) {
            return false;
        }
        // PublicKey ([]byte)
        if (!Buffer.compare(this.public_key, other.public_key)) {
            return false;
        }
        return true;
    }
}
exports.Oracle = Oracle;
// OutputTag A tag or category of an output used to categorize and organize
// outputs from different transactions.
class OutputTag {
    // Serialize returns the byte representation of the message.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Tag (string)
        {
            bytes_1.WriteVarChar(buf, this.tag, 8);
        }
        return buf.buf;
    }
    Write(buf) {
        // Tag (string)
        {
            this.tag = bytes_1.ReadVarChar(buf, 8);
        }
    }
    Validate() {
        // tag (string)
        {
            if (this.tag.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field tag too long %d/%d', this.tag.length, (2 << 8) - 1);
            }
        }
        return;
    }
    Equal(other) {
        // Tag (string)
        if (this.tag !== other.tag) {
            return false;
        }
        return true;
    }
}
exports.OutputTag = OutputTag;
// QuantityIndex A QuantityIndex contains a quantity, and an index. The
// quantity could be used to describe a number of tokens, or a value. The
// index is used to refer to an input or output index position.
class QuantityIndex {
    // Serialize returns the byte representation of the message.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Index (uint16)
        {
            bytes_1.write(buf, this.index, 'uint16');
        }
        // Quantity (uint64)
        {
            bytes_1.write(buf, this.quantity, 'uint64');
        }
        return buf.buf;
    }
    Write(buf) {
        // Index (uint16)
        {
            this.index = bytes_1.read(buf, 'uint16');
        }
        // Quantity (uint64)
        {
            this.quantity = bytes_1.read(buf, 'uint64');
        }
    }
    Validate() {
        // index (uint16)
        {
        }
        // quantity (uint64)
        {
        }
        return;
    }
    Equal(other) {
        // Index (uint16)
        if (this.index !== other.index) {
            return false;
        }
        // Quantity (uint64)
        if (this.quantity !== other.quantity) {
            return false;
        }
        return true;
    }
}
exports.QuantityIndex = QuantityIndex;
// TargetAddress A TargetAddress defines a public address and quantity.
class TargetAddress {
    constructor() {
        // Public address where the token balance will be changed.
        this.address = new protocol_types_1.PublicKeyHash(); // PublicKeyHash `json:"address,omitempty"`
    }
    // Serialize returns the byte representation of the message.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Address (PublicKeyHash)
        {
            buf.write(this.address.Serialize());
        }
        // Quantity (uint64)
        {
            bytes_1.write(buf, this.quantity, 'uint64');
        }
        return buf.buf;
    }
    Write(buf) {
        // Address (PublicKeyHash)
        {
            this.address.Write(buf);
        }
        // Quantity (uint64)
        {
            this.quantity = bytes_1.read(buf, 'uint64');
        }
    }
    Validate() {
        // address (PublicKeyHash)
        {
            const err = this.address.Validate();
            if (err) {
                return sprintf_js_1.sprintf('field address is invalid : %s', err);
            }
        }
        // quantity (uint64)
        {
        }
        return;
    }
    Equal(other) {
        // Address (PublicKeyHash)
        if (!this.address.Equal(other.address)) {
            return false;
        }
        // Quantity (uint64)
        if (this.quantity !== other.quantity) {
            return false;
        }
        return true;
    }
}
exports.TargetAddress = TargetAddress;
// VotingSystem A VotingSystem defines all details of a Voting System.
class VotingSystem {
    // Serialize returns the byte representation of the message.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Name (string)
        {
            bytes_1.WriteVarChar(buf, this.name, 8);
        }
        // VoteType (byte)
        {
            bytes_1.write(buf, this.vote_type, 'byte');
        }
        // TallyLogic (byte)
        {
            bytes_1.write(buf, this.tally_logic, 'byte');
        }
        // ThresholdPercentage (uint8)
        {
            bytes_1.write(buf, this.threshold_percentage, 'uint8');
        }
        // VoteMultiplierPermitted (bool)
        {
            bytes_1.write(buf, this.vote_multiplier_permitted, 'bool');
        }
        // HolderProposalFee (uint64)
        {
            bytes_1.write(buf, this.holder_proposal_fee, 'uint64');
        }
        return buf.buf;
    }
    Write(buf) {
        // Name (string)
        {
            this.name = bytes_1.ReadVarChar(buf, 8);
        }
        // VoteType (byte)
        {
            this.vote_type = bytes_1.read(buf, 'byte');
        }
        // TallyLogic (byte)
        {
            this.tally_logic = bytes_1.read(buf, 'byte');
        }
        // ThresholdPercentage (uint8)
        {
            this.threshold_percentage = bytes_1.read(buf, 'uint8');
        }
        // VoteMultiplierPermitted (bool)
        {
            this.vote_multiplier_permitted = bytes_1.read(buf, 'bool');
        }
        // HolderProposalFee (uint64)
        {
            this.holder_proposal_fee = bytes_1.read(buf, 'uint64');
        }
    }
    Validate() {
        // name (string)
        {
            if (this.name.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field name too long %d/%d', this.name.length, (2 << 8) - 1);
            }
        }
        // vote_type (byte)
        {
        }
        // tally_logic (byte)
        {
        }
        // threshold_percentage (uint8)
        {
        }
        // vote_multiplier_permitted (bool)
        {
        }
        // holder_proposal_fee (uint64)
        {
        }
        return;
    }
    Equal(other) {
        // Name (string)
        if (this.name !== other.name) {
            return false;
        }
        // VoteType (byte)
        if (this.vote_type !== other.vote_type) {
            return false;
        }
        // TallyLogic (byte)
        if (this.tally_logic !== other.tally_logic) {
            return false;
        }
        // ThresholdPercentage (uint8)
        if (this.threshold_percentage !== other.threshold_percentage) {
            return false;
        }
        // VoteMultiplierPermitted (bool)
        if (this.vote_multiplier_permitted !== other.vote_multiplier_permitted) {
            return false;
        }
        // HolderProposalFee (uint64)
        if (this.holder_proposal_fee !== other.holder_proposal_fee) {
            return false;
        }
        return true;
    }
}
exports.VotingSystem = VotingSystem;
