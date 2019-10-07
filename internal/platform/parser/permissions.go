package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/pkg/errors"
)

func ProcessContractPermissionConfigs(actions, assets Schema, path string, outFile string) error {
	f, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}

	f.WriteString("package actions\n")
	f.WriteString("\n")
	f.WriteString("type PermissionConfig struct {\n")
	f.WriteString("    VotingSystems       []VotingSystemField\n")
	f.WriteString("    ContractPermissions Permissions\n")
	f.WriteString("    AssetPermissions    map[string]Permissions\n")
	f.WriteString("}\n")

	path = filepath.FromSlash(path)
	dirs, err := filepath.Glob(filepath.Join(path, "*"))
	if err != nil {
		panic(err)
	}

	for _, dir := range dirs {
		data := PermissionConfig{
			AssetPermissions: make(map[string][]Permission),
		}
		fmt.Printf("Translating %s\n", dir)

		if err := unmarshalFile(filepath.Join(dir, "VotingSystems.yaml"), &data); err != nil {
			fmt.Printf("Failed to unmarshal yaml voting systems : %s\n", err)
			panic(errors.Wrap(err, "unmarshal yaml voting systems"))
		}

		if err := unmarshalFile(filepath.Join(dir, "Contract.yaml"), &data); err != nil {
			fmt.Printf("Failed to unmarshal yaml contract : %s\n", err)
			panic(errors.Wrap(err, "unmarshal yaml contract"))
		}

		assetFiles, err := filepath.Glob(filepath.Join(dir, "assets", "*"))
		if err != nil {
			panic(err)
		}

		assetStruct := struct {
			Name        string       `yaml:"Name"`
			AssetType   string       `yaml:"AssetType"`
			Permissions []Permission `yaml:"Permissions"`
		}{}

		for _, assetFile := range assetFiles {
			if err := unmarshalFile(assetFile, &assetStruct); err != nil {
				fmt.Printf("Failed to unmarshal yaml asset : %s\n", err)
				panic(errors.Wrap(err, "unmarshal yaml asset"))
			}

			data.AssetPermissions[assetStruct.AssetType] = assetStruct.Permissions
		}

		if err := TranslateContractPermissionConfig(actions, assets, data, f); err != nil {
			fmt.Printf("Failed to translate permissions : %s\n", err)
			panic(errors.Wrap(err, fmt.Sprintf("contract %s", dir)))
		}
	}

	return f.Close()
}

func TranslateContractPermissionConfig(actions, assets Schema, data PermissionConfig, file *os.File) error {

	file.WriteString(fmt.Sprintf("var %s = PermissionConfig{\n", data.Name))

	file.WriteString("    VotingSystems: []VotingSystemField{\n")
	for _, votingSystem := range data.VotingSystems {
		file.WriteString(fmt.Sprintf("        VotingSystemField{\n"))
		file.WriteString(fmt.Sprintf("            Name:                   \"%s\",\n", votingSystem.Name))
		file.WriteString(fmt.Sprintf("            VoteType:               \"%s\",\n", votingSystem.VoteType))
		file.WriteString(fmt.Sprintf("            TallyLogic:              %d,\n", votingSystem.TallyLogic))
		file.WriteString(fmt.Sprintf("            ThresholdPercentage:     %d,\n", votingSystem.ThresholdPercentage))
		file.WriteString(fmt.Sprintf("            VoteMultiplierPermitted: %t,\n", votingSystem.VoteMultiplierPermitted))
		file.WriteString(fmt.Sprintf("            HolderProposalFee:       %d,\n", votingSystem.HolderProposalFee))
		file.WriteString(fmt.Sprintf("        },\n"))
	}
	file.WriteString("    },\n")

	file.WriteString("    ContractPermissions: Permissions{\n")
	fips := make([][]int, 0)
	var err error
	for _, permission := range data.ContractPermissions {
		fips, err = TranslatePermission("", actions, assets, actions, permission, data.VotingSystems,
			"ContractOffer", fips, file)
		if err != nil {
			return errors.Wrap(err, "contract permissions")
		}
	}
	file.WriteString("    },\n")

	assetTypes := make([]string, 0, len(data.AssetPermissions))
	for assetType, _ := range data.AssetPermissions {
		assetTypes = append(assetTypes, assetType)
	}
	sort.Strings(assetTypes)

	file.WriteString("    AssetPermissions: map[string]Permissions{\n")
	for _, assetType := range assetTypes {
		file.WriteString(fmt.Sprintf("        \"%s\": Permissions{\n", assetType))
		fips = make([][]int, 0)
		for _, permission := range data.AssetPermissions[assetType] {
			fips, err = TranslatePermission(assetType, actions, assets, assets, permission,
				data.VotingSystems, "AssetDefinition", fips, file)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("asset permissions %s", assetType))
			}
		}
		file.WriteString("        },\n")
	}
	file.WriteString("    },\n")

	file.WriteString("}\n")
	return nil
}

