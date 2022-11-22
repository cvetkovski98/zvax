package delivery

import (
	"context"
	"github.com/cvetkovski98/zvax-common/gen/pbkey"
	"github.com/cvetkovski98/zvax-keys/internal"
	"github.com/cvetkovski98/zvax-keys/internal/mappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type KeyGrpcServerImpl struct {
	pbkey.UnimplementedKeyGrpcServer
	keyService internal.KeyService
}

func (server *KeyGrpcServerImpl) RegisterKey(
	ctx context.Context,
	request *pbkey.RegisterKeyRequest,
) (*pbkey.RegisterKeyResponse, error) {
	key := mappers.NewRegisterKeyInFromRegisterKeyRequest(request)
	created, cert, err := server.keyService.RegisterKey(ctx, key)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to register key: %s", err.Error())
	}
	payload := mappers.NewRegisterKeyResponseFromKey(created, cert)
	return payload, nil
}

func (server *KeyGrpcServerImpl) GetKeys(ctx context.Context, _ *emptypb.Empty) (*pbkey.GetKeysResponse, error) {
	keys, err := server.keyService.ListKeys(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to list keys: %s", err.Error())
	}
	payload := mappers.NewGetKeysResponseFromKeys(keys)
	return payload, nil
}

func (server *KeyGrpcServerImpl) GetKey(ctx context.Context, request *pbkey.GetKeyRequest) (*pbkey.KeyResponse, error) {
	key, err := server.keyService.GetKey(ctx, request.KeyId)
	if err != nil {
		return nil, err
	}
	payload := mappers.NewKeyResponseFromKey(key)
	return payload, nil
}

func NewKeyGrpcImpl(keyService internal.KeyService) pbkey.KeyGrpcServer {
	return &KeyGrpcServerImpl{
		keyService: keyService,
	}
}
