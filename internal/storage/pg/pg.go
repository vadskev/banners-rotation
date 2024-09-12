package pg

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
	"github.com/vadskev/banners-rotation/internal/storage"
)

type dbStorage struct {
	db *pgxpool.Pool
}

var _ storage.Storage = (*dbStorage)(nil)

func New(ctx context.Context, dsn string) (storage.Storage, error) {
	dbc, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, errors.Errorf("failed to connect to db: %v", err)
	}
	return &dbStorage{
		db: dbc,
	}, nil
}

func (d *dbStorage) Ping(ctx context.Context) error {
	return d.db.Ping(ctx)
}

func (d *dbStorage) Close() {
	if d.db != nil {
		defer func() {
			d.db.Close()
		}()
	}
}

func (d *dbStorage) DB() *pgxpool.Pool {
	return d.db
}
