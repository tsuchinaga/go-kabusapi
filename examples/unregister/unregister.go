package main

import (
	"log"
	"os"

	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
)

func main() {
	password := os.Getenv("API_PASSWORD")
	var token string
	{
		req, err := kabus.NewTokenRequester().Exec(kabus.TokenRequest{APIPassword: password})
		if err != nil {
			panic(err)
		}
		token = req.Token
	}

	{
		res, err := kabus.NewRegisterRequester(token).Exec(kabus.RegisterRequest{Symbols: []kabus.RegistSymbol{{Symbol: "9433", Exchange: kabus.ExchangeToushou}}})
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}

	{
		res, err := kabus.NewUnregisterRequester(token).Exec(kabus.UnregisterRequest{Symbols: []kabus.UnregistSymbol{{Symbol: "9433", Exchange: kabus.ExchangeToushou}}})
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}
}
