// todo_usecase.go (internal/app/usecase/todo_usecase.go)

package usecase

import (
	"todoapp/internal/app/domain"
)

// TodoUsecase handles business logic related to Todos
type TodoUsecase struct {
	todoRepo *domain.TodoRepository
}

// NewTodoUsecase creates a new TodoUsecase
func NewTodoUsecase(todoRepo *domain.TodoRepository) *TodoUsecase {
	return &TodoUsecase{todoRepo}
}

func (uc *TodoUsecase) CreateTodo(todo *domain.Todo) error {
	// Implement logic to create a todo
	return uc.todoRepo.Create(todo)
}

func (uc *TodoUsecase) GetTodoByID(id int) (*domain.Todo, error) {
	// Implement logic to get a todo by ID
	return uc.todoRepo.GetByID(id)
}

func (uc *TodoUsecase) UpdateTodo(todo *domain.Todo) error {
	// Implement logic to update a todo
	return uc.todoRepo.Update(todo)
}

func (uc *TodoUsecase) DeleteTodo(id int) error {
	// Implement logic to delete a todo by ID
	return uc.todoRepo.Delete(id)
}
