package repository

import (
	"context"

	auth "github.com/cvetkovski98/zvax-auth/internal"
	"github.com/cvetkovski98/zvax-auth/internal/model"
	"github.com/uptrace/bun"
)

type pg struct {
	db *bun.DB
}

func (repository *pg) InsertOne(ctx context.Context, user *model.User) (*model.User, error) {
	if _, err := repository.db.NewInsert().Model(user).Exec(ctx); err != nil {
		return nil, err
	}
	return user, nil
}

func (repository *pg) FindOneByEmail(ctx context.Context, email string) (*model.User, error) {
	var user = new(model.User)
	var query = repository.db.NewSelect().Model(user).Where("email = ?", email)
	if err := query.Scan(ctx); err != nil {
		return nil, err
	}
	return user, nil
}

func NewPgAuthRepository(db *bun.DB) auth.Repository {
	return &pg{db}
}
