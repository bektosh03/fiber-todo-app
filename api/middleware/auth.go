package middleware

import (
	"fmt"
	"github.com/bektosh/fiber-app/api/errors"
	"github.com/bektosh/fiber-app/config"
	"github.com/bektosh/fiber-app/pkg/jwt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"time"
)

type JWTRoleAuthorizer struct {
	enforcer   *casbin.Enforcer
	SigningKey []byte
	logger     *log.Logger
}

func NewJWTRoleAuthorizer(cfg config.Config, logger *log.Logger, adapter *gormadapter.Adapter) (*JWTRoleAuthorizer, error) {
	enforcer, err := casbin.NewEnforcer(cfg.PathToCasbinConfFile, adapter)
	if err != nil {
		logger.Println("could not initialize new enforcer:", err)
		return nil, err
	}

	return &JWTRoleAuthorizer{
		enforcer:   enforcer,
		SigningKey: []byte(cfg.JWTSigningKey),
		logger:     logger,
	}, nil
}

func NewAuthorizer(jwtra *JWTRoleAuthorizer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := c.Get("Authorization")

		claims, err := jwt.ExtractClaims(accessToken, jwtra.SigningKey)
		if err != nil {
			jwtra.logger.Println("could not extract claims:", err)
			return err
		} else if exp, ok := claims["exp"].(int64); exp < time.Now().Unix() && ok {
			return c.JSON(errors.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "Access token is expired",
			})
		}

		role := claims["role"].(string)
		fmt.Println(role, c.Path(), c.Method())
		ok, err := jwtra.enforcer.Enforce(role, c.Path(), c.Method())
		if err != nil {
			jwtra.logger.Println("could not enforce:", err)
			return err
		}

		if !ok {
			err = c.SendStatus(http.StatusForbidden)
			if err != nil {
				return err
			}
			return c.JSON(errors.ErrorResponse{
				Code:    http.StatusForbidden,
				Message: "Access Denied",
			})
		}

		return c.Next()
	}
}
