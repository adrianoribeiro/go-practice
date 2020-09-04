package domain

import (
	"errors"
	"fmt"
	"log"
)

type userDaoInterface interface {
	GetUser(int64) (*User, error)
}

type userDao struct{}

var (
	users = map[int64]*User{
		1: {Id: 1, FirstName: "Adriano", LastName: "Ribeiro", Email: "adrianomsg@gmail.com"},
		2: {Id: 2, FirstName: "Berenice", LastName: "Bingo", Email: "bere@gmail.com"},
		5: {Id: 5, FirstName: "Eunice", LastName: "Ecoporanga", Email: "eunice2@gmail.com"},
	}

	UserDao userDaoInterface
)

func (userDao *userDao) GetUser(userId int64) (*User, error) {
	log.Println("We're access the database")
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, errors.New(fmt.Sprintf("User %v was not found.", userId))
}
