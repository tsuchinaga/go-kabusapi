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
		res, err := client.SymbolNameOption(token, kabus.SymbolNameOptionRequest{
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
		res, err := client.SendOrderOption(token, kabus.SendOrderOptionRequest{
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
