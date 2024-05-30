package main

import "net/http"

func main() {
	// Meu mux personalizado
	mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World!"))
	// })
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "My Blog"})
	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}

// Diferença entre mux padrão e personalizado:
// Mux padrão: não se preocupa em criar um mux novo. Mas não tem muito controle sobre. Não dá pra criar diversos servidores (portas diferentes).
// Mux personalizado: tem mais controle, personalização, attachar...
