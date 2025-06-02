package db

import (
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB() *sqlx.DB {
	dsn := os.Getenv("AUTH_POSTGRES_DSN")
	var db *sqlx.DB
	var err error

	timeout := time.After(30 * time.Second)
	tick := time.Tick(2 * time.Second)

	for {
		select {
		case <-timeout:
			log.Fatalf("could not connect to Postgres (auth): %v", err)
		case <-tick:
			db, err = sqlx.Connect("postgres", dsn)
			if err == nil {
				log.Println("Connected to Auth Postgres")
				return db
			}
			log.Printf("Waiting for auth Postgres: %v", err)
		}
	}
}
