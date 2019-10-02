package parser

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func ProcessContractPermissionConfigs(actions Schema, path string, outFile string) error {
	f, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}

	f.WriteString("package actions\n")
	f.WriteString("\n")
	f.WriteString("type PermissionConfig struct {\n")
	f.WriteString("    VotingSystems []VotingSystemField\n")
	f.WriteString("    Permissions   Permissions\n")
	f.WriteString("}\n")

	path = filepath.FromSlash(path)
	files, err := filepath.Glob(filepath.Join(path, "*.yaml"))
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		var data PermissionConfig
		fmt.Printf("Translating %s\n", file)
		if err := unmarshalFile(file, &data); err != nil {
			fmt.Printf("Failed to unmarshal yaml : %s\n", err)
			panic(errors.Wrap(err, "unmarshal yaml"))
		}

		if err := TranslateContractPermissionConfig(actions, data, f); err != nil {
			fmt.Printf("Failed to translate permissions : %s\n", err)
			panic(errors.Wrap(err, "translate permissions"))
		}
	}

	return f.Close()
}

func TranslateContractPermissionConfig(actions Schema, data PermissionConfig, file *os.File) error {

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

	file.WriteString("    Permissions: Permissions{\n")
	for _, permission := range data.Permissions {
		file.WriteString(fmt.Sprintf("        Permission{ // %s\n", permission.Name))

		// Authorization flags
		file.WriteString(fmt.Sprintf("            Permitted: %t,\n", permission.Permitted))
		file.WriteString(fmt.Sprintf("            AdministrationProposal: %t,\n", permission.AdministrationProposal))
		file.WriteString(fmt.Sprintf("            HolderProposal: %t,\n", permission.HolderProposal))
		file.WriteString(fmt.Sprintf("            AdministrativeMatter: %t,\n", permission.AdministrativeMatter))

		// VotingSystemsAllowed
		file.WriteString(fmt.Sprintf("            VotingSystemsAllowed: []bool{"))
		first := true
		for _, system := range data.VotingSystems {
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
			// file.WriteString(fmt.Sprintf("            %t,\n", found))
		}
		file.WriteString("},\n")

		// Fields
		file.WriteString("            Fields: []FieldIndexPath{\n")
		for _, fieldNames := range permission.Fields {
			var fields []Field
			found := false
			structName := "ContractFormation"
			for _, m := range actions.Messages {
				if m.Name == structName {
					fields = m.Fields
					found = true
					break
				}
			}
			if !found {
				panic(fmt.Errorf("Failed to find %s message", structName))
			}
			file.WriteString(fmt.Sprintf("                FieldIndexPath{"))
			first := true
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
						found = true
						break
					}
				}
				if !found {
					panic(fmt.Errorf("Failed to find %s field in %s", fieldName, structName))
				}

				if j == len(fieldNames)-1 {
					break
				}
				structName = fieldName
				found = false
				for _, ft := range actions.FieldTypes {
					if ft.Name == fieldType {
						fields = ft.Fields
						found = true
						break
					}
				}
				if !found {
					panic(fmt.Errorf("Failed to find struct %s", fieldType))
				}
			}
			file.WriteString(fmt.Sprintf("}, // %v\n", fieldNames))
		}
		file.WriteString("            },\n")
		file.WriteString("        },\n")
	}
	file.WriteString("    },\n")

	file.WriteString("}\n")
	return nil
}
