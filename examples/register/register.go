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

	symbols := make([]kabus.RegisterSymbol, 0)
	symbols = append(symbols, kabus.RegisterSymbol{Symbol: "9433", Exchange: kabus.ExchangeToushou})
	{
		res, err := client.SymbolNameFuture(token, kabus.SymbolNameFutureRequest{
			FutureCode: kabus.FutureCodeNK225,
			DerivMonth: kabus.YmNUMToday,
		})
		if err != nil {
			panic(err)
		}
		symbols = append(symbols, kabus.RegisterSymbol{Symbol: res.Symbol, Exchange: kabus.ExchangeDaytime})
	}

	{
		res, err := client.Register(token, kabus.RegisterRequest{Symbols: symbols})
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}
}
