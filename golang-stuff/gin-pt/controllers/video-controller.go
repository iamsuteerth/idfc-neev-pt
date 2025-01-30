package controller

import (
	"iamsuteerth/golang-stuff/gin-pt/entity"
	"iamsuteerth/golang-stuff/gin-pt/service"
	"iamsuteerth/golang-stuff/gin-pt/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type VideoController interface {
	FindAll() []entity.Video
	Save(context *gin.Context) error
	ShowAll(context *gin.Context)
}

type videoController struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &videoController{
		service: service,
	}
}

func (vc *videoController) FindAll() []entity.Video {
	return vc.service.FindAll()
}
func (vc *videoController) Save(context *gin.Context) error {
	var video entity.Video
	// Extract payload from context which is JSON
	// Unmarshall the video
	err := context.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	vc.service.Save(video)
	return nil
}
func (vc *videoController) ShowAll(context *gin.Context) {
	videos := vc.service.FindAll()
	// Store all the variables to be used within the templates
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	// Pass this to template
	context.HTML(http.StatusOK, "index.html", data)
}
