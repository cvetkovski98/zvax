package mappers

import (
	"github.com/cvetkovski98/zvax-common/gen/pbslot"
	"github.com/cvetkovski98/zvax-slots/internal/model"
	"github.com/pkg/errors"
	"time"
)

func NewPageRequestFromSlotListRequest(pb *pbslot.SlotListRequest) (*model.PageRequest, error) {
	sd, err := time.Parse(time.RFC3339, pb.StartDate)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse start date")
	}
	ed, err := time.Parse(time.RFC3339, pb.EndDate)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse end date")
	}
	if ed.Before(sd) {
		return nil, errors.New("end date is before start date")
	}
	return &model.PageRequest{
		StartDate: sd,
		EndDate:   ed,
		Location:  pb.Location,
	}, nil
}

func NewSlotListResponseFromPageResponse(page model.Page[model.Slot]) *pbslot.SlotListResponse {
	pbSlots := make([]*pbslot.SlotResponse, len(page.Items))
	for i, slot := range page.Items {
		pbSlots[i] = NewSlotResponseFromSlot(*slot)
	}
	return &pbslot.SlotListResponse{
		Items: pbSlots,
	}
}
