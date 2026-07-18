package model

import (
	"github.com/google/uuid"
)

type User struct {
	UserId   uuid.UUID
	Sub      string
	Username string
}

func NewUser(sub, username string) User {
	return User{uuid.New(), sub, username}
}
