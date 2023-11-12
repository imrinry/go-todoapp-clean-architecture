// todo_repository.go (internal/app/domain/todo_repository.go)

package domain

import (
	"github.com/jmoiron/sqlx"
)

// TodoRepository handles interactions with the Todo database
type TodoRepository struct {
	db *sqlx.DB
}

// NewTodoRepository creates a new TodoRepository
func NewTodoRepository(db *sqlx.DB) *TodoRepository {
	return &TodoRepository{db}
}

func (r *TodoRepository) Create(todo *Todo) error {
	// Implement logic to create a todo in the database
	_, err := r.db.Exec("INSERT INTO todos (title, completed) VALUES ($1, $2)", todo.Title, todo.Completed)
	return err
}

func (r *TodoRepository) GetByID(id int) (*Todo, error) {
	// Implement logic to retrieve a todo from the database by ID
	var todo Todo
	err := r.db.Get(&todo, "SELECT id, title, completed FROM todos WHERE id = $1", id)
	return &todo, err
}

func (r *TodoRepository) Update(todo *Todo) error {
	// Implement logic to update a todo in the database
	_, err := r.db.Exec("UPDATE todos SET title = $1, completed = $2 WHERE id = $3", todo.Title, todo.Completed, todo.ID)
	return err
}

func (r *TodoRepository) Delete(id int) error {
	// Implement logic to delete a todo from the database by ID
	_, err := r.db.Exec("DELETE FROM todos WHERE id = $1", id)
	return err
}
