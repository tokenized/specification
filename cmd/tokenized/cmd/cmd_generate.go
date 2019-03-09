package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/tokenized/specification/internal/golang"
	"github.com/tokenized/specification/internal/markdown"
	"github.com/tokenized/specification/internal/platform/parser"
	"github.com/tokenized/specification/internal/python"

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

		// -------------------------------------------------------------------------
		// Paths

		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		srcPath := dir + "/src"
		distPath := dir + "/dist"

		// -------------------------------------------------------------------------
		// Protocol

		actions := parser.NewProtocolActions(parser.FetchFiles(srcPath, "protocol", "develop/actions"))
		fieldTypes := parser.NewProtocolTypes(parser.FetchFiles(srcPath, "protocol", "develop/types"))

		// Compile languages
		golang.CompileProtocol(distPath, actions, fieldTypes)
		python.CompileProtocol(distPath, actions, fieldTypes)
		markdown.CompileProtocol(distPath, actions, fieldTypes)

		return nil
	},
}

func init() {
	cmdGenerate.Flags().Bool(FlagDebugMode, false, "Debug mode")
}
