package cmd

import (
	"context"
	"time"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/pkg/whatsonchain"
	"github.com/tokenized/specification/dist/golang/print"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var cmdDisplayTx = &cobra.Command{
	Use:   "display_tx txid",
	Short: "Download and display a transaction",
	Args:  cobra.ExactArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Wrong argument count : requires 1")
		}

		txid, err := bitcoin.NewHash32FromStr(args[0])
		if err != nil {
			return errors.Wrap(err, "txid")
		}

		woc := whatsonchain.NewService("", bitcoin.MainNet, 5*time.Second, 30*time.Second)

		tx, err := woc.GetTx(context.Background(), *txid)
		if err != nil {
			return errors.Wrap(err, "get tx")
		}

		if err := print.PrintTx(tx); err != nil {
			return errors.Wrap(err, "print tx")
		}

		return nil
	},
}
