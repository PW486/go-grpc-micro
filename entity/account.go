package entity

import (
	"time"

	"github.com/google/uuid"
)

// Account entity contains personal infomation.
type Account struct {
	ID        uuid.UUID `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Email     string `gorm:"unique;not null"`
	Name      string `gorm:"unique;not null"`
	Password  []byte `gorm:"not null"`
	Match     *uuid.UUID
}
