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

	{
		res, err := kabus.NewRegisterRequester(token, isProd).Exec(kabus.RegisterRequest{Symbols: []kabus.RegistSymbol{{Symbol: "9433", Exchange: kabus.ExchangeToushou}}})
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}

	{
		res, err := kabus.NewUnregisterAllRequester(token, isProd).Exec()
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}
}
