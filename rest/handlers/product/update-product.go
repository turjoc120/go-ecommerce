package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/turjoc120/ecom/repo"
	"github.com/turjoc120/ecom/util"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "give me a valid id", 400)
		return
	}
	var updatedProduct reqProduct
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "give me a valid data", 400)
		return
	}
	product, err := h.productRepo.Update(repo.Product{
		ID:    productId,
		Name:  updatedProduct.Name,
		Price: updatedProduct.Price,
	})

	if err != nil {
		http.Error(w, "internal server erro", http.StatusInternalServerError)
	}
	util.SendData(w, product, http.StatusOK)
}
