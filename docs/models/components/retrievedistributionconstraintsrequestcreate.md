# RetrieveDistributionConstraintsRequestCreate

Request to retrieve retirement distribution constraints


## Fields

| Field                                                                                                                                                | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          | Example                                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Mechanism`                                                                                                                                          | [components.RetrieveDistributionConstraintsRequestCreateMechanism](../../models/components/retrievedistributionconstraintsrequestcreatemechanism.md) | :heavy_check_mark:                                                                                                                                   | Cash transfer mechanism to search constraints for                                                                                                    | ACH                                                                                                                                                  |
| `Name`                                                                                                                                               | *string*                                                                                                                                             | :heavy_check_mark:                                                                                                                                   | Name of the account being queried, for retirement distribution constraints Format: accounts/{account}                                                | accounts/01H8FM6EXVH77SAW3TC8KAWMES                                                                                                                  |