package protocol

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/tokenized/envelope/pkg/golang/envelope"
	envelopeV0 "github.com/tokenized/envelope/pkg/golang/envelope/v0"
	envelopeV1 "github.com/tokenized/envelope/pkg/golang/envelope/v1"
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

	InstrumentCodeSize = 20

	ContractCodeSize = 32

	BSVInstrumentID = "BSV"
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
func Serialize(action actions.Action, isTest bool) (bitcoin.Script, error) {
	message, err := WrapAction(action, isTest)
	if err != nil {
		return nil, errors.Wrap(err, "wrap action")
	}

	var buf bytes.Buffer
	err = message.Serialize(&buf)
	if err != nil {
		return nil, errors.Wrap(err, "serialize envelope")
	}

	return bitcoin.Script(buf.Bytes()), nil
}

// Deserialize reads an action from a Tokenized OP_RETURN script.
func Deserialize(script bitcoin.Script, isTest bool) (actions.Action, error) {
	buf := bytes.NewReader(script)
	message, err := envelope.Deserialize(buf)
	if err == envelope.ErrNotEnvelope {
		return nil, ErrNotTokenized
	} else if err != nil {
		return nil, err
	}

	switch msg := message.(type) {
	case *envelopeV0.Message:
		if !bytes.Equal(msg.PayloadProtocol(), GetProtocolID(isTest)) {
			return nil, ErrNotTokenized
		}

		if msg.PayloadVersion() != Version {
			return nil, ErrUnknownVersion
		}

		return actions.Deserialize(msg.PayloadIdentifier(), msg.Payload())

	default:
		protocols := msg.PayloadProtocols()
		if len(protocols) != 1 {
			return nil, errors.Wrapf(ErrNotTokenized, "wrong protocol count: %d", len(protocols))
		}

		if !bytes.Equal(protocols[0], GetProtocolID(isTest)) {
			return nil, ErrNotTokenized
		}

		if msg.PayloadCount() != 3 {
			return nil, errors.Wrapf(ErrNotTokenized, "wrong payload count: %d",
				msg.PayloadCount())
		}

		version, _, err := bitcoin.ParsePushNumberScript(msg.PayloadAt(0))
		if err != nil {
			return nil, errors.Wrap(ErrNotTokenized, errors.Wrap(err, "version number").Error())
		}
		if version < 0 || uint64(version) != Version {
			return nil, ErrUnknownVersion
		}

		return actions.Deserialize(msg.PayloadAt(1), msg.PayloadAt(2))
	}
}

func ActionCodeForScript(script bitcoin.Script, isTest bool) (string, error) {
	buf := bytes.NewReader(script)
	message, err := envelope.Deserialize(buf)
	if err == envelope.ErrNotEnvelope {
		return "", ErrNotTokenized
	} else if err != nil {
		return "", err
	}

	switch msg := message.(type) {
	case *envelopeV0.Message:
		if !bytes.Equal(msg.PayloadProtocol(), GetProtocolID(isTest)) {
			return "", ErrNotTokenized
		}

		if msg.PayloadVersion() != Version {
			return "", ErrUnknownVersion
		}

		return string(msg.PayloadIdentifier()), nil

	default:
		protocols := msg.PayloadProtocols()
		if len(protocols) != 1 {
			return "", errors.Wrapf(ErrNotTokenized, "wrong protocol count: %d", len(protocols))
		}

		if !bytes.Equal(protocols[0], GetProtocolID(isTest)) {
			return "", ErrNotTokenized
		}

		if msg.PayloadCount() != 3 {
			return "", errors.Wrapf(ErrNotTokenized, "wrong payload count: %d",
				msg.PayloadCount())
		}

		version, _, err := bitcoin.ParsePushNumberScript(msg.PayloadAt(0))
		if err != nil {
			return "", errors.Wrap(ErrNotTokenized, errors.Wrap(err, "version number").Error())
		}
		if version < 0 || uint64(version) != Version {
			return "", ErrUnknownVersion
		}

		return string(msg.PayloadAt(1)), nil
	}
}

// WrapAction wraps an action in an envelope message.
func WrapAction(action actions.Action, isTest bool) (envelope.BaseMessage, error) {
	payload, err := proto.Marshal(action)
	if err != nil {
		return nil, errors.Wrap(err, "serialize action")
	}

	message := envelopeV1.NewMessage([][]byte{GetProtocolID(isTest)},
		[][]byte{bitcoin.PushNumberScript(int64(Version)), []byte(action.Code()), payload})

	return message, nil
}

// SerializeFlagOutputScript creates a locking script containing the flag value for a relationship
//   message.
func SerializeFlagOutputScript(flag []byte) (bitcoin.Script, error) {
	message := envelopeV1.NewMessage([][]byte{[]byte(FlagProtocolID)},
		[][]byte{bitcoin.PushNumberScript(int64(FlagVersion)), flag})

	var buf bytes.Buffer
	if err := message.Serialize(&buf); err != nil {
		return nil, errors.Wrap(err, "Failed to serialize flag envelope")
	}
	return bitcoin.Script(buf.Bytes()), nil
}

