package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"` // tags para lidar com json no marshal/unmarshal
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{
		Numero: 1,
		Saldo:  100,
	}

	res, err := json.Marshal(conta) // Marshal = transformar em json.
	if err != nil {
		panic(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta) // sem guardar nenhuma variável, só cuspir pra algum writer
	if err != nil {
		panic(err)
	}

	jsonPuro := []byte(`{"n": 2, "s": 200}`)
	var conta2 Conta
	err = json.Unmarshal(jsonPuro, &conta2) // Unmarshal = de json para struct
	if err != nil {
		panic(err)
	}
	println(conta2.Numero, conta2.Saldo)
}
