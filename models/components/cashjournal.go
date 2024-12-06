// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/afs-public/ascend-sdk-go/internal/utils"
)

// CashJournalAmount - The amount to transfer in USD
type CashJournalAmount struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *CashJournalAmount) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// PartyType - Whether a cash journal is first party or third party Determined asynchronously when the transfer is processing, and will be set by the time the transfer is posted
type PartyType string

const (
	PartyTypePartyTypeUnspecified PartyType = "PARTY_TYPE_UNSPECIFIED"
	PartyTypeFirstParty           PartyType = "FIRST_PARTY"
	PartyTypeThirdParty           PartyType = "THIRD_PARTY"
)

func (e PartyType) ToPointer() *PartyType {
	return &e
}
func (e *PartyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PARTY_TYPE_UNSPECIFIED":
		fallthrough
	case "FIRST_PARTY":
		fallthrough
	case "THIRD_PARTY":
		*e = PartyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PartyType: %v", v)
	}
}

// CashJournalType - The type of retirement contribution.
type CashJournalType string

const (
	CashJournalTypeTypeUnspecified    CashJournalType = "TYPE_UNSPECIFIED"
	CashJournalTypeRegular            CashJournalType = "REGULAR"
	CashJournalTypeEmployee           CashJournalType = "EMPLOYEE"
	CashJournalTypeEmployer           CashJournalType = "EMPLOYER"
	CashJournalTypeRecharacterization CashJournalType = "RECHARACTERIZATION"
	CashJournalTypeRollover60Day      CashJournalType = "ROLLOVER_60_DAY"
	CashJournalTypeRolloverDirect     CashJournalType = "ROLLOVER_DIRECT"
	CashJournalTypeTransfer           CashJournalType = "TRANSFER"
	CashJournalTypeTrusteeFee         CashJournalType = "TRUSTEE_FEE"
	CashJournalTypeConversion         CashJournalType = "CONVERSION"
	CashJournalTypeRepayment          CashJournalType = "REPAYMENT"
)

func (e CashJournalType) ToPointer() *CashJournalType {
	return &e
}
func (e *CashJournalType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TYPE_UNSPECIFIED":
		fallthrough
	case "REGULAR":
		fallthrough
	case "EMPLOYEE":
		fallthrough
	case "EMPLOYER":
		fallthrough
	case "RECHARACTERIZATION":
		fallthrough
	case "ROLLOVER_60_DAY":
		fallthrough
	case "ROLLOVER_DIRECT":
		fallthrough
	case "TRANSFER":
		fallthrough
	case "TRUSTEE_FEE":
		fallthrough
	case "CONVERSION":
		fallthrough
	case "REPAYMENT":
		*e = CashJournalType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CashJournalType: %v", v)
	}
}

// CashJournalRetirementContribution - The retirement contribution details Must be provided when the destination account is a retirement account
type CashJournalRetirementContribution struct {
	// Tax year for which the contribution is applied. Current year is always valid; prior year is only valid before tax deadline. Must be in "YYYY" format.
	TaxYear *int `json:"tax_year,omitempty"`
	// The type of retirement contribution.
	Type *CashJournalType `json:"type,omitempty"`
}

func (o *CashJournalRetirementContribution) GetTaxYear() *int {
	if o == nil {
		return nil
	}
	return o.TaxYear
}

func (o *CashJournalRetirementContribution) GetType() *CashJournalType {
	if o == nil {
		return nil
	}
	return o.Type
}

// CashJournalRetirementDistributionAmount - Fixed USD amount to withhold for taxes.
type CashJournalRetirementDistributionAmount struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *CashJournalRetirementDistributionAmount) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// CashJournalPercentage - Percentage of total disbursement amount to withhold for taxes.
type CashJournalPercentage struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *CashJournalPercentage) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// CashJournalFederalTaxWithholding - The federal tax withholding.
type CashJournalFederalTaxWithholding struct {
	// Fixed USD amount to withhold for taxes.
	Amount *CashJournalRetirementDistributionAmount `json:"amount,omitempty"`
	// Percentage of total disbursement amount to withhold for taxes.
	Percentage *CashJournalPercentage `json:"percentage,omitempty"`
}

