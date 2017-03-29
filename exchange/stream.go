package exchange

import (
	"github.com/crackcomm/crypto-exchanges/currency"
	"github.com/crackcomm/crypto-exchanges/orderbook"
)

// Stream - Exchange stream.
type Stream interface {
	// Subscribe - Subscribe to currency pair exchanges.
	Subscribe(currency.Pair) error

	// Events - Channel of events.
	Events() <-chan orderbook.Event

	// Errors - Channel of errors.
	Errors() <-chan error

	// Close - Closes a stream.
	Close() error
}
