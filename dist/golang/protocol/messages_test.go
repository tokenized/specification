package protocol

import (
	"reflect"
	"testing"
)


func TestAssetDefinition(t *testing.T) {
	// The hex is the body of the message
	body := "00434f5561706d3271737a6e686b7332337a38643833753431733830313968797269336900000000000011bf01474252010301010100000000000f42404155443ba3d70a3c23d70a0009736f6d652064617461"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewAssetDefinition()

	header, err := NewHeaderForCode(CodeAssetDefinition, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestAssetCreation(t *testing.T) {
	// The hex is the body of the message
	body := "0053484361706d3271737a6e686b7332337a38643833753431733830313968797269336900000000000011bf01474252010301010100000000000f42404155443ba3d70a3c23d70a0009736f6d65206461746100000000000011bf1588f8cc6a16430d"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewAssetCreation()

	header, err := NewHeaderForCode(CodeAssetCreation, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestAssetModification(t *testing.T) {
	// The hex is the body of the message
	body := "00000000000000000000a8700385d4cc62628cc34629862121f84e6237689de8e45e151dcbc8cf30b33d"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewAssetModification()

	header, err := NewHeaderForCode(CodeAssetModification, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestContractOffer(t *testing.T) {
	// The hex is the body of the message
	body := "001d5465736c61202d205368617265686f6c6465722041677265656d656e740100000020c236f77c7abd7249489e7d2bb6c7e46ba3f4095956e78a584af753ece56cf6d1555341000055532d4341000000005c78d0f03468747470733a2f2f746f6b656e697a65642e636f6d2f436f6e74726163742f3371656f534367374a6d66536e4a65734a466f6a6a0a5465736c6120496e632e502368747470733a2f2f6578616d706c652e636f6d2f696d616765732f6c6f676f2e706e6709546f6b656e697a6564492492496db6db6d249b6db6db6db6c0000000000000000001010100010132053133353737000c466169726d6f6e742041766507526f62696e6f684243000000555341053530323130136a616d657340746f6b656e697a65642e636f6d0a303432303139393939390000"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewContractOffer()

	header, err := NewHeaderForCode(CodeContractOffer, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestContractFormation(t *testing.T) {
	// The hex is the body of the message
	body := "001d5465736c61202d205368617265686f6c6465722041677265656d656e740100000020c236f77c7abd7249489e7d2bb6c7e46ba3f4095956e78a584af753ece56cf6d1555341000055532d434100000163400d3f003468747470733a2f2f746f6b656e697a65642e636f6d2f436f6e74726163742f3371656f534367374a6d66536e4a65734a466f6a6a0a5465736c6120496e632e502368747470733a2f2f6578616d706c652e636f6d2f696d616765732f6c6f676f2e706e6709546f6b656e697a6564492492496db6db6d249b6db6db6db6c0000000000000000001010000010132053133353737000c466169726d6f6e742041766507526f62696e6f684243000000555341053530323130136a616d657340746f6b656e697a65642e636f6d0a30343438343834383438000000000000000000001588f8cc6a16430d"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewContractFormation()

	header, err := NewHeaderForCode(CodeContractFormation, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestContractAmendment(t *testing.T) {
	// The hex is the body of the message
	body := "000101002a00a8700385d4cc62628cc34629862121f84e6237689de8e45e151dcbc8cf30b33d"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewContractAmendment()

	header, err := NewHeaderForCode(CodeContractAmendment, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestStaticContractFormation(t *testing.T) {
	// The hex is the body of the message
	body := "001d5465736c61202d205368617265686f6c6465722041677265656d656e744e6f6e2d446973636c6f737572652041677265656d656e740100000020c236f77c7abd7249489e7d2bb6c7e46ba3f4095956e78a584af753ece56cf6d10000555341000055532d434100000163400d3f00000000005c78d0f03468747470733a2f2f746f6b656e697a65642e636f6d2f436f6e74726163742f3371656f534367374a6d66536e4a65734a466f6a6a3c762af9de09dc7403132f4a21bdf8aa02f41db9de7f9dab60409ab8cc907a3f00"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewStaticContractFormation()

	header, err := NewHeaderForCode(CodeStaticContractFormation, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestOrder(t *testing.T) {
	// The hex is the body of the message
	body := "0053484361706d3271737a6e686b7332337a3864383375343173383031396879726933694600002231377a4157616269706355486e35585039773847456333504b764735625947424d652453757072656d6520616e6420446973747269637420436f75727473204272697362616e65010000c236f77c7abd7249489e7d2bb6c7e46ba3f4095956e78a584af753ece56cf6d1f3318be9fb3f73e53b29868beae46b42911c2116f979a5d3284face90746cb3700000166550ce3800023536f7272792c206275742074686520636f757274206f72646572206d616465206d652e"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewOrder()

	header, err := NewHeaderForCode(CodeOrder, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestFreeze(t *testing.T) {
	// The hex is the body of the message
	body := "00001588f8cc6a16430d"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewFreeze()

	header, err := NewHeaderForCode(CodeFreeze, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestThaw(t *testing.T) {
	// The hex is the body of the message
	body := "00001588f8cc6a16430d"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewThaw()

	header, err := NewHeaderForCode(CodeThaw, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestConfiscation(t *testing.T) {
	// The hex is the body of the message
	body := "000000000000000027101588f8cc6a16430d"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewConfiscation()

	header, err := NewHeaderForCode(CodeConfiscation, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestReconciliation(t *testing.T) {
	// The hex is the body of the message
	body := "00001588f8cc6a16430d"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewReconciliation()

	header, err := NewHeaderForCode(CodeReconciliation, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestInitiative(t *testing.T) {
	// The hex is the body of the message
	body := "0053484361706d3271737a6e686b7332337a3864383375343173383031396879726933690100000f4142434445464748494a4b4c4d4e4f0f00204368616e676520746865206e616d65206f662074686520436f6e74726163742e77201b0094f50df309f0343e4f44dae64d0de503c91038faf2c6b039f9f18aec0000016482fd5d80"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewInitiative()

	header, err := NewHeaderForCode(CodeInitiative, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestReferendum(t *testing.T) {
	// The hex is the body of the message
	body := "000152524561706d3271737a6e686b7332337a3864383375343173383031396879726933690100000f4142434445464748494a4b4c4d4e4f0f00204368616e676520746865206e616d65206f662074686520436f6e74726163742e77201b0094f50df309f0343e4f44dae64d0de503c91038faf2c6b039f9f18aec0000016482fd5d80"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewReferendum()

	header, err := NewHeaderForCode(CodeReferendum, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestVote(t *testing.T) {
	// The hex is the body of the message
	body := "1588f8cc6a16430d"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewVote()

	header, err := NewHeaderForCode(CodeVote, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestBallotCast(t *testing.T) {
	// The hex is the body of the message
	body := "52524561706d3271737a6e686b7332337a386438337534317338303139687972693369f3318be9fb3f73e53b29868beae46b42911c2116f979a5d3284face90746cb370141"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewBallotCast()

	header, err := NewHeaderForCode(CodeBallotCast, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestBallotCounted(t *testing.T) {
	// The hex is the body of the message
	body := "1588f8cc6a16430d"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewBallotCounted()

	header, err := NewHeaderForCode(CodeBallotCounted, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestResult(t *testing.T) {
	// The hex is the body of the message
	body := "0053484361706d3271737a6e686b7332337a3864383375343173383031396879726933690000f2318be9fb3f73e53a29868beae46b42911c2116f979a5d3284face90746cb37010000000000000bb801321588f8cc6a16430d"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewResult()

	header, err := NewHeaderForCode(CodeResult, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestMessage(t *testing.T) {
	// The hex is the body of the message
	body := "0002000100023600000c48656c6c6f20776f726c6421"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewMessage()

	header, err := NewHeaderForCode(CodeMessage, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestRejection(t *testing.T) {
	// The hex is the body of the message
	body := "000200010002010024536f7272792c20796f7520646f6e2774206861766520656e6f75676820746f6b656e732e1588f8cc6a16430d"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewRejection()

	header, err := NewHeaderForCode(CodeRejection, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestEstablishment(t *testing.T) {
	// The hex is the body of the message
	body := "0000174e6f72746820416d65726963612057686974656c697374"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewEstablishment()

	header, err := NewHeaderForCode(CodeEstablishment, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestAddition(t *testing.T) {
	// The hex is the body of the message
	body := "000008757365726e616d65"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewAddition()

	header, err := NewHeaderForCode(CodeAddition, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestAlteration(t *testing.T) {
	// The hex is the body of the message
	body := "00001c4368616e67656420436f756e747279206f66205265736964656e6365"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewAlteration()

	header, err := NewHeaderForCode(CodeAlteration, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestRemoval(t *testing.T) {
	// The hex is the body of the message
	body := "00002b52656d6f7665642064756520746f2076696f6c6174696f6e206f6620636f6d70616e7920706f6c6963792e"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewRemoval()

	header, err := NewHeaderForCode(CodeRemoval, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestSend(t *testing.T) {
	// The hex is the body of the message
	body := "0053484361706d3271737a6e686b7332337a3864383375343173383031396879726933690000"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewSend()

	header, err := NewHeaderForCode(CodeSend, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestExchange(t *testing.T) {
	// The hex is the body of the message
	body := "0052524561706d3271737a6e686b7332337a0000000000000000000000000000000000000000016331e3c2004155443ba3d70a3c23d70a31485132554c7544375435796b6175635a334b6d546f3469323939323551613669630000"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewExchange()

	header, err := NewHeaderForCode(CodeExchange, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestSwap(t *testing.T) {
	// The hex is the body of the message
	body := "0052524572616e3271737a6e68697335337a00000000000000000000000000000000000053484361706d3271737a6e686b7332337a0000000000000000000000000000000000000000016331e3c2004155443ba3d70a3c23d70a31485132554c7544375435796b6175635a334b6d546f34693239393235516136696300000000"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewSwap()

	header, err := NewHeaderForCode(CodeSwap, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

func TestSettlement(t *testing.T) {
	// The hex is the body of the message
	body := "005352524561706d3271737a6e686b7332337a38643833753431733830313968797269336953484361706d3271737a6e686b7332337a38643833753431733830313968797269336900001588f8cc6a16430d"

	b, err := hex.DecodeString(body)
	if err != nil {
		t.Fatalf("Invalid hex value : hex=%v : %v", body, err)
	}

	// Create a valid header for the body
	m := NewSettlement()

	header, err := NewHeaderForCode(CodeSettlement, len(b))
	if err != nil {
		t.Fatal(err)
	}

	m.Header = *header

	headerBytes, err := m.Header.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	// This is the target byte payload
	want := headerBytes
	want = append(want, b...)

	n, err := m.Write(want)
	if err != nil {
		t.Fatal(err)
	}

	if n != len(want) {
		t.Fatalf("got %v, want %v", n, len(want))
	}

	// Serializing the message should give us the same bytes
	got, err := m.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}

