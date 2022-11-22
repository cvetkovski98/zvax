package repository

import (
	"context"

	"github.com/cvetkovski98/zvax-auth/internal/model"
	"github.com/uptrace/bun"
)

func RegisterModels(ctx context.Context, db *bun.DB) error {
	db.RegisterModel(
		(*model.RolePermission)(nil),
		(*model.UserRole)(nil),
	)
	return nil
}
