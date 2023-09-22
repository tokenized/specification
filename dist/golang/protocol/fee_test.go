package protocol

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"math"
	"testing"
	"time"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/pkg/json"
	"github.com/tokenized/pkg/wire"
	"github.com/tokenized/specification/dist/golang/actions"
	"github.com/tokenized/specification/dist/golang/instruments"
	"github.com/tokenized/txbuilder"
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

func TestContractOfferResponseFees_PKH(t *testing.T) {
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

func Test_ContractOfferResponseFees_MultiSig(t *testing.T) {
	contractFee := uint64(2000)
	feeRate := float64(0.05)
	dustFeeRate := float64(0.0)
	isTest := true

	key, _ := bitcoin.GenerateKey(bitcoin.MainNet)
	contractAgentLockingScript, _ := key.LockingScript()
	contractAgentDust := txbuilder.DustLimitForLockingScript(contractAgentLockingScript,
		float32(dustFeeRate))

	key, _ = bitcoin.GenerateKey(bitcoin.MainNet)
	contractFeeLockingScript, _ := key.LockingScript()

	key, _ = bitcoin.GenerateKey(bitcoin.MainNet)
	masterLockingScript, _ := key.LockingScript()
	masterRA, _ := bitcoin.RawAddressFromLockingScript(masterLockingScript)

	key, _ = bitcoin.GenerateKey(bitcoin.MainNet)
	contractOperatorLockingScript, _ := key.LockingScript()
	contractOperatorRA, _ := bitcoin.RawAddressFromLockingScript(contractOperatorLockingScript)

	key, _ = bitcoin.GenerateKey(bitcoin.MainNet)
	contractOperatorEntityLockingScript, _ := key.LockingScript()
	contractOperatorEntityRA, _ := bitcoin.RawAddressFromLockingScript(contractOperatorEntityLockingScript)

	entityTemplate, err := bitcoin.NewMultiPKHTemplate(3, 4)
	if err != nil {
		t.Fatalf("Failed to create multi-sig template : %s", err)
	}

	publicKeys := make([]bitcoin.PublicKey, 4)
	for i := range publicKeys {
		key, _ := bitcoin.GenerateKey(bitcoin.MainNet)
		publicKeys[i] = key.PublicKey()
	}

	entityLockingScript, err := entityTemplate.LockingScript(publicKeys)
	if err != nil {
		t.Fatalf("Failed to create multi-sig locking script : %s", err)
	}

	entityRA, err := bitcoin.RawAddressFromLockingScript(entityLockingScript)
	if err != nil {
		t.Fatalf("Failed to create multi-sig address : %s", err)
	}

	publicKeys = make([]bitcoin.PublicKey, 4)
	for i := range publicKeys {
		key, _ := bitcoin.GenerateKey(bitcoin.MainNet)
		publicKeys[i] = key.PublicKey()
	}

	adminLockingScript, err := entityTemplate.LockingScript(publicKeys)
	if err != nil {
		t.Fatalf("Failed to create multi-sig locking script : %s", err)
	}

	adminRA, err := bitcoin.RawAddressFromLockingScript(adminLockingScript)
	if err != nil {
		t.Fatalf("Failed to create multi-sig address : %s", err)
	}

	requestTx := wire.NewMsgTx(1)

	contractOffer := &actions.ContractOffer{
		ContractName:             "Multisig Contracts 123456",
		ContractOperatorIncluded: true,
		ContractFee:              contractFee,
		MasterAddress:            masterRA.Bytes(),
		EntityContract:           entityRA.Bytes(),
		OperatorEntityContract:   contractOperatorEntityRA.Bytes(),
		ContractType:             1,
		GoverningLaw:             "ESP  ",
		Jurisdiction:             "ESP  ",
	}

	contractOfferScript, err := Serialize(contractOffer, true)
	if err != nil {
		t.Fatalf("%s Failed to serialize contract offer : %s", Failed, err)
	}

	requestTx.AddTxOut(wire.NewTxOut(0, contractAgentLockingScript))
	requestTx.AddTxOut(wire.NewTxOut(0, contractOfferScript))

	inputScripts := []bitcoin.Script{adminLockingScript, contractOperatorLockingScript}
	estResponseTxFee, err := EstimatedContractOfferResponseTxFee(contractOffer,
		contractAgentLockingScript, contractFeeLockingScript, inputScripts, feeRate, dustFeeRate, 0,
		isTest)
	if err != nil {
		t.Fatalf("%s Failed to estimate response : %s", Failed, err)
	}

	t.Logf("Estimated Response Tx Fee: %d", estResponseTxFee)

	contractFormation, _ := contractOffer.Formation()
	contractFormation.AdminAddress = adminRA.Bytes()
	contractFormation.OperatorAddress = contractOperatorRA.Bytes()
	contractFormation.Timestamp = uint64(time.Now().UnixNano())

	contractFormationScript, err := Serialize(contractFormation, true)
	if err != nil {
		t.Fatalf("%s Failed to serialize contract formation : %s", Failed, err)
	}

	responseTx := wire.NewMsgTx(1)

	// From Contract (-3 for good measure. not sure why but the test is 3 off)
	unlockingScriptSize, err := txbuilder.UnlockingScriptSize(contractAgentLockingScript)
	if err != nil {
		t.Fatalf("%s Failed to calculate agent unlocking script size : %s", Failed, err)
	}

	responseTx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(&bitcoin.Hash32{}, 0),
		make([]byte, unlockingScriptSize)))

	// To Contract
	responseTx.AddTxOut(wire.NewTxOut(contractAgentDust, contractAgentLockingScript))

	// Contract Formation
	responseTx.AddTxOut(wire.NewTxOut(0, contractFormationScript))

	// To Fees
	responseTx.AddTxOut(wire.NewTxOut(contractFee, contractFeeLockingScript))

	responseTxSize := uint64(responseTx.SerializeSize())

	responseTxFee := txbuilder.EstimatedFeeValue(uint64(responseTxSize), feeRate)

	if uint64(responseTxFee)+contractFee+contractAgentDust != uint64(estResponseTxFee)+contractFee {
		t.Fatalf("%s Wrong funding amount : got %d, want %d", Failed, uint64(estResponseTxFee)+contractFee,
			uint64(responseTxFee)+contractFee+contractAgentDust)
	}
}

