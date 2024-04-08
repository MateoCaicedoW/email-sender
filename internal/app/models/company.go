package models

import (
	"time"

	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid/v5"
)

type Company struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`

	// CreatedAt and UpdatedAt are the timestamps for when the record was created and last updated
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Companies []Company

// CompanyService provides the interface for the company service
type CompanyService interface {
	Create(c *Company) error
	Validate(c Company) *validate.Errors
	FindByName(name string) (*Company, error)
	AddUser(companyID, userID uuid.UUID) error
	FindByID(id uuid.UUID) (Company, error)
	AllByUser(userID uuid.UUID) (Companies, error)
	RemoveUser(companyID, userID uuid.UUID) error
	HasUser(companyID, userID uuid.UUID) (bool, error)
}
