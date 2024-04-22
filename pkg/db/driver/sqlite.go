package driver

import (
	"database/sql"
	"gitea.com/lzhuk/forum/pkg/config"
	_ "github.com/mattn/go-sqlite3"
)

func NewDB(cfg config.Config) (*sql.DB, error) { // DB
	db, err := sql.Open(cfg.Db.Driver, cfg.Db.Dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
