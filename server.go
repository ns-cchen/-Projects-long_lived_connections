package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		// Generate a random delay between 10 and 60 seconds
		delay := rand.Intn(41) + 10
		podID := os.Getenv("POD_ID")
		if podID == "" {
			podID = "unknown"
		}

		fmt.Printf("Pod %s will delay the response for %d seconds.\n", podID, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		response := fmt.Sprintf("Response from Pod %s, delayed by %d seconds", podID, delay)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	})

	// Read port from environment variable or default to 5000
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	fmt.Printf("Server is starting on port %s...\n", port)
	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}
