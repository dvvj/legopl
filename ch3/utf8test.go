package main

import (
	"fmt"
	"unicode/utf8"
)

func onebitTest() {
	var b [1]byte
	var i byte
	for i = 0; i <= 254; i++ {
		b[0] = i

		t, _ := utf8.DecodeRune(b[:])
		fmt.Printf("%02x: %v\n", i, t)
	}
}

func twobitsTest() {
	var b [2]byte
	b[0] = 0
	var i byte
	for i = 0; i <= 254; i++ {
		b[1] = i

		t, _ := utf8.DecodeRune(b[:])
		fmt.Printf("00%02x: %v\n", i, t)
	}
}

func main() {
	//onebitTest()
	twobitsTest()
}
