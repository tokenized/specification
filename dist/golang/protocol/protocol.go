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
	"github.com/tokenized/envelope/pkg/golang/envelope/base"
	envelopeBase "github.com/tokenized/envelope/pkg/golang/envelope/base"
	envelopeV0 "github.com/tokenized/envelope/pkg/golang/envelope/v0"
	envelopeV1 "github.com/tokenized/envelope/pkg/golang/envelope/v1"
	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/pkg/bsor"
	"github.com/tokenized/pkg/json"
	"github.com/tokenized/specification/dist/golang/actions"
	"github.com/tokenized/specification/dist/golang/instruments"
	"github.com/tokenized/specification/dist/golang/messages"
	actionsv0 "github.com/tokenized/specification/dist/golang/v0/actions"
	instrumentsv0 "github.com/tokenized/specification/dist/golang/v0/instruments"
	messagesv0 "github.com/tokenized/specification/dist/golang/v0/messages"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ripemd160"
)

const (
	// Version of the Tokenized protocol.
	Version = uint64(1)

	FlagVersion = uint64(0)

	InstrumentCodeSize = 20

	ContractCodeSize = 32

	BSVInstrumentID = "BSV"
)

var (
	// ProtocolID is the current protocol ID
	ProtocolID = envelopeBase.ProtocolID("TKN")

	// TestProtocolID is the current protocol ID for testing.
	TestProtocolID = envelopeBase.ProtocolID("test.TKN")

	FlagProtocolID = envelopeBase.ProtocolID("flag")

	ErrNotTokenized   = errors.New("Not Tokenized")
	ErrNotFlag        = errors.New("Not Flag")
	ErrUnknownVersion = errors.New("Unknown Version")

	// DefaultEndian specifies the order of bytes for encoding integers.
	DefaultEndian = binary.LittleEndian
)

func GetProtocolID(isTest bool) envelopeBase.ProtocolID {
	if isTest {
		return TestProtocolID
	} else {
		return ProtocolID
	}
}

// Serialize serializes an action into a Tokenized OP_RETURN script.
func Serialize(action actions.Action, isTest bool) (bitcoin.Script, error) {
	scriptItems, err := WrapAction(action, isTest)
	if err != nil {
		return nil, errors.Wrap(err, "wrap action")
	}

	return scriptItems.Script()
}

// DeserializeInstrumentPayload decodes an instrument payload. Use InstrumentPayloadVersion, which
// is set by the decoder, for version to specify which encoding is used.
func DeserializeInstrumentPayload(version uint8, code []byte,
	payload []byte) (instruments.Instrument, error) {

	if version == 0 {
		payloadv0, err := instrumentsv0.DeserializeV0(code, payload)
		if err != nil {
			return nil, err
		}

		return convertInstrumentV0toV1(code, payloadv0)
	}

	if version == 1 {
		return instruments.DeserializeV1(code, payload)
	}

	return nil, errors.Wrap(ErrUnknownVersion, "instrument payload")
}

// convertInstrumentV0toV1 is needed to read the old protobuf structures into the new bsor
// structures. It requires the json field names to match and be the same type.
func convertInstrumentV0toV1(code []byte,
	instrumentv0 instrumentsv0.Instrument) (instruments.Instrument, error) {

	js, err := json.Marshal(instrumentv0)
	if err != nil {
		return nil, errors.Wrap(err, "json marshal")
	}

	instrumentv1 := instruments.NewInstrumentFromCode(string(code))
	if instrumentv1 == nil {
		return nil, fmt.Errorf("No action for code: %s", string(code))
	}

	if err := json.Unmarshal(js, instrumentv1); err != nil {
		return nil, errors.Wrap(err, "json unmarshal")
	}

	return instrumentv1, nil
}

