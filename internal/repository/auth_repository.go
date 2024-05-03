package repository

import (
	"fmt"

	"github.com/aliflazuardi/cats-social/internal/database"
)

func FindUser() {
	name := "alif"
	rows, err := database.DB.Query("SELECT name FROM users WHERE name = $1", name)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(rows)
}
