package pg

import (
	"context"
	"errors"

	"github.com/vadskev/banners-rotation/internal/logger"
	"github.com/vadskev/banners-rotation/internal/models"
)

func (d *dbStorage) AddBanner(ctx context.Context, banner *models.Banner) (*models.Banner, error) {
	var bannerID int32
	stmt := `INSERT INTO banners_table (name) VALUES($1) RETURNING id`

	err := d.db.QueryRow(ctx, stmt, banner.Description).Scan(&bannerID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	banner.ID = bannerID

	return banner, nil
}

func (d *dbStorage) DeleteBanner(ctx context.Context, banner *models.Banner) (string, error) {
	res, err := d.db.Exec(ctx, "DELETE FROM banners_table WHERE id = $1", banner.ID)
	if err != nil {
		return "Error", err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return "Error", errors.New("no banner to delete")
	}

	return "OK", nil
}
