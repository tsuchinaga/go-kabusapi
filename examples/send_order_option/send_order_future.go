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
		res, err := kabus.NewSymbolNameOptionRequester(token, isProd).Exec(kabus.SymbolNameOptionRequest{
			DerivMonth:  kabus.YmNUMToday,
			PutOrCall:   kabus.PutOrCallPut,
			StrikePrice: 0,
		})
		if err != nil {
			panic(err)
		}
		symbol = res.Symbol
	}

	{
		res, err := kabus.NewSendOrderOptionRequester(token, isProd).Exec(kabus.SendOrderOptionRequest{
			Password:       password,
			Symbol:         symbol,
			Exchange:       kabus.OptionExchangeEvening,
			TradeType:      kabus.TradeTypeEntry,
			TimeInForce:    kabus.TimeInForceFAK,
			Side:           kabus.SideBuy,
			Qty:            1,
			FrontOrderType: kabus.OptionFrontOrderTypeMarket,
			ExpireDay:      kabus.YmdNUMToday,
		})
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}
}
