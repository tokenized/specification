package cmd

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/tokenized/pkg/bitcoin"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var cmdGenerateKey = &cobra.Command{
	Use:   "gen_key",
	Short: "Generates a random WIF key",
	Args:  cobra.ExactArgs(0),
	RunE: func(c *cobra.Command, args []string) error {
		network, err := Network()
		if err != nil {
			return errors.Wrap(err, "network")
		}

		key, err := bitcoin.GenerateKey(network)
		if err != nil {
			return errors.Wrap(err, "generate")
		}

		fmt.Printf("Key (WIF): %s\n", key)

		publicKey := key.PublicKey()
		fmt.Printf("Public Key: %s\n", publicKey)
		fmt.Printf("Public Key (Base64): %s\n", base64.StdEncoding.EncodeToString(publicKey.Bytes()))

		ra, err := key.RawAddress()
		if err != nil {
			return errors.Wrap(err, "address")
		}

		fmt.Printf("P2PKH (Pay To Public Key Hash):\n")

		address := bitcoin.NewAddressFromRawAddress(ra, network)
		fmt.Printf("  Address: %s\n", address)

		lockingScript, err := key.LockingScript()
		if err != nil {
			return errors.Wrap(err, "locking script")
		}

		fmt.Printf("  Locking Script: %s\n", lockingScript)

		fmt.Printf("  Raw Address: %s\n", hex.EncodeToString(ra.Bytes()))
		fmt.Printf("  Raw Address (Base64): %s\n", base64.StdEncoding.EncodeToString(ra.Bytes()))

		return nil
	},
}
