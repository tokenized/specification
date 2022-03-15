package markdown

import (
	"fmt"
	"io/ioutil"

	"github.com/tokenized/specification/internal/platform/parser"
	"gopkg.in/yaml.v2"
)

func Compile(
	srcPath, distPath string,
	actions parser.Schema,
	instruments parser.Schema,
	messages parser.Schema,
) {

	templateToFile(distPath, "protocol-actions.tpl", "protocol-actions.md", actions)

	templateToFile(distPath, "protocol-instruments.tpl", "protocol-instruments.md", instruments)

	templateToFile(distPath, "protocol-messages.tpl", "protocol-messages.md", messages)

	resources := fetchResources(srcPath)

	templateToFile(distPath, "protocol-resources.tpl", "protocol-resources.md", resources)
}

func templateToFile(distPath, tplFile, goFile string, data interface{}) {

	tpl := "./internal/markdown/templates/" + tplFile

	path := distPath + "/markdown/" + goFile

	parser.HtmlTemplateToFile(data, tpl, path)
}

func fetchResources(srcPath string) []parser.Resource {
	path := srcPath + "/resources/develop"

	filenames := parser.FetchFiles(path)

	items := []parser.Resource{}

	for _, filename := range filenames {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		m := parser.Resource{}
		if err := yaml.Unmarshal(data, &m); err != nil {
			panic(fmt.Errorf("file %v : %s", filename, err))
		}

		items = append(items, m)
	}

	return items
}
