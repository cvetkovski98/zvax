package auth

import (
	"context"

	"github.com/cvetkovski98/zvax-auth/internal/model"
)

type Repository interface {
	InsertOne(ctx context.Context, user *model.User) (*model.User, error)
	FindOneByEmail(ctx context.Context, email string) (*model.User, error)
}
