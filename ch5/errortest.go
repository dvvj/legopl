package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readTillEOF(f *os.File) (int, error) {
	in := bufio.NewReader(f)
	var c int
	for {
		_, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
		c++
	}
	return c, nil
}

func main() {
	file2Open := os.Args[1]

	f, err := os.Open(file2Open)

	if err != nil {
		fmt.Printf("main failed: %v\n", err)
	} else {
		fmt.Printf("main successfully opened file %s\n", file2Open)
	}

	c, err := readTillEOF(f)
	if err != nil {
		fmt.Printf("main failed: %v\n", err)
	} else {
		fmt.Printf("main rune count: %d\n", c)
	}

	f.Close()

}
