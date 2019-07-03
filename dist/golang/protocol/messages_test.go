package protocol

import (
	"bytes"
	"reflect"
	"testing"
)

func TestPublicMessage(t *testing.T) {
	// Create a randomized object
	initialMessage := PublicMessage{}
	// Version (uint)
	// uint test not setup

	// Timestamp (Timestamp)
	initialMessage.Timestamp = Timestamp{}

	// Subject (varchar)
	initialMessage.Subject = "Text 2"

	// Regarding (Output)
	initialMessage.Regarding = Output{}

	// PublicMessage (Document)
	initialMessage.PublicMessage = Document{}

	// Attachments (Document[])
	for i := 0; i < 2; i++ {
		initialMessage.Attachments = append(initialMessage.Attachments, Document{})
	}

	// Encode message
	initialEncoding, err := initialMessage.Serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := PublicMessage{}

	n, err := decodedMessage.Write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Version (uint)
	// uint test compare not setup

	// Timestamp (Timestamp)
	if initialMessage.Timestamp != decodedMessage.Timestamp {
		t.Errorf("Timestamp doesn't match : %v != %v", initialMessage.Timestamp, decodedMessage.Timestamp)
	}

	// Subject (varchar)
	if initialMessage.Subject != decodedMessage.Subject {
		t.Errorf("Subject doesn't match : %s != %s", initialMessage.Subject, decodedMessage.Subject)
	}

	// Regarding (Output)
	if initialMessage.Regarding != decodedMessage.Regarding {
		t.Errorf("Regarding doesn't match : %v != %v", initialMessage.Regarding, decodedMessage.Regarding)
	}

	// PublicMessage (Document)
	// if initialMessage.PublicMessage != decodedMessage.PublicMessage {
	// 	t.Errorf("PublicMessage doesn't match : %v != %v", initialMessage.PublicMessage, decodedMessage.PublicMessage)
	// }

	// Attachments (Document[])
	if len(initialMessage.Attachments) != len(decodedMessage.Attachments) {
		t.Errorf("Attachments lengths don't match : %d != %d", len(initialMessage.Attachments), len(decodedMessage.Attachments))
	}
}

func TestPrivateMessage(t *testing.T) {
	// Create a randomized object
	initialMessage := PrivateMessage{}
	// Version (uint)
	// uint test not setup

	// Timestamp (Timestamp)
	initialMessage.Timestamp = Timestamp{}

	// Subject (varchar)
	initialMessage.Subject = "Text 2"

	// Regarding (Output)
	initialMessage.Regarding = Output{}

	// PrivateMessage (Document)
	initialMessage.PrivateMessage = Document{}

	// Attachments (Document[])
	for i := 0; i < 2; i++ {
		initialMessage.Attachments = append(initialMessage.Attachments, Document{})
	}

	// Encode message
	initialEncoding, err := initialMessage.Serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := PrivateMessage{}

	n, err := decodedMessage.Write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Version (uint)
	// uint test compare not setup

	// Timestamp (Timestamp)
	if initialMessage.Timestamp != decodedMessage.Timestamp {
		t.Errorf("Timestamp doesn't match : %v != %v", initialMessage.Timestamp, decodedMessage.Timestamp)
	}

	// Subject (varchar)
	if initialMessage.Subject != decodedMessage.Subject {
		t.Errorf("Subject doesn't match : %s != %s", initialMessage.Subject, decodedMessage.Subject)
	}

	// Regarding (Output)
	if initialMessage.Regarding != decodedMessage.Regarding {
		t.Errorf("Regarding doesn't match : %v != %v", initialMessage.Regarding, decodedMessage.Regarding)
	}

	// PrivateMessage (Document)
	// if initialMessage.PrivateMessage != decodedMessage.PrivateMessage {
	// 	t.Errorf("PrivateMessage doesn't match : %v != %v", initialMessage.PrivateMessage, decodedMessage.PrivateMessage)
	// }

	// Attachments (Document[])
	if len(initialMessage.Attachments) != len(decodedMessage.Attachments) {
		t.Errorf("Attachments lengths don't match : %d != %d", len(initialMessage.Attachments), len(decodedMessage.Attachments))
	}
}

func TestRevertedTx(t *testing.T) {
	// Create a randomized object
	initialMessage := RevertedTx{}
	// Version (uint)
	// uint test not setup

	// Timestamp (Timestamp)
	initialMessage.Timestamp = Timestamp{}

	// Transaction (varbin)
	initialMessage.Transaction = make([]byte, 0, 32)
	for i := uint64(0); i < 32; i++ {
		initialMessage.Transaction = append(initialMessage.Transaction, byte(65+i+2))
	}

	// Encode message
	initialEncoding, err := initialMessage.Serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := RevertedTx{}

	n, err := decodedMessage.Write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Version (uint)
	// uint test compare not setup

	// Timestamp (Timestamp)
	if initialMessage.Timestamp != decodedMessage.Timestamp {
		t.Errorf("Timestamp doesn't match : %v != %v", initialMessage.Timestamp, decodedMessage.Timestamp)
	}

	// Transaction (varbin)
	if !bytes.Equal(initialMessage.Transaction, decodedMessage.Transaction) {
		t.Errorf("Transaction doesn't match : %x != %x", initialMessage.Transaction, decodedMessage.Transaction)
	}
}

