package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", Status)
	mux.HandleFunc("GET /ping", Ping)

	log.Printf("server is running on port %v", 8080)

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(`server crashed`)
	}
}

func Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"ok":      true,
		"message": "server is running",
	})
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
