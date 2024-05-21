package main

import "fmt"

// STRUCTS

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar() // QUALQUER STRUCT QUE TIVER O MÉTODO DESATIVAR ESTÁ IMPLEMENTANDO A INTERFACE PESSOA. (implements Pessoa)
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool

	// Endereco // Composição
	Address Endereco // propriedade do tipo endereço.
}

type Empresa struct {
	Nome string
}

func (c Cliente) Desativar() {
	c.Ativo = false
}

func (e Empresa) Desativar() {

}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 20,
		Ativo: true,
	}

	minhaEmpresa := Empresa{
		Nome: "Josias Ltda",
	}

	Desativacao(minhaEmpresa)
	Desativacao(wesley)

	fmt.Printf("Nome: %s\nIdade: %d\nAtivo: %t\n", wesley.Nome, wesley.Idade, wesley.Ativo)
}

// go way
