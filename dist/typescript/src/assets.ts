import {sprintf} from 'sprintf-js';
import _ from '@keyring/util';
import {write, read, ReadVarChar, ReadFixedChar,
	WriteVarChar, WriteFixedChar} from './bytes';
import { Resources } from './resources';


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


// AssetPayload is the interface for payloads within asset actions.
export interface AssetPayload {
	Type(): string;
	Serialize(): Buffer;
	Write(Buffer): number;
	Validate(): string;
}

// AssetTypeMapping holds a mapping of asset codes to asset types.
export function AssetTypeMapping(code: string): AssetPayload {
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


// Coupon asset type.
export class Coupon implements AssetPayload {

// The version number that this asset payload adheres to.
	version; // uint8 `json:"version,omitempty"`
// The entity responsible for redemption of this coupon.
	redeeming_entity; // string `json:"redeeming_entity,omitempty"`
// 
	issue_date; // Timestamp `json:"issue_date,omitempty"`
// 
	expiry_date; // Timestamp `json:"expiry_date,omitempty"`
// 
	value; // uint64 `json:"value,omitempty"`
// International Organization for Standardization code for Currency. Currency for coupon. From resources/currency.
	currency; // [3]byte `json:"currency,omitempty"`
// 
	description; // string `json:"description,omitempty"`


	// Type returns the type identifer for this message.
	Type(): string {
		return CodeCoupon;
	}

	// Len returns the byte size of this message.
	Len(): number {
		return AssetTypeLen;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	Read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// Serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
	

		// Version (uint8)
		{
			write(buf, this.version, 'uint8');
		}
	

		// RedeemingEntity (string)
		{
			WriteVarChar(buf, this.redeeming_entity, 8);
		}
	

		// IssueDate (Timestamp)
		{
			write(buf, this.issue_date, 'Timestamp');
		}
	

		// ExpiryDate (Timestamp)
		{
			write(buf, this.expiry_date, 'Timestamp');
		}
	

		// Value (uint64)
		{
			write(buf, this.value, 'uint64');
		}
	

		// Currency ([3]byte)
		{
			write(buf, this.currency, '[3]byte');
		}
	

		// Description (string)
		{
			WriteVarChar(buf, this.description, 16);
		}
	
		return buf.buf;
	}

	// Write populates the fields in Coupon from the byte slice
	Write(b: Buffer): number {
		const buf = new _.Reader(b);

		// Version (uint8)
		{
			this.version = read(buf, 'uint8');
	
		}
	

		// RedeemingEntity (string)
		{
			this.redeeming_entity = ReadVarChar(buf, 8);
		}
	

		// IssueDate (Timestamp)
		{
			this.issue_date.Write(buf);
	
		}
	

		// ExpiryDate (Timestamp)
		{
			this.expiry_date.Write(buf);
	
		}
	

		// Value (uint64)
		{
			this.value = read(buf, 'uint64');
	
		}
	

		// Currency ([3]byte)
		{
			this.currency = read(buf, '[3]byte');
	
		}
	

		// Description (string)
		{
			this.description = ReadVarChar(buf, 16);
		}
	
		return b.length - buf.length;
	}

	Validate(): string {

		// Version (uint8)
		{
		}
	

		// RedeemingEntity (string)
		{
			if (this.redeeming_entity.length > (2 << 8) - 1) {
				return sprintf('varchar field redeeming_entity too long %d/%d', this.redeeming_entity.length, (2 << 8) - 1);
			}
		}
	

		// IssueDate (Timestamp)
		{
			const err = this.issue_date.Validate();
			if (err) {
				return sprintf('field issue_date is invalid : %s', err);
			}
	
		}
	

		// ExpiryDate (Timestamp)
		{
			const err = this.expiry_date.Validate();
			if (err) {
				return sprintf('field expiry_date is invalid : %s', err);
			}
	
		}
	

		// Value (uint64)
		{
		}
	

		// Currency ([3]byte)
		{
			if (!Resources.GetCurrency(this.currency)) {
				return sprintf('Invalid currency value : %d', this.currency);
			}
		}
	

		// Description (string)
		{
			if (this.description.length > (2 << 16) - 1) {
				return sprintf('varchar field description too long %d/%d', this.description.length, (2 << 16) - 1);
			}
		}
	
	}

	toString(): string {
		const vals = [];
	
		vals.push(sprintf('version:%v', this.version));
		vals.push(sprintf('redeeming_entity:%#+v', this.redeeming_entity));
		vals.push(sprintf('issue_date:%#+v', this.issue_date));
		vals.push(sprintf('expiry_date:%#+v', this.expiry_date));
		vals.push(sprintf('value:%v', this.value));
		vals.push(sprintf('currency:%#+v', this.currency));
		vals.push(sprintf('description:%#+v', this.description));

		return sprintf('{%s}', vals.join(' '));
	}
}


// Currency asset type.
export class Currency implements AssetPayload {

// The version number that this asset payload adheres to.
	version; // uint8 `json:"version,omitempty"`
// International Organization for Standardization code for Currency. (Specification/Resources)
	iso_code; // [3]byte `json:"iso_code,omitempty"`
// 
	monetary_authority; // string `json:"monetary_authority,omitempty"`
// 
	description; // string `json:"description,omitempty"`


	// Type returns the type identifer for this message.
	Type(): string {
		return CodeCurrency;
	}

	// Len returns the byte size of this message.
	Len(): number {
		return AssetTypeLen;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	Read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// Serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
	

		// Version (uint8)
		{
			write(buf, this.version, 'uint8');
		}
	

		// ISOCode ([3]byte)
		{
			write(buf, this.iso_code, '[3]byte');
		}
	

		// MonetaryAuthority (string)
		{
			WriteVarChar(buf, this.monetary_authority, 8);
		}
	

		// Description (string)
		{
			WriteVarChar(buf, this.description, 16);
		}
	
		return buf.buf;
	}

	// Write populates the fields in Currency from the byte slice
	Write(b: Buffer): number {
		const buf = new _.Reader(b);

		// Version (uint8)
		{
			this.version = read(buf, 'uint8');
	
		}
	

		// ISOCode ([3]byte)
		{
			this.iso_code = read(buf, '[3]byte');
	
		}
	

		// MonetaryAuthority (string)
		{
			this.monetary_authority = ReadVarChar(buf, 8);
		}
	

		// Description (string)
		{
			this.description = ReadVarChar(buf, 16);
		}
	
		return b.length - buf.length;
	}

	Validate(): string {

		// Version (uint8)
		{
		}
	

		// ISOCode ([3]byte)
		{
			if (!Resources.GetCurrency(this.iso_code)) {
				return sprintf('Invalid currency value : %d', this.iso_code);
			}
		}
	

		// MonetaryAuthority (string)
		{
			if (this.monetary_authority.length > (2 << 8) - 1) {
				return sprintf('varchar field monetary_authority too long %d/%d', this.monetary_authority.length, (2 << 8) - 1);
			}
		}
	

		// Description (string)
		{
			if (this.description.length > (2 << 16) - 1) {
				return sprintf('varchar field description too long %d/%d', this.description.length, (2 << 16) - 1);
			}
		}
	
	}

	toString(): string {
		const vals = [];
	
		vals.push(sprintf('version:%v', this.version));
		vals.push(sprintf('iso_code:%#+v', this.iso_code));
		vals.push(sprintf('monetary_authority:%#+v', this.monetary_authority));
		vals.push(sprintf('description:%#+v', this.description));

		return sprintf('{%s}', vals.join(' '));
	}
}


// LoyaltyPoints asset type.
export class LoyaltyPoints implements AssetPayload {

// The version number that this asset payload adheres to.
	version; // uint8 `json:"version,omitempty"`
// Age restriction is used to specify required ages for asset ownership.
	age_restriction; // AgeRestriction `json:"age_restriction,omitempty"`
// 
	offer_name; // string `json:"offer_name,omitempty"`
// 
	valid_from; // Timestamp `json:"valid_from,omitempty"`
// 
	expiration_timestamp; // Timestamp `json:"expiration_timestamp,omitempty"`
// 
	description; // string `json:"description,omitempty"`


	// Type returns the type identifer for this message.
	Type(): string {
		return CodeLoyaltyPoints;
	}

	// Len returns the byte size of this message.
	Len(): number {
		return AssetTypeLen;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	Read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// Serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
	

		// Version (uint8)
		{
			write(buf, this.version, 'uint8');
		}
	

		// AgeRestriction (AgeRestriction)
		{
			write(buf, this.age_restriction, 'AgeRestriction');
		}
	

		// OfferName (string)
		{
			WriteVarChar(buf, this.offer_name, 8);
		}
	

		// ValidFrom (Timestamp)
		{
			write(buf, this.valid_from, 'Timestamp');
		}
	

		// ExpirationTimestamp (Timestamp)
		{
			write(buf, this.expiration_timestamp, 'Timestamp');
		}
	

		// Description (string)
		{
			WriteVarChar(buf, this.description, 16);
		}
	
		return buf.buf;
	}

	// Write populates the fields in LoyaltyPoints from the byte slice
	Write(b: Buffer): number {
		const buf = new _.Reader(b);

		// Version (uint8)
		{
			this.version = read(buf, 'uint8');
	
		}
	

		// AgeRestriction (AgeRestriction)
		{
			this.age_restriction.Write(buf);
	
		}
	

		// OfferName (string)
		{
			this.offer_name = ReadVarChar(buf, 8);
		}
	

		// ValidFrom (Timestamp)
		{
			this.valid_from.Write(buf);
	
		}
	

		// ExpirationTimestamp (Timestamp)
		{
			this.expiration_timestamp.Write(buf);
	
		}
	

		// Description (string)
		{
			this.description = ReadVarChar(buf, 16);
		}
	
		return b.length - buf.length;
	}

	Validate(): string {

		// Version (uint8)
		{
		}
	

		// AgeRestriction (AgeRestriction)
		{
			const err = this.age_restriction.Validate();
			if (err) {
				return sprintf('field age_restriction is invalid : %s', err);
			}
	
		}
	

		// OfferName (string)
		{
			if (this.offer_name.length > (2 << 8) - 1) {
				return sprintf('varchar field offer_name too long %d/%d', this.offer_name.length, (2 << 8) - 1);
			}
		}
	

		// ValidFrom (Timestamp)
		{
			const err = this.valid_from.Validate();
			if (err) {
				return sprintf('field valid_from is invalid : %s', err);
			}
	
		}
	

		// ExpirationTimestamp (Timestamp)
		{
			const err = this.expiration_timestamp.Validate();
			if (err) {
				return sprintf('field expiration_timestamp is invalid : %s', err);
			}
	
		}
	

		// Description (string)
		{
			if (this.description.length > (2 << 16) - 1) {
				return sprintf('varchar field description too long %d/%d', this.description.length, (2 << 16) - 1);
			}
		}
	
	}

	toString(): string {
		const vals = [];
	
		vals.push(sprintf('version:%v', this.version));
		vals.push(sprintf('age_restriction:%#+v', this.age_restriction));
		vals.push(sprintf('offer_name:%#+v', this.offer_name));
		vals.push(sprintf('valid_from:%#+v', this.valid_from));
		vals.push(sprintf('expiration_timestamp:%#+v', this.expiration_timestamp));
		vals.push(sprintf('description:%#+v', this.description));

		return sprintf('{%s}', vals.join(' '));
	}
}


// Membership asset type.
export class Membership implements AssetPayload {

// The version number that this asset payload adheres to.
	version; // uint8 `json:"version,omitempty"`
// Age restriction is used to specify required ages for asset ownership.
	age_restriction; // AgeRestriction `json:"age_restriction,omitempty"`
// 
	valid_from; // Timestamp `json:"valid_from,omitempty"`
// 
	expiration_timestamp; // Timestamp `json:"expiration_timestamp,omitempty"`
// 
	id; // string `json:"id,omitempty"`
// 
	membership_class; // string `json:"membership_class,omitempty"`
// 
	role_type; // string `json:"role_type,omitempty"`
// 
	membership_type; // string `json:"membership_type,omitempty"`
// 
	description; // string `json:"description,omitempty"`


	// Type returns the type identifer for this message.
	Type(): string {
		return CodeMembership;
	}

	// Len returns the byte size of this message.
	Len(): number {
		return AssetTypeLen;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	Read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// Serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
	

		// Version (uint8)
		{
			write(buf, this.version, 'uint8');
		}
	

		// AgeRestriction (AgeRestriction)
		{
			write(buf, this.age_restriction, 'AgeRestriction');
		}
	

		// ValidFrom (Timestamp)
		{
			write(buf, this.valid_from, 'Timestamp');
		}
	

		// ExpirationTimestamp (Timestamp)
		{
			write(buf, this.expiration_timestamp, 'Timestamp');
		}
	

		// ID (string)
		{
			WriteVarChar(buf, this.id, 8);
		}
	

		// MembershipClass (string)
		{
			WriteVarChar(buf, this.membership_class, 8);
		}
	

		// RoleType (string)
		{
			WriteVarChar(buf, this.role_type, 8);
		}
	

		// MembershipType (string)
		{
			WriteVarChar(buf, this.membership_type, 8);
		}
	

		// Description (string)
		{
			WriteVarChar(buf, this.description, 16);
		}
	
		return buf.buf;
	}

	// Write populates the fields in Membership from the byte slice
	Write(b: Buffer): number {
		const buf = new _.Reader(b);

		// Version (uint8)
		{
			this.version = read(buf, 'uint8');
	
		}
	

		// AgeRestriction (AgeRestriction)
		{
			this.age_restriction.Write(buf);
	
		}
	

		// ValidFrom (Timestamp)
		{
			this.valid_from.Write(buf);
	
		}
	

		// ExpirationTimestamp (Timestamp)
		{
			this.expiration_timestamp.Write(buf);
	
		}
	

		// ID (string)
		{
			this.id = ReadVarChar(buf, 8);
		}
	

		// MembershipClass (string)
		{
			this.membership_class = ReadVarChar(buf, 8);
		}
	

		// RoleType (string)
		{
			this.role_type = ReadVarChar(buf, 8);
		}
	

		// MembershipType (string)
		{
			this.membership_type = ReadVarChar(buf, 8);
		}
	

		// Description (string)
		{
			this.description = ReadVarChar(buf, 16);
		}
	
		return b.length - buf.length;
	}

	Validate(): string {

		// Version (uint8)
		{
		}
	

		// AgeRestriction (AgeRestriction)
		{
			const err = this.age_restriction.Validate();
			if (err) {
				return sprintf('field age_restriction is invalid : %s', err);
			}
	
		}
	

		// ValidFrom (Timestamp)
		{
			const err = this.valid_from.Validate();
			if (err) {
				return sprintf('field valid_from is invalid : %s', err);
			}
	
		}
	

		// ExpirationTimestamp (Timestamp)
		{
			const err = this.expiration_timestamp.Validate();
			if (err) {
				return sprintf('field expiration_timestamp is invalid : %s', err);
			}
	
		}
	

		// ID (string)
		{
			if (this.id.length > (2 << 8) - 1) {
				return sprintf('varchar field id too long %d/%d', this.id.length, (2 << 8) - 1);
			}
		}
	

		// MembershipClass (string)
		{
			if (this.membership_class.length > (2 << 8) - 1) {
				return sprintf('varchar field membership_class too long %d/%d', this.membership_class.length, (2 << 8) - 1);
			}
		}
	

		// RoleType (string)
		{
			if (this.role_type.length > (2 << 8) - 1) {
				return sprintf('varchar field role_type too long %d/%d', this.role_type.length, (2 << 8) - 1);
			}
		}
	

		// MembershipType (string)
		{
			if (this.membership_type.length > (2 << 8) - 1) {
				return sprintf('varchar field membership_type too long %d/%d', this.membership_type.length, (2 << 8) - 1);
			}
		}
	

		// Description (string)
		{
			if (this.description.length > (2 << 16) - 1) {
				return sprintf('varchar field description too long %d/%d', this.description.length, (2 << 16) - 1);
			}
		}
	
	}

	toString(): string {
		const vals = [];
	
		vals.push(sprintf('version:%v', this.version));
		vals.push(sprintf('age_restriction:%#+v', this.age_restriction));
		vals.push(sprintf('valid_from:%#+v', this.valid_from));
		vals.push(sprintf('expiration_timestamp:%#+v', this.expiration_timestamp));
		vals.push(sprintf('id:%#+v', this.id));
		vals.push(sprintf('membership_class:%#+v', this.membership_class));
		vals.push(sprintf('role_type:%#+v', this.role_type));
		vals.push(sprintf('membership_type:%#+v', this.membership_type));
		vals.push(sprintf('description:%#+v', this.description));

		return sprintf('{%s}', vals.join(' '));
	}
}


// ShareCommon asset type.
export class ShareCommon implements AssetPayload {

// The version number that this asset payload adheres to.
	version; // uint8 `json:"version,omitempty"`
// A period of time where the asset is unable to be transferred.  After the transfer lockout period, the assets can be transferred.
	transfer_lockout; // Timestamp `json:"transfer_lockout,omitempty"`
// Ticker symbol assigned by exchanges to represent the asset.
	ticker; // string `json:"ticker,omitempty"`
// International Securities Identification Number
	isin; // string `json:"isin,omitempty"`
// 
	description; // string `json:"description,omitempty"`


	// Type returns the type identifer for this message.
	Type(): string {
		return CodeShareCommon;
	}

	// Len returns the byte size of this message.
	Len(): number {
		return AssetTypeLen;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	Read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// Serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
	

		// Version (uint8)
		{
			write(buf, this.version, 'uint8');
		}
	

		// TransferLockout (Timestamp)
		{
			write(buf, this.transfer_lockout, 'Timestamp');
		}
	

		// Ticker (string)
		{
			WriteFixedChar(buf, this.ticker, 5);
		}
	

		// ISIN (string)
		{
			WriteFixedChar(buf, this.isin, 12);
		}
	

		// Description (string)
		{
			WriteVarChar(buf, this.description, 16);
		}
	
		return buf.buf;
	}

	// Write populates the fields in ShareCommon from the byte slice
	Write(b: Buffer): number {
		const buf = new _.Reader(b);

		// Version (uint8)
		{
			this.version = read(buf, 'uint8');
	
		}
	

		// TransferLockout (Timestamp)
		{
			this.transfer_lockout.Write(buf);
	
		}
	

		// Ticker (string)
		{
			this.ticker = ReadFixedChar(buf, 5);
		}
	

		// ISIN (string)
		{
			this.isin = ReadFixedChar(buf, 12);
		}
	

		// Description (string)
		{
			this.description = ReadVarChar(buf, 16);
		}
	
		return b.length - buf.length;
	}

	Validate(): string {

		// Version (uint8)
		{
		}
	

		// TransferLockout (Timestamp)
		{
			const err = this.transfer_lockout.Validate();
			if (err) {
				return sprintf('field transfer_lockout is invalid : %s', err);
			}
	
		}
	

		// Ticker (string)
		{
			if (this.ticker.length > 5) {
				return sprintf('fixedchar field ticker too long %d/%d', this.ticker.length, 5);
			}
		}
	

		// ISIN (string)
		{
			if (this.isin.length > 12) {
				return sprintf('fixedchar field isin too long %d/%d', this.isin.length, 12);
			}
		}
	

		// Description (string)
		{
			if (this.description.length > (2 << 16) - 1) {
				return sprintf('varchar field description too long %d/%d', this.description.length, (2 << 16) - 1);
			}
		}
	
	}

	toString(): string {
		const vals = [];
	
		vals.push(sprintf('version:%v', this.version));
		vals.push(sprintf('transfer_lockout:%#+v', this.transfer_lockout));
		vals.push(sprintf('ticker:%#+v', this.ticker));
		vals.push(sprintf('isin:%#+v', this.isin));
		vals.push(sprintf('description:%#+v', this.description));

		return sprintf('{%s}', vals.join(' '));
	}
}


// TicketAdmission asset type.
export class TicketAdmission implements AssetPayload {

// The version number that this asset payload adheres to.
	version; // uint8 `json:"version,omitempty"`
// Age restriction is used to specify required ages for asset ownership.
	age_restriction; // AgeRestriction `json:"age_restriction,omitempty"`
// 
	admission_type; // string `json:"admission_type,omitempty"`
// 
	venue; // string `json:"venue,omitempty"`
// 
	class; // string `json:"class,omitempty"`
// 
	area; // string `json:"area,omitempty"`
// 
	seat; // string `json:"seat,omitempty"`
// 
	start_time_date; // Timestamp `json:"start_time_date,omitempty"`
// 
	valid_from; // Timestamp `json:"valid_from,omitempty"`
// 
	expiration_timestamp; // Timestamp `json:"expiration_timestamp,omitempty"`
// 
	description; // string `json:"description,omitempty"`


	// Type returns the type identifer for this message.
	Type(): string {
		return CodeTicketAdmission;
	}

	// Len returns the byte size of this message.
	Len(): number {
		return AssetTypeLen;
	}

	// Read implements the io.Reader interface, writing the receiver to the
	// []byte.
	Read(b: Buffer): number {
		const data = this.Serialize();

		b.fill(data);

		return data.length;
	}

	// Serialize returns the full OP_RETURN payload bytes.
	Serialize(): Buffer {
		const buf = new _.Writer();
	

		// Version (uint8)
		{
			write(buf, this.version, 'uint8');
		}
	

		// AgeRestriction (AgeRestriction)
		{
			write(buf, this.age_restriction, 'AgeRestriction');
		}
	

		// AdmissionType (string)
		{
			WriteFixedChar(buf, this.admission_type, 3);
		}
	

		// Venue (string)
		{
			WriteVarChar(buf, this.venue, 8);
		}
	

		// Class (string)
		{
			WriteVarChar(buf, this.class, 8);
		}
	

		// Area (string)
		{
			WriteVarChar(buf, this.area, 8);
		}
	

		// Seat (string)
		{
			WriteVarChar(buf, this.seat, 8);
		}
	

		// StartTimeDate (Timestamp)
		{
			write(buf, this.start_time_date, 'Timestamp');
		}
	

		// ValidFrom (Timestamp)
		{
			write(buf, this.valid_from, 'Timestamp');
		}
	

		// ExpirationTimestamp (Timestamp)
		{
			write(buf, this.expiration_timestamp, 'Timestamp');
		}
	

		// Description (string)
		{
			WriteVarChar(buf, this.description, 16);
		}
	
		return buf.buf;
	}

	// Write populates the fields in TicketAdmission from the byte slice
	Write(b: Buffer): number {
		const buf = new _.Reader(b);

		// Version (uint8)
		{
			this.version = read(buf, 'uint8');
	
		}
	

		// AgeRestriction (AgeRestriction)
		{
			this.age_restriction.Write(buf);
	
		}
	

		// AdmissionType (string)
		{
			this.admission_type = ReadFixedChar(buf, 3);
		}
	

		// Venue (string)
		{
			this.venue = ReadVarChar(buf, 8);
		}
	

		// Class (string)
		{
			this.class = ReadVarChar(buf, 8);
		}
	

		// Area (string)
		{
			this.area = ReadVarChar(buf, 8);
		}
	

		// Seat (string)
		{
			this.seat = ReadVarChar(buf, 8);
		}
	

		// StartTimeDate (Timestamp)
		{
			this.start_time_date.Write(buf);
	
		}
	

		// ValidFrom (Timestamp)
		{
			this.valid_from.Write(buf);
	
		}
	

		// ExpirationTimestamp (Timestamp)
		{
			this.expiration_timestamp.Write(buf);
	
		}
	

		// Description (string)
		{
			this.description = ReadVarChar(buf, 16);
		}
	
		return b.length - buf.length;
	}

	Validate(): string {

		// Version (uint8)
		{
		}
	

		// AgeRestriction (AgeRestriction)
		{
			const err = this.age_restriction.Validate();
			if (err) {
				return sprintf('field age_restriction is invalid : %s', err);
			}
	
		}
	

		// AdmissionType (string)
		{
			if (this.admission_type.length > 3) {
				return sprintf('fixedchar field admission_type too long %d/%d', this.admission_type.length, 3);
			}
		}
	

		// Venue (string)
		{
			if (this.venue.length > (2 << 8) - 1) {
				return sprintf('varchar field venue too long %d/%d', this.venue.length, (2 << 8) - 1);
			}
		}
	

		// Class (string)
		{
			if (this.class.length > (2 << 8) - 1) {
				return sprintf('varchar field class too long %d/%d', this.class.length, (2 << 8) - 1);
			}
		}
	

		// Area (string)
		{
			if (this.area.length > (2 << 8) - 1) {
				return sprintf('varchar field area too long %d/%d', this.area.length, (2 << 8) - 1);
			}
		}
	

		// Seat (string)
		{
			if (this.seat.length > (2 << 8) - 1) {
				return sprintf('varchar field seat too long %d/%d', this.seat.length, (2 << 8) - 1);
			}
		}
	

		// StartTimeDate (Timestamp)
		{
			const err = this.start_time_date.Validate();
			if (err) {
				return sprintf('field start_time_date is invalid : %s', err);
			}
	
		}
	

		// ValidFrom (Timestamp)
		{
			const err = this.valid_from.Validate();
			if (err) {
				return sprintf('field valid_from is invalid : %s', err);
			}
	
		}
	

		// ExpirationTimestamp (Timestamp)
		{
			const err = this.expiration_timestamp.Validate();
			if (err) {
				return sprintf('field expiration_timestamp is invalid : %s', err);
			}
	
		}
	

		// Description (string)
		{
			if (this.description.length > (2 << 16) - 1) {
				return sprintf('varchar field description too long %d/%d', this.description.length, (2 << 16) - 1);
			}
		}
	
	}

	toString(): string {
		const vals = [];
	
		vals.push(sprintf('version:%v', this.version));
		vals.push(sprintf('age_restriction:%#+v', this.age_restriction));
		vals.push(sprintf('admission_type:%#+v', this.admission_type));
		vals.push(sprintf('venue:%#+v', this.venue));
		vals.push(sprintf('class:%#+v', this.class));
		vals.push(sprintf('area:%#+v', this.area));
		vals.push(sprintf('seat:%#+v', this.seat));
		vals.push(sprintf('start_time_date:%#+v', this.start_time_date));
		vals.push(sprintf('valid_from:%#+v', this.valid_from));
		vals.push(sprintf('expiration_timestamp:%#+v', this.expiration_timestamp));
		vals.push(sprintf('description:%#+v', this.description));

		return sprintf('{%s}', vals.join(' '));
	}
}

