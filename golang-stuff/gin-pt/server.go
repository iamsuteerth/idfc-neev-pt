package main

import (
	controller "iamsuteerth/golang-stuff/gin-pt/controllers"
	"iamsuteerth/golang-stuff/gin-pt/middlewares"
	"iamsuteerth/golang-stuff/gin-pt/service"
	"io"
	"net/http"
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

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("./templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/test", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "Hello World!",
			})
		})
		apiRoutes.GET("/videos", func(context *gin.Context) {
			context.JSON(200, videController.FindAll())
		})
		apiRoutes.POST("/videos", func(context *gin.Context) {
			err := videController.Save(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"message": "Video input is valid!",
				})
			}
		})
	}
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videController.ShowAll)
	}
	server.Run(":8080")
}
