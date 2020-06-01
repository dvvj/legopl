package thumbnail

import (
	"fmt"
	"strings"
	"time"
)

// simulate thumbnail functionality
func Image(filename string) (string, error) {
	if strings.HasSuffix(filename, ".jpg") {
		return "th_" + filename, nil
	} else {
		return "", fmt.Errorf("suffix not supported")
	}
}

func Images1(filenames []string) {
	for _, f := range filenames {
		go func(f string) {
			th, _ := Image(f)
			fmt.Printf("thumnail for %s: %s\n", f, th)
		}(f)
	}
	time.Sleep(100)
}

func Images2(filenames []string) {
	ch := make(chan struct{})

	for _, f := range filenames {
		go func(f string) {
			th, _ := Image(f)
			fmt.Printf("thumnail for %s: %s\n", f, th)
			ch <- struct{}{}
		}(f)
	}

	for range filenames {
		<-ch
	}
}
