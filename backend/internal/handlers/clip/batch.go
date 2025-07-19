package clip

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shubhamku044/clipture/internal/handlers/common"
)

func (h *Handler) CreateBatchClip(c *gin.Context) {
	var req BatchClipRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.ValidationErrorResponse(c, err)
		return
	}

	// Validate batch size
	if len(req.Clips) > 10 {
		common.ErrorResponse(c, http.StatusBadRequest, "Maximum 10 clips per batch")
		return
	}

	batchID := uuid.New().String()

	responses := make([]ClipResponse, len(req.Clips))
	for i, clip := range req.Clips {
		if err := h.validateClipRequest(clip); err != nil {
			responses[i] = ClipResponse{
				ID:     batchID + "-" + string(rune(i)),
				Status: "error",
				Error:  err.Error(),
			}
			continue
		}

		clipID := batchID + "-" + string(rune(i))
		responses[i] = ClipResponse{
			ID:     clipID,
			Status: "processing",
		}

		// Process in background
		go h.processClipAsync(clip, clipID)
	}

	response := BatchClipResponse{
		BatchID: batchID,
		Clips:   responses,
	}

	common.SuccessResponse(c, response, "Batch processing started")
}
