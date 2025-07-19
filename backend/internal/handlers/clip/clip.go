package clip

import (
	"github.com/shubhamku044/clipture/internal/config"
	"github.com/shubhamku044/clipture/internal/services/video"
)

type Handler struct {
	config       *config.Config
	videoService *video.Service
}

func NewHandler(cfg *config.Config, videoSvc *video.Service) *Handler {
	return &Handler{
		config:       cfg,
		videoService: videoSvc,
	}
}
