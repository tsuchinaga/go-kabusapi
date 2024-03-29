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
	kabucomPassword := os.Getenv("KABUCOM_PASSWORD")

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
		res, err := client.SendOrderStock(token, kabus.SendOrderStockRequest{
			Password:           kabucomPassword,
			Symbol:             "1475",
			Exchange:           kabus.StockExchangeToushou,
			SecurityType:       kabus.SecurityTypeStock,
			Side:               kabus.SideBuy,
			CashMargin:         kabus.CashMarginCash,
			MarginTradeType:    kabus.MarginTradeTypeUnspecified,
			DelivType:          kabus.DelivTypeCash,
			FundType:           kabus.FundTypeTransferMargin,
			AccountType:        kabus.AccountTypeGeneral,
			Qty:                1.0,
			ClosePositionOrder: kabus.ClosePositionOrderUnspecified,
			ClosePositions:     []kabus.ClosePosition{{}},
			FrontOrderType:     kabus.StockFrontOrderTypeReverseLimit,
			Price:              0,
			ExpireDay:          kabus.YmdNUMToday,
			ReverseLimitOrder: &kabus.StockReverseLimitOrder{
				TriggerSec:        kabus.TriggerSecOrderSymbol,
				TriggerPrice:      1970,
				UnderOver:         kabus.UnderOverUnder,
				AfterHitOrderType: kabus.StockAfterHitOrderTypeMarket,
			},
		})
		if err != nil {
			panic(err)
		}
		log.Printf("%+v\n", res)
	}

	{
		res, err := client.Orders(token, kabus.OrdersRequest{
			Product:          kabus.ProductAll,
			IsGetOrderDetail: kabus.IsGetOrderDetailTrue})
		if err != nil {
			panic(err)
		}
		log.Printf("%+v\n", res)
	}
}