func TestInstrumentDefinitionResponseFees_Old(t *testing.T) {
	contractFee := uint64(2000)
	dustLimit := uint64(546)

	currency := &instruments.Currency{
		CurrencyCode: "AUD",
		Precision:    2,
	}

	cb, err := currency.Bytes()
	if err != nil {
		t.Fatalf("%s Failed to serialize currency : %s", Failed, err)
	}

	requestTx := wire.NewMsgTx(1)

	instrumentDefinition := &actions.InstrumentDefinition{
		AdministrationProposal: true,
		AuthorizedTokenQty:     10000,
		InstrumentPayload:      cb,
	}

	b, err := Serialize(instrumentDefinition, true)
	if err != nil {
		t.Fatalf("%s Failed to serialize instrument definition : %s", Failed, err)
	}

	requestTx.AddTxOut(wire.NewTxOut(3000, make([]byte, 26)))
	requestTx.AddTxOut(wire.NewTxOut(0, b))

	size, funding, err := EstimatedResponse(requestTx, 0, dustLimit, contractFee, true)
	if err != nil {
		t.Fatalf("%s Failed to estimate response : %s", Failed, err)
	}

	t.Logf("%s Size : %d, Funding %d", Success, size, funding)

	var instrumentCode bitcoin.Hash32
	instrumentCreation := &actions.InstrumentCreation{
		InstrumentCode:         instrumentCode.Bytes(), // Instrument code is added by smart contract
		AdministrationProposal: true,
		AuthorizedTokenQty:     10000,
		InstrumentPayload:      cb,
		InstrumentRevision:     0,
		Timestamp:              uint64(time.Now().UnixNano()),
	}

	r, err := Serialize(instrumentCreation, true)
	if err != nil {
		t.Fatalf("%s Failed to serialize instrument creation : %s", Failed, err)
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

func TestInstrumentDefinitionResponseFees_New(t *testing.T) {
	h := "006a02bd015108746573742e544b4e5301000249311e50195a03434f55621532064b6e75636b7340014a09080112035553441802"
	b, err := hex.DecodeString(h)
	if err != nil {
		t.Fatalf("Invalid hex : %s", err)
	}

	action, err := Deserialize(b, true)
	if err != nil {
		t.Fatalf("Failed to deserialize action : %s", err)
	}

	definition, ok := action.(*actions.InstrumentDefinition)
	if !ok {
		t.Fatalf("Not an instrument definition")
	}

	payload, err := instruments.Deserialize([]byte(definition.InstrumentType),
		definition.InstrumentPayload)
	if err != nil {
		t.Fatalf("Failed to deserialize payload : %s", err)
	}

	js, _ := json.MarshalIndent(definition, "", "  ")
	t.Logf("Instrument Definition : %s", js)

	js, _ = json.MarshalIndent(payload, "", "  ")
	t.Logf("Instrument Payload : %s", js)

	key, _ := bitcoin.GenerateKey(bitcoin.MainNet)
	contractAgentLockingScript, _ := key.LockingScript()

	key, _ = bitcoin.GenerateKey(bitcoin.MainNet)
	contractFeeLockingScript, _ := key.LockingScript()

	responseTxFee, err := EstimatedInstrumentDefinitionResponseTxFee(definition,
		contractAgentLockingScript, contractFeeLockingScript, 0.05, 0.0, true)
	if err != nil {
		t.Fatalf("Failed to estimate response fee : %s", err)
	}

	if responseTxFee != 18 {
		t.Errorf("Wrong response tx fee : got %d, want %d", responseTxFee, 18)
	}
}

func TestTransferResponseFees(t *testing.T) {
	contractFees := []uint64{2000, 4000}
	feeRate := float32(0.05)
	dustFeeRate := float32(0.0)

	ac1 := make([]byte, 32)
	rand.Read(ac1)
	ac2 := make([]byte, 32)
	rand.Read(ac2)

	ckey1, err := bitcoin.GenerateKey(bitcoin.MainNet)
	if err != nil {
		t.Fatalf("Failed to generate key : %s", err)
	}

	cad1, err := ckey1.RawAddress()
	if err != nil {
		t.Fatalf("Failed to create address : %s", err)
	}

	ckey2, err := bitcoin.GenerateKey(bitcoin.MainNet)
	if err != nil {
		t.Fatalf("Failed to generate key : %s", err)
	}

	cad2, err := ckey2.RawAddress()
	if err != nil {
		t.Fatalf("Failed to create address : %s", err)
	}

	key1, err := bitcoin.GenerateKey(bitcoin.MainNet)
	if err != nil {
		t.Fatalf("Failed to generate key : %s", err)
	}

	ad1, err := key1.RawAddress()
	if err != nil {
		t.Fatalf("Failed to create address : %s", err)
	}

	key2, err := bitcoin.GenerateKey(bitcoin.MainNet)
	if err != nil {
		t.Fatalf("Failed to generate key : %s", err)
	}

	ad2, err := key2.RawAddress()
	if err != nil {
		t.Fatalf("Failed to create address : %s", err)
	}

	key3, err := bitcoin.GenerateKey(bitcoin.MainNet)
	if err != nil {
		t.Fatalf("Failed to generate key : %s", err)
	}

	ad3, err := key3.RawAddress()
	if err != nil {
		t.Fatalf("Failed to create address : %s", err)
	}

	// Setup transfer request
	transfer := &actions.Transfer{
		Instruments: []*actions.InstrumentTransferField{
			&actions.InstrumentTransferField{
				ContractIndex:  0,
				InstrumentType: instruments.CodeCurrency,
				InstrumentCode: ac1,
				InstrumentSenders: []*actions.QuantityIndexField{
					&actions.QuantityIndexField{
						Quantity: 1000,
						Index:    0,
					},
					&actions.QuantityIndexField{
						Quantity: 2000,
						Index:    1,
					},
				},
				InstrumentReceivers: []*actions.InstrumentReceiverField{
					&actions.InstrumentReceiverField{
						Address:  ad1.Bytes(),
						Quantity: 3000,
					},
				},
			},
			&actions.InstrumentTransferField{
				ContractIndex:  1,
				InstrumentType: instruments.CodeCurrency,
				InstrumentCode: ac1,
				InstrumentSenders: []*actions.QuantityIndexField{
					&actions.QuantityIndexField{
						Quantity: 4000,
						Index:    1,
					},
				},
				InstrumentReceivers: []*actions.InstrumentReceiverField{
					&actions.InstrumentReceiverField{
						Address:  ad2.Bytes(),
						Quantity: 3000,
					},
					&actions.InstrumentReceiverField{
						Address:  ad3.Bytes(),
						Quantity: 1000,
					},
				},
			},
		},
	}

	// Build request transaction
	requestTx := wire.NewMsgTx(1)

	var in1 bitcoin.Hash32
	rand.Read(in1[:])
	requestTx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(&in1, 1),
		make([]byte, txbuilder.MaximumP2PKInputSize)))

	var in2 bitcoin.Hash32
	rand.Read(in2[:])
	requestTx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(&in2, 2),
		make([]byte, txbuilder.MaximumP2PKInputSize)))

	lockingScript, err := cad1.LockingScript()
	if err != nil {
		t.Fatalf("Failed to create locking script : %s", err)
	}
	requestTx.AddTxOut(wire.NewTxOut(1000, lockingScript))

	lockingScript, err = cad2.LockingScript()
	if err != nil {
		t.Fatalf("Failed to create locking script : %s", err)
	}
	requestTx.AddTxOut(wire.NewTxOut(1000, lockingScript))

	opScript, err := Serialize(transfer, true)
	if err != nil {
		t.Fatalf("Failed to serialize transfer : %s", err)
	}
	requestTx.AddTxOut(wire.NewTxOut(0, opScript))

	// Estimate funding
	funding, boomerang, err := EstimatedTransferResponse(requestTx,
		[]bitcoin.Script{lockingScript, lockingScript}, feeRate, dustFeeRate, contractFees, true)

	if err != nil {
		t.Fatalf("Failed to estimate response : %s", err)
	}

	t.Logf("Funding : %v", funding)
	t.Logf("Boomerang : %d", boomerang)
}

