# Tokenized Protocol Specification

Reference for the Tokenized Protocol structure. Definitions are stored in YAML which is used to generate protobuf to be compiled to other languages. Definitions are grouped by their purpose and are versioned within.

## Supported Languages

- Go
- Python
- Markdown
- Typescript/Javascript
- _your language here_ (contributions welcome)

## Getting Started

To build the project from source, clone the GitHub repo and in a command line terminal.

    mkdir -p $GOPATH/src/github.com/tokenized
    cd $GOPATH/src/github.com/tokenized
    git clone https://github.com/tokenized/specification.git
    cd specification

Navigate to the root directory and run:

    make

## Project Structure

- `cmd/tokenized` - Command line interface
- `src` - Source files for protocol messages (see Source Definitions)
- `dist` - Latest library for each language
- `internal` - Compiler logic and templates

## Source Definitions

The Tokenized Protocol must retain the entire definition history in order to parse legacy messages from prior versions. The working version for each definition is `develop`. Previous versions are stored in a directory that specifies that release version, for example, `v0001`.

## Protocol

### Serialization

This protocol's main purpose is to provide the ability to parse and create Bitcoin OP_RETURN output scripts containing Tokenized data in a consistent way.
There are several classes of Tokenized data including actions and messages. Some of the types have different types of payloads that further specify data.
To create an OP_RETURN script, first populate a action/message struct, then pass it into the Serialize function, which will return a Bitcoin script.
To read an OP_RETURN script, pass the script into the Deserialize function and if it is valid, then it will return an action/message object.

#### OP_RETURN script format
The OP_RETURN script is encoded using the Tokenized [envelope](https://github.com/tokenized/envelope) system. The Tokenized protocol uses the protocol identifiers `tokenized` and `test.tokenized`. The envelope `PayloadIdentifier` specifies which type of Tokenized message is contained. The envelope `Payload` is encoded using [Protocol Buffers](https://developers.google.com/protocol-buffers/).

### Primitive Types

* `int` is a signed integer. `size` is the number of bytes for the integer. Valid values for `size` are 1, 2, 4, and 8.

* `uint` is an unsigned integer. `size` is the number of bytes for the integer. Valid values for `size` are 1, 2, 4, and 8.

* `float` is a floating point number. `size` is the number of bytes for the float. Valid values for `size` are 4 and 8.

* `bool` is a boolean stored as 1 byte.

* `bin` is fixed length binary data. `size` is the length in bytes of the data.

* `varbin` is variable length binary data.
`varSize` defines the maximum size of the data as defined in `Variable Sizes`.

* `fixedchar` is fixed length text data.
The data is assumed to be UTF-8 unless the first two bytes are a valid UTF-16 BOM (Byte Order Method).
`size` is the length in bytes of the data.

* `varchar` is variable length text data.
The data is assumed to be UTF-8 unless the first two bytes are a valid UTF-16 BOM (Byte Order Method).
`varSize` defines the maximum size of the data as defined in `Variable Sizes`.

## List Types

Fields within the Tokenized protocol can be defined as a list of a specified data type.
This is done by adding a `[]` to the end of the `type` value.
The maximum number of elements in a list are defined by `listSize` as defined in Variable/List Sizes.

### Variable/List Sizes

Fields within the Tokenized protocol can be lists of objects or variable length objects.
The maximum number of elements in a list are defined by `listSize`.
The maximum size of a variable sized object are defined by `varSize`.
The valid values for `listSize` and `varSize` are as follows.
If no value is specified for `listSize` or `varSize` then `tiny` is used.

* `tiny` defines a max size/length of 2^8^-1 or 255.

* `small` defines a max size/length of 2^16^-1 or 65,535.

* `medium` defines a max size/length of 2^32^-1 or 4,294,967,295.

* `large` defines a max size/length of 2^64^-1 or approximately 1.8x10^19^.

### Basic Types

* `Action` implements a base interface for all message types. Provides the ability to serialize the data as a Bitcoin OP_RETURN script.

* `Asset` is the interface for asset payloads. Provides the ability to serialize asset data.

* `Message` is the interface for message payloads. Provides the ability to serialize message data.

* `OpReturnMessage` implements a base interface for all message types.
Provides the ability to serialize the data as a Bitcoin OP_RETURN script and to request the variable payload data.

* `PayloadMessage` is the interface for messages that are derived from payloads, such as asset types and message types.
Provides the ability to serialize the data.

* `TxId` represents a Bitcoin transaction ID.
It is the double SHA256 hash of the serialized transaction.
`size` does not need to be specified and is always 32 bytes.

* `PublicKeyHash` represents a Bitcoin Public Key Hash.
It is the RIPEMD160 hash of the SHA256 of the serialized public key.
Public key hashes are used as an "address" to send/receive transactions, tokens, and messages.
`size` does not need to be specified and is always 20 bytes.

* `AssetCode` represents a unique identifier for an asset/token.
`size` does not need to be specified and is always 32 bytes.

* `ContractCode` represents a unique identifier for a static contract.
`size` does not need to be specified and is always 32 bytes.

* `Timestamp` represents a time.
`size` does not need to be specified and is always 8 bytes.

* `Polity` represents a unique identifier for a nation/state/political entity.
`size` does not need to be specified and is always 3 bytes.

### Arrays/Lists

Arrays/Lists of objects are defined by adding `[]` at the end of the type field.
Arrays/Lists are variable length and have an unsigned integer encoded before the objects to specify the number of items.
`size` for lists represents the bits used to serialize the number of items in the list. Valid values are 1, 2, 4, and 8.
The default, when size is not specified, is 1. A size of 1 means the list can contain a maximum 255 objects.

### Actions

**Actions** are Tokenized messages that request or confirm the creation, modification, or transfer of an object within the Tokenized protocol.
Actions must be serialized as an OP_RETURN output in a Bitcoin transaction with Bitcoin transaction inputs and outputs from/to the appropriate parties.

### Assets

All asset payloads for the Asset related actions.

### Messages

All message payloads for the Message related message actions.

### Resources

General resources used by other actions.

# License

Copyright 2018-2019 nChain Holdings Limited.