func (o *CashJournalFederalTaxWithholding) GetAmount() *CashJournalRetirementDistributionAmount {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *CashJournalFederalTaxWithholding) GetPercentage() *CashJournalPercentage {
	if o == nil {
		return nil
	}
	return o.Percentage
}

// CashJournalRetirementDistributionStateTaxWithholdingAmount - Fixed USD amount to withhold for taxes.
type CashJournalRetirementDistributionStateTaxWithholdingAmount struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *CashJournalRetirementDistributionStateTaxWithholdingAmount) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// CashJournalRetirementDistributionPercentage - Percentage of total disbursement amount to withhold for taxes.
type CashJournalRetirementDistributionPercentage struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *CashJournalRetirementDistributionPercentage) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// CashJournalStateTaxWithholding - The state tax withholding.
type CashJournalStateTaxWithholding struct {
	// Fixed USD amount to withhold for taxes.
	Amount *CashJournalRetirementDistributionStateTaxWithholdingAmount `json:"amount,omitempty"`
	// Percentage of total disbursement amount to withhold for taxes.
	Percentage *CashJournalRetirementDistributionPercentage `json:"percentage,omitempty"`
}

func (o *CashJournalStateTaxWithholding) GetAmount() *CashJournalRetirementDistributionStateTaxWithholdingAmount {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *CashJournalStateTaxWithholding) GetPercentage() *CashJournalRetirementDistributionPercentage {
	if o == nil {
		return nil
	}
	return o.Percentage
}

// CashJournalRetirementDistributionType - The type of retirement distribution.
type CashJournalRetirementDistributionType string

const (
	CashJournalRetirementDistributionTypeTypeUnspecified                            CashJournalRetirementDistributionType = "TYPE_UNSPECIFIED"
	CashJournalRetirementDistributionTypeNormal                                     CashJournalRetirementDistributionType = "NORMAL"
	CashJournalRetirementDistributionTypeDisability                                 CashJournalRetirementDistributionType = "DISABILITY"
	CashJournalRetirementDistributionTypeSosepp                                     CashJournalRetirementDistributionType = "SOSEPP"
	CashJournalRetirementDistributionTypePremature                                  CashJournalRetirementDistributionType = "PREMATURE"
	CashJournalRetirementDistributionTypeDeath                                      CashJournalRetirementDistributionType = "DEATH"
	CashJournalRetirementDistributionTypeExcessContributionRemovalBeforeTaxDeadline CashJournalRetirementDistributionType = "EXCESS_CONTRIBUTION_REMOVAL_BEFORE_TAX_DEADLINE"
	CashJournalRetirementDistributionTypeExcessContributionRemovalAfterTaxDeadline  CashJournalRetirementDistributionType = "EXCESS_CONTRIBUTION_REMOVAL_AFTER_TAX_DEADLINE"
	CashJournalRetirementDistributionTypeRolloverToQualifiedPlan                    CashJournalRetirementDistributionType = "ROLLOVER_TO_QUALIFIED_PLAN"
	CashJournalRetirementDistributionTypeRolloverToIra                              CashJournalRetirementDistributionType = "ROLLOVER_TO_IRA"
	CashJournalRetirementDistributionTypeDistributionTransfer                       CashJournalRetirementDistributionType = "DISTRIBUTION_TRANSFER"
	CashJournalRetirementDistributionTypeRecharacterizationPriorYear                CashJournalRetirementDistributionType = "RECHARACTERIZATION_PRIOR_YEAR"
	CashJournalRetirementDistributionTypeRecharacterizationCurrentYear              CashJournalRetirementDistributionType = "RECHARACTERIZATION_CURRENT_YEAR"
	CashJournalRetirementDistributionTypeDistributionConversion                     CashJournalRetirementDistributionType = "DISTRIBUTION_CONVERSION"
	CashJournalRetirementDistributionTypeManagementFee                              CashJournalRetirementDistributionType = "MANAGEMENT_FEE"
	CashJournalRetirementDistributionTypePlanLoan401K                               CashJournalRetirementDistributionType = "PLAN_LOAN_401K"
	CashJournalRetirementDistributionTypePrematureSimpleIraLessThan2Years           CashJournalRetirementDistributionType = "PREMATURE_SIMPLE_IRA_LESS_THAN_2_YEARS"
	CashJournalRetirementDistributionTypeNormalRothIraGreaterThan5Years             CashJournalRetirementDistributionType = "NORMAL_ROTH_IRA_GREATER_THAN_5_YEARS"
)

