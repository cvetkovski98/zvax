package service

import (
	"context"
	"github.com/cvetkovski98/zvax-users/internal/users"
	"github.com/cvetkovski98/zvax-users/internal/users/model"
	"github.com/pkg/errors"
)

type UserServiceImpl struct {
	userRepository users.UserRepository
}

func (service *UserServiceImpl) ListOne(ctx context.Context, userId int64) (*model.User, error) {
	user, err := service.userRepository.FindByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserServiceImpl) CreateOne(ctx context.Context, user *model.User) (*model.User, error) {
	if user.UserId != nil {
		return nil, errors.New("user id must be nil")
	}
	created, err := service.userRepository.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func NewUserServiceImpl(userRepository users.UserRepository) users.UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}
