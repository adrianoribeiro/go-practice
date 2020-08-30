package app

import (
	"fmt"
	"go-practice/mvc-sample/controllers"
	"log"
	"net/http"
)

func StartApp() {

	http.HandleFunc("/users", controllers.GetUser)

	port := 8080
	log.Printf("Server start on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
