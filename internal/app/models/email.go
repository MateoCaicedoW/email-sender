package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type Email struct {
	ID      uuid.UUID `json:"id" db:"id"`
	Name    string    `json:"name" db:"name"`
	Message string    `json:"message" db:"message"`
	Sent    bool      `json:"sent" db:"sent"`

	// Timestamps
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
