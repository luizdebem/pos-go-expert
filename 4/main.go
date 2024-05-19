package main

import "fmt"

type ID int

func main() {
	var a float64 = 1.337
	var b ID = 5
	fmt.Println("Olá mundo.")
	fmt.Printf("O tipo de a é %T\n", a)
	fmt.Printf("O tipo de b é %T\n", b)
}
