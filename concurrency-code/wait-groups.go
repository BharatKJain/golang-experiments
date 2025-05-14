package main

import (
	"time"
	"sync"
)

func main(){
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(){
		time.Sleep(2 * time.Second)
		println("Completed 1st thread!")
		defer wg.Done()
	}()
	go func(){
		time.Sleep(1 * time.Second)
		println("Completed 2nd thread!")
		defer wg.Done()
	}()
	go func(){
		time.Sleep(1 * time.Second / 2)
		println("Completed 3nd thread!")
		defer wg.Done()
	}()
	println("Program almost reached end without waiting for goroutine.")
	println("Waiting for wait-group to finish...")
	wg.Wait()
	println("Program end.")
}