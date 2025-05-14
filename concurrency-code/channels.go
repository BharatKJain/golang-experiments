package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			ch1 <- i
			time.Sleep(1 * time.Second) // Simulate some work
		}
		close(ch1)
	}()

	go func() {
		for i := 6; i <= 10; i++ {
			ch2 <- i
			time.Sleep(2 * time.Second) // Simulate some work
		}
		close(ch2)
	}()
	wg := &sync.WaitGroup{}


	wg.Add(2)
	go func() {
		for num := range ch1 {
			fmt.Println("Channel 1:", num)
		}
		wg.Done()
	}()
	go func() {
		for num := range ch2 {
			fmt.Println("Channel 2:", num)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("All goroutines finished.")

}