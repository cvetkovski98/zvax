package mapper

import (
	"github.com/cvetkovski98/zvax-keys/internal/dto"
	"github.com/cvetkovski98/zvax-keys/internal/model"
)

func RegisterKeyDtoToModel(dto *dto.RegisterKey) *model.Key {
	return &model.Key{
		Holder:      dto.Holder,
		Affiliation: dto.Affiliation,
		Value:       dto.Value,
	}
}

func KeyDtoToModel(dto *dto.Key) *model.Key {
	return &model.Key{
		Holder:      dto.Holder,
		Affiliation: dto.Affiliation,
		Value:       dto.Value,
	}
}

func KeysDtoToModel(dtos *dto.Keys) []*model.Key {
	var keys []*model.Key
	for _, key := range dtos.Keys {
		keys = append(keys, KeyDtoToModel(key))
	}
	return keys
}
