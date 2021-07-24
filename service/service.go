package service

import (
	"github.com/bektosh/fiber-app/config"
	"github.com/bektosh/fiber-app/database"
	"log"
)

type Service struct {
	storage *database.Storage
	logger  *log.Logger
	config  *config.Config
}

func New(s *database.Storage, l *log.Logger, cfg *config.Config) *Service {
	return &Service{
		storage: s,
		logger:  l,
		config: cfg,
	}
}
