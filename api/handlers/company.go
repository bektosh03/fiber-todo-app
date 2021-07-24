package handlers

import (
	"github.com/bektosh/fiber-app/api/errors"
	"github.com/bektosh/fiber-app/models/requests"
	"github.com/gofiber/fiber/v2"
)

// CreateCompany
// @Security ApiKeyAuth
// @Summary Create a company
// @Tags Company
// @Accept json
// @Produce json
// @Param body body requests.CreateCompany true "body"
// @Success 200 {object} responses.CompanyResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 401 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /company/ [POST]
func (h *Handler) CreateCompany(c *fiber.Ctx) error {
	var body requests.CreateCompany

	err := c.BodyParser(&body)
	if err != nil {
		h.Logger.Println("error while parsing body:", err, string(c.Request().Body()))
		errors.AbortWithBadRequest(c, "Invalid JSON supplied")
		return nil
	}

	claims, err := h.getClaims(c)
	if err != nil {
		h.Logger.Println("error while extracting claims:", err)
		errors.AbortWithUnauthorized(c, err.Error())
		return nil
	}

	res, err, msg := h.service.CreateCompany(body, claims["sub"].(string))
	if errors.AbortWithError(c, err, msg) {
		return nil
	}

	return c.JSON(res)
}
