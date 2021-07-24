package handlers

import (
	"github.com/bektosh/fiber-app/config"
	"github.com/bektosh/fiber-app/database"
	"github.com/bektosh/fiber-app/pkg/jwt"
	"github.com/bektosh/fiber-app/service"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

type Handler struct {
	Logger  *log.Logger
	service *service.Service
	cfg     *config.Config
}

func New(db *sqlx.DB, cfg *config.Config) *Handler {
	logger := log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	return &Handler{
		Logger:  logger,
		service: service.New(database.New(db), logger, cfg),
		cfg:     cfg,
	}
}

func (h *Handler) getClaims(c *fiber.Ctx) (jwtgo.MapClaims, error) {
	token := c.Get("Authorization")
	claims, err := jwt.ExtractClaims(token, []byte(h.cfg.JWTSigningKey))
	if err != nil {
		return nil, err
	}
	return claims, nil
}
