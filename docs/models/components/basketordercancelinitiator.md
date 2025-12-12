# BasketOrderCancelInitiator

Output only field that is required for Equity Orders for any client who is having Apex do CAT reporting on their behalf. This field denotes the initiator of the cancel request. This field will be present when provided on the RemoveOrderRequest


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `BasketOrderCancelInitiatorInitiatorUnspecified` | INITIATOR_UNSPECIFIED                            |
| `BasketOrderCancelInitiatorFirm`                 | FIRM                                             |
| `BasketOrderCancelInitiatorClient`               | CLIENT                                           |