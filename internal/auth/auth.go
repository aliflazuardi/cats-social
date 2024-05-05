package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aliflazuardi/cats-social/internal/model"
	"github.com/aliflazuardi/cats-social/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var u model.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
  
	user, err := repository.FindUser(u.Email)
	if err != nil {
		fmt.Println("error find user in database: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
  
  if user.Email != "" {
    fmt.Println("user already exist")
    w.WriteHeader(http.StatusBadRequest)
    return
  }

	// validate user request email format is correct
	err = validateUser(u)
	if err != nil {
		fmt.Println("user values is incorrect")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u.PasswordHash, err = hashPassword(u.Password)
	if err != nil {
		fmt.Println("error hashing password: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	u.UUID = uuid.New()

	err = repository.InsertUser(u)
	if err != nil {
		fmt.Println("error insert user to database: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func validateUser(u model.User) error {
	if u.Email == "" {
		return fmt.Errorf("email can't be empty")
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}

func Login(w http.ResponseWriter, r *http.Request) {
	var u model.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := repository.FindUser(u.Email)
	if err != nil {
		fmt.Println("error find user in database: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// user not found
	if user.Email == "" {
		fmt.Println("can't find user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check password
	err = compareHashWithPassword(user.PasswordHash, u.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("login failed, password don't match")
		return
	}
	fmt.Println("password matched")
}

func compareHashWithPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
