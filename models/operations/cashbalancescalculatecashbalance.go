// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"ascend-sdk/models/components"
	"encoding/json"
	"fmt"
)

// Mechanism - The withdraw mechanism to calculate the balance for. The mechanism determines what account activity will affect the balance.
type Mechanism string

const (
	MechanismMechanismUnspecified Mechanism = "MECHANISM_UNSPECIFIED"
	MechanismAcat                 Mechanism = "ACAT"
	MechanismAch                  Mechanism = "ACH"
	MechanismCashJournal          Mechanism = "CASH_JOURNAL"
	MechanismCheck                Mechanism = "CHECK"
	MechanismCredit               Mechanism = "CREDIT"
	MechanismFee                  Mechanism = "FEE"
	MechanismIct                  Mechanism = "ICT"
	MechanismPaypal               Mechanism = "PAYPAL"
	MechanismRtp                  Mechanism = "RTP"
	MechanismTpj                  Mechanism = "TPJ"
	MechanismWire                 Mechanism = "WIRE"
)

func (e Mechanism) ToPointer() *Mechanism {
	return &e
}
func (e *Mechanism) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MECHANISM_UNSPECIFIED":
		fallthrough
	case "ACAT":
		fallthrough
	case "ACH":
		fallthrough
	case "CASH_JOURNAL":
		fallthrough
	case "CHECK":
		fallthrough
	case "CREDIT":
		fallthrough
	case "FEE":
		fallthrough
	case "ICT":
		fallthrough
	case "PAYPAL":
		fallthrough
	case "RTP":
		fallthrough
	case "TPJ":
		fallthrough
	case "WIRE":
		*e = Mechanism(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Mechanism: %v", v)
	}
}

type CashBalancesCalculateCashBalanceRequest struct {
	// The account id.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The withdraw mechanism to calculate the balance for. The mechanism determines what account activity will affect the balance.
	Mechanism *Mechanism `queryParam:"style=form,explode=true,name=mechanism"`
}

func (o *CashBalancesCalculateCashBalanceRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *CashBalancesCalculateCashBalanceRequest) GetMechanism() *Mechanism {
	if o == nil {
		return nil
	}
	return o.Mechanism
}

type CashBalancesCalculateCashBalanceResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	CalculateCashBalanceResponse *components.CalculateCashBalanceResponse
	// INVALID_ARGUMENT: The request has an invalid argument.
	Status *components.Status
}

func (o *CashBalancesCalculateCashBalanceResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CashBalancesCalculateCashBalanceResponse) GetCalculateCashBalanceResponse() *components.CalculateCashBalanceResponse {
	if o == nil {
		return nil
	}
	return o.CalculateCashBalanceResponse
}

func (o *CashBalancesCalculateCashBalanceResponse) GetStatus() *components.Status {
	if o == nil {
		return nil
	}
	return o.Status
}
