package thumbnail

import (
	"testing"
)

var imagefiles = []string{
	"a.jpg",
	"b.jpg",
	"c.exe",
}

func TestImages1(t *testing.T) {
	Images1(imagefiles)
}

func TestImages2(t *testing.T) {
	Images2(imagefiles)
}
