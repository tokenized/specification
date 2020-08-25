package protocol

import (
	"context"
	"crypto/sha256"
	"encoding/binary"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/specification/dist/golang/actions"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

// TransferOracleSigHash returns a Double SHA256 of the data required to identify a
//   token receiver approval by an oracle.
// approved = 1 - means approved. any other value is a signature for rejecting approval.
// The block hash of the chain tip - 4 should be used. The signature will be considered valid
//   until 1 hour past the timestamp of the block after the block hash specified (chain tip).
func TransferOracleSigHash(ctx context.Context, contractAddress bitcoin.RawAddress,
	assetCode []byte, receiverAddress bitcoin.RawAddress, blockHash bitcoin.Hash32,
	expiration uint64, approved uint8) ([]byte, error) {

	// Calculate the hash
	digest := sha256.New()

	digest.Write(receiverAddress.Bytes())
	digest.Write(contractAddress.Bytes())
	digest.Write(assetCode)
	digest.Write(blockHash[:])
	binary.Write(digest, DefaultEndian, &expiration)

	binary.Write(digest, DefaultEndian, &approved)

	hash := sha256.Sum256(digest.Sum(nil))
	return hash[:], nil
}

// ContractAdminIdentityOracleSigHash returns a Double SHA256 of the data required to identify a
//   contract issuer and related information by an oracle.
// entities parameter must contain either *actions.EntityField or bitcoin.RawAddress
// approved = 1 - means approved. any other value is a signature for rejecting approval.
// The block hash of the chain tip - 4 should be used. The signature will be considered valid
//   until 1 hour past the timestamp of the block after the block hash specified (chain tip).
//
// The admin address is written into the hash, then the operator address if one exists.
// When a contract provides the EntityContract address instead of an Issuer entity, then that
// address is written into the hash instead of the Issuer entity.
// If there is an operator then the OperatorEntityContract address is written into the hash.
func ContractAdminIdentityOracleSigHash(ctx context.Context, adminAddress bitcoin.RawAddress,
	entity interface{}, blockHash bitcoin.Hash32, expiration uint64, approved uint8) ([]byte, error) {

	// Calculate the hash
	digest := sha256.New()

	digest.Write(adminAddress.Bytes())
	switch e := entity.(type) {
	case *actions.EntityField:
		// TODO There is an issue with protobuf not being deterministic. The marshalled data may not
		// always be marshalled the same, so the signature hash may not match. --ce
		data, err := proto.Marshal(e)
		if err != nil {
			return nil, errors.Wrap(err, "serialize entity")
		}
		digest.Write(data)
	case bitcoin.RawAddress:
		if err := e.Serialize(digest); err != nil {
			return nil, errors.Wrap(err, "serialize entity raw address")
		}
	}
	digest.Write(blockHash[:])
	binary.Write(digest, DefaultEndian, &expiration)

	binary.Write(digest, DefaultEndian, &approved)

	hash := sha256.Sum256(digest.Sum(nil))
	return hash[:], nil
}

// EntityPubKeyOracleSigHash returns a Double SHA256 of the data required to verify an association
//   between an entity and a public key by an oracle.
// approved = 1 - means approved. any other value is a signature for rejecting approval.
// The block hash of the chain tip - 4 should be used. This gives a timestamp to the signature.
func EntityPubKeyOracleSigHash(ctx context.Context, entity *actions.EntityField,
	pubKey bitcoin.PublicKey, blockHash bitcoin.Hash32, approved uint8) ([]byte, error) {

	// Calculate the hash
	digest := sha256.New()

	data, err := proto.Marshal(entity)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to serialize entity")
	}
	digest.Write(data)
	digest.Write(pubKey.Bytes())
	digest.Write(blockHash[:])

	binary.Write(digest, DefaultEndian, &approved)

	hash := sha256.Sum256(digest.Sum(nil))
	return hash[:], nil
}

// EntityXPubOracleSigHash returns a Double SHA256 of the data required to verify an association
//   between an entity and an extended public key by an oracle.
// approved = 1 - means approved. any other value is a signature for rejecting approval.
// The block hash of the chain tip - 4 should be used. This gives a timestamp to the signature.
func EntityXPubOracleSigHash(ctx context.Context, entity *actions.EntityField,
	xpub bitcoin.ExtendedKeys, blockHash bitcoin.Hash32, approved uint8) ([]byte, error) {

	// Calculate the hash
	digest := sha256.New()

	// TODO There is an issue with protobuf not being deterministic. The marshalled data may not
	// always be marshalled the same, so the signature hash may not match. --ce
	data, err := proto.Marshal(entity)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to serialize entity")
	}
	digest.Write(data)
	digest.Write(xpub.Bytes())
	digest.Write(blockHash[:])

	binary.Write(digest, DefaultEndian, &approved)

	hash := sha256.Sum256(digest.Sum(nil))
	return hash[:], nil
}
