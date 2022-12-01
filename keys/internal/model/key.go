package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Key struct {
	bun.BaseModel `bun:"table:keys"`

	Id          int    `bun:"id,pk,nullzero"`
	Holder      string `bun:"holder,nullzero,notnull"`
	Affiliation string `bun:"affiliation,nullzero,notnull"`
	Value       string `bun:"value,nullzero,notnull"`

	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}
