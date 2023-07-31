package mw

import (
	"fmt"
	"net/http"
)

func Hello(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello from middleware")
		next.ServeHTTP(w, r)
	})
}
