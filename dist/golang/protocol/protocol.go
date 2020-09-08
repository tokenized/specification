package protocol

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/tokenized/envelope/pkg/golang/envelope"
	v0 "github.com/tokenized/envelope/pkg/golang/envelope/v0"
	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/specification/dist/golang/actions"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

const (
	// ProtocolID is the current protocol ID
	ProtocolID = "tokenized"

	// TestProtocolID is the current protocol ID for testing.
	TestProtocolID = "test.tokenized"

	// Version of the Tokenized protocol.
	Version = uint64(0)

	FlagProtocolID = "flag"

	FlagVersion = uint64(0)
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

	message := v0.NewMessage(GetProtocolID(isTest), Version, payload)
	message.SetPayloadIdentifier([]byte(action.Code()))

	return message, nil
}

// SerializeFlagOutputScript creates a locking script containing the flag value for a relationship
//   message.
func SerializeFlagOutputScript(flag []byte) ([]byte, error) {
	message := v0.NewMessage([]byte(FlagProtocolID), FlagVersion, flag)
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

// ------------------------------------------------------------------------------------------------
// TxId represents a Bitcoin transaction ID. (Double SHA256 of tx data)
type TxId struct {
	data [32]byte
}

var zeroTxId TxId

func TxIdFromBytes(data []byte) *TxId {
	var result TxId
	copy(result.data[:], data)
	return &result
}

func DeserializeTxId(buf *bytes.Reader) (*TxId, error) {
	var result TxId
	if _, err := buf.Read(result.data[:]); err != nil {
		return nil, err
	}
	return &result, nil
}

// Validate returns an error if the value is invalid
func (id *TxId) Validate() error {
	return nil
}

// IsZero returns true if the tx id is all zeros.
func (id *TxId) IsZero() bool {
	if id == nil {
		return true
	}
	return bytes.Equal(id.data[:], zeroTxId.data[:])
}

// Equal returns true if the specified values are the same.
func (id *TxId) Equal(other TxId) bool {
	if id == nil {
		return id == nil
	}
	return bytes.Equal(id.data[:], other.data[:])
}

// Bytes returns the byte slice for the TxId.
func (id *TxId) Bytes() []byte {
	if id == nil {
		return nil
	}
	return id.data[:]
}

// String converts to a string
func (id *TxId) String() string {
	return fmt.Sprintf("%x", id.data[:])
}

// Serialize serializes a txid into a buffer.
func (id *TxId) Serialize(buf *bytes.Buffer) error {
	_, err := buf.Write(id.data[:])
	return err
}

// MarshalJSON converts to json.
func (id *TxId) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%x\"", id.data)), nil
}

// UnmarshalJSON converts from json.
func (id *TxId) UnmarshalJSON(data []byte) error {
	if len(data) < 2 {
		return fmt.Errorf("Too short for TxId hex data : %d", len(data))
	}
	n, err := hex.Decode(id.data[:], data[1:len(data)-1])
	if err != nil {
		return err
	}
	if n != 32 {
		return fmt.Errorf("Invalid TxId size : %d", n)
	}
	return nil
}

// Set sets the value specified
func (id *TxId) Set(value []byte) error {
	if len(value) != len(id.data) {
		return errors.New("Invalid value length for txid")
	}
	copy(id.data[:], value)
	return nil
}

// AssetID encodes an asset ID.
//
// AssetID = AssetType(3 characters) + base58(AssetCode + checksum)
//
// There is a special case for BSV, which will be returned as BSV.
func AssetID(assetType string, code AssetCode) string {
	if assetType == "BSV" {
		return assetType
	}

	data := code.Bytes()

	// Perform Double SHA-256 hash
	checksum := bitcoin.DoubleSha256(data)

	// Append the first 4 checksum bytes
	data = append(data, checksum[:4]...)

	// Prepend with type and encode as text with base 58.
	return assetType + bitcoin.Base58(data)
}

// DecodeAssetID decodes the asset type and asset code from an asset ID.
func DecodeAssetID(id string) (string, AssetCode, error) {
	if id == "BSV" {
		// Bitcoin asset id. Asset code all zeros.
		return "BSV", *AssetCodeFromBytes(make([]byte, 32)), nil
	}

	if len(id) < 35 {
		return "", AssetCode{}, fmt.Errorf("Asset ID too short : %s", id)
	}

	assetType := id[:3]
	text := id[3:]

	data := bitcoin.Base58Decode(text)

	if len(data) < 5 {
		return "", AssetCode{}, fmt.Errorf("Asset ID too short : %s", id)
	}

	// Verify checksum
	l := len(data)
	checksum := data[l-4:]
	data = data[:l-4]
	hash := bitcoin.DoubleSha256(data)
	if !bytes.Equal(hash[:4], checksum) {
		return "", AssetCode{}, fmt.Errorf("Invalid Asset ID checksum : %s", id)
	}

	return assetType, *AssetCodeFromBytes(data), nil
}

