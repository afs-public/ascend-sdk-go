// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// WireWithdrawalCreate - A withdrawal transfer using the wire mechanism
type WireWithdrawalCreate struct {
	// A representation of a decimal value, such as 2.5. Clients may convert values into language-native decimal formats, such as Java's [BigDecimal][] or Python's [decimal.Decimal][].
	//
	//  [BigDecimal]:
	//  https://docs.oracle.com/en/java/javase/11/docs/api/java.base/java/math/BigDecimal.html
	//  [decimal.Decimal]: https://docs.python.org/3/library/decimal.html
	Amount *DecimalCreate `json:"amount,omitempty"`
	// The person or entity taking receipt of the wired funds
	Beneficiary WireWithdrawalBeneficiaryCreate `json:"beneficiary"`
	// External identifier supplied by the API caller. Each request must have a unique pairing of client_transfer_id and account
	ClientTransferID string `json:"client_transfer_id"`
	// Whether the entire account balance is being withdrawn. If this field is set to True, a transfer amount must not be specified
	FullDisbursement *bool `json:"full_disbursement,omitempty"`
	// An intermediary party
	Intermediary *WireWithdrawalIntermediaryCreate `json:"intermediary,omitempty"`
	// A distribution from a retirement account.
	IraDistribution *RetirementDistributionCreate `json:"ira_distribution,omitempty"`
	// A recipient bank / financial institution
	RecipientBank WireWithdrawalRecipientBankCreate `json:"recipient_bank"`
}

func (o *WireWithdrawalCreate) GetAmount() *DecimalCreate {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *WireWithdrawalCreate) GetBeneficiary() WireWithdrawalBeneficiaryCreate {
	if o == nil {
		return WireWithdrawalBeneficiaryCreate{}
	}
	return o.Beneficiary
}

func (o *WireWithdrawalCreate) GetClientTransferID() string {
	if o == nil {
		return ""
	}
	return o.ClientTransferID
}

func (o *WireWithdrawalCreate) GetFullDisbursement() *bool {
	if o == nil {
		return nil
	}
	return o.FullDisbursement
}

func (o *WireWithdrawalCreate) GetIntermediary() *WireWithdrawalIntermediaryCreate {
	if o == nil {
		return nil
	}
	return o.Intermediary
}

func (o *WireWithdrawalCreate) GetIraDistribution() *RetirementDistributionCreate {
	if o == nil {
		return nil
	}
	return o.IraDistribution
}

func (o *WireWithdrawalCreate) GetRecipientBank() WireWithdrawalRecipientBankCreate {
	if o == nil {
		return WireWithdrawalRecipientBankCreate{}
	}
	return o.RecipientBank
}
