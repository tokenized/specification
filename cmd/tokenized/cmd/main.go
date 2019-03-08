package cmd

import (
	"github.com/spf13/cobra"
)

var scCmd = &cobra.Command{
	Use:   "tokenized",
	Short: "Tokenized CLI",
}

func Execute() {
	scCmd.AddCommand(cmdGenerate)
	scCmd.Execute()
}
