package downloader

import (
	"testing"
)

func TestDownload(t *testing.T) {
	url := "https://youtube.com"

	if got := Download(url); got != nil {
		t.Errorf("Want %s but got %s", url, got)
	}
}