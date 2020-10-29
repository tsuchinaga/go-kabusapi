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

	{
		// フリーETFで日経225に連動する銘柄を選んだだけ
		res, err := kabus.NewRegisterRequester(token, isProd).Exec(kabus.RegisterRequest{Symbols: []kabus.RegisterSymbol{
			{Symbol: "1320", Exchange: kabus.StockExchangeToushou},
			{Symbol: "1329", Exchange: kabus.StockExchangeToushou},
			{Symbol: "1346", Exchange: kabus.StockExchangeToushou},
			{Symbol: "1369", Exchange: kabus.StockExchangeToushou},
			{Symbol: "1397", Exchange: kabus.StockExchangeToushou},
		}})
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}

	{
		onNext := func(msg kabus.PriceMessage) error {
			log.Printf("%+v\n", msg)
			return nil
		}
		ws := kabus.NewWSRequester(isProd, onNext)

		// 5s後にwsをCloseする
		go func() {
			<-time.After(5 * time.Second)
			if err := ws.Close(); err != nil {
				panic(err)
			}
		}()

		if err := ws.Open(); err != nil {
			panic(err)
		}
	}
}
