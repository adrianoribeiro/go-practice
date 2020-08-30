package services

import (
	"go-practice/mvc-sample/domain"
)

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
