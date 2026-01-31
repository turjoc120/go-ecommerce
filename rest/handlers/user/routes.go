package user

import (
	"ecoommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("POST /users", manager.With(
		http.HandlerFunc(h.CreateUser),
	))
	mux.Handle("POST /login", manager.With(
		http.HandlerFunc(h.Login),
	))
}
