package matematica

func Soma[T int | float64](a, b T) T { // S maiúsculo = exportado
	return a + b
}

var A int = 10

type Carro struct {
	Marca string
}
