package product

import (
	"ecoommerce/repo"
	"ecoommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "give me a valid product id", 400)
	}

	var req reqProduct
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "give me a valid data", 400)
		return
	}

	_, err = h.productRepo.Update(repo.Product{
		ID:    productId,
		Name:  req.Name,
		Price: req.Price,
	})

	if err != nil {
		util.SendData(w, http.StatusInternalServerError, "Product updated successfully")
	}
	util.SendData(w, 200, "Product updated successfully")
}
