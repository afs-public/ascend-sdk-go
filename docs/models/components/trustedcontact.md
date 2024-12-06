# TrustedContact

A Trusted Contact is a person designated to verify the well being of the account holder. Only one form of contact information is required; An account may contain zero, one, or many.


## Fields

| Field                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | Example                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `EmailAddress`                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                               | An email address indicated for account communications.                                                                                                                                                                                                                                                                                                                                                                                                                           | example@email.com                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| `FamilyName`                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                               | Family name of a natural person.                                                                                                                                                                                                                                                                                                                                                                                                                                                 | Doe                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| `GivenName`                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                               | The given name of a natural person; Conventionally known as 'first name' in most English-speaking countries.                                                                                                                                                                                                                                                                                                                                                                     | John                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| `MailingAddress`                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [*components.TrustedContactMailingAddress](../../models/components/trustedcontactmailingaddress.md)                                                                                                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                               | The object containing data for the purpose of delivery physical mailings to a party; Typically used for statements, account updates, tax documents, and other postal mailings; May also be used as an alternative identity verification address to personalAddress. If input, the required fields within the `mailing_address` object include:<br/> - `administrative_area`<br/> - `region_code` - 2 character CLDR Code<br/> - `postal_code`<br/> - `locality`<br/> - `address_lines` - max 5 lines |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| `MiddleNames`                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                               | Non-primary names representing a natural person; Name attributed to a person other than "Given" and "Family" names.                                                                                                                                                                                                                                                                                                                                                              | Larry                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| `Name`                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                               | The name field Format: {parent=account/*}/{name=trustedContacts/*}                                                                                                                                                                                                                                                                                                                                                                                                               | accounts/01HC3MAQ4DR9QN1V8MJ4CN1HMK/trustedContacts/8096110d-fb55-4f9d-b883-b84f0b70d3ea                                                                                                                                                                                                                                                                                                                                                                                         |
| `PhoneNumber`                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | [*components.TrustedContactPhoneNumber](../../models/components/trustedcontactphonenumber.md)                                                                                                                                                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                               | The phone number for a party; this value exists on the party record in the context of the account and does not commute to other accounts held by/for the person                                                                                                                                                                                                                                                                                                                  | 555-123-4567                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| `TrustedContactID`                                                                                                                                                                                                                                                                                                                                                                                                                                                               | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                               | A system-generated unique identifier referencing a Trusted Contact person on an account; Used to access the record after creation                                                                                                                                                                                                                                                                                                                                                | 8096110d-fb55-4f9d-b883-b84f0b70d3ea                                                                                                                                                                                                                                                                                                                                                                                                                                             |