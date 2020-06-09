package main

import (
	"fmt"
	"sync"
	"time"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				fmt.Println("[sq] return ...")
				return
			}
		}
	}()

	return out
}

// fan-in
func merge(done <-chan struct{}, in ...<-chan int) chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				fmt.Println("[merge] return ...")
				return
			}
		}

	}
	wg.Add(len(in))

	for _, c := range in {
		go output(c)
	}

	go func() {
		fmt.Println("waiting to close out...")
		wg.Wait()
		fmt.Println("closing out...")
		close(out)
	}()

	return out
}

func main() {
	c := gen(2, 3)
	done := make(chan struct{})
	c1 := sq(done, c)
	c2 := sq(done, c)

	out := merge(done, c1, c2)
	fmt.Println(<-out)

	close(done)

	time.Sleep(100)

	x, ok := <-out
	if ok {
		fmt.Printf("ok: x = %d\n", x)
	} else {
		fmt.Printf("!ok: x = %d\n", x)
	}

	// for {
	// 	if len(out) > 0 {
	// 		fmt.Println("not done ...")
	// 		time.Sleep(1000)
	// 	} else {
	// 		break
	// 	}
	// }
}
