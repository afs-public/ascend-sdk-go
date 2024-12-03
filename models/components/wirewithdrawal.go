// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"ascend-sdk/internal/utils"
	"encoding/json"
	"fmt"
	"time"
)

// WireWithdrawalAmount - A cash amount in the format of decimal value
type WireWithdrawalAmount struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *WireWithdrawalAmount) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// WireWithdrawalAddress - The address of the person or entity taking receipt of the wired funds. This will be populated automatically in the case of a valid first-party wire
type WireWithdrawalAddress struct {
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

func (o *WireWithdrawalAddress) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *WireWithdrawalAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *WireWithdrawalAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *WireWithdrawalAddress) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *WireWithdrawalAddress) GetStreetAddress() []string {
	if o == nil {
		return nil
	}
	return o.StreetAddress
}

// WireWithdrawalBeneficiary - The beneficiary of the wire withdrawal
type WireWithdrawalBeneficiary struct {
	// The bank account of the person or entity taking receipt of the wired funds. Limited to 25 characters if intermediaryDetails.account is set
	Account *string `json:"account,omitempty"`
	// The name of the person or entity taking receipt of the wired funds. This field defaults to the name of the account owner and should only be populated when performing a third party wire transfer
	AccountTitle *string `json:"account_title,omitempty"`
	// The address of the person or entity taking receipt of the wired funds. This will be populated automatically in the case of a valid first-party wire
	Address *WireWithdrawalAddress `json:"address,omitempty"`
	// Indicates if this beneficiary is a third party beneficiary. A wire transfer is considered third party if the beneficiary is not the exact same person and/or entity that the funds originated from. This includes wire transfers where the originator account is an individual account and the beneficiary account is a joint account
	ThirdParty *bool `json:"third_party,omitempty"`
}

func (o *WireWithdrawalBeneficiary) GetAccount() *string {
	if o == nil {
		return nil
	}
	return o.Account
}

func (o *WireWithdrawalBeneficiary) GetAccountTitle() *string {
	if o == nil {
		return nil
	}
	return o.AccountTitle
}

func (o *WireWithdrawalBeneficiary) GetAddress() *WireWithdrawalAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *WireWithdrawalBeneficiary) GetThirdParty() *bool {
	if o == nil {
		return nil
	}
	return o.ThirdParty
}

// WireWithdrawalIntermediaryAddress - The address of the intermediary party
type WireWithdrawalIntermediaryAddress struct {
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

func (o *WireWithdrawalIntermediaryAddress) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *WireWithdrawalIntermediaryAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *WireWithdrawalIntermediaryAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *WireWithdrawalIntermediaryAddress) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *WireWithdrawalIntermediaryAddress) GetStreetAddress() []string {
	if o == nil {
		return nil
	}
	return o.StreetAddress
}

// WireWithdrawalIntermediary - The intermediary party
type WireWithdrawalIntermediary struct {
	// The account number of the intermediary party
	Account *string `json:"account,omitempty"`
	// The name of the intermediary party
	AccountTitle *string `json:"account_title,omitempty"`
	// The address of the intermediary party
	Address *WireWithdrawalIntermediaryAddress `json:"address,omitempty"`
}

func (o *WireWithdrawalIntermediary) GetAccount() *string {
	if o == nil {
		return nil
	}
	return o.Account
}

func (o *WireWithdrawalIntermediary) GetAccountTitle() *string {
	if o == nil {
		return nil
	}
	return o.AccountTitle
}

func (o *WireWithdrawalIntermediary) GetAddress() *WireWithdrawalIntermediaryAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

// WireWithdrawalIraDistributionAmount - Fixed USD amount to withhold for taxes.
type WireWithdrawalIraDistributionAmount struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *WireWithdrawalIraDistributionAmount) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// WireWithdrawalPercentage - Percentage of total disbursement amount to withhold for taxes.
type WireWithdrawalPercentage struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *WireWithdrawalPercentage) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// WireWithdrawalFederalTaxWithholding - The federal tax withholding.
type WireWithdrawalFederalTaxWithholding struct {
	// Fixed USD amount to withhold for taxes.
	Amount *WireWithdrawalIraDistributionAmount `json:"amount,omitempty"`
	// Percentage of total disbursement amount to withhold for taxes.
	Percentage *WireWithdrawalPercentage `json:"percentage,omitempty"`
}

