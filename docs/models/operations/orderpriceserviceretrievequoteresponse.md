# OrderPriceServiceRetrieveQuoteResponse


## Fields

| Field                                                                                                                                                             | Type                                                                                                                                                              | Required                                                                                                                                                          | Description                                                                                                                                                       |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `HTTPMeta`                                                                                                                                                        | [components.HTTPMetadata](../../models/components/httpmetadata.md)                                                                                                | :heavy_check_mark:                                                                                                                                                | N/A                                                                                                                                                               |
| `RetrieveQuoteResponse`                                                                                                                                           | [*components.RetrieveQuoteResponse](../../models/components/retrievequoteresponse.md)                                                                             | :heavy_minus_sign:                                                                                                                                                | OK                                                                                                                                                                |
| `Status`                                                                                                                                                          | [*components.Status](../../models/components/status.md)                                                                                                           | :heavy_minus_sign:                                                                                                                                                | INVALID_ARGUMENT: There was an issue with one or more fields in the request.  The message field will contain details about which field failed validation and why. |