func (e CashJournalRetirementDistributionType) ToPointer() *CashJournalRetirementDistributionType {
	return &e
}
func (e *CashJournalRetirementDistributionType) UnmarshalJSON(data []byte) error {
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
		*e = CashJournalRetirementDistributionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CashJournalRetirementDistributionType: %v", v)
	}
}

// CashJournalRetirementDistribution - The retirement distribution details Must be provided when the source account is a retirement account
type CashJournalRetirementDistribution struct {
	// The federal tax withholding.
	FederalTaxWithholding *CashJournalFederalTaxWithholding `json:"federal_tax_withholding,omitempty"`
	// The institution receiving retirement funds when performing a transfer to an identical retirement account type at a different financial institution. This is required for check and wire withdrawals because we can't always identify the institution using the transfer instructions. For cash journals this value will default to "Apex Clearing", regardless of what is passed in here
	ReceivingInstitution *string `json:"receiving_institution,omitempty"`
	// The state tax withholding.
	StateTaxWithholding *CashJournalStateTaxWithholding `json:"state_tax_withholding,omitempty"`
	// Whether or not this distribution has a state withholding waiver.
	StateWithholdingWaiver *bool `json:"state_withholding_waiver,omitempty"`
	// Tax year for which the distribution is applied.
	TaxYear *int `json:"tax_year,omitempty"`
	// The type of retirement distribution.
	Type *CashJournalRetirementDistributionType `json:"type,omitempty"`
}

func (o *CashJournalRetirementDistribution) GetFederalTaxWithholding() *CashJournalFederalTaxWithholding {
	if o == nil {
		return nil
	}
	return o.FederalTaxWithholding
}

func (o *CashJournalRetirementDistribution) GetReceivingInstitution() *string {
	if o == nil {
		return nil
	}
	return o.ReceivingInstitution
}

func (o *CashJournalRetirementDistribution) GetStateTaxWithholding() *CashJournalStateTaxWithholding {
	if o == nil {
		return nil
	}
	return o.StateTaxWithholding
}

func (o *CashJournalRetirementDistribution) GetStateWithholdingWaiver() *bool {
	if o == nil {
		return nil
	}
	return o.StateWithholdingWaiver
}

func (o *CashJournalRetirementDistribution) GetTaxYear() *int {
	if o == nil {
		return nil
	}
	return o.TaxYear
}

func (o *CashJournalRetirementDistribution) GetType() *CashJournalRetirementDistributionType {
	if o == nil {
		return nil
	}
	return o.Type
}

// CashJournalStateState - The high level state of a transfer, one of:
// - `PROCESSING` - The transfer is being processed and will be posted if successful.
// - `PENDING_REVIEW` - The transfer is pending review and will continue processing if approved.
// - `POSTED` - The transfer has been posted to the ledger and will be completed at the end of the processing window if not canceled first.
// - `COMPLETED` - The transfer has been batched and completed.
// - `REJECTED` - The transfer was rejected.
// - `CANCELED` - The transfer was canceled.
// - `RETURNED` - The transfer was returned.
// - `POSTPONED` - The transfer is postponed and will resume processing during the next processing window.
type CashJournalStateState string

