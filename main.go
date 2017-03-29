package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/k0kubun/pp"

	"github.com/crackcomm/crypto-exchanges/currency"
	"github.com/crackcomm/crypto-exchanges/exchange/poloniex"
)

func main() {
	flag.CommandLine.Parse([]string{"-logtostderr"})

	client, err := poloniex.NewStream()
	if err != nil {
		glog.Fatal(err)
	}
	defer client.Close()

	err = client.Subscribe(currency.Pair{"BTC", "DCR"})
	if err != nil {
		glog.Fatal(err)
	}

	for ev := range client.Events() {
		// switch ev.(type) {
		// case *orderbook.EventModify:
		// case *orderbook.EventRemove:
		// case *orderbook.EventTrade:
		// default:
		// 	panic("")
		// }
		// glog.Infof("%#v", ev)
		pp.Print(ev)
	}

	<-make(chan bool)
}
