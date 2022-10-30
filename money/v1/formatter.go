package moneypb

import (
	"github.com/shopspring/decimal"

	"github.com/sipki-tech/money"
)

// Money returns money format.
func (x *Money) Money() *money.Money {
	moneyType := money.Fiat
	if x.Type != Type_TYPE_FIAT {
		moneyType = money.Crypto
	}

	return &money.Money{
		Code:    x.CurrencyCode.Code(),
		Decimal: decimal.RequireFromString(x.Decimal.Value),
		Type:    moneyType,
	}
}
