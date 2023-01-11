package money

import (
	"errors"
)

var ErrUnknownType = errors.New("unknown type")

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

func NewType(t string) (Type, error) {
	switch t {
	case Fiat.String():
		return Fiat, nil
	case Crypto.String():
		return Crypto, nil
	default:
		return 0, ErrUnknownType
	}
}
