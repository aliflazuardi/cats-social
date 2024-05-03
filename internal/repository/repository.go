package repository

import (
	"database/sql"

	"github.com/aliflazuardi/cats-social/internal/database"
)

var db *sql.DB

func Init() {
	db = database.DB
}
