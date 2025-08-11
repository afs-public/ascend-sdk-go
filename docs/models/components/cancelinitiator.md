# CancelInitiator

Output only field that is required for Equity Orders for any client who is having Apex do CAT reporting on their behalf. This field denotes the initiator of the cancel request. This field will be present when provided on the CancelOrderRequest


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `CancelInitiatorInitiatorUnspecified` | INITIATOR_UNSPECIFIED                 |
| `CancelInitiatorFirm`                 | FIRM                                  |
| `CancelInitiatorClient`               | CLIENT                                |