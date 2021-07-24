package api

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/bektosh/fiber-app/api/docs"
	"github.com/bektosh/fiber-app/api/handlers"
	"github.com/gofiber/fiber/v2"
)

type Body struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

// SetUpRoutes
// @title Fiber ToDo App
// @version 1.0
// @description This is a swagger for todo-app
// @contact.name API Support
// @contact.email bektosh@novalab.uz
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func SetUpRoutes(app *fiber.App, h *handlers.Handler) {
	// Workspace
	app.Post("/workspace/", h.CreateWorkspace)
	app.Get("/workspace/:user-id/", h.GetWorkspaces)
	app.Patch("/workspace/", h.UpdateWorkspace)
	app.Delete("/workspace/:id/", h.DeleteWorkspace)

	// Company
	app.Post("/company/", h.CreateCompany)

	// User
	app.Post("/user/", h.CreateUser)
	//app.Get("/user/:company-id/", h.GetCompanyUsers)

	// Swagger
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))
}
