package clip

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shubhamku044/clipture/internal/handlers/common"
	"github.com/shubhamku044/clipture/internal/services/video"
)

func (h *Handler) CreateClip(c *gin.Context) {
	var req ClipRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.ValidationErrorResponse(c, err)
		return
	}

	// Validate request
	if err := h.validateClipRequest(req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	clipID := uuid.New().String()

	// Set user ID from context (if authenticated)
	if userID, exists := c.Get("user_id"); exists {
		req.UserID = userID.(string)
	}

	// Process clip asynchronously
	response := ClipResponse{
		ID:        clipID,
		Status:    "processing",
		CreatedAt: time.Now().Unix(),
	}

	// Start processing in background
	go h.processClipAsync(req, clipID)

	common.SuccessResponse(c, response, "Clip processing started")
}

func (h *Handler) validateClipRequest(req ClipRequest) error {
	if err := common.ValidateYouTubeURL(req.VideoURL); err != nil {
		return err
	}
	if err := common.ValidatePlatform(req.Platform); err != nil {
		return err
	}
	if err := common.ValidateTimeRange(req.StartTime, req.EndTime); err != nil {
		return err
	}
	return nil
}

func (h *Handler) processClipAsync(req ClipRequest, clipID string) {
	// This would typically update a database or cache with status
	// For now, we'll just process the video
	err := h.videoService.ProcessClip(req.VideoURL, int(req.StartTime), int(req.EndTime), video.ClipPlatform(req.Platform), clipID)
	if err != nil {
		// Update status in database/cache
	} else {
		// Update status in database/cache
	}
}
