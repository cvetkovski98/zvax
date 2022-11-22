package model

import "fmt"

type Key struct {
	KeyId       *int64
	Holder      string
	Affiliation string
	Value       string
}

func (k *Key) String() string {
	return fmt.Sprintf(
		"Key{KeyId: %d, Holder: %s, Affiliation: %s, Value: %s}",
		*k.KeyId, k.Holder, k.Affiliation, k.Value,
	)
}
