package order

import (
	"fmt"
)

// Type - Order type.
type Type uint8

const (
	// Ask - Order ask.
	Ask Type = 0

	// Bid - Order bid.
	Bid Type = 1
)

// TypeFromString - Converts from string to Type.
func TypeFromString(str string) (_ Type, err error) {
	switch str {
	case "ask", "sell":
		return Ask, nil
	case "bid", "buy":
		return Bid, nil
	}
	err = fmt.Errorf("invalid order type %q", str)
	return
}
