package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/tokenized/specification/internal/golang"
	"github.com/tokenized/specification/internal/json"
	"github.com/tokenized/specification/internal/markdown"
	"github.com/tokenized/specification/internal/platform/parser"
	"github.com/tokenized/specification/internal/python"
	"github.com/tokenized/specification/internal/typescript"

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
		// Prepare Values

		fieldTypes := parser.NewProtocolTypes(parser.FetchFiles(srcPath, "protocol", filepath.FromSlash("develop/types")))

		actions := parser.NewProtocolActions(fieldTypes, parser.FetchFiles(srcPath, "protocol", filepath.FromSlash("develop/actions")))

		assets := parser.NewAssets(parser.FetchFiles(srcPath, "assets", "develop"))

		messages := parser.NewProtocolMessages(fieldTypes, parser.FetchFiles(srcPath, "messages", "develop"))

		resources := parser.NewProtocolResources(parser.FetchFiles(srcPath, "resources", "develop"))

		rejectionCodes := parser.NewProtocolRejectionCodes(filepath.FromSlash(srcPath + "/resources/develop/Rejections.yaml"))

		// --------------------------------------------------------------------
		// Compile Languages

		golang.Compile(distPath, actions, messages, fieldTypes, resources, rejectionCodes, assets)

		json.Compile(distPath, actions, messages, fieldTypes, resources, rejectionCodes, assets)

		markdown.Compile(distPath, actions, messages, fieldTypes, resources, rejectionCodes, assets)

		python.Compile(distPath, actions, messages, fieldTypes, resources, rejectionCodes, assets)

		typescript.Compile(distPath, actions, messages, fieldTypes, resources, rejectionCodes, assets)

		return nil
	},
}

func init() {
	cmdGenerate.Flags().Bool(FlagDebugMode, false, "Debug mode")
}
