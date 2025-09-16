package middleware

import (
	"fmt"
	"net/http"
)

func Mid2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Mid2 before")
		next.ServeHTTP(w, r)
		fmt.Println("Mid2 after")
	})
}
func Mid3(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Mid3............. before")
		next.ServeHTTP(w, r)
		fmt.Println("Mid3............. after")
	})
}
