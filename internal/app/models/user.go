package models

import (
	"time"

	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid/v5"
)

type User struct {
	ID        uuid.UUID `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Users []User

type UserService interface {
	FindByEmail(email string) (*User, error)
	Create(user *User) error
	ValidateRegister(user User) *validate.Errors
	FirstCompany(id uuid.UUID) (Company, error)
	FindByID(id uuid.UUID) (User, error)
	List(limit, page int, term string, companyID uuid.UUID) (List, error)
	Validate(user User, companyID uuid.UUID) *validate.Errors
	Update(user User) error
}