// DeserializeFlagOutputScript returns a flag value if the script is a flag script.
func DeserializeFlagOutputScript(script bitcoin.Script) ([]byte, error) {
	buf := bytes.NewReader(script)
	message, err := envelope.Deserialize(buf)
	if err == envelope.ErrNotEnvelope {
		return nil, ErrNotFlag
	} else if err != nil {
		return nil, err
	}

	switch msg := message.(type) {
	case *envelopeV0.Message:
		if !bytes.Equal(msg.PayloadProtocol(), []byte(FlagProtocolID)) {
			return nil, ErrNotFlag
		}

		if msg.PayloadVersion() != FlagVersion {
			return nil, ErrUnknownVersion
		}

		return msg.Payload(), nil

	default:
		protocols := msg.PayloadProtocols()
		if len(protocols) != 1 {
			return nil, errors.Wrapf(ErrNotFlag, "wrong protocol count: %d", len(protocols))
		}

		if !bytes.Equal(protocols[0], []byte(FlagProtocolID)) {
			return nil, ErrNotFlag
		}

		if msg.PayloadCount() != 2 {
			return nil, errors.Wrapf(ErrNotFlag, "wrong payload count: %d",
				msg.PayloadCount())
		}

		version, _, err := bitcoin.ParsePushNumberScript(msg.PayloadAt(0))
		if err != nil {
			return nil, errors.Wrap(ErrNotFlag, errors.Wrap(err, "version number").Error())
		}
		if version < 0 || uint64(version) != FlagVersion {
			return nil, ErrUnknownVersion
		}

		return message.PayloadAt(1), nil
	}
}

// InstrumentCodeFromContract generates a "unique" deterministic instrument code from a contract public key
//   hash and an instrument index.
func InstrumentCodeFromContract(contractAddress bitcoin.RawAddress, index uint64) bitcoin.Hash20 {
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

// InstrumentCodeFromBytes returns a InstrumentCode with the specified bytes.
func InstrumentCodeFromBytes(b []byte) bitcoin.Hash20 {
	var result bitcoin.Hash20
	copy(result[:], b)
	return result
}

// InstrumentID encodes an instrument ID.
//
// InstrumentID = InstrumentType(3 characters) + base58(InstrumentCode + checksum)
//
// There is a special case for BSV, which will be returned as BSV.
func InstrumentID(instrumentType string, code bitcoin.Hash20) string {
	if instrumentType == BSVInstrumentID {
		return instrumentType
	}

	b := code.Bytes()

	// Perform Double SHA-256 hash
	checksum := bitcoin.DoubleSha256(b)

	// Append the first 4 checksum bytes
	b = append(b, checksum[:4]...)

	// Prepend with type and encode as text with base 58.
	return instrumentType + bitcoin.Base58(b)
}

// InstrumentIDForRaw returns the instrument ID for an instrument type and instrument code in byte slice form.
func InstrumentIDForRaw(instrumentType string, instrumentCode []byte) (string, error) {
	if instrumentType == BSVInstrumentID {
		return instrumentType, nil
	}

	code, err := bitcoin.NewHash20(instrumentCode)
	if err != nil {
		return "", errors.Wrap(err, "instrument code")
	}

	return InstrumentID(instrumentType, *code), nil
}

func InstrumentIDForTransfer(f *actions.InstrumentTransferField) (string, error) {
	return InstrumentIDForRaw(f.InstrumentType, f.InstrumentCode)
}

func InstrumentIDForSettlement(f *actions.InstrumentSettlementField) (string, error) {
	return InstrumentIDForRaw(f.InstrumentType, f.InstrumentCode)
}

// DecodeInstrumentID decodes the instrument type and instrument code from an instrument ID.
func DecodeInstrumentID(id string) (string, bitcoin.Hash20, error) {
	if id == BSVInstrumentID {
		// Bitcoin instrument id. Instrument code all zeros.
		return BSVInstrumentID, bitcoin.Hash20{}, nil
	}

	if len(id) < InstrumentCodeSize+7 {
		return "", bitcoin.Hash20{}, fmt.Errorf("Instrument ID too short : %s", id)
	}

	instrumentType := id[:3]
	text := id[3:]

	b := bitcoin.Base58Decode(text)

	if len(b) != InstrumentCodeSize+4 {
		return "", bitcoin.Hash20{}, fmt.Errorf("Instrument code data wrong size : %s : got %d, want %d",
			id, len(b), InstrumentCodeSize+4)
	}

	// Verify checksum
	checksum := b[InstrumentCodeSize:]
	b = b[:InstrumentCodeSize]
	hash := bitcoin.DoubleSha256(b)
	if !bytes.Equal(hash[:4], checksum) {
		return "", bitcoin.Hash20{}, fmt.Errorf("Invalid Instrument ID checksum : %s", id)
	}

	code, err := bitcoin.NewHash20(b)
	if err != nil {
		return "", bitcoin.Hash20{}, errors.Wrap(err, "new hash")
	}

	return instrumentType, *code, nil
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

// InstrumentCodeFromBytes returns a ContractCode with the specified bytes.
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

func DeserializeTimestamp(r io.Reader) (Timestamp, error) {
	var result Timestamp
	if err := binary.Read(r, DefaultEndian, &result.nanoseconds); err != nil {
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
func (time *Timestamp) Serialize(w io.Writer) error {
	if err := binary.Write(w, DefaultEndian, &time.nanoseconds); err != nil {
		return err
	}
	return nil
}

// MarshalJSON converts to json.
func (time Timestamp) MarshalJSON() ([]byte, error) {
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
