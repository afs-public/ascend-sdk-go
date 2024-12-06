# SubscriberListPushSubscriptionsRequest


## Fields

| Field                                                                                                                                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                                                                                                                                                                 | Required                                                                                                                                                                                                                                                                                                                                                                                                                             | Description                                                                                                                                                                                                                                                                                                                                                                                                                          | Example                                                                                                                                                                                                                                                                                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `Filter`                                                                                                                                                                                                                                                                                                                                                                                                                             | **string*                                                                                                                                                                                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                   | A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information; If empty, all subscriptions the user has permission to view will be returned; Filter options include:<br/> `name`<br/> `subscription_id`<br/> `client_id`<br/> `correspondent_id`<br/> `display_name`<br/> `event_types`<br/> `state`<br/> `http_callback.url`<br/> `http_callback.timeout_seconds` | correspondent_id=="01H8MCDXH4HYJJAV921BDKCC83"                                                                                                                                                                                                                                                                                                                                                                                       |
| `PageSize`                                                                                                                                                                                                                                                                                                                                                                                                                           | **int*                                                                                                                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                   | The number of entries to return in a single page; Default = 100; Maximum = 1000                                                                                                                                                                                                                                                                                                                                                      | 50                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| `PageToken`                                                                                                                                                                                                                                                                                                                                                                                                                          | **string*                                                                                                                                                                                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                   | Page token used for pagination; Supplying a page token returns the next page of results                                                                                                                                                                                                                                                                                                                                              | ZXhhbXBsZQo                                                                                                                                                                                                                                                                                                                                                                                                                          |