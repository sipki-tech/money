package moneypb

import (
	"github.com/shopspring/decimal"

	"github.com/sipki-tech/money"
)

// Fiat returns money format.
func (x *Fiat) Fiat() *money.Fiat {
	return &money.Fiat{
		Code:    x.CurrencyCode.Code(),
		Decimal: decimal.RequireFromString(x.Decimal.Value),
	}
}
