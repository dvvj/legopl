package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	if cap(x) > len(x) {
		z = x[:len(x)+1]
	} else {
		zlen := len(x) + 1
		z = make([]int, zlen, 2*zlen)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func main() {
	//a := [...]int{0, 1, 2, 3, 4}
	var x, z []int
	for i := 0; i < 20; i++ {
		z = appendInt(x, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(z), z)
		x = z
	}
}
