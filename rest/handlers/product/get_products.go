package product

import (
	"net/http"

	"github.com/turjoc120/ecom/util"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList, err := h.productRepo.List()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	util.SendData(w, productList, http.StatusOK)
}
