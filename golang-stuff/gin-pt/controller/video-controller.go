package controller

import (
	"iamsuteerth/golang-stuff/gin-pt/entity"
	"iamsuteerth/golang-stuff/gin-pt/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(context *gin.Context) entity.Video
}

type videoController struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &videoController{
		service: service,
	}
}

func (vc *videoController) FindAll() []entity.Video {
	return vc.service.FindAll()
}
func (vc *videoController) Save(context *gin.Context) entity.Video {
	var video entity.Video
	// Extract payload from context which is JSON
	// Unmarshall the video
	context.BindJSON(&video)
	vc.service.Save(video)
	return video
}
