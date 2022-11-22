package service

import (
	"context"
	"fmt"
	"github.com/cvetkovski98/zvax-slots/internal"
	"github.com/cvetkovski98/zvax-slots/internal/model"
	"github.com/pkg/errors"
)

type SlotServiceImpl struct {
	slotRepository internal.SlotRepository
}

func (service *SlotServiceImpl) GetSlotsAtLocationBetween(ctx context.Context, page *model.PageRequest) (*model.Page[model.Slot], error) {
	slots, err := service.slotRepository.FindAllWithDateTimeBetween(ctx, page.StartDate, page.EndDate)
	if err != nil {
		return nil, err
	}
	return &model.Page[model.Slot]{
		Items: slots,
	}, nil
}

func (service *SlotServiceImpl) CreateSlot(ctx context.Context, slot *model.Slot) (*model.Slot, error) {
	if slot.SlotID != nil {
		return nil, errors.New("slot id must be nil")
	}
	return service.slotRepository.InsertOne(ctx, slot)
}

func (service *SlotServiceImpl) CreateReservation(ctx context.Context, slotId string) (*model.Reservation, error) {
	reservation, err := service.slotRepository.ReserveOneByKey(ctx, slotId)
	if err != nil {
		return nil, err
	}
	return reservation, nil
}

func (service *SlotServiceImpl) ConfirmReservation(ctx context.Context, reservationId string) (string, error) {
	reservation, err := service.slotRepository.ConfirmOneByReservationId(ctx, reservationId)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf(
			"error confirming reservation with id=%s",
			*reservation.ReservationID,
		))
	}
	// TODO: generate reservation token
	return *reservation.ReservationID, nil
}

func NewSlotServiceImpl(slotRepository internal.SlotRepository) internal.SlotService {
	return &SlotServiceImpl{
		slotRepository: slotRepository,
	}
}
