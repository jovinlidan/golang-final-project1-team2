package controllers

import (
	"golang-final-project1-team2/httpserver/controllers/params"
	"golang-final-project1-team2/httpserver/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	_ "golang-final-project1-team2/httpserver/controllers/views"
	_ "golang-final-project1-team2/httpserver/repositories/models"
)

type TodoController struct {
	svc *services.TodoSvc
}

func NewTodoController(svc *services.TodoSvc) *TodoController {
	return &TodoController{
		svc: svc,
	}
}

// GetTodos godoc
// @Summary      List all todos
// @Description  get all todos
// @Tags         todos
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Todo
// @Failure      400  {object}  views.Response
// @Failure      404  {object}  views.Response
// @Failure      500  {object}  views.Response
// @Router       /todos [get]
func (s *TodoController) GetTodos(ctx *gin.Context) {
	response := s.svc.GetTodos()
	WriteJsonRespnse(ctx, response)
}

// CreateTodo godoc
// @Summary      Create new todo
// @Description  create a new single todo
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param request body params.TodoCreateRequest true "query params"
// @Success      200  {object}  models.Todo
// @Failure      400  {object}  views.Response
// @Failure      404  {object}  views.Response
// @Failure      500  {object}  views.Response
// @Router       /todos [post]
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
