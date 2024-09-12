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

func (i Implementation) AttachBanner(ctx context.Context, request *desc.AttachBannerRequest) (*desc.AttachBannerResponse, error) {
	if request.BannerId <= 0 {
		return nil, status.Error(codes.Aborted, "Banner Id is required")
	}

	if request.SlotId <= 0 {
		return nil, status.Error(codes.Aborted, "Slot Id is required")
	}

	attach, err := i.storageService.AttachBanner(ctx, converter.ToAttachBannerFromAttachBannerRequest(request))
	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	}

	logger.Info(fmt.Sprintf("attach: %#v", attach))

	return &desc.AttachBannerResponse{
		Status: "OK",
	}, nil
}

func (i Implementation) DetachBanner(ctx context.Context, request *desc.DetachBannerRequest) (*desc.DetachBannerResponse, error) {
	if request.BannerId <= 0 {
		return nil, status.Error(codes.Aborted, "Banner Id is required")
	}
	if request.SlotId <= 0 {
		return nil, status.Error(codes.Aborted, "Slot Id is required")
	}

	st, err := i.storageService.DetachBanner(ctx, converter.ToAttachBannerFromDetachBannerRequest(request))
	if err != nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	logger.Info(fmt.Sprintf("deleted banner: %#v", st))

	return &desc.DetachBannerResponse{
		Status: "OK",
	}, nil
}
