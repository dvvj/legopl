package main

import (
	"fmt"
	"legopl/ch2/initfunc"
	"legopl/ch2/tempconv"
)

func main() {
	c := tempconv.Celsius(10)

	fmt.Println(c)

	t1 := initfunc.GetShortName("PRC1")
	fmt.Println(t1)
}
