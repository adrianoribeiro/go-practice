package services

import (
	"go-practice/mvc-sample/domain"
)

type usersServiceInterface interface {
	GetUser(int64) (*domain.User, error)
}

type usersService struct{}

var (
	UsersService usersServiceInterface
)

func (usersService *usersService) GetUser(userId int64) (*domain.User, error) {
	return domain.UserDao.GetUser(userId)
}
