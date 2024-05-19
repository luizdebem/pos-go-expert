package main

import "fmt"

// Se eu quiser tornar algum valor MUTÁVEL, a melhor forma é com ponteiros.
// Aqui numa função de soma simples não precisaria, por exemplo..

func soma(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	soma := soma(&minhaVar1, &minhaVar2)
	fmt.Printf("a soma é: %d\n", soma)
	fmt.Printf("minhaVar1 é: %d\n", minhaVar1)
	fmt.Printf("minhaVar2 é: %d\n", minhaVar2)
}
