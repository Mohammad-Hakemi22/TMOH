package database

import (
	"log"

	"github.com/Mohammad-Hakemi22/tmoh/model"
	"gorm.io/gorm"
)

func InitialMigration() {
	connection, err := GetDatabase()
	if err != nil {
		log.Fatalln("something wrong in database connection", err)
	}
	defer Closedatabase(connection.Conn)
	connection.Conn.AutoMigrate(model.User{})
}

func Closedatabase(connection *gorm.DB) {
	sqldb, err := connection.DB()
	if err != nil {
		log.Fatalln("something wrong in closing database connection", err)
	}
	sqldb.Close()
}
