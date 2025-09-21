package cmd

import (
	"github.com/turjoc120/ecom/config"
	"github.com/turjoc120/ecom/repo"
	"github.com/turjoc120/ecom/rest"
	"github.com/turjoc120/ecom/rest/handlers/product"
	"github.com/turjoc120/ecom/rest/handlers/user"
	"github.com/turjoc120/ecom/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()

	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()

	middleware := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(productRepo, middleware)
	userHandler := user.NewHandler(userRepo, cnf)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()

}
