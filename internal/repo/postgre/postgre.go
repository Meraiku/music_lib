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

type postgre struct {
	db *bun.DB
	mu *sync.RWMutex
}

func New() (*postgre, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("DB_URL"))))

	if err := sqldb.Ping(); err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, pgdialect.New())

	return &postgre{
		db: db,
		mu: &sync.RWMutex{}}, nil
}
