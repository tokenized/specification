package cmd

import (
	"github.com/spf13/cobra"
)

var cmdDecode = &cobra.Command{
	Use:   "decode",
	Short: "Decode a message",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(c *cobra.Command, args []string) error {

		// raw := args[0]

		// debugMode, _ := c.Flags().GetBool(FlagDebugMode)

		// if debugMode {
		// 	fmt.Println("Debug mode enabled!")
		// }

		// b, err := hex.DecodeString(raw)
		// if err != nil {
		// 	panic(err)
		// }

		// m, err := protocol.New(b)
		// if err != nil {
		// 	panic(err)
		// }

		// fmt.Printf("message (string)\n%#s\n\n", m)
		// fmt.Printf("message (struct)\n%#+v\n\n", m)

		// p, err := m.PayloadMessage()
		// if err != nil {
		// 	panic(err)
		// }

		// fmt.Printf("payload (struct)\n%#+v\n", p)

		return nil
	},
}

func init() {
	cmdDecode.Flags().Bool(FlagDebugMode, false, "Debug mode")
}
