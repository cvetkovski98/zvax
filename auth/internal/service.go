package auth

import (
	"context"

	"github.com/cvetkovski98/zvax-auth/internal/dto"
	"github.com/cvetkovski98/zvax-auth/internal/token"
)

type Service interface {
	Login(ctx context.Context, user *dto.Login) (*dto.User, *token.Token, error)
	Register(ctx context.Context, user *dto.Register) (*dto.User, *token.Token, error)
	Me(ctx context.Context, t *token.Token) (*dto.User, error)
}
