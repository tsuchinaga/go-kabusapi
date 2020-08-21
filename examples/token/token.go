package main

import (
	"fmt"
	"os"

	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
)

func main() {
	password := os.Getenv("API_PASSWORD")
	isProd := false
	if os.Getenv("IS_PROD") != "" {
		isProd = true
	}

	requester := kabus.NewTokenRequester(isProd)
	res, err := requester.Exec(kabus.TokenRequest{APIPassword: password})
	fmt.Printf("res: %+v\nerr: %+v\n", res, err)
}
