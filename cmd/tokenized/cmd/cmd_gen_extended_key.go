package cmd

import (
	"fmt"

	"github.com/tokenized/pkg/bitcoin"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var cmdGenerateExtendedKey = &cobra.Command{
	Use:   "gen_extended_key",
	Short: "Generates a random BIP32 extended private key",
	Args:  cobra.ExactArgs(0),
	RunE: func(c *cobra.Command, args []string) error {
		key, err := bitcoin.GenerateMasterExtendedKey()
		if err != nil {
			return errors.Wrap(err, "generate")
		}

		fmt.Printf("Extended Key (BIP32): %s\n", key.String58())
		return nil
	},
}
