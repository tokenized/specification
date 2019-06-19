
export class RejectionCodeData {
	Code: number;
	Name: string;
	Label: string;
	Description: string;
	// Metadata: string;
}

let rejectionTypes;

export function GetRejectionCodes() {
	if (rejectionTypes) {
		return rejectionTypes;
	}

	// load := make([]RejectionCodeData, 0)

	// if err := yaml.Unmarshal([]byte(yamlRejections), &load); err != nil {
	// 	return nil, errors.Wrap(err, "Failed to unmarshal rejection codes yaml")
	// }

	rejectionTypes = {};
	// for _, value := range load {
	// 	rejectionTypes[value.Code] = value
	// }
	return rejectionTypes;
}

export function GetRejectionCode(code: number): RejectionCodeData {
	const types = GetRejectionCodes();
	return types[code];
}

export class CurrencyTypeData {
	Code: string;
	Name: string;
	Label: string;
	Description: string;
	// Metadata    string;
}

let currencyTypes;

export function GetCurrencies() {
	if (currencyTypes) return currencyTypes;
	currencyTypes = [];
	// load := make([]CurrencyTypeData, 0)

	// if err := yaml.Unmarshal([]byte(yamlCurrencies), &load); err != nil {
	// 	return nil, errors.Wrap(err, "Failed to unmarshal currencyTypes yaml")
	// }

	// currencyTypes = make(map[[3]byte]CurrencyTypeData)
	// for _, value := range load {
	// 	if len(value.Code) != 3 {
	// 		return nil, fmt.Errorf("Currency type incorrect length : %s", value.Code)
	// 	}
	// 	var code [3]byte
	// 	copy(code[:], []byte(value.Code))
	// 	currencyTypes[code] = value
	// }
	return currencyTypes;
}

export function GetCurrency(cur: number): CurrencyTypeData {
	const types = GetCurrencies();
	return types[cur];
}

export class EntityTypeData {
	Code: string;
	Name: string;
	Label: string;
	Description: string;
	// Metadata: string;
}

let entityTypes;

export function GetEntityTypes() {
	if (entityTypes) return entityTypes;
	entityTypes = [];
	// load := make([]EntityTypeData, 0)

	// if err := yaml.Unmarshal([]byte(yamlEntities), &load); err != nil {
	// 	return nil, errors.Wrap(err, "Failed to unmarshal entities yaml")
	// }

	// entityTypes = make(map[byte]EntityTypeData)
	// for _, value := range load {
	// 	if len(value.Code) > 1 {
	// 		return nil, fmt.Errorf("Entity type too long : %s", value.Code)
	// 	}
	// 	entityTypes[byte(value.Code[0])] = value
	// }
	return entityTypes;
}

export function GetEntityType(ent: number): EntityTypeData {
	const types = GetEntityTypes();
	return types[ent];
}

export class PolityType {
	Code: string;
	Name: string;
	Label: string;
	Description: string;
	// Metadata    string
}

let polityTypes;

export function GetPolityTypes() {
	if (polityTypes) return polityTypes;
	polityTypes = [];
	// load := make([]PolityType, 0)

	// if err := yaml.Unmarshal([]byte(yamlPolities), &load); err != nil {
	// 	return nil, errors.Wrap(err, "Failed to unmarshal polities yaml")
	// }

	// polityTypes = make(map[string]PolityType)
	// for _, value := range load {
	// 	polityTypes[value.Code] = value
	// }
	return polityTypes;
}

export function GetPolityType(pol: string): PolityType {
	const types = GetPolityTypes();
	return types[pol];
}

export class RoleType {
	Code: number;
	Name: string;
	Label: string;
	Description: string;
	// Metadata: string;
}

let roleTypes;

export function GetRoleTypes() {
	if (roleTypes) return roleTypes;
	roleTypes = [];
	// load := make([]RoleType, 0)

	// if err := yaml.Unmarshal([]byte(yamlRoles), &load); err != nil {
	// 	return nil, errors.Wrap(err, "Failed to unmarshal roles yaml")
	// }

	// roleTypes = make(map[uint8]RoleType)
	// for _, value := range load {
	// 	roleTypes[value.Code] = value
	// }
	return roleTypes;
}

export function GetRoleType(role: number): RoleType {
	const types = GetRoleTypes();
	return types[role];
}

class TagType {
	Code:  number;
	Label: string;
}

let tagTypes;

export function GetTagTypes() {
	if (tagTypes) return tagTypes;
	tagTypes = [];
	// load := make([]TagType, 0)

	// if err := yaml.Unmarshal([]byte(yamlTags), &load); err != nil {
	// 	return nil, errors.Wrap(err, "Failed to unmarshal tags yaml")
	// }

	// tagTypes = make(map[uint8]TagType)
	// for _, value := range load {
	// 	tagTypes[value.Code] = value
	// }
	return tagTypes;
}

export function GetTagType(tag: number): TagType {
	const types = GetTagTypes();
	return types[tag];
}

{{range .}}
{{comment (print .Name " - " .Metadata.Description) "//"}}
const yaml{{ .Name }} = `
{{ .Data }}
`;
if (0) console.log(yaml{{ .Name }});
{{ end }}