package rotation

import (
	"context"

	desc "github.com/vadskev/banners-rotation/pkg/rotation_v1"
)

func (i Implementation) HitBanner(ctx context.Context, request *desc.HitBannerRequest) (*desc.HitBannerResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (i Implementation) SelectBanner(ctx context.Context, request *desc.SelectBannerRequest) (*desc.SelectBannerResponse, error) {
	//TODO implement me
	panic("implement me")
}
