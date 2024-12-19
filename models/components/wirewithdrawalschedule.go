// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// Address - The address of the person or entity taking receipt of the wired funds. This will be populated automatically in the case of a valid first-party wire
type Address struct {
	// Required: Describes the city in which the entity is located.
	City *string `json:"city,omitempty"`
	// Required: The country code used for geolocation, identity verification, and/or mail delivery purposes.
	Country *string `json:"country,omitempty"`
	// Required: The postal code used for geolocation, identity verification, and/or mail delivery purposes.
	PostalCode *string `json:"postal_code,omitempty"`
	// Required: The state code used for geolocation, identity verification, and/or mail delivery purposes.
	State *string `json:"state,omitempty"`
	// The street name and number relating to a party's legal or mailing address.
	StreetAddress []string `json:"streetAddress,omitempty"`
}

func (o *Address) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *Address) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *Address) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *Address) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *Address) GetStreetAddress() []string {
	if o == nil {
		return nil
	}
	return o.StreetAddress
}

// WireWithdrawalScheduleBeneficiary - The beneficiary of the wire withdrawal
type WireWithdrawalScheduleBeneficiary struct {
	// The bank account of the person or entity taking receipt of the wired funds. Limited to 25 characters if intermediaryDetails.account is set
	Account *string `json:"account,omitempty"`
	// The name of the person or entity taking receipt of the wired funds. This field defaults to the name of the account owner and should only be populated when performing a third party wire transfer
	AccountTitle *string `json:"account_title,omitempty"`
	// The address of the person or entity taking receipt of the wired funds. This will be populated automatically in the case of a valid first-party wire
	Address *Address `json:"address,omitempty"`
	// Indicates if this beneficiary is a third party beneficiary. A wire transfer is considered third party if the beneficiary is not the exact same person and/or entity that the funds originated from. This includes wire transfers where the originator account is an individual account and the beneficiary account is a joint account
	ThirdParty *bool `json:"third_party,omitempty"`
}

func (o *WireWithdrawalScheduleBeneficiary) GetAccount() *string {
	if o == nil {
		return nil
	}
	return o.Account
}

func (o *WireWithdrawalScheduleBeneficiary) GetAccountTitle() *string {
	if o == nil {
		return nil
	}
	return o.AccountTitle
}

func (o *WireWithdrawalScheduleBeneficiary) GetAddress() *Address {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *WireWithdrawalScheduleBeneficiary) GetThirdParty() *bool {
	if o == nil {
		return nil
	}
	return o.ThirdParty
}

// WireWithdrawalScheduleAddress - The address of the intermediary party
type WireWithdrawalScheduleAddress struct {
	// Required: Describes the city in which the entity is located.
	City *string `json:"city,omitempty"`
	// Required: The country code used for geolocation, identity verification, and/or mail delivery purposes.
	Country *string `json:"country,omitempty"`
	// Required: The postal code used for geolocation, identity verification, and/or mail delivery purposes.
	PostalCode *string `json:"postal_code,omitempty"`
	// Required: The state code used for geolocation, identity verification, and/or mail delivery purposes.
	State *string `json:"state,omitempty"`
	// The street name and number relating to a party's legal or mailing address.
	StreetAddress []string `json:"streetAddress,omitempty"`
}

func (o *WireWithdrawalScheduleAddress) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *WireWithdrawalScheduleAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *WireWithdrawalScheduleAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *WireWithdrawalScheduleAddress) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *WireWithdrawalScheduleAddress) GetStreetAddress() []string {
	if o == nil {
		return nil
	}
	return o.StreetAddress
}

// Intermediary - The intermediary party
type Intermediary struct {
	// The account number of the intermediary party
	Account *string `json:"account,omitempty"`
	// The name of the intermediary party
	AccountTitle *string `json:"account_title,omitempty"`
	// The address of the intermediary party
	Address *WireWithdrawalScheduleAddress `json:"address,omitempty"`
}

