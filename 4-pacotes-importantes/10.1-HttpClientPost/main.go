package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}

	jsonVar := bytes.NewBuffer([]byte(`{"nome": "Josias", "idade": 25}`)) // Importante bufferizar

	res, err := c.Post("https://google.com/", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	io.CopyBuffer(os.Stdout, res.Body, nil)
}
