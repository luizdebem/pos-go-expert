package main

import "fmt"

func main() {
	// Memória -> Endereço -> Valor.
	// variável -> ponteiro que tem um endereço na memória -> valor
	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a
	*b = 30
	fmt.Println(*b) // * = qual valor está guardado na memória?
	fmt.Println(a)  // 30
}
