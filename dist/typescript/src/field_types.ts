import {sprintf} from 'sprintf-js';
import _ from '@keyring/util';
import {write, read, ReadVarChar, ReadVariableSize, ReadVarBin, ReadFixedChar,
	WriteVarChar, WriteVariableSize, WriteFixedChar, WriteVarBin} from './bytes';
import { GetRoleType, GetEntityType, } from './resources';


// Administrator Administrator is used to refer to a Administration role in
// an Entity.
export class Administrator {
// Chairman, Director, Managing Partner, etc.. Found in 'Roles' in
	// Specification/Resources
	type; // uint8 `json:"type,omitempty"`
// Length 0-255 bytes. 0 is valid.
	name; // string `json:"name,omitempty"`
// Serialize returns the byte representation of the message.
	Serialize(): Buffer {
		const buf = new _.Writer();

		// Type (uint8)
		{
			write(buf, this.type, 'uint8');
		}

		// Name (string)
		{
			WriteVarChar(buf, this.name, 8);
		}
		return buf.buf;
	}

	Write(buf: _.Reader) {

		// Type (uint8)
		{
			this.type = read(buf, 'uint8');
		}

		// Name (string)
		{
			this.name = ReadVarChar(buf, 8);
		}
	}

	Validate() {

		// type (uint8)
		{
			if (!GetRoleType(this.type)) {
				return sprintf('Invalid role value : %d', this.type);
			}
		}

		// name (string)
		{
			if (this.name.length > (2 << 8) - 1) {
				return sprintf('varchar field name too long %d/%d', this.name.length, (2 << 8) - 1);
			}
		}
		return;
	}

