package main

import (
	"fmt"
	"os"

	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
)

func main() {
	password := os.Getenv("API_PASSWORD")
	requester := kabus.NewTokenRequester()
	res, err := requester.Exec(kabus.TokenRequest{APIPassword: password})
	fmt.Printf("res: %+v\nerr: %+v\n", res, err)
}
