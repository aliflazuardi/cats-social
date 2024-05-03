package repository

import (
	"fmt"
	"time"

	"github.com/aliflazuardi/cats-social/internal/database"
	"github.com/aliflazuardi/cats-social/internal/model"
)

const findUserStatement = `SELECT uuid, email, name, password FROM users WHERE email = $1 limit 1`

const insertUserStatement = `INSERT INTO users (uuid, email, name, password, created_at) VALUES ($1, $2, $3, $4, $5)`

func FindUser(email string) (model.User, error) {
	rows, err := database.DB.Query(findUserStatement, email)
	if err != nil {
		fmt.Println(err.Error())
	}
	var u model.User

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&u.UUID, &u.Email, &u.Name, &u.PasswordHash)
		if err != nil {
			return model.User{}, err
		}
	}
	return u, nil
}

func InsertUser(user model.User) error {
	result, err := database.DB.Exec(insertUserStatement, user.UUID, user.Email, user.Name, user.PasswordHash, time.Now())
	if err != nil {
		return err
	}
	fmt.Println(result)

	return nil
}
