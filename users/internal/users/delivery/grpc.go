package delivery

import (
	"context"
	"github.com/cvetkovski98/zvax-common/gen/pbuser"
	"github.com/cvetkovski98/zvax-users/internal/users"
	"github.com/cvetkovski98/zvax-users/internal/users/mappers"
)

type UserGrpcServerImpl struct {
	pbuser.UnimplementedUserGrpcServer
	userService users.UserService
}

func (server *UserGrpcServerImpl) GetUser(ctx context.Context, request *pbuser.GetUserRequest) (*pbuser.UserResponse, error) {
	user, err := server.userService.ListOne(ctx, request.UserId)
	if err != nil {
		return nil, err
	}
	payload := mappers.NewUserResponseFromUser(user)
	return payload, nil
}

func (server *UserGrpcServerImpl) CreateUser(ctx context.Context, request *pbuser.CreateUserRequest) (*pbuser.UserResponse, error) {
	user := mappers.NewUserFromCreateUserRequest(request)
	created, err := server.userService.CreateOne(ctx, user)
	if err != nil {
		return nil, err
	}
	payload := mappers.NewUserResponseFromUser(created)
	return payload, nil
}

func NewUserGrpcImpl(userService users.UserService) pbuser.UserGrpcServer {
	return &UserGrpcServerImpl{
		userService: userService,
	}
}
