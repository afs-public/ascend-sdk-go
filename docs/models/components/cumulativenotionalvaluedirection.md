# CumulativeNotionalValueDirection

This represents the direction of the total filled notional value, which will be present if the `cumulative_notional_value` is also present. If there are no executions, this value will be absent. When the summed notional value of all CREDIT legs exceeds that of the DEBIT legs a CREDIT will be reported here; otherwise a DEBIT will be reported.


## Values

| Name                                                         | Value                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `CumulativeNotionalValueDirectionDebitCreditTypeUnspecified` | DEBIT_CREDIT_TYPE_UNSPECIFIED                                |
| `CumulativeNotionalValueDirectionDebit`                      | DEBIT                                                        |
| `CumulativeNotionalValueDirectionCredit`                     | CREDIT                                                       |