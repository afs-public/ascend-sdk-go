# ListSnapshotsResponse

Returns the requested snapshots.


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `NextPageToken`                                              | **string*                                                    | :heavy_minus_sign:                                           | The token for retrieving the next page of snapshots.         | ZXhhbXBsZQo                                                  |
| `Snapshots`                                                  | [][components.Snapshot](../../models/components/snapshot.md) | :heavy_minus_sign:                                           | The returned snapshots.                                      |                                                              |
| `TotalSize`                                                  | **string*                                                    | :heavy_minus_sign:                                           | The total number of snapshots to return.                     | 100                                                          |