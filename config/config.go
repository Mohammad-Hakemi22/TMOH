package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	Port        string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_TIMEZONE string
}

var AppConfig config

func SetConfig() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatalln(err)
	}
	AppConfig.Port = os.Getenv("Port")
	AppConfig.DB_HOST = os.Getenv("DB_HOST")
	AppConfig.DB_PORT = os.Getenv("DB_PORT")
	AppConfig.DB_USER = os.Getenv("DB_USER")
	AppConfig.DB_NAME = os.Getenv("DB_NAME")
	AppConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	AppConfig.DB_TIMEZONE = os.Getenv("DB_TIMEZONE")
}
