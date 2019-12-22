package account

import (
	"time"

	"github.com/google/uuid"
)

// CreateAccountDTO is a data transfer object.
type CreateAccountDTO struct {
	Email    string     `json:"email" binding:"required"`
	Name     string     `json:"name" binding:"required"`
	Password string     `json:"password" binding:"required"`
	MatchID  *uuid.UUID `json:"matchId"`
}

// LogInDTO is a data transfer object.
type LogInDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// FindAccountResponse is a reponse object.
type FindAccountResponse struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	Email     string     `json:"email"`
	Name      string     `json:"name"`
	MatchID   *uuid.UUID `json:"matchId"`
}
