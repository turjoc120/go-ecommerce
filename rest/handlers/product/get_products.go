package product

import (
	"ecoommerce/util"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList, err := h.productRepo.List()
	if err != nil {
		util.SendData(w, 200, "no products found")
		return
	}
	util.SendData(w, 200, productList)
}
