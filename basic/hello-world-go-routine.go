package main

import (
	"io"
	"os"
	"time"
	"fmt"
)

func main() {
	go echo(os.Stdin, os.Stdout)
	go printDelayedHello()
	// Blocking code. Makes main routine to wait for 5 seconds
	time.Sleep(5 * time.Second)
}

func echo(out io.Writer, in io.Reader) {
	io.Copy(out, in)
}

func printDelayedHello() {

	time.Sleep(2 * time.Second)
	fmt.Println("Hello")
}
