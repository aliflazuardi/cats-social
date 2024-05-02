package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/aliflazuardi/cats-social/configs"
	"github.com/aliflazuardi/cats-social/internal/routes"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func main() {
	fmt.Println("Welcome to Cats Social Application")

	dbConfigs := configs.GetDBConfig()

	connStr := fmt.Sprintf("dbname=%s user=%s password=%s %s", dbConfigs.DBName, dbConfigs.UserName, dbConfigs.Password, dbConfigs.Params)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	name := "alif"
	rows, err := db.Query("SELECT name FROM users WHERE name = $1", name)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(rows)

	mux := routes.GetRoutesHandler()

	if err := http.ListenAndServe("localhost:8080", mux); err != nil { // need to change to 8080 (got port clash now)
		fmt.Println(err.Error())
	}
}