// DeserializeMessagePayload decodes a message payload. Use MessagePayloadVersion, which is set by
// the decoder, for version to specify which encoding is used.
func DeserializeMessagePayload(version uint8, code uint16,
	payload []byte) (messages.Message, error) {

	if version == 0 {
		payloadv0, err := messagesv0.DeserializeV0(uint32(code), payload)
		if err != nil {
			return nil, err
		}

		return convertMessageV0toV1(code, payloadv0)
	}

	if version == 1 {
		return messages.DeserializeV1(code, payload)
	}

	return nil, errors.Wrap(ErrUnknownVersion, "instrument payload")
}

// convertMessageV0toV1 is needed to read the old protobuf structures into the new bsor
// structures. It requires the json field names to match and be the same type.
func convertMessageV0toV1(code uint16,
	messagev0 messagesv0.Message) (messages.Message, error) {

	js, err := json.Marshal(messagev0)
	if err != nil {
		return nil, errors.Wrap(err, "json marshal")
	}

	messagev1 := messages.NewMessageFromCode(code)
	if messagev1 == nil {
		return nil, fmt.Errorf("No action for code: %d", code)
	}

	if err := json.Unmarshal(js, messagev1); err != nil {
		return nil, errors.Wrap(err, "json unmarshal")
	}

	return messagev1, nil
}

// Deserialize reads an action from a Tokenized OP_RETURN script.
func Deserialize(script bitcoin.Script, isTest bool) (actions.Action, error) {
	r := bytes.NewReader(script)

	envelopeVersion, err := envelopeBase.ParseHeader(r)
	if err == envelopeBase.ErrNotEnvelope {
		return nil, ErrNotTokenized
	} else if err != nil {
		return nil, err
	}

	if envelopeVersion == 0 {
		return parseEnvelopeV0(r, isTest)
	}

	if envelopeVersion == 1 {
		return parseEnvelopeV1(r, isTest)
	}

	return nil, errors.Wrap(ErrUnknownVersion, "envelope")
}

func parseEnvelopeV0(r *bytes.Reader, isTest bool) (actions.Action, error) {
	msg, err := envelopeV0.Deserialize(r)
	if err == envelopeBase.ErrNotEnvelope {
		return nil, ErrNotTokenized
	} else if err != nil {
		return nil, errors.Wrap(err, "envelope v0: parse")
	}

	if !bytes.Equal(msg.PayloadProtocol(), GetProtocolID(isTest)) {
		return nil, ErrNotTokenized
	}

	version := msg.PayloadVersion()
	if version == 0 {
		return parseV0(msg.PayloadIdentifier(), msg.Payload())
	}

	return nil, errors.Wrapf(ErrUnknownVersion, "envelope v0: %d", version)
}

func parseEnvelopeV1(r *bytes.Reader, isTest bool) (actions.Action, error) {
	envelopeData, err := envelopeV1.ParseAfterHeader(r)
	if err != nil {
		return nil, errors.Wrap(err, "envelope v1: protocol ids")
	}

	if len(envelopeData.ProtocolIDs) != 1 ||
		!bytes.Equal(envelopeData.ProtocolIDs[0], GetProtocolID(isTest)) {

		return nil, errors.Wrap(ErrNotTokenized, "envelope v1: protocol id")
	}

	if len(envelopeData.Payload) < 2 {
		return nil, errors.Wrap(ErrNotTokenized, "envelope v1: not enough payloads")
	}

	version, err := bitcoin.ScriptNumberValueUnsigned(envelopeData.Payload[0])
	if err != nil {
		return nil, errors.Wrap(err, "envelope v1: version")
	}

	if version == 0 {
		code := []byte(envelopeData.Payload[1].Data)

		var payload []byte
		if len(envelopeData.Payload) > 2 {
			payload = envelopeData.Payload[2].Data
		}

		return parseV0(code, payload)
	}

	if version == 1 {
		return parseV1(envelopeData.Payload[1:])
	}

	return nil, errors.Wrapf(ErrUnknownVersion, "envelope v1: %d", version)
}

func parseV0(code []byte, payload []byte) (actions.Action, error) {
	actionv0, err := actionsv0.DeserializeV0(code, payload)
	if err != nil {
		return nil, errors.Wrap(err, "deserialize v0")
	}

	// Convert to v1 action
	return convertV0toV1(actionv0)
}

