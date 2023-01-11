package money

import (
	"github.com/shopspring/decimal"
	"github.com/sipki-tech/currency"
)

// Type for money.
type Type uint8

// Type enum.
//
//go:generate stringer -output=stringer.Type.go -type=Type
const (
	_ Type = iota
	Fiat
	Crypto
)

// Money represent fiat money.
type Money struct {
	Code    currency.Code
	Decimal decimal.Decimal
	Type    Type
}
