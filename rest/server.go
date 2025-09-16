package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/turjoc120/ecom/config"
	"github.com/turjoc120/ecom/rest/middleware"
)

func Start(cnf config.Config) {
	mux := http.NewServeMux()

	wrappedMux := middleware.Use(mux, middleware.CorsWithPreflight, middleware.Logger)

	initRoutes(mux)

	fmt.Println("Server started on :8080")

	err := http.ListenAndServe(":"+strconv.Itoa(cnf.HttpPort), wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
