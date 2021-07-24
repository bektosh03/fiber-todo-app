package service

import (
	"github.com/bektosh/fiber-app/api/errors"
	"github.com/bektosh/fiber-app/api/validators"
	"github.com/bektosh/fiber-app/models/requests"
	"github.com/bektosh/fiber-app/models/responses"
	"github.com/bektosh/fiber-app/pkg/jwt"
	"github.com/google/uuid"
	"strings"
)

func (s *Service) CreateCompany(req requests.CreateCompany, ownerID string) (responses.CompanyResponse, error, string) {
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	exists, err := s.storage.Psql.CheckCompanyEmail(req.Email)
	if err != nil {
		s.logger.Println("error while checking company email existence:", err)
		return responses.CompanyResponse{}, errors.Internal, ""
	} else if exists {
		return responses.CompanyResponse{}, errors.AlreadyExists, "Company with this email already exists"
	}
	id, err := uuid.NewRandom()
	if err != nil {
		s.logger.Println("error while generating random UUID:", err)
		return responses.CompanyResponse{}, errors.Internal, ""
	}
	if ok := validators.ValidateEmail(req.Email); !ok {
		return responses.CompanyResponse{}, errors.BadRequest, "Bad value for email"
	}
	req.Name = strings.TrimSpace(req.Name)

	res, err := s.storage.Psql.CreateCompany(responses.CompanyResponse{
		ID:    id.String(),
		Name:  req.Name,
		Email: req.Email,
		Owner: ownerID,
	})
	if err != nil {
		s.logger.Println("error while creating company:", err)
		return responses.CompanyResponse{}, errors.Internal, ""
	}

	res.AccessToken, err = jwt.GenerateAccessJWT(ownerID, false, true, []byte(s.config.JWTSigningKey))
	if err != nil {
		s.logger.Println("error while generating access token:", err)
		return responses.CompanyResponse{}, errors.Internal, ""
	}
	res.RefreshToken, err = jwt.GenerateRefreshJWT(ownerID, false, true, []byte(s.config.JWTSigningKey))
	if err != nil {
		s.logger.Println("error while generating refresh token:", err)
		return responses.CompanyResponse{}, errors.Internal, ""
	}

	return res, nil, ""
}
