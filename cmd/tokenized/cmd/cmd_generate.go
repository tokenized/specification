package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/tokenized/specification/internal/golang"
	"github.com/tokenized/specification/internal/platform/parser"

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

		// Fetch files and messages
		allFiles := parser.FetchFiles(srcPath, "protocol", "develop")
		allMessages := parser.BuildMessages(allFiles, "protocol")

		// fmt.Printf("%+v", allMessages)

		// Compile languages
		golang.Compile(distPath, allMessages)
		// markdown.Compile(distPath, allMessages)
		// python.Compile(distPath, allMessages)

		return nil
	},
}

func init() {
	cmdGenerate.Flags().Bool(FlagDebugMode, false, "Debug mode")
}
