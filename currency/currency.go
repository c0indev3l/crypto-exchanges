package currency

// Currency - Currency symbol.
type Currency string

// String - Currency symbol as a string.
func (currency Currency) String() string {
	return string(currency)
}
