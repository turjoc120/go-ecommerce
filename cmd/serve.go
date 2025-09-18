package cmd

import (
	"github.com/turjoc120/ecom/config"
	"github.com/turjoc120/ecom/rest"
	"github.com/turjoc120/ecom/rest/handlers/product"
	"github.com/turjoc120/ecom/rest/handlers/user"
)

func Serve() {
	cnf := config.GetConfig()
	productHandler := product.NewHandler()
	userHandler := user.NewHandler()

	server := rest.NewServer(productHandler, userHandler)
	server.Start(cnf)

}
