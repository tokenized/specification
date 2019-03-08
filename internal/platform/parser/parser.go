package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func BuildMessages(filenames []string, packageName string) Messages {
	messages := Messages{
		Package: packageName,
	}

	for _, filename := range filenames {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		m := Message{}
		if err := yaml.Unmarshal(b, &m); err != nil {
			panic(fmt.Errorf("file %v : %v", filename, err))
		}

		// This is not one of the message definitions
		if len(m.Metadata.Name) == 0 {
			continue
		}

		messages.Messages = append(messages.Messages, m)
	}

	// Order by message code
	sort.Slice(messages.Messages, func(i, j int) bool {
		return messages.Messages[i].Code < messages.Messages[j].Code
	})

	return messages
}

func FetchFiles(srcPath, packageName, version string) []string {

	dir := srcPath + "/" + packageName + "/" + version
	filenames := []string{}

	fn := func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(path, ".yaml") {
			return nil
		}

		// Do not go in to sub directories
		if filepath.Dir(path) != dir {
			return nil
		}

		filenames = append(filenames, path)

		return nil
	}

	if err := filepath.Walk(dir, fn); err != nil {
		panic(err)
	}

	return filenames
}

func reformat(s string, prefix string) string {
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

	return strings.Join(lines, "\n")
}

func SnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
