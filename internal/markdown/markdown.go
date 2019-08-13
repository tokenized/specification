package markdown

import (
	"fmt"
	"io/ioutil"

	"github.com/tokenized/specification/internal/platform/parser2"
	"gopkg.in/yaml.v2"
)

func Compile(
	srcPath, distPath string,
	actions parser2.Schema,
	assets parser2.Schema,
	messages parser2.Schema,
) {

	templateToFile(distPath, "protocol-actions.tpl", "protocol-actions.md", actions)

	templateToFile(distPath, "protocol-assets.tpl", "protocol-assets.md", assets)

	templateToFile(distPath, "protocol-messages.tpl", "protocol-messages.md", messages)

	resources := fetchResources(srcPath)

	templateToFile(distPath, "protocol-resources.tpl", "protocol-resources.md", resources)
}

func templateToFile(distPath, tplFile, goFile string, data interface{}) {

	tpl := "./internal/markdown/templates/" + tplFile

	path := distPath + "/markdown/" + goFile

	parser2.TemplateToFile(distPath, data, tpl, path)
}

func fetchResources(srcPath string) []parser2.Resource {
	path := srcPath + "/resources/develop"

	filenames := parser2.FetchFiles(path)

	items := []parser2.Resource{}

	for _, filename := range filenames {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		m := parser2.Resource{}
		if err := yaml.Unmarshal(data, &m); err != nil {
			panic(fmt.Errorf("file %v : %s", filename, err))
		}

		items = append(items, m)
	}

	return items
}
