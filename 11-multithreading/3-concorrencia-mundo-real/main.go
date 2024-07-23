package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var number uint64 = 0

func main() {

	// m := sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		atomic.AddUint64(&number, 1) // Operação atômica: resolve problemas de race condition. Não precisa ficar fazendo o mutex manual ali.
		// m.Unlock()
		w.Write([]byte(fmt.Sprintf("Você é o visitante número %d 👍\n", number)))
	})
	http.ListenAndServe(":3000", nil)
}

// Verificar race conditions:
// go run -race main.go
