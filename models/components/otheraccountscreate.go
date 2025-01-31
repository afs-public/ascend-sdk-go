// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// OtherAccountsCreate - A customer-disclosed list of other Apex-held accounts owned by the Entity applicant at the time of this account's application; expressed as zero, one, or many account numbers
type OtherAccountsCreate struct {
	// Other account names held at Apex
	AccountNames []string `json:"account_names,omitempty"`
	// Other account numbers held at Apex
	AccountNumbers []string `json:"account_numbers,omitempty"`
	// The owner has other accounts at Apex
	OwnerHasOtherAccountsAtApex bool `json:"owner_has_other_accounts_at_apex"`
}

func (o *OtherAccountsCreate) GetAccountNames() []string {
	if o == nil {
		return nil
	}
	return o.AccountNames
}

func (o *OtherAccountsCreate) GetAccountNumbers() []string {
	if o == nil {
		return nil
	}
	return o.AccountNumbers
}

func (o *OtherAccountsCreate) GetOwnerHasOtherAccountsAtApex() bool {
	if o == nil {
		return false
	}
	return o.OwnerHasOtherAccountsAtApex
}
