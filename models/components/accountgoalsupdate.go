// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// AccountGoalsUpdateInvestmentObjective - The financial goal or purpose that an investor has in mind when making investment decisions; firms often ask investors to specify their investment objectives when opening an account, in order to provide appropriate investment recommendations and manage risk appropriately
type AccountGoalsUpdateInvestmentObjective string

const (
	AccountGoalsUpdateInvestmentObjectiveInvestmentObjectiveUnspecified AccountGoalsUpdateInvestmentObjective = "INVESTMENT_OBJECTIVE_UNSPECIFIED"
	AccountGoalsUpdateInvestmentObjectiveBalanced                       AccountGoalsUpdateInvestmentObjective = "BALANCED"
	AccountGoalsUpdateInvestmentObjectiveCapitalAppreciation            AccountGoalsUpdateInvestmentObjective = "CAPITAL_APPRECIATION"
	AccountGoalsUpdateInvestmentObjectiveCapitalPreservation            AccountGoalsUpdateInvestmentObjective = "CAPITAL_PRESERVATION"
	AccountGoalsUpdateInvestmentObjectiveGrowth                         AccountGoalsUpdateInvestmentObjective = "GROWTH"
	AccountGoalsUpdateInvestmentObjectiveGrowthAndIncome                AccountGoalsUpdateInvestmentObjective = "GROWTH_AND_INCOME"
	AccountGoalsUpdateInvestmentObjectiveGrowthIncome                   AccountGoalsUpdateInvestmentObjective = "GROWTH_INCOME"
	AccountGoalsUpdateInvestmentObjectiveIncome                         AccountGoalsUpdateInvestmentObjective = "INCOME"
	AccountGoalsUpdateInvestmentObjectiveLongTermGrowthWithGreaterRisk  AccountGoalsUpdateInvestmentObjective = "LONG_TERM_GROWTH_WITH_GREATER_RISK"
	AccountGoalsUpdateInvestmentObjectiveLongTermGrowthWithSafety       AccountGoalsUpdateInvestmentObjective = "LONG_TERM_GROWTH_WITH_SAFETY"
	AccountGoalsUpdateInvestmentObjectiveMaximumGrowth                  AccountGoalsUpdateInvestmentObjective = "MAXIMUM_GROWTH"
	AccountGoalsUpdateInvestmentObjectiveShortTermGrowthWithRisk        AccountGoalsUpdateInvestmentObjective = "SHORT_TERM_GROWTH_WITH_RISK"
	AccountGoalsUpdateInvestmentObjectiveSpeculation                    AccountGoalsUpdateInvestmentObjective = "SPECULATION"
	AccountGoalsUpdateInvestmentObjectiveOther                          AccountGoalsUpdateInvestmentObjective = "OTHER"
)

func (e AccountGoalsUpdateInvestmentObjective) ToPointer() *AccountGoalsUpdateInvestmentObjective {
	return &e
}

// AccountGoalsUpdateLiquidityNeeds - An investor’s short-term cash requirements or the need to access funds quickly; it is important to consider an investor’s liquidity needs to ensure that they have sufficient cash or easily liquidated assets available to meet their financial obligations - this may include holding cash or cash equivalents
type AccountGoalsUpdateLiquidityNeeds string

const (
	AccountGoalsUpdateLiquidityNeedsLiquidityNeedsUnspecified AccountGoalsUpdateLiquidityNeeds = "LIQUIDITY_NEEDS_UNSPECIFIED"
	AccountGoalsUpdateLiquidityNeedsVeryImportant             AccountGoalsUpdateLiquidityNeeds = "VERY_IMPORTANT"
	AccountGoalsUpdateLiquidityNeedsSomewhatImportant         AccountGoalsUpdateLiquidityNeeds = "SOMEWHAT_IMPORTANT"
	AccountGoalsUpdateLiquidityNeedsNotImportant              AccountGoalsUpdateLiquidityNeeds = "NOT_IMPORTANT"
)

func (e AccountGoalsUpdateLiquidityNeeds) ToPointer() *AccountGoalsUpdateLiquidityNeeds {
	return &e
}

