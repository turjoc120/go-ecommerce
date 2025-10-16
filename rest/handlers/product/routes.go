package product

import (
	"net/http"

	"github.com/turjoc120/ecom/rest/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {

	mux.Handle("GET /products",
		middleware.Use(http.HandlerFunc(h.GetProducts)))

	mux.Handle("POST /products",
		middleware.Use(http.HandlerFunc(h.CreateProduct),
			h.middlewares.AuthenticateJWT,
		))

	mux.Handle("GET /products/{id}",
		middleware.Use(http.HandlerFunc(h.GetProductByID)))

	mux.Handle("DELETE /products/{id}",
		middleware.Use(http.HandlerFunc(h.DeleteProduct),
			h.middlewares.AuthenticateJWT,
		))

	mux.Handle("PATCH /products/{id}",
		middleware.Use(http.HandlerFunc(h.UpdateProduct),
			h.middlewares.AuthenticateJWT,
		))

}
