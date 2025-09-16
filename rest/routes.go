package rest

import (
	"net/http"

	"github.com/turjoc120/ecom/rest/handlers"
	"github.com/turjoc120/ecom/rest/middleware"
)

func initRoutes(mux *http.ServeMux) {
	mux.Handle("GET /test",
		middleware.Use(http.HandlerFunc(handlers.TestHandler)))

	mux.Handle("GET /products",
		middleware.Use(http.HandlerFunc(handlers.GetProducts)))

	mux.Handle("POST /products",
		middleware.Use(http.HandlerFunc(handlers.CreateProduct)))

	mux.Handle("GET /products/{id}",
		middleware.Use(http.HandlerFunc(handlers.GetProductByID)))

	mux.Handle("DELETE /products/{id}",
		middleware.Use(http.HandlerFunc(handlers.DeleteProduct)))

	mux.Handle("POST /users",
		middleware.Use(http.HandlerFunc(handlers.CreateUser)))

	mux.Handle("POST /users/login",
		middleware.Use(http.HandlerFunc(handlers.Login)))

	// mux.Handle("GET /products", middleware.Logger(http.HandlerFunc(handlers.GetProducts)))
	// mux.Handle("POST /products", middleware.Logger(http.HandlerFunc(handlers.CreateProduct)))
	// mux.Handle("GET /products/{id}", middleware.Logger(http.HandlerFunc(handlers.GetProductByID)))
}
