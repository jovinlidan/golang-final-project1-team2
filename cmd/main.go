package main

import (
	"golang-final-project1-team2/config"
	"golang-final-project1-team2/httpserver"
	"golang-final-project1-team2/httpserver/controllers"
	"golang-final-project1-team2/httpserver/repositories/gorm"
	"golang-final-project1-team2/httpserver/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectMySQLGORM()
	if err != nil {
		panic(err)
	}
	todoRepo := gorm.NewTodoRepo(db)
	todoSvc := services.NewTodoSvc(todoRepo)
	todoHandler := controllers.NewTodoController(todoSvc)

	router := gin.Default()

	app := httpserver.NewRouter(router, todoHandler)
	app.Start(":4000")

}
