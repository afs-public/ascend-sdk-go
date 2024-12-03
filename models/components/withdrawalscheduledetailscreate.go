// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// WithdrawalScheduleDetailsCreate - Details of withdrawal schedule transfers
type WithdrawalScheduleDetailsCreate struct {
	// A representation of a decimal value, such as 2.5. Clients may convert values into language-native decimal formats, such as Java's [BigDecimal][] or Python's [decimal.Decimal][].
	//
	//  [BigDecimal]:
	//  https://docs.oracle.com/en/java/javase/11/docs/api/java.base/java/math/BigDecimal.html
	//  [decimal.Decimal]: https://docs.python.org/3/library/decimal.html
	Amount *DecimalCreate `json:"amount,omitempty"`
	// External identifier supplied by the API caller. Each request must have a unique pairing of client_schedule_id and account
	ClientScheduleID string `json:"client_schedule_id"`
	// Flag to indicate a full disbursement transfer (mutually exclusive with 'amount')
	FullDisbursement *bool `json:"full_disbursement,omitempty"`
	// Properties common to all transfer schedules
	ScheduleProperties SchedulePropertiesCreate `json:"schedule_properties"`
}

func (o *WithdrawalScheduleDetailsCreate) GetAmount() *DecimalCreate {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *WithdrawalScheduleDetailsCreate) GetClientScheduleID() string {
	if o == nil {
		return ""
	}
	return o.ClientScheduleID
}

func (o *WithdrawalScheduleDetailsCreate) GetFullDisbursement() *bool {
	if o == nil {
		return nil
	}
	return o.FullDisbursement
}

func (o *WithdrawalScheduleDetailsCreate) GetScheduleProperties() SchedulePropertiesCreate {
	if o == nil {
		return SchedulePropertiesCreate{}
	}
	return o.ScheduleProperties
}