func TranslatePermission(assetType string, actions, assets, schema Schema, permission Permission,
	votingSystems []VotingSystem, structName string, fips [][]int, file *os.File) ([][]int, error) {

	file.WriteString(fmt.Sprintf("        Permission{ // %s\n", permission.Name))

	// Authorization flags
	file.WriteString(fmt.Sprintf("            Permitted: %t,\n", permission.Permitted))
	file.WriteString(fmt.Sprintf("            AdministrationProposal: %t,\n", permission.AdministrationProposal))
	file.WriteString(fmt.Sprintf("            HolderProposal: %t,\n", permission.HolderProposal))
	file.WriteString(fmt.Sprintf("            AdministrativeMatter: %t,\n", permission.AdministrativeMatter))

	// VotingSystemsAllowed
	file.WriteString(fmt.Sprintf("            VotingSystemsAllowed: []bool{"))
	first := true
	for _, system := range votingSystems {
		found := false
		for _, name := range permission.VotingSystemsAllowed {
			if name == system.Name {
				found = true
				break
			}
		}
		if first {
			first = false
		} else {
			file.WriteString(", ")
		}
		file.WriteString(fmt.Sprintf("%t", found))
	}
	file.WriteString("},\n")

	// Fields
	file.WriteString("            Fields: []FieldIndexPath{\n")
	for _, fieldNames := range permission.Fields {
		var fields []Field
		found := false
		container := structName
		for _, m := range actions.Messages {
			if m.Name == container {
				fields = m.Fields
				found = true
				break
			}
		}
		if !found {
			return fips, fmt.Errorf("Failed to find %s message", container)
		}
		file.WriteString(fmt.Sprintf("                FieldIndexPath{"))
		first := true
		fieldIndexPath := make([]int, 0, len(permission.Fields))
		for j, fieldName := range fieldNames {
			found := false
			fieldType := ""
			for i, field := range fields {
				if field.Name == fieldName {
					fieldType = field.BaseType()
					if first {
						first = false
					} else {
						file.WriteString(", ")
					}
					file.WriteString(fmt.Sprintf("%d", i+1))
					fieldIndexPath = append(fieldIndexPath, i+1)
					found = true
					break
				}
			}
			if !found {
				return fips, fmt.Errorf("Failed to find %s field in %s", fieldName, container)
			}

			if j == len(fieldNames)-1 {
				break
			}
			container = fieldName
			found = false
			if fieldName == "AssetPayload" {
				for _, m := range assets.Messages {
					if m.Code == assetType {
						fields = m.Fields
						found = true
						break
					}
				}
			} else {
				for _, ft := range schema.FieldTypes {
					if ft.Name == fieldType {
						fields = ft.Fields
						found = true
						break
					}
				}
			}
			if !found {
				return fips, fmt.Errorf("Failed to find struct %s (%s)", fieldType, fieldName)
			}
		}

		// Check for duplicates
		for _, fip := range fips {
			matches := true
			for i, index := range fieldIndexPath {
				if i >= len(fip) {
					if i > 0 {
						return fips, fmt.Errorf("Duplicate partial field index path %v", fieldNames)
					}
					break
				}
				if index != fip[i] {
					matches = false
					break
				}
			}
			if matches {
				return fips, fmt.Errorf("Duplicate field index path %v", fieldNames)
			}
		}

		fips = append(fips, fieldIndexPath)
		file.WriteString(fmt.Sprintf("}, // %v\n", fieldNames))
	}

	file.WriteString("            },\n")
	file.WriteString("        },\n")
	return fips, nil
}
