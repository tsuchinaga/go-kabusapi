package main

import (
	"fmt"

	"gitlab.com/tsuchinaga/go-kabusapi/kabus"
)

func main() {
	requester := kabus.NewTokenRequester()
	res, err := requester.Exec(kabus.TokenRequest{APIPassword: "password"})
	fmt.Printf("res: %+v\nerr: %+v\n", res, err)
}
