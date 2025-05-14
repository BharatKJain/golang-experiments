package main

import (
    "fmt"
    "time"
)

func worker(done chan bool, sleep int) {
    fmt.Println("working...")
    time.Sleep(time.Duration(sleep) * time.Second)
    fmt.Println("done")

    done <- true
}

func main() {

    done := make(chan bool, 0)
	task2 := make(chan bool, 0)
	async_task_3_after_task_1 := make(chan bool, 0)
    go worker(done,1)
	go worker(task2,5)

    response1:= <-done
	if response1 {
		go worker(async_task_3_after_task_1,2)
	}
	response2:= <-task2
	response3 := <-async_task_3_after_task_1

	fmt.Println("task1:", response1)
	fmt.Println("task2:", response2)
	fmt.Println("task3:", response3)
}