package handlers

import (
	"github.com/bektosh/fiber-app/api/errors"
	"github.com/bektosh/fiber-app/api/helpers"
	"github.com/bektosh/fiber-app/models/requests"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// CreateWorkspace
// @Summary Creates a Workspace
// @Tags Workspaces
// @Accept json
// @Produce json
// @Param body body requests.CreateWorkspace true "body"
// @Success 200 {object} responses.Workspace
// @Failure 400 {object} errors.ErrorResponse
// @Failure 401 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /workspace/ [POST]
func (h *Handler) CreateWorkspace(c *fiber.Ctx) error {
	var req requests.CreateWorkspace

	err := c.BodyParser(&req)
	if errors.AbortWithBadRequest(c, err, "Invalid JSON supplied") {
		return nil
	}

	res, err := h.service.CreateWorkspace(req.Name)
	if errors.AbortWithInternal(c, err, errors.InternalMsg) {
		return nil
	}

	c.Status(http.StatusOK)
	return c.JSON(res)
}

// GetWorkspaces
// @Summary lists all workspaces
// @Tags Workspaces
// @Accept json
// @Produce json
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Param name query string false "name"
// @Success 200 {object} responses.WorkspacesResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 401 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /workspaces/ [GET]
func (h *Handler) GetWorkspaces(c *fiber.Ctx) error {
	page, limit, err := helpers.ParsePageAndLimit(c)
	if errors.AbortWithBadRequest(c, err, "Bad values for page or limit") {
		return nil
	}

	name := c.Query("name")

	response, err := h.service.GetWorkspaces(page, limit, name)
	if errors.AbortWithInternal(c, err, errors.InternalMsg) {
		return nil
	}

	return c.JSON(response)
}

// UpdateWorkspace
// @Summary updates a workspace
// @Tags Workspaces
// @Accept json
// @Produce json
// @Param body body requests.UpdateWorkspace true "body"
// @Success 200 {object} responses.Workspace
// @Failure 400 {object} errors.ErrorResponse
// @Failure 401 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /workspace/ [PATCH]
func (h *Handler) UpdateWorkspace(c *fiber.Ctx) error {
	var req requests.UpdateWorkspace

	err := c.BodyParser(&req)
	if errors.AbortWithBadRequest(c, err, "Invalid JSON supplied") {
		return nil
	}

	res, err := h.service.UpdateWorkspace(req)
	if errors.AbortWithInternal(c, err, errors.InternalMsg) {
		return nil
	}

	return c.JSON(res)
}

// DeleteWorkspace
// @Summary deletes a workspace
// @Tags Workspaces
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200
// @Failure 400 {object} errors.ErrorResponse
// @Failure 401 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /workspace/{id}/ [DELETE]
func (h *Handler) DeleteWorkspace(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusBadRequest)
		return c.JSON(errors.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "parameter 'id' not supplied",
		})
	}
	err := h.service.DeleteWorkspace(id)
	if errors.AbortWithInternal(c, err, errors.InternalMsg) {
		return nil
	}
	return nil
}
