# IctDepositTravelRuleCreate

The travel rules associated with an ICT deposit


## Fields

| Field                                                                                             | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `EntityOriginatingParty`                                                                          | [*components.TravelRuleEntityPartyCreate](../../models/components/travelruleentitypartycreate.md) | :heavy_minus_sign:                                                                                | Travel rule entity party                                                                          |
| `EntityRecipientParty`                                                                            | [*components.TravelRuleEntityPartyCreate](../../models/components/travelruleentitypartycreate.md) | :heavy_minus_sign:                                                                                | Travel rule entity party                                                                          |
| `IndividualOriginatingParty`                                                                      | [*components.TravelRulePartyCreate](../../models/components/travelrulepartycreate.md)             | :heavy_minus_sign:                                                                                | Travel rule party                                                                                 |
| `IndividualRecipientParty`                                                                        | [*components.TravelRulePartyCreate](../../models/components/travelrulepartycreate.md)             | :heavy_minus_sign:                                                                                | Travel rule party                                                                                 |
| `OriginatingInstitution`                                                                          | [components.InstitutionCreate](../../models/components/institutioncreate.md)                      | :heavy_check_mark:                                                                                | Institution representing originator or recipient of funds from an Instant Cash Transfer           |