package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "We were not expecting an user with id 0 ")

	assert.NotNil(t, err, "We were expected an error when user id is 0")
}

func TestGetUserFound(t *testing.T) {
	user, err := GetUser(1)

	assert.NotNil(t, user, "We were expecting an user with id 1")
	assert.Nil(t, err, "We were not expecting an error when user is 1")
	assert.EqualValues(t, "Adriano", user.FirstName)

}
