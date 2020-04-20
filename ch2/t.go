package main

import (
	"fmt"
	"legopl/ch2/initfunc"
	"legopl/ch2/tempconv"
	"unicode/utf8"
)

func main() {
	c := tempconv.Celsius(10)

	fmt.Println(c)

	t1 := initfunc.GetShortName("PRC1")
	fmt.Println(t1)

	fmt.Println("=========== String tests")
	s := "hello, 世界"
	fmt.Println(len(s))
	fmt.Println(s[0:9])
	fmt.Println(s[0:10])
	fmt.Println(utf8.RuneCountInString(s))
	for i, r := range s {
		fmt.Printf("%d\t%q\t%#[2]x\n", i, r)
	}

	r := []rune(s)
	fmt.Println(len(r))
	fmt.Printf("r[7]: %s\n", string(r[7]))
}
