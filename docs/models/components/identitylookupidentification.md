# IdentityLookupIdentification

Identification document for verification


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     | Example                                                                         |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `RegionCode`                                                                    | **string*                                                                       | :heavy_minus_sign:                                                              | CLDR format                                                                     | US                                                                              |
| `Type`                                                                          | [*components.IdentityLookupType](../../models/components/identitylookuptype.md) | :heavy_minus_sign:                                                              | The type of identification document                                             | SSN                                                                             |
| `Value`                                                                         | **string*                                                                       | :heavy_minus_sign:                                                              | The value of the identification document                                        | 123-45-6789                                                                     |