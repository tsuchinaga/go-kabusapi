package main

import (
	"log"
	"os"

	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

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

	// NK225Miniを登録
	symbolCode := ""
	{
		res, err := client.SymbolNameFuture(token, kabus.SymbolNameFutureRequest{FutureCode: kabus.FutureCodeNK225Mini, DerivMonth: kabus.YmNUMToday})
		if err != nil {
			panic(err)
		}
		symbolCode = res.Symbol
	}

	{
		res, err := client.Register(token, kabus.RegisterRequest{Symbols: []kabus.RegisterSymbol{{Symbol: symbolCode, Exchange: kabus.ExchangeAll}}})
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}

	{
		ws := kabus.NewWSRequester(isProd)
		ws.SetOnNext(func(msg kabus.PriceMessage) error {
			log.Printf("%+v\n", msg)
			log.Println(ws.IsOpened())
			_ = ws.Close()
			return nil
		})
		for i := 0; i < 5; i++ {
			if err := ws.Open(); err != nil {
				panic(err)
			}
			log.Println(ws.IsOpened())
		}
	}
}
