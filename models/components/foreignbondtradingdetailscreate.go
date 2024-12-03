// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// ForeignBondTradingDetailsCreate - Foreign bond trading detail
type ForeignBondTradingDetailsCreate struct {
	// Does the account anticipate trading in foreign bonds
	ForeignBondTrading bool `json:"foreign_bond_trading"`
	// The foreign bond trading countries details. If yes, than please provide details
	ForeignBondTradingDetail []ForeignBondTradingDetailCreate `json:"foreign_bond_trading_detail,omitempty"`
}

func (o *ForeignBondTradingDetailsCreate) GetForeignBondTrading() bool {
	if o == nil {
		return false
	}
	return o.ForeignBondTrading
}

func (o *ForeignBondTradingDetailsCreate) GetForeignBondTradingDetail() []ForeignBondTradingDetailCreate {
	if o == nil {
		return nil
	}
	return o.ForeignBondTradingDetail
}
