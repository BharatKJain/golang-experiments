package main

import "fmt"
import "time"

func printHello(counter int){
	fmt.Println("Hello, World! with counter: ", counter)
}

func main(){
	// Start a goroutine to run the printHello function
	go printHello(0)
	time.Sleep(1 * time.Microsecond)
	go printHello(1)
	time.Sleep(1 * time.Microsecond)
	go printHello(2)

	// Wait for user input to prevent the program from exiting immediately
	var input string
	fmt.Scanln(&input)
	fmt.Println("Program finished.")
	// The goroutine will run in the background and print "Hello, World!" before the program exits
}