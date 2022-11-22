package mappers

import (
	"github.com/cvetkovski98/zvax-common/gen/pbslot"
	"github.com/cvetkovski98/zvax-slots/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewSlotReservationResponseFromReservation(reservation model.Reservation) *pbslot.SlotReservationResponse {
	return &pbslot.SlotReservationResponse{
		ReservationId: *reservation.ReservationID,
		ValidUntil:    timestamppb.New(reservation.ValidUntil),
	}
}
