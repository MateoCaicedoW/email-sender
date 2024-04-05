package emails

import (
	"fmt"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/jmoiron/sqlx"
)

var _ models.EmailService = (*service)(nil)

type service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *service {
	return &service{db: db}
}

func (s *service) Create(e *models.Email) error {
	_, err := s.db.NamedExec(`INSERT INTO emails (name, message, sent, subject, attachment_name, attachment_content, scheduled, scheduled_at) VALUES (:name, :message, :sent, :subject, :attachment_name, :attachment_content, :scheduled, :scheduled_at)`, e)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) List(perPage, page int, term, status string) (models.List, error) {
	query := `

	SELECT 
		*
	FROM
		emails
	WHERE 
		(emails.name ILIKE '%' || ? || '%')
	`
	var emails models.Emails
	offset := (page - 1) * perPage
	params := []interface{}{term}

	if status != "all" && status != "" {
		query += ` AND emails.sent = ?`
		params = append(params, status)
	}

	var total int
	count := fmt.Sprintf(`SELECT COUNT(*) FROM (%v) items`, query)
	if err := s.db.Get(&total, s.db.Rebind(count), params...); err != nil {
		return models.List{}, err
	}

	query += ` ORDER BY emails.created_at DESC`
	params = append(params, perPage, offset)
	paginated := fmt.Sprintf(`%v LIMIT ? OFFSET ?`, query)

	err := s.db.Select(&emails, s.db.Rebind(paginated), params...)
	if err != nil {
		return models.List{}, err
	}

	return models.List{
		Page:    page,
		PerPage: perPage,
		Total:   total,

		Items: emails,
	}, nil
}

func (s *service) Scheduled() (models.Emails, error) {
	var emails models.Emails
	err := s.db.Select(&emails, `SELECT * FROM emails WHERE scheduled = true AND sent = false AND scheduled_at <= NOW()`)
	if err != nil {
		return nil, err
	}

	return emails, nil
}

func (s *service) Update(e *models.Email) error {
	_, err := s.db.NamedExec(`UPDATE emails SET name = :name, message = :message, sent = :sent, subject = :subject, attachment_name = :attachment_name, attachment_content = :attachment_content, scheduled = :scheduled, scheduled_at = :scheduled_at WHERE id = :id`, e)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Validate(e *models.Email) *validate.Errors {
	verrs := validate.Validate(
		&validators.StringIsPresent{Field: e.Name, Name: "Name"},
		&validators.StringIsPresent{Field: e.Message, Name: "Message"},
		&validators.StringIsPresent{Field: e.Subject, Name: "Subject"},
	)

	if e.Scheduled {
		verrs.Append(validate.Validate(
			&validators.TimeIsPresent{Field: e.ScheduledAt, Name: "ScheduledAt", Message: "This field is required when scheduling an email."},
		))
	}

	return verrs
}
