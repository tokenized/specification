package protocol

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"strconv"
	"time"

	"github.com/tokenized/envelope/pkg/golang/envelope"
	envelopeV0 "github.com/tokenized/envelope/pkg/golang/envelope/v0"
	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/specification/dist/golang/actions"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ripemd160"
)

const (
	// ProtocolID is the current protocol ID
	ProtocolID = "TKN"

	// TestProtocolID is the current protocol ID for testing.
	TestProtocolID = "test.TKN"

	// Version of the Tokenized protocol.
	Version = uint64(0)

	FlagProtocolID = "flag"

	FlagVersion = uint64(0)

	AssetCodeSize = 20

	ContractCodeSize = 32
)

var (
	ErrNotTokenized   = errors.New("Not Tokenized")
	ErrNotFlag        = errors.New("Not Flag")
	ErrUnknownVersion = errors.New("Unknown Version")

	// DefaultEndian specifies the order of bytes for encoding integers.
	DefaultEndian = binary.LittleEndian
)

func GetProtocolID(isTest bool) []byte {
	if isTest {
		return []byte(TestProtocolID)
	} else {
		return []byte(ProtocolID)
	}
}

// Serialize serializes an action into a Tokenized OP_RETURN script.
func Serialize(action actions.Action, isTest bool) ([]byte, error) {
	message, err := WrapAction(action, isTest)
	if err != nil {
		return nil, errors.Wrap(err, "wrap action")
	}

	var buf bytes.Buffer
	err = message.Serialize(&buf)
	if err != nil {
		return nil, errors.Wrap(err, "serialize envelope")
	}

	return buf.Bytes(), nil
}

// Deserialize reads an action from a Tokenized OP_RETURN script.
func Deserialize(script []byte, isTest bool) (actions.Action, error) {
	buf := bytes.NewReader(script)
	message, err := envelope.Deserialize(buf)
	if err == envelope.ErrNotEnvelope {
		return nil, ErrNotTokenized
	} else if err != nil {
		return nil, err
	}

	if !bytes.Equal(message.PayloadProtocol(), GetProtocolID(isTest)) {
		return nil, ErrNotTokenized
	}

	if message.PayloadVersion() != Version {
		return nil, ErrUnknownVersion
	}

	return actions.Deserialize(message.PayloadIdentifier(), message.Payload())
}

// WrapAction wraps an action in an envelope message.
func WrapAction(action actions.Action, isTest bool) (envelope.BaseMessage, error) {
	payload, err := proto.Marshal(action)
	if err != nil {
		return nil, errors.Wrap(err, "serialize action")
	}

	message := envelopeV0.NewMessage(GetProtocolID(isTest), Version, payload)
	message.SetPayloadIdentifier([]byte(action.Code()))

	return message, nil
}

// SerializeFlagOutputScript creates a locking script containing the flag value for a relationship
//   message.
func SerializeFlagOutputScript(flag []byte) ([]byte, error) {
	message := envelopeV0.NewMessage([]byte(FlagProtocolID), FlagVersion, flag)
	var buf bytes.Buffer
	if err := message.Serialize(&buf); err != nil {
		return nil, errors.Wrap(err, "Failed to serialize flag envelope")
	}
	return buf.Bytes(), nil
}

// DeserializeFlagOutputScript returns a flag value if the script is a flag script.
func DeserializeFlagOutputScript(script []byte) ([]byte, error) {
	buf := bytes.NewReader(script)
	message, err := envelope.Deserialize(buf)
	if err == envelope.ErrNotEnvelope {
		return nil, ErrNotFlag
	} else if err != nil {
		return nil, err
	}

	if !bytes.Equal(message.PayloadProtocol(), []byte(FlagProtocolID)) {
		return nil, ErrNotFlag
	}

	if message.PayloadVersion() != FlagVersion {
		return nil, ErrUnknownVersion
	}

	return message.Payload(), nil
}

// AssetCodeFromContract generates a "unique" deterministic asset code from a contract public key
//   hash and an asset index.
func AssetCodeFromContract(contractAddress bitcoin.RawAddress, index uint64) bitcoin.Hash20 {
	hash256 := sha256.New()
	hash256.Write(contractAddress.Bytes())
	binary.Write(hash256, DefaultEndian, &index)
	hash := hash256.Sum(nil)

	hash160 := ripemd160.New()
	hash160.Write(hash)
	hash = hash160.Sum(nil)

	var result bitcoin.Hash20
	copy(result[:], hash)
	return result
}

// AssetCodeFromBytes returns a AssetCode with the specified bytes.
func AssetCodeFromBytes(b []byte) bitcoin.Hash20 {
	var result bitcoin.Hash20
	copy(result[:], b)
	return result
}

