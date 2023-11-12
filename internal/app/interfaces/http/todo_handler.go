// todo_handler.go (internal/app/interfaces/http/todo_handler.go)

package http

import (
	"todoapp/internal/app/usecase"

	"github.com/gofiber/fiber/v2"
)

// TodoHandler handles HTTP requests for Todos
type TodoHandler struct {
	todoUsecase *usecase.TodoUsecase
}

// NewTodoHandler creates a new TodoHandler
func NewTodoHandler(todoUsecase *usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{todoUsecase}
}

func (h *TodoHandler) CreateTodoHandler(c *fiber.Ctx) error {
	// Implement logic to create a todo via HTTP
	// Use c.BodyParser(&todo) to parse the incoming JSON into a todo struct
	// Call the appropriate method in the usecase
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *TodoHandler) GetTodoHandler(c *fiber.Ctx) error {
	// Implement logic to get a todo by ID via HTTP
	// Parse the ID from the URL params using c.Params("id")
	// Call the appropriate method in the usecase
	return c.SendStatus(fiber.StatusOK)
}

func (h *TodoHandler) UpdateTodoHandler(c *fiber.Ctx) error {
	// Implement logic to update a todo via HTTP
	// Parse the ID from the URL params using c.Params("id")
	// Use c.BodyParser(&todo) to parse the incoming JSON into a todo struct
	// Call the appropriate method in the usecase
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *TodoHandler) DeleteTodoHandler(c *fiber.Ctx) error {
	// Implement logic to delete a todo by ID via HTTP
	// Parse the ID from the URL params using c.Params("id")
	// Call the appropriate method in the usecase
	return c.SendStatus(fiber.StatusNoContent)
}
