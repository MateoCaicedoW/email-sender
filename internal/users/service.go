package users

import (
	"fmt"

	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid/v5"
	"github.com/jmoiron/sqlx"
)

var _ models.UserService = (*service)(nil)

type service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *service {
	return &service{db: db}
}

func (s *service) FindByEmail(email string) (*models.User, error) {
	query := `SELECT * FROM users WHERE email = $1`
	user := &models.User{}
	err := s.db.Get(user, query, email)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}

	return user, nil
}

func (s *service) Create(user *models.User) error {
	query := `INSERT INTO users (first_name, last_name, email) VALUES ($1, $2, $3) RETURNING id`
	err := s.db.QueryRow(query, user.FirstName, user.LastName, user.Email).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) FindByID(id uuid.UUID) (models.User, error) {
	query := `SELECT * FROM users WHERE id = $1`
	user := models.User{}
	err := s.db.Get(&user, query, id)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return user, err
	}

	return user, nil
}

func (s *service) List(perPage, page int, term string) (models.List, error) {
	query := `

	SELECT 
		users.*
	FROM
		users
	WHERE
		((users.first_name ILIKE '%' || ? || '%') 
	OR 
		(users.last_name ILIKE '%' || $1 || '%') 
	OR 
		(users.email ILIKE '%' || $1 || '%'))
	`
	var users models.Users
	offset := (page - 1) * perPage
	params := []interface{}{term}

	var total int
	count := fmt.Sprintf(`SELECT COUNT(*) FROM (%v) items`, query)
	if err := s.db.Get(&total, s.db.Rebind(count), params...); err != nil {
		return models.List{}, err
	}

	query += ` ORDER BY users.created_at DESC`
	params = append(params, perPage, offset)
	paginated := fmt.Sprintf(`%v LIMIT ? OFFSET ?`, query)

	err := s.db.Select(&users, s.db.Rebind(paginated), params...)
	if err != nil {
		return models.List{}, err
	}

	return models.List{
		Page:    page,
		PerPage: perPage,
		Total:   total,

		Items: users,
	}, nil
}

func (s *service) Validate(user models.User) *validate.Errors {
	verrs := validate.Validate(
		&validators.StringIsPresent{Field: user.FirstName, Name: "First Name"},
		&validators.StringIsPresent{Field: user.LastName, Name: "Last Name"},
		&validators.EmailIsPresent{Field: user.Email, Name: "Email"},
	)

	if user.Email != "" {
		existing, err := s.FindByEmail(user.Email)
		if err != nil {
			return verrs
		}

		if existing.ID != user.ID {
			verrs.Add("email", "User with this email already exists")
		}
	}

	return verrs
}

func (s *service) Update(user models.User) error {
	query := `UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE id = $4`
	_, err := s.db.Exec(query, user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Delete(id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
