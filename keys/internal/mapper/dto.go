package mapper

import (
	"github.com/cvetkovski98/zvax-keys/internal/dto"
	"github.com/cvetkovski98/zvax-keys/internal/model"
)

func KeyModelToDto(model *model.Key) *dto.Key {
	return &dto.Key{
		Holder:      model.Holder,
		Affiliation: model.Affiliation,
		Value:       model.Value,
	}
}

func KeysModelToDto(models []*model.Key) *dto.Keys {
	var keys []*dto.Key
	for _, key := range models {
		keys = append(keys, KeyModelToDto(key))
	}
	return &dto.Keys{
		Keys: keys,
	}
}
