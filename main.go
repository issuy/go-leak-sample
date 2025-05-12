package main

import (
	"log"
	"net/http"
)

func main() {
	// Define the health check handler
	http.HandleFunc("/health", healthCheckHandler)

	// Start the HTTP server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// healthCheckHandler responds with a 200 OK status code to indicate that the service is healthy
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
