package config

import (
	"log"
	"os"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

type Config struct{
	DBURL string
}

func Load() *Config{
	err:=godotenv.Load()
	if err!=nil{
		log.Println("No env variable detected")
	}

	cfg:=&Config{
		DBURL: os.Getenv("DB_URL"),
	}

	if cfg.DBURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	return cfg
}


// package database

// import (
// 	"database/sql"
// 	"log"
// 	"mini_search_engine/config"

// 	_ "github.com/lib/pq"
// )

// func Initializer(cfg *config.Config) *sql.DB {
// 	db, err := sql.Open("postgres", cfg.DATABASE_URL)
// 	if err != nil {
// 		log.Fatalf("Failed to open database: %v", err)
// 	}

// 	if err := db.Ping(); err != nil {
// 		log.Fatalf("Failed to connect to Neon: %v", err)
// 	}

// 	log.Println("Neon PostgreSQL connected successfully")
// 	return db
// }