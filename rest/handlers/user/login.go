package user

import (
	"ecoommerce/util"

	"encoding/json"
	"net/http"
)

type reqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqUser reqLogin
	err := json.NewDecoder(r.Body).Decode(&reqUser)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.Get(reqUser.Email, reqUser.Password)
	if err != nil {
		util.SendData(w, 404, "user not found")
		return
	}

	accessToken, err := util.CreateJwt(h.cnf.JwtSecret, util.Payload{
		Sub:   user.Username,
		Name:  user.Username,
		Admin: user.IsAdmin,
	})
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	util.SendData(w, 200, accessToken)
}
