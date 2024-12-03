# AuditTrail

Audit trail details


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   | Example                                                       |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `AuditType`                                                   | [*components.AuditType](../../models/components/audittype.md) | :heavy_minus_sign:                                            | The type of audit that was performed on the investigation     | INVESTIGATION_REQUEST_UPDATE                                  |
| `Comment`                                                     | **string*                                                     | :heavy_minus_sign:                                            | Comment relating to why the audit was saved                   | Updating family name                                          |
| `Field`                                                       | **string*                                                     | :heavy_minus_sign:                                            | The name of the field that has been updated                   | investigation_request.person_investigation.family_name        |
| `NewValue`                                                    | **string*                                                     | :heavy_minus_sign:                                            | The new value for the field that was updated                  | Doe                                                           |
| `OldValue`                                                    | **string*                                                     | :heavy_minus_sign:                                            | The prior value for the field that was updated                | Dough                                                         |
| `UpdateTime`                                                  | [*time.Time](https://pkg.go.dev/time#Time)                    | :heavy_minus_sign:                                            | The date the user updated the investigation                   | 2023-06-13 23:48:58.343 +0000 UTC                             |
| `UpdateUser`                                                  | **string*                                                     | :heavy_minus_sign:                                            | The user that made the update to the investigation            | jsmith                                                        |