# BasketOrdersServiceGetBasketResponse


## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `HTTPMeta`                                                                                       | [components.HTTPMetadata](../../models/components/httpmetadata.md)                               | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `Basket`                                                                                         | [*components.Basket](../../models/components/basket.md)                                          | :heavy_minus_sign:                                                                               | OK                                                                                               |
| `Status`                                                                                         | [*components.Status](../../models/components/status.md)                                          | :heavy_minus_sign:                                                                               | INVALID_ARGUMENT: The correspondent_id or the basket_id could not be determined for the request. |