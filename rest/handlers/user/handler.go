package user

import (
	"github.com/turjoc120/ecom/config"
	"github.com/turjoc120/ecom/repo"
)

type Handler struct {
	userRepo repo.UserRepo
	cnf      *config.Config
}

func NewHandler(userRepo repo.UserRepo, cnf *config.Config) *Handler {
	return &Handler{
		userRepo: userRepo,
		cnf:      cnf,
	}
}
