package math

var X string = "Hello, World!"

type math struct {
	a int
	b int
	C int
}

func NewMath(a int, b int) math {
	return math{
		a: a,
		b: b,
	}
}

func (m math) Add() int {
	return m.a + m.b
}

func (m math) Sub() int {
	return m.a - m.b
}
