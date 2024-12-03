// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ExternalAccountCreate - The external account information
type ExternalAccountCreate struct {
	// The account identifier The account number for external communications
	AccountNumber string `json:"account_number"`
	// The NSCC brokerage / clearing house identifier
	ParticipantNumber string `json:"participant_number"`
}

func (o *ExternalAccountCreate) GetAccountNumber() string {
	if o == nil {
		return ""
	}
	return o.AccountNumber
}

func (o *ExternalAccountCreate) GetParticipantNumber() string {
	if o == nil {
		return ""
	}
	return o.ParticipantNumber
}
