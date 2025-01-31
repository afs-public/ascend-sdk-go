// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ContributionConstraintsContributionTypeInfoType - Retirement contribution type
type ContributionConstraintsContributionTypeInfoType string

const (
	ContributionConstraintsContributionTypeInfoTypeTypeUnspecified    ContributionConstraintsContributionTypeInfoType = "TYPE_UNSPECIFIED"
	ContributionConstraintsContributionTypeInfoTypeRegular            ContributionConstraintsContributionTypeInfoType = "REGULAR"
	ContributionConstraintsContributionTypeInfoTypeEmployee           ContributionConstraintsContributionTypeInfoType = "EMPLOYEE"
	ContributionConstraintsContributionTypeInfoTypeEmployer           ContributionConstraintsContributionTypeInfoType = "EMPLOYER"
	ContributionConstraintsContributionTypeInfoTypeRecharacterization ContributionConstraintsContributionTypeInfoType = "RECHARACTERIZATION"
	ContributionConstraintsContributionTypeInfoTypeRollover60Day      ContributionConstraintsContributionTypeInfoType = "ROLLOVER_60_DAY"
	ContributionConstraintsContributionTypeInfoTypeRolloverDirect     ContributionConstraintsContributionTypeInfoType = "ROLLOVER_DIRECT"
	ContributionConstraintsContributionTypeInfoTypeTransfer           ContributionConstraintsContributionTypeInfoType = "TRANSFER"
	ContributionConstraintsContributionTypeInfoTypeTrusteeFee         ContributionConstraintsContributionTypeInfoType = "TRUSTEE_FEE"
	ContributionConstraintsContributionTypeInfoTypeConversion         ContributionConstraintsContributionTypeInfoType = "CONVERSION"
	ContributionConstraintsContributionTypeInfoTypeRepayment          ContributionConstraintsContributionTypeInfoType = "REPAYMENT"
)

func (e ContributionConstraintsContributionTypeInfoType) ToPointer() *ContributionConstraintsContributionTypeInfoType {
	return &e
}

// ContributionConstraintsContributionTypeInfo - Retirement contribution type info
type ContributionConstraintsContributionTypeInfo struct {
	// Whether this specific retirement contribution may be allowed for the previous year, without consideration of the tax deadline
	PreviousYearAllowed *bool `json:"previous_year_allowed,omitempty"`
	// Retirement contribution type
	Type *ContributionConstraintsContributionTypeInfoType `json:"type,omitempty"`
}

func (o *ContributionConstraintsContributionTypeInfo) GetPreviousYearAllowed() *bool {
	if o == nil {
		return nil
	}
	return o.PreviousYearAllowed
}

func (o *ContributionConstraintsContributionTypeInfo) GetType() *ContributionConstraintsContributionTypeInfoType {
	if o == nil {
		return nil
	}
	return o.Type
}
