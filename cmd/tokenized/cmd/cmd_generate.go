package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/tokenized/specification/internal/markdown"
	"github.com/tokenized/specification/internal/platform/parser2"

	"github.com/spf13/cobra"
)

var cmdGenerate = &cobra.Command{
	Use:   "generate",
	Short: "Generate code distribution",
	RunE: func(c *cobra.Command, args []string) error {
		debugMode, _ := c.Flags().GetBool(FlagDebugMode)

		if debugMode {
			fmt.Println("Debug mode enabled!")
		}

		// --------------------------------------------------------------------
		// Paths

		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		srcPath := filepath.FromSlash(dir + "/src")
		distPath := filepath.FromSlash(dir + "/dist")

		// --------------------------------------------------------------------
		// Prepare Schemas

		actions, _ := parser2.NewSchema(filepath.FromSlash(srcPath + "/actions/develop"))

		assets, _ := parser2.NewSchema(filepath.FromSlash(srcPath + "/assets/develop"))

		messages, _ := parser2.NewSchema(filepath.FromSlash(srcPath + "/messages/develop"))

		// fmt.Printf("%+v", actions)
		// fmt.Printf("%+v", assets)
		// fmt.Printf("%+v", messages)

		// --------------------------------------------------------------------
		// Prepare Values

		// fieldTypes := parser.NewProtocolTypes(parser.FetchFiles(srcPath, "protocol", filepath.FromSlash("develop/types")))

		// actions := parser.NewProtocolActions(fieldTypes, parser.FetchFiles(srcPath, "protocol", filepath.FromSlash("develop/actions")))

		// assets := parser.NewAssets(parser.FetchFiles(srcPath, "assets", "develop"))

		// messages := parser.NewProtocolMessages(fieldTypes, parser.FetchFiles(srcPath, "messages", "develop"))

		// resources := parser.NewProtocolResources(parser.FetchFiles(srcPath, "resources", "develop"))

		// rejectionCodes := parser.NewProtocolRejectionCodes(filepath.FromSlash(srcPath + "/resources/develop/Rejections.yaml"))

		// --------------------------------------------------------------------
		// Compile Languages

		// golang.Compile(distPath, actions, messages, fieldTypes, resources, rejectionCodes, assets)

		// json.Compile(distPath, actions, messages, fieldTypes, resources, rejectionCodes, assets)

		markdown.Compile(distPath, *actions, *assets, *messages)

		// python.Compile(distPath, actions, messages, fieldTypes, resources, rejectionCodes, assets)

		// typescript.Compile(distPath, actions, messages, fieldTypes, resources, rejectionCodes, assets)

		return nil
	},
}

func init() {
	cmdGenerate.Flags().Bool(FlagDebugMode, false, "Debug mode")
}
