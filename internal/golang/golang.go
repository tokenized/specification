package golang

import (
	"fmt"
	"os"
	"text/template"

	"github.com/tokenized/specification/internal/platform/parser"
)

func Compile(distPath string, msgs parser.Messages) {

	templateToFile(distPath, msgs, "protocol.tpl", "protocol.go")

	pms := msgs.ProtocolMessages()
	for _, pm := range pms {
		// fmt.Printf("\n\n%+v", pm)
		templateToFile(distPath, pm, "message.tpl", parser.SnakeCase(pm.Name())+".go")
		// templateToFile("protocol", distPath, msgs)
	}

}

func templateToFile(distPath string, data interface{}, tplFile, goFile string) {
	tpl := "./internal/golang/templates/" + tplFile

	path := distPath + "/golang/protocol/" + goFile

	fmt.Println(path)

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseFiles(tpl))
	if err := tmpl.Execute(f, data); err != nil {
		panic(err)
	}
}
