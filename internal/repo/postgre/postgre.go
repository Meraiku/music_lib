package postgre

import (
	"database/sql"
	"os"
	"sync"

	"github.com/meraiku/music_lib/internal/repo"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var _ repo.MusicRepository = (*postgre)(nil)

var dsn string

type postgre struct {
	db *bun.DB
	mu *sync.RWMutex
}

func New() (*postgre, error) {
	dsn = os.Getenv("POSTGRES_DSN")
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	if err := sqldb.Ping(); err != nil {
		return nil, err
	}

	if err := Songs(); err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, pgdialect.New())

	return &postgre{
		db: db,
		mu: &sync.RWMutex{}}, nil
}
