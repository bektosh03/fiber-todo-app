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
// @host localhost:8000
// @BasePath /
func SetUpRoutes(app *fiber.App, h *handlers.Handler) {
	app.Post("/workspace/", h.CreateWorkspace)
	app.Get("/workspaces/", h.GetWorkspaces)
	app.Patch("/workspace/", h.UpdateWorkspace)
	app.Delete("/workspace/:id/", h.DeleteWorkspace)
	//app.Post("/user/", h.)
	//app.Get("/todos/", h.GetToDos)
	//app.Post("/todos/", h.CreateToDo)
	//app.Delete("/todos/:id/", h.DeleteToDo)
	//app.Put("/todos/:id/", h.UpdateToDo)
	//app.Get("/todos/:id/", h.GetToDoByID)
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))
}
