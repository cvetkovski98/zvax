package internal

import (
	"context"
	"github.com/cvetkovski98/zvax-slots/internal/model"
)

type SlotService interface {
	GetSlotsAtLocationBetween(ctx context.Context, page *model.PageRequest) (*model.Page[model.Slot], error)
	CreateSlot(ctx context.Context, slot *model.Slot) (*model.Slot, error)
	CreateReservation(ctx context.Context, slotId string) (*model.Reservation, error)
	ConfirmReservation(ctx context.Context, reservationId string) (string, error)
}
