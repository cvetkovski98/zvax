package internal

import (
	"context"
	"github.com/cvetkovski98/zvax-slots/internal/model"
	"time"
)

type SlotRepository interface {
	FindOneByKey(ctx context.Context, key string) (*model.Slot, error)
	FindAllWithDateTimeBetween(ctx context.Context, from time.Time, to time.Time) ([]*model.Slot, error)
	InsertOne(ctx context.Context, slot *model.Slot) (*model.Slot, error)
	ReserveOneByKey(ctx context.Context, key string) (*model.Reservation, error)
	ConfirmOneByReservationId(ctx context.Context, reservationId string) (*model.Reservation, error)
	CleanUpExpiredReservations(ctx context.Context) error
}