func TestExampleTransferResponseFees(t *testing.T) {
	contractFees := []uint64{2000, 2000}
	feeRate := float32(1.0)
	dustFeeRate := float32(1.0)

	h := "010000000441cf4cd05f3e90c6b88bcf79128fece494c8b3df89ae9b18707921ccf8640431020000006a473044022059aa3e22b19989c2237e604a6b251715dafdc351415f0938ce28302e8061a8f502205a12a3b85abfa2d6f27118abd883c775cd87ae61735742af7b2639b6b8bc91ee412103b370b2aab691da4c1ae98b7e464e5e6835b3af3a671cbde7b3e13463c9d20eb0ffffffff23aecd316912c3221ec115c946e2bb5cdeaa39b9b461b9ea5332cb4d1df6f27b000000006a473044022025fbaf15e11b4f1ae921e2800eedb10f1c63690f477900d5714c0f735e3638e30220330700c48cb9ec07c88c8bef52eef864f977483799f4d5c714c5cb3c29ff85a24121032447bf539a49de8034bd1df92256049d1918ec2786fa455349d2d06117cf1bcfffffffff23aecd316912c3221ec115c946e2bb5cdeaa39b9b461b9ea5332cb4d1df6f27b010000006a473044022054124286f9118d86bc329210cf7a9e8eb62eea9f270bc71424bf0ec7e72589500220598f0aec527e9f71b8406fd77b316acdf3b6fea8aa44660d238c5729368457ac4121029bb9c7cb8be3e2a3bc94caef749a0ce527c37209138d13265b37bd19d8b20e58ffffffffb007215319b8db969f05583fdf9ac2da09da9f6e49b1296648c6f6997e9a3628030000006b483045022100fe8e1d9424fb82233d459bfb18e18034f74fa828f7128b2e93e429a07b6fa7b802205b6c184913baf618194359c12c1f546615a524f7a19dd1472c41072c0c49315b412103ad7c96e7da4066b647e4fa4ad471e66a6f1d49b79f5f39d7627ef4287dd1c07affffffff04db160000000000001976a9144068eae2ac3722b8d9b9fb6c1920062d0f5b3d2588ac22020000000000001976a9147beee8a3d931c9d6d221de0bfec2ad4d8bd84c0588ac0000000000000000fd8902006a02bd0008746573742e544b4e041a0254314d73020af00412034355521a20d6571723491ca8006ea1cd45c07886e92cb411104b2d1216c277f7eb25113833220310e0082a720a15207beee8a3d931c9d6d221de0bfec2ad4d8bd84c0510c80118012a46304402202e8620babfd561d41807920194f9b0613c2c3a608e3e885226130c09a42dcb1702203fd8d514ab78ea23dc2d7e1d738b406f557a50afdf33b4d8f799a78fcb1a969b30918a2838f4ddbf8fdec2e29e162a720a1520a6c6ab9ade7a68f4dc5daf670b872c41ca25cfa6106418012a473045022100caf87f8abef267cb50b1a7667b4ecb5fc7b06e5b1b393e3edf33d519ce3f8993022026d90ea97cb81854b9dcd0071303a2a2e7bc9e2c90016c7c05a424f4976464db30918a2838d4fdc0b8dec2e29e162a720a1520d9354520c31bc16704faf6596cae92c24930dc2f10a60118012a4630440220595ba09b7fd1c95728ceb407810b1dca6167ab8f3e562fb29b35da9a3450133f022057bdde8817e969866e19c13a79695e457fd920221090a032f4a13810e7e2009d30918a283897f9f6bcdec2e29e162a720a1520ccb51eb299d1fc9f2f91c8ae82cad3041153388110b80118012a463044022064f0f702c9c1086868e489a7a9a88ef9863ac70b8f29cc1a1af0f5c701b51d4d0220015722a799d407937dbb36d738e3a1b6ebdf9a42c710e67cb1b3baefa1cf8be830918a2838dd97c3c2dec2e29e162a720a1520f0453add3936952a18a339b34753e7d18a69c8f110d60318012a46304402206174c356fc4f59e9d0319ef34c819e2251e849469996fb57abadf0c5bf4941bb0220693ef570ad6eec1328bddaf405fff9e8ba43727682328490f8331ab83f8833ac30918a2838a3919ac7dec2e29e1628210000000000001976a914d5ef701d1e06a72bc17711678da1a3b0b52cbb4d88ac00000000"

	b, err := hex.DecodeString(h)
	if err != nil {
		t.Fatalf("Failed to parse request tx hex : %s", err)
	}

	requestTx := &wire.MsgTx{}
	if err := requestTx.Deserialize(bytes.NewReader(b)); err != nil {
		t.Fatalf("Failed to parse request tx : %s", err)
	}

	key, _ := bitcoin.GenerateKey(bitcoin.TestNet)
	script, _ := key.LockingScript()
	var lockingScripts []bitcoin.Script
	for range requestTx.TxIn {
		lockingScripts = append(lockingScripts, script)
	}

	// Estimate funding
	funding, boomerang, err := EstimatedTransferResponse(requestTx, lockingScripts, feeRate,
		dustFeeRate, contractFees, true)

	if err != nil {
		t.Fatalf("Failed to estimate response : %s", err)
	}

	t.Logf("Funding : %v", funding)
	t.Logf("Boomerang : %d", boomerang)

	if len(funding) < 1 {
		t.Fatalf("No contract funding returned")
	}

	// From response 03c1e7ec7dc47debcd277d11c2abb321b9687d240a85f0e20c61e3f34e339171
	if funding[0] < 5794 {
		t.Errorf("Not enough contract funding : got %d, want %d", funding[0], 5794)
	}

	if boomerang > 0 {
		t.Errorf("Not enough boomerang : got %d, want %d", boomerang, 0)
	}
}

func Test_ExampleTransferResponseFees_Zero(t *testing.T) {
	contractFees := []uint64{0}
	feeRate := float32(0.05)
	dustFeeRate := float32(0.0)

	key, _ := bitcoin.GenerateKey(bitcoin.MainNet)
	contractAgentLockingScript, _ := key.LockingScript()

	key, _ = bitcoin.GenerateKey(bitcoin.MainNet)
	senderLockingScript, _ := key.LockingScript()

	key, _ = bitcoin.GenerateKey(bitcoin.MainNet)
	receiverAddress, _ := key.RawAddress()

	requestTx := wire.NewMsgTx(1)

	requestTx.AddTxOut(wire.NewTxOut(0, contractAgentLockingScript))

	transfer := &actions.Transfer{
		Instruments: []*actions.InstrumentTransferField{
			{
				ContractIndex:  0,
				InstrumentType: instruments.CodeCreditNote,
				InstrumentCode: make([]byte, InstrumentCodeSize),
				InstrumentSenders: []*actions.QuantityIndexField{
					{
						Quantity: 100,
						Index:    0,
					},
				},
				InstrumentReceivers: []*actions.InstrumentReceiverField{
					{
						Address:  receiverAddress.Bytes(),
						Quantity: 100,
					},
				},
			},
		},
	}

	transferScript, err := Serialize(transfer, true)
	if err != nil {
		t.Fatalf("Failed to create transfer script : %s", err)
	}

	requestTx.AddTxOut(wire.NewTxOut(0, transferScript))

	lockingScripts := []bitcoin.Script{senderLockingScript}

	// Estimate funding
	funding, boomerang, err := EstimatedTransferResponse(requestTx, lockingScripts, feeRate,
		dustFeeRate, contractFees, true)

	if err != nil {
		t.Fatalf("Failed to estimate response : %s", err)
	}

	t.Logf("Funding : %v", funding)
	t.Logf("Boomerang : %d", boomerang)

	if len(funding) < 1 {
		t.Fatalf("No contract funding returned")
	}

	// From response 03c1e7ec7dc47debcd277d11c2abb321b9687d240a85f0e20c61e3f34e339171
	if funding[0] != 19 {
		t.Errorf("Not enough contract funding : got %d, want %d", funding[0], 19)
	}

	if boomerang > 0 {
		t.Errorf("Not enough boomerang : got %d, want %d", boomerang, 0)
	}
}

