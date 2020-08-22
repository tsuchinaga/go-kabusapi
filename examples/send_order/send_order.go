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
		res, err := kabus.NewSendOrderRequester(token, isProd).Exec(kabus.SendOrderRequest{
			Password:           password,
			Symbol:             "1320",
			Exchange:           kabus.ExchangeToushou,
			SecurityType:       kabus.SecurityTypeKabu,
			Side:               kabus.SideBuy,
			CashMargin:         kabus.CashMarginCash,
			MarginTradeType:    kabus.MarginTradeTypeUnspecified,
			DelivType:          kabus.DelivTypeCash,
			FundType:           kabus.FundTypeTransferMargin,
			AccountType:        kabus.AccountTypeGeneral,
			Qty:                1.0,
			ClosePositionOrder: kabus.ClosePositionOrderUnspecified,
			ClosePositions:     []kabus.ClosePosition{},
			Price:              0,
			ExpireDay:          kabus.YmdNUM{Time: time.Date(2020, 8, 24, 0, 0, 0, 0, time.Local)},
			FrontOrderType:     kabus.FrontOrderTypeMarket,
		})
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}
}