func (o *Intermediary) GetAccount() *string {
	if o == nil {
		return nil
	}
	return o.Account
}

func (o *Intermediary) GetAccountTitle() *string {
	if o == nil {
		return nil
	}
	return o.AccountTitle
}

func (o *Intermediary) GetAddress() *WireWithdrawalScheduleAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

// WireWithdrawalScheduleRecipientBankType - The type of bank identifier specified
type WireWithdrawalScheduleRecipientBankType string

const (
	WireWithdrawalScheduleRecipientBankTypeTypeUnspecified WireWithdrawalScheduleRecipientBankType = "TYPE_UNSPECIFIED"
	WireWithdrawalScheduleRecipientBankTypeAba             WireWithdrawalScheduleRecipientBankType = "ABA"
	WireWithdrawalScheduleRecipientBankTypeBic             WireWithdrawalScheduleRecipientBankType = "BIC"
)

func (e WireWithdrawalScheduleRecipientBankType) ToPointer() *WireWithdrawalScheduleRecipientBankType {
	return &e
}

// BankID - An identifier that represents ABA routing number for domestic wire or BIC for foreign wire
type BankID struct {
	// The bank identifier
	ID *string `json:"id,omitempty"`
	// The type of bank identifier specified
	Type *WireWithdrawalScheduleRecipientBankType `json:"type,omitempty"`
}

func (o *BankID) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BankID) GetType() *WireWithdrawalScheduleRecipientBankType {
	if o == nil {
		return nil
	}
	return o.Type
}

// WireWithdrawalScheduleRecipientBankAddress - The address of the recipient bank / financial institution
type WireWithdrawalScheduleRecipientBankAddress struct {
	// Required: Describes the city in which the entity is located.
	City *string `json:"city,omitempty"`
	// Required: The country code used for geolocation, identity verification, and/or mail delivery purposes.
	Country *string `json:"country,omitempty"`
	// Required: The postal code used for geolocation, identity verification, and/or mail delivery purposes.
	PostalCode *string `json:"postal_code,omitempty"`
	// Required: The state code used for geolocation, identity verification, and/or mail delivery purposes.
	State *string `json:"state,omitempty"`
	// The street name and number relating to a party's legal or mailing address.
	StreetAddress []string `json:"streetAddress,omitempty"`
}

func (o *WireWithdrawalScheduleRecipientBankAddress) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *WireWithdrawalScheduleRecipientBankAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *WireWithdrawalScheduleRecipientBankAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *WireWithdrawalScheduleRecipientBankAddress) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *WireWithdrawalScheduleRecipientBankAddress) GetStreetAddress() []string {
	if o == nil {
		return nil
	}
	return o.StreetAddress
}

// InternationalBankDetails - Bank details required in the case of an international wire transfer
type InternationalBankDetails struct {
	// Any additional information to be communicated to the recipient bank, such as intermediary banks to be used.
	AdditionalInfo *string `json:"additional_info,omitempty"`
	// The address of the recipient bank / financial institution
	Address *WireWithdrawalScheduleRecipientBankAddress `json:"address,omitempty"`
	// The name of the recipient bank / financial institution
	BankName *string `json:"bank_name,omitempty"`
}

func (o *InternationalBankDetails) GetAdditionalInfo() *string {
	if o == nil {
		return nil
	}
	return o.AdditionalInfo
}

func (o *InternationalBankDetails) GetAddress() *WireWithdrawalScheduleRecipientBankAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *InternationalBankDetails) GetBankName() *string {
	if o == nil {
		return nil
	}
	return o.BankName
}

// RecipientBank - The recipient bank / financial institution
type RecipientBank struct {
	// An identifier that represents ABA routing number for domestic wire or BIC for foreign wire
	BankID *BankID `json:"bank_id,omitempty"`
	// Bank details required in the case of an international wire transfer
	InternationalBankDetails *InternationalBankDetails `json:"international_bank_details,omitempty"`
}

