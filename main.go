package main

import (
	"brainstorm-press-api/controllers"
	"brainstorm-press-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	db := models.SetupModels()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/posts", controllers.FindPosts)
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts/:id", controllers.FindPost)
	r.PATCH("/posts/:id", controllers.UpdatePost)
	r.DELETE("posts/:id", controllers.DeletePost)

	r.Run()
}