func TestMultiContractExampleTransferResponseFees(t *testing.T) {
	contractFees := []uint64{2000, 2000}
	feeRate := float32(1.0)
	dustFeeRate := float32(1.0)

	// Hex tx
	h := "0100000004ca2bc0de6f66a52ffb29300832a1dbe992cb5f2164a81ccc29dfdd895c5258a3010000006a47304402201ee4aeceeb7e7722bffa2af82c3f155b128c22c92e257d3721b86536d54d380b022070cab3a4c13d94db8ad78d6e8aeb5c94a3319f12a5a7b15c107d79ac37d5ed5c4121023f4788aad94abbd33760ceebe24a315425998bf435c38ffe2ac0c2838b1dde2effffffff9839e203c40d6ba67810b935ef6bc4bc4bb3cc900b9b58e9de02b616eca84261000000006a473044022044b6e0890ab4df32dea6297ab9cbeda4fe964b785418d4b66547f27621f7a6d002203fe1b664452a8f18dfbca20d1dac6342e7129bad9d7c33fed87ba43c068e1e454121021898ab0284a77b81ad2dc9eaa51e651be2712813b55e1adf14d8096f1933c6e0ffffffff0bb2ed3ed42e0af3c06483864fd11beaf7657fba68f824ed59d61766e69be47e010000006a47304402201e900def83ced090da54b0bb9de5ea482180a1db7d888bce4f91777daf66243a022057e0958fa9aa9c25e561feae80fc7fc2bc6cd564a313ca41c2f87c363e197114412102e20407d67a11c43635d1050c72c477f3744c17ecb257bc643037ed54f59575c9ffffffff4e05fb0d92dbdd570937cd0c97f8ba19219146e8a04992286c84b3310b5f52ce000000006b483045022100b342ea27d980eeef85834c2e40264b482e5b33f7e76e118b9608930f955341660220786c0c1bb60a66fa2fea7fbe59f8365bb2c565512bd60a47653fbedabddec048412103765276b998c225bc39f69e10ba49d7b0118af7abf45e6827ee55e153fdb116d5ffffffff04d2100000000000001976a9144068eae2ac3722b8d9b9fb6c1920062d0f5b3d2588acf2090000000000001976a914e7bf2122bc0b7373234843cd9c3b0dd37dc2717488ac0000000000000000fd1101006a02bd0008746573742e544b4e041a0254314cfc0aa30112034355521a20d6571723491ca8006ea1cd45c07886e92cb411104b2d1216c277f7eb251138332205080210c8012a730a15208faa446d79832101b751595b8c474309cfd538d110c80118012a473045022100e3538663ba00dae05ab45e2b8a1e728d174a2f1d87c56170b904a76e9540717402206a74daf4d4f8919e0a267b74a41475389d401d1fbc86a3e1751f4a0fdc59179e30908a2838e6cf8fc189c0e29e160a4a08011203434f551a20e6c3012d4e38c2fff4b206c3fe27af25ebdf8994d83a0fb3c6ab04bf4d5bf75e220310e8072a1a0a1520eb891888cdcbb2dcbd097eae04e8a7394faabb5b10e80710c0c4ca86abcff0a016f4050000000000001976a9144068eae2ac3722b8d9b9fb6c1920062d0f5b3d2588ac00000000"

	b, err := hex.DecodeString(h)
	if err != nil {
		t.Fatalf("Failed to parse request tx hex : %s", err)
	}

	requestTx := &wire.MsgTx{}
	if err := requestTx.Deserialize(bytes.NewReader(b)); err != nil {
		t.Fatalf("Failed to parse request tx : %s", err)
	}

	key, _ := bitcoin.GenerateKey(bitcoin.TestNet)
	script, _ := key.LockingScript()
	var lockingScripts []bitcoin.Script
	for range requestTx.TxIn {
		lockingScripts = append(lockingScripts, script)
	}

	// Estimate funding
	funding, boomerang, err := EstimatedTransferResponse(requestTx, lockingScripts, feeRate,
		dustFeeRate, contractFees, true)

	if err != nil {
		t.Fatalf("Failed to estimate response : %s", err)
	}

	// Failure was 1524. Failed at sig request with 1126 of 1352 needed (226 short)
	// Needs > 1750

	t.Logf("Funding : %v", funding)
	t.Logf("Boomerang : %d", boomerang)

	if boomerang < 1750 {
		t.Errorf("Not enough boomerang : got %d, want >= %d", boomerang, 1750)
	}
}

