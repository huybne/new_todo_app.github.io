package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/huybne/todo-app-api/database"
	"github.com/huybne/todo-app-api/model"
)

//	var todos = []*Todo{
//		{},
//	}
var Todos []model.Todo

func GetTodos(c *fiber.Ctx) error {
	db := database.DB.Db
	var todos []model.Todo
	db.Find(&todos)
	if len(todos) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"Status":  "error",
			"message": "Todos not found",
			"data":    "nil"})
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": " All Todos found",
		"data":    todos})

}

func GetTodo(c *fiber.Ctx) error {
	db := database.DB.Db
	var Todo model.Todo
	id := c.Params("id")
	db.Find(&Todo, "id = ?", id)
	if Todo.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Todo not found",
			"data":    nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success",
		"message": "Todo Found",
		"data":    Todo})
}
func CreateTodo(c *fiber.Ctx) error {

	db := database.DB.Db
	todo := new(model.Todo)

	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Something's wrong with your input",
			"data":    err})
	}
	err = db.Create(&todo).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create todo",
			"data":    err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success",
		"message": "Todo has created",
		"data":    todo})
}
func UpdateTodo(c *fiber.Ctx) error {
	type UpdateTodo struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	db := database.DB.Db
	var todo model.Todo
	// get id params
	id := c.Params("id")
	// find single todo in the database by id
	db.First(&todo, "id = ?", id)
	if todo.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Todo not found",
			"data":    nil})
	}
	var updateTodoData UpdateTodo
	err := c.BodyParser(&updateTodoData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Something's wrong with your input",
			"data":    err})
	}
	todo.Title = updateTodoData.Title
	todo.Completed = updateTodoData.Completed
	// Save the Changes
	db.Save(&todo)
	// Return the updated todo
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Todo Found",
		"data":    todo})
}

func DeleteTodo(c *fiber.Ctx) error {
	db := database.DB.Db
	var todo model.Todo
	id := c.Params("id")
	db.Find(&todo, "id = ?", id)
	if todo.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
			"data":    nil})
	}
	err := db.Delete(&todo, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to delete user",
			"data":    nil})
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Todo deleted"})
}
