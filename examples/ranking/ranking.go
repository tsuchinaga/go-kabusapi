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

	// 値上がり率
	{
		res, err := client.Ranking(token, kabus.RankingRequest{Type: kabus.RankingTypePriceIncreaseRate, ExchangeDivision: kabus.ExchangeDivisionToushou})
		if err != nil {
			panic(err)
		}
		log.Printf("%+v\n", res)
	}

	// TICK回数
	{
		res, err := client.Ranking(token, kabus.RankingRequest{Type: kabus.RankingTypeTickCount, ExchangeDivision: kabus.ExchangeDivisionToushou})
		if err != nil {
			panic(err)
		}
		log.Printf("%+v\n", res)
	}

	// 売買高急増
	{
		res, err := client.Ranking(token, kabus.RankingRequest{Type: kabus.RankingTypeVolumeRapidIncrease, ExchangeDivision: kabus.ExchangeDivisionToushou})
		if err != nil {
			panic(err)
		}
		log.Printf("%+v\n", res)
	}

	// 売買代金急増
	{
		res, err := client.Ranking(token, kabus.RankingRequest{Type: kabus.RankingTypeValueRapidIncrease, ExchangeDivision: kabus.ExchangeDivisionToushou})
		if err != nil {
			panic(err)
		}
		log.Printf("%+v\n", res)
	}

	// 信用高倍率
	{
		res, err := client.Ranking(token, kabus.RankingRequest{Type: kabus.RankingTypeMarginHighMagnification, ExchangeDivision: kabus.ExchangeDivisionToushou})
		if err != nil {
			panic(err)
		}
		log.Printf("%+v\n", res)
	}

	// 業種別値上がり率
	{
		res, err := client.Ranking(token, kabus.RankingRequest{Type: kabus.RankingTypePriceIncreaseRateByCategory, ExchangeDivision: kabus.ExchangeDivisionToushou})
		if err != nil {
			panic(err)
		}
		log.Printf("%+v\n", res)
	}
}
