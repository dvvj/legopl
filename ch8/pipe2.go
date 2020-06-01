package main

import "fmt"

func main() {
	naturalsChan := make(chan int)
	squarerChan := make(chan int)

	go counter(naturalsChan)
	go squarer(squarerChan, naturalsChan)

	printer(squarerChan)
}

func counter(out chan<- int) {
	for i := 1; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}
