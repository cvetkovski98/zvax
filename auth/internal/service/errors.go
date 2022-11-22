package service

import "github.com/pkg/errors"

var (
	ErrHashingPassword   = errors.New("error hashing password")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrCreatingUser      = errors.New("error creating user")
	ErrCreatingUserJWT   = errors.New("error creating user jwt")
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrFindingUser       = errors.New("error finding user")
)
