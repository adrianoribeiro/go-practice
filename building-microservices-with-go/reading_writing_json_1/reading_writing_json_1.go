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
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloWorldHandle(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello world man."}
	data, err := json.Marshal(response)

	if err != nil {
		panic("Oops")
	}

	fmt.Fprint(w, string(data))
}
