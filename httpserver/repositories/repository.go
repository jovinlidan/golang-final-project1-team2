package repositories

import "golang-final-project1-team2/httpserver/repositories/models"

type TodoRepo interface {
	CreateTodo(todo *models.Todo) error
	GetTodos() (*[]models.Todo, error)
	GetTodoByID(id int) (*models.Todo, error)
	UpdateTodo(id int, todo *models.Todo) error
	DeleteTodo(id int) error
}
