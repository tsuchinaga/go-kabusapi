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
		res, err := kabus.NewRegisterRequester(token, isProd).Exec(kabus.RegisterRequest{Symbols: []kabus.RegisterSymbol{{Symbol: symbol, Exchange: kabus.ExchangeAll}}})
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}

	for {
		res, err := kabus.NewBoardRequester(token, isProd).Exec(kabus.BoardRequest{Symbol: symbol, Exchange: kabus.ExchangeAll})
		if err != nil {
			panic(err)
		}
		log.Printf("now: %s, CurrentPriceTime: %s\n", time.Now(), res.CurrentPriceTime)
		<-time.After(10 * time.Second)
	}
}
