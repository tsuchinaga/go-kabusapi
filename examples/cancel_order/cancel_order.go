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

	orderID := ""
	{
		res, err := client.SendOrderStock(token, kabus.SendOrderStockRequest{
			Password:           password,
			Symbol:             "1320",
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
			ClosePositions:     []kabus.ClosePosition{},
			Price:              0,
			ExpireDay:          kabus.YmdNUMToday,
			FrontOrderType:     kabus.StockFrontOrderTypeMOCA, // 後場引成 = 約定しないように
		})
		if err != nil {
			panic(err)
		}
		log.Println(res)
		orderID = res.OrderID
	}

	time.Sleep(3 * time.Second) // kabuステーションで注文がどう動くかを確認するために待機時間を入れてます
	{
		res, err := client.CancelOrder(token, kabus.CancelOrderRequest{
			OrderID:  orderID,
			Password: password,
		})
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}
}
