package memo

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func incomingUrls() []string {
	return []string{
		"https://google.com",
		"https://golang.org",
		"https://play.golang.org",
		"https://golang.org",
		"https://play.golang.org",
		"https://godoc.org",
		"https://play.golang.org",
	}
}
func TestMemo1(t *testing.T) {
	m := New(httpGetBody)

	for _, url := range incomingUrls() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}

func TestMemo2(t *testing.T) {
	m := New(httpGetBody)

	var n sync.WaitGroup

	for _, url1 := range incomingUrls() {
		n.Add(1)

		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url1)
	}

	n.Wait()
}
