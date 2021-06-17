package errors

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

const (
	InternalMsg = "Internal Server Error"
	NotEnoughRights = "Not Enough Rights"
)

type ErrorResponse struct {
	Code    int
	Message string
}

func AbortWithBadRequest(c *fiber.Ctx, err error, msg string) bool {
	if err == nil {
		return false
	}
	c.Status(http.StatusBadRequest)
	err = c.JSON(ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: msg,
	})
	return true
}

func AbortWithInternal(c *fiber.Ctx, err error, msg string) bool {
	if err == nil {
		return false
	}
	c.Status(http.StatusInternalServerError)
	err = c.JSON(ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: msg,
	})
	return true
}

func AbortWithUnauthorized(c *fiber.Ctx, err error, msg string) bool {
	if err == nil {
		return false
	}
	c.Status(http.StatusUnauthorized)
	err = c.JSON(ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: msg,
	})
	return true
}