func (o *WireWithdrawalFederalTaxWithholding) GetAmount() *WireWithdrawalIraDistributionAmount {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *WireWithdrawalFederalTaxWithholding) GetPercentage() *WireWithdrawalPercentage {
	if o == nil {
		return nil
	}
	return o.Percentage
}

// WireWithdrawalIraDistributionStateTaxWithholdingAmount - Fixed USD amount to withhold for taxes.
type WireWithdrawalIraDistributionStateTaxWithholdingAmount struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *WireWithdrawalIraDistributionStateTaxWithholdingAmount) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// WireWithdrawalIraDistributionPercentage - Percentage of total disbursement amount to withhold for taxes.
type WireWithdrawalIraDistributionPercentage struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *WireWithdrawalIraDistributionPercentage) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// WireWithdrawalStateTaxWithholding - The state tax withholding.
type WireWithdrawalStateTaxWithholding struct {
	// Fixed USD amount to withhold for taxes.
	Amount *WireWithdrawalIraDistributionStateTaxWithholdingAmount `json:"amount,omitempty"`
	// Percentage of total disbursement amount to withhold for taxes.
	Percentage *WireWithdrawalIraDistributionPercentage `json:"percentage,omitempty"`
}

func (o *WireWithdrawalStateTaxWithholding) GetAmount() *WireWithdrawalIraDistributionStateTaxWithholdingAmount {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *WireWithdrawalStateTaxWithholding) GetPercentage() *WireWithdrawalIraDistributionPercentage {
	if o == nil {
		return nil
	}
	return o.Percentage
}

// WireWithdrawalType - The type of retirement distribution.
type WireWithdrawalType string

const (
	WireWithdrawalTypeTypeUnspecified                            WireWithdrawalType = "TYPE_UNSPECIFIED"
	WireWithdrawalTypeNormal                                     WireWithdrawalType = "NORMAL"
	WireWithdrawalTypeDisability                                 WireWithdrawalType = "DISABILITY"
	WireWithdrawalTypeSosepp                                     WireWithdrawalType = "SOSEPP"
	WireWithdrawalTypePremature                                  WireWithdrawalType = "PREMATURE"
	WireWithdrawalTypeDeath                                      WireWithdrawalType = "DEATH"
	WireWithdrawalTypeExcessContributionRemovalBeforeTaxDeadline WireWithdrawalType = "EXCESS_CONTRIBUTION_REMOVAL_BEFORE_TAX_DEADLINE"
	WireWithdrawalTypeExcessContributionRemovalAfterTaxDeadline  WireWithdrawalType = "EXCESS_CONTRIBUTION_REMOVAL_AFTER_TAX_DEADLINE"
	WireWithdrawalTypeRolloverToQualifiedPlan                    WireWithdrawalType = "ROLLOVER_TO_QUALIFIED_PLAN"
	WireWithdrawalTypeRolloverToIra                              WireWithdrawalType = "ROLLOVER_TO_IRA"
	WireWithdrawalTypeDistributionTransfer                       WireWithdrawalType = "DISTRIBUTION_TRANSFER"
	WireWithdrawalTypeRecharacterizationPriorYear                WireWithdrawalType = "RECHARACTERIZATION_PRIOR_YEAR"
	WireWithdrawalTypeRecharacterizationCurrentYear              WireWithdrawalType = "RECHARACTERIZATION_CURRENT_YEAR"
	WireWithdrawalTypeDistributionConversion                     WireWithdrawalType = "DISTRIBUTION_CONVERSION"
	WireWithdrawalTypeManagementFee                              WireWithdrawalType = "MANAGEMENT_FEE"
	WireWithdrawalTypePlanLoan401K                               WireWithdrawalType = "PLAN_LOAN_401K"
	WireWithdrawalTypePrematureSimpleIraLessThan2Years           WireWithdrawalType = "PREMATURE_SIMPLE_IRA_LESS_THAN_2_YEARS"
	WireWithdrawalTypeNormalRothIraGreaterThan5Years             WireWithdrawalType = "NORMAL_ROTH_IRA_GREATER_THAN_5_YEARS"
)

