package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aliflazuardi/cats-social/configs"
)

var DB *sql.DB

func ConnectDB() {
	dbConfigs := configs.GetDBConfig()

	connStr := fmt.Sprintf("dbname=%s user=%s password=%s %s", dbConfigs.DBName, dbConfigs.UserName, dbConfigs.Password, dbConfigs.Params)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
