package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	hostPort := os.Args[1]
	conn, err := net.Dial("tcp", hostPort)

	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} // signal end
	}()

	copyFrom(conn, os.Stdin)
	conn.Close()
	<-done
}

func copyFrom(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
