package service

import (
	"github.com/bektosh/fiber-app/database"
	"log"
)

type Service struct {
	storage *database.Storage
	logger  *log.Logger
}

func New(s *database.Storage, l *log.Logger) *Service {
	return &Service{
		storage: s,
		logger:  l,
	}
}
