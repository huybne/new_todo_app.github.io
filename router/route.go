package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huybne/todo-app-api/handler"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	v1 := api.Group("/todo")
	// routes
	v1.Get("/", handler.GetTodos)
	v1.Get("/:id", handler.GetTodo)
	v1.Post("/", handler.CreateTodo)
	v1.Put("/:id", handler.UpdateTodo)
	v1.Delete("/:id", handler.DeleteTodo)
}
