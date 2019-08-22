package parser

import (
	"html/template"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// TemplateToFile renders a template to a file
func TemplateToFile(data interface{}, inFile, outFile string) {
	f, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}

	tmplFuncs := MakeTemplateFuncs()

	tmpl := template.Must(template.New(path.Base(inFile)).Funcs(tmplFuncs).ParseFiles(inFile))

	if err := tmpl.Execute(f, data); err != nil {
		panic(err)
	}
}

// FetchFiles fetches all the file names in a directory
func FetchFiles(dirPath string) []string {
	filenames := []string{}

	fn := func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(path, ".yaml") {
			return nil
		}

		// Do not go in to sub directories
		if filepath.Dir(path) != dirPath {
			return nil
		}

		filenames = append(filenames, path)

		return nil
	}

	if err := filepath.Walk(dirPath, fn); err != nil {
		panic(err)
	}

	return filenames
}
