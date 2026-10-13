# IneligibleHolding

IneligibleHolding is a holding that can't transfer, with the reason(s) why.


## Fields

| Field                                                                                                 | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `Failures`                                                                                            | [][components.AcatsAssetFailure](../../models/components/acatsassetfailure.md)                        | :heavy_minus_sign:                                                                                    | Why the holding is ineligible; at least one entry.                                                    |
| `Holding`                                                                                             | [*components.AcatsIneligibleHoldingHolding](../../models/components/acatsineligibleholdingholding.md) | :heavy_minus_sign:                                                                                    | The holding.                                                                                          |