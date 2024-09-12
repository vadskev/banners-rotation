package pg

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vadskev/banners-rotation/internal/logger"
	"github.com/vadskev/banners-rotation/internal/models"
)

func (d *dbStorage) AttachBanner(ctx context.Context, attach *models.AttachBanner) (*models.AttachBanner, error) {

	stmt := `INSERT INTO banner_slot_table (banner_id, slot_id) VALUES($1, $2)`

	res, err := d.db.Exec(ctx, stmt, attach.BannerID, attach.SlotID)
	if err != nil {
		logger.Error(errors.New("error to Attach").Error())
		return nil, errors.New("error to Attach")
	}
	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return nil, errors.New("failed to Attach")
	}

	return attach, nil
}

func (d *dbStorage) DetachBanner(ctx context.Context, attach *models.AttachBanner) (string, error) {
	res, err := d.db.Exec(ctx, "DELETE FROM banner_slot_table WHERE banner_id = $1 AND slot_id = $2", attach.BannerID, attach.SlotID)
	if err != nil {
		logger.Error(errors.New("error to Detach").Error())
		return "Error", err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return "Error", errors.New("failed to Detach")
	}
	return "OK", nil
}
