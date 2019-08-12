package parser2

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	strMatchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	strMatchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func StrComment(s string, prefix string) string {
	parts := strings.Split(s, " ")

	lines := []string{}

	line := prefix

	for _, p := range parts {
		if len(p) == 0 {
			continue
		}

		if len(line)+len(p) > 74 {
			// line length exceeded. Add the line to our lines
			lines = append(lines, line)

			// start a new line
			line = prefix
		}

		// append the word to the line
		line = fmt.Sprintf("%v %v", line, p)
	}

	// make sure to append any remaining non-empty line
	if line != prefix {
		lines = append(lines, line)
	}

	return strings.TrimSpace(strings.Join(lines, "\n"))
}

func StrCamelCase(str string) string {
	return strings.ToLower(str[0:1]) + str[1:]
}

func StrSnakeCase(str string) string {
	result := strMatchFirstCap.ReplaceAllString(str, "${1}_${2}")
	result = strMatchAllCap.ReplaceAllString(result, "${1}_${2}")
	return strings.ToLower(result)
}

func StrKebabCase(str string) string {
	result := strMatchFirstCap.ReplaceAllString(str, "${1}-${2}")
	result = strMatchAllCap.ReplaceAllString(result, "${1}-${2}")
	return strings.ToLower(result)
}
