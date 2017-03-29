package poloniex

import (
	"fmt"
	"time"

	"github.com/crackcomm/crypto-exchanges/common"
	"github.com/crackcomm/crypto-exchanges/currency"
	"github.com/crackcomm/crypto-exchanges/order"
	"github.com/crackcomm/crypto-exchanges/orderbook"
)

var handlers = map[string]func(currency.Pair, map[string]interface{}) (orderbook.Event, error){
	"orderBookRemove": parseRemove,
	"orderBookModify": parseModify,
	"newTrade":        parseTrade,
}

func parseEvent(pair currency.Pair, typ string, data map[string]interface{}) (_ orderbook.Event, err error) {
	// Get handler by message type
	handler, ok := handlers[typ]
	if !ok {
		return nil, fmt.Errorf("unknown type: %q", typ)
	}

	// Handle message
	return handler(pair, data)
}

func parseRemove(pair currency.Pair, data map[string]interface{}) (_ orderbook.Event, err error) {
	res := new(orderbook.EventRemove)
	res.Type, err = order.TypeFromString(data["type"].(string))
	if err != nil {
		return
	}
	res.Rate, err = common.ParseIVolume(currency.Currency("BTC"), data["rate"])
	if err != nil {
		return
	}
	return res, nil
}

func parseModify(pair currency.Pair, data map[string]interface{}) (_ orderbook.Event, err error) {
	res := new(orderbook.EventModify)
	res.Type, err = order.TypeFromString(data["type"].(string))
	if err != nil {
		return
	}
	res.Rate, err = common.ParseIVolume(pair.One(), data["rate"])
	if err != nil {
		return
	}
	res.Volume, err = common.ParseIVolume(pair.Two(), data["amount"])
	if err != nil {
		return
	}
	return res, nil
}

func parseTrade(pair currency.Pair, data map[string]interface{}) (_ orderbook.Event, err error) {
	res := new(orderbook.EventTrade)
	res.ID, err = common.ParseIInt64(data["tradeID"])
	if err != nil {
		return
	}
	res.Type, err = order.TypeFromString(data["type"].(string))
	if err != nil {
		return
	}
	res.Rate, err = common.ParseIVolume(pair.One(), data["rate"])
	if err != nil {
		return
	}
	res.Volume, err = common.ParseIVolume(pair.Two(), data["amount"])
	if err != nil {
		return
	}
	res.Time, err = time.Parse("2006-01-02 15:04:05", data["date"].(string))
	if err != nil {
		return
	}
	return res, nil
}
