// main.go

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import PostgreSQL driver

	"todoapp/internal/app/domain"
	"todoapp/internal/app/interfaces/http"
	"todoapp/internal/app/usecase"
)

func main() {
	app := fiber.New()

	// Initialize the database connection
	db, err := sqlx.Open("postgres", "user=youruser password=yourpassword dbname=yourdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize repository and usecase
	todoRepository := domain.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUsecase(todoRepository)

	// Set up Todo routes
	http.SetTodoRoutes(app, todoUsecase)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
