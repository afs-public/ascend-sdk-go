// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// TransferAccountCreate - The delivering/receiving party information
type TransferAccountCreate struct {
	// The internal apex account id
	ApexAccountID string `json:"apex_account_id"`
	// The external account information
	ExternalAccount ExternalAccountCreate `json:"external_account"`
}

func (o *TransferAccountCreate) GetApexAccountID() string {
	if o == nil {
		return ""
	}
	return o.ApexAccountID
}

func (o *TransferAccountCreate) GetExternalAccount() ExternalAccountCreate {
	if o == nil {
		return ExternalAccountCreate{}
	}
	return o.ExternalAccount
}
