package main

import (
	"fmt"
	"log"
	"net/http"
)

func logMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("REQUEST: %s %s \n", r.URL.Path, r.Method)
		next.ServeHTTP(w, r)
	})
}

func isAuthenticated(r *http.Request) bool {
	token := r.Header.Get("Authorization")
	validToken := "secret???token=12345"

	return token == validToken
}

func authMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !isAuthenticated(r) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	router := http.NewServeMux()

	handler := logMiddleWare(authMiddleWare(http.HandlerFunc(mainHandler)))

	router.Handle("/", handler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
