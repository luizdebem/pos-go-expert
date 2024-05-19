package main

const a = "Hello, World!"

// escopo global
var (
	b bool    // inicia como false (default value)
	c int     = 10
	d string  = "Josias"
	e float64 = 1.2
)

func main() {
	// escopo local
	a := "X" // short declaration com := e inferência (por causa da inicialização)
	a = "Y"  // na reatribuição não é :=, é =.
	println(a)
}

func x() {

}
