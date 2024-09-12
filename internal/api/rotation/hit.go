package rotation

import (
	"context"
	"time"

	"github.com/vadskev/banners-rotation/internal/models"
	desc "github.com/vadskev/banners-rotation/pkg/rotation_v1"
)

func (i Implementation) HitBanner(ctx context.Context, request *desc.HitBannerRequest) (*desc.HitBannerResponse, error) {

	i.kafkaProducer.SendMessage(
		models.Message{
			Type:          "show",
			SlotID:        12,
			BannerID:      13,
			SocialGroupID: 14,
			CreatedAt:     time.Now().Unix(),
		})

	return &desc.HitBannerResponse{}, nil
}

func (i Implementation) SelectBanner(ctx context.Context, request *desc.SelectBannerRequest) (*desc.SelectBannerResponse, error) {
	//TODO implement me
	panic("implement me")
}
