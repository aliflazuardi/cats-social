package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBName   string
	Host     string
	Port     string
	UserName string
	Password string
	Params   string
}

func GetDBConfig() *DBConfig {
	err := godotenv.Load("../../configs/.env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err.Error())
	}

	dbConfig := DBConfig{
		DBName:   os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		UserName: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Params:   os.Getenv("DB_PARAMS"),
	}

	return &dbConfig
}