	Equal(other: Administrator): boolean {

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


// AgeRestriction Age restriction is used to specify required ages for
// asset ownership.
export class AgeRestriction {
// The lowest age valid to own asset. Zero for no restriction.
	lower; // uint8 `json:"lower,omitempty"`
// The highest age valid to own asset. Zero for no restriction.
	upper; // uint8 `json:"upper,omitempty"`
// Serialize returns the byte representation of the message.
	Serialize(): Buffer {
		const buf = new _.Writer();

		// Lower (uint8)
		{
			write(buf, this.lower, 'uint8');
		}

		// Upper (uint8)
		{
			write(buf, this.upper, 'uint8');
		}
		return buf.buf;
	}

	Write(buf: _.Reader) {

		// Lower (uint8)
		{
			this.lower = read(buf, 'uint8');
		}

		// Upper (uint8)
		{
			this.upper = read(buf, 'uint8');
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

	Equal(other: AgeRestriction): boolean {

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


// Amendment An Amendment is used to describe the modification of a single
// field in a Contract or Asset, as defined in the ContractFormation and
// AssetCreation messages.
export class Amendment {
// Index of the field to be amended.
	field_index; // uint8 `json:"field_index,omitempty"`
// Specifies the element of the complex array type to be amended. This
	// only applies to array types, and has no meaning for a simple type such
	// as uint64, string, byte or byte[]. Specifying a value > 0 for a simple
	// type will result in a Rejection.
	element; // uint16 `json:"element,omitempty"`
// Index of the subfield to be amended. This only applies to specific
	// fields containing complex types with subfields. This is used to specify
	// which field of the object the amendment applies to.
	subfield_index; // uint8 `json:"subfield_index,omitempty"`
// Specifies the element of the complex array type to be amended. This
	// only applies to array types, and has no meaning for a simple type such
	// as uint64, string, byte or byte[]. Specifying a value > 0 for a simple
	// type will result in a Rejection.
	subfield_element; // uint16 `json:"subfield_element,omitempty"`
// 0 = Modify. 1 = Add an element to the data to the array of elements. 2
	// = Delete the element listed in the Element field. The Add and Delete
	// operations only apply to a particilar element of a complex array type.
	// For example, it could be used to remove a particular VotingSystem from
	// a Contract.
	operation; // uint8 `json:"operation,omitempty"`
// New data for the amended subfield. Data type depends on the the type of
	// the field being amended. The value should be serialize as defined by
	// the protocol.
	data; // []byte `json:"data,omitempty"`
// Serialize returns the byte representation of the message.
	Serialize(): Buffer {
		const buf = new _.Writer();

		// FieldIndex (uint8)
		{
			write(buf, this.field_index, 'uint8');
		}

		// Element (uint16)
		{
			write(buf, this.element, 'uint16');
		}

		// SubfieldIndex (uint8)
		{
			write(buf, this.subfield_index, 'uint8');
		}

		// SubfieldElement (uint16)
		{
			write(buf, this.subfield_element, 'uint16');
		}

		// Operation (uint8)
		{
			write(buf, this.operation, 'uint8');
		}

		// Data ([]byte)
		{
			WriteVarBin(buf, this.data, 32);
		}
		return buf.buf;
	}

	Write(buf: _.Reader) {

		// FieldIndex (uint8)
		{
			this.field_index = read(buf, 'uint8');
		}

		// Element (uint16)
		{
			this.element = read(buf, 'uint16');
		}

		// SubfieldIndex (uint8)
		{
			this.subfield_index = read(buf, 'uint8');
		}

		// SubfieldElement (uint16)
		{
			this.subfield_element = read(buf, 'uint16');
		}

		// Operation (uint8)
		{
			this.operation = read(buf, 'uint8');
		}

		// Data ([]byte)
		{
			this.data = ReadVarBin(buf, 32);
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
				return sprintf('varbin field data too long %d/%d', this.data.length, (2 << 32) - 1);
			}
		}
		return;
	}

	Equal(other: Amendment): boolean {

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


// AssetReceiver An AssetReceiver is a quantity, address, and oracle
// signature. The quantity could be used to describe a number of tokens, or
// a value. The address is where the asset will be sent.
export class AssetReceiver {
// The address receiving the tokens
	address; // PublicKeyHash `json:"address,omitempty"`
// Number of tokens to be received by address at Output X
	quantity; // uint64 `json:"quantity,omitempty"`
// 0 = No Oracle-signed Message (OracleConfirmationSig skipped in
	// serialization), 1 = ECDSA+secp256k1. If the contract for the asset
	// being received has oracles, then a signature is required from one of
	// them.
	oracle_sig_algorithm; // uint8 `json:"oracle_sig_algorithm,omitempty"`
// Specifies the index into the list of oracles in the contract offer that
	// was used to authorize this transfer.
	oracle_index; // uint8 `json:"oracle_index,omitempty"`
// Length 0-255 bytes. If restricted to a oracle (whitelist) or has
	// transfer restrictions (age, location, investor status): ECDSA+secp256k1
	// (or the like) signed message provided by an approved/trusted oracle
	// through an API signature of the defined message. If no transfer
	// restrictions(trade restriction/age restriction fields in the Asset Type
	// payload. or restricted to a whitelist by the Contract Auth Flags, it is
	// a NULL field.
	oracle_confirmation_sig; // []byte `json:"oracle_confirmation_sig,omitempty"`
// The block height of the block hash used in the oracle signature.
	oracle_sig_block_height; // uint32 `json:"oracle_sig_block_height,omitempty"`
// Serialize returns the byte representation of the message.
	Serialize(): Buffer {
		const buf = new _.Writer();

		// Address (PublicKeyHash)
		{
			buf.write(this.address.Serialize());
		}

		// Quantity (uint64)
		{
			write(buf, this.quantity, 'uint64');
		}

		// OracleSigAlgorithm (uint8)
		{
			write(buf, this.oracle_sig_algorithm, 'uint8');
		}

		// OracleIndex (uint8)
		if ( this.oracle_sig_algorithm === 1) {
			write(buf, this.oracle_index, 'uint8');
		}

		// OracleConfirmationSig ([]byte)
		if ( this.oracle_sig_algorithm === 1) {
			WriteVarBin(buf, this.oracle_confirmation_sig, 8);
		}

		// OracleSigBlockHeight (uint32)
		if ( this.oracle_sig_algorithm === 1) {
			write(buf, this.oracle_sig_block_height, 'uint32');
		}
		return buf.buf;
	}

	Write(buf: _.Reader) {

		// Address (PublicKeyHash)
		{
			this.address.Write(buf);
		}

		// Quantity (uint64)
		{
			this.quantity = read(buf, 'uint64');
		}

		// OracleSigAlgorithm (uint8)
		{
			this.oracle_sig_algorithm = read(buf, 'uint8');
		}

		// OracleIndex (uint8)
		if ( this.oracle_sig_algorithm === 1) {
			this.oracle_index = read(buf, 'uint8');
		}

		// OracleConfirmationSig ([]byte)
		if ( this.oracle_sig_algorithm === 1) {
			this.oracle_confirmation_sig = ReadVarBin(buf, 8);
		}

		// OracleSigBlockHeight (uint32)
		if ( this.oracle_sig_algorithm === 1) {
			this.oracle_sig_block_height = read(buf, 'uint32');
		}
	}

	Validate() {

		// address (PublicKeyHash)
		{
			const err = this.address.Validate();
			if (err) {
				return sprintf('field address is invalid : %s', err);
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
				return sprintf('varbin field oracle_confirmation_sig too long %d/%d', this.oracle_confirmation_sig.length, (2 << 8) - 1);
			}
		}

		// oracle_sig_block_height (uint32)
		{
		}
		return;
	}

	Equal(other: AssetReceiver): boolean {

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


// AssetSettlement AssetSettlement is the data required to settle an asset
// transfer.
export class AssetSettlement {
// Index of input containing the contract's address for this offset
	contract_index; // uint16 `json:"contract_index,omitempty"`
// Three letter character that specifies the asset type. Example: COU
	asset_type; // string `json:"asset_type,omitempty"`
// A unique code that is used to identify the asset. It is generated by
	// hashing the contract public key hash and the asset index.
	// SHA256(contract PKH + asset index)
	asset_code; // AssetCode `json:"asset_code,omitempty"`
// Each element contains the resulting token balance of Asset X for the
	// output Address, which is referred to by the index.
	settlements; // []QuantityIndex `json:"settlements,omitempty"`
// Serialize returns the byte representation of the message.
	Serialize(): Buffer {
		const buf = new _.Writer();

		// ContractIndex (uint16)
		{
			write(buf, this.contract_index, 'uint16');
		}

		// AssetType (string)
		{
			WriteFixedChar(buf, this.asset_type, 3);
		}

		// AssetCode (AssetCode)
		{
			buf.write(this.asset_code.Serialize());
		}

		// Settlements ([]QuantityIndex)
		{
			WriteVariableSize(buf, this.settlements.length, 8, 8);
			this.settlements.forEach(value => {
				buf.write(value.Serialize());
			});
		}
		return buf.buf;
	}

	Write(buf: _.Reader) {

		// ContractIndex (uint16)
		{
			this.contract_index = read(buf, 'uint16');
		}

		// AssetType (string)
		{
			this.asset_type = ReadFixedChar(buf, 3);
		}

		// AssetCode (AssetCode)
		{
			this.asset_code.Write(buf);
		}

		// Settlements ([]QuantityIndex)
		{
			const size = ReadVariableSize(buf, 8, 8);
			this.settlements = [Array(size)]
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
				return sprintf('fixedchar field asset_type too long %d/%d', this.asset_type.length, 3);
			}
		}

		// asset_code (AssetCode)
		{
			const err = this.asset_code.Validate();
			if (err) {
				return sprintf('field asset_code is invalid : %s', err);
			}
		}

		// settlements ([]QuantityIndex)
		{
			if (this.settlements.length > (2 << 8) - 1) {
				return sprintf('list field settlements has too many items %d/%d', this.settlements.length, (2 << 8) - 1);
			}

			this.settlements.forEach((value, i) => {
				const err = value.Validate();
				if (err) {
					return sprintf('list field settlements[%d] is invalid : %s', i, err);
				}
			});
		}
		return;
	}

	Equal(other: AssetSettlement): boolean {

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
		if (this.settlements.length !== 0 || other.settlements.length !== 0 ) {
			if (this.settlements.length !== other.settlements.length) {
				return false;
			}

			if (!this.settlements.every((value, i) => {
				if (!value.Equal(other.settlements[i])) {
					return false;
				}
				return true;
			})) return false;
		}
		return true;
	}

}


// AssetTransfer AssetTransfer is the data required to transfer an asset.
export class AssetTransfer {
// Index of output containing the contract's address for this offset
	contract_index; // uint16 `json:"contract_index,omitempty"`
// Three letter character that specifies the asset type. Example: COU
	asset_type; // string `json:"asset_type,omitempty"`
// A unique code that is used to identify the asset. It is generated by
	// hashing the contract public key hash and the asset index.
	// SHA256(contract PKH + asset index)
	asset_code; // AssetCode `json:"asset_code,omitempty"`
// Each element has the value of tokens to be spent from the input
	// address, which is referred to by the index.
	asset_senders; // []QuantityIndex `json:"asset_senders,omitempty"`
// Each element has the value of tokens to be received, the address, and
	// an oracle signature if required.
	asset_receivers; // []AssetReceiver `json:"asset_receivers,omitempty"`
// Serialize returns the byte representation of the message.
	Serialize(): Buffer {
		const buf = new _.Writer();

		// ContractIndex (uint16)
		{
			write(buf, this.contract_index, 'uint16');
		}

		// AssetType (string)
		{
			WriteFixedChar(buf, this.asset_type, 3);
		}

		// AssetCode (AssetCode)
		{
			buf.write(this.asset_code.Serialize());
		}

		// AssetSenders ([]QuantityIndex)
		{
			WriteVariableSize(buf, this.asset_senders.length, 8, 8);
			this.asset_senders.forEach(value => {
				buf.write(value.Serialize());
			});
		}

		// AssetReceivers ([]AssetReceiver)
		{
			WriteVariableSize(buf, this.asset_receivers.length, 8, 8);
			this.asset_receivers.forEach(value => {
				buf.write(value.Serialize());
			});
		}
		return buf.buf;
	}

	Write(buf: _.Reader) {

		// ContractIndex (uint16)
		{
			this.contract_index = read(buf, 'uint16');
		}

		// AssetType (string)
		{
			this.asset_type = ReadFixedChar(buf, 3);
		}

		// AssetCode (AssetCode)
		{
			this.asset_code.Write(buf);
		}

		// AssetSenders ([]QuantityIndex)
		{
			const size = ReadVariableSize(buf, 8, 8);
			this.asset_senders = [Array(size)]
				.map(() => {
					const newValue = new QuantityIndex();
					newValue.Write(buf);
					return newValue;
				});
		}

		// AssetReceivers ([]AssetReceiver)
		{
			const size = ReadVariableSize(buf, 8, 8);
			this.asset_receivers = [Array(size)]
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
				return sprintf('fixedchar field asset_type too long %d/%d', this.asset_type.length, 3);
			}
		}

		// asset_code (AssetCode)
		{
			const err = this.asset_code.Validate();
			if (err) {
				return sprintf('field asset_code is invalid : %s', err);
			}
		}

		// asset_senders ([]QuantityIndex)
		{
			if (this.asset_senders.length > (2 << 8) - 1) {
				return sprintf('list field asset_senders has too many items %d/%d', this.asset_senders.length, (2 << 8) - 1);
			}

			this.asset_senders.forEach((value, i) => {
				const err = value.Validate();
				if (err) {
					return sprintf('list field asset_senders[%d] is invalid : %s', i, err);
				}
			});
		}

		// asset_receivers ([]AssetReceiver)
		{
			if (this.asset_receivers.length > (2 << 8) - 1) {
				return sprintf('list field asset_receivers has too many items %d/%d', this.asset_receivers.length, (2 << 8) - 1);
			}

			this.asset_receivers.forEach((value, i) => {
				const err = value.Validate();
				if (err) {
					return sprintf('list field asset_receivers[%d] is invalid : %s', i, err);
				}
			});
		}
		return;
	}

	Equal(other: AssetTransfer): boolean {

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
		if (this.asset_senders.length !== 0 || other.asset_senders.length !== 0 ) {
			if (this.asset_senders.length !== other.asset_senders.length) {
				return false;
			}

			if (!this.asset_senders.every((value, i) => {
				if (!value.Equal(other.asset_senders[i])) {
					return false;
				}
				return true;
			})) return false;
		}

		// AssetReceivers ([]AssetReceiver)
		if (this.asset_receivers.length !== 0 || other.asset_receivers.length !== 0 ) {
			if (this.asset_receivers.length !== other.asset_receivers.length) {
				return false;
			}

			if (!this.asset_receivers.every((value, i) => {
				if (!value.Equal(other.asset_receivers[i])) {
					return false;
				}
				return true;
			})) return false;
		}
		return true;
	}

}


// Document A file containing data.
export class Document {
// Full name, including file extension, of the file. Length 0-255 bytes. 0
	// is valid.
	name; // string `json:"name,omitempty"`
// MIME type of the file. Length 0-255 bytes. 0 is valid.
	type; // string `json:"type,omitempty"`
// Compression/encryption algorithm used on file contents. Compression is
	// encoded before encryption, and the reverse for decoding. 0 is no
	// compression or encryption, 1 is AES256 encryption with no compression,
	// 2 is LZMA2 with no encryption, 3 is LZMA2 with AES256 encryption, 4 is
	// Deflate with no encryption, 5 is LZMA2 with AES256 encryption.
	algorithm; // uint8 `json:"algorithm,omitempty"`
// The contents of the file.
	contents; // []byte `json:"contents,omitempty"`
// Serialize returns the byte representation of the message.
	Serialize(): Buffer {
		const buf = new _.Writer();

		// Name (string)
		{
			WriteVarChar(buf, this.name, 8);
		}

		// Type (string)
		{
			WriteVarChar(buf, this.type, 8);
		}

		// Algorithm (uint8)
		{
			write(buf, this.algorithm, 'uint8');
		}

		// Contents ([]byte)
		{
			WriteVarBin(buf, this.contents, 32);
		}
		return buf.buf;
	}

	Write(buf: _.Reader) {

		// Name (string)
		{
			this.name = ReadVarChar(buf, 8);
		}

		// Type (string)
		{
			this.type = ReadVarChar(buf, 8);
		}

		// Algorithm (uint8)
		{
			this.algorithm = read(buf, 'uint8');
		}

		// Contents ([]byte)
		{
			this.contents = ReadVarBin(buf, 32);
		}
	}

	Validate() {

		// name (string)
		{
			if (this.name.length > (2 << 8) - 1) {
				return sprintf('varchar field name too long %d/%d', this.name.length, (2 << 8) - 1);
			}
		}

		// type (string)
		{
			if (this.type.length > (2 << 8) - 1) {
				return sprintf('varchar field type too long %d/%d', this.type.length, (2 << 8) - 1);
			}
		}

		// algorithm (uint8)
		{
		}

		// contents ([]byte)
		{
			if (this.contents.length > (2 << 32) - 1) {
				return sprintf('varbin field contents too long %d/%d', this.contents.length, (2 << 32) - 1);
			}
		}
		return;
	}

	Equal(other: Document): boolean {

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


// Entity Entity represents the details of a legal Entity, such as a public
// or private company, government agency, or and individual.
export class Entity {
// Length 1-255 bytes (0 is not valid). Issuing entity (company,
	// organization, individual). Can be any unique identifying string,
	// including human readable names for branding/vanity purposes.
	name; // string `json:"name,omitempty"`
// The type of entity. (i.e Public Company, Individual)
	// (Specification/Resources).
	type; // byte `json:"type,omitempty"`
// Null is valid. A Legal Entity Identifier (or LEI) is an international
	// identifier made up of a 20-character identifier that identifies
	// distinct legal entities that engage in financial transactions. It is
	// defined by ISO 17442.[1] Natural persons are not required to have an
	// LEI; they’re eligible to have one issued, however, but only if they
	// act in an independent business capacity.[2] The LEI is a global
	// standard, designed to be non-proprietary data that is freely accessible
	// to all.[3] As of December 2018, over 1,300,000 legal entities from more
	// than 200 countries have now been issued with LEIs.
	lei; // string `json:"lei,omitempty"`
// Registered/Physical/mailing address(HQ). Y-1/N-0, N means there is no
	// issuer address.
	address_included; // bool `json:"address_included,omitempty"`
// Issuer/Entity/Contracting Party X Address Details (eg. HQ)
	unit_number; // string `json:"unit_number,omitempty"`

	building_number; // string `json:"building_number,omitempty"`

	street; // string `json:"street,omitempty"`

	suburb_city; // string `json:"suburb_city,omitempty"`

	territory_state_province_code; // string `json:"territory_state_province_code,omitempty"`

	country_code; // string `json:"country_code,omitempty"`

	postal_zip_code; // string `json:"postal_zip_code,omitempty"`
// Length 0-255 bytes. Address for text-based communication: eg. email
	// address, Bitcoin address
	email_address; // string `json:"email_address,omitempty"`
// Length 0-50 bytes. 0 is valid. Phone Number for Entity.
	phone_number; // string `json:"phone_number,omitempty"`
// A list of people that are in Administrative Roles for the Entity. eg.
	// Chair, Director, Managing Partner, etc.
	administration; // []Administrator `json:"administration,omitempty"`
// A list of people in Management Roles for the Entity. e.g CEO, COO, CTO,
	// CFO, Secretary, Executive, etc.
	management; // []Manager `json:"management,omitempty"`
// Serialize returns the byte representation of the message.
	Serialize(): Buffer {
		const buf = new _.Writer();

		// Name (string)
		{
			WriteVarChar(buf, this.name, 8);
		}

		// Type (byte)
		{
			write(buf, this.type, 'byte');
		}

		// LEI (string)
		{
			WriteFixedChar(buf, this.lei, 20);
		}

		// AddressIncluded (bool)
		{
			write(buf, this.address_included, 'bool');
		}

		// UnitNumber (string)
		if (this.address_included) {
			WriteVarChar(buf, this.unit_number, 8);
		}

		// BuildingNumber (string)
		if (this.address_included) {
			WriteVarChar(buf, this.building_number, 8);
		}

		// Street (string)
		if (this.address_included) {
			WriteVarChar(buf, this.street, 16);
		}

		// SuburbCity (string)
		if (this.address_included) {
			WriteVarChar(buf, this.suburb_city, 8);
		}

		// TerritoryStateProvinceCode (string)
		if (this.address_included) {
			WriteFixedChar(buf, this.territory_state_province_code, 5);
		}

		// CountryCode (string)
		if (this.address_included) {
			WriteFixedChar(buf, this.country_code, 3);
		}

		// PostalZIPCode (string)
		if (this.address_included) {
			WriteFixedChar(buf, this.postal_zip_code, 12);
		}

		// EmailAddress (string)
		{
			WriteVarChar(buf, this.email_address, 8);
		}

		// PhoneNumber (string)
		{
			WriteVarChar(buf, this.phone_number, 8);
		}

		// Administration ([]Administrator)
		{
			WriteVariableSize(buf, this.administration.length, 8, 8);
			this.administration.forEach(value => {
				buf.write(value.Serialize());
			});
		}

		// Management ([]Manager)
		{
			WriteVariableSize(buf, this.management.length, 8, 8);
			this.management.forEach(value => {
				buf.write(value.Serialize());
			});
		}
		return buf.buf;
	}

	Write(buf: _.Reader) {

		// Name (string)
		{
			this.name = ReadVarChar(buf, 8);
		}

		// Type (byte)
		{
			this.type = read(buf, 'byte');
		}

		// LEI (string)
		{
			this.lei = ReadFixedChar(buf, 20);
		}

		// AddressIncluded (bool)
		{
			this.address_included = read(buf, 'bool');
		}

		// UnitNumber (string)
		if (this.address_included) {
			this.unit_number = ReadVarChar(buf, 8);
		}

		// BuildingNumber (string)
		if (this.address_included) {
			this.building_number = ReadVarChar(buf, 8);
		}

		// Street (string)
		if (this.address_included) {
			this.street = ReadVarChar(buf, 16);
		}

		// SuburbCity (string)
		if (this.address_included) {
			this.suburb_city = ReadVarChar(buf, 8);
		}

		// TerritoryStateProvinceCode (string)
		if (this.address_included) {
			this.territory_state_province_code = ReadFixedChar(buf, 5);
		}

		// CountryCode (string)
		if (this.address_included) {
			this.country_code = ReadFixedChar(buf, 3);
		}

		// PostalZIPCode (string)
		if (this.address_included) {
			this.postal_zip_code = ReadFixedChar(buf, 12);
		}

		// EmailAddress (string)
		{
			this.email_address = ReadVarChar(buf, 8);
		}

		// PhoneNumber (string)
		{
			this.phone_number = ReadVarChar(buf, 8);
		}

		// Administration ([]Administrator)
		{
			const size = ReadVariableSize(buf, 8, 8);
			this.administration = [Array(size)]
				.map(() => {
					const newValue = new Administrator();
					newValue.Write(buf);
					return newValue;
				});
		}

		// Management ([]Manager)
		{
			const size = ReadVariableSize(buf, 8, 8);
			this.management = [Array(size)]
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
				return sprintf('varchar field name too long %d/%d', this.name.length, (2 << 8) - 1);
			}
		}

		// type (byte)
		{
			if (!GetEntityType(this.type)) {
				return sprintf('Invalid entity type value : %c', this.type);
			}
		}

		// lei (string)
		{
			if (this.lei.length > 20) {
				return sprintf('fixedchar field lei too long %d/%d', this.lei.length, 20);
			}
		}

		// address_included (bool)
		{
		}

		// unit_number (string)
		if (this.address_included) {
			if (this.unit_number.length > (2 << 8) - 1) {
				return sprintf('varchar field unit_number too long %d/%d', this.unit_number.length, (2 << 8) - 1);
			}
		}

		// building_number (string)
		if (this.address_included) {
			if (this.building_number.length > (2 << 8) - 1) {
				return sprintf('varchar field building_number too long %d/%d', this.building_number.length, (2 << 8) - 1);
			}
		}

		// street (string)
		if (this.address_included) {
			if (this.street.length > (2 << 16) - 1) {
				return sprintf('varchar field street too long %d/%d', this.street.length, (2 << 16) - 1);
			}
		}

		// suburb_city (string)
		if (this.address_included) {
			if (this.suburb_city.length > (2 << 8) - 1) {
				return sprintf('varchar field suburb_city too long %d/%d', this.suburb_city.length, (2 << 8) - 1);
			}
		}

		// territory_state_province_code (string)
		if (this.address_included) {
			if (this.territory_state_province_code.length > 5) {
				return sprintf('fixedchar field territory_state_province_code too long %d/%d', this.territory_state_province_code.length, 5);
			}
		}

		// country_code (string)
		if (this.address_included) {
			if (this.country_code.length > 3) {
				return sprintf('fixedchar field country_code too long %d/%d', this.country_code.length, 3);
			}
		}

		// postal_zip_code (string)
		if (this.address_included) {
			if (this.postal_zip_code.length > 12) {
				return sprintf('fixedchar field postal_zip_code too long %d/%d', this.postal_zip_code.length, 12);
			}
		}

		// email_address (string)
		{
			if (this.email_address.length > (2 << 8) - 1) {
				return sprintf('varchar field email_address too long %d/%d', this.email_address.length, (2 << 8) - 1);
			}
		}

		// phone_number (string)
		{
			if (this.phone_number.length > (2 << 8) - 1) {
				return sprintf('varchar field phone_number too long %d/%d', this.phone_number.length, (2 << 8) - 1);
			}
		}

		// administration ([]Administrator)
		{
			if (this.administration.length > (2 << 8) - 1) {
				return sprintf('list field administration has too many items %d/%d', this.administration.length, (2 << 8) - 1);
			}

			this.administration.forEach((value, i) => {
				const err = value.Validate();
				if (err) {
					return sprintf('list field administration[%d] is invalid : %s', i, err);
				}
			});
		}

		// management ([]Manager)
		{
			if (this.management.length > (2 << 8) - 1) {
				return sprintf('list field management has too many items %d/%d', this.management.length, (2 << 8) - 1);
			}

			this.management.forEach((value, i) => {
				const err = value.Validate();
				if (err) {
					return sprintf('list field management[%d] is invalid : %s', i, err);
				}
			});
		}
		return;
	}

	Equal(other: Entity): boolean {

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
		if (this.administration.length !== 0 || other.administration.length !== 0 ) {
			if (this.administration.length !== other.administration.length) {
				return false;
			}

			if (!this.administration.every((value, i) => {
				if (!value.Equal(other.administration[i])) {
					return false;
				}
				return true;
			})) return false;
		}

		// Management ([]Manager)
		if (this.management.length !== 0 || other.management.length !== 0 ) {
			if (this.management.length !== other.management.length) {
				return false;
			}

			if (!this.management.every((value, i) => {
				if (!value.Equal(other.management[i])) {
					return false;
				}
				return true;
			})) return false;
		}
		return true;
	}

}


// Manager Manager is used to refer to a role that is responsible for the
// Management of an Entity.
export class Manager {
// CEO, COO, CFO, etc. Found in 'Roles' in Specification/Resources
	type; // uint8 `json:"type,omitempty"`
// Length 0-255 bytes. 0 is valid.
	name; // string `json:"name,omitempty"`
// Serialize returns the byte representation of the message.
	Serialize(): Buffer {
		const buf = new _.Writer();

		// Type (uint8)
		{
			write(buf, this.type, 'uint8');
		}

		// Name (string)
		{
			WriteVarChar(buf, this.name, 8);
		}
		return buf.buf;
	}

	Write(buf: _.Reader) {

		// Type (uint8)
		{
			this.type = read(buf, 'uint8');
		}

		// Name (string)
		{
			this.name = ReadVarChar(buf, 8);
		}
	}

	Validate() {

		// type (uint8)
		{
			if (!GetRoleType(this.type)) {
				return sprintf('Invalid role value : %d', this.type);
			}
		}

		// name (string)
		{
			if (this.name.length > (2 << 8) - 1) {
				return sprintf('varchar field name too long %d/%d', this.name.length, (2 << 8) - 1);
			}
		}
		return;
	}

	Equal(other: Manager): boolean {

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


// Oracle A Oracle defines the details of a public Oracle.
export class Oracle {
// Length 0-255 bytes. 0 is valid. Oracle X Name (eg. Coinbase, Tokenized,
	// etc.)
	name; // string `json:"name,omitempty"`
// Length 0-255 bytes. 0 is valid. If applicable: URL for REST/RPC
	// Endpoint
	url; // string `json:"url,omitempty"`
// Length 0-255 bytes. 0 is not valid. Oracle Public Key (eg. Bitcoin
	// Public key), used to confirm digital signed proofs for transfers. Can
	// also be the same public address that controls a Tokenized Oracle.
	public_key; // []byte `json:"public_key,omitempty"`
// Serialize returns the byte representation of the message.
	Serialize(): Buffer {
		const buf = new _.Writer();

		// Name (string)
		{
			WriteVarChar(buf, this.name, 8);
		}

		// URL (string)
		{
			WriteVarChar(buf, this.url, 8);
		}

		// PublicKey ([]byte)
		{
			WriteVarBin(buf, this.public_key, 8);
		}
		return buf.buf;
	}

	Write(buf: _.Reader) {

		// Name (string)
		{
			this.name = ReadVarChar(buf, 8);
		}

		// URL (string)
		{
			this.url = ReadVarChar(buf, 8);
		}

		// PublicKey ([]byte)
		{
			this.public_key = ReadVarBin(buf, 8);
		}
	}

	Validate() {

		// name (string)
		{
			if (this.name.length > (2 << 8) - 1) {
				return sprintf('varchar field name too long %d/%d', this.name.length, (2 << 8) - 1);
			}
		}

		// url (string)
		{
			if (this.url.length > (2 << 8) - 1) {
				return sprintf('varchar field url too long %d/%d', this.url.length, (2 << 8) - 1);
			}
		}

		// public_key ([]byte)
		{
			if (this.public_key.length > (2 << 8) - 1) {
				return sprintf('varbin field public_key too long %d/%d', this.public_key.length, (2 << 8) - 1);
			}
		}
		return;
	}

	Equal(other: Oracle): boolean {

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


// OutputTag A tag or category of an output used to categorize and organize
// outputs from different transactions.
export class OutputTag {
// The text of the tag.
	tag; // string `json:"tag,omitempty"`
// Serialize returns the byte representation of the message.
	Serialize(): Buffer {
		const buf = new _.Writer();

		// Tag (string)
		{
			WriteVarChar(buf, this.tag, 8);
		}
		return buf.buf;
	}

	Write(buf: _.Reader) {

		// Tag (string)
		{
			this.tag = ReadVarChar(buf, 8);
		}
	}

	Validate() {

		// tag (string)
		{
			if (this.tag.length > (2 << 8) - 1) {
				return sprintf('varchar field tag too long %d/%d', this.tag.length, (2 << 8) - 1);
			}
		}
		return;
	}

	Equal(other: OutputTag): boolean {

		// Tag (string)
		if (this.tag !== other.tag) {
			return false;
		}
		return true;
	}

}


// QuantityIndex A QuantityIndex contains a quantity, and an index. The
// quantity could be used to describe a number of tokens, or a value. The
// index is used to refer to an input or output index position.
export class QuantityIndex {
// The index of the input/output (depending on context) sending/receiving
	// the tokens.
	index; // uint16 `json:"index,omitempty"`
// Number of tokens being sent
	quantity; // uint64 `json:"quantity,omitempty"`
// Serialize returns the byte representation of the message.
	Serialize(): Buffer {
		const buf = new _.Writer();

		// Index (uint16)
		{
			write(buf, this.index, 'uint16');
		}

		// Quantity (uint64)
		{
			write(buf, this.quantity, 'uint64');
		}
		return buf.buf;
	}

	Write(buf: _.Reader) {

		// Index (uint16)
		{
			this.index = read(buf, 'uint16');
		}

		// Quantity (uint64)
		{
			this.quantity = read(buf, 'uint64');
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

	Equal(other: QuantityIndex): boolean {

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


// TargetAddress A TargetAddress defines a public address and quantity.
export class TargetAddress {
// Public address where the token balance will be changed.
	address; // PublicKeyHash `json:"address,omitempty"`
// Qty of tokens to be frozen, thawed, confiscated or reconciled. For
	// Contract-wide freezes 0 will be used.
	quantity; // uint64 `json:"quantity,omitempty"`
// Serialize returns the byte representation of the message.
	Serialize(): Buffer {
		const buf = new _.Writer();

		// Address (PublicKeyHash)
		{
			buf.write(this.address.Serialize());
		}

		// Quantity (uint64)
		{
			write(buf, this.quantity, 'uint64');
		}
		return buf.buf;
	}

	Write(buf: _.Reader) {

		// Address (PublicKeyHash)
		{
			this.address.Write(buf);
		}

		// Quantity (uint64)
		{
			this.quantity = read(buf, 'uint64');
		}
	}

	Validate() {

		// address (PublicKeyHash)
		{
			const err = this.address.Validate();
			if (err) {
				return sprintf('field address is invalid : %s', err);
			}
		}

		// quantity (uint64)
		{
		}
		return;
	}

	Equal(other: TargetAddress): boolean {

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


// VotingSystem A VotingSystem defines all details of a Voting System.
export class VotingSystem {
// eg. Special Resolutions, Ordinary Resolutions, Fundamental Matters,
	// General Matters, Directors' Vote, Poll, etc.
	name; // string `json:"name,omitempty"`
// R - Relative Threshold, A - Absolute Threshold, P - Plurality,
	// (Relative Threshold means the number of counted votes must exceed the
	// threshold % of total ballots cast. Abstentations/spoiled votes do not
	// detract from the liklihood of a vote passing as they are not included
	// in the denominator. Absolute Threshold requires the number of ballots
	// counted to exceed the threshold value when compared to the total
	// outstanding tokens. Abstentations/spoiled votes detract from the
	// liklihood of the vote passing. For example, in an absolute threshold
	// vote, if the threshold was 50% and 51% of the total outstanding tokens
	// did not vote, then the vote cannot pass. 50% of all tokens would have
	// had to vote for one vote option for the vote to be successful.
	// Plurality means the most favoured option is selected, regardless of the
	// number of votes cast or the percentage relative to other choices.
	vote_type; // byte `json:"vote_type,omitempty"`
// 0 - Standard Scoring (+1 * # of tokens owned), 1 - Weighted Scoring
	// (1st choice * Vote Max * # of tokens held, 2nd choice * Vote Max-1 * #
	// of tokens held,..etc.)
	tally_logic; // byte `json:"tally_logic,omitempty"`
// This field is ignored when VoteType is P (Plurality). Otherwise it is
	// the percentage of ballots required for a proposal to pass. Valid values
	// are greater than 0 and less than 100. 75 means 75% and greater. Only
	// applicable to Relative and Absolute Threshold vote methods. The
	// Plurality vote method requires no threshold value (NULL), as the
	// successful vote option is simply selected on the basis of highest
	// ballots cast for it.
	threshold_percentage; // uint8 `json:"threshold_percentage,omitempty"`
// Where an asset has a vote multiplier, true must be set here for the
	// vote multiplier to count, otherwise votes are simply treated as 1x per
	// token.
	vote_multiplier_permitted; // bool `json:"vote_multiplier_permitted,omitempty"`
// Token Owners must pay the fee amount to broadcast a valid Proposal. If
	// the Proposal action is valid, the smart contract will start a vote. 0
	// is valid.
	holder_proposal_fee; // uint64 `json:"holder_proposal_fee,omitempty"`
// Serialize returns the byte representation of the message.
	Serialize(): Buffer {
		const buf = new _.Writer();

		// Name (string)
		{
			WriteVarChar(buf, this.name, 8);
		}

		// VoteType (byte)
		{
			write(buf, this.vote_type, 'byte');
		}

		// TallyLogic (byte)
		{
			write(buf, this.tally_logic, 'byte');
		}

		// ThresholdPercentage (uint8)
		{
			write(buf, this.threshold_percentage, 'uint8');
		}

		// VoteMultiplierPermitted (bool)
		{
			write(buf, this.vote_multiplier_permitted, 'bool');
		}

		// HolderProposalFee (uint64)
		{
			write(buf, this.holder_proposal_fee, 'uint64');
		}
		return buf.buf;
	}

	Write(buf: _.Reader) {

		// Name (string)
		{
			this.name = ReadVarChar(buf, 8);
		}

		// VoteType (byte)
		{
			this.vote_type = read(buf, 'byte');
		}

		// TallyLogic (byte)
		{
			this.tally_logic = read(buf, 'byte');
		}

		// ThresholdPercentage (uint8)
		{
			this.threshold_percentage = read(buf, 'uint8');
		}

		// VoteMultiplierPermitted (bool)
		{
			this.vote_multiplier_permitted = read(buf, 'bool');
		}

		// HolderProposalFee (uint64)
		{
			this.holder_proposal_fee = read(buf, 'uint64');
		}
	}

	Validate() {

		// name (string)
		{
			if (this.name.length > (2 << 8) - 1) {
				return sprintf('varchar field name too long %d/%d', this.name.length, (2 << 8) - 1);
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

	Equal(other: VotingSystem): boolean {

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


