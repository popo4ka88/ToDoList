package server

import (
	"ToDoList/internal/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func homeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tpl", gin.H{})
}

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("static/home.tpl")

	router.GET("/", homeHandler)
	router.GET("/todo", handlers.FetchTodos)
	router.POST("/todo", handlers.CreateTodo)
	router.PUT("/todo/:id", handlers.UpdateTodo)
	router.DELETE("/todo/:id", handlers.DeleteTodo)
	return router
}
