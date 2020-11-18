package parser

import (
	"fmt"
	"html"
	htmlTemplate "html/template"
	"math"
	"regexp"
	"strings"
	"text/template"
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
		"stripwhitespace": func(str string) string {
			return StrStripWhiteSpace(str)
		},
		"yamldescription": func(str, pad string) string {
			return StrYamlDescription(str, pad)
		},
		"fixedhex": func(size int) string {
			return FixedHex(size)
		},
	}
}

func MakeHtmlTemplateFuncs() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap{
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
		"stripwhitespace": func(str string) string {
			return StrStripWhiteSpace(str)
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

func StrYamlDescription(s, pad string) string {
	padLen := len(pad)
	s = strings.TrimSpace(s)
	if len(s) < 100-padLen-len("description") {
		if strings.Contains(s, ":") {
			return fmt.Sprintf("\"%s\"", s)
		}
		return s
	}

	parts := strings.Split(s, " ")

	lines := []string{" >"}

	line := ""

	for _, p := range parts {
		if len(p) == 0 {
			line += " "
			continue
		}

		if len(line)+len(p) > 100-padLen {
			// line length exceeded. Add the line to our lines
			lines = append(lines, pad+line)
			line = ""
		}

		// append the word to the line
		if len(line) == 0 {
			line = p
		} else {
			line = fmt.Sprintf("%v %v", line, p)
		}

		if line[len(line)-1] == '\n' {
			lines = append(lines, pad+line[:len(line)-1])
			line = ""
		}
	}

	// make sure to append any remaining non-empty line
	if line != pad {
		lines = append(lines, pad+line)
	}

	return strings.Join(lines, "\n")
}

func StrCamelCase(str string) string {
	return strings.ToLower(str[0:1]) + str[1:]
}

func StrStripWhiteSpace(str string) string {
	result := ""
	for _, c := range str {
		switch c {
		case ' ':
		case '-':
		case ',':
		case '\'':
		case '(':
		case ')':
		case '.':
		default:
			result += string(c)
		}
	}
	return result
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

func FixedHex(size int) string {
	b := make([]byte, size)
	for i, _ := range b {
		b[i] = byte((i + 3) * 7)
	}
	return fmt.Sprintf("%x", b)
}
