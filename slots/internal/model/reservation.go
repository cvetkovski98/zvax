package model

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

type Reservation struct {
	ReservationID *string   `json:"reservationId" redis:"reservationId"`
	Slot          *Slot     `json:"slot" redis:"slot"`
	ValidUntil    time.Time `json:"validUntil" redis:"validUntil"`
}

func (reservation *Reservation) ToMap() map[string]string {
	return map[string]string{
		"reservationId": *reservation.ReservationID,
		"slotId":        *reservation.Slot.SlotID,
		"validUntil":    reservation.ValidUntil.Format(time.RFC3339),
	}
}

func NewReservationFromMaps(rMap map[string]string, sMap map[string]string) (*Reservation, error) {
	reservationId, ok := rMap["reservationId"]
	if !ok {
		return nil, errors.New("reservationId is not in hash")
	}
	validUntil, ok := rMap["validUntil"]
	if !ok {
		return nil, errors.New("valid_until is not in hash")
	}
	dateTime, err := time.Parse(time.RFC3339, validUntil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse valid_until")
	}
	slot, err := NewSlotFromMap(sMap)
	return &Reservation{
		ReservationID: &reservationId,
		ValidUntil:    dateTime,
		Slot:          slot,
	}, nil
}

func NewReservationRedisId(slotId string) string {
	return fmt.Sprintf("reservation:%s", slotId)
}
