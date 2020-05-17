package main

import (
	"fmt"
	"os"
	"strings"
)

func defer1() {
	fmt.Println("in defer1")
}

func defer2() {
	fmt.Println("in defer2")
}

func testDefer(file string) (string, error) {
	f, err := os.Open(file)
	defer defer1()
	defer defer2()

	if err != nil {
		return "", err
	}
	defer f.Close()

	if strings.HasSuffix(f.Name(), ".txt") {
		return "ok: is text file", nil
	} else if strings.HasSuffix(f.Name(), ".out") {
		panic(fmt.Errorf("should not be an .out file"))
	} else {
		return "", fmt.Errorf("failed: not a txt file")
	}
}

func main() {
	s, err := testDefer(os.Args[1])
	fmt.Printf("%s/%v\n", s, err)
}
