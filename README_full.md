# Introducing the Apex Go SDK

In today's fast-paced digital ecosystem, developers need tools that not only streamline the development process but also unlock new possibilities for innovation and efficiency.

Enter the Apex Go SDK, a cutting-edge software development kit designed to empower fintech app developers like never before.
With our SDK, you can more easily integrate new account creation, optimize trading, and elevate your applications with realtime buying power, all through a seamless SDK interface.

Whether you're building complex, data-driven platforms or simplified, user-centric applications, Apex Go SDK was created to offer the flexibility, power, and ease of use to bring your visions to life faster and more effectively.
Join us in redefining the boundaries of what your applications can achieve.
Start your journey with Apex today.

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/afs-public/ascend-sdk-go
```
<!-- End SDK Installation [installation] -->

## Initializing the SDK

The following sample shows how to initialise the SDK, using the API Key and Service Account Credentials you received during sign-up:
```go
package main

import (
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"context"
	"fmt"
	"os"
)

func main() {
    ctx := context.Background()
    s := ascendsdkgo.New(
        ascendsdkgo.WithServerURL("https://uat.apexapis.com"),
        ascendsdkgo.WithSecurity(components.Security{
            APIKey: ascendsdkgo.String(os.Getenv("API_KEY")),
            ServiceAccountCreds: &components.ServiceAccountCreds{
                PrivateKey:   os.Getenv("SERVICE_ACCOUNT_CREDS_PRIVATE_KEY"),
                Name:         os.Getenv("SERVICE_ACCOUNT_CREDS_NAME"),
                Organization: os.Getenv("SERVICE_ACCOUNT_CREDS_ORGANIZATION"),
                Type:         "serviceAccount",
            },
        }),
    )

    res, err := s.AccountCreation.GetAccount(ctx, "VALID_ACCOUNT_ID", nil)
    if err != nil {
        fmt.Printf("Error was: %s\n", err)
    } else {
        fmt.Printf("Account ID: %s\n", res.Account.AccountID)
    }
}
```

<!-- No SDK Example Usage [usage] -->


<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/models/operations"
	"github.com/afs-public/ascend-sdk-go/retry"
	"log"
	"models/operations"
)

func main() {
	s := ascendsdkgo.New()

	ctx := context.Background()
	res, err := s.Authentication.GenerateServiceAccountToken(ctx, components.GenerateServiceAccountTokenRequestCreate{
		Jws: "eyJhbGciOiAiUlMyNTYifQ.eyJuYW1lIjogImpkb3VnaCIsIm9yZ2FuaXphdGlvbiI6ICJjb3JyZXNwb25kZW50cy8xMjM0NTY3OC0xMjM0LTEyMzQtMTIzNC0xMjM0NTY3ODkxMDEiLCJkYXRldGltZSI6ICIyMDI0LTAyLTA1VDIxOjAyOjI3LjkwMTE4MFoifQ.IMy3KmYoG8Ppf+7hXN7tm7J4MrNpQLGL7WCWvhh4nZWAVKkluL3/u3KC6hZ6Mb/5p7Y54CgZ68aWT2BcP5y4VtzIZR1Chm5pxbLfgE4aJuk+FnF6K3Gc3bBjOWCL58pxY2aTb0iU/exDEA1cbMDvbCzmY5kRefDvorLOqgUS/tS2MJ2jv4RlZFPlmHv5PtOruJ8xUW19gEgGhsPXYYeSHFTE1ZlaDvyXrKtpOvlf+FVc2RTuEw529LZnzwH4/eJJR3BpSpHyJTjQqiaMT3wzpXXYKfCRqnDkSSKJDzCzTb0/uWK/Lf0uafxPXk5YLdis+dbo1zNQhVVKjwnMpk1vLw",
	}, operations.AuthenticationGenerateServiceAccountTokenSecurity{
		APIKeyAuth: "<YOUR_API_KEY_HERE>",
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Token != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/models/operations"
	"github.com/afs-public/ascend-sdk-go/retry"
	"log"
)

func main() {
	s := ascendsdkgo.New(
		ascendsdkgo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
	)

	ctx := context.Background()
	res, err := s.Authentication.GenerateServiceAccountToken(ctx, components.GenerateServiceAccountTokenRequestCreate{
		Jws: "eyJhbGciOiAiUlMyNTYifQ.eyJuYW1lIjogImpkb3VnaCIsIm9yZ2FuaXphdGlvbiI6ICJjb3JyZXNwb25kZW50cy8xMjM0NTY3OC0xMjM0LTEyMzQtMTIzNC0xMjM0NTY3ODkxMDEiLCJkYXRldGltZSI6ICIyMDI0LTAyLTA1VDIxOjAyOjI3LjkwMTE4MFoifQ.IMy3KmYoG8Ppf+7hXN7tm7J4MrNpQLGL7WCWvhh4nZWAVKkluL3/u3KC6hZ6Mb/5p7Y54CgZ68aWT2BcP5y4VtzIZR1Chm5pxbLfgE4aJuk+FnF6K3Gc3bBjOWCL58pxY2aTb0iU/exDEA1cbMDvbCzmY5kRefDvorLOqgUS/tS2MJ2jv4RlZFPlmHv5PtOruJ8xUW19gEgGhsPXYYeSHFTE1ZlaDvyXrKtpOvlf+FVc2RTuEw529LZnzwH4/eJJR3BpSpHyJTjQqiaMT3wzpXXYKfCRqnDkSSKJDzCzTb0/uWK/Lf0uafxPXk5YLdis+dbo1zNQhVVKjwnMpk1vLw",
	}, operations.AuthenticationGenerateServiceAccountTokenSecurity{
		APIKeyAuth: "<YOUR_API_KEY_HERE>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Token != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `GetAccount` function may return the following errors:

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| sdkerrors.Status        | 400, 403, 404, 500, 503 | application/json        |
| sdkerrors.SDKError      | 4XX, 5XX                | \*/\*                   |

### Example

```go
package main

