# ScheduledRetirementContributionCreate

The retirement contribution details for a scheduled deposit


## Fields

| Field                                                                                                                                             | Type                                                                                                                                              | Required                                                                                                                                          | Description                                                                                                                                       | Example                                                                                                                                           |
| ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- |
| `TaxYear`                                                                                                                                         | **int*                                                                                                                                            | :heavy_minus_sign:                                                                                                                                | An explicit tax year value. The current year is always valid; and the prior year is valid only before the tax deadline. Must be in "YYYY" format. | 2024                                                                                                                                              |
| `TemporalTaxYear`                                                                                                                                 | [*components.TemporalTaxYear](../../models/components/temporaltaxyear.md)                                                                         | :heavy_minus_sign:                                                                                                                                | A temporal tax year value. This will always evaluate to a year based on the date the transfer was initiated.                                      | CURRENT_CALENDAR_YEAR                                                                                                                             |
| `Type`                                                                                                                                            | [components.ScheduledRetirementContributionCreateType](../../models/components/scheduledretirementcontributioncreatetype.md)                      | :heavy_check_mark:                                                                                                                                | The type of retirement contribution.                                                                                                              | REGULAR                                                                                                                                           |