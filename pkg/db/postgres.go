package db

import (
	"fmt"
	"log"
	"time"

	"github.com/R3iwan/soundScroll/pkg/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectPostgres(cfg config.PostgresConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode,
	)

	var db *sqlx.DB
	var err error

	for i := 0; i < 5; i++ { // Retry 5 times with a 2-second delay
		db, err = sqlx.Connect("postgres", dsn)
		if err == nil {
			if err = db.Ping(); err == nil {
				log.Println("Connected to Postgres")
				return db, nil
			}
		}
		log.Printf("Postgres connection failed: %v. Retrying in 2 seconds...", err)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("could not connect to Postgres: %w", err)
}
