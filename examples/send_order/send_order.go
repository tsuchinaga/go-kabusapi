package main

import (
	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
	"log"
	"os"
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
		res, err := kabus.NewSendOrderRequester(token, isProd).Exec(kabus.SendOrderRequest{
			Password:           password,
			Symbol:             "1320",
			Exchange:           kabus.ExchangeToushou,
			SecurityType:       kabus.SecurityTypeKabu,
			Side:               kabus.SideBuy,
			CashMargin:         kabus.CashMarginMarginEntry,
			MarginTradeType:    kabus.MarginTradeTypeSystem,
			DelivType:          kabus.DelivTypeUnspecified,
			FundType:           kabus.FundTypeTransferMargin,
			AccountType:        kabus.AccountTypeGeneral,
			Qty:                1.0,
			ClosePositionOrder: kabus.ClosePositionOrderUnspecified,
			ClosePositions:     []kabus.ClosePosition{{}},
			Price:              0,
			ExpireDay:          kabus.YmdNUMToday,
			FrontOrderType:     kabus.FrontOrderTypeMarket,
		})
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}
}
