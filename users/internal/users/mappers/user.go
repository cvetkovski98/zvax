package mappers

import (
	"github.com/cvetkovski98/zvax-common/gen/pbuser"
	"github.com/cvetkovski98/zvax-users/internal/users/model"
)

func NewUserResponseFromUser(user *model.User) *pbuser.UserResponse {
	return &pbuser.UserResponse{
		UserId: *user.UserId,
		Name:   user.Name,
		Email:  user.Email,
		Phone:  user.Phone,
	}
}

func NewUserFromCreateUserRequest(request *pbuser.CreateUserRequest) *model.User {
	return &model.User{
		UserId: nil,
		Name:   request.Name,
		Email:  request.Email,
		Phone:  request.Phone,
	}
}
