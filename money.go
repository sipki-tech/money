package money

import (
	"github.com/shopspring/decimal"
	"github.com/sipki-tech/currency"
)

// Money represent fiat money.
type Money struct {
	Code    currency.Code
	Decimal decimal.Decimal
	Type    Type
}
