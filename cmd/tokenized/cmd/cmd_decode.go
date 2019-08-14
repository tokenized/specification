package cmd

import (
	"github.com/spf13/cobra"
)

var cmdDecode = &cobra.Command{
	Use:   "decode",
	Short: "Decode a message",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
		/*
			raw := args[0]

			testMode, _ := c.Flags().GetBool(FlagTestMode)

			if testMode {
				fmt.Println("Test mode enabled!")
			}

			b, err := hex.DecodeString(raw)
			if err != nil {
				panic(err)
			}

			m, err := protocol.Deserialize(b, testMode)
			if err != nil {
				panic(err)
			}

			fmt.Printf("message (string)\n%#s\n\n", m)
			fmt.Printf("message (struct)\n%#+v\n\n", m)

			// TODO: Asset payloads

			// p := m.Type()
			// if err != nil {
			// 	panic(err)
			// }

			// fmt.Printf("payload (struct)\n%#+v\n", p)
		*/
		return nil
	},
}

func init() {
	cmdDecode.Flags().Bool(FlagTestMode, false, "Test mode")
}
