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

	{
		res, err := client.UnregisterAll(token)
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}

	ws := kabus.NewWSRequester(isProd)
	{
		ws.SetOnNext(func(msg kabus.PriceMessage) error {
			log.Printf("%+v\n", msg)
			return nil
		})

		go func() {
			if err := ws.Open(); err != nil {
				panic(err)
			}
		}()
	}

	registerSymbols := []kabus.RegisterSymbol{
		{Symbol: "1320", Exchange: kabus.ExchangeToushou},
		{Symbol: "1329", Exchange: kabus.ExchangeToushou},
		{Symbol: "1346", Exchange: kabus.ExchangeToushou},
		{Symbol: "1369", Exchange: kabus.ExchangeToushou},
		{Symbol: "1397", Exchange: kabus.ExchangeToushou},
	}

	for _, s := range registerSymbols {
		res, err := client.Register(token, kabus.RegisterRequest{Symbols: []kabus.RegisterSymbol{s}})
		if err != nil {
			panic(err)
		}
		log.Println(res)

		<-time.After(10 * time.Second)
	}

	if err := ws.Close(); err != nil {
		panic(err)
	}
}