func (o *RecipientBank) GetBankID() *BankID {
	if o == nil {
		return nil
	}
	return o.BankID
}

func (o *RecipientBank) GetInternationalBankDetails() *InternationalBankDetails {
	if o == nil {
		return nil
	}
	return o.InternationalBankDetails
}

// WireWithdrawalScheduleRetirementDistributionFederalTaxWithholdingAmount - Fixed USD amount to withhold for taxes.
type WireWithdrawalScheduleRetirementDistributionFederalTaxWithholdingAmount struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *WireWithdrawalScheduleRetirementDistributionFederalTaxWithholdingAmount) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// WireWithdrawalSchedulePercentage - Percentage of total disbursement amount to withhold for taxes.
type WireWithdrawalSchedulePercentage struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *WireWithdrawalSchedulePercentage) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// WireWithdrawalScheduleFederalTaxWithholding - The federal tax withholding.
type WireWithdrawalScheduleFederalTaxWithholding struct {
	// Fixed USD amount to withhold for taxes.
	Amount *WireWithdrawalScheduleRetirementDistributionFederalTaxWithholdingAmount `json:"amount,omitempty"`
	// Percentage of total disbursement amount to withhold for taxes.
	Percentage *WireWithdrawalSchedulePercentage `json:"percentage,omitempty"`
}

func (o *WireWithdrawalScheduleFederalTaxWithholding) GetAmount() *WireWithdrawalScheduleRetirementDistributionFederalTaxWithholdingAmount {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *WireWithdrawalScheduleFederalTaxWithholding) GetPercentage() *WireWithdrawalSchedulePercentage {
	if o == nil {
		return nil
	}
	return o.Percentage
}

// WireWithdrawalScheduleRetirementDistributionAmount - Fixed USD amount to withhold for taxes.
type WireWithdrawalScheduleRetirementDistributionAmount struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *WireWithdrawalScheduleRetirementDistributionAmount) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// WireWithdrawalScheduleRetirementDistributionPercentage - Percentage of total disbursement amount to withhold for taxes.
type WireWithdrawalScheduleRetirementDistributionPercentage struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *WireWithdrawalScheduleRetirementDistributionPercentage) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// WireWithdrawalScheduleStateTaxWithholding - The state tax withholding.
type WireWithdrawalScheduleStateTaxWithholding struct {
	// Fixed USD amount to withhold for taxes.
	Amount *WireWithdrawalScheduleRetirementDistributionAmount `json:"amount,omitempty"`
	// Percentage of total disbursement amount to withhold for taxes.
	Percentage *WireWithdrawalScheduleRetirementDistributionPercentage `json:"percentage,omitempty"`
}

func (o *WireWithdrawalScheduleStateTaxWithholding) GetAmount() *WireWithdrawalScheduleRetirementDistributionAmount {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *WireWithdrawalScheduleStateTaxWithholding) GetPercentage() *WireWithdrawalScheduleRetirementDistributionPercentage {
	if o == nil {
		return nil
	}
	return o.Percentage
}

// WireWithdrawalScheduleType - The type of retirement distribution.
type WireWithdrawalScheduleType string

