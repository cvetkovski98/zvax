package users

import (
	"context"
	"github.com/cvetkovski98/zvax-users/internal/users/model"
)

type UserService interface {
	ListOne(ctx context.Context, userId int64) (*model.User, error)
	CreateOne(ctx context.Context, user *model.User) (*model.User, error)
}
