package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/turjoc120/ecom/database"
	"github.com/turjoc120/ecom/util"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin
	err := json.NewDecoder(r.Body).Decode(&reqLogin)
	if err != nil {
		http.Error(w, "give me a valid data", http.StatusBadRequest)
		return
	}

	usr := database.Find(reqLogin.Email, reqLogin.Password)
	if usr == nil {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}

	util.SendData(w, usr, http.StatusCreated)
}
