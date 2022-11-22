package delivery

import (
	"context"
	"github.com/cvetkovski98/zvax-common/gen/pbslot"
	"github.com/cvetkovski98/zvax-slots/internal"
	"github.com/cvetkovski98/zvax-slots/internal/mappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SlotGrpcServerImpl struct {
	pbslot.UnimplementedSlotGrpcServer
	slotService internal.SlotService
}

func (s SlotGrpcServerImpl) GetSlotList(
	ctx context.Context,
	request *pbslot.SlotListRequest,
) (*pbslot.SlotListResponse, error) {
	pageRequest, err := mappers.NewPageRequestFromSlotListRequest(request)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}
	payload, err := s.slotService.GetSlotsAtLocationBetween(ctx, pageRequest)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get slots: %v", err)
	}
	return mappers.NewSlotListResponseFromPageResponse(*payload), nil
}

func (s SlotGrpcServerImpl) CreateSlotReservation(
	ctx context.Context,
	request *pbslot.SlotReservationRequest,
) (*pbslot.SlotReservationResponse, error) {
	reservation, err := s.slotService.CreateReservation(ctx, request.SlotId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create slot: %v", err)
	}
	return mappers.NewSlotReservationResponseFromReservation(*reservation), nil
}

func (s SlotGrpcServerImpl) ConfirmSlotReservation(
	ctx context.Context,
	request *pbslot.SlotConfirmationRequest,
) (*pbslot.SlotConfirmationResponse, error) {
	token, err := s.slotService.ConfirmReservation(ctx, request.ReservationId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to confirm slot: %v", err)
	}
	return &pbslot.SlotConfirmationResponse{
		SlotConfirmationToken: token,
	}, nil
}

func NewSlotGrpcServerImpl(slotService internal.SlotService) pbslot.SlotGrpcServer {
	return &SlotGrpcServerImpl{
		slotService: slotService,
	}
}