const (
	CashJournalStateStateStateUnspecified CashJournalStateState = "STATE_UNSPECIFIED"
	CashJournalStateStateProcessing       CashJournalStateState = "PROCESSING"
	CashJournalStateStatePendingReview    CashJournalStateState = "PENDING_REVIEW"
	CashJournalStateStatePosted           CashJournalStateState = "POSTED"
	CashJournalStateStateCompleted        CashJournalStateState = "COMPLETED"
	CashJournalStateStateRejected         CashJournalStateState = "REJECTED"
	CashJournalStateStateCanceled         CashJournalStateState = "CANCELED"
	CashJournalStateStateReturned         CashJournalStateState = "RETURNED"
	CashJournalStateStatePostponed        CashJournalStateState = "POSTPONED"
)

func (e CashJournalStateState) ToPointer() *CashJournalStateState {
	return &e
}
func (e *CashJournalStateState) UnmarshalJSON(data []byte) error {
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
		*e = CashJournalStateState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CashJournalStateState: %v", v)
	}
}

// CashJournalState - The current state of the cash journal
type CashJournalState struct {
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
	State *CashJournalStateState `json:"state,omitempty"`
	// The time of the state update.
	UpdateTime *time.Time `json:"update_time,omitempty"`
}

func (c CashJournalState) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CashJournalState) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CashJournalState) GetActor() *string {
	if o == nil {
		return nil
	}
	return o.Actor
}

func (o *CashJournalState) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CashJournalState) GetMetadata() map[string]any {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *CashJournalState) GetState() *CashJournalStateState {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *CashJournalState) GetUpdateTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdateTime
}

// CashJournal - A cash journal transfer. Funds are moved from a source account to a destination account
type CashJournal struct {
	// The amount to transfer in USD
	Amount *CashJournalAmount `json:"amount,omitempty"`
	// The external identifier supplied by the API caller Each request must have a unique pairing of `client_transfer_id` and `source_account`
	ClientTransferID *string `json:"client_transfer_id,omitempty"`
	// The account that funds will be moved to
	DestinationAccount *string `json:"destination_account,omitempty"`
	// The resource name of the cash journal
	Name *string `json:"name,omitempty"`
	// Whether a cash journal is first party or third party Determined asynchronously when the transfer is processing, and will be set by the time the transfer is posted
	PartyType *PartyType `json:"party_type,omitempty"`
	// The retirement contribution details Must be provided when the destination account is a retirement account
	RetirementContribution *CashJournalRetirementContribution `json:"retirement_contribution,omitempty"`
	// The retirement distribution details Must be provided when the source account is a retirement account
	RetirementDistribution *CashJournalRetirementDistribution `json:"retirement_distribution,omitempty"`
	// The account that funds will be moved from
	SourceAccount *string `json:"source_account,omitempty"`
	// The current state of the cash journal
	State *CashJournalState `json:"state,omitempty"`
}

func (o *CashJournal) GetAmount() *CashJournalAmount {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *CashJournal) GetClientTransferID() *string {
	if o == nil {
		return nil
	}
	return o.ClientTransferID
}

func (o *CashJournal) GetDestinationAccount() *string {
	if o == nil {
		return nil
	}
	return o.DestinationAccount
}

func (o *CashJournal) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CashJournal) GetPartyType() *PartyType {
	if o == nil {
		return nil
	}
	return o.PartyType
}

func (o *CashJournal) GetRetirementContribution() *CashJournalRetirementContribution {
	if o == nil {
		return nil
	}
	return o.RetirementContribution
}

func (o *CashJournal) GetRetirementDistribution() *CashJournalRetirementDistribution {
	if o == nil {
		return nil
	}
	return o.RetirementDistribution
}

func (o *CashJournal) GetSourceAccount() *string {
	if o == nil {
		return nil
	}
	return o.SourceAccount
}

func (o *CashJournal) GetState() *CashJournalState {
	if o == nil {
		return nil
	}
	return o.State
}