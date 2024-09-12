package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vadskev/banners-rotation/internal/models"
)

type Storage interface {
	AddBanner(ctx context.Context, banner *models.Banner) (*models.Banner, error)
	DeleteBanner(ctx context.Context, banner *models.Banner) (string, error)

	AddSlot(ctx context.Context, slot *models.Slot) (*models.Slot, error)
	DeleteSlot(ctx context.Context, slot *models.Slot) (string, error)

	AddSocialGroup(ctx context.Context, socialGroup *models.SocialGroup) (*models.SocialGroup, error)
	DeleteSocialGroup(ctx context.Context, socialGroup *models.SocialGroup) (string, error)

	AttachBanner(ctx context.Context) (models.AttachBanner, error)
	DetachBanner(ctx context.Context) string

	HitBanner(ctx context.Context) (models.HitBanner, error)
	SelectBanner(ctx context.Context) (models.Banner, error)

	Ping(ctx context.Context) error
	Close()
	DB() *pgxpool.Pool
}
