package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/pkg/errors"
)

func ProcessContractPermissionConfigs(actions, instruments Schema, path string, outFile string) error {
	f, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}

	f.WriteString("package actions\n")
	f.WriteString("\n")
	f.WriteString("import \"github.com/tokenized/specification/dist/golang/permissions\"\n")
	f.WriteString("\n")
	f.WriteString("type PermissionConfig struct {\n")
	f.WriteString("    VotingSystems       []VotingSystemField\n")
	f.WriteString("    ContractPermissions permissions.Permissions\n")
	f.WriteString("    InstrumentPermissions    map[string]permissions.Permissions\n")
	f.WriteString("}\n")

	path = filepath.FromSlash(path)
	dirs, err := filepath.Glob(filepath.Join(path, "*"))
	if err != nil {
		panic(err)
	}

	for _, dir := range dirs {
		data := PermissionConfig{
			InstrumentPermissions: make(map[string][]Permission),
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

		instrumentFiles, err := filepath.Glob(filepath.Join(dir, "instruments", "*"))
		if err != nil {
			panic(err)
		}

		instrumentStruct := struct {
			Name           string       `yaml:"Name"`
			InstrumentType string       `yaml:"InstrumentType"`
			Permissions    []Permission `yaml:"Permissions"`
		}{}

		for _, instrumentFile := range instrumentFiles {
			if err := unmarshalFile(instrumentFile, &instrumentStruct); err != nil {
				fmt.Printf("Failed to unmarshal yaml instrument : %s\n", err)
				panic(errors.Wrap(err, "unmarshal yaml instrument"))
			}

			data.InstrumentPermissions[instrumentStruct.InstrumentType] = instrumentStruct.Permissions
		}

		if err := TranslateContractPermissionConfig(actions, instruments, data, f); err != nil {
			fmt.Printf("Failed to translate permissions : %s\n", err)
			panic(errors.Wrap(err, fmt.Sprintf("contract %s", dir)))
		}
	}

	return f.Close()
}

func TranslateContractPermissionConfig(actions, instruments Schema, data PermissionConfig, file *os.File) error {

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

	file.WriteString("    ContractPermissions: permissions.Permissions{\n")
	fips := make([][]int, 0)
	var err error
	for _, permission := range data.ContractPermissions {
		fips, err = TranslatePermission("", actions, instruments, actions, permission, data.VotingSystems,
			"ContractOffer", fips, file)
		if err != nil {
			return errors.Wrap(err, "contract permissions")
		}
	}
	file.WriteString("    },\n")

	instrumentTypes := make([]string, 0, len(data.InstrumentPermissions))
	for instrumentType, _ := range data.InstrumentPermissions {
		instrumentTypes = append(instrumentTypes, instrumentType)
	}
	sort.Strings(instrumentTypes)

	file.WriteString("    InstrumentPermissions: map[string]permissions.Permissions{\n")
	for _, instrumentType := range instrumentTypes {
		file.WriteString(fmt.Sprintf("        \"%s\": permissions.Permissions{\n", instrumentType))
		fips = make([][]int, 0)
		for _, permission := range data.InstrumentPermissions[instrumentType] {
			fips, err = TranslatePermission(instrumentType, actions, instruments, instruments, permission,
				data.VotingSystems, "InstrumentDefinition", fips, file)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("instrument permissions %s", instrumentType))
			}
		}
		file.WriteString("        },\n")
	}
	file.WriteString("    },\n")

	file.WriteString("}\n")
	return nil
}

func TranslatePermission(instrumentType string, actions, instruments, schema Schema, permission Permission,
	votingSystems []VotingSystem, structName string, fips [][]int, file *os.File) ([][]int, error) {

	file.WriteString(fmt.Sprintf("        permissions.Permission{ // %s\n", permission.Name))

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
	file.WriteString("            Fields: []permissions.FieldIndexPath{\n")
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
		file.WriteString(fmt.Sprintf("                permissions.FieldIndexPath{"))
		first := true
		fieldIndexPath := make([]int, 0, len(permission.Fields))
		for j, fieldName := range fieldNames {
			if fieldName == "0" {
				// match all for list index
				if first {
					first = false
				} else {
					file.WriteString(", ")
				}
				file.WriteString("0")
				fieldIndexPath = append(fieldIndexPath, 0)
				continue
			}

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
			if fieldName == "InstrumentPayload" {
				for _, m := range instruments.Messages {
					if m.Code == instrumentType {
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
