package services

import (
	"errors"
	"fmt"
	"go-practice/mvc-sample/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	userDaoMock usersDaoMock
	userAdriano = domain.User{
		Id:        1,
		FirstName: "Adriano",
		LastName:  "Ribeiro",
		Email:     "adrianomsg@gmail.com",
	}
)

type usersDaoMock struct{}

func TestGetUserNoUserFound(t *testing.T) {
	user, err := userDaoMock.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestGetUserFound(t *testing.T) {
	user, err := userDaoMock.GetUser(1)
	assert.EqualValues(t, userAdriano.Id, user.Id)
	assert.EqualValues(t, userAdriano.FirstName, user.FirstName)
	assert.EqualValues(t, userAdriano.LastName, user.LastName)
	assert.EqualValues(t, userAdriano.Email, user.Email)
	assert.Nil(t, err)
}

func (u *usersDaoMock) GetUser(userId int64) (*domain.User, error) {
	if userId == 0 {

		return nil, errors.New(fmt.Sprintf("User %v was not found.", userId))
	}

	return &userAdriano, nil
}
