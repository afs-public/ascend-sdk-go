// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// Rate - The rate at which monies were withheld, expressed as a value between 0-1
type Rate struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *Rate) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

type WithholdingDataState string

const (
	WithholdingDataStateWithholdingStateUnspecified WithholdingDataState = "WITHHOLDING_STATE_UNSPECIFIED"
	WithholdingDataStateCa                          WithholdingDataState = "CA"
	WithholdingDataStateMe                          WithholdingDataState = "ME"
	WithholdingDataStateVt                          WithholdingDataState = "VT"
	WithholdingDataStateAl                          WithholdingDataState = "AL"
	WithholdingDataStateAk                          WithholdingDataState = "AK"
	WithholdingDataStateAz                          WithholdingDataState = "AZ"
	WithholdingDataStateAr                          WithholdingDataState = "AR"
	WithholdingDataStateCo                          WithholdingDataState = "CO"
	WithholdingDataStateCt                          WithholdingDataState = "CT"
	WithholdingDataStateDe                          WithholdingDataState = "DE"
	WithholdingDataStateFl                          WithholdingDataState = "FL"
	WithholdingDataStateGa                          WithholdingDataState = "GA"
	WithholdingDataStateHi                          WithholdingDataState = "HI"
	WithholdingDataStateID                          WithholdingDataState = "ID"
	WithholdingDataStateIl                          WithholdingDataState = "IL"
	WithholdingDataStateIn                          WithholdingDataState = "IN"
	WithholdingDataStateIa                          WithholdingDataState = "IA"
	WithholdingDataStateKs                          WithholdingDataState = "KS"
	WithholdingDataStateKy                          WithholdingDataState = "KY"
	WithholdingDataStateLa                          WithholdingDataState = "LA"
	WithholdingDataStateMd                          WithholdingDataState = "MD"
	WithholdingDataStateMa                          WithholdingDataState = "MA"
	WithholdingDataStateMi                          WithholdingDataState = "MI"
	WithholdingDataStateMn                          WithholdingDataState = "MN"
	WithholdingDataStateMs                          WithholdingDataState = "MS"
	WithholdingDataStateMo                          WithholdingDataState = "MO"
	WithholdingDataStateMt                          WithholdingDataState = "MT"
	WithholdingDataStateNe                          WithholdingDataState = "NE"
	WithholdingDataStateNv                          WithholdingDataState = "NV"
	WithholdingDataStateNh                          WithholdingDataState = "NH"
	WithholdingDataStateNj                          WithholdingDataState = "NJ"
	WithholdingDataStateNm                          WithholdingDataState = "NM"
	WithholdingDataStateNy                          WithholdingDataState = "NY"
	WithholdingDataStateNc                          WithholdingDataState = "NC"
	WithholdingDataStateNd                          WithholdingDataState = "ND"
	WithholdingDataStateOh                          WithholdingDataState = "OH"
	WithholdingDataStateOk                          WithholdingDataState = "OK"
	WithholdingDataStateOr                          WithholdingDataState = "OR"
	WithholdingDataStatePa                          WithholdingDataState = "PA"
	WithholdingDataStateRi                          WithholdingDataState = "RI"
	WithholdingDataStateSc                          WithholdingDataState = "SC"
	WithholdingDataStateSd                          WithholdingDataState = "SD"
	WithholdingDataStateTn                          WithholdingDataState = "TN"
	WithholdingDataStateTx                          WithholdingDataState = "TX"
	WithholdingDataStateUt                          WithholdingDataState = "UT"
	WithholdingDataStateVa                          WithholdingDataState = "VA"
	WithholdingDataStateWa                          WithholdingDataState = "WA"
	WithholdingDataStateWv                          WithholdingDataState = "WV"
	WithholdingDataStateWi                          WithholdingDataState = "WI"
	WithholdingDataStateWy                          WithholdingDataState = "WY"
)

func (e WithholdingDataState) ToPointer() *WithholdingDataState {
	return &e
}

// WithholdingDataType - Provides more detail on the type of the withholding (Federal, State, etc.)
type WithholdingDataType string

const (
	WithholdingDataTypeWithholdingTypeUnspecified WithholdingDataType = "WITHHOLDING_TYPE_UNSPECIFIED"
	WithholdingDataTypeFederal                    WithholdingDataType = "FEDERAL"
	WithholdingDataTypeState                      WithholdingDataType = "STATE"
	WithholdingDataTypeForeignSecurity            WithholdingDataType = "FOREIGN_SECURITY"
	WithholdingDataTypeFederalIra                 WithholdingDataType = "FEDERAL_IRA"
	WithholdingDataTypeStateIra                   WithholdingDataType = "STATE_IRA"
	WithholdingDataTypeNonResidentAlien           WithholdingDataType = "NON_RESIDENT_ALIEN"
)

func (e WithholdingDataType) ToPointer() *WithholdingDataType {
	return &e
}

type WithholdingData struct {
	// The rate at which monies were withheld, expressed as a value between 0-1
	Rate *Rate `json:"rate,omitempty"`
	// 2-Letter alpha code for the names of countries, dependent territories, and special areas of geographical interest. Complies with ISO-3166 Alpha-2 Codes
	RegionCode *string               `json:"region_code,omitempty"`
	State      *WithholdingDataState `json:"state,omitempty"`
	// The tax year associated with the withholding
	TaxYear *int `json:"tax_year,omitempty"`
	// Provides more detail on the type of the withholding (Federal, State, etc.)
	Type *WithholdingDataType `json:"type,omitempty"`
}

func (o *WithholdingData) GetRate() *Rate {
	if o == nil {
		return nil
	}
	return o.Rate
}

func (o *WithholdingData) GetRegionCode() *string {
	if o == nil {
		return nil
	}
	return o.RegionCode
}

func (o *WithholdingData) GetState() *WithholdingDataState {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *WithholdingData) GetTaxYear() *int {
	if o == nil {
		return nil
	}
	return o.TaxYear
}

func (o *WithholdingData) GetType() *WithholdingDataType {
	if o == nil {
		return nil
	}
	return o.Type
}