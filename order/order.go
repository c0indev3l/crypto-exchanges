package order

import "github.com/crackcomm/crypto-exchanges/currency"

// Order - Currency order.
type Order struct {
	// Volume - Volume of order.
	Volume currency.Volume `json:"volume,omitempty"`

	// Rate - Price rate.
	Rate currency.Volume `json:"rate,omitempty"`

	// Type - Order type (ask or bid).
	Type Type `json:"type,omitempty"`
}
