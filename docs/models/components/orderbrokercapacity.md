# OrderBrokerCapacity

Defaults to "AGENCY" for all asset types except Fixed Income. For Equities: Only "AGENCY" is allowed. For Mutual Funds: Only "AGENCY" is allowed. For Fixed Income: Either "AGENCY" or "PRINCIPAL" are allowed.  If not specified, defaults to the capacity configured for  the correspondent by Apex. For Event Contracts: Only "AGENCY" is allowed.


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `OrderBrokerCapacityBrokerCapacityUnspecified` | BROKER_CAPACITY_UNSPECIFIED                    |
| `OrderBrokerCapacityAgency`                    | AGENCY                                         |
| `OrderBrokerCapacityPrincipal`                 | PRINCIPAL                                      |