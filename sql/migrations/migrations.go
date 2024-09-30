package migrations

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var sqlMigrations embed.FS

func Songs(db *sql.DB) error {

	goose.SetBaseFS(sqlMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db, "."); err != nil {
		return err
	}

	return nil
}
