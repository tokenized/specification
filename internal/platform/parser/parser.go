package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
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
