# File

The file containing the snapshot data.


## Fields

| Field                                            | Type                                             | Required                                         | Description                                      |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `DownloadURI`                                    | **string*                                        | :heavy_minus_sign:                               | The signed download uri for the file.            |
| `URIExpiryTime`                                  | [*time.Time](https://pkg.go.dev/time#Time)       | :heavy_minus_sign:                               | The timestamp at which the download uri expires. |