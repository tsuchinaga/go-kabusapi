package main

import (
	"log"
	"os"

	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
)

func main() {
	password := os.Getenv("API_PASSWORD")
	var token string
	{
		req, err := kabus.NewTokenRequester().Exec(kabus.TokenRequest{APIPassword: password})
		if err != nil {
			panic(err)
		}
		token = req.Token
	}

	// TODO エラーで動かない https://github.com/kabucom/kabusapi/issues/14
	{
		res, err := kabus.NewPositionsRequester(token).Exec(kabus.PositionsRequest{Product: kabus.ProductAll})
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}
}
