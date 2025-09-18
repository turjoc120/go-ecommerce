package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/turjoc120/ecom/config"
	"github.com/turjoc120/ecom/rest/handlers/product"
	"github.com/turjoc120/ecom/rest/handlers/user"
	"github.com/turjoc120/ecom/rest/middleware"
)

type Server struct {
	cnf            *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(cnf *config.Config, productHandler *product.Handler, userHandler *user.Handler) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Start() {
	mux := http.NewServeMux()

	wrappedMux := middleware.Use(mux, middleware.CorsWithPreflight, middleware.Logger)

	server.productHandler.RegisterRoutes(mux)
	server.userHandler.RegisterRoutes(mux)

	fmt.Println("Server started on :8080")

	err := http.ListenAndServe(":"+strconv.Itoa(server.cnf.HttpPort), wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
