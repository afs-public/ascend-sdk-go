// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// InvestmentProfileUpdate - Investor profile.
type InvestmentProfileUpdate struct {
	// The account goals on an investor profile.
	AccountGoals *AccountGoalsUpdate `json:"account_goals,omitempty"`
	// A detailed summary of financial and personal details of an investor, to help understand the investor's financial standing, investment experience and risk tolerance.
	CustomerProfile *CustomerProfileUpdate `json:"customer_profile,omitempty"`
}

func (o *InvestmentProfileUpdate) GetAccountGoals() *AccountGoalsUpdate {
	if o == nil {
		return nil
	}
	return o.AccountGoals
}

func (o *InvestmentProfileUpdate) GetCustomerProfile() *CustomerProfileUpdate {
	if o == nil {
		return nil
	}
	return o.CustomerProfile
}