// ------------------------------------------------------------------------------------------------
// AssetCode represents a unique identifier for a Tokenized asset.
type AssetCode struct {
	data [32]byte
}

// Validate returns an error if the value is invalid
func (code *AssetCode) Validate() error {
	return nil
}

// IsZero returns true if the AssetCode is all zeroes. (empty)
func (code *AssetCode) IsZero() bool {
	if code == nil {
		return true
	}
	zero := make([]byte, 32)
	return bytes.Equal(code.data[:], zero)
}

// Equal returns true if the specified asset code is the same value.
func (code *AssetCode) Equal(other AssetCode) bool {
	return bytes.Equal(code.data[:], other.data[:])
}

// AssetCodeFromContract generates a "unique" deterministic asset code from a contract public key
//   hash and an asset index.
func AssetCodeFromContract(contractAddress bitcoin.RawAddress, index uint64) *AssetCode {
	hash256 := sha256.New()
	hash256.Write(contractAddress.Bytes())
	binary.Write(hash256, DefaultEndian, &index)
	hash := hash256.Sum(nil)

	var result AssetCode
	copy(result.data[:], hash)
	return &result
}

// AssetCodeFromBytes returns a AssetCode with the specified bytes.
func AssetCodeFromBytes(data []byte) *AssetCode {
	var result AssetCode
	copy(result.data[:], data)
	return &result
}

func DeserializeAssetCode(buf *bytes.Reader) (*AssetCode, error) {
	var result AssetCode
	if _, err := buf.Read(result.data[:]); err != nil {
		return nil, err
	}
	return &result, nil
}

// Bytes returns a byte slice containing the AssetCode.
func (code *AssetCode) Bytes() []byte {
	if code == nil {
		return nil
	}
	return code.data[:]
}

// String converts to a string
func (code *AssetCode) String() string {
	return fmt.Sprintf("%x", code.data[:])
}

// Serialize serializes an asset code into a buffer.
func (code *AssetCode) Serialize(buf *bytes.Buffer) error {
	_, err := buf.Write(code.data[:])
	return err
}

// MarshalJSON converts to json.
func (code *AssetCode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%x\"", code.data)), nil
}

// UnmarshalJSON converts from json.
func (code *AssetCode) UnmarshalJSON(data []byte) error {
	if len(data) < 2 {
		return fmt.Errorf("Too short for AssetCode hex data : %d", len(data))
	}
	n, err := hex.Decode(code.data[:], data[1:len(data)-1])
	if err != nil {
		return err
	}
	if n != 32 {
		return fmt.Errorf("Invalid AssetCode size : %d", n)
	}
	return nil
}

// ------------------------------------------------------------------------------------------------
// ContractCode represents a unique identifier for a Tokenized static contract.
type ContractCode struct {
	data [32]byte
}

// Validate returns an error if the value is invalid
func (code *ContractCode) Validate() error {
	return nil
}

// IsZero returns true if the ContractCode is all zeroes. (empty)
func (code *ContractCode) IsZero() bool {
	if code == nil {
		return true
	}
	zero := make([]byte, 32, 32)
	return bytes.Equal(code.data[:], zero)
}

// Equal returns true if the specified values are the same.
func (code *ContractCode) Equal(other ContractCode) bool {
	return bytes.Equal(code.data[:], other.data[:])
}

// AssetCodeFromBytes returns a ContractCode with the specified bytes.
func ContractCodeFromBytes(data []byte) *ContractCode {
	var result ContractCode
	copy(result.data[:], data)
	return &result
}

func DeserializeContractCode(buf *bytes.Reader) (*ContractCode, error) {
	var result ContractCode
	if _, err := buf.Read(result.data[:]); err != nil {
		return nil, err
	}
	return &result, nil
}

// Bytes returns a byte slice containing the ContractCode.
func (code *ContractCode) Bytes() []byte {
	if code == nil {
		return nil
	}
	return code.data[:]
}

// String converts to a string
func (code *ContractCode) String() string {
	return fmt.Sprintf("%x", code.data[:])
}

// Serialize serializes a contract code into a buffer.
func (code *ContractCode) Serialize(buf *bytes.Buffer) error {
	_, err := buf.Write(code.data[:])
	return err
}

// MarshalJSON converts to json.
func (code *ContractCode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%x\"", code.data)), nil
}

// UnmarshalJSON converts from json.
func (code *ContractCode) UnmarshalJSON(data []byte) error {
	if len(data) < 2 {
		return fmt.Errorf("Too short for ContractCode hex data : %d", len(data))
	}
	n, err := hex.Decode(code.data[:], data[1:len(data)-1])
	if err != nil {
		return err
	}
	if n != 32 {
		return fmt.Errorf("Invalid ContractCode size : %d", n)
	}
	return nil
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
