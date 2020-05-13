package main

import (
	"fmt"
	"strings"
)

func sqrt(n int) int { return n * n }
func neg(n int) int  { return -n }

func test1() {
	var f func(int) int

	//fmt.Println(f(20)) // panic

	f = sqrt
	fmt.Println(f(20))
	f = neg
	fmt.Println(f(20))
}

func add1(r rune) rune { return r + 1 }

func testStringsMap() {
	fmt.Println(strings.Map(add1, "abc"))
	fmt.Println(strings.Map(add1, "123"))
	fmt.Println(strings.Map(add1, "你好"))
}

func main() {
	test1()
	testStringsMap()
}
