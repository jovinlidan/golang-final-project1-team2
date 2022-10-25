package gorm

import (
	"golang-final-project1-team2/httpserver/repositories"
	"golang-final-project1-team2/httpserver/repositories/models"

	"github.com/jinzhu/gorm"
)

type TodoRepo struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) repositories.TodoRepo {
	return &TodoRepo{
		db: db,
	}
}

func (s *TodoRepo) GetTodos() (*[]models.Todo, error) {
	var todos []models.Todo
	err := s.db.Model(&models.Todo{}).Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return &todos, nil
}

func (s *TodoRepo) CreateTodo(todo *models.Todo) error {
	err := s.db.Create(&todo).Error
	return err
}

func (s *TodoRepo) GetTodoByID(id int) (*models.Todo, error) {
	var todo models.Todo
	err := s.db.Model(&models.Todo{}).Where("id = ?", id).Find(&todo).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (s *TodoRepo) UpdateTodo(id int, todo *models.Todo) error {
	err := s.db.Model(&models.Todo{}).Where("id = ?", id).Update(todo).Error
	return err
}

func (s *TodoRepo) DeleteTodo(id int) error {
	err := s.db.Model(&models.Todo{}).Where("id = ?", id).Delete(&models.Todo{}).Error
	return err
}
