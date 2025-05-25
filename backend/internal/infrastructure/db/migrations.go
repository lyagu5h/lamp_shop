package infrastructure

import (
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

func RunMigration(db *sqlx.DB, dir string) error {
	goose.SetDialect("postgres")
	sqlDB := db.DB
	return goose.Up(sqlDB, dir)
}