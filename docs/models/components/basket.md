# Basket

The message describing a basket


## Fields

| Field                                                                                                                | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          | Example                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `BasketID`                                                                                                           | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | System generated unique id for the basket                                                                            | fffd326-72fa-4d2b-bd1f-45384fe5d521                                                                                  |
| `BasketOrderCount`                                                                                                   | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | Number of orders in the basket                                                                                       | 30000                                                                                                                |
| `BasketState`                                                                                                        | [*components.BasketState](../../models/components/basketstate.md)                                                    | :heavy_minus_sign:                                                                                                   | The processing status of the basket                                                                                  | SUBMITTED                                                                                                            |
| `ClientBasketID`                                                                                                     | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | User-supplied unique basket ID. Cannot be more than 40 characters long.                                              | 39347a0d-860b-48e8-a04d-511133f057e3                                                                                 |
| `CompleteTime`                                                                                                       | [*time.Time](https://pkg.go.dev/time#Time)                                                                           | :heavy_minus_sign:                                                                                                   | Time the basket was completed                                                                                        | 2023-09-21 16:55:27.58 +0000 UTC                                                                                     |
| `CompressedOrderCount`                                                                                               | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | Number of compressed orders in the basket that will go to market. This number is calculated after basket submission. | 35                                                                                                                   |
| `CorrespondentID`                                                                                                    | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | The unique id for the correspondent for the basket                                                                   | 01HPMZZM6RKMVZA1JQ63RQKJRP                                                                                           |
| `CreateTime`                                                                                                         | [*time.Time](https://pkg.go.dev/time#Time)                                                                           | :heavy_minus_sign:                                                                                                   | Time of the basket creation                                                                                          | 2023-09-21 16:55:27.58 +0000 UTC                                                                                     |
| `DistinctAccountCount`                                                                                               | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | Number of distinct accounts in the basket.                                                                           | 500                                                                                                                  |
| `LastUpdateTime`                                                                                                     | [*time.Time](https://pkg.go.dev/time#Time)                                                                           | :heavy_minus_sign:                                                                                                   | Time of the last basket update                                                                                       | 2023-09-21 16:55:27.58 +0000 UTC                                                                                     |
| `Name`                                                                                                               | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | System generated name of the basket                                                                                  | correspondents/01HPMZZM6RKMVZA1JQ63RQKJRP/baskets/fffd326-72fa-4d2b-bd1f-45384fe5d521                                |
| `RejectedAccountCount`                                                                                               | **string*                                                                                                            | :heavy_minus_sign:                                                                                                   | Number of accounts that did not pass risk checks and their orders were rejected.                                     | 2                                                                                                                    |
| `SubmitTime`                                                                                                         | [*time.Time](https://pkg.go.dev/time#Time)                                                                           | :heavy_minus_sign:                                                                                                   | Time the basket was submitted                                                                                        | 2023-09-21 16:55:27.58 +0000 UTC                                                                                     |