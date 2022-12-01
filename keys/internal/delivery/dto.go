package delivery

import (
	"github.com/cvetkovski98/zvax-common/gen/pbkey"
	"github.com/cvetkovski98/zvax-keys/internal/dto"
)

func RegisterKeyRequestToDto(request *pbkey.RegisterKeyRequest) *dto.RegisterKey {
	return &dto.RegisterKey{
		Holder:      request.Holder,
		Affiliation: request.Affiliation,
		Value:       request.Value,
	}
}

func KeyDtoToRegisterResponse(dto *dto.Key) *pbkey.RegisterKeyResponse {
	return &pbkey.RegisterKeyResponse{
		Holder:      dto.Holder,
		Affiliation: dto.Affiliation,
		Value:       dto.Value,
	}
}

func KeyDtoToResponse(dto *dto.Key) *pbkey.KeyResponse {
	return &pbkey.KeyResponse{
		Holder:      dto.Holder,
		Affiliation: dto.Affiliation,
		Value:       dto.Value,
	}
}

func KeysDtoToResponse(dtos *dto.Keys) *pbkey.KeysResponse {
	var keys []*pbkey.KeyResponse
	for _, key := range dtos.Keys {
		keys = append(keys, KeyDtoToResponse(key))
	}
	return &pbkey.KeysResponse{
		Keys: keys,
	}
}
