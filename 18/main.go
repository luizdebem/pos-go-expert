package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/luizdebem/curso-go/matematica"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Volkswagen"}
	fmt.Printf("Resultado: %d\n", s)
	fmt.Printf("A: %d\n", matematica.A)
	fmt.Printf("Carro: %s\n", carro.Marca)
	fmt.Println(uuid.New().String())
}
