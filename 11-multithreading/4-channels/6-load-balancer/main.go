package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)

	WorkerAmount := 10_000
	DataAmount := 100_000

	for i := 0; i < WorkerAmount; i++ {
		go worker(i, data)
	}

	for i := 0; i < DataAmount; i++ {
		data <- i
	}
}
