import YAML from 'yaml';
import R from 'ramda';
import { Data } from './resource_data';

export class RejectionCodeData {
	code: number;
	name: string;
	label: string;
	description: string;
	// Metadata: string;
}

export class CurrencyTypeData {
	code: string;
	name: string;
	label: string;
	description: string;
	// Metadata    string;
}

export class EntityTypeData {
	code: string;
	name: string;
	label: string;
	description: string;
	// Metadata: string;
}

export class PolityType {
	code: string;
	name: string;
	label: string;
	description: string;
	// Metadata    string
}

export class RoleType {
	code: number;
	name: string;
	label: string;
	description: string;
	// Metadata: string;
}

class TagType {
	code:  number;
	label: string;
}

let tagTypes;
let roleTypes;
let polityTypes;
let entityTypes;
let currencyTypes;
let rejectionTypes;

export class Resources {

	static GetRejectionCodes() {
		if (rejectionTypes) return rejectionTypes;
		rejectionTypes = R.indexBy((v: RejectionCodeData) => ''+v.code, YAML.parse(Data.yamlRejections));
		return rejectionTypes;
	}

	static GetRejectionCode(code: number): RejectionCodeData {
		const types = Resources.GetRejectionCodes();
		return types[''+code];
	}

	static GetCurrencies() {
		if (currencyTypes) return currencyTypes;
		currencyTypes = R.indexBy((v: CurrencyTypeData) => v.code, YAML.parse(Data.yamlCurrencies));
		return currencyTypes;
	}

	static GetCurrency(cur: string): CurrencyTypeData {
		const types = Resources.GetCurrencies();
		return types[cur];
	}

	static GetEntityTypes() {
		if (entityTypes) return entityTypes;
		entityTypes = R.indexBy((v: EntityTypeData) => v.code.charAt(0), YAML.parse(Data.yamlEntities));
		return entityTypes;
	}

	static GetEntityType(ent: number): EntityTypeData {
		const types = Resources.GetEntityTypes();
		return types[String.fromCharCode(ent)];
	}

	static GetPolityTypes() {
		if (polityTypes) return polityTypes;
		polityTypes = R.indexBy((v: PolityType) => v.code, YAML.parse(Data.yamlPolities));
		return polityTypes;
	}

	static GetPolityType(pol: string): PolityType {
		const types = Resources.GetPolityTypes();
		return types[pol];
	}

	static GetRoleTypes() {
		if (roleTypes) return roleTypes;
		roleTypes = R.indexBy((v: RoleType) => ''+v.code, YAML.parse(Data.yamlRoles));
		return roleTypes;
	}

	static GetRoleType(role: number): RoleType {
		const types = Resources.GetRoleTypes();
		return types[''+role];
	}

	static GetTagTypes() {
		if (tagTypes) return tagTypes;
		tagTypes = R.indexBy((v: RoleType) => ''+v.code, YAML.parse(Data.yamlTags));
		return tagTypes;
	}

	static GetTagType(tag: number): TagType {
		const types = Resources.GetTagTypes();
		return types[''+tag];
	}
}
