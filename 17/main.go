package main

// GENERICS

type MyNumber int

type Number interface {
	~int | ~float64 // o ~ "libera" usar o MyNumber acima, que é um int por baixo dos panos
}

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

// duas funções iguais pro mesmo tipo?!

// Usando generics dá pra usar só uma.
func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

func Compara[T Number](a T, b T) bool {
	if a > b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Josias": 1000, "Joaquim": 2000, "Jorge": 3000}
	m2 := map[string]float64{"Josias": 1000.0, "Joaquim": 2000.0, "Jorge": 3000.0}
	m3 := map[string]MyNumber{"Josias": 1000, "Joaquim": 2000, "Jorge": 3000}
	println(SomaInteiro(m))
	println(SomaFloat(m2))
	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3)) // ~
	println(Compara(1, 2))
}
