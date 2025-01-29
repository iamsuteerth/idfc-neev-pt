package main

import (
	"iamsuteerth/golang-stuff/gin-pt/controller"
	"iamsuteerth/golang-stuff/gin-pt/middlewares"
	"iamsuteerth/golang-stuff/gin-pt/service"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService   service.VideoService       = service.New()
	videController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	file, _ := os.Create("./logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}

func main() {
	// server := gin.Default()
	setupLogOutput() // Logging to both stdout and the file in the path specified
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())
	server.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	server.GET("/videos", func(context *gin.Context) {
		context.JSON(200, videController.FindAll())
	})
	server.POST("/videos", func(context *gin.Context) {
		context.JSON(201, videController.Save(context))
	})
	server.Run(":8080")
}
