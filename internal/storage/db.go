package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"booking/internal/config"
)

var DB *sql.DB

func InitDB(cfg *config.DBConfig) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	
	if err = db.Ping(); err != nil {
		return err
	}
	
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS bookings (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL,
			place_id INTEGER NOT NULL,
			time_from TIMESTAMP NOT NULL,
			time_to TIMESTAMP NOT NULL
		)
	`); err != nil {
		return fmt.Errorf("table creation failed: %w", err)
	}

	DB = db
	return nil
}
