package models

import (
	"time"

	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid/v5"
)

type Subscriber struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CompanyID uuid.UUID `json:"company_id" db:"company_id"`
	Email     string    `json:"email" db:"email"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Active    bool      `json:"active" db:"active"`

	//Timestamps
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdateAt  time.Time `json:"updated_at" db:"updated_at"`
}

type Subscribers []Subscriber

type SubscriberService interface {
	Validate(sub *Subscriber) *validate.Errors
	Create(sub *Subscriber) error
	FindByEmail(email string, companyID uuid.UUID) (*Subscriber, error)
	Find(id uuid.UUID, companyID uuid.UUID) (*Subscriber, error)
	List(perPage, page int, term, status string, companyID uuid.UUID) (List, error)
	Update(sub *Subscriber) error
	Delete(id uuid.UUID, companyID uuid.UUID) error
	ActiveCount(companyID uuid.UUID) (int, error)
	InactiveCount(companyID uuid.UUID) (int, error)
	All(companyID uuid.UUID) (Subscribers, error)
}
