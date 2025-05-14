package main

import (
	"sync"
	"time"
)

func main() {
	ch := make(chan string, 1)
	writerWg := sync.WaitGroup{}
	readerWg := sync.WaitGroup{}
	readerWg.Add(2)
	writerWg.Add(2)
	go func() {
		for i := 0; i < 2; i++ {
			// In the second iteration, it will block because the channel is full
			ch <- "hello"
			println("Added value to channel")
			writerWg.Done()
		}
	}()
	go func() {
		time.Sleep(5 * time.Second)
		for i := range ch {
			println("Channel:", i)
			readerWg.Done()
		}
	}()
	writerWg.Wait()
	readerWg.Wait()
}
