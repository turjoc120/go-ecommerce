package rest

import (
	"ecoommerce/config"
	"ecoommerce/rest/handlers/product"
	"ecoommerce/rest/handlers/user"
	"ecoommerce/rest/middleware"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	cnf            *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(cnf *config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (s *Server) Start() {
	manager := middleware.NewManager()
	manager.Use(middleware.CorsWithPreflight)

	mux := http.NewServeMux()
	wrapedMux := manager.WrapMux(mux)

	s.productHandler.RegisterRoutes(mux, manager)
	s.userHandler.RegisterRoutes(mux, manager)

	fmt.Println("Starting server on port :", s.cnf.HttpPort)
	err := http.ListenAndServe(":"+strconv.Itoa(s.cnf.HttpPort), wrapedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
