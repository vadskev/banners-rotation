package pg

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vadskev/banners-rotation/internal/logger"
	"github.com/vadskev/banners-rotation/internal/models"
)

func (d *dbStorage) AddSlot(ctx context.Context, slot *models.Slot) (*models.Slot, error) {
	var slotID int32
	stmt := `INSERT INTO slots_table (name) VALUES($1) RETURNING id`

	err := d.db.QueryRow(ctx, stmt, slot.Description).Scan(&slotID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	slot.ID = slotID

	return slot, nil
}

func (d *dbStorage) DeleteSlot(ctx context.Context, slot *models.Slot) (string, error) {
	res, err := d.db.Exec(ctx, "DELETE FROM slots_table WHERE id = $1", slot.ID)
	if err != nil {
		return "Error", err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return "Error", errors.New("no banner to delete")
	}

	return "OK", nil
}
