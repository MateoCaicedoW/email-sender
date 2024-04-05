package models

import (
	"time"

	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid/v5"
)

type Email struct {
	ID      uuid.UUID `json:"id" db:"id"`
	Name    string    `json:"name" db:"name"`
	Message string    `json:"message" db:"message"`
	Sent    bool      `json:"sent" db:"sent"`
	Subject string    `json:"subject" db:"subject"`

	// Attachment Attributes
	AttachmentName    string `db:"attachment_name"`
	AttachmentContent []byte `db:"attachment_content"`

	// Scheduled Attributes
	Scheduled   bool      `json:"scheduled" db:"scheduled"`
	ScheduledAt time.Time `json:"scheduled_at" db:"scheduled_at"`

	// Timestamps
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Emails []Email

// EmailService provides the interface for the email service
type EmailService interface {
	Create(e *Email) error
	List(perPage, page int, term, status string) (List, error)
	Scheduled() (Emails, error)
	Update(e *Email) error
	Validate(e *Email) *validate.Errors
}
