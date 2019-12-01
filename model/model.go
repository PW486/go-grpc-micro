package model

import (
	"time"

	"github.com/google/uuid"
)

// Account is main model
type Account struct {
	ID        uuid.UUID `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Email     string `json:"email" binding:"required" gorm:"unique;not null"`
	Name      string `json:"name" binding:"required" gorm:"unique;not null"`
	Password  string
	Match     uuid.UUID
}
