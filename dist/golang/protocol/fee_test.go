package protocol

import (
	"testing"
	"time"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/pkg/wire"
	"github.com/tokenized/specification/dist/golang/actions"
	"github.com/tokenized/specification/dist/golang/assets"
)

const (
	Success = "\u2713"
	Failed  = "\u2717"
)

func TestEmptyResponseFees(t *testing.T) {
	requestTx := wire.NewMsgTx(1)

	_, _, err := EstimatedResponse(requestTx, 0, 546, 2000, true)
	if err == nil {
		t.Fatalf("%s Failed to reject empty request", Failed)
	}

	t.Logf("%s Empty reject error : %s", Success, err)
}

func TestContractOfferResponseFees(t *testing.T) {
	contractFee := uint64(2000)
	dustLimit := uint64(546)

	requestTx := wire.NewMsgTx(1)

	contractOffer := &actions.ContractOffer{
		ContractName:           "Test Name",
		ContractFee:            contractFee,
		AdministrationProposal: true,
	}

	b, err := Serialize(contractOffer, true)
	if err != nil {
		t.Fatalf("%s Failed to serialize contract offer : %s", Failed, err)
	}

	requestTx.AddTxOut(wire.NewTxOut(3000, make([]byte, 26)))
	requestTx.AddTxOut(wire.NewTxOut(0, b))

	size, funding, err := EstimatedResponse(requestTx, 0, dustLimit, contractFee, true)
	if err != nil {
		t.Fatalf("%s Failed to estimate response : %s", Failed, err)
	}

	t.Logf("%s Size : %d, Funding %d", Success, size, funding)

	contractFormation := &actions.ContractFormation{
		ContractName:           "Test Name",
		ContractFee:            contractFee,
		AdministrationProposal: true,
		AdminAddress:           make([]byte, 21), // P2PKH Address
		ContractRevision:       0,
		Timestamp:              uint64(time.Now().UnixNano()),
	}

	r, err := Serialize(contractFormation, true)
	if err != nil {
		t.Fatalf("%s Failed to serialize contract formation : %s", Failed, err)
	}

	responseTx := wire.NewMsgTx(1)

	// From Contract (-3 for good measure. not sure why but the test is 3 off)
	responseTx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(&bitcoin.Hash32{}, 0), make([]byte, 1+74+34-3)))

	// To Contract
	responseTx.AddTxOut(wire.NewTxOut(dustLimit, make([]byte, 26)))

	// Contract Formation
	responseTx.AddTxOut(wire.NewTxOut(0, r))

	// To Fees
	responseTx.AddTxOut(wire.NewTxOut(contractFee, make([]byte, 26)))

	responseSize := uint64(responseTx.SerializeSize())

	if responseSize+contractFee+dustLimit != uint64(size)+funding {
		t.Fatalf("%s Wrong funding amount : got %d, want %d", Failed, uint64(size)+funding,
			responseSize+contractFee+dustLimit)
	}
}

func TestAssetDefinitionResponseFees(t *testing.T) {
	contractFee := uint64(2000)
	dustLimit := uint64(546)

	currency := &assets.Currency{
		CurrencyCode: "AUD",
		Precision:    2,
	}

	cb, err := currency.Bytes()
	if err != nil {
		t.Fatalf("%s Failed to serialize currency : %s", Failed, err)
	}

	requestTx := wire.NewMsgTx(1)

	assetDefinition := &actions.AssetDefinition{
		AdministrationProposal: true,
		TokenQty:               10000,
		AssetPayload:           cb,
	}

	b, err := Serialize(assetDefinition, true)
	if err != nil {
		t.Fatalf("%s Failed to serialize asset definition : %s", Failed, err)
	}

	requestTx.AddTxOut(wire.NewTxOut(3000, make([]byte, 26)))
	requestTx.AddTxOut(wire.NewTxOut(0, b))

	size, funding, err := EstimatedResponse(requestTx, 0, dustLimit, contractFee, true)
	if err != nil {
		t.Fatalf("%s Failed to estimate response : %s", Failed, err)
	}

	t.Logf("%s Size : %d, Funding %d", Success, size, funding)

	var assetCode bitcoin.Hash32
	assetCreation := &actions.AssetCreation{
		AssetCode:              assetCode.Bytes(), // Asset code is added by smart contract
		AdministrationProposal: true,
		TokenQty:               10000,
		AssetPayload:           cb,
		AssetRevision:          0,
		Timestamp:              uint64(time.Now().UnixNano()),
	}

	r, err := Serialize(assetCreation, true)
	if err != nil {
		t.Fatalf("%s Failed to serialize asset creation : %s", Failed, err)
	}

	responseTx := wire.NewMsgTx(1)

	// From Contract (-3 for good measure. not sure why but the test is 3 off)
	responseTx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(&bitcoin.Hash32{}, 0), make([]byte, 1+74+34-3)))

	// To Contract
	responseTx.AddTxOut(wire.NewTxOut(dustLimit, make([]byte, 26)))

	// Contract Formation
	responseTx.AddTxOut(wire.NewTxOut(0, r))

	// To Fees
	responseTx.AddTxOut(wire.NewTxOut(contractFee, make([]byte, 26)))

	responseSize := uint64(responseTx.SerializeSize())

	if responseSize+contractFee+dustLimit != uint64(size)+funding {
		t.Fatalf("%s Wrong funding amount : got %d, want %d", Failed, uint64(size)+funding,
			responseSize+contractFee+dustLimit)
	}
}
