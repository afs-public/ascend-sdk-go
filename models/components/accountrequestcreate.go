// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// CatAccountHolderType - The FINRA CAT classification for the Account Holder; Is set automatically based on attributes of the owners and account type
type CatAccountHolderType string

const (
	CatAccountHolderTypeCatAccountHolderTypeUnspecified CatAccountHolderType = "CAT_ACCOUNT_HOLDER_TYPE_UNSPECIFIED"
	CatAccountHolderTypeAInstitutionalCustomer          CatAccountHolderType = "A_INSTITUTIONAL_CUSTOMER"
	CatAccountHolderTypeEEmployeeAccount                CatAccountHolderType = "E_EMPLOYEE_ACCOUNT"
	CatAccountHolderTypeFForeign                        CatAccountHolderType = "F_FOREIGN"
	CatAccountHolderTypeIIndividual                     CatAccountHolderType = "I_INDIVIDUAL"
	CatAccountHolderTypeOMarketMaking                   CatAccountHolderType = "O_MARKET_MAKING"
	CatAccountHolderTypeVAgencyAveragePriceAccount      CatAccountHolderType = "V_AGENCY_AVERAGE_PRICE_ACCOUNT"
	CatAccountHolderTypePOtherProprietary               CatAccountHolderType = "P_OTHER_PROPRIETARY"
	CatAccountHolderTypeXErrorAccount                   CatAccountHolderType = "X_ERROR_ACCOUNT"
)

func (e CatAccountHolderType) ToPointer() *CatAccountHolderType {
	return &e
}

// AccountRequestCreate - A single record representing an owner or manager of an Account.
type AccountRequestCreate struct {
	// Indicates if the issuer of a security held by the account is permitted to communicate directly with the shareholder versus through the brokerage firm; This can include sending proxy statements, annual reports, and other important information directly to the shareholder's address on file with the brokerage firm
	AcceptsIssuerDirectCommunication *bool `json:"accepts_issuer_direct_communication,omitempty"`
	// An Account Group is a way of segmenting accounts within a Correspondent; It is up to the client to define what these groups are and AFS Operations is responsible for configuring them; If the client requests additional groups/codes, they can be added; Examples of Account Groups could hypothetically include HNW (High Net Worth), GOLD (Gold Status Customer), and NWC (Northwest Branch Customer)
	AccountGroupID string `json:"account_group_id"`
	// A boolean to indicate if an account is advised
	Advised *bool `json:"advised,omitempty"`
	// The FINRA CAT classification for the Account Holder; Is set automatically based on attributes of the owners and account type
	CatAccountHolderType *CatAccountHolderType `json:"cat_account_holder_type,omitempty"`
	// A unique identifier referencing a Correspondent; A Client may have several operating Correspondents within its purview.
	CorrespondentID string `json:"correspondent_id"`
	// A list of identifiers associated with the account
	Identifiers []IdentifierCreate `json:"identifiers,omitempty"`
	// A list of natural persons indicated to receive selected account documents such as account statements
	InterestedParties []InterestedPartyCreate `json:"interested_parties,omitempty"`
	// Investor profile.
	InvestmentProfile *InvestmentProfileCreate `json:"investment_profile,omitempty"`
	// A boolean to indicate if an account is managed
	Managed *bool `json:"managed,omitempty"`
	// Parties associated with the account (e.g. custodian).
	Parties []PartyRequestCreate `json:"parties"`
	// The primary registered representative for the account
	PrimaryRegisteredRepID *string `json:"primary_registered_rep_id,omitempty"`
	// The account tax profile.
	TaxProfile *AccountTaxProfileCreate `json:"tax_profile,omitempty"`
	// A list of persons designated to verify the well being of the account holder.
	TrustedContacts []TrustedContactCreate `json:"trusted_contacts,omitempty"`
	// A boolean to indicate if an account is a wrap brokerage account
	WrapFeeBilled *bool `json:"wrap_fee_billed,omitempty"`
}

func (o *AccountRequestCreate) GetAcceptsIssuerDirectCommunication() *bool {
	if o == nil {
		return nil
	}
	return o.AcceptsIssuerDirectCommunication
}

func (o *AccountRequestCreate) GetAccountGroupID() string {
	if o == nil {
		return ""
	}
	return o.AccountGroupID
}

func (o *AccountRequestCreate) GetAdvised() *bool {
	if o == nil {
		return nil
	}
	return o.Advised
}

func (o *AccountRequestCreate) GetCatAccountHolderType() *CatAccountHolderType {
	if o == nil {
		return nil
	}
	return o.CatAccountHolderType
}

func (o *AccountRequestCreate) GetCorrespondentID() string {
	if o == nil {
		return ""
	}
	return o.CorrespondentID
}

func (o *AccountRequestCreate) GetIdentifiers() []IdentifierCreate {
	if o == nil {
		return nil
	}
	return o.Identifiers
}

func (o *AccountRequestCreate) GetInterestedParties() []InterestedPartyCreate {
	if o == nil {
		return nil
	}
	return o.InterestedParties
}

func (o *AccountRequestCreate) GetInvestmentProfile() *InvestmentProfileCreate {
	if o == nil {
		return nil
	}
	return o.InvestmentProfile
}

func (o *AccountRequestCreate) GetManaged() *bool {
	if o == nil {
		return nil
	}
	return o.Managed
}

func (o *AccountRequestCreate) GetParties() []PartyRequestCreate {
	if o == nil {
		return []PartyRequestCreate{}
	}
	return o.Parties
}

func (o *AccountRequestCreate) GetPrimaryRegisteredRepID() *string {
	if o == nil {
		return nil
	}
	return o.PrimaryRegisteredRepID
}

func (o *AccountRequestCreate) GetTaxProfile() *AccountTaxProfileCreate {
	if o == nil {
		return nil
	}
	return o.TaxProfile
}

func (o *AccountRequestCreate) GetTrustedContacts() []TrustedContactCreate {
	if o == nil {
		return nil
	}
	return o.TrustedContacts
}

func (o *AccountRequestCreate) GetWrapFeeBilled() *bool {
	if o == nil {
		return nil
	}
	return o.WrapFeeBilled
}