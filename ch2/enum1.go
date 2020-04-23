package main

import (
	"fmt"
)

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	ZiB
	YiB
)

func main() {
	t1 := ZiB / TiB
	fmt.Println(t1)
}