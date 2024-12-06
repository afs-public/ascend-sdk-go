# PushSubscriptionCreate

Configuration information about a push subscription


## Fields

| Field                                                                                                                                                                                                         | Type                                                                                                                                                                                                          | Required                                                                                                                                                                                                      | Description                                                                                                                                                                                                   | Example                                                                                                                                                                                                       |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ClientID`                                                                                                                                                                                                    | **string*                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                            | The client that owns the subscription. A client subscription will receive events for it and all of its correspondents. This can only be set at creation time and is mutually exclusive with correspondent_id. |                                                                                                                                                                                                               |
| `CorrespondentID`                                                                                                                                                                                             | **string*                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                            | The correspondent that owns the subscription. A correspondent subscription will receive events only for itself. This can only be set at creation time and is mutually exclusive with client_id.               | 01H8MCDXH4HYJJAV921BDKCC83                                                                                                                                                                                    |
| `DisplayName`                                                                                                                                                                                                 | *string*                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                            | The user-defined name for the subscription                                                                                                                                                                    | This is an example HTTP configuration.                                                                                                                                                                        |
| `EventTypes`                                                                                                                                                                                                  | []*string*                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                            | Filter for event types; ["\*"] matches all values; Suffix wildcards using "\*" (e.g. ["account.\*"]) are supported                                                                                            | [<br/>"position.v1.updated"<br/>]                                                                                                                                                                             |
| `HTTPCallback`                                                                                                                                                                                                | [*components.HTTPPushCallbackCreate](../../models/components/httppushcallbackcreate.md)                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                            | Configuration information about an HTTP target callback                                                                                                                                                       |                                                                                                                                                                                                               |