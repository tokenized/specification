package protocol

import (
	"reflect"
	"testing"
	"time"
)

func TestAssetDefinition(t *testing.T) {
	// Create a randomized object
	initialMessage := AssetDefinition{}
	// AssetType (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+0))
		}
		initialMessage.AssetType = string(text)
	}

	// AssetID (fixedchar)
	{
		text := make([]byte, 0, 32)
		for i := uint64(0); i < 32; i++ {
			text = append(text, byte(65+i+1))
		}
		initialMessage.AssetID = string(text)
	}

	// AssetAuthFlags (bin)
	// bin test not setup

	// TransfersPermitted (bool)
	// bool test not setup

	// TradeRestrictions (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+4))
		}
		initialMessage.TradeRestrictions = string(text)
	}

	// EnforcementOrdersPermitted (bool)
	// bool test not setup

	// VoteMultiplier (uint)
	// uint test not setup

	// ReferendumProposal (bool)
	// bool test not setup

	// InitiativeProposal (bool)
	// bool test not setup

	// AssetModificationGovernance (bool)
	// bool test not setup

	// TokenQty (uint)
	// uint test not setup

	// ContractFeeCurrency (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+11))
		}
		initialMessage.ContractFeeCurrency = string(text)
	}

	// ContractFeeVar (float)
	// float test not setup

	// ContractFeeFixed (float)
	// float test not setup

	// AssetPayload (varbin)
	initialMessage.AssetPayload = make([]byte, 0, 16)
	for i := uint64(0); i < 16; i++ {
		initialMessage.AssetPayload = append(initialMessage.AssetPayload, byte(65+i+14))
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := AssetDefinition{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestAssetCreation(t *testing.T) {
	// Create a randomized object
	initialMessage := AssetCreation{}
	// AssetType (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+0))
		}
		initialMessage.AssetType = string(text)
	}

	// AssetID (fixedchar)
	{
		text := make([]byte, 0, 32)
		for i := uint64(0); i < 32; i++ {
			text = append(text, byte(65+i+1))
		}
		initialMessage.AssetID = string(text)
	}

	// AssetAuthFlags (bin)
	// bin test not setup

	// TransfersPermitted (bool)
	// bool test not setup

	// TradeRestrictions (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+4))
		}
		initialMessage.TradeRestrictions = string(text)
	}

	// EnforcementOrdersPermitted (bool)
	// bool test not setup

	// VoteMultiplier (uint)
	// uint test not setup

	// ReferendumProposal (bool)
	// bool test not setup

	// InitiativeProposal (bool)
	// bool test not setup

	// AssetModificationGovernance (bool)
	// bool test not setup

	// TokenQty (uint)
	// uint test not setup

	// ContractFeeCurrency (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+11))
		}
		initialMessage.ContractFeeCurrency = string(text)
	}

	// ContractFeeVar (float)
	// float test not setup

	// ContractFeeFixed (float)
	// float test not setup

	// AssetPayload (varbin)
	initialMessage.AssetPayload = make([]byte, 0, 16)
	for i := uint64(0); i < 16; i++ {
		initialMessage.AssetPayload = append(initialMessage.AssetPayload, byte(65+i+14))
	}

	// AssetRevision (uint)
	// uint test not setup

	// Timestamp (timestamp)
	initialMessage.Timestamp = uint64(time.Now().UnixNano())

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := AssetCreation{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestAssetModification(t *testing.T) {
	// Create a randomized object
	initialMessage := AssetModification{}
	// AssetType (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+0))
		}
		initialMessage.AssetType = string(text)
	}

	// AssetID (fixedchar)
	{
		text := make([]byte, 0, 32)
		for i := uint64(0); i < 32; i++ {
			text = append(text, byte(65+i+1))
		}
		initialMessage.AssetID = string(text)
	}

	// AssetRevision (uint)
	// uint test not setup

	// Modifications (Amendment[])
	// Amendment[] test not setup

	// RefTxID (sha256)
	// sha256 test not setup

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := AssetModification{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestContractOffer(t *testing.T) {
	// Create a randomized object
	initialMessage := ContractOffer{}
	// ContractName (varchar)
	initialMessage.ContractName = "Text 0"

	// ContractFileType (uint)
	// uint test not setup

	// ContractFile (varbin)
	initialMessage.ContractFile = make([]byte, 0, 32)
	for i := uint64(0); i < 32; i++ {
		initialMessage.ContractFile = append(initialMessage.ContractFile, byte(65+i+2))
	}

	// GoverningLaw (fixedchar)
	{
		text := make([]byte, 0, 5)
		for i := uint64(0); i < 5; i++ {
			text = append(text, byte(65+i+3))
		}
		initialMessage.GoverningLaw = string(text)
	}

	// Jurisdiction (fixedchar)
	{
		text := make([]byte, 0, 5)
		for i := uint64(0); i < 5; i++ {
			text = append(text, byte(65+i+4))
		}
		initialMessage.Jurisdiction = string(text)
	}

	// ContractExpiration (time)
	// time test not setup

	// ContractURI (varchar)
	initialMessage.ContractURI = "Text 6"

	// IssuerName (varchar)
	initialMessage.IssuerName = "Text 7"

	// IssuerType (fixedchar)
	// fixedchar test not setup

	// IssuerLogoURL (varchar)
	initialMessage.IssuerLogoURL = "Text 9"

	// ContractOperatorID (varchar)
	initialMessage.ContractOperatorID = "Text 10"

	// ContractAuthFlags (bin)
	// bin test not setup

	// VotingSystems (VotingSystem[])
	// VotingSystem[] test not setup

	// RestrictedQtyAssets (uint)
	// uint test not setup

	// ReferendumProposal (bool)
	// bool test not setup

	// InitiativeProposal (bool)
	// bool test not setup

	// Registries (Registry[])
	// Registry[] test not setup

	// IssuerAddress (bool)
	// bool test not setup

	// UnitNumber (varchar)
	initialMessage.UnitNumber = "Text 18"

	// BuildingNumber (varchar)
	initialMessage.BuildingNumber = "Text 19"

	// Street (varchar)
	initialMessage.Street = "Text 20"

	// SuburbCity (varchar)
	initialMessage.SuburbCity = "Text 21"

	// TerritoryStateProvinceCode (fixedchar)
	{
		text := make([]byte, 0, 5)
		for i := uint64(0); i < 5; i++ {
			text = append(text, byte(65+i+22))
		}
		initialMessage.TerritoryStateProvinceCode = string(text)
	}

	// CountryCode (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+23))
		}
		initialMessage.CountryCode = string(text)
	}

	// PostalZIPCode (varchar)
	initialMessage.PostalZIPCode = "Text 24"

	// EmailAddress (varchar)
	initialMessage.EmailAddress = "Text 25"

	// PhoneNumber (varchar)
	initialMessage.PhoneNumber = "Text 26"

	// KeyRoles (KeyRole[])
	// KeyRole[] test not setup

	// NotableRoles (NotableRole[])
	// NotableRole[] test not setup

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := ContractOffer{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestContractFormation(t *testing.T) {
	// Create a randomized object
	initialMessage := ContractFormation{}
	// ContractName (varchar)
	initialMessage.ContractName = "Text 0"

	// ContractFileType (uint)
	// uint test not setup

	// ContractFile (varbin)
	initialMessage.ContractFile = make([]byte, 0, 32)
	for i := uint64(0); i < 32; i++ {
		initialMessage.ContractFile = append(initialMessage.ContractFile, byte(65+i+2))
	}

	// GoverningLaw (fixedchar)
	{
		text := make([]byte, 0, 5)
		for i := uint64(0); i < 5; i++ {
			text = append(text, byte(65+i+3))
		}
		initialMessage.GoverningLaw = string(text)
	}

	// Jurisdiction (fixedchar)
	{
		text := make([]byte, 0, 5)
		for i := uint64(0); i < 5; i++ {
			text = append(text, byte(65+i+4))
		}
		initialMessage.Jurisdiction = string(text)
	}

	// ContractExpiration (time)
	// time test not setup

	// ContractURI (varchar)
	initialMessage.ContractURI = "Text 6"

	// IssuerName (varchar)
	initialMessage.IssuerName = "Text 7"

	// IssuerType (fixedchar)
	// fixedchar test not setup

	// IssuerLogoURL (varchar)
	initialMessage.IssuerLogoURL = "Text 9"

	// ContractOperatorID (varchar)
	initialMessage.ContractOperatorID = "Text 10"

	// ContractAuthFlags (bin)
	// bin test not setup

	// VotingSystems (VotingSystem[])
	// VotingSystem[] test not setup

	// RestrictedQtyAssets (uint)
	// uint test not setup

	// ReferendumProposal (bool)
	// bool test not setup

	// InitiativeProposal (bool)
	// bool test not setup

	// Registries (Registry[])
	// Registry[] test not setup

	// IssuerAddress (bool)
	// bool test not setup

	// UnitNumber (varchar)
	initialMessage.UnitNumber = "Text 18"

	// BuildingNumber (varchar)
	initialMessage.BuildingNumber = "Text 19"

	// Street (varchar)
	initialMessage.Street = "Text 20"

	// SuburbCity (varchar)
	initialMessage.SuburbCity = "Text 21"

	// TerritoryStateProvinceCode (fixedchar)
	{
		text := make([]byte, 0, 5)
		for i := uint64(0); i < 5; i++ {
			text = append(text, byte(65+i+22))
		}
		initialMessage.TerritoryStateProvinceCode = string(text)
	}

	// CountryCode (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+23))
		}
		initialMessage.CountryCode = string(text)
	}

	// PostalZIPCode (varchar)
	initialMessage.PostalZIPCode = "Text 24"

	// EmailAddress (varchar)
	initialMessage.EmailAddress = "Text 25"

	// PhoneNumber (varchar)
	initialMessage.PhoneNumber = "Text 26"

	// KeyRoles (KeyRole[])
	// KeyRole[] test not setup

	// NotableRoles (NotableRole[])
	// NotableRole[] test not setup

	// ContractRevision (uint)
	// uint test not setup

	// Timestamp (timestamp)
	initialMessage.Timestamp = uint64(time.Now().UnixNano())

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := ContractFormation{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestContractAmendment(t *testing.T) {
	// Create a randomized object
	initialMessage := ContractAmendment{}
	// ChangeIssuerAddress (bool)
	// bool test not setup

	// ChangeOperatorAddress (bool)
	// bool test not setup

	// ContractRevision (uint)
	// uint test not setup

	// Amendments (Amendment[])
	// Amendment[] test not setup

	// RefTxID (SHA256)
	// SHA256 test not setup

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := ContractAmendment{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestStaticContractFormation(t *testing.T) {
	// Create a randomized object
	initialMessage := StaticContractFormation{}
	// ContractName (varchar)
	initialMessage.ContractName = "Text 0"

	// ContractType (varchar)
	initialMessage.ContractType = "Text 1"

	// ContractFileType (uint)
	// uint test not setup

	// ContractFile (varbin)
	initialMessage.ContractFile = make([]byte, 0, 32)
	for i := uint64(0); i < 32; i++ {
		initialMessage.ContractFile = append(initialMessage.ContractFile, byte(65+i+3))
	}

	// ContractRevision (uint)
	// uint test not setup

	// GoverningLaw (fixedchar)
	{
		text := make([]byte, 0, 5)
		for i := uint64(0); i < 5; i++ {
			text = append(text, byte(65+i+5))
		}
		initialMessage.GoverningLaw = string(text)
	}

	// Jurisdiction (fixedchar)
	{
		text := make([]byte, 0, 5)
		for i := uint64(0); i < 5; i++ {
			text = append(text, byte(65+i+6))
		}
		initialMessage.Jurisdiction = string(text)
	}

	// EffectiveDate (time)
	// time test not setup

	// ContractExpiration (time)
	// time test not setup

	// ContractURI (varchar)
	initialMessage.ContractURI = "Text 9"

	// PrevRevTxID (sha256)
	// sha256 test not setup

	// Entities (Entity[])
	// Entity[] test not setup

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := StaticContractFormation{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestOrder(t *testing.T) {
	// Create a randomized object
	initialMessage := Order{}
	// AssetType (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+0))
		}
		initialMessage.AssetType = string(text)
	}

	// AssetID (fixedchar)
	{
		text := make([]byte, 0, 32)
		for i := uint64(0); i < 32; i++ {
			text = append(text, byte(65+i+1))
		}
		initialMessage.AssetID = string(text)
	}

	// ComplianceAction (fixedchar)
	// fixedchar test not setup

	// TargetAddresses (TargetAddress[])
	// TargetAddress[] test not setup

	// DepositAddress (Address)
	// Address test not setup

	// AuthorityName (varchar)
	initialMessage.AuthorityName = "Text 5"

	// SigAlgoAddressList (uint)
	// uint test not setup

	// AuthorityPublicKey (varchar)
	initialMessage.AuthorityPublicKey = "Text 7"

	// OrderSignature (varchar)
	initialMessage.OrderSignature = "Text 8"

	// SupportingEvidenceHash (sha256)
	// sha256 test not setup

	// RefTxnID (sha256)
	// sha256 test not setup

	// FreezePeriod (time)
	// time test not setup

	// Message (varchar)
	initialMessage.Message = "Text 12"

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Order{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestFreeze(t *testing.T) {
	// Create a randomized object
	initialMessage := Freeze{}
	// Addresses (Address[])
	// Address[] test not setup

	// Timestamp (timestamp)
	initialMessage.Timestamp = uint64(time.Now().UnixNano())

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Freeze{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestThaw(t *testing.T) {
	// Create a randomized object
	initialMessage := Thaw{}
	// Addresses (Address[])
	// Address[] test not setup

	// RefTxnID (sha256)
	// sha256 test not setup

	// Timestamp (timestamp)
	initialMessage.Timestamp = uint64(time.Now().UnixNano())

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Thaw{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestConfiscation(t *testing.T) {
	// Create a randomized object
	initialMessage := Confiscation{}
	// Addresses (Address[])
	// Address[] test not setup

	// DepositQty (uint)
	// uint test not setup

	// Timestamp (timestamp)
	initialMessage.Timestamp = uint64(time.Now().UnixNano())

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Confiscation{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestReconciliation(t *testing.T) {
	// Create a randomized object
	initialMessage := Reconciliation{}
	// Addresses (Address[])
	// Address[] test not setup

	// Timestamp (timestamp)
	initialMessage.Timestamp = uint64(time.Now().UnixNano())

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Reconciliation{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestInitiative(t *testing.T) {
	// Create a randomized object
	initialMessage := Initiative{}
	// AssetType (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+0))
		}
		initialMessage.AssetType = string(text)
	}

	// AssetID (fixedchar)
	{
		text := make([]byte, 0, 32)
		for i := uint64(0); i < 32; i++ {
			text = append(text, byte(65+i+1))
		}
		initialMessage.AssetID = string(text)
	}

	// VoteSystem (uint)
	// uint test not setup

	// Proposal (bool)
	// bool test not setup

	// ProposedChanges (Amendment[])
	// Amendment[] test not setup

	// VoteOptions (varchar)
	initialMessage.VoteOptions = "Text 5"

	// VoteMax (uint)
	// uint test not setup

	// ProposalDescription (varchar)
	initialMessage.ProposalDescription = "Text 7"

	// ProposalDocumentHash (sha256)
	// sha256 test not setup

	// VoteCutOffTimestamp (time)
	// time test not setup

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Initiative{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestReferendum(t *testing.T) {
	// Create a randomized object
	initialMessage := Referendum{}
	// AssetSpecificVote (bool)
	// bool test not setup

	// AssetType (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+1))
		}
		initialMessage.AssetType = string(text)
	}

	// AssetID (fixedchar)
	{
		text := make([]byte, 0, 32)
		for i := uint64(0); i < 32; i++ {
			text = append(text, byte(65+i+2))
		}
		initialMessage.AssetID = string(text)
	}

	// VoteSystem (uint)
	// uint test not setup

	// Proposal (bool)
	// bool test not setup

	// ProposedChanges (Amendment[])
	// Amendment[] test not setup

	// VoteOptions (varchar)
	initialMessage.VoteOptions = "Text 6"

	// VoteMax (uint)
	// uint test not setup

	// ProposalDescription (varchar)
	initialMessage.ProposalDescription = "Text 8"

	// ProposalDocumentHash (sha256)
	// sha256 test not setup

	// VoteCutOffTimestamp (time)
	// time test not setup

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Referendum{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestVote(t *testing.T) {
	// Create a randomized object
	initialMessage := Vote{}
	// Timestamp (timestamp)
	initialMessage.Timestamp = uint64(time.Now().UnixNano())

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Vote{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestBallotCast(t *testing.T) {
	// Create a randomized object
	initialMessage := BallotCast{}
	// AssetType (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+0))
		}
		initialMessage.AssetType = string(text)
	}

	// AssetID (fixedchar)
	{
		text := make([]byte, 0, 32)
		for i := uint64(0); i < 32; i++ {
			text = append(text, byte(65+i+1))
		}
		initialMessage.AssetID = string(text)
	}

	// VoteTxnID (sha256)
	// sha256 test not setup

	// Vote (varchar)
	initialMessage.Vote = "Text 3"

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := BallotCast{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestBallotCounted(t *testing.T) {
	// Create a randomized object
	initialMessage := BallotCounted{}
	// Timestamp (timestamp)
	initialMessage.Timestamp = uint64(time.Now().UnixNano())

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := BallotCounted{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestResult(t *testing.T) {
	// Create a randomized object
	initialMessage := Result{}
	// AssetType (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+0))
		}
		initialMessage.AssetType = string(text)
	}

	// AssetID (fixedchar)
	{
		text := make([]byte, 0, 32)
		for i := uint64(0); i < 32; i++ {
			text = append(text, byte(65+i+1))
		}
		initialMessage.AssetID = string(text)
	}

	// Proposal (bool)
	// bool test not setup

	// ProposedChanges (Amendment[])
	// Amendment[] test not setup

	// VoteTxnID (sha256)
	// sha256 test not setup

	// VoteOptionsCount (uint)
	// uint test not setup

	// OptionXTally (uint)
	// uint test not setup

	// Result (varchar)
	initialMessage.Result = "Text 7"

	// Timestamp (timestamp)
	initialMessage.Timestamp = uint64(time.Now().UnixNano())

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Result{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestMessage(t *testing.T) {
	// Create a randomized object
	initialMessage := Message{}
	// QtyReceivingAddresses (uint8)
	// uint8 test not setup

	// AddressIndexes (uint16[])
	// uint16[] test not setup

	// MessageType (fixedchar)
	{
		text := make([]byte, 0, 2)
		for i := uint64(0); i < 2; i++ {
			text = append(text, byte(65+i+2))
		}
		initialMessage.MessageType = string(text)
	}

	// MessagePayload (varchar)
	initialMessage.MessagePayload = "Text 3"

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Message{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestRejection(t *testing.T) {
	// Create a randomized object
	initialMessage := Rejection{}
	// QtyReceivingAddresses (uint8)
	// uint8 test not setup

	// AddressIndexes (uint16[])
	// uint16[] test not setup

	// RejectionType (uint)
	// uint test not setup

	// MessagePayload (varchar)
	initialMessage.MessagePayload = "Text 3"

	// Timestamp (timestamp)
	initialMessage.Timestamp = uint64(time.Now().UnixNano())

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Rejection{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestEstablishment(t *testing.T) {
	// Create a randomized object
	initialMessage := Establishment{}
	// Message (varchar)
	initialMessage.Message = "Text 0"

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Establishment{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestAddition(t *testing.T) {
	// Create a randomized object
	initialMessage := Addition{}
	// Message (varchar)
	initialMessage.Message = "Text 0"

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Addition{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestAlteration(t *testing.T) {
	// Create a randomized object
	initialMessage := Alteration{}
	// Message (varchar)
	initialMessage.Message = "Text 0"

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Alteration{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestRemoval(t *testing.T) {
	// Create a randomized object
	initialMessage := Removal{}
	// Message (varchar)
	initialMessage.Message = "Text 0"

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Removal{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestTransfer(t *testing.T) {
	// Create a randomized object
	initialMessage := Transfer{}
	// Assets (AssetTransfer[])
	// AssetTransfer[] test not setup

	// OfferExpiry (time)
	// time test not setup

	// ExchangeFeeCurrency (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+2))
		}
		initialMessage.ExchangeFeeCurrency = string(text)
	}

	// ExchangeFeeVar (float)
	// float test not setup

	// ExchangeFeeFixed (float)
	// float test not setup

	// ExchangeFeeAddress (Address)
	// Address test not setup

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Transfer{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}

func TestSettlement(t *testing.T) {
	// Create a randomized object
	initialMessage := Settlement{}
	// Assets (AssetSettlement[])
	// AssetSettlement[] test not setup

	// Timestamp (timestamp)
	initialMessage.Timestamp = uint64(time.Now().UnixNano())

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Settlement{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	if !reflect.DeepEqual(initialMessage, decodedMessage) {
		t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	}
}
