# CustomerProfileCreate

A detailed summary of financial and personal details of an investor, to help understand the investor's financial standing, investment experience and risk tolerance.


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             | Example                                                                                 |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `AnnualIncomeRangeUsd`                                                                  | [*components.AnnualIncomeRangeUsd](../../models/components/annualincomerangeusd.md)     | :heavy_minus_sign:                                                                      | Annual income range; the low number is exclusive, the high number is inclusive          | FROM_100K_TO_200K                                                                       |
| `FederalTaxBracket`                                                                     | **float64*                                                                              | :heavy_minus_sign:                                                                      | Federal tax bracket percent.                                                            | 1.5                                                                                     |
| `InvestmentExperience`                                                                  | [*components.InvestmentExperience](../../models/components/investmentexperience.md)     | :heavy_minus_sign:                                                                      | Investment experience.                                                                  | GOOD                                                                                    |
| `LiquidNetWorthRangeUsd`                                                                | [*components.LiquidNetWorthRangeUsd](../../models/components/liquidnetworthrangeusd.md) | :heavy_minus_sign:                                                                      | Liquid net worth range; the low number is exclusive, the high number is inclusive       | FROM_100K_TO_200K                                                                       |
| `TotalNetWorthRangeUsd`                                                                 | [*components.TotalNetWorthRangeUsd](../../models/components/totalnetworthrangeusd.md)   | :heavy_minus_sign:                                                                      | Total net worth range; the low number is exclusive, the high number is inclusive        | FROM_100K_TO_200K                                                                       |