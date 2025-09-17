package main

import (
	"fmt"

	"github.com/turjoc120/ecom/util"
)

func main() {
	// cmd.Serve()
	payload := util.Payload{
		Sub:         23,
		FirstName:   "turjo",
		LastName:    "chow",
		Email:       "turjoc1@gmail.com",
		IsShopOwner: true,
	}

	jwt, err := util.CreateJwt("my-secret", payload)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jwt)
}