func parseV1(payload bitcoin.ScriptItems) (actions.Action, error) {
	return actions.DeserializeV1(payload)
}

// convertV0toV1 is needed to read the old protobuf structures into the new bsor structures. It
// requires the json field names to match and be the same type.
func convertV0toV1(actionv0 actionsv0.Action) (actions.Action, error) {
	js, err := json.Marshal(actionv0)
	if err != nil {
		return nil, errors.Wrap(err, "json marshal")
	}

	actionv1 := actions.NewActionFromCode(actionv0.Code())
	if actionv1 == nil {
		return nil, fmt.Errorf("No action for code: %s", actionv0.Code())
	}

	if err := json.Unmarshal(js, actionv1); err != nil {
		return nil, errors.Wrap(err, "json unmarshal")
	}

	return actionv1, nil
}

func ActionCodeForScript(script bitcoin.Script, isTest bool) (string, error) {
	r := bytes.NewReader(script)

	envelopeVersion, err := envelopeBase.ParseHeader(r)
	if err == envelopeBase.ErrNotEnvelope {
		return "", ErrNotTokenized
	} else if err != nil {
		return "", err
	}

	if envelopeVersion == 0 {
		msg, err := envelopeV0.Deserialize(r)
		if err == envelopeBase.ErrNotEnvelope {
			return "", ErrNotTokenized
		} else if err != nil {
			return "", errors.Wrap(err, "envelope v0: parse")
		}

		if !bytes.Equal(msg.PayloadProtocol(), GetProtocolID(isTest)) {
			return "", ErrNotTokenized
		}

		version := msg.PayloadVersion()

		if version == 0 {
			return string(msg.PayloadIdentifier()), nil
		}

		return "", errors.Wrapf(ErrUnknownVersion, "envelope v0: %d", version)
	}

	if envelopeVersion == 1 {
		envelopeData, err := envelopeV1.ParseAfterHeader(r)
		if err != nil {
			return "", errors.Wrap(err, "envelope v1: protocol ids")
		}

		if len(envelopeData.ProtocolIDs) != 1 ||
			!bytes.Equal(envelopeData.ProtocolIDs[0], GetProtocolID(isTest)) {
			return "", errors.Wrap(ErrNotTokenized, "envelope v1: protocol id")
		}

		if len(envelopeData.Payload) < 2 {
			return "", errors.Wrap(ErrNotTokenized, "envelope v1: not enough payloads")
		}

		version, err := bitcoin.ScriptNumberValueUnsigned(envelopeData.Payload[0])
		if err != nil {
			return "", errors.Wrap(err, "envelope v1: version")
		}

		if version == 0 {
			return string([]byte(envelopeData.Payload[1].Data)), nil
		}

		if version == 1 {
			return string([]byte(envelopeData.Payload[1].Data)), nil
		}
	}

	return "", errors.Wrap(ErrUnknownVersion, "envelope")
}

// WrapAction wraps an action in an envelope message.
func WrapAction(action actions.Action, isTest bool) (bitcoin.ScriptItems, error) {
	scriptItems := envelopeV1.HeaderScriptItems(base.ProtocolIDs{GetProtocolID(isTest)})

	payloadScriptItems, err := bsor.Marshal(action)
	if err != nil {
		return nil, errors.Wrap(err, "serialize action")
	}

	scriptItems = append(scriptItems, bitcoin.PushNumberScriptItem(int64(len(payloadScriptItems)+2)))

	scriptItems = append(scriptItems, bitcoin.PushNumberScriptItem(int64(Version)))

	scriptItems = append(scriptItems, bitcoin.NewPushDataScriptItem([]byte(action.Code())))

	return append(scriptItems, payloadScriptItems...), nil
}

// SerializeFlagOutputScript creates a locking script containing the flag value for a relationship
//   message.
func SerializeFlagOutputScript(flag []byte) (bitcoin.Script, error) {
	message := envelopeV1.NewMessage(base.ProtocolIDs{FlagProtocolID},
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
	if err == envelopeBase.ErrNotEnvelope {
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
