package dto

import (
	"github.com/google/uuid"

	model "github.com/sklyar-vlad/selfDev/internal/model/user"
)

type UserResponse struct {
	UserId   uuid.UUID `json:"user_id"`
	Role     string    `json:"role"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

func ToUserResponse(p *model.User) UserResponse {
	resp := UserResponse{
		UserId:   p.UserId,
		Username: p.Username,
		Email:    p.Email,
	}

	return resp
}
