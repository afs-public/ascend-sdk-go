// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// InvestmentObjective - The financial goal or purpose that an investor has in mind when making investment decisions; firms often ask investors to specify their investment objectives when opening an account, in order to provide appropriate investment recommendations and manage risk appropriately
type InvestmentObjective string

const (
	InvestmentObjectiveInvestmentObjectiveUnspecified InvestmentObjective = "INVESTMENT_OBJECTIVE_UNSPECIFIED"
	InvestmentObjectiveBalanced                       InvestmentObjective = "BALANCED"
	InvestmentObjectiveCapitalAppreciation            InvestmentObjective = "CAPITAL_APPRECIATION"
	InvestmentObjectiveCapitalPreservation            InvestmentObjective = "CAPITAL_PRESERVATION"
	InvestmentObjectiveGrowth                         InvestmentObjective = "GROWTH"
	InvestmentObjectiveGrowthAndIncome                InvestmentObjective = "GROWTH_AND_INCOME"
	InvestmentObjectiveGrowthIncome                   InvestmentObjective = "GROWTH_INCOME"
	InvestmentObjectiveIncome                         InvestmentObjective = "INCOME"
	InvestmentObjectiveLongTermGrowthWithGreaterRisk  InvestmentObjective = "LONG_TERM_GROWTH_WITH_GREATER_RISK"
	InvestmentObjectiveLongTermGrowthWithSafety       InvestmentObjective = "LONG_TERM_GROWTH_WITH_SAFETY"
	InvestmentObjectiveMaximumGrowth                  InvestmentObjective = "MAXIMUM_GROWTH"
	InvestmentObjectiveShortTermGrowthWithRisk        InvestmentObjective = "SHORT_TERM_GROWTH_WITH_RISK"
	InvestmentObjectiveSpeculation                    InvestmentObjective = "SPECULATION"
	InvestmentObjectiveOther                          InvestmentObjective = "OTHER"
)

func (e InvestmentObjective) ToPointer() *InvestmentObjective {
	return &e
}

// LiquidityNeeds - An investor’s short-term cash requirements or the need to access funds quickly; it is important to consider an investor’s liquidity needs to ensure that they have sufficient cash or easily liquidated assets available to meet their financial obligations - this may include holding cash or cash equivalents
type LiquidityNeeds string

const (
	LiquidityNeedsLiquidityNeedsUnspecified LiquidityNeeds = "LIQUIDITY_NEEDS_UNSPECIFIED"
	LiquidityNeedsVeryImportant             LiquidityNeeds = "VERY_IMPORTANT"
	LiquidityNeedsSomewhatImportant         LiquidityNeeds = "SOMEWHAT_IMPORTANT"
	LiquidityNeedsNotImportant              LiquidityNeeds = "NOT_IMPORTANT"
)

func (e LiquidityNeeds) ToPointer() *LiquidityNeeds {
	return &e
}

// RiskTolerance - An investor’s willingness and ability to tolerate risk when making investment decisions; reflects the investor’s comfort level with the potential ups and downs of the market and their ability to withstand potential losses
type RiskTolerance string

const (
	RiskToleranceRiskToleranceUnspecified RiskTolerance = "RISK_TOLERANCE_UNSPECIFIED"
	RiskToleranceLow                      RiskTolerance = "LOW"
	RiskToleranceMedium                   RiskTolerance = "MEDIUM"
	RiskToleranceHigh                     RiskTolerance = "HIGH"
)

func (e RiskTolerance) ToPointer() *RiskTolerance {
	return &e
}

// TimeHorizon - TThe length of time an investor expects to hold an investment before selling it; this can affect the appropriate asset allocation and risk level for the portfolio
type TimeHorizon string

const (
	TimeHorizonTimeHorizonUnspecified TimeHorizon = "TIME_HORIZON_UNSPECIFIED"
	TimeHorizonShort                  TimeHorizon = "SHORT"
	TimeHorizonAverage                TimeHorizon = "AVERAGE"
	TimeHorizonLong                   TimeHorizon = "LONG"
)

func (e TimeHorizon) ToPointer() *TimeHorizon {
	return &e
}

// AccountGoalsCreate - The account goals on an investor profile.
type AccountGoalsCreate struct {
	// The financial goal or purpose that an investor has in mind when making investment decisions; firms often ask investors to specify their investment objectives when opening an account, in order to provide appropriate investment recommendations and manage risk appropriately
	InvestmentObjective *InvestmentObjective `json:"investment_objective,omitempty"`
	// An investor’s short-term cash requirements or the need to access funds quickly; it is important to consider an investor’s liquidity needs to ensure that they have sufficient cash or easily liquidated assets available to meet their financial obligations - this may include holding cash or cash equivalents
	LiquidityNeeds *LiquidityNeeds `json:"liquidity_needs,omitempty"`
	// An investor’s willingness and ability to tolerate risk when making investment decisions; reflects the investor’s comfort level with the potential ups and downs of the market and their ability to withstand potential losses
	RiskTolerance *RiskTolerance `json:"risk_tolerance,omitempty"`
	// TThe length of time an investor expects to hold an investment before selling it; this can affect the appropriate asset allocation and risk level for the portfolio
	TimeHorizon *TimeHorizon `json:"time_horizon,omitempty"`
}

func (o *AccountGoalsCreate) GetInvestmentObjective() *InvestmentObjective {
	if o == nil {
		return nil
	}
	return o.InvestmentObjective
}

func (o *AccountGoalsCreate) GetLiquidityNeeds() *LiquidityNeeds {
	if o == nil {
		return nil
	}
	return o.LiquidityNeeds
}

func (o *AccountGoalsCreate) GetRiskTolerance() *RiskTolerance {
	if o == nil {
		return nil
	}
	return o.RiskTolerance
}

func (o *AccountGoalsCreate) GetTimeHorizon() *TimeHorizon {
	if o == nil {
		return nil
	}
	return o.TimeHorizon
}
