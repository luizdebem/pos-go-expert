package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

// Main é a thread 1 (auto-gerada)
// GC é a thread 2 (auto-gerada)
func main() {

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)

	// Thread 3
	go task("A", &waitGroup)
	// Thread 4
	go task("B", &waitGroup)
	// Thread 5
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task anonymous is running\n", i)
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()

	waitGroup.Wait()

}

// Go Routines:
// Rodar simultaneamente várias tarefas.
