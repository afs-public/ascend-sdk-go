# BrokerCapacity

Defaults to "AGENCY" if not specified, except for Fixed Income orders from RIA correspondents which default to "PRINCIPAL" when not specified. For Equities: Only "AGENCY" is allowed. For Mutual Funds: Only "AGENCY" is allowed. For Fixed Income: Either "AGENCY" or "PRINCIPAL" are allowed.  - RIA correspondents: Defaults to "PRINCIPAL" if not specified.  - Other correspondents: Defaults to "AGENCY" if not specified. For Event Contracts: Only "AGENCY" is allowed.


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `BrokerCapacityBrokerCapacityUnspecified` | BROKER_CAPACITY_UNSPECIFIED               |
| `BrokerCapacityAgency`                    | AGENCY                                    |
| `BrokerCapacityPrincipal`                 | PRINCIPAL                                 |