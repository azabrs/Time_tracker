package time_tracker_db

import (
	"Time-tracker/internal/config"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewDB(cfg config.DB) (repository, error) {
	db, err := sqlx.Connect(cfg.DriverName, cfg.DSN)
	if err != nil {
		return repository{}, fmt.Errorf("cannot open database connection: %w", err)
	}
	if err = db.Ping(); err != nil {
		return repository{}, fmt.Errorf("cannot connect to database: %w", err)
	}

	return repository{
		db: db,
	}, nil
}
