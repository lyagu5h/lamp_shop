package infrastructure

import (
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func NewPostgresDB() *sqlx.DB {
    dsn := os.Getenv("ORDERS_POSTGRES_DSN")
    if dsn == "" {
        log.Fatal("ORDERS_POSTGRES_DSN must be set")
    }

    var db *sqlx.DB
    var err error

    timeout := time.After(30 * time.Second)
    tick := time.Tick(2 * time.Second)

    for {
        select {
        case <-timeout:
            log.Fatalf("could not connect to orders DB: %v", err)
        case <-tick:
            db, err = sqlx.Connect("postgres", dsn)
            if err == nil {
                log.Println("Connected to orders PostgreSQL")
                db.SetMaxOpenConns(10)
                return db
            }
            log.Printf("Waiting for orders-postgres to be ready: %v", err)
        }
    }
}

func RunMigration(db *sqlx.DB, dir string) error {
    goose.SetDialect("postgres")
    sqlDB := db.DB
    return goose.Up(sqlDB, dir)
}