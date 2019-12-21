package account

import "github.com/google/uuid"

// CreateAccountDTO is awesome
type CreateAccountDTO struct {
	Email    string    `json:"email" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	Password string    `json:"password" binding:"required"`
	Match    uuid.UUID `json:"match"`
}

// LogInDTO is awesome
type LogInDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
