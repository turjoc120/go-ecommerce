package cmd

import (
	"fmt"
	"os"

	"github.com/turjoc120/ecom/config"
	"github.com/turjoc120/ecom/infra/db"
	"github.com/turjoc120/ecom/repo"
	"github.com/turjoc120/ecom/rest"
	"github.com/turjoc120/ecom/rest/handlers/product"
	"github.com/turjoc120/ecom/rest/handlers/user"
	"github.com/turjoc120/ecom/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()

	db, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	productRepo := repo.NewProductRepo(db)
	userRepo := repo.NewUserRepo(db)

	middleware := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(productRepo, middleware)
	userHandler := user.NewHandler(userRepo, cnf)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()

}
