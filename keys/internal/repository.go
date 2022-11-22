package internal

import (
	"context"

	"github.com/cvetkovski98/zvax-keys/internal/model"
)

type KeyRepository interface {
	InsertOne(ctx context.Context, key *model.Key) (*model.Key, error)
	FindAll(ctx context.Context) ([]*model.Key, error)
	FindOneById(ctx context.Context, keyId int64) (*model.Key, error)
}
