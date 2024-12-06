// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// Program - The name of the program the ICT deposit is associated with
type Program string

const (
	ProgramIctProgramUnspecified Program = "ICT_PROGRAM_UNSPECIFIED"
	ProgramBrokerPartner         Program = "BROKER_PARTNER"
	ProgramDepositOnly           Program = "DEPOSIT_ONLY"
	ProgramBankingPartner        Program = "BANKING_PARTNER"
	ProgramMoneyTransmitter      Program = "MONEY_TRANSMITTER"
	ProgramWithdrawalOnly        Program = "WITHDRAWAL_ONLY"
	ProgramDigitalPartner        Program = "DIGITAL_PARTNER"
)

func (e Program) ToPointer() *Program {
	return &e
}
func (e *Program) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ICT_PROGRAM_UNSPECIFIED":
		fallthrough
	case "BROKER_PARTNER":
		fallthrough
	case "DEPOSIT_ONLY":
		fallthrough
	case "BANKING_PARTNER":
		fallthrough
	case "MONEY_TRANSMITTER":
		fallthrough
	case "WITHDRAWAL_ONLY":
		fallthrough
	case "DIGITAL_PARTNER":
		*e = Program(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Program: %v", v)
	}
}

// IctDepositCreate - An Instant Cash Transfer. Funds are moved from a configured Firm account to a customer's brokerage account.
type IctDepositCreate struct {
	// A representation of a decimal value, such as 2.5. Clients may convert values into language-native decimal formats, such as Java's [BigDecimal][] or Python's [decimal.Decimal][].
	//
	//  [BigDecimal]:
	//  https://docs.oracle.com/en/java/javase/11/docs/api/java.base/java/math/BigDecimal.html
	//  [decimal.Decimal]: https://docs.python.org/3/library/decimal.html
	Amount DecimalCreate `json:"amount"`
	// External identifier supplied by the API caller. Each request must have a unique pairing of client_transfer_id and account.
	ClientTransferID string `json:"client_transfer_id"`
	// The name of the program the ICT deposit is associated with
	Program Program `json:"program"`
	// A contribution to a retirement account.
	RetirementContribution *RetirementContributionCreate `json:"retirement_contribution,omitempty"`
	// The travel rules associated with an ICT deposit
	TravelRule IctDepositTravelRuleCreate `json:"travel_rule"`
}

func (o *IctDepositCreate) GetAmount() DecimalCreate {
	if o == nil {
		return DecimalCreate{}
	}
	return o.Amount
}

func (o *IctDepositCreate) GetClientTransferID() string {
	if o == nil {
		return ""
	}
	return o.ClientTransferID
}

func (o *IctDepositCreate) GetProgram() Program {
	if o == nil {
		return Program("")
	}
	return o.Program
}

func (o *IctDepositCreate) GetRetirementContribution() *RetirementContributionCreate {
	if o == nil {
		return nil
	}
	return o.RetirementContribution
}

func (o *IctDepositCreate) GetTravelRule() IctDepositTravelRuleCreate {
	if o == nil {
		return IctDepositTravelRuleCreate{}
	}
	return o.TravelRule
}