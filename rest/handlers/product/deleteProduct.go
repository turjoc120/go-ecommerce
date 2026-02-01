package product

import (
	"ecoommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	pId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		util.SendData(w, 400, "give me a valid  id")
		return
	}

	err = h.productRepo.Delete(pId)
	if err != nil {
		util.SendData(w, http.StatusInternalServerError, "product not found")
		return
	}

	util.SendData(w, 200, "Product deleted successfully")
}
