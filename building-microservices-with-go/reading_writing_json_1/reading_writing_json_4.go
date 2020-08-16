package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandle)
	log.Printf("Server start on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func helloWorldHandle(w http.ResponseWriter, r *http.Request) {

	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: request.Name}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}
