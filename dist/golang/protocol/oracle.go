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
// The block hash of the chain tip - 4 should be used. The signature will be considered valid
//   until 1 hour past the timestamp of the block after the block hash specified (chain tip).
func TransferOracleSigHash(ctx context.Context, contractAddress bitcoin.RawAddress, assetCode []byte,
	receiverAddress bitcoin.RawAddress, quantity uint64, blockHash *bitcoin.Hash32, approved uint8) ([]byte, error) {

	// Calculate the hash
	digest := sha256.New()

	digest.Write(receiverAddress.Bytes())
	digest.Write(contractAddress.Bytes())
	digest.Write(assetCode)
	binary.Write(digest, DefaultEndian, &quantity)
	digest.Write(blockHash[:])

	binary.Write(digest, DefaultEndian, &approved)

	hash := sha256.Sum256(digest.Sum(nil))
	return hash[:], nil
}

// ContractOracleSigHash returns a Double SHA256 of the data required to identify a
//   contract creator by an oracle.
// The block hash of the chain tip - 4 should be used. The signature will be considered valid
//   until 1 hour past the timestamp of the block after the block hash specified (chain tip).
func ContractOracleSigHash(ctx context.Context, adminAddresses []bitcoin.RawAddress,
	entities []*actions.EntityField, blockHash *bitcoin.Hash32, approved uint8) ([]byte, error) {

	// Calculate the hash
	digest := sha256.New()

	for _, admin := range adminAddresses {
		digest.Write(admin.Bytes())
	}
	for _, entity := range entities {
		data, err := proto.Marshal(entity)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to serialize entity")
		}
		digest.Write(data)
	}
	digest.Write(blockHash[:])

	binary.Write(digest, DefaultEndian, &approved)

	hash := sha256.Sum256(digest.Sum(nil))
	return hash[:], nil
}

// EntityPubKeyOracleSigHash returns a Double SHA256 of the data required to verify an association
//   between an entity and a public key by an oracle.
// The block hash of the chain tip - 4 should be used. This gives a timestamp to the signature.
func EntityPubKeyOracleSigHash(ctx context.Context, entity *actions.EntityField,
	pubKey bitcoin.PublicKey, blockHash *bitcoin.Hash32, approved uint8) ([]byte, error) {

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
// The block hash of the chain tip - 4 should be used. This gives a timestamp to the signature.
func EntityXPubOracleSigHash(ctx context.Context, entity *actions.EntityField,
	xpub bitcoin.ExtendedKeys, blockHash *bitcoin.Hash32, approved uint8) ([]byte, error) {

	// Calculate the hash
	digest := sha256.New()

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
