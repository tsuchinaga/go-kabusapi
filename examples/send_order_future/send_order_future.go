package main

import (
	"log"
	"os"

	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
)

func main() {
	password := os.Getenv("API_PASSWORD")
	isProd := false
	if os.Getenv("IS_PROD") != "" {
		isProd = true
	}

	client := kabus.NewRESTClient(isProd)

	var token string
	{
		req, err := client.Token(kabus.TokenRequest{APIPassword: password})
		if err != nil {
			panic(err)
		}
		token = req.Token
	}

	var symbol string
	{
		res, err := client.SymbolNameFuture(token, kabus.SymbolNameFutureRequest{
			FutureCode: kabus.FutureCodeNK225Mini,
			DerivMonth: kabus.YmNUMToday,
		})
		if err != nil {
			panic(err)
		}
		symbol = res.Symbol
	}

	{
		res, err := client.SendOrderFuture(token, kabus.SendOrderFutureRequest{
			Password:       password,
			Symbol:         symbol,
			Exchange:       kabus.FutureExchangeEvening,
			TradeType:      kabus.TradeTypeEntry,
			TimeInForce:    kabus.TimeInForceFAK,
			Side:           kabus.SideBuy,
			Qty:            1,
			FrontOrderType: kabus.FutureFrontOrderTypeMarket,
			ExpireDay:      kabus.YmdNUMToday,
		})
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}
}
