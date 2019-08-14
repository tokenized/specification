package parser

import (
	"fmt"
	"html"
	"html/template"
	"math"
	"regexp"
	"strings"
)

var (
	strMatchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	strMatchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func MakeTemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"minus": func(a, b int) int {
			return a - b
		},
		"add": func(a, b int) int {
			return a + b
		},
		"padding": func(str string, size int) string {
			return strings.Repeat(" ", int(math.Max(float64(size-len(str)), 0)))
		},
		"comment": func(str, chr string) string {
			return StrComment(html.UnescapeString(str), chr)
		},
		"snakecase": func(str string) string {
			return StrSnakeCase(str)
		},
		"kebabcase": func(str string) string {
			return StrKebabCase(str)
		},
		"camelcase": func(str string) string {
			return StrCamelCase(str)
		},
	}
}

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
