package user

import (
	"net/http"

	"github.com/turjoc120/ecom/rest/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {

	mux.Handle("POST /users",
		middleware.Use(http.HandlerFunc(h.CreateUser)))

	mux.Handle("POST /users/login",
		middleware.Use(http.HandlerFunc(h.Login)))

}
