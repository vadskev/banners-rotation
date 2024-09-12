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

func (i Implementation) AddBanner(ctx context.Context, request *desc.AddBannerRequest) (*desc.AddBannerResponse, error) {
	if len(request.Description) == 0 {
		return nil, status.Error(codes.Aborted, "description is required")
	}

	banner, err := i.storageService.AddBanner(ctx, converter.ToBannerFromAddBannerRequest(request))
	if err != nil {
		return nil, status.Error(codes.Canceled, err.Error())
	}

	logger.Info(fmt.Sprintf("inserted banner: %#v", banner))

	return &desc.AddBannerResponse{
		Id: banner.ID,
	}, nil
}

func (i Implementation) DeleteBanner(ctx context.Context, request *desc.DeleteBannerRequest) (*desc.DeleteBannerResponse, error) {
	if request.Id <= 0 {
		return nil, status.Error(codes.Aborted, "id is required")

	}

	st, err := i.storageService.DeleteBanner(ctx, converter.ToBannerFromDeleteBannerRequest(request))
	if err != nil {
		return nil, status.Error(codes.NotFound, "banner not found")
	}

	logger.Info(fmt.Sprintf("deleted banner: %#v", st))

	return &desc.DeleteBannerResponse{
		Status: "OK",
	}, nil
}
