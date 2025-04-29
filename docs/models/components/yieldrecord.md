# YieldRecord

Contains details about the yields associated with a trade in fixed income instruments


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         | Example                                                                             |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `YieldPercent`                                                                      | [*components.YieldPercent](../../models/components/yieldpercent.md)                 | :heavy_minus_sign:                                                                  | The yield percentage at which the transaction was effected                          | {<br/>"value": "0.25"<br/>}                                                         |
| `YieldType`                                                                         | [*components.YieldRecordYieldType](../../models/components/yieldrecordyieldtype.md) | :heavy_minus_sign:                                                                  | Characterization of a yield percentage                                              | YIELD_TO_CALL                                                                       |