// Exercício de Ponteiros em Go

// Defina uma função chamada modifyValue que recebe um ponteiro para um inteiro como argumento. Esta função deve dobrar o valor do inteiro apontado pelo ponteiro.

// No main(), declare uma variável inteira e atribua a ela um valor inicial.

// Declare um ponteiro para a variável inteira e atribua o endereço da variável.

// Chame a função modifyValue, passando o ponteiro como argumento.

// Imprima o novo valor da variável.

package main

import "fmt"

func modifyValue(value *int) {
	*value = *value * 2
}

func main() {
	value := 10
	pointerToValue := &value // aqui to apontando para o endereço de value.
	modifyValue(pointerToValue)
	fmt.Println(value)
}