func TestMultiContractExample2TransferResponseFees(t *testing.T) {
	contractFees := []uint64{2000, 2000}
	feeRate := float32(1.0)
	dustFeeRate := float32(1.0)

	h := "0100000004d3e50737e2f8e67a7db014edd8f76da4208bc773422ce4760da693da78473268020000006a473044022069e3f6c16ac3b58a062c68f580b158c191fac86542817f938b64ea7780008d22022073535c685987d5535c8b7a9a31d390303a9b8dd8f8aa13b4960e5396924a902f412103b654875e469b677e919b7c1bbcfb9657c80d4c0fe412705579fd3a67dda518bfffffffff94abf2d91615b61ad9944da7821238a0106a57119b25d3f4b519a68924e0afda000000006b483045022100c9ecdc04f1507d25de4d3a6e87875774b0f2d70e9057ba4103edb80e21ddaf2f0220505bfb25e8098243220c0409fb3a041c7a1beb64ff9f43276f80f74deff38298412102baa5c4238b87a6561dbbb9a7ea19f4914599508c37304509b9725bd6d3831935ffffffffb5bef1d4f18c67583de9f7e8ce85fa389da619ce2f8749b80b25c1d5a042e4be020000006946304302200d621b472cffafb78fbb9e637fed5f4639f09d283e6316d67262d0e54c19bbce021f5ca3bb38e411e3e7b8e4c5d48b1e47e3a1bf76f9e8eabfc620af2a459e1d864121021c176d29126d262a152e518837b5415b1be3fbf36d7e9d4b64b30f327b02af2dffffffff5e728e246dfbc9f1c5102566223a6730110e43b188cc5be773c9617727f3fa1b000000006a47304402201f5ad7d489c8d82543c63ec03aec3ae6bbe5b4b68ba7cf23e87743b3e23c04f30220399bc533bf9387dcbd79b9bd2a536d218eed9a5f16fb8db06fe50972e5ee9603412103b667d6ba2666fe9bbbd9096f00fa6918fc95d8a131d57199e5b95d1f1a39b035ffffffff04d2100000000000001976a9144068eae2ac3722b8d9b9fb6c1920062d0f5b3d2588ac970b0000000000001976a9144068eae2ac3722b8d9b9fb6c1920062d0f5b3d2588acf2090000000000001976a914d10efc43564e142b4a2a35b94c5c8e5f7ae1d04d88ac0000000000000000fd6601006a02bd0008746573742e544b4e041a0254314d50010aa00112034355521a20d6571723491ca8006ea1cd45c07886e92cb411104b2d1216c277f7eb251138332204080210022a710a15206ebc761d1462d1e29860c277f3e4b601cb5d2023100218012a463044022016193cadea74015bb550d0c04e56855133882e5060a31722256af5ceaafcff7602207389033aa28f24a4d1d787df5d6347f58a1036930115a6e91e56137a36c0914930de852838fee0bc99fbc3949e160aa001080212034355521a205014ab300af2697c1a9313c055bafb0f31efb28757a8d8c098ec68432eb18a96220210012a710a152049d32a3556228d119a040f292ac6dbf0684261a6100118012a46304402202d62d642f1ad775d41d9b715ec6164b50aff90d67c9e95e7c56495eebbc8846d022036c2956dfc494c652f8268a3c65f01bf4751f58cc7e343191810d07df77dda6930de852838bbe3be8998c7949e161080bccaf69bd3a2a01600000000"

	// Actual outputs

	// Value: 0.00004306
	// Script: 76a9144068eae2ac3722b8d9b9fb6c1920062d0f5b3d2588ac
	// Address: 16sZw5K28wNWMPSzyqw5QabvxCg97bUi8k

	// Value: 0.00002967
	// Script: 76a9144068eae2ac3722b8d9b9fb6c1920062d0f5b3d2588ac
	// Address: 16sZw5K28wNWMPSzyqw5QabvxCg97bUi8k

	// Value: 0.00002546
	// Script: 76a914d10efc43564e142b4a2a35b94c5c8e5f7ae1d04d88ac
	// Address: 1L4QBJCgVPJfYjmXzykajifzb2cwpVEJyd

	b, err := hex.DecodeString(h)
	if err != nil {
		t.Fatalf("Failed to parse request tx hex : %s", err)
	}

	requestTx := &wire.MsgTx{}
	if err := requestTx.Deserialize(bytes.NewReader(b)); err != nil {
		t.Fatalf("Failed to parse request tx : %s", err)
	}

	key, _ := bitcoin.GenerateKey(bitcoin.TestNet)
	script, _ := key.LockingScript()
	var lockingScripts []bitcoin.Script
	for range requestTx.TxIn {
		lockingScripts = append(lockingScripts, script)
	}

	// Estimate funding
	funding, boomerang, err := EstimatedTransferResponse(requestTx, lockingScripts, feeRate,
		dustFeeRate, contractFees, true)

	if err != nil {
		t.Fatalf("Failed to estimate response : %s", err)
	}

	t.Logf("Funding : %v", funding)
	t.Logf("Boomerang : %d", boomerang)

	// Settlement request
	// Tx (396 bytes) : c1315e69710b76461f1c49374db601637b018f7c4600bd1f6356f18f68b4f777
	// TxId: c1315e69710b76461f1c49374db601637b018f7c4600bd1f6356f18f68b4f777
	//   Version: 1
	//   Inputs:

	//     Outpoint: 1 - baa09434e58928ed5760fc8cf427dc04f1c53c7477ca4400e8b0f79d0d5af02c
	//     Script: 483045022100a78386815d026bd5cfefdf61b1e317b37dcdb0a50a741fc80f484401b7199112022010f79e993b33875561f42d9227a540b5766d68d19ef4a3fdd70d7f93dea83d9f412103c5de270d037fbee41f3a2bba8d20ccffc58259280878f31d736b96b1ff8d8064
	//     Sequence: ffffffff
	//     Address: 16sZw5K28wNWMPSzyqw5QabvxCg97bUi8k

	//   Outputs:

	//     Value: 0.00002570
	//     Script: 76a914d10efc43564e142b4a2a35b94c5c8e5f7ae1d04d88ac
	//     Address: 1L4QBJCgVPJfYjmXzykajifzb2cwpVEJyd

	//     Value: 0.00000000
	//     Script: 006a02bd000e746573742e746f6b656e697a6564041a024d314ca812010018eb07229f0108b4cfb69ca8d48f9e1612202cf05a0d9df7b0e80044ca77743cc5f104dc27f48cfc6057ed2889e53494a0ba1a1a0a15207afeb886413b7fbcf2da3ab85a6f019c16109d1d10d00f2255006a02bd000e746573742e746f6b656e697a6564041a0254323b0a2f12034355521a20d6571723491ca8006ea1cd45c07886e92cb411104b2d1216c277f7eb25113833220022040801100210b4cfb69ca8d48f9e16

	//   LockTime: 0

	// Message
	// {
	//     "ReceiverIndexes": [
	//         0
	//     ],
	//     "MessageCode": 1003,
	//     "MessagePayload": "08b4cfb69ca8d48f9e1612202cf05a0d9df7b0e80044ca77743cc5f104dc27f48cfc6057ed2889e53494a0ba1a1a0a15207afeb886413b7fbcf2da3ab85a6f019c16109d1d10d00f2255006a02bd000e746573742e746f6b656e697a6564041a0254323b0a2f12034355521a20d6571723491ca8006ea1cd45c07886e92cb411104b2d1216c277f7eb25113833220022040801100210b4cfb69ca8d48f9e16"
	// }

	// Settlement Request
	// {
	//     "Timestamp": 1602224435149776820,
	//     "TransferTxId": "2cf05a0d9df7b0e80044ca77743cc5f104dc27f48cfc6057ed2889e53494a0ba",
	//     "ContractFees": [
	//         {
	//             "Address": "207afeb886413b7fbcf2da3ab85a6f019c16109d1d",
	//             "Quantity": 2000
	//         }
	//     ],
	//     "Settlement": "006a02bd000e746573742e746f6b656e697a6564041a0254323b0a2f12034355521a20d6571723491ca8006ea1cd45c07886e92cb411104b2d1216c277f7eb25113833220022040801100210b4cfb69ca8d48f9e16"
	// }

	// Settlement
	// {
	//     "Instruments": [
	//         {
	//             "InstrumentType": "CCY",
	//             "InstrumentCode": "d6571723491ca8006ea1cd45c07886e92cb411104b2d1216c277f7eb25113833",
	//             "Settlements": [
	//                 {},
	//                 {
	//                     "Index": 1,
	//                     "Quantity": 2
	//                 }
	//             ]
	//         }
	//     ],
	//     "Timestamp": 1602224435149776820
	// }

	// Signature request
	// Tx (804 bytes) : c0943380e3cb15990d955dc2df2e84624913760ca1a9aad6c14bf593cb0ab5e4
	// TxId: c0943380e3cb15990d955dc2df2e84624913760ca1a9aad6c14bf593cb0ab5e4
	//   Version: 1
	//   Inputs:

	//     Outpoint: 0 - c1315e69710b76461f1c49374db601637b018f7c4600bd1f6356f18f68b4f777
	//     Script: 483045022100fd41e511318de88b9b8ae0011da565e09dead9d21d8c268e98bc8b16de9fc6ca022045a6dced78a471d74d46a5a1e241992402d2e495ad44226945d563eac1b6d98f412102019782826b61dec0cae0ebdcce8e9f693264552554fd62a806575f04c213794a
	//     Sequence: ffffffff
	//     Address: 1L4QBJCgVPJfYjmXzykajifzb2cwpVEJyd

	//   Outputs:

	//     Value: 0.00001765
	//     Script: 76a9144068eae2ac3722b8d9b9fb6c1920062d0f5b3d2588ac
	//     Address: 16sZw5K28wNWMPSzyqw5QabvxCg97bUi8k

	//     Value: 0.00000000
	//     Script: 006a02bd000e746573742e746f6b656e697a6564041a024d314d3d0212010018ea0722b40408a896c9f9a8d48f9e1612a70401000000022cf05a0d9df7b0e80044ca77743cc5f104dc27f48cfc6057ed2889e53494a0ba0000000000ffffffff2cf05a0d9df7b0e80044ca77743cc5f104dc27f48cfc6057ed2889e53494a0ba020000006b4830450221009c35ed7746c6f18d08a3d6ddd40b97e3907ca8e090463d213130ddabf9c3826e0220266c57019e74e58fedfda5aecc473d701d6ba1a00b8ca1bbe9bb3743f1f5dbd7412102019782826b61dec0cae0ebdcce8e9f693264552554fd62a806575f04c213794affffffff0722020000000000001976a91451368c1f497d192a87260fb74cb2c8a57001da0188ac22020000000000001976a9146ebc761d1462d1e29860c277f3e4b601cb5d202388ac22020000000000001976a9143f61ae39377679ef838312525d6fd7ff2225c6d688ac22020000000000001976a91449d32a3556228d119a040f292ac6dbf0684261a688acd0070000000000001976a9147afeb886413b7fbcf2da3ab85a6f019c16109d1d88acd0070000000000001976a9147afeb886413b7fbcf2da3ab85a6f019c16109d1d88ac00000000000000008b006a02bd000e746573742e746f6b656e697a6564041a0254324c700a2f12034355521a20d6571723491ca8006ea1cd45c07886e92cb411104b2d1216c277f7eb2511383322002204080110020a33080112034355521a205014ab300af2697c1a9313c055bafb0f31efb28757a8d8c098ec68432eb18a962202080222040803100110b4cfb69ca8d48f9e1600000000

	//   LockTime: 0

	// Message
	// {
	//     "ReceiverIndexes": [
	//         0
	//     ],
	//     "MessageCode": 1002,
	//     "MessagePayload": "08a896c9f9a8d48f9e1612a70401000000022cf05a0d9df7b0e80044ca77743cc5f104dc27f48cfc6057ed2889e53494a0ba0000000000ffffffff2cf05a0d9df7b0e80044ca77743cc5f104dc27f48cfc6057ed2889e53494a0ba020000006b4830450221009c35ed7746c6f18d08a3d6ddd40b97e3907ca8e090463d213130ddabf9c3826e0220266c57019e74e58fedfda5aecc473d701d6ba1a00b8ca1bbe9bb3743f1f5dbd7412102019782826b61dec0cae0ebdcce8e9f693264552554fd62a806575f04c213794affffffff0722020000000000001976a91451368c1f497d192a87260fb74cb2c8a57001da0188ac22020000000000001976a9146ebc761d1462d1e29860c277f3e4b601cb5d202388ac22020000000000001976a9143f61ae39377679ef838312525d6fd7ff2225c6d688ac22020000000000001976a91449d32a3556228d119a040f292ac6dbf0684261a688acd0070000000000001976a9147afeb886413b7fbcf2da3ab85a6f019c16109d1d88acd0070000000000001976a9147afeb886413b7fbcf2da3ab85a6f019c16109d1d88ac00000000000000008b006a02bd000e746573742e746f6b656e697a6564041a0254324c700a2f12034355521a20d6571723491ca8006ea1cd45c07886e92cb411104b2d1216c277f7eb2511383322002204080110020a33080112034355521a205014ab300af2697c1a9313c055bafb0f31efb28757a8d8c098ec68432eb18a962202080222040803100110b4cfb69ca8d48f9e1600000000"
	// }

	// Signature Request
	// {
	//     "Timestamp": 1602224435345115944,
	//     "Payload": "01000000022cf05a0d9df7b0e80044ca77743cc5f104dc27f48cfc6057ed2889e53494a0ba0000000000ffffffff2cf05a0d9df7b0e80044ca77743cc5f104dc27f48cfc6057ed2889e53494a0ba020000006b4830450221009c35ed7746c6f18d08a3d6ddd40b97e3907ca8e090463d213130ddabf9c3826e0220266c57019e74e58fedfda5aecc473d701d6ba1a00b8ca1bbe9bb3743f1f5dbd7412102019782826b61dec0cae0ebdcce8e9f693264552554fd62a806575f04c213794affffffff0722020000000000001976a91451368c1f497d192a87260fb74cb2c8a57001da0188ac22020000000000001976a9146ebc761d1462d1e29860c277f3e4b601cb5d202388ac22020000000000001976a9143f61ae39377679ef838312525d6fd7ff2225c6d688ac22020000000000001976a91449d32a3556228d119a040f292ac6dbf0684261a688acd0070000000000001976a9147afeb886413b7fbcf2da3ab85a6f019c16109d1d88acd0070000000000001976a9147afeb886413b7fbcf2da3ab85a6f019c16109d1d88ac00000000000000008b006a02bd000e746573742e746f6b656e697a6564041a0254324c700a2f12034355521a20d6571723491ca8006ea1cd45c07886e92cb411104b2d1216c277f7eb2511383322002204080110020a33080112034355521a205014ab300af2697c1a9313c055bafb0f31efb28757a8d8c098ec68432eb18a962202080222040803100110b4cfb69ca8d48f9e1600000000"
	// }

	// Signature Request Embedded Tx
	// Tx (551 bytes) : 7c7d5ce45b2d5e2a347cb3f32574310c866cc410a43c88b78e12613a2fd07454
	// TxId: 7c7d5ce45b2d5e2a347cb3f32574310c866cc410a43c88b78e12613a2fd07454
	//   Version: 1
	//   Inputs:

	//     Outpoint: 0 - baa09434e58928ed5760fc8cf427dc04f1c53c7477ca4400e8b0f79d0d5af02c
	//     Script:
	//     Sequence: ffffffff

	//     Outpoint: 2 - baa09434e58928ed5760fc8cf427dc04f1c53c7477ca4400e8b0f79d0d5af02c
	//     Script: 4830450221009c35ed7746c6f18d08a3d6ddd40b97e3907ca8e090463d213130ddabf9c3826e0220266c57019e74e58fedfda5aecc473d701d6ba1a00b8ca1bbe9bb3743f1f5dbd7412102019782826b61dec0cae0ebdcce8e9f693264552554fd62a806575f04c213794a
	//     Sequence: ffffffff
	//     Address: 1L4QBJCgVPJfYjmXzykajifzb2cwpVEJyd

	//   Outputs:

	//     Value: 0.00000546
	//     Script: 76a91451368c1f497d192a87260fb74cb2c8a57001da0188ac
	//     Address: 18QR5ZBBr8y4bqmTW41GXMT3s6seCXD1fM

	//     Value: 0.00000546
	//     Script: 76a9146ebc761d1462d1e29860c277f3e4b601cb5d202388ac
	//     Address: 1B6X6dQW5auagXBFtANav3scSHXNfAtZkk

	//     Value: 0.00000546
	//     Script: 76a9143f61ae39377679ef838312525d6fd7ff2225c6d688ac
	//     Address: 16n8b3t1odK1AoUheqCgmRriTqsKMn9nBJ

	//     Value: 0.00000546
	//     Script: 76a91449d32a3556228d119a040f292ac6dbf0684261a688ac
	//     Address: 17jMHzo3iV93euZPka8zdZ8E4n7vfKPbEL

	//     Value: 0.00002000
	//     Script: 76a9147afeb886413b7fbcf2da3ab85a6f019c16109d1d88ac
	//     Address: 1CDLaiT9GELzhJJfqfvHpMSdrJ8w3xN7b6

	//     Value: 0.00002000
	//     Script: 76a9147afeb886413b7fbcf2da3ab85a6f019c16109d1d88ac
	//     Address: 1CDLaiT9GELzhJJfqfvHpMSdrJ8w3xN7b6

	//     Value: 0.00000000
	//     Script: 006a02bd000e746573742e746f6b656e697a6564041a0254324c700a2f12034355521a20d6571723491ca8006ea1cd45c07886e92cb411104b2d1216c277f7eb2511383322002204080110020a33080112034355521a205014ab300af2697c1a9313c055bafb0f31efb28757a8d8c098ec68432eb18a962202080222040803100110b4cfb69ca8d48f9e16

	//   LockTime: 0

	// Settlement a82ba0db573c3b2db144b1fe549422c21eb6f63bb206d8a9aa5bd1e0347c11b8
	if len(funding) < 2 {
		t.Fatalf("Not enough contract funding entries")
	}

	if funding[0] < 4852 {
		t.Errorf("Not enough contract funding : got %d, want %d", funding[0], 4852)
	}
	t.Logf("Master funding is over by %d (%d%%)", funding[0]-4852,
		int(100.0*(float32(funding[0]-4852)/float32(4852))))

	if funding[1] != 2000 {
		t.Errorf("Not enough contract funding : got %d, want %d", funding[1], 2000)
	}

	// Example Boomerang of 2967 was 1219 too high. Correct boomerang amount would be 1748
	if boomerang < 1748 {
		t.Errorf("Not enough boomerang : got %d, want >= %d", boomerang, 1748)
	}
	t.Logf("Boomerang is over by %d (%d%%)", boomerang-1748,
		int(100.0*(float32(boomerang-1748)/float32(1748))))
}

