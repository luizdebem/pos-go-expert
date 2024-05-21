package main

import "fmt"

func main() {
	salarios := map[string]int{"Josias": 1000, "Joaquim": 2000, "Jorge": 3000}
	fmt.Println(salarios["Josias"])
	delete(salarios, "Josias")
	fmt.Println(salarios["Josias"])
	salarios["Jos"] = 1250
	fmt.Println(salarios["Jos"])

	// sal := make(map[string]int)
	// sal1 := map[string]int{}

	for nome, salario := range salarios {
		fmt.Printf("Nome: %s - Salario: %d\n", nome, salario)
	}

	// blank identifier
	for _, salario := range salarios {
		fmt.Printf("Salario: %d\n", salario)
	}
}
