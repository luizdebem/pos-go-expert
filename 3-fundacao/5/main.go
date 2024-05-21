package main

import "fmt"

var ()

func main() {
	var meuArray [3]int // criando um array chamado meuArray de 3 posições fixas.
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30
	// meuArray[3] = 40 // Out of bounds

	fmt.Println(meuArray[len(meuArray)-1])

	// percorrer
	for i, v := range meuArray {
		fmt.Printf("O valor do index é %d e o valor é %d\n", i, v)
	}

}
