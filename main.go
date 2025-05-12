package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/leak-test", leakTestHandler)

	// Start the HTTP server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func leakTestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
