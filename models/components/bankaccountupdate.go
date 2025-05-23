// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// BankAccountUpdateType - The bank account type.
type BankAccountUpdateType string

const (
	BankAccountUpdateTypeTypeUnspecified BankAccountUpdateType = "TYPE_UNSPECIFIED"
	BankAccountUpdateTypeChecking        BankAccountUpdateType = "CHECKING"
	BankAccountUpdateTypeSavings         BankAccountUpdateType = "SAVINGS"
)

func (e BankAccountUpdateType) ToPointer() *BankAccountUpdateType {
	return &e
}

// BankAccountUpdate - A representation of a bank account.
type BankAccountUpdate struct {
	// The bank account number. This value will be masked in responses.
	AccountNumber *string `json:"account_number,omitempty"`
	// The name of the bank account owner.
	Owner *string `json:"owner,omitempty"`
	// The bank routing number (either ABA or BIC).
	RoutingNumber *string `json:"routing_number,omitempty"`
	// The bank account type.
	Type *BankAccountUpdateType `json:"type,omitempty"`
}

func (o *BankAccountUpdate) GetAccountNumber() *string {
	if o == nil {
		return nil
	}
	return o.AccountNumber
}

func (o *BankAccountUpdate) GetOwner() *string {
	if o == nil {
		return nil
	}
	return o.Owner
}

func (o *BankAccountUpdate) GetRoutingNumber() *string {
	if o == nil {
		return nil
	}
	return o.RoutingNumber
}

func (o *BankAccountUpdate) GetType() *BankAccountUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}
