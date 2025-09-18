package cmd

import (
	"github.com/turjoc120/ecom/config"
	"github.com/turjoc120/ecom/rest"
	"github.com/turjoc120/ecom/rest/handlers/product"
	"github.com/turjoc120/ecom/rest/handlers/user"
	"github.com/turjoc120/ecom/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()
	middleware := middleware.NewMiddlewares(cnf)
	productHandler := product.NewHandler(middleware)
	userHandler := user.NewHandler()

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()

}
