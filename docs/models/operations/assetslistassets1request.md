# AssetsListAssets1Request


## Fields

| Field                                                                                                                                                                                                                                                             | Type                                                                                                                                                                                                                                                              | Required                                                                                                                                                                                                                                                          | Description                                                                                                                                                                                                                                                       | Example                                                                                                                                                                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Parent`                                                                                                                                                                                                                                                          | **string*                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                | The parent resource name, which is the correspondent ID.                                                                                                                                                                                                          | correspondents/1234                                                                                                                                                                                                                                               |
| `PageSize`                                                                                                                                                                                                                                                        | **int*                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                | The maximum number of assets to return. The service may return fewer than this value. Default is 100 (subject to change) The maximum is 1000, values exceeding this will be set to 1000 (subject to change)                                                       | 100                                                                                                                                                                                                                                                               |
| `PageToken`                                                                                                                                                                                                                                                       | **string*                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                | A page token, received from a previous `ListAssets` call. Provide this to retrieve the subsequent page. When paginating, all other parameters provided to `ListAssets` must match the call that provided the page token in order to maintain a stable result set. | Mv-BAwEBCVBhZ2VUb2tlbgH_ggABAgEPUmVxdWVzdENoZWNrc3VtAQYAAQJJZAEMAAAAD_-CAfzrRtzkAQQ1MDA3AA==                                                                                                                                                                      |
| `Filter`                                                                                                                                                                                                                                                          | **string*                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                | A CEL string to filter results; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information;                                                                                              | (symbol == 'IBM' && usable) \|\| symbol == 'USD'                                                                                                                                                                                                                  |