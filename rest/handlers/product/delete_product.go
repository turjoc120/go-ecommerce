package product

import (
	"net/http"
	"strconv"

	"github.com/turjoc120/ecom/util"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	pId := r.PathValue("id")
	productId, err := strconv.Atoi(pId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = h.productRepo.Delete(productId)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	util.SendData(w, "Product deleted successfully", http.StatusOK)

}
