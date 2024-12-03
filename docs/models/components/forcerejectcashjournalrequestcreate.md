# ForceRejectCashJournalRequestCreate

Request to force reject a pending cash journal


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            | Example                                                |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Name`                                                 | *string*                                               | :heavy_check_mark:                                     | The name of the cash journal to force reject           | cashJournals/20230817000319                            |
| `Reason`                                               | **string*                                              | :heavy_minus_sign:                                     | The optional reason for force rejecting a cash journal | Forced rejection test                                  |