func TestOffer(t *testing.T) {
	// Create a randomized object
	initialMessage := Offer{}
	// Version (uint)
	// uint test not setup

	// Timestamp (Timestamp)
	initialMessage.Timestamp = Timestamp{}

	// Payload (varbin)
	initialMessage.Payload = make([]byte, 0, 32)
	for i := uint64(0); i < 32; i++ {
		initialMessage.Payload = append(initialMessage.Payload, byte(65+i+2))
	}

	// Encode message
	initialEncoding, err := initialMessage.Serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Offer{}

	n, err := decodedMessage.Write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Version (uint)
	// uint test compare not setup

	// Timestamp (Timestamp)
	if initialMessage.Timestamp != decodedMessage.Timestamp {
		t.Errorf("Timestamp doesn't match : %v != %v", initialMessage.Timestamp, decodedMessage.Timestamp)
	}

	// Payload (varbin)
	if !bytes.Equal(initialMessage.Payload, decodedMessage.Payload) {
		t.Errorf("Payload doesn't match : %x != %x", initialMessage.Payload, decodedMessage.Payload)
	}
}

func TestSignatureRequest(t *testing.T) {
	// Create a randomized object
	initialMessage := SignatureRequest{}
	// Version (uint)
	// uint test not setup

	// Timestamp (Timestamp)
	initialMessage.Timestamp = Timestamp{}

	// Payload (varbin)
	initialMessage.Payload = make([]byte, 0, 32)
	for i := uint64(0); i < 32; i++ {
		initialMessage.Payload = append(initialMessage.Payload, byte(65+i+2))
	}

	// Encode message
	initialEncoding, err := initialMessage.Serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := SignatureRequest{}

	n, err := decodedMessage.Write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Version (uint)
	// uint test compare not setup

	// Timestamp (Timestamp)
	if initialMessage.Timestamp != decodedMessage.Timestamp {
		t.Errorf("Timestamp doesn't match : %v != %v", initialMessage.Timestamp, decodedMessage.Timestamp)
	}

	// Payload (varbin)
	if !bytes.Equal(initialMessage.Payload, decodedMessage.Payload) {
		t.Errorf("Payload doesn't match : %x != %x", initialMessage.Payload, decodedMessage.Payload)
	}
}

func TestSettlementRequest(t *testing.T) {
	// Create a randomized object
	initialMessage := SettlementRequest{}
	// Version (uint)
	// uint test not setup

	// Timestamp (Timestamp)
	initialMessage.Timestamp = Timestamp{}

	// TransferTxId (TxId)
	initialMessage.TransferTxId = TxId{}

	// ContractFees (TargetAddress[])
	for i := 0; i < 2; i++ {
		initialMessage.ContractFees = append(initialMessage.ContractFees, TargetAddress{})
	}

	// Settlement (varbin)
	initialMessage.Settlement = make([]byte, 0, 32)
	for i := uint64(0); i < 32; i++ {
		initialMessage.Settlement = append(initialMessage.Settlement, byte(65+i+4))
	}

	// Encode message
	initialEncoding, err := initialMessage.Serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := SettlementRequest{}

	n, err := decodedMessage.Write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Version (uint)
	// uint test compare not setup

	// Timestamp (Timestamp)
	if initialMessage.Timestamp != decodedMessage.Timestamp {
		t.Errorf("Timestamp doesn't match : %v != %v", initialMessage.Timestamp, decodedMessage.Timestamp)
	}

	// TransferTxId (TxId)
	if initialMessage.TransferTxId != decodedMessage.TransferTxId {
		t.Errorf("TransferTxId doesn't match : %v != %v", initialMessage.TransferTxId, decodedMessage.TransferTxId)
	}

	// ContractFees (TargetAddress[])
	if len(initialMessage.ContractFees) != len(decodedMessage.ContractFees) {
		t.Errorf("ContractFees lengths don't match : %d != %d", len(initialMessage.ContractFees), len(decodedMessage.ContractFees))
	}

	// Settlement (varbin)
	if !bytes.Equal(initialMessage.Settlement, decodedMessage.Settlement) {
		t.Errorf("Settlement doesn't match : %x != %x", initialMessage.Settlement, decodedMessage.Settlement)
	}
}

func TestOutputMetadata(t *testing.T) {
	// Create a randomized object
	initialMessage := OutputMetadata{}
	// Version (uint)
	// uint test not setup

	// OutputDescription (varchar)
	initialMessage.OutputDescription = "Text 1"

	// Tags (Tag[])
	for i := 0; i < 5; i++ {
		var item uint8
		initialMessage.Tags = append(initialMessage.Tags, item)
	}

	// CustomTags (OutputTag[])
	for i := 0; i < 2; i++ {
		initialMessage.CustomTags = append(initialMessage.CustomTags, OutputTag{})
	}

	// Encode message
	initialEncoding, err := initialMessage.Serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := OutputMetadata{}

	n, err := decodedMessage.Write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Version (uint)
	// uint test compare not setup

	// OutputDescription (varchar)
	if initialMessage.OutputDescription != decodedMessage.OutputDescription {
		t.Errorf("OutputDescription doesn't match : %s != %s", initialMessage.OutputDescription, decodedMessage.OutputDescription)
	}

	// Tags (Tag[])
	if len(initialMessage.Tags) != len(decodedMessage.Tags) {
		t.Errorf("Tags lengths don't match : %d != %d", len(initialMessage.Tags), len(decodedMessage.Tags))
	}
	for i, value := range initialMessage.Tags {
		if value != decodedMessage.Tags[i] {
			t.Errorf("Tags value %d doesn't match : %v != %v", i, value, decodedMessage.Tags[i])
		}
	}

	// CustomTags (OutputTag[])
	if len(initialMessage.CustomTags) != len(decodedMessage.CustomTags) {
		t.Errorf("CustomTags lengths don't match : %d != %d", len(initialMessage.CustomTags), len(decodedMessage.CustomTags))
	}
}
