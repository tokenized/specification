package cmd

import (
	"fmt"

	"github.com/specification/internal/golang"
	"github.com/specification/internal/markdown"
	"github.com/specification/internal/python"

	"github.com/spf13/cobra"
)

const (
	FlagDebugMode = "debug"
)

var cmdGenerate = &cobra.Command{
	Use:   "generate",
	Short: "Generate code distribution",
	RunE: func(c *cobra.Command, args []string) error {
		debugMode, _ := c.Flags().GetBool(FlagDebugMode)

		if debugMode {
			fmt.Println("Debug mode enabled!")
		}

		golang.Compile()
		markdown.Compile()
		python.Compile()

		return nil
	},
}

func init() {
	cmdGenerate.Flags().Bool(FlagDebugMode, false, "Debug mode")
}
