package user

import (
	"encoding/json"
	"net/http"

	"github.com/turjoc120/ecom/util"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin
	err := json.NewDecoder(r.Body).Decode(&reqLogin)
	if err != nil {
		http.Error(w, "give me a valid data", http.StatusBadRequest)
		return
	}

	usr, err := h.userRepo.Get(reqLogin.Email, reqLogin.Password)
	if err != nil {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}

	accessToken, err := util.CreateJwt(h.cnf.JwtSecretKey, util.Payload{Sub: usr.ID,
		FirstName:   usr.FirstName,
		LastName:    usr.LastName,
		Email:       usr.Email,
		IsShopOwner: true})
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	util.SendData(w, accessToken, http.StatusCreated)
}
