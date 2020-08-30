package domain

import (
	"errors"
	"fmt"
)

var (
	users = map[int64]*User{
		1: {Id: 1, FirstName: "Adriano", LastName: "Ribeiro", Email: "adrianomsg@gmail.com"},
		2: {Id: 2, FirstName: "Berenice", LastName: "Bingo", Email: "bere@gmail.com"},
		5: {Id: 5, FirstName: "Eunice", LastName: "Ecoporanga", Email: "eunice2@gmail.com"},
	}
)

func GetUser(userId int64) (*User, error) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, errors.New(fmt.Sprintf("User %v was not found.", userId))
}
