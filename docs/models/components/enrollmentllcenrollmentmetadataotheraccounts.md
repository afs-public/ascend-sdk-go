# EnrollmentLlcEnrollmentMetadataOtherAccounts

A customer-disclosed list of other Apex-held accounts owned by the Entity applicant at the time of this account's application; expressed as zero, one, or many account numbers


## Fields

| Field                                | Type                                 | Required                             | Description                          | Example                              |
| ------------------------------------ | ------------------------------------ | ------------------------------------ | ------------------------------------ | ------------------------------------ |
| `AccountNames`                       | []*string*                           | :heavy_minus_sign:                   | Other account names held at Apex     | [<br/>"John Doe Trading"<br/>]       |
| `AccountNumbers`                     | []*string*                           | :heavy_minus_sign:                   | Other account numbers held at Apex   | [<br/>"N6D8ZJP"<br/>]                |
| `OwnerHasOtherAccountsAtApex`        | **bool*                              | :heavy_minus_sign:                   | The owner has other accounts at Apex | true                                 |