func (e WireWithdrawalType) ToPointer() *WireWithdrawalType {
	return &e
}
func (e *WireWithdrawalType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TYPE_UNSPECIFIED":
		fallthrough
	case "NORMAL":
		fallthrough
	case "DISABILITY":
		fallthrough
	case "SOSEPP":
		fallthrough
	case "PREMATURE":
		fallthrough
	case "DEATH":
		fallthrough
	case "EXCESS_CONTRIBUTION_REMOVAL_BEFORE_TAX_DEADLINE":
		fallthrough
	case "EXCESS_CONTRIBUTION_REMOVAL_AFTER_TAX_DEADLINE":
		fallthrough
	case "ROLLOVER_TO_QUALIFIED_PLAN":
		fallthrough
	case "ROLLOVER_TO_IRA":
		fallthrough
	case "DISTRIBUTION_TRANSFER":
		fallthrough
	case "RECHARACTERIZATION_PRIOR_YEAR":
		fallthrough
	case "RECHARACTERIZATION_CURRENT_YEAR":
		fallthrough
	case "DISTRIBUTION_CONVERSION":
		fallthrough
	case "MANAGEMENT_FEE":
		fallthrough
	case "PLAN_LOAN_401K":
		fallthrough
	case "PREMATURE_SIMPLE_IRA_LESS_THAN_2_YEARS":
		fallthrough
	case "NORMAL_ROTH_IRA_GREATER_THAN_5_YEARS":
		*e = WireWithdrawalType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WireWithdrawalType: %v", v)
	}
}

// WireWithdrawalIraDistribution - IRA distribution details for withdrawal from retirement account
type WireWithdrawalIraDistribution struct {
	// The federal tax withholding.
	FederalTaxWithholding *WireWithdrawalFederalTaxWithholding `json:"federal_tax_withholding,omitempty"`
	// The institution receiving retirement funds when performing a transfer to an identical retirement account type at a different financial institution. This is required for check and wire withdrawals because we can't always identify the institution using the transfer instructions. For cash journals this value will default to "Apex Clearing", regardless of what is passed in here
	ReceivingInstitution *string `json:"receiving_institution,omitempty"`
	// The state tax withholding.
	StateTaxWithholding *WireWithdrawalStateTaxWithholding `json:"state_tax_withholding,omitempty"`
	// Whether or not this distribution has a state withholding waiver.
	StateWithholdingWaiver *bool `json:"state_withholding_waiver,omitempty"`
	// Tax year for which the distribution is applied.
	TaxYear *int `json:"tax_year,omitempty"`
	// The type of retirement distribution.
	Type *WireWithdrawalType `json:"type,omitempty"`
}

func (o *WireWithdrawalIraDistribution) GetFederalTaxWithholding() *WireWithdrawalFederalTaxWithholding {
	if o == nil {
		return nil
	}
	return o.FederalTaxWithholding
}

func (o *WireWithdrawalIraDistribution) GetReceivingInstitution() *string {
	if o == nil {
		return nil
	}
	return o.ReceivingInstitution
}

func (o *WireWithdrawalIraDistribution) GetStateTaxWithholding() *WireWithdrawalStateTaxWithholding {
	if o == nil {
		return nil
	}
	return o.StateTaxWithholding
}

func (o *WireWithdrawalIraDistribution) GetStateWithholdingWaiver() *bool {
	if o == nil {
		return nil
	}
	return o.StateWithholdingWaiver
}

func (o *WireWithdrawalIraDistribution) GetTaxYear() *int {
	if o == nil {
		return nil
	}
	return o.TaxYear
}

func (o *WireWithdrawalIraDistribution) GetType() *WireWithdrawalType {
	if o == nil {
		return nil
	}
	return o.Type
}

// WireWithdrawalRecipientBankType - The type of bank identifier specified
type WireWithdrawalRecipientBankType string

const (
	WireWithdrawalRecipientBankTypeTypeUnspecified WireWithdrawalRecipientBankType = "TYPE_UNSPECIFIED"
	WireWithdrawalRecipientBankTypeAba             WireWithdrawalRecipientBankType = "ABA"
	WireWithdrawalRecipientBankTypeBic             WireWithdrawalRecipientBankType = "BIC"
)

func (e WireWithdrawalRecipientBankType) ToPointer() *WireWithdrawalRecipientBankType {
	return &e
}
func (e *WireWithdrawalRecipientBankType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TYPE_UNSPECIFIED":
		fallthrough
	case "ABA":
		fallthrough
	case "BIC":
		*e = WireWithdrawalRecipientBankType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WireWithdrawalRecipientBankType: %v", v)
	}
}

// WireWithdrawalBankID - An identifier that represents ABA routing number for domestic wire or BIC for foreign wire
type WireWithdrawalBankID struct {
	// The bank identifier
	ID *string `json:"id,omitempty"`
	// The type of bank identifier specified
	Type *WireWithdrawalRecipientBankType `json:"type,omitempty"`
}

