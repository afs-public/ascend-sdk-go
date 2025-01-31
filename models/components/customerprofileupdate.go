// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// CustomerProfileUpdateAnnualIncomeRangeUsd - Annual income range; the low number is exclusive, the high number is inclusive
type CustomerProfileUpdateAnnualIncomeRangeUsd string

const (
	CustomerProfileUpdateAnnualIncomeRangeUsdUsdRangeUnspecified CustomerProfileUpdateAnnualIncomeRangeUsd = "USD_RANGE_UNSPECIFIED"
	CustomerProfileUpdateAnnualIncomeRangeUsdUnder25K            CustomerProfileUpdateAnnualIncomeRangeUsd = "UNDER_25K"
	CustomerProfileUpdateAnnualIncomeRangeUsdFrom25KTo50K        CustomerProfileUpdateAnnualIncomeRangeUsd = "FROM_25K_TO_50K"
	CustomerProfileUpdateAnnualIncomeRangeUsdFrom50KTo100K       CustomerProfileUpdateAnnualIncomeRangeUsd = "FROM_50K_TO_100K"
	CustomerProfileUpdateAnnualIncomeRangeUsdFrom100KTo200K      CustomerProfileUpdateAnnualIncomeRangeUsd = "FROM_100K_TO_200K"
	CustomerProfileUpdateAnnualIncomeRangeUsdFrom200KTo300K      CustomerProfileUpdateAnnualIncomeRangeUsd = "FROM_200K_TO_300K"
	CustomerProfileUpdateAnnualIncomeRangeUsdFrom300KTo500K      CustomerProfileUpdateAnnualIncomeRangeUsd = "FROM_300K_TO_500K"
	CustomerProfileUpdateAnnualIncomeRangeUsdFrom500KTo1M        CustomerProfileUpdateAnnualIncomeRangeUsd = "FROM_500K_TO_1M"
	CustomerProfileUpdateAnnualIncomeRangeUsdFrom1MTo5M          CustomerProfileUpdateAnnualIncomeRangeUsd = "FROM_1M_TO_5M"
	CustomerProfileUpdateAnnualIncomeRangeUsdOver5M              CustomerProfileUpdateAnnualIncomeRangeUsd = "OVER_5M"
)

func (e CustomerProfileUpdateAnnualIncomeRangeUsd) ToPointer() *CustomerProfileUpdateAnnualIncomeRangeUsd {
	return &e
}

// CustomerProfileUpdateInvestmentExperience - Investment experience.
type CustomerProfileUpdateInvestmentExperience string

const (
	CustomerProfileUpdateInvestmentExperienceInvestmentExperienceUnspecified CustomerProfileUpdateInvestmentExperience = "INVESTMENT_EXPERIENCE_UNSPECIFIED"
	CustomerProfileUpdateInvestmentExperienceNone                            CustomerProfileUpdateInvestmentExperience = "NONE"
	CustomerProfileUpdateInvestmentExperienceLimited                         CustomerProfileUpdateInvestmentExperience = "LIMITED"
	CustomerProfileUpdateInvestmentExperienceGood                            CustomerProfileUpdateInvestmentExperience = "GOOD"
	CustomerProfileUpdateInvestmentExperienceExtensive                       CustomerProfileUpdateInvestmentExperience = "EXTENSIVE"
)

func (e CustomerProfileUpdateInvestmentExperience) ToPointer() *CustomerProfileUpdateInvestmentExperience {
	return &e
}

// CustomerProfileUpdateLiquidNetWorthRangeUsd - Liquid net worth range; the low number is exclusive, the high number is inclusive
type CustomerProfileUpdateLiquidNetWorthRangeUsd string

const (
	CustomerProfileUpdateLiquidNetWorthRangeUsdUsdRangeUnspecified CustomerProfileUpdateLiquidNetWorthRangeUsd = "USD_RANGE_UNSPECIFIED"
	CustomerProfileUpdateLiquidNetWorthRangeUsdUnder25K            CustomerProfileUpdateLiquidNetWorthRangeUsd = "UNDER_25K"
	CustomerProfileUpdateLiquidNetWorthRangeUsdFrom25KTo50K        CustomerProfileUpdateLiquidNetWorthRangeUsd = "FROM_25K_TO_50K"
	CustomerProfileUpdateLiquidNetWorthRangeUsdFrom50KTo100K       CustomerProfileUpdateLiquidNetWorthRangeUsd = "FROM_50K_TO_100K"
	CustomerProfileUpdateLiquidNetWorthRangeUsdFrom100KTo200K      CustomerProfileUpdateLiquidNetWorthRangeUsd = "FROM_100K_TO_200K"
	CustomerProfileUpdateLiquidNetWorthRangeUsdFrom200KTo300K      CustomerProfileUpdateLiquidNetWorthRangeUsd = "FROM_200K_TO_300K"
	CustomerProfileUpdateLiquidNetWorthRangeUsdFrom300KTo500K      CustomerProfileUpdateLiquidNetWorthRangeUsd = "FROM_300K_TO_500K"
	CustomerProfileUpdateLiquidNetWorthRangeUsdFrom500KTo1M        CustomerProfileUpdateLiquidNetWorthRangeUsd = "FROM_500K_TO_1M"
	CustomerProfileUpdateLiquidNetWorthRangeUsdFrom1MTo5M          CustomerProfileUpdateLiquidNetWorthRangeUsd = "FROM_1M_TO_5M"
	CustomerProfileUpdateLiquidNetWorthRangeUsdOver5M              CustomerProfileUpdateLiquidNetWorthRangeUsd = "OVER_5M"
)

