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

// OrderAuthoritySigHash returns a Double SHA256 of the data in an order op return that needs to
//   be signed by an authority for the OrderSignature field.
// The Order must be fully populated including the AuthorityName.
func OrderAuthoritySigHash(ctx context.Context, contractAddress bitcoin.RawAddress,
	order *actions.Order) (*bitcoin.Hash32, error) {

	// Calculate the hash
	digest := sha256.New()

	digest.Write([]byte(order.AuthorityName))
	digest.Write(contractAddress.Bytes())
	digest.Write([]byte{byte(order.ComplianceAction[0])})
	digest.Write(order.AssetCode)
	digest.Write(order.SupportingEvidence[:])

	for _, refTx := range order.ReferenceTransactions {
		data, err := proto.Marshal(refTx)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to serialize reference transaction")
		}
		digest.Write(data)
	}

	err := binary.Write(digest, binary.LittleEndian, order.FreezePeriod)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to serialize freeze timeout")
	}

	for _, target := range order.TargetAddresses {
		data, err := proto.Marshal(target)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to serialize target address")
		}
		digest.Write(data)
	}

	digest.Write(order.DepositAddress)

	hash := bitcoin.Hash32(sha256.Sum256(digest.Sum(nil)))
	return &hash, nil
}
