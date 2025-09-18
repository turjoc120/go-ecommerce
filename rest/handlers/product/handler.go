package product

import "github.com/turjoc120/ecom/rest/middleware"

type Handler struct {
	middlewares *middleware.Middlewares
}

func NewHandler(middlewares *middleware.Middlewares) *Handler {
	return &Handler{
		middlewares: middlewares,
	}
}
