package companies

import (
	"github.com/MateoCaicedoW/email-sender/internal/app/models"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid/v5"
	"github.com/jmoiron/sqlx"
)

var _ models.CompanyService = (*service)(nil)

type service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *service {
	return &service{db: db}
}

func (s *service) Create(company *models.Company) error {
	query := `INSERT INTO companies (name) VALUES ($1) RETURNING id`
	err := s.db.QueryRow(query, company.Name).Scan(&company.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Validate(company models.Company) *validate.Errors {
	verrs := validate.Validate(
		&validators.StringIsPresent{Field: company.Name, Name: "Name"},
	)

	if company.Name != "" {
		verrs.Append(
			validate.Validate(
				&validators.FuncValidator{
					Name:    "Name",
					Message: "%sCompany with this name already exists.",
					Fn: func() bool {
						c, err := s.FindByName(company.Name)
						return err == nil && c.ID.IsNil()
					},
				},
			),
		)
	}

	return verrs
}

func (s *service) FindByName(name string) (*models.Company, error) {
	query := `SELECT * FROM companies WHERE name = $1`
	company := &models.Company{}
	err := s.db.Get(company, query, name)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}

	return company, nil
}

func (s *service) AddUser(companyID, userID uuid.UUID) error {
	query := `INSERT INTO user_companies (company_id, user_id) VALUES ($1, $2)`
	_, err := s.db.Exec(query, companyID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) FindByID(id uuid.UUID) (models.Company, error) {
	query := `SELECT * FROM companies WHERE id = $1`
	company := models.Company{}
	err := s.db.Get(&company, query, id)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return models.Company{}, err
	}

	return company, nil
}

func (s *service) AllByUser(userID uuid.UUID) (models.Companies, error) {
	query := `
	SELECT 
		companies.*
	FROM
		companies
	JOIN user_companies ON companies.id = user_companies.company_id
	WHERE user_companies.user_id = $1
	`
	companies := models.Companies{}
	err := s.db.Select(&companies, query, userID)
	if err != nil {
		return nil, err
	}

	return companies, nil
}

func (s *service) RemoveUser(companyID, userID uuid.UUID) error {
	query := `DELETE FROM user_companies WHERE company_id = $1 AND user_id = $2`
	_, err := s.db.Exec(query, companyID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) HasUser(companyID, userID uuid.UUID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM user_companies WHERE company_id = $1 AND user_id = $2)`
	var uc bool
	err := s.db.Get(&uc, query, companyID, userID)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return false, err
	}

	return true, nil
}
