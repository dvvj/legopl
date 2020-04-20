package main

import (
	"fmt"
	"os"
	"strings"
)

func last1(s string, b byte, front bool) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == b {
			if front {
				return s[:i]
			} else {
				return s[i+1:]
			}
		}
	}
	return s
}

func bn1(s string) string {
	s = last1(s, '/', false)
	return last1(s, '.', true)
}

func bn2(s string) string {
	slash := strings.LastIndex(s, "/")
	if slash >= 0 {
		s = s[slash+1:]
	}
	dot := strings.LastIndex(s, ".")
	if dot >= 0 {
		return s[:dot]
	}
	return s
}

func main() {
	for _, s := range os.Args {
		fmt.Println("bn1: " + bn1(s))
		fmt.Println("bn2: " + bn2(s))
	}
}
