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

	var token string
	{
		req, err := kabus.NewTokenRequester(isProd).Exec(kabus.TokenRequest{APIPassword: password})
		if err != nil {
			panic(err)
		}
		token = req.Token
	}

	var symbol string
	{
		res, err := kabus.NewSymbolNameFutureRequester(token, isProd).Exec(kabus.SymbolNameFutureRequest{
			FutureCode: kabus.FutureCodeNK225Mini,
			DerivMonth: kabus.YmNUMToday,
		})
		if err != nil {
			panic(err)
		}
		symbol = res.Symbol
	}

	{
		res, err := kabus.NewSendOrderFutureRequester(token, isProd).Exec(kabus.SendOrderFutureRequest{
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
