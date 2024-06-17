package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if err == context.DeadlineExceeded {
			fmt.Println("A requisição de cotação do dólar teve tempo excedido.")
			panic(err)
		} else {
			panic(err)
		}

	}
	defer res.Body.Close()

	f, err := os.OpenFile("./cotacao.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.Copy(f, res.Body)
	if _, err := fmt.Fprintln(f, ""); err != nil {
		panic(err)
	}
}
