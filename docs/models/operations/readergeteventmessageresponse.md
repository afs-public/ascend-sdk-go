# ReaderGetEventMessageResponse


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `HTTPMeta`                                                          | [components.HTTPMetadata](../../models/components/httpmetadata.md)  | :heavy_check_mark:                                                  | N/A                                                                 |
| `EventMessage`                                                      | [*components.EventMessage](../../models/components/eventmessage.md) | :heavy_minus_sign:                                                  | OK                                                                  |
| `Status`                                                            | [*components.Status](../../models/components/status.md)             | :heavy_minus_sign:                                                  | INVALID_ARGUMENT: The request was not well formed.                  |