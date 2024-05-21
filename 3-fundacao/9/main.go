package main

import "fmt"

// Funções variádicas

func main() {
	fmt.Println(sum(1, 2, 3, 5, 5, 3, 2, 1, 2, 3, 4))
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
