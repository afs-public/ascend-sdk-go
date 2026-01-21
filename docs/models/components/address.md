# Address

**Field Dependencies:**

Address is required for third party beneficiaries

Required if `third_party` is `true`.


## Fields

| Field                                                                                                  | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `City`                                                                                                 | **string*                                                                                              | :heavy_minus_sign:                                                                                     | Required: Describes the city in which the entity is located.                                           |
| `Country`                                                                                              | **string*                                                                                              | :heavy_minus_sign:                                                                                     | Required: The country code used for geolocation, identity verification, and/or mail delivery purposes. |
| `PostalCode`                                                                                           | **string*                                                                                              | :heavy_minus_sign:                                                                                     | Required: The postal code used for geolocation, identity verification, and/or mail delivery purposes.  |
| `State`                                                                                                | **string*                                                                                              | :heavy_minus_sign:                                                                                     | Required: The state code used for geolocation, identity verification, and/or mail delivery purposes.   |
| `StreetAddress`                                                                                        | []*string*                                                                                             | :heavy_minus_sign:                                                                                     | The street name and number relating to a party's legal or mailing address.                             |