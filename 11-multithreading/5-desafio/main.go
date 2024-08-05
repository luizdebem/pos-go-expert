/*
Neste desafio você terá que usar o que aprendemos com Multithreading e APIs
para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

https://brasilapi.com.br/api/cep/v1/01153000 + cep

http://viacep.com.br/ws/" + cep + "/json/

Os requisitos para este desafio são:

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line com os dados do endereço,
bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.
*/

package main

import (
	"io"
	"net/http"
	"time"
)

var cep = "22431050"

func BrasilAPIRequest() string {
	println("Iniciando request para BrasilAPI....")
	req, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	return "Resposta de BrasilAPI: " + string(res)
}

func ViaCEPRequest() string {
	println("Iniciando request para ViaCEP...")
	req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	return "Resposta de ViaCEP:" + string(res)
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- BrasilAPIRequest()
	}()

	go func() {
		ch2 <- ViaCEPRequest()
	}()

	select {
	case msg := <-ch1:
		println(msg)
	case msg := <-ch2:
		println(msg)
	case <-time.After(time.Second):
		println("TIMEOUT: Limite de tempo de resposta das APIs excedido.")
	}
}
