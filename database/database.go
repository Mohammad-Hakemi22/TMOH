package database

import (
	"fmt"

	"github.com/Mohammad-Hakemi22/tmoh/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	Conn *gorm.DB
}


func GetDatabase() (Connection, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s timezone=%s sslmode=disable",
		config.AppConfig.DB_HOST,
		config.AppConfig.DB_USER,
		config.AppConfig.DB_PASSWORD,
		config.AppConfig.DB_NAME,
		config.AppConfig.DB_PORT,
		config.AppConfig.DB_TIMEZONE,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return Connection{Conn: nil}, err
	}
	return Connection{Conn: db}, nil
}
