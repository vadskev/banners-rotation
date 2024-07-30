package rotation

import (
	"context"

	"github.com/vadskev/banners-rotation/internal/logger"
	desc "github.com/vadskev/banners-rotation/pkg/rotation_v1"
)

func (i Implementation) AttachBanner(ctx context.Context, request *desc.AttachBannerRequest) (*desc.AttachBannerResponse, error) {
	//TODO implement me
	logger.Info("ololo")
	panic("implement me")
}

func (i Implementation) DetachBanner(ctx context.Context, request *desc.DetachBannerRequest) (*desc.DetachBannerResponse, error) {
	//TODO implement me
	panic("implement me")
}
