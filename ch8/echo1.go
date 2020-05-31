package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {

	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}

	err := c.Close()
	if err != nil {
		log.Printf("error closing connection:%v", err)
	} else {
		log.Println("connection closed")
	}
}

func echo(c net.Conn, text string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(text))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", text)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(text))
}

func echoErrorCheck(c net.Conn, text string, delay time.Duration) {
	_, err := fmt.Fprintln(c, "\t", strings.ToUpper(text))
	if err != nil {
		log.Printf("error echoing '%s':%v", text, err)
		return
	}
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", text)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(text))
}
