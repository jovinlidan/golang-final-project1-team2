package services

import (
	"database/sql"
	"golang-final-project1-team2/httpserver/controllers/params"
	"golang-final-project1-team2/httpserver/controllers/views"
	"golang-final-project1-team2/httpserver/repositories"
	"golang-final-project1-team2/httpserver/repositories/models"
	"strings"
)



type TodoSvc struct {
	repo repositories.TodoRepo
}

func NewTodoSvc(repo repositories.TodoRepo) *TodoSvc {
	return &TodoSvc{
		repo: repo,
	}
}


func (s *TodoSvc) GetTodos() *views.Response {
	todos, err := s.repo.GetTodos()
	if err != nil {
		if err == sql.ErrNoRows {
			return views.DataNotFound(err)
		}
		return views.InternalServerError(err)
	}

	return views.SuccessFindAllResponse(parseModelToTodos(todos), "GET_ALL_TODOS")
}

func (s *TodoSvc) CreateTodo(req *params.TodoCreateRequest) *views.Response {
	todo := parseCreateRequestToModel(req)
	
	err := s.repo.CreateTodo(todo)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return views.DataConflict(err)
		}
		return views.InternalServerError(err)
	}
	return views.SuccessCreateResponse(todo, "CREATE_TODO")
}

/// Create
func parseCreateRequestToModel(req *params.TodoCreateRequest) *models.Todo {
	return &models.Todo{
		Title: req.Title,
		Description: req.Description,
		CompletedAt: req.CompletedAt,
	}
}
///

func parseModelToTodos(mod *[]models.Todo) *[]views.Todo {
	var s []views.Todo
	for _, st := range *mod {
		s = append(s, views.Todo{
			ID: st.ID,
			Title: st.Title,
			Description: st.Description,
			CompletedAt: st.CompletedAt,
			
		})
	}
	return &s
}

