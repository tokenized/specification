package messages

import (
	"fmt"
)

const (
	max1ByteInteger = 255
	max2ByteInteger = 65535
	max4ByteInteger = 4294967295
)

func (a *PublicMessage) Validate() error {
	if a == nil {
		return nil
	}

	// Field Timestamp - uint

	// Field Subject - varchar
	if len(a.Subject) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Subject), max2ByteInteger)
	}

	// Field Regarding - Outpoint

	// Field PublicMessage - Document

	// Field Attachments - Document
	if len(a.Attachments) > max4ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.Attachments), max4ByteInteger)
	}
	for i, v := range a.Attachments {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("Attachments[%d] invalid : %s", i, err)
		}
	}

	return nil
}

func (a *PrivateMessage) Validate() error {
	if a == nil {
		return nil
	}

	// Field Timestamp - uint

	// Field Subject - varchar
	if len(a.Subject) > max2ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Subject), max2ByteInteger)
	}

	// Field Regarding - Outpoint

	// Field PrivateMessage - Document

	// Field Attachments - Document
	if len(a.Attachments) > max4ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.Attachments), max4ByteInteger)
	}
	for i, v := range a.Attachments {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("Attachments[%d] invalid : %s", i, err)
		}
	}

	return nil
}

func (a *RevertedTx) Validate() error {
	if a == nil {
		return nil
	}

	// Field Timestamp - uint

	// Field Transaction - varbin
	if len(a.Transaction) > max4ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Transaction), max4ByteInteger)
	}

	return nil
}

func (a *Offer) Validate() error {
	if a == nil {
		return nil
	}

	// Field Timestamp - uint

	// Field Payload - varbin
	if len(a.Payload) > max4ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Payload), max4ByteInteger)
	}

	return nil
}

func (a *SignatureRequest) Validate() error {
	if a == nil {
		return nil
	}

	// Field Timestamp - uint

	// Field Payload - varbin
	if len(a.Payload) > max4ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Payload), max4ByteInteger)
	}

	return nil
}

func (a *SettlementRequest) Validate() error {
	if a == nil {
		return nil
	}

	// Field Timestamp - uint

	// Field TransferTxId - bin
	if len(a.TransferTxId) != 0 && len(a.TransferTxId) != 32 {
		return fmt.Errorf("Fixed width field TransferTxId wrong size : %d should be %d",
			len(a.TransferTxId), 32)
	}

	// Field ContractFees - TargetAddress
	if len(a.ContractFees) > max2ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.ContractFees), max2ByteInteger)
	}
	for i, v := range a.ContractFees {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("ContractFees[%d] invalid : %s", i, err)
		}
	}

	// Field Settlement - varbin
	if len(a.Settlement) > max4ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Settlement), max4ByteInteger)
	}

	return nil
}

func (a *OutputMetadata) Validate() error {
	if a == nil {
		return nil
	}

	// Field OutputDescription - varchar
	if len(a.OutputDescription) > max4ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.OutputDescription), max4ByteInteger)
	}

	// Field Tags - uint
	if len(a.Tags) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.Tags), max1ByteInteger)
	}

	// Field CustomTags - OutputTag
	if len(a.CustomTags) > max1ByteInteger {
		return fmt.Errorf("List over max length : %d > %d", len(a.CustomTags), max1ByteInteger)
	}
	for i, v := range a.CustomTags {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("CustomTags[%d] invalid : %s", i, err)
		}
	}

	return nil
}

func (a *OutpointField) Validate() error {
	if a == nil {
		return nil
	}

	// Field TxId - bin
	if len(a.TxId) != 0 && len(a.TxId) != 32 {
		return fmt.Errorf("Fixed width field TxId wrong size : %d should be %d",
			len(a.TxId), 32)
	}

	// Field OutputIndex - uint

	return nil
}

func (a *OutputTagField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Tag - varchar
	if len(a.Tag) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Tag), max1ByteInteger)
	}

	return nil
}

func (a *DocumentField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Name - varchar
	if len(a.Name) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Name), max1ByteInteger)
	}

	// Field Type - varchar
	if len(a.Type) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Type), max1ByteInteger)
	}

	// Field Contents - varbin
	if len(a.Contents) > max4ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Contents), max4ByteInteger)
	}

	return nil
}

func (a *TargetAddressField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Address - varbin
	if len(a.Address) > max1ByteInteger {
		return fmt.Errorf("variable size over max value : %d > %d", len(a.Address), max1ByteInteger)
	}

	// Field Quantity - uint

	return nil
}
