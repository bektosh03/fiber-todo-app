package handlers

import (
	errs "github.com/bektosh/fiber-app/api/errors"
	"github.com/bektosh/fiber-app/api/helpers"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetToDos(c *fiber.Ctx) error {
	page, limit, err := helpers.ParsePageAndLimit(c)
	h.Logger.Println(page, limit)
	if errs.AbortWithBadRequest(c, err, "Bad values for page or limit") {
		return nil
	}

	return nil
}

func (h *Handler) CreateToDo(c *fiber.Ctx) error {
	return nil
}
