package cmd

import (
	"ecoommerce/config"
	"ecoommerce/infra/db"
	"ecoommerce/repo"
	"ecoommerce/rest"
	"ecoommerce/rest/handlers/product"
	"ecoommerce/rest/handlers/user"
	"ecoommerce/rest/middleware"
	"os"
)

func Serve() {

	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		os.Exit(1)
	}

	err = db.MigrateDb(dbCon, "./migrations")
	if err != nil {
		os.Exit(1)
	}

	middlewares := middleware.NewMiddleWares(cnf)
	productRepo := repo.NewProductRepo(dbCon)
	productHandler := product.NewHandler(middlewares, productRepo)

	userRepo := repo.NewUserRepo(dbCon)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()
}
