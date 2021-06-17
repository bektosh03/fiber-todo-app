package handlers

import (
	"github.com/bektosh/fiber-app/database"
	"github.com/bektosh/fiber-app/service"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

type Handler struct {
	Logger  *log.Logger
	service *service.Service
}

func New(db *sqlx.DB) *Handler {
	logger := log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	return &Handler{
		Logger:  logger,
		service: service.New(database.New(db), logger),
	}
}
