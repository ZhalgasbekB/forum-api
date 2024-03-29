package driver

import (
	"context"
	"database/sql"
	"log"

	"gitea.com/lzhuk/forum/pkg/config"
	_ "github.com/mattn/go-sqlite3"
)

func NewDB(ctx context.Context, cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open(cfg.Db.Driver, cfg.Db.Dsn)
	if err != nil {
		log.Println("Ошибка инциализации базы данных %w", err)
		return nil, err
	}
	return db, nil
}
