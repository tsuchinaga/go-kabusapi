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
		res, err := client.WalletMargin(token)
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}

	{
		res, err := client.WalletMarginSymbol(token, kabus.WalletMarginSymbolRequest{Symbol: "9433", Exchange: kabus.StockExchangeToushou})
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}
}
