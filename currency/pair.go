package currency

import "fmt"

// Pair - Currency pair.
type Pair [2]Currency

// NewPair - Creates a new currency pair.
func NewPair(one, two string) Pair {
	return Pair{Currency(one), Currency(two)}
}

// One - First currency.
func (pair Pair) One() Currency {
	return pair[0]
}

// Two - Second currency.
func (pair Pair) Two() Currency {
	return pair[1]
}

// String - Currency pair separated by underscore.
func (pair Pair) String() string {
	return fmt.Sprintf("%s_%s", pair[0], pair[1])
}
