// router.go (internal/app/interfaces/http/router.go)

package http

import (
	"todoapp/internal/app/usecase"

	"github.com/gofiber/fiber/v2"
)

// SetTodoRoutes sets up Todo routes
func SetTodoRoutes(app *fiber.App, todoUsecase *usecase.TodoUsecase) {
	todoHandler := NewTodoHandler(todoUsecase)

	api := app.Group("/api/v1")

	api.Post("/todos", todoHandler.CreateTodoHandler)
	api.Get("/todos/:id", todoHandler.GetTodoHandler)
	api.Put("/todos/:id", todoHandler.UpdateTodoHandler)
	api.Delete("/todos/:id", todoHandler.DeleteTodoHandler)
}
