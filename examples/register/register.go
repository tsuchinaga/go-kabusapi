package main

import (
	"log"

	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
)

func main() {
	var token string
	{
		req, err := kabus.NewTokenRequester().Exec(kabus.TokenRequest{APIPassword: "Password"})
		if err != nil {
			panic(err)
		}
		token = req.Token
	}

	{
		res, err := kabus.NewRegisterRequester().Exec(token, kabus.RegisterRequest{Symbols: []kabus.RegistSymbol{{Symbol: "9433", Exchange: kabus.ExchangeToushou}}})
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}
}
