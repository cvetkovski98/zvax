package delivery

import (
	"context"

	"github.com/cvetkovski98/zvax-common/gen/pbkey"
	keys "github.com/cvetkovski98/zvax-keys/internal"
)

type server struct {
	ks keys.Service

	pbkey.UnimplementedKeyServer
}

func (s *server) RegisterKey(ctx context.Context, request *pbkey.RegisterKeyRequest) (*pbkey.RegisterKeyResponse, error) {
	var dto = RegisterKeyRequestToDto(request)
	key, _, err := s.ks.RegisterKey(ctx, dto)
	if err != nil {
		return nil, err
	}
	return KeyDtoToRegisterResponse(key), nil
}

func (s *server) GetKeys(ctx context.Context, request *pbkey.KeysRequest) (*pbkey.KeysResponse, error) {
	keys, err := s.ks.ListKeys(ctx, request.Holder)
	if err != nil {
		return nil, err
	}
	return KeysDtoToResponse(keys), nil
}

func NewKeyServer(keyService keys.Service) pbkey.KeyServer {
	return &server{
		ks: keyService,
	}
}
