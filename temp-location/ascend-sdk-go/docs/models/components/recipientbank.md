# RecipientBank

The recipient bank / financial institution


## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `BankID`                                                                                    | [*components.BankID](../../models/components/bankid.md)                                     | :heavy_minus_sign:                                                                          | An identifier that represents ABA routing number for domestic wire or BIC for foreign wire  |
| `InternationalBankDetails`                                                                  | [*components.InternationalBankDetails](../../models/components/internationalbankdetails.md) | :heavy_minus_sign:                                                                          | Bank details required in the case of an international wire transfer                         |