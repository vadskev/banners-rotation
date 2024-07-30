package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vadskev/banners-rotation/internal/models"
)

type Storage interface {
	AddBanner(ctx context.Context) (models.Banner, error)
	DeleteBanner(ctx context.Context) (string, error)

	AddSlot(ctx context.Context) (models.Slot, error)
	DeleteSlot(ctx context.Context) string

	AddSocialGroup(ctx context.Context) (models.SocialGroup, error)
	DeleteSocialGroup(ctx context.Context) string

	AttachBanner(ctx context.Context) (models.AttachBanner, error)
	DetachBanner(ctx context.Context) string

	HitBanner(ctx context.Context) (models.HitBanner, error)
	SelectBanner(ctx context.Context) (models.Banner, error)

	Ping(ctx context.Context) error
	Close()
	DB() *pgxpool.Pool
}
