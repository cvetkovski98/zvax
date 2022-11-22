package delivery

import (
	"context"
	"fmt"
	"strings"

	auth "github.com/cvetkovski98/zvax-auth/internal"
	"github.com/cvetkovski98/zvax-auth/internal/token"
	"github.com/cvetkovski98/zvax-common/gen/pbauth"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type key int

const (
	setCookieHeaderKey string = "set-cookie"
	getCookieHeaderKey string = "cookie"
	ctxTokenKey        key    = iota
)

type server struct {
	as auth.Service

	pbauth.UnimplementedAuthServer
}

func setTokenAsCookie(ctx context.Context, t *token.Token) {
	tokenValue := t.ToCookieValue()
	headerValue := fmt.Sprintf("token=%s; HttpOnly; Path=/", tokenValue)
	grpc.SetHeader(ctx, metadata.Pairs(setCookieHeaderKey, headerValue))
}

func getTokenFromCookie(ctx context.Context) (*token.Token, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrMissingMetadata
	}
	cookie := md.Get(getCookieHeaderKey)
	if len(cookie) == 0 {
		return nil, ErrMissingCookies
	}
	for _, cookie := range cookie {
		values := strings.Split(cookie, ";")
		for _, value := range values {
			if strings.HasPrefix(value, "token=") {
				tokenValue := strings.TrimPrefix(value, "token=")
				return token.FromCookieValue(tokenValue)
			}
		}
	}
	return nil, ErrMissingToken
}

func (s *server) Me(ctx context.Context, _ *empty.Empty) (*pbauth.AuthResponse, error) {
	t, err := getTokenFromCookie(ctx)
	if err != nil {
		return nil, err
	}
	auth, err := s.as.Me(ctx, t)
	if err != nil {
		return nil, err
	}
	return UserDtoToResponse(auth), nil
}

func (s *server) Login(ctx context.Context, request *pbauth.LoginRequest) (*pbauth.AuthResponse, error) {
	payload := LoginRequestToDto(request)
	auth, t, err := s.as.Login(ctx, payload)
	if err != nil {
		return nil, err
	}
	setTokenAsCookie(ctx, t)
	return UserDtoToResponse(auth), nil
}

func (s *server) Register(ctx context.Context, request *pbauth.RegisterRequest) (*pbauth.AuthResponse, error) {
	payload := RegisterRequestToDto(request)
	auth, t, err := s.as.Register(ctx, payload)
	if err != nil {
		return nil, err
	}
	setTokenAsCookie(ctx, t)
	return UserDtoToResponse(auth), nil
}

func NewAuthServer(authService auth.Service) pbauth.AuthServer {
	return &server{as: authService}
}