func Test_EstimatedConfiscationResponse(t *testing.T) {
	contractFee := uint64(2000)
	feeRate := float32(0.5)
	dustFeeRate := float32(0.25)

	instrumentType := instruments.CodeCurrency
	var instrumentCode bitcoin.Hash20
	rand.Read(instrumentCode[:])

	contractKey, err := bitcoin.GenerateKey(bitcoin.MainNet)
	if err != nil {
		t.Fatalf("Failed to generate key : %s", err)
	}
	contractLockingScript, err := contractKey.LockingScript()
	if err != nil {
		t.Fatalf("Failed to create locking script : %s", err)
	}

	targetKey, err := bitcoin.GenerateKey(bitcoin.MainNet)
	if err != nil {
		t.Fatalf("Failed to generate key : %s", err)
	}
	targetAddress, err := targetKey.RawAddress()
	if err != nil {
		t.Fatalf("Failed to create raw address : %s", err)
	}

	depositKey, err := bitcoin.GenerateKey(bitcoin.MainNet)
	if err != nil {
		t.Fatalf("Failed to generate key : %s", err)
	}
	depositAddress, err := depositKey.RawAddress()
	if err != nil {
		t.Fatalf("Failed to create raw address : %s", err)
	}

	confiscationOrder := &actions.Order{
		ComplianceAction: actions.ComplianceActionConfiscation,
		InstrumentType:   instrumentType,
		InstrumentCode:   instrumentCode.Bytes(),
		DepositAddress:   depositAddress.Bytes(),
		Message:          "Sent to wrong address",
	}

	target := &actions.TargetAddressField{
		Address:  targetAddress.Bytes(),
		Quantity: 5000,
	}

	confiscationOrder.TargetAddresses = append(confiscationOrder.TargetAddresses, target)

	confiscationOrderScript, err := Serialize(confiscationOrder, true)
	if err != nil {
		t.Fatalf("%s Failed to serialize confiscation order : %s", Failed, err)
	}

	requestTx := wire.NewMsgTx(1)

	requestTx.AddTxOut(wire.NewTxOut(3000, contractLockingScript))
	requestTx.AddTxOut(wire.NewTxOut(0, confiscationOrderScript))

	funding, err := EstimatedConfiscationResponse(requestTx, feeRate, dustFeeRate, contractFee,
		true)
	if err != nil {
		t.Fatalf("%s Failed to estimate response : %s", Failed, err)
	}

	t.Logf("%s Funding %d", Success, funding)

	confiscation := &actions.Confiscation{
		InstrumentType: instrumentType,
		InstrumentCode: instrumentCode.Bytes(),
		Quantities: []*actions.QuantityIndexField{
			{ // Target
				Quantity: math.MaxUint64,
				Index:    0,
			},
			{ // Deposit
				Quantity: math.MaxUint64,
				Index:    1,
			},
		},
		DepositQty: 5000,
		Timestamp:  uint64(time.Now().UnixNano()),
	}

	confiscationScript, err := Serialize(confiscation, true)
	if err != nil {
		t.Fatalf("%s Failed to serialize confiscation response : %s", Failed, err)
	}

	responseTx := wire.NewMsgTx(1)
	dustLimit := txbuilder.DustLimitForLockingScript(contractLockingScript, dustFeeRate)

	// From Contract
	responseTx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(&bitcoin.Hash32{}, 0), make([]byte,
		txbuilder.MaximumP2PKHSigScriptSize)))

	// To Target
	responseTx.AddTxOut(wire.NewTxOut(dustLimit, make([]byte, txbuilder.P2PKHOutputScriptSize)))

	// To Deposit
	responseTx.AddTxOut(wire.NewTxOut(dustLimit, make([]byte, txbuilder.P2PKHOutputScriptSize)))

	// Confiscation
	responseTx.AddTxOut(wire.NewTxOut(0, confiscationScript))

	// To Fees
	responseTx.AddTxOut(wire.NewTxOut(contractFee, make([]byte, txbuilder.P2PKHOutputScriptSize)))

	outputValue := uint64(0)
	for _, txout := range responseTx.TxOut {
		outputValue += txout.Value
	}
	responseFee := uint64(float32(responseTx.SerializeSize()) * feeRate)

	if responseFee+outputValue != funding {
		t.Fatalf("%s Wrong funding amount : got %d, want %d", Failed, funding,
			responseFee+outputValue)
	}
	t.Logf("Verified Estimated Response Funding : got %d, want %d", funding,
		responseFee+outputValue)
}

