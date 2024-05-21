package main

import "fmt"

func main() {

	// Closures, Funções Anônimas...
	total := func() int {
		return sum(5, 5, 5, 5, 5, 5) * 2
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
