package keys

import (
	"context"

	"github.com/cvetkovski98/zvax-keys/internal/dto"
)

type Service interface {
	RegisterKey(ctx context.Context, key *dto.RegisterKey) (*dto.Key, string, error)
	ListKeys(ctx context.Context, holder string) (*dto.Keys, error)
	GetKey(ctx context.Context, id int) (*dto.Key, error)
}
