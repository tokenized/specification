package cmd

import (
	"github.com/spf13/cobra"
)

var scCmd = &cobra.Command{
	Use:   "tokenized",
	Short: "Tokenized CLI",
}

func Execute() {
	scCmd.AddCommand(cmdDecode)
	scCmd.AddCommand(cmdGenerateKey)
	scCmd.AddCommand(cmdGenerateExtendedKey)
	scCmd.AddCommand(cmdCreateContract)
	scCmd.AddCommand(cmdCreateInstrument)
	scCmd.AddCommand(cmdCreateTransfer)
	scCmd.AddCommand(cmdDisplayTx)
	scCmd.AddCommand(cmdDisplayScript)
	scCmd.Execute()
}
