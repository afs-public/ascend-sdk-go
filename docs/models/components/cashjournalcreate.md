# CashJournalCreate

A cash journal transfer. Funds are moved from a source account to a destination account


## Fields

| Field                                                                                                                                                                                                                                                                                                                                                        | Type                                                                                                                                                                                                                                                                                                                                                         | Required                                                                                                                                                                                                                                                                                                                                                     | Description                                                                                                                                                                                                                                                                                                                                                  | Example                                                                                                                                                                                                                                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `Amount`                                                                                                                                                                                                                                                                                                                                                     | [*components.DecimalCreate](../../models/components/decimalcreate.md)                                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                           | A representation of a decimal value, such as 2.5. Clients may convert values into language-native decimal formats, such as Java's [BigDecimal][] or Python's [decimal.Decimal][].<br/><br/> [BigDecimal]:<br/> https://docs.oracle.com/en/java/javase/11/docs/api/java.base/java/math/BigDecimal.html<br/> [decimal.Decimal]: https://docs.python.org/3/library/decimal.html |                                                                                                                                                                                                                                                                                                                                                              |
| `ClientTransferID`                                                                                                                                                                                                                                                                                                                                           | *string*                                                                                                                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                           | The external identifier supplied by the API caller Each request must have a unique pairing of `client_transfer_id` and `source_account`                                                                                                                                                                                                                      | 113bw03-49f8-4525-934c-560fb39dg2kd                                                                                                                                                                                                                                                                                                                          |
| `DestinationAccount`                                                                                                                                                                                                                                                                                                                                         | *string*                                                                                                                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                           | The account that funds will be moved to                                                                                                                                                                                                                                                                                                                      | accounts/01H8FB90ZRRFWXB4XC2JPJ1D4Y                                                                                                                                                                                                                                                                                                                          |
| `FullDisbursement`                                                                                                                                                                                                                                                                                                                                           | **bool*                                                                                                                                                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                           | Whether the entire source account balance is being withdrawn Must not be true if the `amount` is specified                                                                                                                                                                                                                                                   | false                                                                                                                                                                                                                                                                                                                                                        |
| `RetirementContribution`                                                                                                                                                                                                                                                                                                                                     | [*components.RetirementContributionCreate](../../models/components/retirementcontributioncreate.md)                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                           | A contribution to a retirement account.                                                                                                                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                                                                                                                              |
| `RetirementDistribution`                                                                                                                                                                                                                                                                                                                                     | [*components.RetirementDistributionCreate](../../models/components/retirementdistributioncreate.md)                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                           | A distribution from a retirement account.                                                                                                                                                                                                                                                                                                                    |                                                                                                                                                                                                                                                                                                                                                              |
| `SourceAccount`                                                                                                                                                                                                                                                                                                                                              | *string*                                                                                                                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                           | The account that funds will be moved from                                                                                                                                                                                                                                                                                                                    | accounts/01H8FM6EXVH77SAW3TC8KAWMES                                                                                                                                                                                                                                                                                                                          |