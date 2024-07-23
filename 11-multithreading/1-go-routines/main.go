package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// Main é a thread 1 (auto-gerada)
// GC é a thread 2 (auto-gerada)
func main() {
	// Thread 3
	go task("A")
	// Thread 4
	go task("B")
	// Thread 5
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task anonymous is running\n", i)
			time.Sleep(1 * time.Second)
		}
	}()
	// Mais nada aqui.. sair
}

// Go Routines:
// Rodar simultaneamente várias tarefas.
