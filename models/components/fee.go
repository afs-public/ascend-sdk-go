// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// FeeAmount - Monetary amount associated with the fee
type FeeAmount struct {
	// The decimal value, as a string; Refer to [Google’s Decimal type protocol buffer](https://github.com/googleapis/googleapis/blob/40203ca1880849480bbff7b8715491060bbccdf1/google/type/decimal.proto#L33) for details
	Value *string `json:"value,omitempty"`
}

func (o *FeeAmount) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

// FeeType - The type of fee being assessed
type FeeType string

const (
	FeeTypeFeeTypeUnspecified             FeeType = "FEE_TYPE_UNSPECIFIED"
	FeeTypeClientClearing                 FeeType = "CLIENT_CLEARING"
	FeeTypeLiquidity                      FeeType = "LIQUIDITY"
	FeeTypeGeneralPurpose                 FeeType = "GENERAL_PURPOSE"
	FeeTypeCommission                     FeeType = "COMMISSION"
	FeeTypeOrf                            FeeType = "ORF"
	FeeTypeTaf                            FeeType = "TAF"
	FeeTypeSec                            FeeType = "SEC"
	FeeTypeAccountClosing                 FeeType = "ACCOUNT_CLOSING"
	FeeTypeAccountIra                     FeeType = "ACCOUNT_IRA"
	FeeTypeAchReturn                      FeeType = "ACH_RETURN"
	FeeTypeAdvisory                       FeeType = "ADVISORY"
	FeeTypeCheckFee                       FeeType = "CHECK_FEE"
	FeeTypeExchange                       FeeType = "EXCHANGE"
	FeeTypeManagement                     FeeType = "MANAGEMENT"
	FeeTypeOvernight                      FeeType = "OVERNIGHT"
	FeeTypePlatform                       FeeType = "PLATFORM"
	FeeTypeStatement                      FeeType = "STATEMENT"
	FeeTypeStopPayment                    FeeType = "STOP_PAYMENT"
	FeeTypeWireFee                        FeeType = "WIRE_FEE"
	FeeTypeInactivity                     FeeType = "INACTIVITY"
	FeeTypeAmaService                     FeeType = "AMA_SERVICE"
	FeeTypeNoticeOfChange                 FeeType = "NOTICE_OF_CHANGE"
	FeeTypeAccountTransfer                FeeType = "ACCOUNT_TRANSFER"
	FeeTypeAgencyProcessing               FeeType = "AGENCY_PROCESSING"
	FeeTypeRtpFee                         FeeType = "RTP_FEE"
	FeeTypeDomesticWireDepositFee         FeeType = "DOMESTIC_WIRE_DEPOSIT_FEE"
	FeeTypeDomesticWireWithdrawalFee      FeeType = "DOMESTIC_WIRE_WITHDRAWAL_FEE"
	FeeTypeInternationalWireDepositFee    FeeType = "INTERNATIONAL_WIRE_DEPOSIT_FEE"
	FeeTypeInternationalWireWithdrawalFee FeeType = "INTERNATIONAL_WIRE_WITHDRAWAL_FEE"
	FeeTypeBrokerFee                      FeeType = "BROKER_FEE"
)

func (e FeeType) ToPointer() *FeeType {
	return &e
}

type Fee struct {
	// Monetary amount associated with the fee
	Amount *FeeAmount `json:"amount,omitempty"`
	// The type of fee being assessed
	Type *FeeType `json:"type,omitempty"`
}

func (o *Fee) GetAmount() *FeeAmount {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *Fee) GetType() *FeeType {
	if o == nil {
		return nil
	}
	return o.Type
}
