// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// AccountRequestUpdateCatAccountHolderType - The FINRA CAT classification for the Account Holder; Is set automatically based on attributes of the owners and account type
type AccountRequestUpdateCatAccountHolderType string

const (
	AccountRequestUpdateCatAccountHolderTypeCatAccountHolderTypeUnspecified AccountRequestUpdateCatAccountHolderType = "CAT_ACCOUNT_HOLDER_TYPE_UNSPECIFIED"
	AccountRequestUpdateCatAccountHolderTypeAInstitutionalCustomer          AccountRequestUpdateCatAccountHolderType = "A_INSTITUTIONAL_CUSTOMER"
	AccountRequestUpdateCatAccountHolderTypeEEmployeeAccount                AccountRequestUpdateCatAccountHolderType = "E_EMPLOYEE_ACCOUNT"
	AccountRequestUpdateCatAccountHolderTypeFForeign                        AccountRequestUpdateCatAccountHolderType = "F_FOREIGN"
	AccountRequestUpdateCatAccountHolderTypeIIndividual                     AccountRequestUpdateCatAccountHolderType = "I_INDIVIDUAL"
	AccountRequestUpdateCatAccountHolderTypeOMarketMaking                   AccountRequestUpdateCatAccountHolderType = "O_MARKET_MAKING"
	AccountRequestUpdateCatAccountHolderTypeVAgencyAveragePriceAccount      AccountRequestUpdateCatAccountHolderType = "V_AGENCY_AVERAGE_PRICE_ACCOUNT"
	AccountRequestUpdateCatAccountHolderTypePOtherProprietary               AccountRequestUpdateCatAccountHolderType = "P_OTHER_PROPRIETARY"
	AccountRequestUpdateCatAccountHolderTypeXErrorAccount                   AccountRequestUpdateCatAccountHolderType = "X_ERROR_ACCOUNT"
)

func (e AccountRequestUpdateCatAccountHolderType) ToPointer() *AccountRequestUpdateCatAccountHolderType {
	return &e
}
func (e *AccountRequestUpdateCatAccountHolderType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CAT_ACCOUNT_HOLDER_TYPE_UNSPECIFIED":
		fallthrough
	case "A_INSTITUTIONAL_CUSTOMER":
		fallthrough
	case "E_EMPLOYEE_ACCOUNT":
		fallthrough
	case "F_FOREIGN":
		fallthrough
	case "I_INDIVIDUAL":
		fallthrough
	case "O_MARKET_MAKING":
		fallthrough
	case "V_AGENCY_AVERAGE_PRICE_ACCOUNT":
		fallthrough
	case "P_OTHER_PROPRIETARY":
		fallthrough
	case "X_ERROR_ACCOUNT":
		*e = AccountRequestUpdateCatAccountHolderType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountRequestUpdateCatAccountHolderType: %v", v)
	}
}

// AccountRequestUpdate - A single record representing an owner or manager of an Account.
type AccountRequestUpdate struct {
	// Indicates if the issuer of a security held by the account is permitted to communicate directly with the shareholder versus through the brokerage firm; This can include sending proxy statements, annual reports, and other important information directly to the shareholder's address on file with the brokerage firm
	AcceptsIssuerDirectCommunication *bool `json:"accepts_issuer_direct_communication,omitempty"`
	// A boolean to indicate if an account is advised
	Advised *bool `json:"advised,omitempty"`
	// The FINRA CAT classification for the Account Holder; Is set automatically based on attributes of the owners and account type
	CatAccountHolderType *AccountRequestUpdateCatAccountHolderType `json:"cat_account_holder_type,omitempty"`
	// A list of identifiers associated with the account
	Identifiers []IdentifierUpdate `json:"identifiers,omitempty"`
	// A list of natural persons indicated to receive selected account documents such as account statements
	InterestedParties []InterestedPartyUpdate `json:"interested_parties,omitempty"`
	// Investor profile.
	InvestmentProfile *InvestmentProfileUpdate `json:"investment_profile,omitempty"`
	// A boolean to indicate if an account is managed
	Managed *bool `json:"managed,omitempty"`
	// Parties associated with the account (e.g. custodian).
	Parties []PartyRequestUpdate `json:"parties,omitempty"`
	// The primary registered representative for the account
	PrimaryRegisteredRepID *string `json:"primary_registered_rep_id,omitempty"`
	// The account tax profile.
	TaxProfile *AccountTaxProfileUpdate `json:"tax_profile,omitempty"`
	// A list of persons designated to verify the well being of the account holder.
	TrustedContacts []TrustedContactUpdate `json:"trusted_contacts,omitempty"`
	// A boolean to indicate if an account is a wrap brokerage account
	WrapFeeBilled *bool `json:"wrap_fee_billed,omitempty"`
}

func (o *AccountRequestUpdate) GetAcceptsIssuerDirectCommunication() *bool {
	if o == nil {
		return nil
	}
	return o.AcceptsIssuerDirectCommunication
}

func (o *AccountRequestUpdate) GetAdvised() *bool {
	if o == nil {
		return nil
	}
	return o.Advised
}

func (o *AccountRequestUpdate) GetCatAccountHolderType() *AccountRequestUpdateCatAccountHolderType {
	if o == nil {
		return nil
	}
	return o.CatAccountHolderType
}

func (o *AccountRequestUpdate) GetIdentifiers() []IdentifierUpdate {
	if o == nil {
		return nil
	}
	return o.Identifiers
}

func (o *AccountRequestUpdate) GetInterestedParties() []InterestedPartyUpdate {
	if o == nil {
		return nil
	}
	return o.InterestedParties
}

func (o *AccountRequestUpdate) GetInvestmentProfile() *InvestmentProfileUpdate {
	if o == nil {
		return nil
	}
	return o.InvestmentProfile
}

func (o *AccountRequestUpdate) GetManaged() *bool {
	if o == nil {
		return nil
	}
	return o.Managed
}

func (o *AccountRequestUpdate) GetParties() []PartyRequestUpdate {
	if o == nil {
		return nil
	}
	return o.Parties
}

func (o *AccountRequestUpdate) GetPrimaryRegisteredRepID() *string {
	if o == nil {
		return nil
	}
	return o.PrimaryRegisteredRepID
}

func (o *AccountRequestUpdate) GetTaxProfile() *AccountTaxProfileUpdate {
	if o == nil {
		return nil
	}
	return o.TaxProfile
}

func (o *AccountRequestUpdate) GetTrustedContacts() []TrustedContactUpdate {
	if o == nil {
		return nil
	}
	return o.TrustedContacts
}

func (o *AccountRequestUpdate) GetWrapFeeBilled() *bool {
	if o == nil {
		return nil
	}
	return o.WrapFeeBilled
}
