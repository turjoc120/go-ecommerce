package handlers

import (
	"net/http"

	"github.com/turjoc120/ecom/database"
	"github.com/turjoc120/ecom/util"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.List(), 200)
}
