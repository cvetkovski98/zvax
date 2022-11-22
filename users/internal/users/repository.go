package users

import (
	"context"
	"github.com/cvetkovski98/zvax-users/internal/users/model"
)

type UserRepository interface {
	FindByID(ctx context.Context, userId int64) (*model.User, error)
	InsertOne(ctx context.Context, user *model.User) (*model.User, error)
}
