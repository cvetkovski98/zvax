package keys

import (
	"context"

	"github.com/cvetkovski98/zvax-keys/internal/model"
)

type Repository interface {
	InsertOne(ctx context.Context, key *model.Key) (*model.Key, error)
	FindAllByHolder(ctx context.Context, holder string) ([]*model.Key, error)
	FindOneById(ctx context.Context, id int) (*model.Key, error)
}
