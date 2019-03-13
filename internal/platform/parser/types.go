package parser

import (
	"fmt"
	"strings"
)

func GoType(typeName string, size uint64) string {
	name := typeName

	prefix := ""
	if strings.HasSuffix(name, "[]") {
		prefix = "[]"
		name = typeName[:len(name)-2]
	}

	switch name {
	// Temporary
	case "polity":
		return "[]byte"
	case "dropdown":
		return "[]byte"

	case "varchar":
		return "string"
	case "fixedchar":
		if size == 1 {
			return "byte"
		}
		return "string"

	case "opcode":
		return "byte"

	case "sha256",
		"SHA256":
		return "[32]byte"

	case "sha",
		"SHA",
		"bin[var]",
		"bin",
		"pushdata_length",
		"payload":
		return "[]byte"

	case "time", "timestamp":
		return "uint64"

	case "uint", "int", "float":
		return fmt.Sprintf("%v%s%v", prefix, strings.ToLower(name), size*8)

	case "nvarchar8",
		"nvarchar16",
		"nvarchar32",
		"nvarchar64":
		return prefix + strings.Title(typeName)

	case "Header",
		"header":
		return "Header"

	}

	return fmt.Sprintf("%s%s", prefix, name)
}
