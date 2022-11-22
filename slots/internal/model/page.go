package model

import (
	"time"
)

type Model interface {
	Slot
}

type Page[T Model] struct {
	Items []*T
}

type PageRequest struct {
	StartDate time.Time
	EndDate   time.Time
	Location  string
}
