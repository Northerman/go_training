package main

import (
	"github.com/gin-gonic/gin"
	"github.com/northerman/myapi/middleware"
	"github.com/northerman/myapi/task"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	apiV1 := r.Group("/api/v1")

	apiV1.Use(middleware.Auth)

	apiV1.GET("/todos", task.GetTodosHandler)
	apiV1.GET("/todos/:id", task.GetTodoByIdHandler)
	apiV1.POST("/todos", task.CreateTodosHandler)
	apiV1.PUT("/todos/:id", task.UpdateTodosHandler)
	apiV1.DELETE("/todos/:id", task.DeleteTodosHandler)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":1234")
}
