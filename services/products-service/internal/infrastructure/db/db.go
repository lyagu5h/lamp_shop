package infrastructure

import (
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB() *sqlx.DB {
    dsn := os.Getenv("PRODUCTS_POSTGRES_DSN")
    if dsn == "" {
        log.Fatal("PRODUCTS_POSTGRES_DSN must be set")
    }

    var db *sqlx.DB
    var err error

    timeout := time.After(30 * time.Second)
    tick := time.Tick(2 * time.Second)

    for {
        select {
        case <-timeout:
            log.Fatalf("could not connect to products DB: %v", err)
        case <-tick:
            db, err = sqlx.Connect("postgres", dsn)
            if err == nil {
                log.Println("Connected to products PostgreSQL")
                db.SetMaxOpenConns(10)
                return db
            }
            log.Printf("Waiting for products-postgres to be ready: %v", err)
        }
    }
}
