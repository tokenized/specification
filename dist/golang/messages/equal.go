package messages

import (
	"bytes"

	proto "github.com/golang/protobuf/proto"
)

func (l *PublicMessage) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*PublicMessage)
	if !ok {
		return false
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	// Field Subject - varchar
	if l.Subject != r.Subject {
		return false // fmt.Errorf("Subject string mismatched")
	}

	// Field Regarding - Outpoint

	// Field PublicMessage - Document

	// Field Attachments - Document
	if len(l.Attachments) != len(r.Attachments) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.Attachments {
		if !v.Equal(r.Attachments[i]) {
			return false // fmt.Errorf("Attachments[%d] : %s", i, err)
		}
	}

	return true
}

func (l *PrivateMessage) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*PrivateMessage)
	if !ok {
		return false
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	// Field Subject - varchar
	if l.Subject != r.Subject {
		return false // fmt.Errorf("Subject string mismatched")
	}

	// Field Regarding - Outpoint

	// Field PrivateMessage - Document

	// Field Attachments - Document
	if len(l.Attachments) != len(r.Attachments) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.Attachments {
		if !v.Equal(r.Attachments[i]) {
			return false // fmt.Errorf("Attachments[%d] : %s", i, err)
		}
	}

	return true
}

func (l *RevertedTx) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*RevertedTx)
	if !ok {
		return false
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	// Field Transaction - varbin
	if !bytes.Equal(l.Transaction, r.Transaction) {
		return false // fmt.Errorf("Transaction bytes mismatched")
	}

	return true
}

func (l *Offer) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*Offer)
	if !ok {
		return false
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	// Field Payload - varbin
	if !bytes.Equal(l.Payload, r.Payload) {
		return false // fmt.Errorf("Payload bytes mismatched")
	}

	return true
}

func (l *SignatureRequest) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*SignatureRequest)
	if !ok {
		return false
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	// Field Payload - varbin
	if !bytes.Equal(l.Payload, r.Payload) {
		return false // fmt.Errorf("Payload bytes mismatched")
	}

	return true
}

func (l *SettlementRequest) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*SettlementRequest)
	if !ok {
		return false
	}

	// Field Timestamp - uint
	if l.Timestamp != r.Timestamp {
		return false // fmt.Errorf("Timestamp integer mismatched")
	}

	// Field TransferTxId - bin
	if !bytes.Equal(l.TransferTxId, r.TransferTxId) {
		return false // fmt.Errorf("TransferTxId bytes mismatched")
	}

	// Field ContractFees - TargetAddress
	if len(l.ContractFees) != len(r.ContractFees) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.ContractFees {
		if !v.Equal(r.ContractFees[i]) {
			return false // fmt.Errorf("ContractFees[%d] : %s", i, err)
		}
	}

	// Field Settlement - varbin
	if !bytes.Equal(l.Settlement, r.Settlement) {
		return false // fmt.Errorf("Settlement bytes mismatched")
	}

	return true
}

func (l *OutputMetadata) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*OutputMetadata)
	if !ok {
		return false
	}

	// Field OutputDescription - varchar
	if l.OutputDescription != r.OutputDescription {
		return false // fmt.Errorf("OutputDescription string mismatched")
	}

	// Field Tags - uint
	if len(l.Tags) != len(r.Tags) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.Tags {
		if v != r.Tags[i] {
			return false // fmt.Errorf("Element Tags integer mismatched")
		}
	}

	// Field CustomTags - OutputTag
	if len(l.CustomTags) != len(r.CustomTags) {
		return false // fmt.Errorf("List length mismatched")
	}
	for i, v := range l.CustomTags {
		if !v.Equal(r.CustomTags[i]) {
			return false // fmt.Errorf("CustomTags[%d] : %s", i, err)
		}
	}

	return true
}

func (l *OutpointField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*OutpointField)
	if !ok {
		return false
	}

	// Field TxId - bin
	if !bytes.Equal(l.TxId, r.TxId) {
		return false // fmt.Errorf("TxId bytes mismatched")
	}

	// Field OutputIndex - uint
	if l.OutputIndex != r.OutputIndex {
		return false // fmt.Errorf("OutputIndex integer mismatched")
	}

	return true
}

func (l *OutputTagField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*OutputTagField)
	if !ok {
		return false
	}

	// Field Tag - varchar
	if l.Tag != r.Tag {
		return false // fmt.Errorf("Tag string mismatched")
	}

	return true
}

func (l *DocumentField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*DocumentField)
	if !ok {
		return false
	}

	// Field Name - varchar
	if l.Name != r.Name {
		return false // fmt.Errorf("Name string mismatched")
	}

	// Field Type - varchar
	if l.Type != r.Type {
		return false // fmt.Errorf("Type string mismatched")
	}

	// Field Contents - varbin
	if !bytes.Equal(l.Contents, r.Contents) {
		return false // fmt.Errorf("Contents bytes mismatched")
	}

	return true
}

func (l *TargetAddressField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*TargetAddressField)
	if !ok {
		return false
	}

	// Field Address - varbin
	if !bytes.Equal(l.Address, r.Address) {
		return false // fmt.Errorf("Address bytes mismatched")
	}

	// Field Quantity - uint
	if l.Quantity != r.Quantity {
		return false // fmt.Errorf("Quantity integer mismatched")
	}

	return true
}
