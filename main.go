package main

import (
	"net/http"
	"todo-list/config"
	"todo-list/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.InitDatabase()

	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Todo List",
		})
	})

	todoRoutes := r.Group("/todos")
	{
		todoRoutes.GET("/", controllers.GetTodos)
		todoRoutes.POST("/create", controllers.CreateTodo)
		todoRoutes.GET("/:id", controllers.GetTodoById)
		todoRoutes.PUT("/:id", controllers.UpdateTodo)
		todoRoutes.DELETE("/:id", controllers.DeleteTodo)
	}

	r.Run()
}
