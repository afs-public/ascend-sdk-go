// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// AssetCategory - The NSCC asset category
type AssetCategory string

const (
	AssetCategoryAssetCategoryUnspecified  AssetCategory = "ASSET_CATEGORY_UNSPECIFIED"
	AssetCategoryAlternativeInvestment     AssetCategory = "ALTERNATIVE_INVESTMENT"
	AssetCategoryAnnuity                   AssetCategory = "ANNUITY"
	AssetCategoryAuctionRatePreferredUit   AssetCategory = "AUCTION_RATE_PREFERRED_UIT"
	AssetCategoryCdAndCommercialPaper      AssetCategory = "CD_AND_COMMERCIAL_PAPER"
	AssetCategoryCorporateBond             AssetCategory = "CORPORATE_BOND"
	AssetCategoryDeferredSaleChargeUit     AssetCategory = "DEFERRED_SALE_CHARGE_UIT"
	AssetCategoryEquity                    AssetCategory = "EQUITY"
	AssetCategoryForeignCurrency           AssetCategory = "FOREIGN_CURRENCY"
	AssetCategoryForeignDebt               AssetCategory = "FOREIGN_DEBT"
	AssetCategoryForeignEquity             AssetCategory = "FOREIGN_EQUITY"
	AssetCategoryLifeInsurance             AssetCategory = "LIFE_INSURANCE"
	AssetCategoryLimitedPartnership        AssetCategory = "LIMITED_PARTNERSHIP"
	AssetCategoryMortgageBackedSecurity    AssetCategory = "MORTGAGE_BACKED_SECURITY"
	AssetCategoryMutualFundMoneyMarket     AssetCategory = "MUTUAL_FUND_MONEY_MARKET"
	AssetCategoryMutualFundNonMoneyMarket  AssetCategory = "MUTUAL_FUND_NON_MONEY_MARKET"
	AssetCategoryMunicipalBond             AssetCategory = "MUNICIPAL_BOND"
	AssetCategoryOption                    AssetCategory = "OPTION"
	AssetCategoryRealEstateInvestmentTrust AssetCategory = "REAL_ESTATE_INVESTMENT_TRUST"
	AssetCategoryRight                     AssetCategory = "RIGHT"
	AssetCategoryUsGovernment              AssetCategory = "US_GOVERNMENT"
	AssetCategoryUnitInvestmentTrust       AssetCategory = "UNIT_INVESTMENT_TRUST"
	AssetCategoryUnit                      AssetCategory = "UNIT"
	AssetCategoryWarrant                   AssetCategory = "WARRANT"
	AssetCategoryZeroCouponBond            AssetCategory = "ZERO_COUPON_BOND"
)

func (e AssetCategory) ToPointer() *AssetCategory {
	return &e
}

// AcatsAssetQuantity - The quantity of the asset, or the amount if the asset is cash; negative quantity denotes short position Fractional amounts only supported for certain asset types
type AcatsAssetQuantity struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *AcatsAssetQuantity) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// AcatsAssetPosition - The position or amount of the asset
type AcatsAssetPosition struct {
	// The quantity of the asset, or the amount if the asset is cash; negative quantity denotes short position Fractional amounts only supported for certain asset types
	Quantity *AcatsAssetQuantity `json:"quantity,omitempty"`
}

func (o *AcatsAssetPosition) GetQuantity() *AcatsAssetQuantity {
	if o == nil {
		return nil
	}
	return o.Quantity
}

// AcatsAssetType - The asset identifier type
type AcatsAssetType string

const (
	AcatsAssetTypeIdentifierTypeUnspecified AcatsAssetType = "IDENTIFIER_TYPE_UNSPECIFIED"
	AcatsAssetTypeCurrencyCode              AcatsAssetType = "CURRENCY_CODE"
	AcatsAssetTypeCusip                     AcatsAssetType = "CUSIP"
	AcatsAssetTypeSymbol                    AcatsAssetType = "SYMBOL"
	AcatsAssetTypeIsin                      AcatsAssetType = "ISIN"
	AcatsAssetTypeAssetID                   AcatsAssetType = "ASSET_ID"
)

func (e AcatsAssetType) ToPointer() *AcatsAssetType {
	return &e
}

// AcatsAsset - The asset being transferred If cash, the asset_id is the currency code (e.g. USD) and the position is the amount
type AcatsAsset struct {
	// The NSCC asset category
	AssetCategory *AssetCategory `json:"asset_category,omitempty"`
	// The asset identifier
	AssetID *string `json:"asset_id,omitempty"`
	// The asset identifier
	Identifier *string `json:"identifier,omitempty"`
	// The position or amount of the asset
	Position *AcatsAssetPosition `json:"position,omitempty"`
	// The asset identifier type
	Type *AcatsAssetType `json:"type,omitempty"`
}

func (o *AcatsAsset) GetAssetCategory() *AssetCategory {
	if o == nil {
		return nil
	}
	return o.AssetCategory
}

func (o *AcatsAsset) GetAssetID() *string {
	if o == nil {
		return nil
	}
	return o.AssetID
}

func (o *AcatsAsset) GetIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.Identifier
}

func (o *AcatsAsset) GetPosition() *AcatsAssetPosition {
	if o == nil {
		return nil
	}
	return o.Position
}

func (o *AcatsAsset) GetType() *AcatsAssetType {
	if o == nil {
		return nil
	}
	return o.Type
}
