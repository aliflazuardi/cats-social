package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/aliflazuardi/cats-social/internal/database"
	"github.com/aliflazuardi/cats-social/internal/repository"
	"github.com/aliflazuardi/cats-social/internal/routes"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func main() {
	fmt.Println("Welcome to Cats Social Application")

	database.ConnectDB()
	repository.Init()

	mux := routes.GetRoutesHandler()

	if err := http.ListenAndServe("localhost:8080", mux); err != nil { // need to change to 8080 (got port clash now)
		fmt.Println(err.Error())
	}
}
