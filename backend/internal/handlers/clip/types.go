package clip

type ClipRequest struct {
	VideoURL   string  `json:"video_url" binding:"required"`
	StartTime  float64 `json:"start_time" binding:"required"`
	EndTime    float64 `json:"end_time" binding:"required"`
	Platform   string  `json:"platform" binding:"required"`
	OutputName string  `json:"output_name,omitempty"`
	UserID     string  `json:"-"` // Set from auth middleware
}

type BatchClipRequest struct {
	Clips []ClipRequest `json:"clips" binding:"required"`
}

type ClipResponse struct {
	ID          string `json:"id"`
	Status      string `json:"status"`
	OutputPath  string `json:"output_path,omitempty"`
	DownloadURL string `json:"download_url,omitempty"`
	Error       string `json:"error,omitempty"`
	CreatedAt   int64  `json:"created_at,omitempty"`
	ProcessedAt int64  `json:"processed_at,omitempty"`
}

type BatchClipResponse struct {
	BatchID string         `json:"batch_id"`
	Clips   []ClipResponse `json:"clips"`
}

type StatusResponse struct {
	ID       string `json:"id"`
	Status   string `json:"status"`
	Progress int    `json:"progress,omitempty"`
	Error    string `json:"error,omitempty"`
}
