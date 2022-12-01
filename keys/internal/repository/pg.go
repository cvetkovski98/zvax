package repository

import (
	"context"

	keys "github.com/cvetkovski98/zvax-keys/internal"
	"github.com/cvetkovski98/zvax-keys/internal/model"
	"github.com/uptrace/bun"
)

type pg struct {
	db *bun.DB
}

func (repository *pg) InsertOne(ctx context.Context, key *model.Key) (*model.Key, error) {
	if _, err := repository.db.NewInsert().Model(key).Exec(ctx); err != nil {
		return nil, err
	}
	return key, nil
}

func (repository *pg) FindAllByHolder(ctx context.Context, holder string) ([]*model.Key, error) {
	var keys []*model.Key
	var query = repository.db.NewSelect().Model(&keys).Where("holder = ?", holder)
	if err := query.Scan(ctx); err != nil {
		return nil, err
	}
	return keys, nil
}

func (repository *pg) FindOneById(ctx context.Context, id int) (*model.Key, error) {
	var key = new(model.Key)
	var query = repository.db.NewSelect().Model(key).Where("id = ?", id)
	if err := query.Scan(ctx); err != nil {
		return nil, err
	}
	return key, nil
}

func NewPgKeyRepository(db *bun.DB) keys.Repository {
	return &pg{db: db}
}
