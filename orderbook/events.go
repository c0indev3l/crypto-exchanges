package orderbook

import (
	"time"

	"github.com/crackcomm/crypto-exchanges/currency"
	"github.com/crackcomm/crypto-exchanges/order"
)

// Event - Event interface.
type Event interface {
	// Name - Name of an event.
	Name() string
}

// EventRemove - Remove from order book.
type EventRemove struct {
	// Rate - Price rate.
	Rate currency.Volume `json:"rate,omitempty"`

	// Type - Order type (ask or bid).
	Type order.Type `json:"type,omitempty"`
}

// Name - Name of an event ("remove").
func (event *EventRemove) Name() string {
	return "remove"
}

// EventModify - Modification in order book.
type EventModify struct {
	// Rate - Price rate.
	Rate currency.Volume `json:"rate,omitempty"`

	// Volume - New volume of rate.
	Volume currency.Volume `json:"volume,omitempty"`

	// Type - Order type (ask or bid).
	Type order.Type `json:"type,omitempty"`
}

// Name - Name of an event ("modify").
func (event *EventModify) Name() string {
	return "modify"
}

// EventTrade - New trade in an order book.
type EventTrade struct {
	// ID - ID of a trade.
	ID int64 `json:"id,omitempty"`

	// Rate - Price rate.
	Rate currency.Volume `json:"rate,omitempty"`

	// Volume - New volume of rate.
	Volume currency.Volume `json:"volume,omitempty"`

	// Type - Order type (ask or bid).
	Type order.Type `json:"type,omitempty"`

	// Time - Time of the trade.
	Time time.Time `json:"time,omitempty"`
}

// Name - Name of an event ("trade").
func (event *EventTrade) Name() string {
	return "trade"
}
