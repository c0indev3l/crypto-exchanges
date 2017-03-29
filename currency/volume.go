package currency

// Volume - Currency volume.
type Volume struct {
	// Amount - Amount of currency.
	Amount float64 `json:"amount,omitempty"`

	// Currency - Currency symbol.
	Currency `json:"currency,omitempty"`
}
