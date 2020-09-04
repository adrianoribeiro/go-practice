package controllers

import (
	"encoding/json"
	"go-practice/mvc-sample/services"
	"net/http"
	"strconv"
)

func GetUser(respWriter http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {
		respWriter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		respWriter.WriteHeader(http.StatusBadRequest)
		respWriter.Write([]byte("User id must be a number."))
		return
	}

	user, err := services.UsersService.GetUser(userID)

	if err != nil {
		respWriter.WriteHeader(http.StatusNotFound)
		respWriter.Write([]byte(err.Error()))
		return
	}

	respWriter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(respWriter)
	encoder.Encode(&user)
}
