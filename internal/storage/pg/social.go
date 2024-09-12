package pg

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vadskev/banners-rotation/internal/logger"
	"github.com/vadskev/banners-rotation/internal/models"
)

func (d *dbStorage) AddSocialGroup(ctx context.Context, socialGroup *models.SocialGroup) (*models.SocialGroup, error) {
	var socialGroupID int32
	stmt := `INSERT INTO social_group_table (name) VALUES($1) RETURNING id`

	err := d.db.QueryRow(ctx, stmt, socialGroup.Description).Scan(&socialGroupID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	socialGroup.ID = socialGroupID

	return socialGroup, nil
}

func (d *dbStorage) DeleteSocialGroup(ctx context.Context, socialGroup *models.SocialGroup) (string, error) {
	res, err := d.db.Exec(ctx, "DELETE FROM social_group_table WHERE id = $1", socialGroup.ID)
	if err != nil {
		return "Error", err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return "Error", errors.New("no banner to delete")
	}

	return "OK", nil
}