import (
	"context"
	"errors"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/models/operations"
	"github.com/afs-public/ascend-sdk-go/models/sdkerrors"
	"log"
)

func main() {
	s := ascendsdkgo.New(
		ascendsdkgo.WithSecurity(components.Security{
			APIKey: ascendsdkgo.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
			ServiceAccountCreds: &components.ServiceAccountCreds{
				PrivateKey:   "-----BEGIN PRIVATE KEY--{OMITTED FOR BREVITY}",
				Name:         "FinFirm",
				Organization: "correspondents/00000000-0000-0000-0000-000000000000",
				Type:         "serviceAccount",
			},
		}),
	)

	ctx := context.Background()
	res, err := s.AccountCreation.GetAccount(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", operations.QueryParamViewFull.ToPointer())
	if err != nil {

		var e *sdkerrors.Status
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Name

You can override the default server globally using the `WithServer` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the names associated with the available servers:

| Name | Server | Variables |
| ----- | ------ | --------- |
| `uat` | `https://uat.apexapis.com` | None |
| `prod` | `https://api.apexapis.com` | None |
| `sbx` | `https://sbx.apexapis.com` | None |

#### Example

```go
package main

import (
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/models/operations"
	"log"
)

func main() {
	s := ascendsdkgo.New(
		ascendsdkgo.WithServer("sbx"),
		ascendsdkgo.WithSecurity(components.Security{
			APIKey: ascendsdkgo.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
			ServiceAccountCreds: &components.ServiceAccountCreds{
				PrivateKey:   "-----BEGIN PRIVATE KEY--{OMITTED FOR BREVITY}",
				Name:         "FinFirm",
				Organization: "correspondents/00000000-0000-0000-0000-000000000000",
				Type:         "serviceAccount",
			},
		}),
	)

	ctx := context.Background()
	res, err := s.AccountCreation.GetAccount(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", operations.QueryParamViewFull.ToPointer())
	if err != nil {
		log.Fatal(err)
	}
	if res.Account != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	ascendsdkgo "github.com/afs-public/ascend-sdk-go"
	"github.com/afs-public/ascend-sdk-go/models/components"
	"github.com/afs-public/ascend-sdk-go/models/operations"
	"log"
)

func main() {
	s := ascendsdkgo.New(
		ascendsdkgo.WithServerURL("https://uat.apexapis.com"),
		ascendsdkgo.WithSecurity(components.Security{
			APIKey: ascendsdkgo.String("ABCDEFGHIJ0123456789abcdefghij0123456789"),
			ServiceAccountCreds: &components.ServiceAccountCreds{
				PrivateKey:   "-----BEGIN PRIVATE KEY--{OMITTED FOR BREVITY}",
				Name:         "FinFirm",
				Organization: "correspondents/00000000-0000-0000-0000-000000000000",
				Type:         "serviceAccount",
			},
		}),
	)

	ctx := context.Background()
	res, err := s.AccountCreation.GetAccount(ctx, "01HC3MAQ4DR9QN1V8MJ4CN1HMK", operations.QueryParamViewFull.ToPointer())
	if err != nil {
		log.Fatal(err)
	}
	if res.Account != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- No Authentication [security] -->

<!-- No Table of Contents [toc] -->

<!-- No Summary -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [AccountCreation](docs/sdks/accountcreation/README.md)

* [CreateAccount](docs/sdks/accountcreation/README.md#createaccount) - Create Account
* [GetAccount](docs/sdks/accountcreation/README.md#getaccount) - Get Account

### [AccountManagement](docs/sdks/accountmanagement/README.md)

* [ListAccounts](docs/sdks/accountmanagement/README.md#listaccounts) - List Accounts
* [UpdateAccount](docs/sdks/accountmanagement/README.md#updateaccount) - Update Account
* [AddParty](docs/sdks/accountmanagement/README.md#addparty) - Add Party
* [UpdateParty](docs/sdks/accountmanagement/README.md#updateparty) - Update Party
* [ReplaceParty](docs/sdks/accountmanagement/README.md#replaceparty) - Replace Party
* [RemoveParty](docs/sdks/accountmanagement/README.md#removeparty) - Remove Party
* [CloseAccount](docs/sdks/accountmanagement/README.md#closeaccount) - Close Account
* [CreateTrustedContact](docs/sdks/accountmanagement/README.md#createtrustedcontact) - Create Trusted Contact
* [UpdateTrustedContact](docs/sdks/accountmanagement/README.md#updatetrustedcontact) - Update Trusted Contact
* [DeleteTrustedContact](docs/sdks/accountmanagement/README.md#deletetrustedcontact) - Delete Trusted Contact
* [CreateInterestedParty](docs/sdks/accountmanagement/README.md#createinterestedparty) - Create Interested Party
* [UpdateInterestedParty](docs/sdks/accountmanagement/README.md#updateinterestedparty) - Update Interested Party
* [DeleteInterestedParty](docs/sdks/accountmanagement/README.md#deleteinterestedparty) - Delete Interested Party
* [ListAvailableRestrictions](docs/sdks/accountmanagement/README.md#listavailablerestrictions) - List Available Restrictions
* [CreateRestriction](docs/sdks/accountmanagement/README.md#createrestriction) - Create Restriction
* [EndRestriction](docs/sdks/accountmanagement/README.md#endrestriction) - End Restriction

### [AccountTransfers](docs/sdks/accounttransfers/README.md)

* [CreateTransfer](docs/sdks/accounttransfers/README.md#createtransfer) - Create Transfer
* [GetTransfer](docs/sdks/accounttransfers/README.md#gettransfer) - Get Transfer

### [ACHTransfers](docs/sdks/achtransfers/README.md)

* [CreateAchDeposit](docs/sdks/achtransfers/README.md#createachdeposit) - Create ACH Deposit
* [GetAchDeposit](docs/sdks/achtransfers/README.md#getachdeposit) - Get ACH Deposit
* [CancelAchDeposit](docs/sdks/achtransfers/README.md#cancelachdeposit) - Cancel ACH Deposit
* [CreateAchWithdrawal](docs/sdks/achtransfers/README.md#createachwithdrawal) - Create ACH Withdrawal
* [GetAchWithdrawal](docs/sdks/achtransfers/README.md#getachwithdrawal) - Get ACH Withdrawal
* [CancelAchWithdrawal](docs/sdks/achtransfers/README.md#cancelachwithdrawal) - Cancel ACH Withdrawal

### [Assets](docs/sdks/assets/README.md)

* [ListAssets](docs/sdks/assets/README.md#listassets) - List Assets
* [GetAsset](docs/sdks/assets/README.md#getasset) - Get Asset
* [ListAssetsCorrespondent](docs/sdks/assets/README.md#listassetscorrespondent) - List Assets (By Correspondent)
* [GetAssetCorrespondent](docs/sdks/assets/README.md#getassetcorrespondent) - Get Asset (By Correspondent)

### [Authentication](docs/sdks/authentication/README.md)

* [GenerateServiceAccountToken](docs/sdks/authentication/README.md#generateserviceaccounttoken) - Generate Service Account Token
* [ListSigningKeys](docs/sdks/authentication/README.md#listsigningkeys) - List Signing Keys

### [BankRelationships](docs/sdks/bankrelationships/README.md)

* [CreateBankRelationship](docs/sdks/bankrelationships/README.md#createbankrelationship) - Create Bank Relationship
* [ListBankRelationships](docs/sdks/bankrelationships/README.md#listbankrelationships) - List Bank Relationships
* [GetBankRelationship](docs/sdks/bankrelationships/README.md#getbankrelationship) - Get Bank Relationship
* [UpdateBankRelationship](docs/sdks/bankrelationships/README.md#updatebankrelationship) - Update Bank Relationship
* [CancelBankRelationship](docs/sdks/bankrelationships/README.md#cancelbankrelationship) - Cancel Bank Relationship
* [VerifyMicroDeposits](docs/sdks/bankrelationships/README.md#verifymicrodeposits) - Verify Micro Deposits
* [ReissueMicroDeposits](docs/sdks/bankrelationships/README.md#reissuemicrodeposits) - Reissue Micro Deposits

### [BasketOrders](docs/sdks/basketorders/README.md)

* [CreateBasket](docs/sdks/basketorders/README.md#createbasket) - Create Basket
* [AddOrders](docs/sdks/basketorders/README.md#addorders) - Add Orders
* [GetBasket](docs/sdks/basketorders/README.md#getbasket) - Get Basket
* [SubmitBasket](docs/sdks/basketorders/README.md#submitbasket) - Submit Basket
* [ListBasketOrders](docs/sdks/basketorders/README.md#listbasketorders) - List Basket Orders
* [ListCompressedOrders](docs/sdks/basketorders/README.md#listcompressedorders) - List Compressed Orders

### [CashBalances](docs/sdks/cashbalances/README.md)

* [CalculateCashBalance](docs/sdks/cashbalances/README.md#calculatecashbalance) - Get Cash Balance

### [CreateOrder](docs/sdks/createorder/README.md)

* [CreateOrder](docs/sdks/createorder/README.md#createorder) - Create Order
* [GetOrder](docs/sdks/createorder/README.md#getorder) - Get Order
* [CancelOrder](docs/sdks/createorder/README.md#cancelorder) - Cancel Order

### [DataRetrieval](docs/sdks/dataretrieval/README.md)

* [ListSnapshots](docs/sdks/dataretrieval/README.md#listsnapshots) - List Snapshots

### [EnrollmentsAndAgreements](docs/sdks/enrollmentsandagreements/README.md)

* [EnrollAccount](docs/sdks/enrollmentsandagreements/README.md#enrollaccount) - Enroll Account
* [ListAvailableEnrollments](docs/sdks/enrollmentsandagreements/README.md#listavailableenrollments) - List Available Enrollments
* [DeactivateEnrollment](docs/sdks/enrollmentsandagreements/README.md#deactivateenrollment) - Deactivate Enrollment
* [ListEnrollments](docs/sdks/enrollmentsandagreements/README.md#listenrollments) - List Account Enrollments
* [AffirmAgreements](docs/sdks/enrollmentsandagreements/README.md#affirmagreements) - Affirm Agreements
* [ListAgreements](docs/sdks/enrollmentsandagreements/README.md#listagreements) - List Account Agreements
* [ListEntitlements](docs/sdks/enrollmentsandagreements/README.md#listentitlements) - List Account Entitlements

### [FeesAndCredits](docs/sdks/feesandcredits/README.md)

* [CreateFee](docs/sdks/feesandcredits/README.md#createfee) - Create Fee
* [GetFee](docs/sdks/feesandcredits/README.md#getfee) - Get Fee
* [CancelFee](docs/sdks/feesandcredits/README.md#cancelfee) - Cancel Fee
* [CreateCredit](docs/sdks/feesandcredits/README.md#createcredit) - Create Credit
* [GetCredit](docs/sdks/feesandcredits/README.md#getcredit) - Get Credit
* [CancelCredit](docs/sdks/feesandcredits/README.md#cancelcredit) - Cancel Credit

### [FixedIncomePricing](docs/sdks/fixedincomepricing/README.md)

* [PreviewOrderCost](docs/sdks/fixedincomepricing/README.md#previewordercost) - Preview Order Cost
* [RetrieveQuote](docs/sdks/fixedincomepricing/README.md#retrievequote) - Retrieve Quote
* [RetrieveFixedIncomeMarks](docs/sdks/fixedincomepricing/README.md#retrievefixedincomemarks) - Retrieve Fixed Income Marks

### [InstantCashTransferICT](docs/sdks/instantcashtransferict/README.md)

* [CreateIctDeposit](docs/sdks/instantcashtransferict/README.md#createictdeposit) - Create ICT Deposit
* [GetIctDeposit](docs/sdks/instantcashtransferict/README.md#getictdeposit) - Get ICT Deposit
* [CancelIctDeposit](docs/sdks/instantcashtransferict/README.md#cancelictdeposit) - Cancel ICT Deposit
* [CreateIctWithdrawal](docs/sdks/instantcashtransferict/README.md#createictwithdrawal) - Create ICT Withdrawal
* [GetIctWithdrawal](docs/sdks/instantcashtransferict/README.md#getictwithdrawal) - Get ICT Withdrawal
* [CancelIctWithdrawal](docs/sdks/instantcashtransferict/README.md#cancelictwithdrawal) - Cancel ICT Withdrawal
* [LocateIctReport](docs/sdks/instantcashtransferict/README.md#locateictreport) - Locate ICT Report

### [Investigations](docs/sdks/investigations/README.md)

* [GetInvestigation](docs/sdks/investigations/README.md#getinvestigation) - Get Investigations
* [UpdateInvestigation](docs/sdks/investigations/README.md#updateinvestigation) - Update Investigation 
* [ListInvestigations](docs/sdks/investigations/README.md#listinvestigations) - List Investigations
* [LinkDocuments](docs/sdks/investigations/README.md#linkdocuments) - Link Documents
* [GetWatchlistItem](docs/sdks/investigations/README.md#getwatchlistitem) - Get Watchlist Item

### [InvestorDocs](docs/sdks/investordocs/README.md)

* [BatchCreateUploadLinks](docs/sdks/investordocs/README.md#batchcreateuploadlinks) - Batch Create Upload Links
* [ListDocuments](docs/sdks/investordocs/README.md#listdocuments) - List Documents

### [Journals](docs/sdks/journals/README.md)

* [CreateCashJournal](docs/sdks/journals/README.md#createcashjournal) - Create Cash Journal
* [GetCashJournal](docs/sdks/journals/README.md#getcashjournal) - Get Cash Journal
* [CancelCashJournal](docs/sdks/journals/README.md#cancelcashjournal) - Cancel Cash Journal

### [Ledger](docs/sdks/ledger/README.md)

* [ListEntries](docs/sdks/ledger/README.md#listentries) - List Entries
* [ListActivities](docs/sdks/ledger/README.md#listactivities) - List Activities
* [ListPositions](docs/sdks/ledger/README.md#listpositions) - List Positions
* [GetActivity](docs/sdks/ledger/README.md#getactivity) - Get Activity
* [GetEntry](docs/sdks/ledger/README.md#getentry) - Get Entry

### [Margins](docs/sdks/margins/README.md)

* [GetBuyingPower](docs/sdks/margins/README.md#getbuyingpower) - Get Buying Power

### [PersonManagement](docs/sdks/personmanagement/README.md)

* [CreateLegalNaturalPerson](docs/sdks/personmanagement/README.md#createlegalnaturalperson) - Create Legal Natural Person
* [ListLegalNaturalPersons](docs/sdks/personmanagement/README.md#listlegalnaturalpersons) - List Legal Natural Persons
* [GetLegalNaturalPerson](docs/sdks/personmanagement/README.md#getlegalnaturalperson) - Get Legal Natural Persons
* [UpdateLegalNaturalPerson](docs/sdks/personmanagement/README.md#updatelegalnaturalperson) - Update Legal Natural Person
* [AssignLargeTrader](docs/sdks/personmanagement/README.md#assignlargetrader) - Assign Large Trader
* [EndLargeTraderLegalNaturalPerson](docs/sdks/personmanagement/README.md#endlargetraderlegalnaturalperson) - End Large Trader
* [CreateLegalEntity](docs/sdks/personmanagement/README.md#createlegalentity) - Create Legal Entity
* [ListLegalEntities](docs/sdks/personmanagement/README.md#listlegalentities) - List Legal Entity
* [GetLegalEntity](docs/sdks/personmanagement/README.md#getlegalentity) - Get Legal Entity
* [UpdateLegalEntity](docs/sdks/personmanagement/README.md#updatelegalentity) - Update Legal Entity
* [AssignLargeTraderLegalEntity](docs/sdks/personmanagement/README.md#assignlargetraderlegalentity) - Assign Entity Large Trader
* [EndLargeTrader](docs/sdks/personmanagement/README.md#endlargetrader) - End Entity Large Trader

### [Reader](docs/sdks/reader/README.md)

* [ListEventMessages](docs/sdks/reader/README.md#listeventmessages) - List Event Messages
* [GetEventMessage](docs/sdks/reader/README.md#geteventmessage) - Get Event Message

### [Retirements](docs/sdks/retirements/README.md)

* [ListContributionSummaries](docs/sdks/retirements/README.md#listcontributionsummaries) - List Contribution Summaries
* [RetrieveContributionConstraints](docs/sdks/retirements/README.md#retrievecontributionconstraints) - Retrieve Contribution Constraints
* [RetrieveDistributionConstraints](docs/sdks/retirements/README.md#retrievedistributionconstraints) - Retrieve Distribution Constraints

### [ScheduleTransfers](docs/sdks/scheduletransfers/README.md)

* [ListScheduleSummaries](docs/sdks/scheduletransfers/README.md#listschedulesummaries) - List Schedule Summaries
* [CreateAchDepositSchedule](docs/sdks/scheduletransfers/README.md#createachdepositschedule) - Create ACH Deposit Schedule
* [ListAchDepositSchedules](docs/sdks/scheduletransfers/README.md#listachdepositschedules) - List ACH Deposit Schedules
* [GetAchDepositSchedule](docs/sdks/scheduletransfers/README.md#getachdepositschedule) - Get ACH Deposit Schedule
* [UpdateAchDepositSchedule](docs/sdks/scheduletransfers/README.md#updateachdepositschedule) - Update ACH Deposit Schedules
* [CancelAchDepositSchedule](docs/sdks/scheduletransfers/README.md#cancelachdepositschedule) - Cancel ACH Deposit Schedule
* [CreateAchWithdrawalSchedule](docs/sdks/scheduletransfers/README.md#createachwithdrawalschedule) - Create ACH Withdrawal Schedule
* [ListAchWithdrawalSchedules](docs/sdks/scheduletransfers/README.md#listachwithdrawalschedules) - List ACH Withdrawal Schedules
* [GetAchWithdrawalSchedule](docs/sdks/scheduletransfers/README.md#getachwithdrawalschedule) - Get ACH Withdrawal Schedule
* [UpdateAchWithdrawalSchedule](docs/sdks/scheduletransfers/README.md#updateachwithdrawalschedule) - Update ACH Withdrawal Schedule
* [CancelAchWithdrawalSchedule](docs/sdks/scheduletransfers/README.md#cancelachwithdrawalschedule) - Cancel ACH Withdrawal Schedule
* [CreateWireWithdrawalSchedule](docs/sdks/scheduletransfers/README.md#createwirewithdrawalschedule) - Create Wire Withdrawal Schedule
* [ListWireWithdrawalSchedules](docs/sdks/scheduletransfers/README.md#listwirewithdrawalschedules) - List Wire Withdrawal Schedules
* [GetWireWithdrawalSchedule](docs/sdks/scheduletransfers/README.md#getwirewithdrawalschedule) - Get Wire Withdrawal Schedule
* [UpdateWireWithdrawalSchedule](docs/sdks/scheduletransfers/README.md#updatewirewithdrawalschedule) - Update Wire Withdrawal Schedule
* [CancelWireWithdrawalSchedule](docs/sdks/scheduletransfers/README.md#cancelwirewithdrawalschedule) - Cancel Wire Withdrawal Schedule


### [Subscriber](docs/sdks/subscriber/README.md)

* [CreatePushSubscription](docs/sdks/subscriber/README.md#createpushsubscription) - Create Push Subscription
* [ListPushSubscriptions](docs/sdks/subscriber/README.md#listpushsubscriptions) - List Push Subscriptions
* [GetPushSubscription](docs/sdks/subscriber/README.md#getpushsubscription) - Get Push Subscription
* [UpdatePushSubscription](docs/sdks/subscriber/README.md#updatepushsubscription) - Update Subscription
* [DeletePushSubscription](docs/sdks/subscriber/README.md#deletepushsubscription) - Delete Subscription
* [GetPushSubscriptionDelivery](docs/sdks/subscriber/README.md#getpushsubscriptiondelivery) - Get Subscription Event Delivery
* [ListPushSubscriptionDeliveries](docs/sdks/subscriber/README.md#listpushsubscriptiondeliveries) - List Push Subscription Event Deliveries

### [TestSimulation](docs/sdks/testsimulation/README.md)

* [ForceApproveAchDeposit](docs/sdks/testsimulation/README.md#forceapproveachdeposit) - ACH Deposit Approval
* [ForceNocAchDeposit](docs/sdks/testsimulation/README.md#forcenocachdeposit) - NOC for a Deposit
* [ForceRejectAchDeposit](docs/sdks/testsimulation/README.md#forcerejectachdeposit) - ACH Deposit Rejection
* [ForceReturnAchDeposit](docs/sdks/testsimulation/README.md#forcereturnachdeposit) - ACH Deposit Return
* [ForceApproveAchWithdrawal](docs/sdks/testsimulation/README.md#forceapproveachwithdrawal) - ACH Withdrawal Approval
* [ForceNocAchWithdrawal](docs/sdks/testsimulation/README.md#forcenocachwithdrawal) - ACH Withdrawal NOC
* [ForceRejectAchWithdrawal](docs/sdks/testsimulation/README.md#forcerejectachwithdrawal) - ACH Withdrawal Rejection
* [ForceReturnAchWithdrawal](docs/sdks/testsimulation/README.md#forcereturnachwithdrawal) - ACH Withdrawal Return
* [GetMicroDepositAmounts](docs/sdks/testsimulation/README.md#getmicrodepositamounts) - Get Relationship Micro Deposit Verification
* [ForceApproveIctDeposit](docs/sdks/testsimulation/README.md#forceapproveictdeposit) - Force Approve ICT Deposit
* [ForceRejectIctDeposit](docs/sdks/testsimulation/README.md#forcerejectictdeposit) - Force Reject ICT Deposit
* [ForceApproveIctWithdrawal](docs/sdks/testsimulation/README.md#forceapproveictwithdrawal) - Force Approve ICT Withdrawal
* [ForceRejectIctWithdrawal](docs/sdks/testsimulation/README.md#forcerejectictwithdrawal) - Force Reject ICT Withdrawal
* [ForceApproveWireWithdrawal](docs/sdks/testsimulation/README.md#forceapprovewirewithdrawal) - Force Approve Wire Withdrawal
* [ForceRejectWireWithdrawal](docs/sdks/testsimulation/README.md#forcerejectwirewithdrawal) - Force Reject Wire Withdrawal
* [ForceApproveCashJournal](docs/sdks/testsimulation/README.md#forceapprovecashjournal) - Force Approve Cash Journal
* [ForceRejectCashJournal](docs/sdks/testsimulation/README.md#forcerejectcashjournal) - Force Reject Cash Journal

### [Wires](docs/sdks/wires/README.md)

* [CreateWireWithdrawal](docs/sdks/wires/README.md#createwirewithdrawal) - Create Wire Transfer
* [GetWireWithdrawal](docs/sdks/wires/README.md#getwirewithdrawal) - Get Wire Transfer
* [CancelWireWithdrawal](docs/sdks/wires/README.md#cancelwirewithdrawal) - Cancel Wire Transfer

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->
