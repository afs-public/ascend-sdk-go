# Drip

Object containing metadata for reserving cash until the DRIP trades are executed


## Fields

| Field                                                                                                                                       | Type                                                                                                                                        | Required                                                                                                                                    | Description                                                                                                                                 | Example                                                                                                                                     |
| ------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- |
| `Action`                                                                                                                                    | [*components.EntryAction](../../models/components/entryaction.md)                                                                           | :heavy_minus_sign:                                                                                                                          | Indicates whether the drip memo activity is reserving cash (DRIP_PENDING) or removing the reservation after a successful reinvestment trade | DRIP_PENDING                                                                                                                                |