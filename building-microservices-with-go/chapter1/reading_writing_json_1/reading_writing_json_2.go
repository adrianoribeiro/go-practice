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

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandle)
	log.Printf("Server start on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

//Using Encoder rather than marshalling to a byte array is nearly 50% faster.
func helloWorldHandle(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello world man."}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}
