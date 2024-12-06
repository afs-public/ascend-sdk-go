# ActivityAccountTransfer

Object containing metadata for account transfer movements


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             | Example                                                                                 |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `AcatsAssetSequenceNumber`                                                              | **string*                                                                               | :heavy_minus_sign:                                                                      | Sequence number assigned by the DTCC ACATS transfer system for each asset transferred   | 20240424178509                                                                          |
| `AcatsControlNumber`                                                                    | **string*                                                                               | :heavy_minus_sign:                                                                      | Unique Identifier generated by the NSCC ACATS when a transfer is initiated or submitted | 20240360002172                                                                          |
| `Action`                                                                                | [*components.ActivityAction](../../models/components/activityaction.md)                 | :heavy_minus_sign:                                                                      | Denotes whether the shares are incoming or outgoing                                     | INCOMING                                                                                |
| `AdditionalInstructions`                                                                | **string*                                                                               | :heavy_minus_sign:                                                                      | Free form text field containing additional information about a transaction              | Account Transfer instruction                                                            |
| `ContraPartyAccountNumber`                                                              | **string*                                                                               | :heavy_minus_sign:                                                                      | Account number at the contra firm                                                       | DBtvTOGIqBu5Pmz9Y14laM6G5jWTACMvwCV22nLYteo                                             |
| `ContraPartyID`                                                                         | **string*                                                                               | :heavy_minus_sign:                                                                      | Contra party identifier                                                                 | 9999                                                                                    |
| `Institution`                                                                           | **string*                                                                               | :heavy_minus_sign:                                                                      | Contra party institution for the account transfer                                       | Schwab                                                                                  |
| `Method`                                                                                | [*components.ActivityMethod](../../models/components/activitymethod.md)                 | :heavy_minus_sign:                                                                      | The method used for the account transfer                                                | ACATS                                                                                   |