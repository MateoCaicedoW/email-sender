package subscribers

import (
	"database/sql"
	"fmt"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid/v5"
	"github.com/jmoiron/sqlx"
)

var _ models.SubscriberService = (*service)(nil)

type service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *service {
	return &service{db: db}
}

func (s *service) Create(sub *models.Subscriber) error {
	query := `INSERT INTO subs (email, first_name, last_name, active, company_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := s.db.QueryRow(query, sub.Email, sub.FirstName, sub.LastName, sub.Active, sub.CompanyID).Scan(&sub.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) FindByEmail(email string, companyID uuid.UUID) (*models.Subscriber, error) {
	query := `SELECT * FROM subs WHERE email = $1 AND company_id = $2`
	sub := &models.Subscriber{}
	err := s.db.Get(sub, query, email, companyID)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (s *service) Validate(sub *models.Subscriber) *validate.Errors {

	errs := validate.Validate(
		&validators.EmailIsPresent{Field: sub.Email, Name: "Email"},
		&validators.StringIsPresent{Field: sub.FirstName, Name: "First Name"},
		&validators.StringIsPresent{Field: sub.LastName, Name: "Last Name"},
	)

	if sub.Email != "" {
		existing, err := s.FindByEmail(sub.Email, sub.CompanyID)
		if err != nil {
			return errs
		}

		if existing != nil && existing.ID != sub.ID {
			errs.Add("email", "Subscriber with this email already exists")
		}
	}

	return errs
}

func (s *service) List(perPage, page int, term, status string, companyID uuid.UUID) (models.List, error) {
	query := `

	SELECT 
		*
	FROM
		subs
	WHERE 
		(subs.company_id = ?)
	AND
		((subs.first_name ILIKE '%' || ? || '%') 
	OR 
		(subs.last_name ILIKE '%' || $2 || '%') 
	OR 
		(subs.email ILIKE '%' || $2 || '%'))
	`
	var subs models.Subscribers
	offset := (page - 1) * perPage
	params := []interface{}{companyID, term}

	if status != "all" && status != "" {
		query += ` AND subs.active = ?`
		params = append(params, status)
	}

	var total int
	count := fmt.Sprintf(`SELECT COUNT(*) FROM (%v) items`, query)
	if err := s.db.Get(&total, s.db.Rebind(count), params...); err != nil {
		return models.List{}, err
	}

	query += ` ORDER BY subs.created_at DESC`
	params = append(params, perPage, offset)
	paginated := fmt.Sprintf(`%v LIMIT ? OFFSET ?`, query)

	err := s.db.Select(&subs, s.db.Rebind(paginated), params...)
	if err != nil {
		return models.List{}, err
	}

	return models.List{
		Page:    page,
		PerPage: perPage,
		Total:   total,

		Items: subs,
	}, nil
}

func (s *service) Find(id, companyID uuid.UUID) (*models.Subscriber, error) {
	query := `SELECT * FROM subs WHERE id = $1 AND company_id = $2`
	sub := &models.Subscriber{}
	err := s.db.Get(sub, query, id, companyID)
	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		return sub, err
	}

	return sub, nil
}

func (s *service) Update(sub *models.Subscriber) error {
	query := `UPDATE subs SET email = $1, first_name = $2, last_name = $3, active = $4 WHERE id = $5 AND company_id = $6`
	_, err := s.db.Exec(query, sub.Email, sub.FirstName, sub.LastName, sub.Active, sub.ID, sub.CompanyID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Delete(id, companyID uuid.UUID) error {
	query := `DELETE FROM subs WHERE id = $1 AND company_id = $2`
	_, err := s.db.Exec(query, id, companyID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) ActiveCount(companyID uuid.UUID) (int, error) {
	query := `SELECT COUNT(*) FROM subs WHERE active = true AND company_id = $1`
	var count int
	err := s.db.Get(&count, query, companyID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *service) InactiveCount(companyID uuid.UUID) (int, error) {
	query := `SELECT COUNT(*) FROM subs WHERE active = false AND company_id = $1`
	var count int
	err := s.db.Get(&count, query, companyID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *service) All(companyID uuid.UUID) (models.Subscribers, error) {
	query := `SELECT * FROM subs WHERE company_id = $1`
	var subs models.Subscribers
	err := s.db.Select(&subs, query, companyID)
	if err != nil {
		return nil, err
	}

	return subs, nil
}
