// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// AchWithdrawalCreate - A withdrawal transfer using the ACH mechanism.
type AchWithdrawalCreate struct {
	// A representation of a decimal value, such as 2.5. Clients may convert values into language-native decimal formats, such as Java's [BigDecimal][] or Python's [decimal.Decimal][].
	//
	//  [BigDecimal]:
	//  https://docs.oracle.com/en/java/javase/11/docs/api/java.base/java/math/BigDecimal.html
	//  [decimal.Decimal]: https://docs.python.org/3/library/decimal.html
	Amount *DecimalCreate `json:"amount,omitempty"`
	// The bank relationship to be used for the ACH withdrawal.
	BankRelationship string `json:"bank_relationship"`
	// The external identifier supplied by the API caller. Each request must have a unique pairing of `client_transfer_id` and `account`.
	ClientTransferID string `json:"client_transfer_id"`
	// Whether the entire account balance is being withdrawn. Must not be true if the `amount` is specified.
	FullDisbursement *bool `json:"full_disbursement,omitempty"`
	// The memo that will appear on the customer's bank statement.
	Memo *string `json:"memo,omitempty"`
	// A distribution from a retirement account.
	RetirementDistribution *RetirementDistributionCreate `json:"retirement_distribution,omitempty"`
}

func (o *AchWithdrawalCreate) GetAmount() *DecimalCreate {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *AchWithdrawalCreate) GetBankRelationship() string {
	if o == nil {
		return ""
	}
	return o.BankRelationship
}

func (o *AchWithdrawalCreate) GetClientTransferID() string {
	if o == nil {
		return ""
	}
	return o.ClientTransferID
}

func (o *AchWithdrawalCreate) GetFullDisbursement() *bool {
	if o == nil {
		return nil
	}
	return o.FullDisbursement
}

func (o *AchWithdrawalCreate) GetMemo() *string {
	if o == nil {
		return nil
	}
	return o.Memo
}

func (o *AchWithdrawalCreate) GetRetirementDistribution() *RetirementDistributionCreate {
	if o == nil {
		return nil
	}
	return o.RetirementDistribution
}