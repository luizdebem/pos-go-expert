package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("./arquivo.txt")

	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("Escrevendo aqui legal\n"))
	// tamanho, err := f.WriteString("Hello, World!\n")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso. Tamanho: %d\n", tamanho)

	f.Close()

	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	// leitura de pouco em pouco abrindo o arquivo
	// buffers
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
