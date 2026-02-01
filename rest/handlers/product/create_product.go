package product

import (
	"ecoommerce/repo"
	"ecoommerce/util"
	"encoding/json"
	"net/http"
)

type reqProduct struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (h *Handler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct reqProduct
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	createdProduct, err := h.productRepo.Create(repo.Product{
		Name:  newProduct.Name,
		Price: newProduct.Price,
	})

	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	util.SendData(w, 200, createdProduct)
}
