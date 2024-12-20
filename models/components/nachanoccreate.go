// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// Code - The notice of change reason code.
type Code string

const (
	CodeCodeUnspecified Code = "CODE_UNSPECIFIED"
	CodeC01             Code = "C01"
	CodeC02             Code = "C02"
	CodeC03             Code = "C03"
	CodeC04             Code = "C04"
	CodeC05             Code = "C05"
	CodeC06             Code = "C06"
	CodeC07             Code = "C07"
	CodeC08             Code = "C08"
	CodeC09             Code = "C09"
	CodeC13             Code = "C13"
	CodeC14             Code = "C14"
	CodeC61             Code = "C61"
	CodeC62             Code = "C62"
	CodeC63             Code = "C63"
	CodeC64             Code = "C64"
	CodeC65             Code = "C65"
	CodeC66             Code = "C66"
	CodeC67             Code = "C67"
	CodeC68             Code = "C68"
	CodeC69             Code = "C69"
)

func (e Code) ToPointer() *Code {
	return &e
}

// UpdatedBankAccountType - The updated bank account type. Should only be set when the code is for an incorrect transaction code.
type UpdatedBankAccountType string

const (
	UpdatedBankAccountTypeTypeUnspecified UpdatedBankAccountType = "TYPE_UNSPECIFIED"
	UpdatedBankAccountTypeChecking        UpdatedBankAccountType = "CHECKING"
	UpdatedBankAccountTypeSavings         UpdatedBankAccountType = "SAVINGS"
)

func (e UpdatedBankAccountType) ToPointer() *UpdatedBankAccountType {
	return &e
}

// NachaNocCreate - A notice of change (NOC) on an ACH transfer from the Nacha network.
type NachaNocCreate struct {
	// The notice of change reason code.
	Code Code `json:"code"`
	// The updated bank account number. Should only be set when the code is for an incorrect DFI account number.
	UpdatedBankAccountNumber *string `json:"updated_bank_account_number,omitempty"`
	// The updated bank account type. Should only be set when the code is for an incorrect transaction code.
	UpdatedBankAccountType *UpdatedBankAccountType `json:"updated_bank_account_type,omitempty"`
	// The updated bank routing number. Should only be set when the code is for an incorrect routing number.
	UpdatedBankRoutingNumber *string `json:"updated_bank_routing_number,omitempty"`
}

func (o *NachaNocCreate) GetCode() Code {
	if o == nil {
		return Code("")
	}
	return o.Code
}

func (o *NachaNocCreate) GetUpdatedBankAccountNumber() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedBankAccountNumber
}

func (o *NachaNocCreate) GetUpdatedBankAccountType() *UpdatedBankAccountType {
	if o == nil {
		return nil
	}
	return o.UpdatedBankAccountType
}

func (o *NachaNocCreate) GetUpdatedBankRoutingNumber() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedBankRoutingNumber
}
