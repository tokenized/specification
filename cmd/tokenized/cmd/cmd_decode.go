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
	Args:  cobra.MinimumNArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Wrong argument count : requires 1")
		}

		b, hexErr := hex.DecodeString(args[0])
		if hexErr == nil {
			if err := print.PrintBytes(b); err != nil {
				return errors.Wrap(err, "decode")
			}

			return nil
		}

		// Concatenate all arguments
		// TODO This still trims the quotes around strings out because of bash. --ce
		s := ""
		first := true
		for _, arg := range args {
			if first {
				s = arg
				first = false
			} else {
				s += " " + arg
			}
		}

		script, scriptErr := bitcoin.StringToScript(s)
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
