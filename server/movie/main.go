package main

import (
	controller "github.com/RNekoCloud/gin-dockerized/controller/movie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Route GIN
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to my app!",
		})
	})

	r.POST("/movie", controller.AddMovie)

	// Initialize Gin App
	r.Run(":3003")
}
