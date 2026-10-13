# AcatsAssetFailure

AcatsAssetFailure is one reason a specific holding can't transfer.


## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               | Example                                                                                   |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `Message`                                                                                 | **string*                                                                                 | :heavy_minus_sign:                                                                        | Human-readable explanation.                                                               | Asset could not be identified in the Apex catalog.                                        |
| `Reason`                                                                                  | [*components.AcatsAssetFailureReason](../../models/components/acatsassetfailurereason.md) | :heavy_minus_sign:                                                                        | Machine-readable failure reason.                                                          | ASSET_NOT_IDENTIFIED                                                                      |