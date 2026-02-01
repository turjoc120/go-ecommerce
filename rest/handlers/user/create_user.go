package user

import (
	"ecoommerce/repo"
	"ecoommerce/util"
	"encoding/json"
	"net/http"
)

type reqUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser reqUser
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	createdUser, err := h.userRepo.Create(repo.User{
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: newUser.Password,
		IsAdmin:  newUser.IsAdmin,
	})

	if err != nil {
		util.SendData(w, http.StatusInternalServerError, "unable to create the user")
		return
	}
	util.SendData(w, 201, createdUser)
}
