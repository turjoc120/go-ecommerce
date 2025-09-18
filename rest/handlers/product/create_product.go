package product

import (
	"encoding/json"
	"net/http"

	"github.com/turjoc120/ecom/database"
	"github.com/turjoc120/ecom/util"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "give me a valid data", 400)
		return
	}
	product := database.Store(newProduct)

	util.SendData(w, product, http.StatusCreated)
}
