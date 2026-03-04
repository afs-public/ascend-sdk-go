# AlternativeOrderDistributionType

Identifies the distribution type for buy side orders (that is, when `side` value = `BUY`). - Orders will be rejected if the alternative investment asset does not allow the requested distribution type. - Not relevant for sell orders (that is, when `side` value = `SELL`). - Confirm the asset’s `cash_distribution_allowed` and `reinvestment_distribution_allowed` properties to know what it allows.


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `AlternativeOrderDistributionTypeDistributionTypeUnspecified` | DISTRIBUTION_TYPE_UNSPECIFIED                                 |
| `AlternativeOrderDistributionTypeCash`                        | CASH                                                          |
| `AlternativeOrderDistributionTypeReinvestment`                | REINVESTMENT                                                  |