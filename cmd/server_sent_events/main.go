package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/sse", sseHandler)
	http.ListenAndServe(":8080", nil)
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	// Set the response headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Create a new event stream
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-r.Context().Done():
			// The client connection is closed
			fmt.Println("Client connection closed.")
			return
		case <-ticker.C:
			// Send an SSE event to the client
			event := fmt.Sprintf("data: %s\n\n", time.Now().Format("2006-01-02 15:04:05"))
			_, err := w.Write([]byte(event))
			if err != nil {
				fmt.Println("Error writing SSE event:", err)
				return
			}
			w.(http.Flusher).Flush()
		}
	}
}
