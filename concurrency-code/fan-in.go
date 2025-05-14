package main

import (
	"sync"
)

func main(){
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func(){
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()

	go func(){
		for num := range ch {
			println("Channel:", num)
		}
		wg.Done()
	}()
	wg.Wait()
	println("All goroutines finished.")
}