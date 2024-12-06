# TravelRule

The travel rule information for the ICT deposit


## Fields

| Field                                                                                                                        | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  | Example                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `EntityOriginatingParty`                                                                                                     | [*components.EntityOriginatingParty](../../models/components/entityoriginatingparty.md)                                      | :heavy_minus_sign:                                                                                                           | An entity originating party                                                                                                  |                                                                                                                              |
| `EntityRecipientParty`                                                                                                       | [*components.EntityRecipientParty](../../models/components/entityrecipientparty.md)                                          | :heavy_minus_sign:                                                                                                           | An entity recipient party                                                                                                    |                                                                                                                              |
| `IndividualOriginatingParty`                                                                                                 | [*components.IndividualOriginatingParty](../../models/components/individualoriginatingparty.md)                              | :heavy_minus_sign:                                                                                                           | An individual originating party                                                                                              |                                                                                                                              |
| `IndividualRecipientParty`                                                                                                   | [*components.IndividualRecipientParty](../../models/components/individualrecipientparty.md)                                  | :heavy_minus_sign:                                                                                                           | An individual recipient party                                                                                                |                                                                                                                              |
| `OriginatingInstitution`                                                                                                     | [*components.OriginatingInstitution](../../models/components/originatinginstitution.md)                                      | :heavy_minus_sign:                                                                                                           | The name of the external financial institution and account that is the source of the funds                                   | {<br/>"account_id": "0987654321",<br/>"title": "Bank of New York"<br/>}                                                      |
| `RecipientInstitution`                                                                                                       | [*components.RecipientInstitution](../../models/components/recipientinstitution.md)                                          | :heavy_minus_sign:                                                                                                           | The name and account id of institution receiving the funds. Always 'Apex Clearing' and investor account id for ICT deposits; | {<br/>"account_id": "01H8FB90ZRRFWXB4XC2JPJ1D4Y",<br/>"title": "Apex Clearing"<br/>}                                         |