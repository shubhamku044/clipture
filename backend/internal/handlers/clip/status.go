package clip

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shubhamku044/clipture/internal/handlers/common"
)

func (h *Handler) GetClipStatus(c *gin.Context) {
	clipID := c.Param("id")

	// TODO: Get status from database/cache
	// For now, return a mock response
	status := StatusResponse{
		ID:       clipID,
		Status:   "completed",
		Progress: 100,
	}

	common.SuccessResponse(c, status, "Status retrieved")
}

func (h *Handler) GetUserClips(c *gin.Context) {
	_, exists := c.Get("user_id")
	if !exists {
		common.ErrorResponse(c, http.StatusUnauthorized, "Authentication required")
		return
	}

	// TODO: Get user's clips from database
	clips := []ClipResponse{
		// Mock data
	}

	common.SuccessResponse(c, clips, "User clips retrieved")
}
