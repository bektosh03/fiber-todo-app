package postgres

import (
	"github.com/bektosh/fiber-app/api/errors"
	"github.com/bektosh/fiber-app/models/responses"
)

func (p *Postgres) CheckCompanyEmail(email string) (bool, error) {
	var exists bool

	err := p.db.Get(
		&exists,
		`SELECT count(1) = 1 FROM companies WHERE email = $1`,
		email,
	)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (p *Postgres) CreateCompany(company responses.CompanyResponse) (responses.CompanyResponse, error) {
	result, err := p.db.Exec(
		`INSERT INTO companies (id, name, email, owner) VALUES ($1, $2, $3, $4)`,
		company.ID, company.Name, company.Email, company.Owner,
	)

	if err != nil {
		return responses.CompanyResponse{}, err
	} else if n, _ := result.RowsAffected(); n == 0 {
		return responses.CompanyResponse{}, errors.AlreadyExists
	}
	return company, nil
}
