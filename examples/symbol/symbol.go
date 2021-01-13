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

	{
		res, err := client.Symbol(token, kabus.SymbolRequest{Symbol: "9433", Exchange: kabus.ExchangeToushou})
		if err != nil {
			panic(err)
		}
		log.Printf("%+v\n", res)
	}

	// 先物
	futureCode := ""
	{
		res, err := client.SymbolNameFuture(token, kabus.SymbolNameFutureRequest{
			FutureCode: kabus.FutureCodeNK225Mini,
			DerivMonth: kabus.YmNUMToday,
		})
		if err != nil {
			panic(err)
		}
		futureCode = res.Symbol
	}

	{
		res, err := client.Symbol(token, kabus.SymbolRequest{Symbol: futureCode, Exchange: kabus.ExchangeAll})
		if err != nil {
			panic(err)
		}
		log.Printf("%+v\n", res)
	}

	// オプション
	optionCode := ""
	{
		res, err := client.SymbolNameOption(token, kabus.SymbolNameOptionRequest{
			DerivMonth:  kabus.YmNUMToday,
			PutOrCall:   kabus.PutOrCallPut,
			StrikePrice: 0,
		})
		if err != nil {
			panic(err)
		}
		optionCode = res.Symbol
	}

	{
		res, err := client.Symbol(token, kabus.SymbolRequest{Symbol: optionCode, Exchange: kabus.ExchangeAll})
		if err != nil {
			panic(err)
		}
		log.Printf("%+v\n", res)
	}
}
