package common

import (
	"fmt"
	"net/url"
	"strings"
)

func ValidateYouTubeURL(videoURL string) error {
	if videoURL == "" {
		return fmt.Errorf("video URL is required")
	}

	parsedURL, err := url.Parse(videoURL)
	if err != nil {
		return fmt.Errorf("invalid URL format")
	}

	allowedHosts := []string{"youtube.com", "www.youtube.com", "youtu.be", "m.youtube.com"}
	isValid := false
	for _, host := range allowedHosts {
		if strings.Contains(parsedURL.Host, host) {
			isValid = true
			break
		}
	}

	if !isValid {
		return fmt.Errorf("only YouTube URLs are supported")
	}

	return nil
}

func ValidatePlatform(platform string) error {
	validPlatforms := []string{"instagram", "tiktok", "youtube"}
	for _, p := range validPlatforms {
		if p == platform {
			return nil
		}
	}
	return fmt.Errorf("invalid platform. Supported: %s", strings.Join(validPlatforms, ", "))
}

func ValidateTimeRange(start, end float64) error {
	if start < 0 {
		return fmt.Errorf("start time cannot be negative")
	}
	if end <= start {
		return fmt.Errorf("end time must be greater than start time")
	}
	if end-start > 300 { // 5 minutes max
		return fmt.Errorf("clip duration cannot exceed 5 minutes")
	}
	return nil
}
