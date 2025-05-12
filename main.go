package main

import (
	"log"
	"net/http"
	"net/http/pprof"

	firebase "firebase.google.com/go/v4"
)

func main() {
	http.HandleFunc("/leak-test", leakTestHandler)
	http.HandleFunc("/debug", pprof.Index)

	// Start the HTTP server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func leakTestHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	_, err = app.Auth(ctx)
	if err != nil {
		log.Fatalf("error initializing auth client: %v\n", err)
	}

	// do something

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
