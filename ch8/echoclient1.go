package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	go copyFrom(os.Stdout, conn)
	copyFrom(conn, os.Stdin)
}

func copyFrom(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
