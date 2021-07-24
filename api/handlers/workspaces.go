package handlers

import (
	"fmt"
	"github.com/bektosh/fiber-app/api/errors"
	"github.com/bektosh/fiber-app/api/helpers"
	"github.com/bektosh/fiber-app/models/requests"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Filters struct {
	Page  uint64 `json:"page"`
	Limit uint64 `json:"limit"`
	Name  string `json:"name"`
}

// CreateWorkspace
// @Security ApiKeyAuth
// @Summary Create a workspace
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
	if err != nil {
		h.Logger.Println("error while parsing body:", err)
		errors.AbortWithBadRequest(c, "Invalid JSON supplied")
	}

	res, err := h.service.CreateWorkspace(req.Name)
	if errors.AbortWithError(c, err) {
		return nil
	}

	c.Status(http.StatusOK)
	return c.JSON(res)
}

// GetWorkspaces
// @Security ApiKeyAuth
// @Summary List all workspaces of the user
// @Tags Workspaces
// @Accept json
// @Produce json
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Param name query string false "name"
// @Param user-id path string true "user-id"
// @Success 200 {object} responses.WorkspacesResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 401 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /workspace/{user-id}/ [GET]
func (h *Handler) GetWorkspaces(c *fiber.Ctx) error {
	var filters Filters
	err := c.QueryParser(&filters)
	fmt.Println(filters)
	page, limit, err := helpers.ParsePageAndLimit(c)
	if err != nil {
		h.Logger.Println("error while parsing page and limit:", err, page, limit)
		errors.AbortWithBadRequest(c, "Bad values for page and limit")
	}

	userID := c.Params("user-id")
	if userID == "" {
		errors.AbortWithBadRequest(c, "Path parameter 'user-id' must be provided")
	}
	name := c.Query("name")

	response, err := h.service.GetWorkspaces(page, limit, name)
	if errors.AbortWithError(c, err) {
		return nil
	}

	return c.JSON(response)
}

// UpdateWorkspace
// @Security ApiKeyAuth
// @Summary Update a workspace
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
	if err != nil {
		h.Logger.Println("error while parsing body:", err)
		errors.AbortWithBadRequest(c, "Invalid JSON supplied")
	}

	res, err := h.service.UpdateWorkspace(req)
	if errors.AbortWithError(c, err) {
		return nil
	}

	return c.JSON(res)
}

// DeleteWorkspace
// @Security ApiKeyAuth
// @Summary Delete a workspace
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
	if errors.AbortWithError(c, err) {
		return nil
	}
	return nil
}
