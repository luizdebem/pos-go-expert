package main

import (
	"fmt"

	"github.com/luizdebem/goexpoert/8-packaging/1-introducao/math"
)

func main() {
	fmt.Println("Hello, World!")
	a := math.Math{
		A: 10,
		B: 20,
	}
	fmt.Println(a.Add())
}
