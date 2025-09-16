package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Println("before logger")
		next.ServeHTTP(w, r)
		diff := time.Since(start)
		fmt.Printf("%s %s %s\n", r.Method, r.RequestURI, diff)
	})
}
