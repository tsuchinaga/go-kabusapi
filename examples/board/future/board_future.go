package main

import (
	"log"
	"os"
	"time"

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
		res, err := client.Register(token, kabus.RegisterRequest{Symbols: []kabus.RegisterSymbol{{Symbol: symbol, Exchange: kabus.ExchangeAll}}})
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}

	for {
		res, err := client.Board(token, kabus.BoardRequest{Symbol: symbol, Exchange: kabus.ExchangeAll})
		if err != nil {
			panic(err)
		}
		log.Printf("now: %s, CurrentPriceTime: %s\n", time.Now(), res.CurrentPriceTime)
		<-time.After(10 * time.Second)
	}
}
