package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

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

var imagefiles = []string{
	"a.jpg",
	"d.jpg",
	"b.png",
}

func testImage1() {
	Images1(imagefiles)

	for {
		time.Sleep(1000)
	}
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

var imagefiles2 = []string{
	"b.png",
	"a.jpg",
	"c.png",
	"d.jpg",
}

func Images3(filenames []string) error {
	errors := make(chan error)

	for _, file := range filenames {
		go func(f string) {
			th, err := Image(f)
			fmt.Printf("thumnail for %s: %s\n", f, th)
			if err == nil {
				fmt.Printf("sleeping for %s: %s\n", f, th)
				time.Sleep(1 * time.Second)
				fmt.Printf("done sleeping for %s: %s\n", f, th)
			}
			errors <- err
		}(file)
	}

	for range filenames {
		if err := <-errors; err != nil {
			fmt.Printf("len(errors): %d\n", len(errors))
			close(errors)
			//return err
		}
	}

	return nil
}

func Images_WaitGroup(filenames []string) {
	thumbnames := make(chan string)

	var wg sync.WaitGroup

	filechan := make(chan string)

	go func() {
		for _, f := range filenames {
			filechan <- f
		}
		close(filechan)
	}()

	for file := range filechan {
		wg.Add(1)
		log.Printf("processing %s...\n", file)
		go func(f string) {
			defer wg.Done()

			th, err := Image(f)
			if err != nil {
				log.Println(err)
				return
			}
			thumbnames <- th
		}(file)
	}

	go func() {
		wg.Wait()
		close(thumbnames)
	}()

	for thnames := range thumbnames {
		fmt.Println("thumb: " + thnames)
	}
}

func main() {
	//Images2(imagefiles)
	//Images3(imagefiles2)
	Images_WaitGroup(imagefiles2)
}
