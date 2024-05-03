package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID         uuid.UUID `json:"uuid"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	PasswordHash string    `json:"-"`
	Name         string    `json:"name,omitempty"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
