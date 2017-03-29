package poloniex

import (
	"fmt"

	"github.com/beatgammit/turnpike"

	"github.com/crackcomm/crypto-exchanges/currency"
	"github.com/crackcomm/crypto-exchanges/orderbook"
)

// Stream - Poloniex stream.
type Stream struct {
	client *turnpike.Client
	events chan orderbook.Event
	errors chan error
}

const (
	// WebsocketAddress - Poloniex Websocket address.
	WebsocketAddress = "wss://api.poloniex.com"

	// WebsocketRealm - Poloniex Websocket realm name.
	WebsocketRealm = "realm1"
)

// NewStream - Creates a new connected poloniex stream.
func NewStream() (stream *Stream, err error) {
	client, err := turnpike.NewWebsocketClient(turnpike.JSON, WebsocketAddress, nil)
	if err != nil {
		return
	}

	_, err = client.JoinRealm(WebsocketRealm, nil)
	if err != nil {
		return
	}

	return &Stream{
		client: client,
		events: make(chan orderbook.Event, 1000),
		errors: make(chan error, 10),
	}, nil
}

// Subscribe - Subscribes to currency pair order book.
func (stream *Stream) Subscribe(pair currency.Pair) error {
	return stream.client.Subscribe(pair.String(), stream.eventHandler(pair))
}

// Events - Channel of events.
func (stream *Stream) Events() <-chan orderbook.Event {
	return stream.events
}

// Errors - Channel of errors.
func (stream *Stream) Errors() <-chan error {
	return stream.errors
}

// Close - Closes a stream.
func (stream *Stream) Close() error {
	return stream.client.Close()
}

// eventHandler - Creates stream event handler for currency pair.
func (stream *Stream) eventHandler(pair currency.Pair) turnpike.EventHandler {
	return func(args []interface{}, kwargs map[string]interface{}) {
		for _, v := range args {
			// Cast to underlying value type
			value := v.(map[string]interface{})

			// Message type
			typ := value["type"].(string)

			// Message data
			data := value["data"].(map[string]interface{})

			// Handle message
			res, err := parseEvent(pair, typ, data)
			if err != nil {
				stream.errors <- fmt.Errorf("unhandled poloniex message: %v", err)
			} else {
				stream.events <- res
			}
		}
	}
}
