package main

import "fmt"

type Cliente struct {
	Nome string
}

func (c Cliente) andou() {
	c.Nome = "Wesley Willians"
	fmt.Printf("O cliente %s andou\n", c.Nome)
}

type Conta struct {
	saldo float64
}

func NewConta() *Conta {
	return &Conta{saldo: 0.0}
}

func (c *Conta) simular(valor float64) float64 {
	c.saldo += valor
	fmt.Printf("Saldo atual simulando: %.2f\n", c.saldo)
	return c.saldo
}

func main() {
	wesley := Cliente{
		Nome: "Wesley",
	}
	wesley.andou()
	fmt.Printf("O valor da struct com nome %s\n", wesley.Nome)

	conta := Conta{saldo: 100.0}
	conta.simular(200.0)
	fmt.Printf("Saldo: %.2f\n", conta.saldo)
}
