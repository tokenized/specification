package cmd

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/tokenized/logger"
	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/pkg/wire"
	"github.com/tokenized/specification/dist/golang/print"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	ErrNotFound = errors.New("Not Found")
	ErrTimeout  = errors.New("Timeout")
)

type TxHash struct {
	BlockHeight uint           `json:"height"`  // block height
	Hash        bitcoin.Hash32 `json:"tx_hash"` // hash of transaction
}

type HTTPError struct {
	Status  int
	Message string
}

func (err HTTPError) Error() string {
	if len(err.Message) > 0 {
		return fmt.Sprintf("HTTP Status %d : %s", err.Status, err.Message)
	}

	return fmt.Sprintf("HTTP Status %d", err.Status)
}

var cmdDisplayScript = &cobra.Command{
	Use:   "script address|hex script|asm script",
	Short: "Fetches historical transactions for a script",
	RunE: func(c *cobra.Command, args []string) error {
		ctx := logger.ContextWithLogger(context.Background(), true, false, "")

		arg := strings.Join(args, " ")
		var lockingScript bitcoin.Script

		ad, err := bitcoin.DecodeAddress(arg)
		if err == nil {
			ra := bitcoin.NewRawAddressFromAddress(ad)
			lockingScript, err = ra.LockingScript()
			if err != nil {
				return errors.Wrap(err, "locking script from address")
			}
		}

		if len(lockingScript) == 0 {
			ls, err := bitcoin.StringToScript(arg)
			if err == nil {
				lockingScript = ls
			}
		}

		if len(lockingScript) == 0 {
			b, err := hex.DecodeString(arg)
			if err == nil {
				lockingScript = bitcoin.Script(b)
			}
		}

		if len(lockingScript) == 0 {
			return fmt.Errorf("Unknown script format : \"%s\"", arg)
		}

		fmt.Printf("Script : %s\n", lockingScript)
		fmt.Printf("Script hex : %x\n", []byte(lockingScript))
		scriptHash := bitcoin.Hash32(sha256.Sum256(lockingScript))
		fmt.Printf("Script Hash : %s\n", scriptHash)

		const urlTemplate = "https://api.whatsonchain.com/v1/bsv/%s/script/%s/history"
		url := fmt.Sprintf(urlTemplate, "main", scriptHash)

		var txHashes []TxHash
		if err := getWithToken(ctx, url, "", &txHashes, time.Second*10,
			time.Second*30); err != nil {

			if httpErr, ok := errors.Cause(err).(HTTPError); ok &&
				httpErr.Status == http.StatusNotFound {
				return ErrNotFound
			}

			return errors.Wrap(err, "get script history")
		}

		fmt.Printf("%d transactions\n", len(txHashes))
		for _, txHash := range txHashes {
			tx, err := getTx(ctx, txHash.Hash)
			if err != nil {
				return errors.Wrapf(err, "get tx: %s", txHash.Hash)
			}

			fmt.Printf("\n")
			print.PrintTx(tx)
		}

		return nil
	},
}

func getTx(ctx context.Context, txid bitcoin.Hash32) (*wire.MsgTx, error) {
	const urlTemplate = "https://api.whatsonchain.com/v1/bsv/%s/tx/%s/hex"
	url := fmt.Sprintf(urlTemplate, "main", txid)

	var txBytes []byte
	if err := getWithToken(ctx, url, "", &txBytes, time.Second*10,
		time.Second*30); err != nil {

		if httpErr, ok := errors.Cause(err).(HTTPError); ok &&
			httpErr.Status == http.StatusNotFound {
			return nil, ErrNotFound
		}

		return nil, errors.Wrap(err, "get")
	}

	tx := &wire.MsgTx{}
	if err := tx.Deserialize(bytes.NewReader(txBytes)); err != nil {
		return nil, errors.Wrap(err, "deserialize")
	}

	return tx, nil
}

func getWithToken(ctx context.Context, url, token string, response interface{},
	dialTimeout, requestTimeout time.Duration) error {

	var transport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: dialTimeout,
		}).Dial,
		TLSHandshakeTimeout: dialTimeout,
	}

	var client = &http.Client{
		Timeout:   requestTimeout,
		Transport: transport,
	}

	httpRequest, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return errors.Wrap(err, "create request")
	}

	if len(token) > 0 {
		httpRequest.Header.Add("woc-api-key", token)
	}

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		if errors.Cause(err) == context.DeadlineExceeded {
			return errors.Wrap(ErrTimeout, errors.Wrap(err, "http post").Error())
		}

		return errors.Wrap(err, "http post")
	}

	if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		if httpResponse.Body != nil {
			b, rerr := ioutil.ReadAll(httpResponse.Body)
			if rerr == nil {
				return HTTPError{
					Status:  httpResponse.StatusCode,
					Message: string(b),
				}
			}
		}

		return HTTPError{Status: httpResponse.StatusCode}
	}

	defer httpResponse.Body.Close()

	if response != nil {
		if responseString, isString := response.(*string); isString {
			b, err := ioutil.ReadAll(httpResponse.Body)
			if err != nil {
				return errors.Wrap(err, "read body")
			}
			*responseString = string(b)
			return nil
		}

		if responseBytes, isBytes := response.(*[]byte); isBytes {
			h, err := ioutil.ReadAll(httpResponse.Body)
			if err != nil {
				return errors.Wrap(err, "read body")
			}

			b, err := hex.DecodeString(string(h))
			if err != nil {
				return errors.Wrap(err, "hex")
			}

			*responseBytes = b
			return nil
		}

		if err := json.NewDecoder(httpResponse.Body).Decode(response); err != nil {
			return errors.Wrap(err, "decode response")
		}
	}

	return nil
}
