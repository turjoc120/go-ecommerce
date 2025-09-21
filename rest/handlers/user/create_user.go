package user

import (
	"encoding/json"
	"net/http"

	"github.com/turjoc120/ecom/repo"
	"github.com/turjoc120/ecom/util"
)

type reqUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var reqUser reqUser
	err := json.NewDecoder(r.Body).Decode(&reqUser)
	if err != nil {
		http.Error(w, "give me a valid data", http.StatusBadRequest)
		return
	}
	user, err := h.userRepo.Create(repo.User{
		FirstName:   reqUser.FirstName,
		LastName:    reqUser.LastName,
		Email:       reqUser.Email,
		Password:    reqUser.Password,
		IsShopOwner: reqUser.IsShopOwner,
	})

	util.SendData(w, user, http.StatusCreated)
}
