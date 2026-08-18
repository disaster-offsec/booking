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
	
	DB = db
	return nil
}
