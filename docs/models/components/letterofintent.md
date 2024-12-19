# LetterOfIntent

Letter of Intent (LOI). An LOI allows investors to receive sales charge discounts based on a commitment to buy a specified monetary amount of shares over a period of time, usually 13 months.


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `Amount`                                                                           | [*components.OrderAmount](../../models/components/orderamount.md)                  | :heavy_minus_sign:                                                                 | The amount of the LOI. This is a monetary value in the same currency as the order. | {<br/>"value": "30.57"<br/>}                                                       |
| `PeriodStartDate`                                                                  | [*components.PeriodStartDate](../../models/components/periodstartdate.md)          | :heavy_minus_sign:                                                                 | The period start date of the LOI.                                                  | {<br/>"day": 25,<br/>"month": 4,<br/>"year": 2024<br/>}                            |