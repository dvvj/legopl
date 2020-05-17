package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func simSlowOps(secs int) {
	defer trace("simSlowOps")()

	time.Sleep(time.Second * time.Duration(secs))
}

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("number expected but got '%s'", os.Args[1])
		os.Exit(1)
	}

	simSlowOps(n)
}
