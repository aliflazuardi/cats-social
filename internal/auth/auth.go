package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UUID         uuid.UUID `json:"uuid"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	passwordHash string    `json:"-"`
	Name         string    `json:"name,omitempty"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var u User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u.passwordHash, err = hashPassword(u.Password)
	if err != nil {
		fmt.Println("error hashing password: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	u.UUID = uuid.New()

	fmt.Println(u)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}

func compareHashWithPassword(password string) error {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
