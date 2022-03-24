package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/renatodalitba/books-solid-go-lang/pkg/config"
	"github.com/renatodalitba/books-solid-go-lang/pkg/errors"

	_ "github.com/jackc/pgx"
)

func Connect(cfg *config.Database) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		return nil, errors.WrapErrorf(err, errors.ErrCodeUnknown, "failed to connect to database")
	}

	return db, nil
}
