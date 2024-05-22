package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://google.com/")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close() // o defer vai executar por último
	// por isso já usa o close agora pra não esquecer, e no final da execução vai rodar o close

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}
