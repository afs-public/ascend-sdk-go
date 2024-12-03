// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// WireWithdrawalRecipientBankCreate - A recipient bank / financial institution
type WireWithdrawalRecipientBankCreate struct {
	// A bank identifier
	BankID RecipientBankBankIDCreate `json:"bank_id"`
	// Bank details
	InternationalBankDetails *RecipientBankBankDetailsCreate `json:"international_bank_details,omitempty"`
}

func (o *WireWithdrawalRecipientBankCreate) GetBankID() RecipientBankBankIDCreate {
	if o == nil {
		return RecipientBankBankIDCreate{}
	}
	return o.BankID
}

func (o *WireWithdrawalRecipientBankCreate) GetInternationalBankDetails() *RecipientBankBankDetailsCreate {
	if o == nil {
		return nil
	}
	return o.InternationalBankDetails
}
