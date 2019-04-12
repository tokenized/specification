package parser

import (
	"fmt"
	"strings"
)

func GoType(typeName string, size uint64, options []string) string {
	name := typeName

	prefix := ""
	if strings.HasSuffix(name, "[]") {
		prefix = "[]"
		name = typeName[:len(name)-2]
	}

	switch name {

	case "varchar":
		return "string"
	case "fixedchar":
		if size == 1 {
			return "byte"
		}
		return "string"

	case "varbin":
		return "[]byte"

	case "bin":
		return fmt.Sprintf("[%d]byte", size)

	case "uint", "int", "float":
		return fmt.Sprintf("%v%s%v", prefix, strings.ToLower(name), size*8)

	case "bool":
		return "bool"

	case "Header",
		"header":
		return "Header"

	case "Role":
		return "uint8"

	case "MessageType":
		return "uint16"

	case "EntityType":
		return "byte" // single character

	case "Currency":
		return "[3]byte"

	case "RejectionCode":
		return "uint8"
	}

	return fmt.Sprintf("%s%s", prefix, name)
}

func IsInternalType(typeName string, size uint64) bool {
	name := typeName
	if strings.HasSuffix(name, "[]") {
		name = typeName[:len(name)-2]
	}

	switch name {

	case "varchar":
		return false
	case "fixedchar":
		return false

	case "bin",
		"varbin":
		return false

	case "uint", "uint8", "uint16", "uint32", "uint64", "int", "int8", "int16", "int32", "int64", "float", "float32", "float64":
		return false

	case "bool":
		return false

	case "Role",
		"EntityType",
		"Currency",
		"MessageType",
		"RejectionCode":
		return false

	}

	return true
}

func IsResource(typeName string) bool {
	switch typeName {
	case "Role",
		"EntityType",
		"Currency",
		"MessageType",
		"RejectionCode":
		return true
	}

	return false
}
