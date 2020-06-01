package main

import "fmt"

func main() {
	naturals := make(chan int)
	squarer := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			naturals <- i
		}
		close(naturals)
	}()

	go func() {
		for i := range naturals {
			squarer <- i * i
		}
		close(squarer)
	}()

	for r := range squarer {
		fmt.Println(r)
	}
}
