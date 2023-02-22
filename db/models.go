package db

import "github.com/google/uuid"

type user struct {
	Name     string    `json:"name"`
	ID       uuid.UUID `json:"id"`
	Password byte      `json:"password"`
	Email    string    `json:"email"`
}