func (o *WireWithdrawalBankID) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *WireWithdrawalBankID) GetType() *WireWithdrawalRecipientBankType {
	if o == nil {
		return nil
	}
	return o.Type
}

// WireWithdrawalRecipientBankAddress - The address of the recipient bank / financial institution
type WireWithdrawalRecipientBankAddress struct {
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

func (o *WireWithdrawalRecipientBankAddress) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *WireWithdrawalRecipientBankAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *WireWithdrawalRecipientBankAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *WireWithdrawalRecipientBankAddress) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *WireWithdrawalRecipientBankAddress) GetStreetAddress() []string {
	if o == nil {
		return nil
	}
	return o.StreetAddress
}

// WireWithdrawalInternationalBankDetails - Bank details required in the case of an international wire transfer
type WireWithdrawalInternationalBankDetails struct {
	// Any additional information to be communicated to the recipient bank, such as intermediary banks to be used.
	AdditionalInfo *string `json:"additional_info,omitempty"`
	// The address of the recipient bank / financial institution
	Address *WireWithdrawalRecipientBankAddress `json:"address,omitempty"`
	// The name of the recipient bank / financial institution
	BankName *string `json:"bank_name,omitempty"`
}

func (o *WireWithdrawalInternationalBankDetails) GetAdditionalInfo() *string {
	if o == nil {
		return nil
	}
	return o.AdditionalInfo
}

func (o *WireWithdrawalInternationalBankDetails) GetAddress() *WireWithdrawalRecipientBankAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *WireWithdrawalInternationalBankDetails) GetBankName() *string {
	if o == nil {
		return nil
	}
	return o.BankName
}

// WireWithdrawalRecipientBank - The recipient bank / financial institution
type WireWithdrawalRecipientBank struct {
	// An identifier that represents ABA routing number for domestic wire or BIC for foreign wire
	BankID *WireWithdrawalBankID `json:"bank_id,omitempty"`
	// Bank details required in the case of an international wire transfer
	InternationalBankDetails *WireWithdrawalInternationalBankDetails `json:"international_bank_details,omitempty"`
}

func (o *WireWithdrawalRecipientBank) GetBankID() *WireWithdrawalBankID {
	if o == nil {
		return nil
	}
	return o.BankID
}

func (o *WireWithdrawalRecipientBank) GetInternationalBankDetails() *WireWithdrawalInternationalBankDetails {
	if o == nil {
		return nil
	}
	return o.InternationalBankDetails
}

// WireWithdrawalStateState - The high level state of a transfer, one of:
// - `PROCESSING` - The transfer is being processed and will be posted if successful.
// - `PENDING_REVIEW` - The transfer is pending review and will continue processing if approved.
// - `POSTED` - The transfer has been posted to the ledger and will be completed at the end of the processing window if not canceled first.
// - `COMPLETED` - The transfer has been batched and completed.
// - `REJECTED` - The transfer was rejected.
// - `CANCELED` - The transfer was canceled.
// - `RETURNED` - The transfer was returned.
// - `POSTPONED` - The transfer is postponed and will resume processing during the next processing window.
type WireWithdrawalStateState string

const (
	WireWithdrawalStateStateStateUnspecified WireWithdrawalStateState = "STATE_UNSPECIFIED"
	WireWithdrawalStateStateProcessing       WireWithdrawalStateState = "PROCESSING"
	WireWithdrawalStateStatePendingReview    WireWithdrawalStateState = "PENDING_REVIEW"
	WireWithdrawalStateStatePosted           WireWithdrawalStateState = "POSTED"
	WireWithdrawalStateStateCompleted        WireWithdrawalStateState = "COMPLETED"
	WireWithdrawalStateStateRejected         WireWithdrawalStateState = "REJECTED"
	WireWithdrawalStateStateCanceled         WireWithdrawalStateState = "CANCELED"
	WireWithdrawalStateStateReturned         WireWithdrawalStateState = "RETURNED"
	WireWithdrawalStateStatePostponed        WireWithdrawalStateState = "POSTPONED"
)

func (e WireWithdrawalStateState) ToPointer() *WireWithdrawalStateState {
	return &e
}
func (e *WireWithdrawalStateState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STATE_UNSPECIFIED":
		fallthrough
	case "PROCESSING":
		fallthrough
	case "PENDING_REVIEW":
		fallthrough
	case "POSTED":
		fallthrough
	case "COMPLETED":
		fallthrough
	case "REJECTED":
		fallthrough
	case "CANCELED":
		fallthrough
	case "RETURNED":
		fallthrough
	case "POSTPONED":
		*e = WireWithdrawalStateState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WireWithdrawalStateState: %v", v)
	}
}

