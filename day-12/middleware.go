package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Middleware function
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the request details
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log the processing time
		duration := time.Since(start)
		log.Printf("Completed request in %v\n", duration)
	})
}

// Main handler function
func MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	// Create the main handler
	mainHandler := http.HandlerFunc(MainHandler)

	// Wrap the main handler with the middleware
	handler := LoggingMiddleware(mainHandler)

	// Set up the server with the handler
	http.Handle("/", handler)

	// Start the server
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
