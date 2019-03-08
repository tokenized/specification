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

		// Determine paths
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		srcPath := dir + "/src"
		distPath := dir + "/dist"

		// Fetch files and actions
		files := parser.FetchFiles(srcPath, "protocol", "develop/actions")
		actions := parser.BuildActions(files, "protocol")

		// fmt.Printf("%+v", actions)

		// Compile languages
		golang.Compile(distPath, actions)
		python.Compile(distPath, actions)
		markdown.Compile(distPath, actions)

		return nil
	},
}

func init() {
	cmdGenerate.Flags().Bool(FlagDebugMode, false, "Debug mode")
}
