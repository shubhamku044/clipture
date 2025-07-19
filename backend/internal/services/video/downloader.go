package video

import (
	"fmt"
	"os/exec"
)

type Service struct {
	// Add dependencies like logger, config, etc.
}

type ClipPlatform string

const (
	PlatformInstagram ClipPlatform = "instagram"
	PlatformTikTok    ClipPlatform = "tiktok"
	PlatformYouTube   ClipPlatform = "youtube"
)

func (s *Service) ProcessClip(videoURL string, startTime, endTime int, platform ClipPlatform, clipID string) error {
	// Step 1: Download video from YouTube
	downloadPath, err := downloadYouTubeVideo(videoURL, clipID)
	if err != nil {
		return fmt.Errorf("failed to download video: %w", err)
	}

	// Step 2: Clip the video using ffmpeg
	outputPath := fmt.Sprintf("clips/%s.mp4", clipID)
	err = clipVideo(downloadPath, outputPath, startTime, endTime)
	if err != nil {
		return fmt.Errorf("failed to clip video: %w", err)
	}

	// Step 3: (Optional) Transcode/rescale according to platform
	// TODO: Platform-specific processing

	return nil
}

func downloadYouTubeVideo(url string, clipID string) (string, error) {
	output := fmt.Sprintf("videos/%s.mp4", clipID)

	cmd := exec.Command("yt-dlp", "-f", "mp4", "-o", output, url)
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return output, nil
}

func clipVideo(inputPath, outputPath string, start, end int) error {
	cmd := exec.Command("ffmpeg", "-ss", fmt.Sprintf("%d", start), "-to", fmt.Sprintf("%d", end),
		"-i", inputPath, "-c", "copy", outputPath)
	return cmd.Run()
}
