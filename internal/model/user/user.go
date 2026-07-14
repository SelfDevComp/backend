package model

import (
	"github.com/google/uuid"
)

type User struct {
	UserId   uuid.UUID
	Sub      string
	Username string
	Email    string
}

func NewUser(sub, username, email string) User {
	return User{uuid.New(), sub, username, email}
}