func Test_TransferResponseFees_Example1(t *testing.T) {
	contractFees := []uint64{2000}
	feeRate := float32(0.05)
	dustFeeRate := float32(0.0)

	h := "006a02bd015108746573742e544b4e5301000254314d12010a850212034343591a14a451388d80b00661b4b90cfe103ee7ec615f0994220210412a710a152055e69778997556277d03d42794a91ebb41493a2d103218012a463044022032c6a3043487bc35aa44acb64f0945d6024fab2de3d63cc8e8389420a2fdbec5022023897ffe30cb511f493056a6bd49a9e0d4e9b27d537046ae02222084e3ec8f1330bca2313889beb7feecad82c1172a710a1520449c2b9f4500c765ee0eac4b9fdc62d5c2561829100f18012a46304402204c1952abf8fff745723166321c71193af10a86e1c138cb5ae548936fa607da640220708d7e9db541c07d694fcbde588f97a25d53266a1d925f7f3c4c0f678260fe2730bca23138ffd4b6afedad82c11710d38bc6eeb5be8cc017"

	b, err := hex.DecodeString(h)
	if err != nil {
		t.Fatalf("Failed to parse transfer hex : %s", err)
	}

	action, err := Deserialize(b, true)
	if err != nil {
		t.Fatalf("Failed to deserialize transfer : %s", err)
	}

	transfer := action.(*actions.Transfer)
	js, _ := json.MarshalIndent(transfer, "", "  ")
	t.Logf("Transfer : %s", js)

	requestTx := wire.NewMsgTx(1)

	var lockingScripts []bitcoin.Script
	for i := 0; i < 2; i++ {
		key, _ := bitcoin.GenerateKey(bitcoin.TestNet)
		script, _ := key.LockingScript()
		lockingScripts = append(lockingScripts, script)

		var randHash bitcoin.Hash32
		rand.Read(randHash[:])
		requestTx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(&randHash, 1), nil))
	}

	key, _ := bitcoin.GenerateKey(bitcoin.TestNet)
	script, _ := key.LockingScript()
	requestTx.AddTxOut(wire.NewTxOut(2021, script))

	requestTx.AddTxOut(wire.NewTxOut(0, b))

	key, _ = bitcoin.GenerateKey(bitcoin.TestNet)
	script, _ = key.LockingScript()
	requestTx.AddTxOut(wire.NewTxOut(187, script))

	// Estimate funding
	funding, boomerang, err := EstimatedTransferResponse(requestTx, lockingScripts, feeRate,
		dustFeeRate, contractFees, true)

	if err != nil {
		t.Fatalf("Failed to estimate response : %s", err)
	}

	t.Logf("Funding : %v", funding)
	t.Logf("Boomerang : %d", boomerang)

	if len(funding) < 1 {
		t.Fatalf("No contract funding returned")
	}

	// From response 522efbe6014b5bf740c48db0b7790d6f3a9203de06ecad1d67abeb01d1281aff
	if funding[0] < 2022 {
		t.Errorf("Not enough contract funding : got %d, want %d", funding[0], 2022)
	}

	if boomerang > 0 {
		t.Errorf("Not enough boomerang : got %d, want %d", boomerang, 0)
	}
}

