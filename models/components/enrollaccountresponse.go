// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// EnrollAccountResponse - The response for enrolling an account.
type EnrollAccountResponse struct {
	// The collection of legal agreements that require affirmation to enroll the account in a program.
	Agreements []Agreement `json:"agreements,omitempty"`
}

func (o *EnrollAccountResponse) GetAgreements() []Agreement {
	if o == nil {
		return nil
	}
	return o.Agreements
}