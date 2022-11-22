package internal

import (
	"context"
	"github.com/cvetkovski98/zvax-keys/internal/model"
	"github.com/cvetkovski98/zvax-keys/internal/model/dto"
)

type KeyService interface {
	RegisterKey(ctx context.Context, key *dto.RegisterKeyInDto) (*model.Key, *string, error)
	ListKeys(ctx context.Context) ([]*model.Key, error)
	GetKey(ctx context.Context, keyId int64) (*model.Key, error)
}
