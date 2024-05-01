package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Welcome to Cats Social Application")

	connStr := "dbname=cats_social_db user=admin password=admin123 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	age := 21
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(rows)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Cats")
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil { // need to change to 8080 (got port clash now)
		fmt.Println(err.Error())
	}
}
