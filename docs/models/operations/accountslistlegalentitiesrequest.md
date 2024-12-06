# AccountsListLegalEntitiesRequest


## Fields

| Field                                                                                                                                                                                                                                                                            | Type                                                                                                                                                                                                                                                                             | Required                                                                                                                                                                                                                                                                         | Description                                                                                                                                                                                                                                                                      | Example                                                                                                                                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `PageSize`                                                                                                                                                                                                                                                                       | **int*                                                                                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The maximum number of legal entities to return. The service may return fewer than this value. If unspecified, at most 25 legal entities will be returned. The maximum value is 100; values above 100 will be coerced to 100.                                                     | 25                                                                                                                                                                                                                                                                               |
| `PageToken`                                                                                                                                                                                                                                                                      | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | A page token, received from a previous `ListLegalEntities` call. Provide this to retrieve the subsequent page.<br/><br/> When paginating, all other parameters provided to `ListLegalEntities` must match the call that provided the page token.                                 | AbTYnwAkMjIyZDNjYTAtZmVjZS00N2Q5LTgyMDctNzI3MDdkMjFiZj3h                                                                                                                                                                                                                         |
| `OrderBy`                                                                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | The order in which legal entities are listed.                                                                                                                                                                                                                                    |                                                                                                                                                                                                                                                                                  |
| `Filter`                                                                                                                                                                                                                                                                         | **string*                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                               | A CEL string to filter results; Use `upperAscii()` for case-insensitive searches; E.g. `entity_name.upperAscii()=="AcMe,InC".upperAscii()`; See the [CEL Search](https://developer.apexclearing.com/apex-fintech-solutions/docs/cel-search) page in Guides for more information; |                                                                                                                                                                                                                                                                                  |