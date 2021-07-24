package handlers

import (
	"github.com/bektosh/fiber-app/api/errors"
	"github.com/bektosh/fiber-app/api/helpers"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetToDos(c *fiber.Ctx) error {
	page, limit, err := helpers.ParsePageAndLimit(c)
	if err != nil {
		h.Logger.Println("error while parsing page and limit:", err, page, limit)
		errors.AbortWithBadRequest(c, "Bad values for page and limit")
	}

	return nil
}

func (h *Handler) CreateToDo(c *fiber.Ctx) error {
	return nil
}
