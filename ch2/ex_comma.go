package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func addComma(s string) string {
	slices := len(s) / 3
	var buf bytes.Buffer

	startIndex := len(s) - slices*3

	if startIndex > 0 {
		buf.WriteString(s[:startIndex])
	}

	for i := 0; i < slices; i++ {
		if startIndex > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(s[startIndex : startIndex+3])
		startIndex += 3
	}

	return buf.String()
}

func main() {
	for _, s := range os.Args[1:] {
		v, _ := strconv.Atoi(s)
		fmt.Printf("arg: %s, v: %d, %s\n", s, v, addComma(s))
	}
}