const (
	WireWithdrawalScheduleTypeTypeUnspecified                            WireWithdrawalScheduleType = "TYPE_UNSPECIFIED"
	WireWithdrawalScheduleTypeNormal                                     WireWithdrawalScheduleType = "NORMAL"
	WireWithdrawalScheduleTypeDisability                                 WireWithdrawalScheduleType = "DISABILITY"
	WireWithdrawalScheduleTypeSosepp                                     WireWithdrawalScheduleType = "SOSEPP"
	WireWithdrawalScheduleTypePremature                                  WireWithdrawalScheduleType = "PREMATURE"
	WireWithdrawalScheduleTypeDeath                                      WireWithdrawalScheduleType = "DEATH"
	WireWithdrawalScheduleTypeExcessContributionRemovalBeforeTaxDeadline WireWithdrawalScheduleType = "EXCESS_CONTRIBUTION_REMOVAL_BEFORE_TAX_DEADLINE"
	WireWithdrawalScheduleTypeExcessContributionRemovalAfterTaxDeadline  WireWithdrawalScheduleType = "EXCESS_CONTRIBUTION_REMOVAL_AFTER_TAX_DEADLINE"
	WireWithdrawalScheduleTypeRolloverToQualifiedPlan                    WireWithdrawalScheduleType = "ROLLOVER_TO_QUALIFIED_PLAN"
	WireWithdrawalScheduleTypeRolloverToIra                              WireWithdrawalScheduleType = "ROLLOVER_TO_IRA"
	WireWithdrawalScheduleTypeDistributionTransfer                       WireWithdrawalScheduleType = "DISTRIBUTION_TRANSFER"
	WireWithdrawalScheduleTypeRecharacterizationPriorYear                WireWithdrawalScheduleType = "RECHARACTERIZATION_PRIOR_YEAR"
	WireWithdrawalScheduleTypeRecharacterizationCurrentYear              WireWithdrawalScheduleType = "RECHARACTERIZATION_CURRENT_YEAR"
	WireWithdrawalScheduleTypeDistributionConversion                     WireWithdrawalScheduleType = "DISTRIBUTION_CONVERSION"
	WireWithdrawalScheduleTypeManagementFee                              WireWithdrawalScheduleType = "MANAGEMENT_FEE"
	WireWithdrawalScheduleTypePlanLoan401K                               WireWithdrawalScheduleType = "PLAN_LOAN_401K"
	WireWithdrawalScheduleTypePrematureSimpleIraLessThan2Years           WireWithdrawalScheduleType = "PREMATURE_SIMPLE_IRA_LESS_THAN_2_YEARS"
	WireWithdrawalScheduleTypeNormalRothIraGreaterThan5Years             WireWithdrawalScheduleType = "NORMAL_ROTH_IRA_GREATER_THAN_5_YEARS"
)

func (e WireWithdrawalScheduleType) ToPointer() *WireWithdrawalScheduleType {
	return &e
}

// WireWithdrawalScheduleRetirementDistribution - The distribution info for a retirement account
type WireWithdrawalScheduleRetirementDistribution struct {
	// The federal tax withholding.
	FederalTaxWithholding *WireWithdrawalScheduleFederalTaxWithholding `json:"federal_tax_withholding,omitempty"`
	// The institution receiving retirement funds when performing a transfer to an identical retirement account type at a different financial institution. This is required for check and wire withdrawals because we can't always identify the institution using the transfer instructions. For cash journals this value will default to "Apex Clearing", regardless of what is passed in here
	ReceivingInstitution *string `json:"receiving_institution,omitempty"`
	// The state tax withholding.
	StateTaxWithholding *WireWithdrawalScheduleStateTaxWithholding `json:"state_tax_withholding,omitempty"`
	// Whether or not this distribution has a state withholding waiver.
	StateWithholdingWaiver *bool `json:"state_withholding_waiver,omitempty"`
	// Tax year for which the distribution is applied.
	TaxYear *int `json:"tax_year,omitempty"`
	// The type of retirement distribution.
	Type *WireWithdrawalScheduleType `json:"type,omitempty"`
}

func (o *WireWithdrawalScheduleRetirementDistribution) GetFederalTaxWithholding() *WireWithdrawalScheduleFederalTaxWithholding {
	if o == nil {
		return nil
	}
	return o.FederalTaxWithholding
}

func (o *WireWithdrawalScheduleRetirementDistribution) GetReceivingInstitution() *string {
	if o == nil {
		return nil
	}
	return o.ReceivingInstitution
}

func (o *WireWithdrawalScheduleRetirementDistribution) GetStateTaxWithholding() *WireWithdrawalScheduleStateTaxWithholding {
	if o == nil {
		return nil
	}
	return o.StateTaxWithholding
}

