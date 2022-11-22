package mappers

import (
	"github.com/cvetkovski98/zvax-common/gen/pbkey"
	"github.com/cvetkovski98/zvax-keys/internal/model"
	"github.com/cvetkovski98/zvax-keys/internal/model/dto"
)

func NewRegisterKeyInFromRegisterKeyRequest(request *pbkey.RegisterKeyRequest) *dto.RegisterKeyInDto {
	return &dto.RegisterKeyInDto{
		Holder:      request.Holder,
		Affiliation: request.Affiliation,
		PublicKey:   request.PublicKey,
	}
}

func NewKeyResponseFromKey(key *model.Key) *pbkey.KeyResponse {
	return &pbkey.KeyResponse{
		KeyId:       *key.KeyId,
		Holder:      key.Holder,
		Affiliation: key.Affiliation,
		Value:       key.Value,
	}
}

func NewRegisterKeyResponseFromKey(key *model.Key, cert *string) *pbkey.RegisterKeyResponse {
	return &pbkey.RegisterKeyResponse{
		Key:         NewKeyResponseFromKey(key),
		Certificate: *cert,
	}
}

func NewGetKeysResponseFromKeys(keys []*model.Key) *pbkey.GetKeysResponse {
	payload := make([]*pbkey.KeyResponse, len(keys))
	for i, key := range keys {
		payload[i] = NewKeyResponseFromKey(key)
	}
	return &pbkey.GetKeysResponse{
		Keys: payload,
	}
}
