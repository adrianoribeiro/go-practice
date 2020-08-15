package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandle)
	log.Printf("Server start on port %d", port)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloWorldHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world\n")
}
