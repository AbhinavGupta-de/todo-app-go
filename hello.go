package main

import (
	"gupta/first/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	var g *gin.Engine
	g = gin.Default()
	
	g.GET("/", controllers.Index)

	g.GET("/tasks", controllers.Tasks)
	g.POST("/task", controllers.PostTask)
	g.PUT("/task/:id", controllers.PutTask)

	g.Run(":3000")
}