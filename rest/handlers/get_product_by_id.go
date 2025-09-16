package handlers

import (
	"net/http"
	"strconv"

	"github.com/turjoc120/ecom/database"
	"github.com/turjoc120/ecom/util"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "give me a valid id", 400)
		return
	}

	product := database.Get(productId)
	if product == nil {
		util.SendData(w, "product not found", 404)
	}
	util.SendData(w, product, 200)
}
