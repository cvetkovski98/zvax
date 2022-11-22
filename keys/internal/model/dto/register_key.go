package dto

import (
	"fmt"
)

type RegisterKeyInDto struct {
	Holder      string
	Affiliation string
	PublicKey   string
}

type RegisterKeyOutDto struct {
	KeyId       int64
	Holder      string
	Affiliation string
	PublicKey   string
	Certificate string
}

func (dto *RegisterKeyInDto) String() string {
	return fmt.Sprintf(
		"RegisterKeyInDto{Holder: %s, Affiliation: %s, PublicKey: %s}",
		dto.Holder, dto.Affiliation, dto.PublicKey,
	)
}
