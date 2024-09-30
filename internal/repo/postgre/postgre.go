package postgre

import (
	"database/sql"
	"os"
	"sync"

	"github.com/meraiku/music_lib/internal/repo"
	"github.com/meraiku/music_lib/pkg/logging"
	migrations "github.com/meraiku/music_lib/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var _ repo.MusicRepository = (*postgre)(nil)

var dsn string

type postgre struct {
	db  *bun.DB
	mu  *sync.RWMutex
	log *logging.Logger
}

func New(logger *logging.Logger) (*postgre, error) {
	dsn = os.Getenv("POSTGRES_DSN")
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	if err := sqldb.Ping(); err != nil {
		return nil, err
	}

	if err := migrations.Songs(sqldb); err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, pgdialect.New())

	return &postgre{
		db:  db,
		mu:  &sync.RWMutex{},
		log: logger,
	}, nil
}
