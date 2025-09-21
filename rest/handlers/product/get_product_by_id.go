package product

import (
	"net/http"
	"strconv"

	"github.com/turjoc120/ecom/util"
)

func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "give me a valid id", 400)
		return
	}

	product, err := h.productRepo.Get(productId)
	if err != nil {
		util.SendData(w, "product not found", http.StatusInternalServerError)
	}
	util.SendData(w, product, http.StatusOK)
}