func (e CustomerProfileUpdateLiquidNetWorthRangeUsd) ToPointer() *CustomerProfileUpdateLiquidNetWorthRangeUsd {
	return &e
}

// CustomerProfileUpdateTotalNetWorthRangeUsd - Total net worth range; the low number is exclusive, the high number is inclusive
type CustomerProfileUpdateTotalNetWorthRangeUsd string

const (
	CustomerProfileUpdateTotalNetWorthRangeUsdUsdRangeUnspecified CustomerProfileUpdateTotalNetWorthRangeUsd = "USD_RANGE_UNSPECIFIED"
	CustomerProfileUpdateTotalNetWorthRangeUsdUnder25K            CustomerProfileUpdateTotalNetWorthRangeUsd = "UNDER_25K"
	CustomerProfileUpdateTotalNetWorthRangeUsdFrom25KTo50K        CustomerProfileUpdateTotalNetWorthRangeUsd = "FROM_25K_TO_50K"
	CustomerProfileUpdateTotalNetWorthRangeUsdFrom50KTo100K       CustomerProfileUpdateTotalNetWorthRangeUsd = "FROM_50K_TO_100K"
	CustomerProfileUpdateTotalNetWorthRangeUsdFrom100KTo200K      CustomerProfileUpdateTotalNetWorthRangeUsd = "FROM_100K_TO_200K"
	CustomerProfileUpdateTotalNetWorthRangeUsdFrom200KTo300K      CustomerProfileUpdateTotalNetWorthRangeUsd = "FROM_200K_TO_300K"
	CustomerProfileUpdateTotalNetWorthRangeUsdFrom300KTo500K      CustomerProfileUpdateTotalNetWorthRangeUsd = "FROM_300K_TO_500K"
	CustomerProfileUpdateTotalNetWorthRangeUsdFrom500KTo1M        CustomerProfileUpdateTotalNetWorthRangeUsd = "FROM_500K_TO_1M"
	CustomerProfileUpdateTotalNetWorthRangeUsdFrom1MTo5M          CustomerProfileUpdateTotalNetWorthRangeUsd = "FROM_1M_TO_5M"
	CustomerProfileUpdateTotalNetWorthRangeUsdOver5M              CustomerProfileUpdateTotalNetWorthRangeUsd = "OVER_5M"
)

func (e CustomerProfileUpdateTotalNetWorthRangeUsd) ToPointer() *CustomerProfileUpdateTotalNetWorthRangeUsd {
	return &e
}

// CustomerProfileUpdate - A detailed summary of financial and personal details of an investor, to help understand the investor's financial standing, investment experience and risk tolerance.
type CustomerProfileUpdate struct {
	// Annual income range; the low number is exclusive, the high number is inclusive
	AnnualIncomeRangeUsd *CustomerProfileUpdateAnnualIncomeRangeUsd `json:"annual_income_range_usd,omitempty"`
	// Federal tax bracket percent.
	FederalTaxBracket *float64 `json:"federal_tax_bracket,omitempty"`
	// Investment experience.
	InvestmentExperience *CustomerProfileUpdateInvestmentExperience `json:"investment_experience,omitempty"`
	// Liquid net worth range; the low number is exclusive, the high number is inclusive
	LiquidNetWorthRangeUsd *CustomerProfileUpdateLiquidNetWorthRangeUsd `json:"liquid_net_worth_range_usd,omitempty"`
	// Total net worth range; the low number is exclusive, the high number is inclusive
	TotalNetWorthRangeUsd *CustomerProfileUpdateTotalNetWorthRangeUsd `json:"total_net_worth_range_usd,omitempty"`
}

func (o *CustomerProfileUpdate) GetAnnualIncomeRangeUsd() *CustomerProfileUpdateAnnualIncomeRangeUsd {
	if o == nil {
		return nil
	}
	return o.AnnualIncomeRangeUsd
}

func (o *CustomerProfileUpdate) GetFederalTaxBracket() *float64 {
	if o == nil {
		return nil
	}
	return o.FederalTaxBracket
}

func (o *CustomerProfileUpdate) GetInvestmentExperience() *CustomerProfileUpdateInvestmentExperience {
	if o == nil {
		return nil
	}
	return o.InvestmentExperience
}

func (o *CustomerProfileUpdate) GetLiquidNetWorthRangeUsd() *CustomerProfileUpdateLiquidNetWorthRangeUsd {
	if o == nil {
		return nil
	}
	return o.LiquidNetWorthRangeUsd
}

func (o *CustomerProfileUpdate) GetTotalNetWorthRangeUsd() *CustomerProfileUpdateTotalNetWorthRangeUsd {
	if o == nil {
		return nil
	}
	return o.TotalNetWorthRangeUsd
}
