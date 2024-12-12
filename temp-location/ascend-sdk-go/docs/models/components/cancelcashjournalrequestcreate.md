# CancelCashJournalRequestCreate

Request to cancel an existing cash journal


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       | Example                                                                           |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `Name`                                                                            | *string*                                                                          | :heavy_check_mark:                                                                | The name of the cash journal to cancel                                            | cashJournals/20240717000319                                                       |
| `Reason`                                                                          | **string*                                                                         | :heavy_minus_sign:                                                                | The reason for canceling the cash journal Maximum of 100 characters are supported | Cancel journal                                                                    |