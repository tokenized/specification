package protocol

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestPublicMessage(t *testing.T) {
	// Create a randomized object
	initialMessage := PublicMessage{}
	// Version (uint8)
	// uint8 test not setup

	// Timestamp (timestamp)
	initialMessage.Timestamp = uint64(time.Now().UnixNano())

	// PublicMessage (varchar)
	initialMessage.PublicMessage = "Text 2"

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
	// Version (uint8)
	// uint8 test compare not setup

	// Timestamp (timestamp)
	if initialMessage.Timestamp != decodedMessage.Timestamp {
		t.Errorf("Timestamp doesn't match : %d != %d", initialMessage.Timestamp, decodedMessage.Timestamp)
	}

	// PublicMessage (varchar)
	if initialMessage.PublicMessage != decodedMessage.PublicMessage {
		t.Errorf("PublicMessage doesn't match : %s != %s", initialMessage.PublicMessage, decodedMessage.PublicMessage)
	}
}

func TestPrivateMessage(t *testing.T) {
	// Create a randomized object
	initialMessage := PrivateMessage{}
	// Version (uint8)
	// uint8 test not setup

	// Timestamp (timestamp)
	initialMessage.Timestamp = uint64(time.Now().UnixNano())

	// PrivateMessage (varbin)
	initialMessage.PrivateMessage = make([]byte, 0, 32)
	for i := uint64(0); i < 32; i++ {
		initialMessage.PrivateMessage = append(initialMessage.PrivateMessage, byte(65+i+2))
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
	// Version (uint8)
	// uint8 test compare not setup

	// Timestamp (timestamp)
	if initialMessage.Timestamp != decodedMessage.Timestamp {
		t.Errorf("Timestamp doesn't match : %d != %d", initialMessage.Timestamp, decodedMessage.Timestamp)
	}

	// PrivateMessage (varbin)
	if !bytes.Equal(initialMessage.PrivateMessage, decodedMessage.PrivateMessage) {
		t.Errorf("PrivateMessage doesn't match : %x != %x", initialMessage.PrivateMessage, decodedMessage.PrivateMessage)
	}
}

func TestOffer(t *testing.T) {
	// Create a randomized object
	initialMessage := Offer{}
	// Version (uint8)
	// uint8 test not setup

	// Timestamp (timestamp)
	initialMessage.Timestamp = uint64(time.Now().UnixNano())

	// Offer (varbin)
	initialMessage.Offer = make([]byte, 0, 32)
	for i := uint64(0); i < 32; i++ {
		initialMessage.Offer = append(initialMessage.Offer, byte(65+i+2))
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
	// Version (uint8)
	// uint8 test compare not setup

	// Timestamp (timestamp)
	if initialMessage.Timestamp != decodedMessage.Timestamp {
		t.Errorf("Timestamp doesn't match : %d != %d", initialMessage.Timestamp, decodedMessage.Timestamp)
	}

	// Offer (varbin)
	if !bytes.Equal(initialMessage.Offer, decodedMessage.Offer) {
		t.Errorf("Offer doesn't match : %x != %x", initialMessage.Offer, decodedMessage.Offer)
	}
}

func TestSignatureRequest(t *testing.T) {
	// Create a randomized object
	initialMessage := SignatureRequest{}
	// Version (uint8)
	// uint8 test not setup

	// Timestamp (timestamp)
	initialMessage.Timestamp = uint64(time.Now().UnixNano())

	// SigRequest (varbin)
	initialMessage.SigRequest = make([]byte, 0, 32)
	for i := uint64(0); i < 32; i++ {
		initialMessage.SigRequest = append(initialMessage.SigRequest, byte(65+i+2))
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
	// Version (uint8)
	// uint8 test compare not setup

	// Timestamp (timestamp)
	if initialMessage.Timestamp != decodedMessage.Timestamp {
		t.Errorf("Timestamp doesn't match : %d != %d", initialMessage.Timestamp, decodedMessage.Timestamp)
	}

	// SigRequest (varbin)
	if !bytes.Equal(initialMessage.SigRequest, decodedMessage.SigRequest) {
		t.Errorf("SigRequest doesn't match : %x != %x", initialMessage.SigRequest, decodedMessage.SigRequest)
	}
}
