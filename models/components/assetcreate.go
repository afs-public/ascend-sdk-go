// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// AssetCreateType - The asset identifier type
type AssetCreateType string

const (
	AssetCreateTypeIdentifierTypeUnspecified AssetCreateType = "IDENTIFIER_TYPE_UNSPECIFIED"
	AssetCreateTypeCurrencyCode              AssetCreateType = "CURRENCY_CODE"
	AssetCreateTypeCusip                     AssetCreateType = "CUSIP"
	AssetCreateTypeSymbol                    AssetCreateType = "SYMBOL"
	AssetCreateTypeIsin                      AssetCreateType = "ISIN"
	AssetCreateTypeAssetID                   AssetCreateType = "ASSET_ID"
)

func (e AssetCreateType) ToPointer() *AssetCreateType {
	return &e
}

// AssetCreate - The asset being transferred If cash, the asset_id is the currency code (e.g. USD) and the position is the amount
type AssetCreate struct {
	// The asset identifier
	Identifier string `json:"identifier"`
	// The position or amount of the asset
	Position PositionCreate `json:"position"`
	// The asset identifier type
	Type AssetCreateType `json:"type"`
}

func (o *AssetCreate) GetIdentifier() string {
	if o == nil {
		return ""
	}
	return o.Identifier
}

func (o *AssetCreate) GetPosition() PositionCreate {
	if o == nil {
		return PositionCreate{}
	}
	return o.Position
}

func (o *AssetCreate) GetType() AssetCreateType {
	if o == nil {
		return AssetCreateType("")
	}
	return o.Type
}