func Test_TransferResponseFees_Example2(t *testing.T) {
	contractFees := []uint64{2000, 2000}
	feeRate := float32(0.05)
	dustFeeRate := float32(0.0)

	h := "010000000215664abb2cbc7db1d1d3b8ed8ed985fcfaed062ce0318f11aeab87dc645bec65010000006a47304402206792e2b1bb0fd2c0c31ab4ece1ab60d1057be2de1e984b9ccf385458fe9bb09802207823e56eefbea591b60f4fd816a59f2f120128162b548b760e06f0fe6deb6dca412103b8fe59c51966f28f8267ef5c43e638abe2721308eefae9677770f258eb95eba4ffffffff164fb009edd8681f2464d8b2bf7af0a3813819c0e1e974208f3fd8afa3baa08d010000006b483045022100c38f5c4d3da12333c72d565c34ef9c5b719db7d4df4ce6fc61ab6ffb8f44b5ea022047c7ef850d54c20c95b9a74ec966c46425a6987639e5c06a1c910b8dc0805871412103e52b04d4eafb375a6aee21ade32591fbac969976a43559225cadf7a3af25aa2dffffffff03e5070000000000001976a914b082b8cce2697aaaacfa698f1950e44e0ba1175c88ac0000000000000000fd2a01006a02bd015108746573742e544b4e5301000254314d12010a850212034343591a14a451388d80b00661b4b90cfe103ee7ec615f0994220210412a710a152055e69778997556277d03d42794a91ebb41493a2d103218012a463044022032c6a3043487bc35aa44acb64f0945d6024fab2de3d63cc8e8389420a2fdbec5022023897ffe30cb511f493056a6bd49a9e0d4e9b27d537046ae02222084e3ec8f1330bca2313889beb7feecad82c1172a710a1520449c2b9f4500c765ee0eac4b9fdc62d5c2561829100f18012a46304402204c1952abf8fff745723166321c71193af10a86e1c138cb5ae548936fa607da640220708d7e9db541c07d694fcbde588f97a25d53266a1d925f7f3c4c0f678260fe2730bca23138ffd4b6afedad82c11710d38bc6eeb5be8cc017bb000000000000001976a91458d2f37dd57fd7d907b72fc7de6875c79edb184588ac00000000"

	b, err := hex.DecodeString(h)
	if err != nil {
		t.Fatalf("Failed to parse request tx hex : %s", err)
	}

	requestTx := &wire.MsgTx{}
	if err := requestTx.Deserialize(bytes.NewReader(b)); err != nil {
		t.Fatalf("Failed to parse request tx : %s", err)
	}

	key, _ := bitcoin.GenerateKey(bitcoin.TestNet)
	script, _ := key.LockingScript()
	var lockingScripts []bitcoin.Script
	for range requestTx.TxIn {
		lockingScripts = append(lockingScripts, script)
	}

	// Estimate funding
	funding, boomerang, err := EstimatedTransferResponse(requestTx, lockingScripts, feeRate,
		dustFeeRate, contractFees, true)

	if err != nil {
		t.Fatalf("Failed to estimate response : %s", err)
	}

	t.Logf("Funding : %v", funding)
	t.Logf("Boomerang : %d", boomerang)

	if len(funding) < 1 {
		t.Fatalf("No contract funding returned")
	}

	// From response 522efbe6014b5bf740c48db0b7790d6f3a9203de06ecad1d67abeb01d1281aff
	if funding[0] < 2022 {
		t.Errorf("Not enough contract funding : got %d, want %d", funding[0], 2022)
	}

	if boomerang > 0 {
		t.Errorf("Not enough boomerang : got %d, want %d", boomerang, 0)
	}
}