// AccountGoalsUpdateRiskTolerance - An investor’s willingness and ability to tolerate risk when making investment decisions; reflects the investor’s comfort level with the potential ups and downs of the market and their ability to withstand potential losses
type AccountGoalsUpdateRiskTolerance string

const (
	AccountGoalsUpdateRiskToleranceRiskToleranceUnspecified AccountGoalsUpdateRiskTolerance = "RISK_TOLERANCE_UNSPECIFIED"
	AccountGoalsUpdateRiskToleranceLow                      AccountGoalsUpdateRiskTolerance = "LOW"
	AccountGoalsUpdateRiskToleranceMedium                   AccountGoalsUpdateRiskTolerance = "MEDIUM"
	AccountGoalsUpdateRiskToleranceHigh                     AccountGoalsUpdateRiskTolerance = "HIGH"
)

func (e AccountGoalsUpdateRiskTolerance) ToPointer() *AccountGoalsUpdateRiskTolerance {
	return &e
}

// AccountGoalsUpdateTimeHorizon - TThe length of time an investor expects to hold an investment before selling it; this can affect the appropriate asset allocation and risk level for the portfolio
type AccountGoalsUpdateTimeHorizon string

const (
	AccountGoalsUpdateTimeHorizonTimeHorizonUnspecified AccountGoalsUpdateTimeHorizon = "TIME_HORIZON_UNSPECIFIED"
	AccountGoalsUpdateTimeHorizonShort                  AccountGoalsUpdateTimeHorizon = "SHORT"
	AccountGoalsUpdateTimeHorizonAverage                AccountGoalsUpdateTimeHorizon = "AVERAGE"
	AccountGoalsUpdateTimeHorizonLong                   AccountGoalsUpdateTimeHorizon = "LONG"
)

func (e AccountGoalsUpdateTimeHorizon) ToPointer() *AccountGoalsUpdateTimeHorizon {
	return &e
}

// AccountGoalsUpdate - The account goals on an investor profile.
type AccountGoalsUpdate struct {
	// The financial goal or purpose that an investor has in mind when making investment decisions; firms often ask investors to specify their investment objectives when opening an account, in order to provide appropriate investment recommendations and manage risk appropriately
	InvestmentObjective *AccountGoalsUpdateInvestmentObjective `json:"investment_objective,omitempty"`
	// An investor’s short-term cash requirements or the need to access funds quickly; it is important to consider an investor’s liquidity needs to ensure that they have sufficient cash or easily liquidated assets available to meet their financial obligations - this may include holding cash or cash equivalents
	LiquidityNeeds *AccountGoalsUpdateLiquidityNeeds `json:"liquidity_needs,omitempty"`
	// An investor’s willingness and ability to tolerate risk when making investment decisions; reflects the investor’s comfort level with the potential ups and downs of the market and their ability to withstand potential losses
	RiskTolerance *AccountGoalsUpdateRiskTolerance `json:"risk_tolerance,omitempty"`
	// TThe length of time an investor expects to hold an investment before selling it; this can affect the appropriate asset allocation and risk level for the portfolio
	TimeHorizon *AccountGoalsUpdateTimeHorizon `json:"time_horizon,omitempty"`
}

func (o *AccountGoalsUpdate) GetInvestmentObjective() *AccountGoalsUpdateInvestmentObjective {
	if o == nil {
		return nil
	}
	return o.InvestmentObjective
}

func (o *AccountGoalsUpdate) GetLiquidityNeeds() *AccountGoalsUpdateLiquidityNeeds {
	if o == nil {
		return nil
	}
	return o.LiquidityNeeds
}

func (o *AccountGoalsUpdate) GetRiskTolerance() *AccountGoalsUpdateRiskTolerance {
	if o == nil {
		return nil
	}
	return o.RiskTolerance
}

func (o *AccountGoalsUpdate) GetTimeHorizon() *AccountGoalsUpdateTimeHorizon {
	if o == nil {
		return nil
	}
	return o.TimeHorizon
}