// WireWithdrawalState - The current state of the wire withdrawal
type WireWithdrawalState struct {
	// The user or service that triggered the state update.
	Actor *string `json:"actor,omitempty"`
	// Additional description of the transfer state.
	Message *string `json:"message,omitempty"`
	// Additional metadata relating to the transfer state. Included data depends on the state, e.g.:
	//  - Rejection reasons are included when the `state` is `REJECTED`
	//  - Reason and comment are included when `state` is `CANCELED`
	Metadata map[string]any `json:"metadata,omitempty"`
	// The high level state of a transfer, one of:
	// - `PROCESSING` - The transfer is being processed and will be posted if successful.
	// - `PENDING_REVIEW` - The transfer is pending review and will continue processing if approved.
	// - `POSTED` - The transfer has been posted to the ledger and will be completed at the end of the processing window if not canceled first.
	// - `COMPLETED` - The transfer has been batched and completed.
	// - `REJECTED` - The transfer was rejected.
	// - `CANCELED` - The transfer was canceled.
	// - `RETURNED` - The transfer was returned.
	// - `POSTPONED` - The transfer is postponed and will resume processing during the next processing window.
	State *WireWithdrawalStateState `json:"state,omitempty"`
	// The time of the state update.
	UpdateTime *time.Time `json:"update_time,omitempty"`
}

func (w WireWithdrawalState) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WireWithdrawalState) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WireWithdrawalState) GetActor() *string {
	if o == nil {
		return nil
	}
	return o.Actor
}

func (o *WireWithdrawalState) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *WireWithdrawalState) GetMetadata() map[string]any {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *WireWithdrawalState) GetState() *WireWithdrawalStateState {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *WireWithdrawalState) GetUpdateTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdateTime
}

// WireWithdrawal - A withdrawal transfer using the wire mechanism
type WireWithdrawal struct {
	// A cash amount in the format of decimal value
	Amount *WireWithdrawalAmount `json:"amount,omitempty"`
	// The beneficiary of the wire withdrawal
	Beneficiary *WireWithdrawalBeneficiary `json:"beneficiary,omitempty"`
	// External identifier supplied by the API caller. Each request must have a unique pairing of client_transfer_id and account
	ClientTransferID *string `json:"client_transfer_id,omitempty"`
	// The FedWire reference number for the withdrawal. Only set after the transfer is completed.
	FedReferenceNumber *string `json:"fed_reference_number,omitempty"`
	// The intermediary party
	Intermediary *WireWithdrawalIntermediary `json:"intermediary,omitempty"`
	// IRA distribution details for withdrawal from retirement account
	IraDistribution *WireWithdrawalIraDistribution `json:"ira_distribution,omitempty"`
	// The service generated name of the wire withdrawal
	Name *string `json:"name,omitempty"`
	// The recipient bank / financial institution
	RecipientBank *WireWithdrawalRecipientBank `json:"recipient_bank,omitempty"`
	// The current state of the wire withdrawal
	State *WireWithdrawalState `json:"state,omitempty"`
}

func (o *WireWithdrawal) GetAmount() *WireWithdrawalAmount {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *WireWithdrawal) GetBeneficiary() *WireWithdrawalBeneficiary {
	if o == nil {
		return nil
	}
	return o.Beneficiary
}

func (o *WireWithdrawal) GetClientTransferID() *string {
	if o == nil {
		return nil
	}
	return o.ClientTransferID
}

func (o *WireWithdrawal) GetFedReferenceNumber() *string {
	if o == nil {
		return nil
	}
	return o.FedReferenceNumber
}

func (o *WireWithdrawal) GetIntermediary() *WireWithdrawalIntermediary {
	if o == nil {
		return nil
	}
	return o.Intermediary
}

func (o *WireWithdrawal) GetIraDistribution() *WireWithdrawalIraDistribution {
	if o == nil {
		return nil
	}
	return o.IraDistribution
}

func (o *WireWithdrawal) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *WireWithdrawal) GetRecipientBank() *WireWithdrawalRecipientBank {
	if o == nil {
		return nil
	}
	return o.RecipientBank
}

func (o *WireWithdrawal) GetState() *WireWithdrawalState {
	if o == nil {
		return nil
	}
	return o.State
}
