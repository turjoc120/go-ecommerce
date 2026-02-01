package product

import (
	"ecoommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		util.SendData(w, 400, "give me valid product id")
		return
	}
	product, err := h.productRepo.Get(productId)
	if err != nil {
		util.SendData(w, 404, "product pawa jay ni")
		return
	}
	util.SendData(w, 200, product)
}
