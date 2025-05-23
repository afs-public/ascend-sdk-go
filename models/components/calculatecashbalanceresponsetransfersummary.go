// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// CalculateCashBalanceResponseTransferSummaryAmount - The amount of the transfer in USD. The value will be positive for deposits and negative for withdrawals.
type CalculateCashBalanceResponseTransferSummaryAmount struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *CalculateCashBalanceResponseTransferSummaryAmount) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// CalculateCashBalanceResponseTransferSummaryMechanism - The mechanism of the transfer.
type CalculateCashBalanceResponseTransferSummaryMechanism string

const (
	CalculateCashBalanceResponseTransferSummaryMechanismMechanismUnspecified CalculateCashBalanceResponseTransferSummaryMechanism = "MECHANISM_UNSPECIFIED"
	CalculateCashBalanceResponseTransferSummaryMechanismAcat                 CalculateCashBalanceResponseTransferSummaryMechanism = "ACAT"
	CalculateCashBalanceResponseTransferSummaryMechanismAch                  CalculateCashBalanceResponseTransferSummaryMechanism = "ACH"
	CalculateCashBalanceResponseTransferSummaryMechanismCashJournal          CalculateCashBalanceResponseTransferSummaryMechanism = "CASH_JOURNAL"
	CalculateCashBalanceResponseTransferSummaryMechanismCheck                CalculateCashBalanceResponseTransferSummaryMechanism = "CHECK"
	CalculateCashBalanceResponseTransferSummaryMechanismCredit               CalculateCashBalanceResponseTransferSummaryMechanism = "CREDIT"
	CalculateCashBalanceResponseTransferSummaryMechanismFee                  CalculateCashBalanceResponseTransferSummaryMechanism = "FEE"
	CalculateCashBalanceResponseTransferSummaryMechanismIct                  CalculateCashBalanceResponseTransferSummaryMechanism = "ICT"
	CalculateCashBalanceResponseTransferSummaryMechanismPaypal               CalculateCashBalanceResponseTransferSummaryMechanism = "PAYPAL"
	CalculateCashBalanceResponseTransferSummaryMechanismRtp                  CalculateCashBalanceResponseTransferSummaryMechanism = "RTP"
	CalculateCashBalanceResponseTransferSummaryMechanismTpj                  CalculateCashBalanceResponseTransferSummaryMechanism = "TPJ"
	CalculateCashBalanceResponseTransferSummaryMechanismWire                 CalculateCashBalanceResponseTransferSummaryMechanism = "WIRE"
	CalculateCashBalanceResponseTransferSummaryMechanismExternalAch          CalculateCashBalanceResponseTransferSummaryMechanism = "EXTERNAL_ACH"
)

func (e CalculateCashBalanceResponseTransferSummaryMechanism) ToPointer() *CalculateCashBalanceResponseTransferSummaryMechanism {
	return &e
}

// CalculateCashBalanceResponseTransferSummary - A summary of a transfer.
type CalculateCashBalanceResponseTransferSummary struct {
	// The amount of the transfer in USD. The value will be positive for deposits and negative for withdrawals.
	Amount *CalculateCashBalanceResponseTransferSummaryAmount `json:"amount,omitempty"`
	// The identifier of the transfer. The value is the last part of the transfer resource name.
	ID *string `json:"id,omitempty"`
	// The mechanism of the transfer.
	Mechanism *CalculateCashBalanceResponseTransferSummaryMechanism `json:"mechanism,omitempty"`
}

func (o *CalculateCashBalanceResponseTransferSummary) GetAmount() *CalculateCashBalanceResponseTransferSummaryAmount {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *CalculateCashBalanceResponseTransferSummary) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CalculateCashBalanceResponseTransferSummary) GetMechanism() *CalculateCashBalanceResponseTransferSummaryMechanism {
	if o == nil {
		return nil
	}
	return o.Mechanism
}
