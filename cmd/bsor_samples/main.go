package main

import (
	"encoding/json"
	"os"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/pkg/bsor"
	"github.com/tokenized/specification/dist/golang/actions"
	"github.com/tokenized/specification/dist/golang/protocol"
)

func main() {
	publicKey, err := bitcoin.PublicKeyFromStr("02dd698867b0a0f7c1e70ad36db7e4860febabbebc599c1bbdd1ce7737ca922014")
	if err != nil {
		panic(err)
	}

	co := &actions.ContractOffer{
		ContractName: "Test Contract Name",
		ContractFee:  2000,
		Services: []*actions.ServiceField{
			{
				Type:      actions.ServiceTypeIdentityOracle, // 0 will not be encoded
				URL:       "test://test_url.com",
				PublicKey: publicKey.Bytes(),
			},
			{
				Type: actions.ServiceTypeAuthorityOracle, // 1
				URL:  "test://test_url.com",
			},
		},
	}

	b, err := bsor.MarshalBinary(co)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("contract_offer.bsor")
	if err != nil {
		panic(err)
	}
	if _, err := file.Write(b); err != nil {
		panic(err)
	}
	file.Close()

	js, err := json.MarshalIndent(co, "", "  ")
	if err != nil {
		panic(err)
	}

	file, err = os.Create("contract_offer.json")
	if err != nil {
		panic(err)
	}
	if _, err := file.Write(js); err != nil {
		panic(err)
	}
	file.Close()

	script, err := protocol.Serialize(co, true)
	if err != nil {
		panic(err)
	}

	file, err = os.Create("contract_offer.envelope")
	if err != nil {
		panic(err)
	}
	if _, err := file.Write(script); err != nil {
		panic(err)
	}
	file.Close()
}
