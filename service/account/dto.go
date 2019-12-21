package account

import "github.com/google/uuid"

// CreateAccountDTO is a data transfer object.
type CreateAccountDTO struct {
	Email    string    `json:"email" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	Password string    `json:"password" binding:"required"`
	Match    uuid.UUID `json:"match"`
}

// LogInDTO is a data transfer object.
type LogInDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
