package middleware

import (
	"fmt"
	"net/http"
)

func Mid1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Mid1 before")
		next.ServeHTTP(w, r)
		fmt.Println("Mid1 after")
	})
}