func (o *WireWithdrawalScheduleRetirementDistribution) GetStateWithholdingWaiver() *bool {
	if o == nil {
		return nil
	}
	return o.StateWithholdingWaiver
}

func (o *WireWithdrawalScheduleRetirementDistribution) GetTaxYear() *int {
	if o == nil {
		return nil
	}
	return o.TaxYear
}

func (o *WireWithdrawalScheduleRetirementDistribution) GetType() *WireWithdrawalScheduleType {
	if o == nil {
		return nil
	}
	return o.Type
}

// WireWithdrawalScheduleAmount - A cash amount in the format of decimal value (mutually exclusive with 'full_disbursement')
type WireWithdrawalScheduleAmount struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *WireWithdrawalScheduleAmount) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// WireWithdrawalScheduleStartDate - The schedule start date
type WireWithdrawalScheduleStartDate struct {
	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.
	Day *int `json:"day,omitempty"`
	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
	Month *int `json:"month,omitempty"`
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
	Year *int `json:"year,omitempty"`
}

func (o *WireWithdrawalScheduleStartDate) GetDay() *int {
	if o == nil {
		return nil
	}
	return o.Day
}

func (o *WireWithdrawalScheduleStartDate) GetMonth() *int {
	if o == nil {
		return nil
	}
	return o.Month
}

func (o *WireWithdrawalScheduleStartDate) GetYear() *int {
	if o == nil {
		return nil
	}
	return o.Year
}

// WireWithdrawalScheduleState - The state of the represented schedule
type WireWithdrawalScheduleState string

const (
	WireWithdrawalScheduleStateActive    WireWithdrawalScheduleState = "ACTIVE"
	WireWithdrawalScheduleStateCanceled  WireWithdrawalScheduleState = "CANCELED"
	WireWithdrawalScheduleStateCompleted WireWithdrawalScheduleState = "COMPLETED"
)

func (e WireWithdrawalScheduleState) ToPointer() *WireWithdrawalScheduleState {
	return &e
}

// WireWithdrawalScheduleTimeUnit - The time unit used to calculate the interval between transfers. The time period between transfers in a scheduled series is the unit of time times the multiplier
type WireWithdrawalScheduleTimeUnit string

const (
	WireWithdrawalScheduleTimeUnitDay   WireWithdrawalScheduleTimeUnit = "DAY"
	WireWithdrawalScheduleTimeUnitWeek  WireWithdrawalScheduleTimeUnit = "WEEK"
	WireWithdrawalScheduleTimeUnitMonth WireWithdrawalScheduleTimeUnit = "MONTH"
)

func (e WireWithdrawalScheduleTimeUnit) ToPointer() *WireWithdrawalScheduleTimeUnit {
	return &e
}

// WireWithdrawalScheduleScheduleProperties - Common schedule properties
type WireWithdrawalScheduleScheduleProperties struct {
	// The number of occurrences (empty or 0 indicates unlimited occurrences)
	Occurrences *int `json:"occurrences,omitempty"`
	// The schedule start date
	StartDate *WireWithdrawalScheduleStartDate `json:"start_date,omitempty"`
	// The state of the represented schedule
	State *WireWithdrawalScheduleState `json:"state,omitempty"`
	// The time unit used to calculate the interval between transfers. The time period between transfers in a scheduled series is the unit of time times the multiplier
	TimeUnit *WireWithdrawalScheduleTimeUnit `json:"time_unit,omitempty"`
	// The multiplier used to determine the length of the interval between transfers. The time period between transfers in a scheduled series is the unit of time times the multiplier
	UnitMultiplier *int `json:"unit_multiplier,omitempty"`
}

func (o *WireWithdrawalScheduleScheduleProperties) GetOccurrences() *int {
	if o == nil {
		return nil
	}
	return o.Occurrences
}

