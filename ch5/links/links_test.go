package links

import (
	"testing"
)

func TestExtractFile(t *testing.T) {
	file := "/home/devvj/golang/src/legopl/ch5/index.html"
	links, err := ExtractFile(file)

	if len(links) <= 0 || err != nil {
		t.Errorf("failed to extract link from file: %s", file)
	}
}

func TestExtractUrl(t *testing.T) {
	url := "https://google.com"

	links, err := ExtractUrl(url)
	if len(links) <= 0 || err != nil {
		t.Errorf("failed to extract link from url: %s", url)
	}
}