// AssetID encodes an asset ID.
//
// AssetID = AssetType(3 characters) + base58(AssetCode + checksum)
//
// There is a special case for BSV, which will be returned as BSV.
func AssetID(assetType string, code bitcoin.Hash20) string {
	if assetType == "BSV" {
		return assetType
	}

	b := code.Bytes()

	// Perform Double SHA-256 hash
	checksum := bitcoin.DoubleSha256(b)

	// Append the first 4 checksum bytes
	b = append(b, checksum[:4]...)

	// Prepend with type and encode as text with base 58.
	return assetType + bitcoin.Base58(b)
}

// DecodeAssetID decodes the asset type and asset code from an asset ID.
func DecodeAssetID(id string) (string, bitcoin.Hash20, error) {
	if id == "BSV" {
		// Bitcoin asset id. Asset code all zeros.
		return "BSV", bitcoin.Hash20{}, nil
	}

	if len(id) < AssetCodeSize+3 {
		return "", bitcoin.Hash20{}, fmt.Errorf("Asset ID too short : %s", id)
	}

	assetType := id[:3]
	text := id[3:]

	b := bitcoin.Base58Decode(text)

	if len(b) != AssetCodeSize+4 {
		return "", bitcoin.Hash20{}, fmt.Errorf("Asset ID too short : %s", id)
	}

	// Verify checksum
	checksum := b[AssetCodeSize:]
	b = b[:AssetCodeSize]
	hash := bitcoin.DoubleSha256(b)
	if !bytes.Equal(hash[:4], checksum) {
		return "", bitcoin.Hash20{}, fmt.Errorf("Invalid Asset ID checksum : %s", id)
	}

	code, err := bitcoin.NewHash20(b)
	if err != nil {
		return "", bitcoin.Hash20{}, errors.Wrap(err, "new hash")
	}

	return assetType, *code, nil
}

// ------------------------------------------------------------------------------------------------
// ContractCode represents a unique identifier for a Tokenized static contract.
type ContractCode bitcoin.Hash32

// Validate returns an error if the value is invalid
func (code *ContractCode) Validate() error {
	return nil
}

// IsZero returns true if the ContractCode is all zeroes. (empty)
func (code *ContractCode) IsZero() bool {
	if code == nil {
		return true
	}
	zero := make([]byte, ContractCodeSize)
	return bytes.Equal(code[:], zero)
}

// AssetCodeFromBytes returns a ContractCode with the specified bytes.
func ContractCodeFromBytes(b []byte) *ContractCode {
	var result ContractCode
	copy(result[:], b)
	return &result
}

// ------------------------------------------------------------------------------------------------
// Timestamp represents a time in the Tokenized protocol.
type Timestamp struct {
	nanoseconds uint64 // nanoseconds since Unix epoch
}

// NewTimestamp returns a new timestamp from nanoseconds.
func NewTimestamp(value uint64) Timestamp {
	return Timestamp{nanoseconds: value}
}

func DeserializeTimestamp(buf *bytes.Reader) (Timestamp, error) {
	var result Timestamp
	if err := binary.Read(buf, DefaultEndian, &result.nanoseconds); err != nil {
		return result, err
	}
	return result, nil
}

// CurrentTimestamp returns a Timestamp containing the current time.
func CurrentTimestamp() Timestamp {
	return Timestamp{nanoseconds: uint64(time.Now().UnixNano())}
}

// Validate returns an error if the value is invalid
func (time *Timestamp) Validate() error {
	return nil
}

// Equal returns true if the specified values are the same.
func (time *Timestamp) Equal(other Timestamp) bool {
	return time.nanoseconds == other.nanoseconds
}

// Nano returns the nanoseconds since the Unix epoch for the Timestamp.
func (time *Timestamp) Nano() uint64 {
	return time.nanoseconds
}

// Nano returns the seconds since the Unix epoch for the Timestamp.
func (time *Timestamp) Seconds() uint32 {
	return uint32(time.nanoseconds / 1000000000)
}

// String converts to a string
func (t *Timestamp) String() string {
	return time.Unix(int64(t.nanoseconds)/1000000000, 0).String()
}

// Serialize serializes a timestamp into a buffer.
func (time *Timestamp) Serialize(buf *bytes.Buffer) error {
	if err := binary.Write(buf, DefaultEndian, &time.nanoseconds); err != nil {
		return err
	}
	return nil
}

// MarshalJSON converts to json.
func (time *Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatUint(time.nanoseconds, 10)), nil
}

// UnmarshalJSON converts from json.
func (time *Timestamp) UnmarshalJSON(data []byte) error {
	if len(data) == 2 && data[0] == '{' && data[1] == '}' {
		time.nanoseconds = 0
		return nil
	}

	var err error
	time.nanoseconds, err = strconv.ParseUint(string(data), 10, 64)
	return err
}
