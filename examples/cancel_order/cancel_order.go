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

	var token string
	{
		req, err := kabus.NewTokenRequester(isProd).Exec(kabus.TokenRequest{APIPassword: password})
		if err != nil {
			panic(err)
		}
		token = req.Token
	}

	// TODO 注文処理を入れる
	orderID := ""
	{

	}

	{
		res, err := kabus.NewCancelOrderRequester(token, isProd).Exec(kabus.CancelOrderRequest{
			OrderID:  orderID,
			Password: password,
		})
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}
}
