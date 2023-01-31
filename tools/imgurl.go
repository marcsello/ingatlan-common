package tools

import (
	"io"
	"net/http"
	"strings"
)

// IsValidImageUrl does some http requests to check if the url is available, and it resolves to a valid image by checking the mime type
func IsValidImageUrl(url string) bool {

	if url == "" {
		return false
	}

	resp, err := http.Get(url)
	if err != nil {
		return false
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	mimeType := http.DetectContentType(body)
	mimeLeft := strings.Split(mimeType, "/")[0] // 0th element is always valid, it's the string itself
	return mimeLeft == "image"
}
