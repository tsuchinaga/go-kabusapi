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
		// フリーETFで日経225に連動する銘柄を選んだだけ
		res, err := client.Register(token, kabus.RegisterRequest{Symbols: []kabus.RegisterSymbol{
			{Symbol: "1320", Exchange: kabus.ExchangeToushou},
			{Symbol: "1329", Exchange: kabus.ExchangeToushou},
			{Symbol: "1346", Exchange: kabus.ExchangeToushou},
			{Symbol: "1369", Exchange: kabus.ExchangeToushou},
			{Symbol: "1397", Exchange: kabus.ExchangeToushou},
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
		ws := kabus.NewWSRequester(isProd)
		ws.SetOnNext(onNext)

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
