package psql

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/likhithkp/banking-ledger-service/config"
)

var (
	DB   *pgxpool.Pool
	Once sync.Once
)

func InitDB() {
	Once.Do(func() {
		dbURL := config.Config()

		config, err := pgxpool.ParseConfig(dbURL)
		if err != nil {
			log.Fatalf("Failed to parse the dbURL: %v", err)
		}

		config.MaxConns = 10
		config.MinConns = 2
		config.HealthCheckPeriod = 1 * time.Second

		pool, err := pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			log.Fatalf("Failed to create new pool: %v", err)
		}

		DB = pool
		log.Println("Connected to database")
	})
}

func GetDB() {
	if DB == nil {
		InitDB()
	}
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}
