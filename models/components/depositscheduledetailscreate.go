// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// DepositScheduleDetailsCreate - Details of deposit schedule transfers
type DepositScheduleDetailsCreate struct {
	// A representation of a decimal value, such as 2.5. Clients may convert values into language-native decimal formats, such as Java's [BigDecimal][] or Python's [decimal.Decimal][].
	//
	//  [BigDecimal]:
	//  https://docs.oracle.com/en/java/javase/11/docs/api/java.base/java/math/BigDecimal.html
	//  [decimal.Decimal]: https://docs.python.org/3/library/decimal.html
	Amount DecimalCreate `json:"amount"`
	// External identifier supplied by the API caller. Each request must have a unique pairing of client_schedule_id and account
	ClientScheduleID string `json:"client_schedule_id"`
	// Properties common to all transfer schedules
	ScheduleProperties SchedulePropertiesCreate `json:"schedule_properties"`
}

func (o *DepositScheduleDetailsCreate) GetAmount() DecimalCreate {
	if o == nil {
		return DecimalCreate{}
	}
	return o.Amount
}

func (o *DepositScheduleDetailsCreate) GetClientScheduleID() string {
	if o == nil {
		return ""
	}
	return o.ClientScheduleID
}

func (o *DepositScheduleDetailsCreate) GetScheduleProperties() SchedulePropertiesCreate {
	if o == nil {
		return SchedulePropertiesCreate{}
	}
	return o.ScheduleProperties
}
