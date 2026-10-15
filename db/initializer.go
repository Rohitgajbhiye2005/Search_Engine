package db

import (
	"database/sql"
	"log"
	"mini_search_engine/config"
	_ "github.com/lib/pq"
)





func Initializer(cfg *config.Config) *sql.DB {
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to Neon: %v", err)
	}

	log.Println("Neon PostgreSQL connected successfully")
	return db
}