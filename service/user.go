package service

import (
	"github.com/bektosh/fiber-app/api/errors"
	"github.com/bektosh/fiber-app/models/requests"
	"github.com/bektosh/fiber-app/models/responses"
	"github.com/bektosh/fiber-app/pkg/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func (s *Service) CreateUser(req requests.CreateUser) (responses.User, error, string) {
	req.Password = strings.TrimSpace(req.Password)
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	if ok := req.Validate(); !ok {
		return responses.User{}, errors.BadRequest, "Not valid email or password (must contain at least 8 characters," +
			"1 uppercase letter, 1 lowercase letter and a symbol)"
	}

	exists, err := s.storage.Psql.CheckUserEmail(req.Email)
	if err != nil {
		s.logger.Println("error while checking email:", err)
		return responses.User{}, errors.Internal, ""
	} else if exists {
		return responses.User{}, errors.AlreadyExists, "User with this email already exists"
	}

	id, err := uuid.NewUUID()
	if err != nil {
		s.logger.Println("Error while generating UUID:", err)
		return responses.User{}, errors.Internal, ""
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.FirstName = strings.TrimSpace(req.FirstName)
	req.LastName = strings.TrimSpace(req.LastName)

	accessToken, err := jwt.GenerateAccessJWT(id.String(), false, false, []byte(s.config.JWTSigningKey))
	if err != nil {
		s.logger.Println("error while generating access token:", err)
		return responses.User{}, errors.Internal, ""
	}
	refreshToken, err := jwt.GenerateRefreshJWT(id.String(), false, false, []byte(s.config.JWTSigningKey))
	if err != nil {
		s.logger.Println("error while generating refresh token:", err)
		return responses.User{}, errors.Internal, ""
	}

	user, err := s.storage.Psql.CreateUser(responses.User{
		ID:           id.String(),
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		Age:          req.Age,
		Gender:       req.Gender,
		ProfilePhoto: req.ProfilePhoto,
		Username:     req.Username,
		CompanyID:    req.CompanyID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, string(hash))
	if err != nil {
		s.logger.Println("error while creating user:", err)
		return responses.User{}, errors.Internal, ""
	}

	return user, nil, ""
}
