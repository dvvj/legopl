package main

import "fmt"

func main() {
	q := [4]int{1, 2, 3}

	for iq, vq := range q {
		fmt.Printf("%d %d\n", iq, vq)
	}

	q2 := [...]int{1, 2, 3}
	for iq, vq := range q2 {
		fmt.Printf("%d %d\n", iq, vq)
	}
}