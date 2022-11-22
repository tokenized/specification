package cmd

import (
	"encoding/hex"
	"fmt"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/specification/dist/golang/print"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var cmdDecode = &cobra.Command{
	Use:   "decode hex_or_asm",
	Short: "Decode a transaction or output script (hex or asm)",
	Args:  cobra.ExactArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Wrong argument count : requires 1")
		}

		b, hexErr := hex.DecodeString(args[0])
		if hexErr == nil {
			if err := print.PrintBytes(b); err != nil {
				return errors.Wrap(err, "decode")
			}

			return nil
		}

		script, scriptErr := bitcoin.StringToScript(args[0])
		if scriptErr == nil {
			if err := print.PrintScript(script, false); err != nil {
				return errors.Wrap(err, "script")
			}

			return nil
		}

		return fmt.Errorf("Not hex or asm script: hex error (%s), asm error (%s)", hexErr,
			scriptErr)
	},
}
