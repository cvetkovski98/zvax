package service

import (
	"context"
	"github.com/cvetkovski98/zvax-users/internal/users"
	"github.com/cvetkovski98/zvax-users/internal/users/model"
	"github.com/cvetkovski98/zvax-users/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

var source map[int64]*model.User

func init() {
	source = map[int64]*model.User{
		1: {
			UserId: nil,
			Name:   "John Doe",
			Email:  "email1",
			Phone:  "123456789",
		},
		2: {
			UserId: nil,
			Name:   "Jane Doe",
			Email:  "email2",
			Phone:  "123456789",
		},
		3: {
			UserId: nil,
			Name:   "Jack Doe",
			Email:  "email3",
			Phone:  "123456789",
		},
	}
	for i, user := range source {
		id := i + 1
		user.UserId = &id
	}
}

func getTestUserService() users.UserService {
	mockUserRepository := test.NewMockUserRepository(source)
	return NewUserServiceImpl(mockUserRepository)
}

func TestUserServiceListOne(t *testing.T) {
	userService := getTestUserService()
	searchedId := int64(1)
	expected := source[searchedId]
	actual, err := userService.ListOne(context.Background(), searchedId)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestUserServiceCreateOne(t *testing.T) {
	userService := getTestUserService()
	user := &model.User{
		UserId: nil,
		Name:   "John Doe",
		Email:  "email1",
		Phone:  "123456789",
	}
	created, err := userService.CreateOne(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, user, created)
}

func TestUserServiceCreatOne_WithId(t *testing.T) {
	userService := getTestUserService()
	user := &model.User{
		UserId: nil,
		Name:   "John Doe",
		Email:  "email1",
		Phone:  "123456789",
	}
	id := int64(1)
	user.UserId = &id
	_, err := userService.CreateOne(context.Background(), user)
	assert.Error(t, err)
}

func TestUserServiceListOneNotFound(t *testing.T) {
	userService := getTestUserService()
	searchedId := int64(0)
	_, err := userService.ListOne(context.Background(), searchedId)
	assert.Error(t, err)
}
