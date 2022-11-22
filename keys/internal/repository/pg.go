package repository

import (
	"context"
	"github.com/cvetkovski98/zvax-keys/internal"
	"github.com/cvetkovski98/zvax-keys/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type pgKeyRepository struct {
	pool *pgxpool.Pool
}

func (repository *pgKeyRepository) InsertOne(ctx context.Context, key *model.Key) (*model.Key, error) {
	err := repository.pool.QueryRow(
		ctx,
		"INSERT INTO keys (holder, affiliation, value) VALUES ($1, $2, $3) RETURNING id, holder, affiliation, value",
		key.Holder, key.Affiliation, key.Value,
	).Scan(&key.KeyId, &key.Holder, &key.Affiliation, &key.Value)
	if err != nil {
		return nil, errors.Wrap(err, "error inserting key")
	}
	return key, nil
}

func (repository *pgKeyRepository) FindAll(ctx context.Context) ([]*model.Key, error) {
	var keys []*model.Key
	rows, err := repository.pool.Query(ctx, "SELECT * FROM keys")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var key = new(model.Key)
		err := rows.Scan(&key.KeyId, &key.Holder, &key.Affiliation, &key.Value)
		if err != nil {
			return nil, err
		}
		keys = append(keys, key)
	}
	return keys, nil
}

func (repository *pgKeyRepository) FindOneById(ctx context.Context, keyId int64) (*model.Key, error) {
	var key = new(model.Key)
	err := repository.pool.QueryRow(
		ctx, "SELECT * FROM keys WHERE id = $1", keyId,
	).Scan(&key.KeyId, &key.Holder, &key.Affiliation, &key.Value)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func NewPgKeyRepository(pool *pgxpool.Pool) internal.KeyRepository {
	return &pgKeyRepository{pool: pool}
}
