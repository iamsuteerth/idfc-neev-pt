package service

import "iamsuteerth/golang-stuff/gin-pt/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}
type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}
func (vs *videoService) Save(video entity.Video) entity.Video {
	vs.videos = append(vs.videos, video)
	return video
}
func (vs *videoService) FindAll() []entity.Video {
	return vs.videos
}
