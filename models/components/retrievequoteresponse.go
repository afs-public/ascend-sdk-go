// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/afs-public/ascend-sdk-go/internal/utils"
)

// AskMinimumQuantity - The best ask minimum quantity. This will be absent if no ask information is available
type AskMinimumQuantity struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *AskMinimumQuantity) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// AskQuantity - The best ask quantity. This will be absent if no ask information is available
type AskQuantity struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *AskQuantity) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// RetrieveQuoteResponseAssetType - The type of asset referenced by the security identifier
type RetrieveQuoteResponseAssetType string

const (
	RetrieveQuoteResponseAssetTypeAssetTypeUnspecified RetrieveQuoteResponseAssetType = "ASSET_TYPE_UNSPECIFIED"
	RetrieveQuoteResponseAssetTypeEquity               RetrieveQuoteResponseAssetType = "EQUITY"
	RetrieveQuoteResponseAssetTypeFixedIncome          RetrieveQuoteResponseAssetType = "FIXED_INCOME"
	RetrieveQuoteResponseAssetTypeMutualFund           RetrieveQuoteResponseAssetType = "MUTUAL_FUND"
)

func (e RetrieveQuoteResponseAssetType) ToPointer() *RetrieveQuoteResponseAssetType {
	return &e
}
func (e *RetrieveQuoteResponseAssetType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASSET_TYPE_UNSPECIFIED":
		fallthrough
	case "EQUITY":
		fallthrough
	case "FIXED_INCOME":
		fallthrough
	case "MUTUAL_FUND":
		*e = RetrieveQuoteResponseAssetType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveQuoteResponseAssetType: %v", v)
	}
}

// BidMinimumQuantity - The best bid minimum quantity. This will be absent if no bid information is available
type BidMinimumQuantity struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *BidMinimumQuantity) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// BidQuantity - The best bid quantity. This will be absent if no bid information is available
type BidQuantity struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *BidQuantity) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// RetrieveQuoteResponseBrokerCapacity - Capacity used in determining bid and ask prices. Defaults to "AGENCY" if no value specified.
type RetrieveQuoteResponseBrokerCapacity string

const (
	RetrieveQuoteResponseBrokerCapacityBrokerCapacityUnspecified RetrieveQuoteResponseBrokerCapacity = "BROKER_CAPACITY_UNSPECIFIED"
	RetrieveQuoteResponseBrokerCapacityAgency                    RetrieveQuoteResponseBrokerCapacity = "AGENCY"
	RetrieveQuoteResponseBrokerCapacityPrincipal                 RetrieveQuoteResponseBrokerCapacity = "PRINCIPAL"
)

func (e RetrieveQuoteResponseBrokerCapacity) ToPointer() *RetrieveQuoteResponseBrokerCapacity {
	return &e
}
func (e *RetrieveQuoteResponseBrokerCapacity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BROKER_CAPACITY_UNSPECIFIED":
		fallthrough
	case "AGENCY":
		fallthrough
	case "PRINCIPAL":
		*e = RetrieveQuoteResponseBrokerCapacity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveQuoteResponseBrokerCapacity: %v", v)
	}
}

// RetrieveQuoteResponseIdentifierType - The identifier type of the asset for which the best bid and offer is returned. This will be the same value as what was sent on the request.
type RetrieveQuoteResponseIdentifierType string

const (
	RetrieveQuoteResponseIdentifierTypeAssetID RetrieveQuoteResponseIdentifierType = "ASSET_ID"
	RetrieveQuoteResponseIdentifierTypeCusip   RetrieveQuoteResponseIdentifierType = "CUSIP"
	RetrieveQuoteResponseIdentifierTypeIsin    RetrieveQuoteResponseIdentifierType = "ISIN"
)

func (e RetrieveQuoteResponseIdentifierType) ToPointer() *RetrieveQuoteResponseIdentifierType {
	return &e
}
func (e *RetrieveQuoteResponseIdentifierType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASSET_ID":
		fallthrough
	case "CUSIP":
		fallthrough
	case "ISIN":
		*e = RetrieveQuoteResponseIdentifierType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RetrieveQuoteResponseIdentifierType: %v", v)
	}
}

