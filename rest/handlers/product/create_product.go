package product

import (
	"encoding/json"
	"net/http"

	"github.com/turjoc120/ecom/repo"
	"github.com/turjoc120/ecom/util"
)

type reqProduct struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct reqProduct
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "give me a valid data", 400)
		return
	}
	product, err := h.productRepo.Create(repo.Product{
		Name:  newProduct.Name,
		Price: newProduct.Price,
	})

	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	util.SendData(w, product, http.StatusCreated)
}
