package main

import "fmt"

type Currency int

const (
	RMB Currency = iota
	USD
	EUR
)

func arrInit() {
	q := [4]int{1, 2, 3}

	for iq, vq := range q {
		fmt.Printf("%d %d\n", iq, vq)
	}

	q2 := [...]int{1, 2, 3}
	for iq, vq := range q2 {
		fmt.Printf("%d %d\n", iq, vq)
	}

	currs := [...]string{RMB: "￥", EUR: "$"}
	for is, vs := range currs {
		fmt.Printf("%d %s\n", is, vs)
	}
}

func arrComp() {
	a1 := [2]int{1, 3}
	a2 := [...]int{1, 3}
	a3 := [2]int{2, 3}

	fmt.Println(a1 == a2, a1 == a3)
}

func main() {
	arrInit()
	arrComp()
}
