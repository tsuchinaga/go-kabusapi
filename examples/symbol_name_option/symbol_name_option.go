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
		res, err := client.SymbolNameOption(token, kabus.SymbolNameOptionRequest{
			DerivMonth:  kabus.YmNUMToday,
			PutOrCall:   kabus.PutOrCallPut,
			StrikePrice: 0,
		})
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}
}
