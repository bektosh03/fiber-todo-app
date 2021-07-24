package errors

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	NotFound      = Error("Not Found")
	AlreadyExists = Error("Already Exists")
	Internal      = Error("Internal Server Error")
	Forbidden     = Error("Forbidden")
	BadRequest    = Error("Bad Request")
	Unauthorized  = Error("Unauthorized")
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func AbortWithError(c *fiber.Ctx, err error, msg ...string) bool {
	switch err {
	case nil:
		return false
	case Internal:
		AbortWithInternal(c, msg[0])
	case BadRequest:
		AbortWithBadRequest(c, msg[0])
	case AlreadyExists:
		AbortWithBadRequest(c, msg[0])
	case NotFound:
		AbortWithNotFound(c, msg[0])
	case Unauthorized:
		AbortWithUnauthorized(c, msg[0])
	default:
		AbortWithInternal(c, msg[0])
	}
	return true
}

func AbortWithBadRequest(c *fiber.Ctx, msg ...string) bool {
	if len(msg) == 0 || len(msg[0]) == 0 {
		msg[0] = BadRequest.Error()
	}
	c.Status(http.StatusBadRequest)
	_ = c.JSON(ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: msg[0],
	})
	return true
}

func AbortWithInternal(c *fiber.Ctx, msg ...string) bool {
	if len(msg) == 0 || len(msg[0]) == 0 {
		msg[0] = Internal.Error()
	}
	c.Status(http.StatusInternalServerError)
	_ = c.JSON(ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: msg[0],
	})
	return true
}

func AbortWithUnauthorized(c *fiber.Ctx, msg ...string) bool {
	if len(msg) == 0 || len(msg[0]) == 0 {
		msg[0] = Unauthorized.Error()
	}
	c.Status(http.StatusUnauthorized)
	_ = c.JSON(ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: msg[0],
	})
	return true
}

func AbortWithNotFound(c *fiber.Ctx, msg ...string) bool {
	if len(msg) == 0 || len(msg[0]) == 0 {
		msg[0] = NotFound.Error()
	}
	c.Status(http.StatusNotFound)
	_ = c.JSON(ErrorResponse{
		Code:    http.StatusNotFound,
		Message: msg[0],
	})
	return true
}
