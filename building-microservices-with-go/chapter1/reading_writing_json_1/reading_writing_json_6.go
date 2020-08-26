package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type validationContextKey string

type helloWorldResponse struct {
	Message string `json:"message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

type validationHandler struct {
	next http.Handler
}

type helloWorldHandler struct{}

func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

func main() {
	port := 8080

	handler := newValidationHandler(newHelloWorldHandler())

	http.Handle("/helloworld", handler)
	log.Printf("Server start on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad Request", http.StatusBadGateway)
		return
	}

	if len(request.Name) < 5 {
		http.Error(rw, "The name must be more than 5 characters", http.StatusBadGateway)
		return
	}

	//Context to access request-scoped data
	c := context.WithValue(r.Context(), validationContextKey("name"), request.Name)
	r = r.WithContext(c)

	h.next.ServeHTTP(rw, r)
}

func (h helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	name := r.Context().Value(validationContextKey("name")).(string)
	response := helloWorldResponse{Message: name}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}
