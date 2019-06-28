"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const sprintf_js_1 = require("sprintf-js");
const util_1 = __importDefault(require("@keyring/util"));
const bytes_1 = require("./bytes");
const resources_1 = require("./resources");
const protocol_types_1 = require("../src/protocol_types");
const field_types_1 = require("../src/field_types");
// AssetTypeLen is the size in bytes of all asset type variants.
const AssetTypeLen = 152;
// CodeCoupon identifies data as a Coupon message.
const CodeCoupon = 'COU';
// CodeCurrency identifies data as a Currency message.
const CodeCurrency = 'CUR';
// CodeLoyaltyPoints identifies data as a LoyaltyPoints message.
const CodeLoyaltyPoints = 'LOY';
// CodeMembership identifies data as a Membership message.
const CodeMembership = 'MEM';
// CodeShareCommon identifies data as a ShareCommon message.
const CodeShareCommon = 'SHC';
// CodeTicketAdmission identifies data as a TicketAdmission message.
const CodeTicketAdmission = 'TIC';
// AssetTypeMapping holds a mapping of asset codes to asset types.
function AssetTypeMapping(code) {
    switch (code) {
        case CodeCoupon:
            return new Coupon();
        case CodeCurrency:
            return new Currency();
        case CodeLoyaltyPoints:
            return new LoyaltyPoints();
        case CodeMembership:
            return new Membership();
        case CodeShareCommon:
            return new ShareCommon();
        case CodeTicketAdmission:
            return new TicketAdmission();
        default:
            return null;
    }
}
exports.AssetTypeMapping = AssetTypeMapping;
// Coupon asset type.
class Coupon {
    // Type returns the type identifer for this message.
    Type() {
        return CodeCoupon;
    }
    // Len returns the byte size of this message.
    Len() {
        return AssetTypeLen;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    Read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // Serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Version (uint8)
        {
            bytes_1.write(buf, this.version, 'uint8');
        }
        // RedeemingEntity (string)
        {
            bytes_1.WriteVarChar(buf, this.redeeming_entity, 8);
        }
        // IssueDate (Timestamp)
        {
            bytes_1.write(buf, this.issue_date, 'Timestamp');
        }
        // ExpiryDate (Timestamp)
        {
            bytes_1.write(buf, this.expiry_date, 'Timestamp');
        }
        // Value (uint64)
        {
            bytes_1.write(buf, this.value, 'uint64');
        }
        // Currency ([3]byte)
        {
            bytes_1.write(buf, this.currency, '[3]byte');
        }
        // Description (string)
        {
            bytes_1.WriteVarChar(buf, this.description, 16);
        }
        return buf.buf;
    }
    // Write populates the fields in Coupon from the byte slice
    Write(b) {
        const buf = new util_1.default.Reader(b);
        // Version (uint8)
        {
            this.version = bytes_1.read(buf, 'uint8');
        }
        // RedeemingEntity (string)
        {
            this.redeeming_entity = bytes_1.ReadVarChar(buf, 8);
        }
        // IssueDate (Timestamp)
        {
            this.issue_date = new protocol_types_1.Timestamp();
            this.issue_date.Write(buf);
        }
        // ExpiryDate (Timestamp)
        {
            this.expiry_date = new protocol_types_1.Timestamp();
            this.expiry_date.Write(buf);
        }
        // Value (uint64)
        {
            this.value = bytes_1.read(buf, 'uint64');
        }
        // Currency ([3]byte)
        {
            this.currency = bytes_1.read(buf, '[3]byte');
        }
        // Description (string)
        {
            this.description = bytes_1.ReadVarChar(buf, 16);
        }
        return b.length - buf.length;
    }
    Validate() {
        // Version (uint8)
        {
        }
        // RedeemingEntity (string)
        {
            if (this.redeeming_entity.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field redeeming_entity too long %d/%d', this.redeeming_entity.length, (2 << 8) - 1);
            }
        }
        // IssueDate (Timestamp)
        {
            const err = this.issue_date.Validate();
            if (err) {
                return sprintf_js_1.sprintf('field issue_date is invalid : %s', err);
            }
        }
        // ExpiryDate (Timestamp)
        {
            const err = this.expiry_date.Validate();
            if (err) {
                return sprintf_js_1.sprintf('field expiry_date is invalid : %s', err);
            }
        }
        // Value (uint64)
        {
        }
        // Currency ([3]byte)
        {
            if (!resources_1.Resources.GetCurrency(this.currency)) {
                return sprintf_js_1.sprintf('Invalid currency value : %d', this.currency);
            }
        }
        // Description (string)
        {
            if (this.description.length > (2 << 16) - 1) {
                return sprintf_js_1.sprintf('varchar field description too long %d/%d', this.description.length, (2 << 16) - 1);
            }
        }
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('version:%v', this.version));
        vals.push(sprintf_js_1.sprintf('redeeming_entity:%#+v', this.redeeming_entity));
        vals.push(sprintf_js_1.sprintf('issue_date:%#+v', this.issue_date));
        vals.push(sprintf_js_1.sprintf('expiry_date:%#+v', this.expiry_date));
        vals.push(sprintf_js_1.sprintf('value:%v', this.value));
        vals.push(sprintf_js_1.sprintf('currency:%#+v', this.currency));
        vals.push(sprintf_js_1.sprintf('description:%#+v', this.description));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
}
exports.Coupon = Coupon;
// Currency asset type.
class Currency {
    // Type returns the type identifer for this message.
    Type() {
        return CodeCurrency;
    }
    // Len returns the byte size of this message.
    Len() {
        return AssetTypeLen;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    Read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // Serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Version (uint8)
        {
            bytes_1.write(buf, this.version, 'uint8');
        }
        // ISOCode ([3]byte)
        {
            bytes_1.write(buf, this.iso_code, '[3]byte');
        }
        // MonetaryAuthority (string)
        {
            bytes_1.WriteVarChar(buf, this.monetary_authority, 8);
        }
        // Description (string)
        {
            bytes_1.WriteVarChar(buf, this.description, 16);
        }
        return buf.buf;
    }
    // Write populates the fields in Currency from the byte slice
    Write(b) {
        const buf = new util_1.default.Reader(b);
        // Version (uint8)
        {
            this.version = bytes_1.read(buf, 'uint8');
        }
        // ISOCode ([3]byte)
        {
            this.iso_code = bytes_1.read(buf, '[3]byte');
        }
        // MonetaryAuthority (string)
        {
            this.monetary_authority = bytes_1.ReadVarChar(buf, 8);
        }
        // Description (string)
        {
            this.description = bytes_1.ReadVarChar(buf, 16);
        }
        return b.length - buf.length;
    }
    Validate() {
        // Version (uint8)
        {
        }
        // ISOCode ([3]byte)
        {
            if (!resources_1.Resources.GetCurrency(this.iso_code)) {
                return sprintf_js_1.sprintf('Invalid currency value : %d', this.iso_code);
            }
        }
        // MonetaryAuthority (string)
        {
            if (this.monetary_authority.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field monetary_authority too long %d/%d', this.monetary_authority.length, (2 << 8) - 1);
            }
        }
        // Description (string)
        {
            if (this.description.length > (2 << 16) - 1) {
                return sprintf_js_1.sprintf('varchar field description too long %d/%d', this.description.length, (2 << 16) - 1);
            }
        }
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('version:%v', this.version));
        vals.push(sprintf_js_1.sprintf('iso_code:%#+v', this.iso_code));
        vals.push(sprintf_js_1.sprintf('monetary_authority:%#+v', this.monetary_authority));
        vals.push(sprintf_js_1.sprintf('description:%#+v', this.description));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
}
exports.Currency = Currency;
// LoyaltyPoints asset type.
class LoyaltyPoints {
    // Type returns the type identifer for this message.
    Type() {
        return CodeLoyaltyPoints;
    }
    // Len returns the byte size of this message.
    Len() {
        return AssetTypeLen;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    Read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // Serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Version (uint8)
        {
            bytes_1.write(buf, this.version, 'uint8');
        }
        // AgeRestriction (AgeRestriction)
        {
            bytes_1.write(buf, this.age_restriction, 'AgeRestriction');
        }
        // OfferName (string)
        {
            bytes_1.WriteVarChar(buf, this.offer_name, 8);
        }
        // ValidFrom (Timestamp)
        {
            bytes_1.write(buf, this.valid_from, 'Timestamp');
        }
        // ExpirationTimestamp (Timestamp)
        {
            bytes_1.write(buf, this.expiration_timestamp, 'Timestamp');
        }
        // Description (string)
        {
            bytes_1.WriteVarChar(buf, this.description, 16);
        }
        return buf.buf;
    }
    // Write populates the fields in LoyaltyPoints from the byte slice
    Write(b) {
        const buf = new util_1.default.Reader(b);
        // Version (uint8)
        {
            this.version = bytes_1.read(buf, 'uint8');
        }
        // AgeRestriction (AgeRestriction)
        {
            this.age_restriction = new field_types_1.AgeRestriction();
            this.age_restriction.Write(buf);
        }
        // OfferName (string)
        {
            this.offer_name = bytes_1.ReadVarChar(buf, 8);
        }
        // ValidFrom (Timestamp)
        {
            this.valid_from = new protocol_types_1.Timestamp();
            this.valid_from.Write(buf);
        }
        // ExpirationTimestamp (Timestamp)
        {
            this.expiration_timestamp = new protocol_types_1.Timestamp();
            this.expiration_timestamp.Write(buf);
        }
        // Description (string)
        {
            this.description = bytes_1.ReadVarChar(buf, 16);
        }
        return b.length - buf.length;
    }
    Validate() {
        // Version (uint8)
        {
        }
        // AgeRestriction (AgeRestriction)
        {
            const err = this.age_restriction.Validate();
            if (err) {
                return sprintf_js_1.sprintf('field age_restriction is invalid : %s', err);
            }
        }
        // OfferName (string)
        {
            if (this.offer_name.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field offer_name too long %d/%d', this.offer_name.length, (2 << 8) - 1);
            }
        }
        // ValidFrom (Timestamp)
        {
            const err = this.valid_from.Validate();
            if (err) {
                return sprintf_js_1.sprintf('field valid_from is invalid : %s', err);
            }
        }
        // ExpirationTimestamp (Timestamp)
        {
            const err = this.expiration_timestamp.Validate();
            if (err) {
                return sprintf_js_1.sprintf('field expiration_timestamp is invalid : %s', err);
            }
        }
        // Description (string)
        {
            if (this.description.length > (2 << 16) - 1) {
                return sprintf_js_1.sprintf('varchar field description too long %d/%d', this.description.length, (2 << 16) - 1);
            }
        }
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('version:%v', this.version));
        vals.push(sprintf_js_1.sprintf('age_restriction:%#+v', this.age_restriction));
        vals.push(sprintf_js_1.sprintf('offer_name:%#+v', this.offer_name));
        vals.push(sprintf_js_1.sprintf('valid_from:%#+v', this.valid_from));
        vals.push(sprintf_js_1.sprintf('expiration_timestamp:%#+v', this.expiration_timestamp));
        vals.push(sprintf_js_1.sprintf('description:%#+v', this.description));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
}
exports.LoyaltyPoints = LoyaltyPoints;
// Membership asset type.
class Membership {
    // Type returns the type identifer for this message.
    Type() {
        return CodeMembership;
    }
    // Len returns the byte size of this message.
    Len() {
        return AssetTypeLen;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    Read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // Serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Version (uint8)
        {
            bytes_1.write(buf, this.version, 'uint8');
        }
        // AgeRestriction (AgeRestriction)
        {
            bytes_1.write(buf, this.age_restriction, 'AgeRestriction');
        }
        // ValidFrom (Timestamp)
        {
            bytes_1.write(buf, this.valid_from, 'Timestamp');
        }
        // ExpirationTimestamp (Timestamp)
        {
            bytes_1.write(buf, this.expiration_timestamp, 'Timestamp');
        }
        // ID (string)
        {
            bytes_1.WriteVarChar(buf, this.id, 8);
        }
        // MembershipClass (string)
        {
            bytes_1.WriteVarChar(buf, this.membership_class, 8);
        }
        // RoleType (string)
        {
            bytes_1.WriteVarChar(buf, this.role_type, 8);
        }
        // MembershipType (string)
        {
            bytes_1.WriteVarChar(buf, this.membership_type, 8);
        }
        // Description (string)
        {
            bytes_1.WriteVarChar(buf, this.description, 16);
        }
        return buf.buf;
    }
    // Write populates the fields in Membership from the byte slice
    Write(b) {
        const buf = new util_1.default.Reader(b);
        // Version (uint8)
        {
            this.version = bytes_1.read(buf, 'uint8');
        }
        // AgeRestriction (AgeRestriction)
        {
            this.age_restriction = new field_types_1.AgeRestriction();
            this.age_restriction.Write(buf);
        }
        // ValidFrom (Timestamp)
        {
            this.valid_from = new protocol_types_1.Timestamp();
            this.valid_from.Write(buf);
        }
        // ExpirationTimestamp (Timestamp)
        {
            this.expiration_timestamp = new protocol_types_1.Timestamp();
            this.expiration_timestamp.Write(buf);
        }
        // ID (string)
        {
            this.id = bytes_1.ReadVarChar(buf, 8);
        }
        // MembershipClass (string)
        {
            this.membership_class = bytes_1.ReadVarChar(buf, 8);
        }
        // RoleType (string)
        {
            this.role_type = bytes_1.ReadVarChar(buf, 8);
        }
        // MembershipType (string)
        {
            this.membership_type = bytes_1.ReadVarChar(buf, 8);
        }
        // Description (string)
        {
            this.description = bytes_1.ReadVarChar(buf, 16);
        }
        return b.length - buf.length;
    }
    Validate() {
        // Version (uint8)
        {
        }
        // AgeRestriction (AgeRestriction)
        {
            const err = this.age_restriction.Validate();
            if (err) {
                return sprintf_js_1.sprintf('field age_restriction is invalid : %s', err);
            }
        }
        // ValidFrom (Timestamp)
        {
            const err = this.valid_from.Validate();
            if (err) {
                return sprintf_js_1.sprintf('field valid_from is invalid : %s', err);
            }
        }
        // ExpirationTimestamp (Timestamp)
        {
            const err = this.expiration_timestamp.Validate();
            if (err) {
                return sprintf_js_1.sprintf('field expiration_timestamp is invalid : %s', err);
            }
        }
        // ID (string)
        {
            if (this.id.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field id too long %d/%d', this.id.length, (2 << 8) - 1);
            }
        }
        // MembershipClass (string)
        {
            if (this.membership_class.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field membership_class too long %d/%d', this.membership_class.length, (2 << 8) - 1);
            }
        }
        // RoleType (string)
        {
            if (this.role_type.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field role_type too long %d/%d', this.role_type.length, (2 << 8) - 1);
            }
        }
        // MembershipType (string)
        {
            if (this.membership_type.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field membership_type too long %d/%d', this.membership_type.length, (2 << 8) - 1);
            }
        }
        // Description (string)
        {
            if (this.description.length > (2 << 16) - 1) {
                return sprintf_js_1.sprintf('varchar field description too long %d/%d', this.description.length, (2 << 16) - 1);
            }
        }
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('version:%v', this.version));
        vals.push(sprintf_js_1.sprintf('age_restriction:%#+v', this.age_restriction));
        vals.push(sprintf_js_1.sprintf('valid_from:%#+v', this.valid_from));
        vals.push(sprintf_js_1.sprintf('expiration_timestamp:%#+v', this.expiration_timestamp));
        vals.push(sprintf_js_1.sprintf('id:%#+v', this.id));
        vals.push(sprintf_js_1.sprintf('membership_class:%#+v', this.membership_class));
        vals.push(sprintf_js_1.sprintf('role_type:%#+v', this.role_type));
        vals.push(sprintf_js_1.sprintf('membership_type:%#+v', this.membership_type));
        vals.push(sprintf_js_1.sprintf('description:%#+v', this.description));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
}
exports.Membership = Membership;
// ShareCommon asset type.
class ShareCommon {
    // Type returns the type identifer for this message.
    Type() {
        return CodeShareCommon;
    }
    // Len returns the byte size of this message.
    Len() {
        return AssetTypeLen;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    Read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // Serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Version (uint8)
        {
            bytes_1.write(buf, this.version, 'uint8');
        }
        // TransferLockout (Timestamp)
        {
            bytes_1.write(buf, this.transfer_lockout, 'Timestamp');
        }
        // Ticker (string)
        {
            bytes_1.WriteFixedChar(buf, this.ticker, 5);
        }
        // ISIN (string)
        {
            bytes_1.WriteFixedChar(buf, this.isin, 12);
        }
        // Description (string)
        {
            bytes_1.WriteVarChar(buf, this.description, 16);
        }
        return buf.buf;
    }
    // Write populates the fields in ShareCommon from the byte slice
    Write(b) {
        const buf = new util_1.default.Reader(b);
        // Version (uint8)
        {
            this.version = bytes_1.read(buf, 'uint8');
        }
        // TransferLockout (Timestamp)
        {
            this.transfer_lockout = new protocol_types_1.Timestamp();
            this.transfer_lockout.Write(buf);
        }
        // Ticker (string)
        {
            this.ticker = bytes_1.ReadFixedChar(buf, 5);
        }
        // ISIN (string)
        {
            this.isin = bytes_1.ReadFixedChar(buf, 12);
        }
        // Description (string)
        {
            this.description = bytes_1.ReadVarChar(buf, 16);
        }
        return b.length - buf.length;
    }
    Validate() {
        // Version (uint8)
        {
        }
        // TransferLockout (Timestamp)
        {
            const err = this.transfer_lockout.Validate();
            if (err) {
                return sprintf_js_1.sprintf('field transfer_lockout is invalid : %s', err);
            }
        }
        // Ticker (string)
        {
            if (this.ticker.length > 5) {
                return sprintf_js_1.sprintf('fixedchar field ticker too long %d/%d', this.ticker.length, 5);
            }
        }
        // ISIN (string)
        {
            if (this.isin.length > 12) {
                return sprintf_js_1.sprintf('fixedchar field isin too long %d/%d', this.isin.length, 12);
            }
        }
        // Description (string)
        {
            if (this.description.length > (2 << 16) - 1) {
                return sprintf_js_1.sprintf('varchar field description too long %d/%d', this.description.length, (2 << 16) - 1);
            }
        }
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('version:%v', this.version));
        vals.push(sprintf_js_1.sprintf('transfer_lockout:%#+v', this.transfer_lockout));
        vals.push(sprintf_js_1.sprintf('ticker:%#+v', this.ticker));
        vals.push(sprintf_js_1.sprintf('isin:%#+v', this.isin));
        vals.push(sprintf_js_1.sprintf('description:%#+v', this.description));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
}
exports.ShareCommon = ShareCommon;
// TicketAdmission asset type.
class TicketAdmission {
    // Type returns the type identifer for this message.
    Type() {
        return CodeTicketAdmission;
    }
    // Len returns the byte size of this message.
    Len() {
        return AssetTypeLen;
    }
    // Read implements the io.Reader interface, writing the receiver to the
    // []byte.
    Read(b) {
        const data = this.Serialize();
        b.fill(data);
        return data.length;
    }
    // Serialize returns the full OP_RETURN payload bytes.
    Serialize() {
        const buf = new util_1.default.Writer();
        // Version (uint8)
        {
            bytes_1.write(buf, this.version, 'uint8');
        }
        // AgeRestriction (AgeRestriction)
        {
            bytes_1.write(buf, this.age_restriction, 'AgeRestriction');
        }
        // AdmissionType (string)
        {
            bytes_1.WriteFixedChar(buf, this.admission_type, 3);
        }
        // Venue (string)
        {
            bytes_1.WriteVarChar(buf, this.venue, 8);
        }
        // Class (string)
        {
            bytes_1.WriteVarChar(buf, this.class, 8);
        }
        // Area (string)
        {
            bytes_1.WriteVarChar(buf, this.area, 8);
        }
        // Seat (string)
        {
            bytes_1.WriteVarChar(buf, this.seat, 8);
        }
        // StartTimeDate (Timestamp)
        {
            bytes_1.write(buf, this.start_time_date, 'Timestamp');
        }
        // ValidFrom (Timestamp)
        {
            bytes_1.write(buf, this.valid_from, 'Timestamp');
        }
        // ExpirationTimestamp (Timestamp)
        {
            bytes_1.write(buf, this.expiration_timestamp, 'Timestamp');
        }
        // Description (string)
        {
            bytes_1.WriteVarChar(buf, this.description, 16);
        }
        return buf.buf;
    }
    // Write populates the fields in TicketAdmission from the byte slice
    Write(b) {
        const buf = new util_1.default.Reader(b);
        // Version (uint8)
        {
            this.version = bytes_1.read(buf, 'uint8');
        }
        // AgeRestriction (AgeRestriction)
        {
            this.age_restriction = new field_types_1.AgeRestriction();
            this.age_restriction.Write(buf);
        }
        // AdmissionType (string)
        {
            this.admission_type = bytes_1.ReadFixedChar(buf, 3);
        }
        // Venue (string)
        {
            this.venue = bytes_1.ReadVarChar(buf, 8);
        }
        // Class (string)
        {
            this.class = bytes_1.ReadVarChar(buf, 8);
        }
        // Area (string)
        {
            this.area = bytes_1.ReadVarChar(buf, 8);
        }
        // Seat (string)
        {
            this.seat = bytes_1.ReadVarChar(buf, 8);
        }
        // StartTimeDate (Timestamp)
        {
            this.start_time_date = new protocol_types_1.Timestamp();
            this.start_time_date.Write(buf);
        }
        // ValidFrom (Timestamp)
        {
            this.valid_from = new protocol_types_1.Timestamp();
            this.valid_from.Write(buf);
        }
        // ExpirationTimestamp (Timestamp)
        {
            this.expiration_timestamp = new protocol_types_1.Timestamp();
            this.expiration_timestamp.Write(buf);
        }
        // Description (string)
        {
            this.description = bytes_1.ReadVarChar(buf, 16);
        }
        return b.length - buf.length;
    }
    Validate() {
        // Version (uint8)
        {
        }
        // AgeRestriction (AgeRestriction)
        {
            const err = this.age_restriction.Validate();
            if (err) {
                return sprintf_js_1.sprintf('field age_restriction is invalid : %s', err);
            }
        }
        // AdmissionType (string)
        {
            if (this.admission_type.length > 3) {
                return sprintf_js_1.sprintf('fixedchar field admission_type too long %d/%d', this.admission_type.length, 3);
            }
        }
        // Venue (string)
        {
            if (this.venue.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field venue too long %d/%d', this.venue.length, (2 << 8) - 1);
            }
        }
        // Class (string)
        {
            if (this.class.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field class too long %d/%d', this.class.length, (2 << 8) - 1);
            }
        }
        // Area (string)
        {
            if (this.area.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field area too long %d/%d', this.area.length, (2 << 8) - 1);
            }
        }
        // Seat (string)
        {
            if (this.seat.length > (2 << 8) - 1) {
                return sprintf_js_1.sprintf('varchar field seat too long %d/%d', this.seat.length, (2 << 8) - 1);
            }
        }
        // StartTimeDate (Timestamp)
        {
            const err = this.start_time_date.Validate();
            if (err) {
                return sprintf_js_1.sprintf('field start_time_date is invalid : %s', err);
            }
        }
        // ValidFrom (Timestamp)
        {
            const err = this.valid_from.Validate();
            if (err) {
                return sprintf_js_1.sprintf('field valid_from is invalid : %s', err);
            }
        }
        // ExpirationTimestamp (Timestamp)
        {
            const err = this.expiration_timestamp.Validate();
            if (err) {
                return sprintf_js_1.sprintf('field expiration_timestamp is invalid : %s', err);
            }
        }
        // Description (string)
        {
            if (this.description.length > (2 << 16) - 1) {
                return sprintf_js_1.sprintf('varchar field description too long %d/%d', this.description.length, (2 << 16) - 1);
            }
        }
    }
    toString() {
        const vals = [];
        vals.push(sprintf_js_1.sprintf('version:%v', this.version));
        vals.push(sprintf_js_1.sprintf('age_restriction:%#+v', this.age_restriction));
        vals.push(sprintf_js_1.sprintf('admission_type:%#+v', this.admission_type));
        vals.push(sprintf_js_1.sprintf('venue:%#+v', this.venue));
        vals.push(sprintf_js_1.sprintf('class:%#+v', this.class));
        vals.push(sprintf_js_1.sprintf('area:%#+v', this.area));
        vals.push(sprintf_js_1.sprintf('seat:%#+v', this.seat));
        vals.push(sprintf_js_1.sprintf('start_time_date:%#+v', this.start_time_date));
        vals.push(sprintf_js_1.sprintf('valid_from:%#+v', this.valid_from));
        vals.push(sprintf_js_1.sprintf('expiration_timestamp:%#+v', this.expiration_timestamp));
        vals.push(sprintf_js_1.sprintf('description:%#+v', this.description));
        return sprintf_js_1.sprintf('{%s}', vals.join(' '));
    }
}
exports.TicketAdmission = TicketAdmission;
