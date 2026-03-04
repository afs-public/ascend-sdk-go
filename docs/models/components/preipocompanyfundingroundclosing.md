# PreIpoCompanyFundingRoundClosing

Closing details of a Pre IPO Company Funding Round.


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     | Example                                                         |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `IssuePrice`                                                    | [*components.IssuePrice](../../models/components/issueprice.md) | :heavy_minus_sign:                                              | The issue price of the closing.                                 | {<br/>"value": "695.44"<br/>}                                   |
| `IssueTime`                                                     | [*time.Time](https://pkg.go.dev/time#Time)                      | :heavy_minus_sign:                                              | The time shares were issued.                                    | 2025-09-30T04:00:00Z                                            |
| `Name`                                                          | **string*                                                       | :heavy_minus_sign:                                              | The name of the closing.                                        | Series E-5                                                      |
| `UpdateTime`                                                    | [*time.Time](https://pkg.go.dev/time#Time)                      | :heavy_minus_sign:                                              | The time the closing was updated.                               | 2025-09-28T10:00:46.611232Z                                     |