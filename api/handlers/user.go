package handlers

import (
	"github.com/bektosh/fiber-app/api/errors"
	"github.com/bektosh/fiber-app/models/requests"
	"github.com/gofiber/fiber/v2"
)

// CreateUser
// @Summary Register user
// @Tags User
// @Accept json
// @Produce json
// @Param body body requests.CreateUser true "body"
// @Success 200 {object} responses.User
// @Failure 400 {object} errors.ErrorResponse
// @Failure 401 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /user/ [POST]
func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var req requests.CreateUser

	err := c.BodyParser(&req)
	if err != nil {
		h.Logger.Println("error while parsing body:", err)
		errors.AbortWithBadRequest(c, "Invalid JSON supplied")
	}

	user, err, msg := h.service.CreateUser(req)
	if errors.AbortWithError(c, err, msg) {
		return nil
	}

	return c.JSON(user)
}

// GetCompanyUsers
// @Security ApiKeyAuth
// @Summary List all users of the workspace, also have search by name functionality
// @Tags User
// @Accept json
// @Produce json
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Param name query string false "name"
// @Param company-id path string true "company-id"
// @Success 200 {object} responses.WorkspacesResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 401 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /user/{company-id}/ [GET]
//func (h *Handler) GetCompanyUsers(c *fiber.Ctx) error {
//	page, limit, err := helpers.ParsePageAndLimit(c)
//	if err != nil {
//		h.Logger.Println("error while parsing page and limit:", err, page, limit)
//		errors.AbortWithBadRequest(c, "Bad values for page and limit")
//	}
//	companyID := c.Params("company-id")
//	if companyID == "" {
//		errors.AbortWithBadRequest(c, "Path parameter 'company-id' must be provided")
//	}
//	name := c.Query("name")
//}
