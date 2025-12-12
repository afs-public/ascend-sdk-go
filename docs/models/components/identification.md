# Identification

Represents an identification document


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     | Example                                                                         |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `RegionCode`                                                                    | **string*                                                                       | :heavy_minus_sign:                                                              | CLDR format                                                                     | US                                                                              |
| `Type`                                                                          | [*components.IdentificationType](../../models/components/identificationtype.md) | :heavy_minus_sign:                                                              | The type of identification document                                             | SSN                                                                             |
| `Value`                                                                         | **string*                                                                       | :heavy_minus_sign:                                                              | The value of the identification document                                        | 123-45-6789                                                                     |