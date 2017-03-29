package poloniex

import (
	"testing"

	"github.com/crackcomm/crypto-exchanges/currency"
	"github.com/crackcomm/crypto-exchanges/order"
	"github.com/crackcomm/crypto-exchanges/orderbook"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type ParseSuite struct{}

var _ = Suite(&ParseSuite{})

func (s *ParseSuite) TestParseRemove(c *C) {
	e, err := parseEvent(currency.NewPair("BTC", "DCR"), "orderBookRemove", map[string]interface{}{
		"type": "ask",
		"rate": "0.01227622",
	})
	c.Assert(err, Equals, nil)

	ev := e.(*orderbook.EventRemove)
	c.Assert(ev.Type, Equals, order.Ask)
	c.Assert(ev.Rate.Amount, Equals, 0.01227622)
	c.Assert(ev.Rate.Currency, Equals, currency.Currency("BTC"))
}

func (s *ParseSuite) TestParseModify(c *C) {
	e, err := parseEvent(currency.NewPair("BTC", "DCR"), "orderBookModify", map[string]interface{}{
		"type":   "ask",
		"rate":   "0.01300000",
		"amount": "1147.68249538",
	})
	c.Assert(err, Equals, nil)

	ev := e.(*orderbook.EventModify)
	c.Assert(ev.Type, Equals, order.Ask)
	c.Assert(ev.Rate.Amount, Equals, 0.01300000)
	c.Assert(ev.Rate.Currency, Equals, currency.Currency("BTC"))
	c.Assert(ev.Volume.Amount, Equals, 1147.68249538)
	c.Assert(ev.Volume.Currency, Equals, currency.Currency("DCR"))
}

func (s *ParseSuite) TestParseTrade(c *C) {
	e, err := parseEvent(currency.NewPair("BTC", "DCR"), "newTrade", map[string]interface{}{
		"amount":  "4.15417989",
		"date":    "2017-03-29 05:19:25",
		"rate":    "0.01203607",
		"total":   "0.04999999",
		"tradeID": "587591",
		"type":    "sell",
	})
	c.Assert(err, Equals, nil)

	ev := e.(*orderbook.EventTrade)
	c.Assert(ev.ID, Equals, int64(587591))
	c.Assert(ev.Type, Equals, order.Ask)
	c.Assert(ev.Rate.Amount, Equals, 0.01203607)
	c.Assert(ev.Rate.Currency, Equals, currency.Currency("BTC"))
	c.Assert(ev.Volume.Amount, Equals, 4.15417989)
	c.Assert(ev.Volume.Currency, Equals, currency.Currency("DCR"))
}