func (o *WireWithdrawalScheduleScheduleProperties) GetStartDate() *WireWithdrawalScheduleStartDate {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *WireWithdrawalScheduleScheduleProperties) GetState() *WireWithdrawalScheduleState {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *WireWithdrawalScheduleScheduleProperties) GetTimeUnit() *WireWithdrawalScheduleTimeUnit {
	if o == nil {
		return nil
	}
	return o.TimeUnit
}

func (o *WireWithdrawalScheduleScheduleProperties) GetUnitMultiplier() *int {
	if o == nil {
		return nil
	}
	return o.UnitMultiplier
}

// WireWithdrawalScheduleScheduleDetails - The transfer schedule details
type WireWithdrawalScheduleScheduleDetails struct {
	// A cash amount in the format of decimal value (mutually exclusive with 'full_disbursement')
	Amount *WireWithdrawalScheduleAmount `json:"amount,omitempty"`
	// External identifier supplied by the API caller. Each request must have a unique pairing of client_schedule_id and account
	ClientScheduleID *string `json:"client_schedule_id,omitempty"`
	// Flag to indicate a full disbursement transfer (mutually exclusive with 'amount')
	FullDisbursement *bool `json:"full_disbursement,omitempty"`
	// Common schedule properties
	ScheduleProperties *WireWithdrawalScheduleScheduleProperties `json:"schedule_properties,omitempty"`
}

func (o *WireWithdrawalScheduleScheduleDetails) GetAmount() *WireWithdrawalScheduleAmount {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *WireWithdrawalScheduleScheduleDetails) GetClientScheduleID() *string {
	if o == nil {
		return nil
	}
	return o.ClientScheduleID
}

func (o *WireWithdrawalScheduleScheduleDetails) GetFullDisbursement() *bool {
	if o == nil {
		return nil
	}
	return o.FullDisbursement
}

func (o *WireWithdrawalScheduleScheduleDetails) GetScheduleProperties() *WireWithdrawalScheduleScheduleProperties {
	if o == nil {
		return nil
	}
	return o.ScheduleProperties
}

// WireWithdrawalSchedule - A withdrawal transfer schedule using the Wire mechanism
type WireWithdrawalSchedule struct {
	// The beneficiary of the wire withdrawal
	Beneficiary *WireWithdrawalScheduleBeneficiary `json:"beneficiary,omitempty"`
	// The intermediary party
	Intermediary *Intermediary `json:"intermediary,omitempty"`
	// The name of the Wire Withdrawal transfer schedule
	Name *string `json:"name,omitempty"`
	// The recipient bank / financial institution
	RecipientBank *RecipientBank `json:"recipient_bank,omitempty"`
	// The distribution info for a retirement account
	RetirementDistribution *WireWithdrawalScheduleRetirementDistribution `json:"retirement_distribution,omitempty"`
	// The transfer schedule details
	ScheduleDetails *WireWithdrawalScheduleScheduleDetails `json:"schedule_details,omitempty"`
}

func (o *WireWithdrawalSchedule) GetBeneficiary() *WireWithdrawalScheduleBeneficiary {
	if o == nil {
		return nil
	}
	return o.Beneficiary
}

func (o *WireWithdrawalSchedule) GetIntermediary() *Intermediary {
	if o == nil {
		return nil
	}
	return o.Intermediary
}

func (o *WireWithdrawalSchedule) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *WireWithdrawalSchedule) GetRecipientBank() *RecipientBank {
	if o == nil {
		return nil
	}
	return o.RecipientBank
}

func (o *WireWithdrawalSchedule) GetRetirementDistribution() *WireWithdrawalScheduleRetirementDistribution {
	if o == nil {
		return nil
	}
	return o.RetirementDistribution
}

func (o *WireWithdrawalSchedule) GetScheduleDetails() *WireWithdrawalScheduleScheduleDetails {
	if o == nil {
		return nil
	}
	return o.ScheduleDetails
}