// MinimumQuantity - The minimum quantity specified in the request (if any). For Fixed Income: Expressed in the par (face-value) amount and may not exceed two decimal places for USD-based currencies.
type MinimumQuantity struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *MinimumQuantity) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// RetrieveQuoteResponse - The current best bid and offer for a specified security
type RetrieveQuoteResponse struct {
	// The account id used in calculating order costs
	AccountID *string `json:"account_id,omitempty"`
	// The best ask minimum quantity. This will be absent if no ask information is available
	AskMinimumQuantity *AskMinimumQuantity `json:"ask_minimum_quantity,omitempty"`
	// The price information of the best ask matching the request criteria. If no ask information is available at all, this will be empty. If yield information is not available, then this will only contain a PERCENTAGE_OF_PAR-typed price.
	AskPrices []BidAskPrice `json:"ask_prices,omitempty"`
	// The best ask quantity. This will be absent if no ask information is available
	AskQuantity *AskQuantity `json:"ask_quantity,omitempty"`
	// Apex Asset ID for this asset.
	AssetID *string `json:"asset_id,omitempty"`
	// The type of asset referenced by the security identifier
	AssetType *RetrieveQuoteResponseAssetType `json:"asset_type,omitempty"`
	// The best bid minimum quantity. This will be absent if no bid information is available
	BidMinimumQuantity *BidMinimumQuantity `json:"bid_minimum_quantity,omitempty"`
	// The price information of the best bid matching the request criteria. If no bid information is available at all, this will be empty. If yield information is not available, then this will only contain a PERCENTAGE_OF_PAR-typed price.
	BidPrices []BidAskPrice `json:"bid_prices,omitempty"`
	// The best bid quantity. This will be absent if no bid information is available
	BidQuantity *BidQuantity `json:"bid_quantity,omitempty"`
	// Capacity used in determining bid and ask prices. Defaults to "AGENCY" if no value specified.
	BrokerCapacity *RetrieveQuoteResponseBrokerCapacity `json:"broker_capacity,omitempty"`
	// Identifier of the asset (of the type specified in `identifier_type`).
	Identifier *string `json:"identifier,omitempty"`
	// The identifier type of the asset for which the best bid and offer is returned. This will be the same value as what was sent on the request.
	IdentifierType *RetrieveQuoteResponseIdentifierType `json:"identifier_type,omitempty"`
	// The minimum quantity specified in the request (if any). For Fixed Income: Expressed in the par (face-value) amount and may not exceed two decimal places for USD-based currencies.
	MinimumQuantity *MinimumQuantity `json:"minimum_quantity,omitempty"`
	// The time the response was generated by the system
	ResponseGenerationTime *time.Time `json:"response_generation_time,omitempty"`
}

func (r RetrieveQuoteResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetrieveQuoteResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetrieveQuoteResponse) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *RetrieveQuoteResponse) GetAskMinimumQuantity() *AskMinimumQuantity {
	if o == nil {
		return nil
	}
	return o.AskMinimumQuantity
}

func (o *RetrieveQuoteResponse) GetAskPrices() []BidAskPrice {
	if o == nil {
		return nil
	}
	return o.AskPrices
}

func (o *RetrieveQuoteResponse) GetAskQuantity() *AskQuantity {
	if o == nil {
		return nil
	}
	return o.AskQuantity
}

func (o *RetrieveQuoteResponse) GetAssetID() *string {
	if o == nil {
		return nil
	}
	return o.AssetID
}

func (o *RetrieveQuoteResponse) GetAssetType() *RetrieveQuoteResponseAssetType {
	if o == nil {
		return nil
	}
	return o.AssetType
}

func (o *RetrieveQuoteResponse) GetBidMinimumQuantity() *BidMinimumQuantity {
	if o == nil {
		return nil
	}
	return o.BidMinimumQuantity
}

func (o *RetrieveQuoteResponse) GetBidPrices() []BidAskPrice {
	if o == nil {
		return nil
	}
	return o.BidPrices
}

func (o *RetrieveQuoteResponse) GetBidQuantity() *BidQuantity {
	if o == nil {
		return nil
	}
	return o.BidQuantity
}

func (o *RetrieveQuoteResponse) GetBrokerCapacity() *RetrieveQuoteResponseBrokerCapacity {
	if o == nil {
		return nil
	}
	return o.BrokerCapacity
}

func (o *RetrieveQuoteResponse) GetIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.Identifier
}

func (o *RetrieveQuoteResponse) GetIdentifierType() *RetrieveQuoteResponseIdentifierType {
	if o == nil {
		return nil
	}
	return o.IdentifierType
}

func (o *RetrieveQuoteResponse) GetMinimumQuantity() *MinimumQuantity {
	if o == nil {
		return nil
	}
	return o.MinimumQuantity
}

func (o *RetrieveQuoteResponse) GetResponseGenerationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.ResponseGenerationTime
}