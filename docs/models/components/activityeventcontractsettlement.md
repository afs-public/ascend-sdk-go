# ActivityEventContractSettlement

Used to record the settlement/payout of event contracts based on real-world event outcomes


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               | Example                                                                   |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `Exchange`                                                                | **string*                                                                 | :heavy_minus_sign:                                                        | The exchange that issued the event contract                               | KALSHI                                                                    |
| `Outcome`                                                                 | [*components.ActivityOutcome](../../models/components/activityoutcome.md) | :heavy_minus_sign:                                                        | The determined outcome of the event                                       | FAVORABLE                                                                 |