package cmd

import (
	"encoding/hex"
	"fmt"

	"github.com/specification/dist/golang/protocol"

	"github.com/spf13/cobra"
)

const (
	FlagDebugMode = "debug"
)

var cmdDecode = &cobra.Command{
	Use:   "decode",
	Short: "Decode a message",
	RunE: func(c *cobra.Command, args []string) error {
		debugMode, _ := c.Flags().GetBool(FlagDebugMode)

		if debugMode {
			fmt.Println("Debug mode enabled!")
		}

		b, err := hex.DecodeString(raw)
		if err != nil {
			panic(err)
		}

		m, err := protocol.New(b)
		if err != nil {
			panic(err)
		}

		fmt.Printf("message (string)\n%#s\n\n", m)
		fmt.Printf("message (struct)\n%#+v\n\n", m)

		p, err := m.PayloadMessage()
		if err != nil {
			panic(err)
		}

		fmt.Printf("payload (struct)\n%#+v\n", p)

		return nil
	},
}

func init() {
	cmdDecode.Flags().Bool(FlagDebugMode, false, "Debug mode")
}
