package controllers

import (
	"golang-final-project1-team2/httpserver/controllers/params"
	"golang-final-project1-team2/httpserver/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TodoController struct {
	svc *services.TodoSvc
}

func NewTodoController(svc *services.TodoSvc) *TodoController {
	return &TodoController{
		svc: svc,
	}
}

func (s *TodoController) GetTodos(ctx *gin.Context) {
	response := s.svc.GetTodos()
	WriteJsonRespnse(ctx, response)
}


func (s *TodoController) CreateTodo(ctx *gin.Context) {
	var req params.TodoCreateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := s.svc.CreateTodo(&req)
	WriteJsonRespnse(ctx, response)

}
