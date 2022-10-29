package money

import (
	"github.com/shopspring/decimal"
	"github.com/sipki-tech/currency"
)

// Fiat represent fiat money.
type Fiat struct {
	Code    currency.Code
	Decimal decimal.Decimal
}
