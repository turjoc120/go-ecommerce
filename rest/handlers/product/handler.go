package product

import (
	"github.com/turjoc120/ecom/repo"
	"github.com/turjoc120/ecom/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	productRepo repo.ProductRepo
}

func NewHandler(productRepo repo.ProductRepo, middlewares *middleware.Middlewares) *Handler {
	return &Handler{
		middlewares: middlewares,
		productRepo: productRepo,
	}
}
