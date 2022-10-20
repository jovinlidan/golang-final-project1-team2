package httpserver

import (
	"golang-final-project1-team2/httpserver/controllers"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)
type Router struct {
	router  *gin.Engine
	todo *controllers.TodoController
}

func NewRouter(router *gin.Engine, todo *controllers.TodoController) *Router {
	return &Router{
		router:  router,
		todo: todo,
	}
}

func (r *Router) Start(port string) {
	r.router.POST("/todos", r.todo.CreateTodo)
	r.router.GET("/todos", r.todo.GetTodos)
	r.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.router.Run(port)
}