package clip

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shubhamku044/clipture/internal/handlers/common"
)

func (h *Handler) DownloadClip(c *gin.Context) {
	filename := c.Param("filename")

	// Security validation
	if strings.Contains(filename, "..") || strings.Contains(filename, "/") || strings.Contains(filename, "\\") {
		common.ErrorResponse(c, http.StatusBadRequest, "Invalid filename")
		return
	}

	filePath := filepath.Join(h.config.OutputDir, filename)

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		common.ErrorResponse(c, http.StatusNotFound, "File not found")
		return
	}

	// Optional: Check user permissions here
	// if userID, exists := c.Get("user_id"); exists {
	//     if !h.hasPermissionToDownload(userID.(string), filename) {
	//         common.ErrorResponse(c, http.StatusForbidden, "Access denied")
	//         return
	//     }
	// }

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.File(filePath)
}
