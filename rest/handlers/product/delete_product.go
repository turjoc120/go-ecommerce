package product

import (
	"net/http"
	"strconv"

	"github.com/turjoc120/ecom/database"
	"github.com/turjoc120/ecom/util"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	pId := r.PathValue("id")
	productId, err := strconv.Atoi(pId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	product := database.Get(productId)
	if product == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	database.Delete(productId)
	util.SendData(w, "Product deleted successfully", 201)

}
