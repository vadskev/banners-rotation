package rotation

import (
	"context"
	"fmt"

	"github.com/vadskev/banners-rotation/internal/converter"
	"github.com/vadskev/banners-rotation/internal/logger"
	desc "github.com/vadskev/banners-rotation/pkg/rotation_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i Implementation) AddSlot(ctx context.Context, request *desc.AddSlotRequest) (*desc.AddSlotResponse, error) {

	if len(request.Description) == 0 {
		return nil, status.Error(codes.Aborted, "description is required")
	}

	slot, err := i.storageService.AddSlot(ctx, converter.ToSlotFromAddSlotRequest(request))
	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	}

	logger.Info(fmt.Sprintf("inserted slot: %#v", slot))

	return &desc.AddSlotResponse{
		Id: slot.ID,
	}, nil
}

func (i Implementation) DeleteSlot(ctx context.Context, request *desc.DeleteSlotRequest) (*desc.DeleteSlotResponse, error) {
	if request.Id <= 0 {
		return nil, status.Error(codes.Aborted, "id is required")

	}

	st, err := i.storageService.DeleteSlot(ctx, converter.ToSlotFromDeleteSlotRequest(request))
	if err != nil {
		return nil, status.Error(codes.NotFound, "slot not found")
	}

	logger.Info(fmt.Sprintf("deleted slot: %#v", st))

	return &desc.DeleteSlotResponse{
		Status: "OK",
	}, nil
}
