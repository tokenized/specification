# Tokenized Protocol Specification

Reference for the Tokenized Protocol structure. Definitions are stored in YAML to be parsed and compiled to other languages. Definitions are grouped by their purpose and are versioned within.

## Supported Languages

- Go
- Python
- Markdown
- _your language here_ (contributions welcome)

## Getting Started

To build the project from source, clone the GitHub repo and in a command line terminal.

    mkdir -p $GOPATH/src/github.com/tokenized
    cd $GOPATH/src/github.com/tokenized
    git clone git@github.com:tokenized/specification
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

### Basic Types

`OpReturnMessage` implements a base interface for all message types.  
Provides the ability to serialize the data as a Bitcoin OP_RETURN script and to request the variable payload data.

`PayloadMessage` is the interface for messages that are derived from payloads, such as asset types and message types.  
Provides the ability to serialize the data.

`TxId` represents a Bitcoin transaction ID.  
It is the double SHA256 hash of the serialized transaction.

`PublicKeyHash` represents a Bitcoin Public Key Hash.  
It is the RIPEMD160 hash of the SHA256 of the serialized public key. Public key hashes are used as an "address" to send/receive transactions, tokens, and messages.

`AssetCode` represents a unique identifier for an asset/token.

`Timestamp` represents a time.

### Actions

**Actions** are Tokenized messages that request or confirm the creation, modification, or transfer of an object within the Tokenized protocol.
Actions must be serialized as an OP_RETURN output in a Bitcoin transaction with Bitcoin transaction inputs and outputs from/to the appropriate parties.

### Assets

All asset payloads for the Asset related actions.

### Resources

All resources and message payloads for message related actions.

# License

Copyright 2018-2019 nChain Holdings Limited.
