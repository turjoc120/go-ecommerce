package user

import (
	"encoding/json"
	"net/http"

	"github.com/turjoc120/ecom/database"
	"github.com/turjoc120/ecom/util"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "give me a valid data", http.StatusBadRequest)
		return
	}
	user := newUser.Store()

	util.SendData(w, user, http.StatusCreated)
}
